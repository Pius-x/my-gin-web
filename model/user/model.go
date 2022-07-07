package user

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/my-gin-web/utils/dbInstance"
	"github.com/pkg/errors"
	"strings"
	"time"
)

type Model struct {
	db *sqlx.DB
}

func NewModel() *Model {
	return &Model{
		db: dbInstance.SelectConn(),
	}
}

func (This *Model) GetOneUserInfo(condition ReqLogin) (userInfo TUsers, err error) {
	var sql = "Select * From crazy_cms.cms_user Where `account` = ? Or `mobile` = ? LIMIT 1 "
	err = This.db.Get(&userInfo, sql, condition.Username, condition.Username)

	return userInfo, err
}

func (This *Model) UpdateLastLoginTime(id int64) {
	var sql = "UPDATE crazy_cms.cms_user set last_login_time = ? where id = ?"
	if _, err := This.db.Exec(sql, time.Now().Unix(), id); err != nil {
		panic(errors.Wrap(err, "更新最后登录时间失败"))
	}
}

func (This *Model) GetPageUserInfo(userList *[]TUsers, total *int64, gidList []int64, limit int64, offset int64) {

	var sql = "SELECT COUNT(1) FROM crazy_cms.cms_user WHERE gid IN (?)"
	sql, args, _ := sqlx.In(sql, gidList)

	if err := This.db.Get(total, sql, args...); err != nil {
		panic(errors.Wrap(err, "列表条数获取失败"))
	}

	sql = "SELECT * FROM crazy_cms.cms_user WHERE gid IN (?) LIMIT ? OFFSET ?"
	sql, args, _ = sqlx.In(sql, gidList, limit, offset)

	if err := This.db.Select(userList, sql, args...); err != nil {
		panic(errors.Wrap(err, "用户列表获取失败"))
	}
}

func (This *Model) GetPageUserInfoByKey(userList *[]TUsers, total *int64, gidList []int64, key string, limit, offset int64) {

	key = fmt.Sprintf("%%%s%%", strings.TrimSpace(key))

	var sql = "SELECT COUNT(1) FROM crazy_cms.cms_user WHERE gid IN (?) AND (account LIKE ? OR mobile LIKE ?)"
	sql, args, _ := sqlx.In(sql, gidList, key, key)
	if err := This.db.Get(total, sql, args...); err != nil {
		panic(errors.Wrap(err, "列表条数获取失败"))
	}

	sql = "SELECT * FROM crazy_cms.cms_user WHERE gid IN (?) AND (account LIKE ? OR mobile LIKE ?) LIMIT ? OFFSET ?"
	sql, args, _ = sqlx.In(sql, gidList, key, key, limit, offset)

	if err := This.db.Select(userList, sql, args...); err != nil {
		panic(errors.Wrap(err, "用户列表获取失败"))
	}
}

func (This *Model) UpdateUserInfo(updateInfo map[string]any) {
	sql := "UPDATE crazy_cms.cms_user set gid = :gid ,`name` = :name ,mobile = :mobile WHERE id = :id"
	if _, err := This.db.NamedExec(sql, updateInfo); err != nil {
		panic(errors.Wrap(err, "更新用户信息失败"))
	}
}

func (This *Model) GetUserListByGid(gid int64) (tUsers []TUsers, total int64) {
	sql := "SELECT COUNT(1) FROM crazy_cms.cms_user WHERE gid = ?"
	if err := This.db.Get(&total, sql, gid); err != nil {
		panic(errors.Wrap(err, "列表条数获取失败"))
	}

	sql = "SELECT * FROM crazy_cms.cms_user WHERE gid = ?"
	if err := This.db.Select(&tUsers, sql, gid); err != nil {
		panic(errors.Wrap(err, "用户列表获取失败"))
	}

	return tUsers, total
}

func (This *Model) UpdateUserGid(id int64, gid int64) {
	sql := "UPDATE crazy_cms.cms_user set gid = ? WHERE id = ?"
	if _, err := This.db.Exec(sql, gid, id); err != nil {
		panic(errors.Wrap(err, "更新用户分组Id失败"))
	}
}

func (This *Model) DeleteUser(id int64) {
	sql := "DELETE FROM crazy_cms.cms_user WHERE id = ?"

	if _, err := This.db.Exec(sql, id); err != nil {
		panic(errors.Wrap(err, "删除用户失败"))
	}
}

func (This *Model) UpdateUserPwd(id int64, pwd string) {
	sql := "UPDATE crazy_cms.cms_user set password = ? WHERE id = ?"
	if _, err := This.db.Exec(sql, pwd, id); err != nil {
		panic(errors.Wrap(err, "更新密码失败"))
	}
}

func (This *Model) GetUserCountByAccount(account string) (total int64) {
	if err := This.db.Get(&total, "SELECT COUNT(1) from crazy_cms.cms_user where account = ?", account); err != nil {
		return 0
	}

	return total
}

func (This *Model) GetUserCountByMobile(mobile string) (total int64) {
	if err := This.db.Get(&total, "SELECT COUNT(1) from crazy_cms.cms_user where mobile = ?", mobile); err != nil {
		return 0
	}

	return total
}

func (This *Model) InsertNewUser(user TUsers) {
	var sql = `INSERT INTO crazy_cms.cms_user (account, password, gid, name, head_pic, mobile, create_time, update_time, last_login_time,
                         create_by)
VALUES (:account, :password, :gid, :name, :head_pic, :mobile, :create_time, :update_time, :last_login_time, :create_by)`

	if _, err := This.db.NamedExec(sql, user); err != nil {
		panic(errors.Wrap(err, "新增用户失败"))
	}
}

func (This *Model) GetUserPwd(id int64) (pwd string) {

	if err := This.db.Get(&pwd, "SELECT password FROM crazy_cms.cms_user WHERE id = ?", id); err != nil {
		return ""
	}

	return pwd
}

func (This *Model) UpdateUserHeadPic(id int64, pic int64) {
	sql := "UPDATE crazy_cms.cms_user set head_pic = ? WHERE id = ?"

	if _, err := This.db.Exec(sql, pic, id); err != nil {
		panic(errors.Wrap(err, "更新头像失败"))
	}
}

func (This *Model) GetOneUserInfoById(id int64) (userInfo TUsers, err error) {
	var sql = "Select * From crazy_cms.cms_user Where id = ? LIMIT 1 "
	err = This.db.Get(&userInfo, sql, id)

	return userInfo, err
}

func (This *Model) GetGidById(id int64) (gid int64, err error) {
	var sql = "Select gid From crazy_cms.cms_user Where id = ? LIMIT 1 "
	err = This.db.Get(&gid, sql, id)

	return gid, err
}

func (This *Model) UpdateUserBindFs(id int64, bindFs bool) {
	sql := "UPDATE crazy_cms.cms_user set bind_fs = ? WHERE id = ?"

	if _, err := This.db.Exec(sql, bindFs, id); err != nil {
		panic(errors.Wrap(err, "更新绑定信息失败"))
	}
}

func (This *Model) UpdateUserName(id int64, name string) {
	sql := "UPDATE crazy_cms.cms_user set name = ? WHERE id = ?"

	if _, err := This.db.Exec(sql, name, id); err != nil {
		panic(errors.Wrap(err, "更新用户昵称失败"))
	}
}
