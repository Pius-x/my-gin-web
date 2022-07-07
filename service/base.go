package service

import (
	"database/sql"
	"encoding/json"
	"github.com/my-gin-web/global"
	"github.com/my-gin-web/model/feishu"
	"github.com/my-gin-web/model/group"
	"github.com/my-gin-web/model/user"
	"github.com/my-gin-web/utils"
	"github.com/pkg/errors"
)

type BaseService struct{}

func (This *BaseService) Login(u user.ReqLogin) (err error, res user.ResLogin) {

	// 获取用户个人信息
	userInfo, err := userModel.GetOneUserInfo(u)
	if err != nil {
		return errors.New("账号不存在"), res
	}

	// 密码校验
	if ok := utils.BcryptCheck(u.Password, userInfo.Password); !ok {
		return errors.New("密码错误"), res
	}

	// 登录信息
	var loginInfo = user.LoginInfo{
		TUsers:     userInfo,
		RouterList: []group.RouteInfo{},
	}

	// 解析权限信息
	var groupInfoString = groupModel.GetRouterList(userInfo.Gid)
	if err = json.Unmarshal([]byte(groupInfoString), &loginInfo.RouterList); err != nil {
		panic(errors.Wrap(err, "权限列表解析失败"))
	}

	// 更新最后一次登录时间
	userModel.UpdateLastLoginTime(loginInfo.Id)

	// 登录以后签发jwt
	res = This.tokenNext(loginInfo)

	return err, res
}

// TokenNext 登录以后签发jwt
func (This *BaseService) tokenNext(loginInfo user.LoginInfo) (response user.ResLogin) {

	j := &utils.JWT{SigningKey: []byte(global.Config.JWT.SigningKey)} // 唯一签名
	claims := j.CreateClaims(user.BaseClaims{
		ID:       loginInfo.Id,
		Account:  loginInfo.Account,
		Name:     loginInfo.Name,
		Password: loginInfo.Password,
	})

	token, err := j.CreateToken(claims)
	if err != nil {
		panic(errors.Wrap(err, "获取token失败"))
	}

	return user.ResLogin{
		User:      loginInfo,
		Token:     token,
		ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
	}
}

// GetFsUserInfo 获取飞书用户信息
func (This BaseService) GetFsUserInfo(code string, uri string) (fsUserInfo feishu.FsUserInfo, err error) {
	fsUserInfo, err = fsUserModel.FsUserInfo(code, uri)
	if fsUserInfo.UnionId == "" || err != nil {
		return fsUserInfo, errors.New("获取飞书用户信息失败")
	}

	return fsUserInfo, err
}

// FsLogin 飞书登录
func (This *BaseService) FsLogin(code string) (res user.ResLogin, err error) {

	FsUserInfoRes, err := This.GetFsUserInfo(code, "/base/fsLogin")
	if err != nil {
		return res, errors.WithStack(err)
	}

	fsUserInfo, err := fsUserModel.GetFsUserInfoByUnionId(FsUserInfoRes.UnionId)
	if err != nil {
		return res, errors.New("此飞书账号未关联管理后台账号")
	}

	// 更新飞书的用户信息
	FsUserInfoRes.FsId = fsUserInfo.FsId
	if err = fsUserModel.UpdateFsUserInfo(FsUserInfoRes); err != nil {
		return res, errors.New("更新飞书用户信息失败")
	}

	userInfo, err := userModel.GetOneUserInfoById(fsUserInfo.CmsUserId)
	if err != nil {
		return res, errors.New("未找到账号信息")
	}

	// 登录信息
	var loginInfo = user.LoginInfo{
		TUsers:     userInfo,
		RouterList: []group.RouteInfo{},
	}

	// 解析权限信息
	var groupInfoString = groupModel.GetRouterList(userInfo.Gid)
	if err = json.Unmarshal([]byte(groupInfoString), &loginInfo.RouterList); err != nil {
		return res, errors.New("权限列表解析失败")
	}

	// 更新最后一次登录时间
	userModel.UpdateLastLoginTime(loginInfo.Id)

	// 登录以后签发jwt
	res = This.tokenNext(loginInfo)

	return res, err
}

// BindFs 绑定飞书账号
func (This *BaseService) BindFs(code string, id int64) (fsUserInfo feishu.FsUserInfo, err error) {

	fsUserInfo, err = This.GetFsUserInfo(code, "/base/fsBind")
	if err != nil {
		return fsUserInfo, errors.WithStack(err)
	}

	fsUserInfo.CmsUserId = id
	_, err = fsUserModel.GetFsUserInfoByCmsUserId(id)

	// 找到记录，说明重复绑定
	if err == nil {
		return fsUserInfo, errors.New("请勿重复绑定")
	}

	// 其他错误
	if !errors.Is(err, sql.ErrNoRows) {
		return fsUserInfo, errors.WithStack(err)
	}

	// 未找到绑定记录，继续绑定逻辑
	if _, err = fsUserModel.GetFsUserInfoByUnionId(fsUserInfo.UnionId); err == nil {
		return fsUserInfo, errors.New("此飞书已绑定其他账号")
	}

	err = fsUserModel.InsertNewFsUser(fsUserInfo)
	if err != nil {
		return fsUserInfo, errors.WithStack(err)
	}

	// 更新用户绑定状态
	userModel.UpdateUserBindFs(fsUserInfo.CmsUserId, true)

	// 更新用户昵称
	userModel.UpdateUserName(fsUserInfo.CmsUserId, fsUserInfo.Name)

	return fsUserInfo, err
}

// UnbindFs 解绑飞书关联
func (This *UserService) UnbindFs(id int64) {

	// 删除飞书用户信息
	fsUserModel.DelFsUserByCmsUserId(id)

	// 重置用户绑定信息为0
	userModel.UpdateUserBindFs(id, false)
}

// BuildFsBindInfo 构建飞书绑定的Token信息
func (This *BaseService) BuildFsBindInfo(fsUserInfo feishu.FsUserInfo) (fsLogin user.FsBindToken) {

	return user.FsBindToken{
		BindFs:    true,
		FsHeadPic: fsUserInfo.Picture,
		FsName:    fsUserInfo.Name,
	}
}

// BuildFsLoginInfo 构建飞书登录的Token信息
func (This *BaseService) BuildFsLoginInfo(res user.ResLogin) (fsLogin user.FsLoginToken) {

	loginInfo := This.BuildLoginInfo(res)

	loginToken, err := json.Marshal(loginInfo)
	if err != nil {
		return fsLogin
	}

	return user.FsLoginToken{
		UserInfo: string(loginToken),
	}
}

// BuildLoginInfo 构建登录的Token信息
func (This *BaseService) BuildLoginInfo(res user.ResLogin) user.LoginInfoRes {

	resInfo := This.paddingLoginResInfo(res)
	fsUserInfo, err := fsUserModel.GetFsUserInfoByCmsUserId(resInfo.Id)
	if err == nil {
		resInfo.BindFs = true
		resInfo.FsHeadPic = fsUserInfo.Picture
		resInfo.FsName = fsUserInfo.Name
	}
	return resInfo
}

// paddingLoginResInfo 填充登录返回信息
func (This BaseService) paddingLoginResInfo(login user.ResLogin) user.LoginInfoRes {

	return user.LoginInfoRes{
		Id:            login.User.Id,
		Account:       login.User.Account,
		Gid:           login.User.Gid,
		Name:          login.User.Name,
		HeadPic:       login.User.HeadPic,
		LastLoginTime: login.User.LastLoginTime,
		RouterList:    login.User.RouterList,
		Token:         login.Token,
		ExpiresAt:     login.ExpiresAt,
	}
}
