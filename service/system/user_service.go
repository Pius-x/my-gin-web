package system

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	systemRes "github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"gorm.io/gorm"
	"strings"
	"time"
)

//@function: Register
//@description: 用户注册
//@param: u model.SysUser
//@return: err error, userInter model.SysUser

type UserService struct{}

func (This *UserService) CreateUserInfo(u system.TUsers) (err error, userInter system.TUsers) {
	var user system.TUsers

	// 判断用户名是否注册
	if !errors.Is(global.GVA_DB.Where("account = ?", u.Account).First(&user).Error, gorm.ErrRecordNotFound) {
		return errors.New("账户名已注册"), userInter
	}

	// 判断电话号码是否注册
	if !errors.Is(global.GVA_DB.Where("mobile = ?", u.Mobile).First(&user).Error, gorm.ErrRecordNotFound) {
		return errors.New("手机号码已注册"), userInter
	}

	// 密码hash加密 注册
	u.Password = utils.BcryptHash(u.Password)
	err = global.GVA_DB.Create(&u).Error
	return err, u
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetUserInfo
//@description: 更新用户信息
//@param: reqUser model.SysUser
//@return: err error, user model.SysUser

func (This *UserService) MultiUpdateUserGid(userInfo systemReq.MultiUpdateUserGid) (err error) {
	for _, item := range userInfo.UserGidList {
		updateFields := map[string]interface{}{
			"gid": item.Gid,
		}

		fmt.Printf("updateFields: %+v \n", updateFields)
		if err = global.GVA_DB.Where("id = ?", item.Id).First(&system.TUsers{}).Updates(&updateFields).Error; err != nil {
			fmt.Printf("err %+v \n", err)
			return errors.New("批量更新失败")
		}
	}

	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetUserInfo
//@description: 更新用户信息
//@param: reqUser model.SysUser
//@return: err error, user model.SysUser

func (This *UserService) UpdateUserInfo(updateInfo map[string]interface{}, id uint64) error {
	return global.GVA_DB.Model(&system.TUsers{}).Where("id = ?", id).Updates(&updateInfo).Error
}

//@function: ChangePassword
//@description: 修改用户密码
//@param: u *model.SysUser, newPassword string
//@return: err error, userInter *model.SysUser

func (This *UserService) ChangePassword(u *system.TUsers, newPassword string) (err error, userInter *system.TUsers) {
	var user system.TUsers
	err = global.GVA_DB.Where("account = ?", u.Account).First(&user).Error
	if err != nil {
		return err, nil
	}
	if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
		return errors.New("原密码错误"), nil
	}
	user.Password = utils.BcryptHash(newPassword)
	err = global.GVA_DB.Save(&user).Error
	return err, &user

}

//@function: UpdateHeadPic
//@description: 更新头像
//@param: u *model.SysUser, newPassword string
//@return: err error, userInter *model.SysUser

func (This *UserService) UpdateHeadPic(u *system.TUsers, newHeadPic uint64) (err error, userInter *system.TUsers) {
	var user system.TUsers
	err = global.GVA_DB.Where("account = ?", u.Account).First(&user).Error
	if err != nil {
		return err, nil
	}
	user.HeadPic = newHeadPic
	err = global.GVA_DB.Save(&user).Error
	return err, &user
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetUserInfoList
//@description: 分页获取数据
//@param: info request.PageInfo
//@return: err error, list interface{}, total int64

func (This *UserService) GetUserInfoList(info request.UserList) (err error, list []systemRes.UserListResponse, total int64) {
	limit := 10
	offset := 0
	if info.Keyword == "" {
		limit = int(info.PageSize)
		offset = int(info.PageSize * (info.Page - 1))
	}

	groupService := GroupService{}
	var gidList = groupService.GetChildrenIdListByGid(info.Gid)

	fmt.Printf("groups %+v \n", gidList)
	db := global.GVA_DB.Model(&system.TUsers{})
	var users []system.TUsers

	var userList []systemRes.UserListResponse

	if info.Keyword == "" {
		err = db.Where("gid in ?", gidList).Count(&total).Limit(limit).Offset(offset).Find(&users).Error
		if err != nil {
			return err, nil, 0
		}
	} else {
		key := fmt.Sprintf("%%%s%%", strings.TrimSpace(info.Keyword))
		err = db.Where("gid in ?", gidList).Where("account LIKE ? or mobile LIKE ?", key, key).
			Count(&total).Limit(limit).Offset(offset).Find(&users).Error
		if err != nil {
			return err, nil, 0
		}
	}

	for _, user := range users {
		var oneUserInfo systemRes.UserListResponse
		oneUserInfo.Id = user.Id
		oneUserInfo.Account = user.Account
		oneUserInfo.Gid = user.Gid
		oneUserInfo.Name = user.Name
		oneUserInfo.Mobile = user.Mobile
		oneUserInfo.CreateTime = time.Unix(int64(user.CreateTime), 0).Format("2006-01-02 15:04:05")
		oneUserInfo.CreateBy = user.CreateBy
		userList = append(userList, oneUserInfo)
	}

	return err, userList, total
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetUserInfoList
//@description: 分页获取数据
//@param: info request.PageInfo
//@return: err error, list interface{}, total int64

func (This *UserService) GetUserListByGid(info request.UserListByGid) (err error, list []systemRes.UserListResponse, total int64) {

	db := global.GVA_DB.Model(&system.TUsers{})
	var users []system.TUsers

	var userList []systemRes.UserListResponse

	err = db.Where("gid = ?", info.Gid).Count(&total).Find(&users).Error
	if err != nil {
		return err, nil, 0
	}

	for _, user := range users {
		var oneUserInfo systemRes.UserListResponse
		oneUserInfo.Id = user.Id
		oneUserInfo.Account = user.Account
		oneUserInfo.Gid = user.Gid
		oneUserInfo.Name = user.Name
		oneUserInfo.Mobile = user.Mobile
		oneUserInfo.CreateTime = time.Unix(int64(user.CreateTime), 0).Format("2006-01-02 15:04:05")
		oneUserInfo.CreateBy = user.CreateBy
		userList = append(userList, oneUserInfo)
	}

	return err, userList, total
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteUser
//@description: 删除用户
//@param: id float64
//@return: err error

func (This *UserService) DeleteUser(id uint64) (err error) {
	var user system.TUsers
	err = global.GVA_DB.Where("id = ?", id).Delete(&user).Error

	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: resetPassword
//@description: 修改用户密码
//@param: ID uint
//@return: err error

func (This *UserService) ResetPassword(ID uint64) (err error) {
	err = global.GVA_DB.Model(&system.TUsers{}).Where("id = ?", ID).Update("password", utils.BcryptHash("123456")).Error
	return err
}
