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
	var sql = "Select * From gva.t_users Where `account` = ? Or `mobile` = ? LIMIT 1 "
	err = This.db.Get(&userInfo, sql, condition.Username, condition.Username)

	return userInfo, err
}

func (This *Model) UpdateLastLoginTime(id int64) {
	var sql = "UPDATE gva.t_users set last_login_time = ? where id = ?"
	if _, err := This.db.Exec(sql, time.Now().Unix(), id); err != nil {
		panic(errors.Wrap(err, "更新最后登录时间失败"))
	}
}

func (This *Model) GetPageUserInfo(userList *[]TUsers, total *int64, gidList []string, limit int64, offset int64) {

	var sql = "SELECT COUNT(1) FROM gva.t_users WHERE gid IN (?)"
	sql, args, _ := sqlx.In(sql, gidList)

	if err := This.db.Get(total, sql, args...); err != nil {
		panic(errors.Wrap(err, "列表条数获取失败"))
	}

	sql = "SELECT * FROM gva.t_users WHERE gid IN (?) LIMIT ? OFFSET ?"
	sql, args, _ = sqlx.In(sql, gidList, limit, offset)

	if err := This.db.Select(userList, sql, args...); err != nil {
		panic(errors.Wrap(err, "用户列表获取失败"))
	}
}

func (This *Model) GetPageUserInfoByKey(userList *[]TUsers, total *int64, gidList []string, key string, limit, offset int64) {

	key = fmt.Sprintf("%%%s%%", strings.TrimSpace(key))

	var sql = "SELECT COUNT(1) FROM gva.t_users WHERE gid IN (?) AND (account LIKE ? OR mobile LIKE ?)"
	sql, args, _ := sqlx.In(sql, gidList, key, key)
	if err := This.db.Get(total, sql, args...); err != nil {
		panic(errors.Wrap(err, "列表条数获取失败"))
	}

	sql = "SELECT * FROM gva.t_users WHERE gid IN (?) AND (account LIKE ? OR mobile LIKE ?) LIMIT ? OFFSET ?"
	sql, args, _ = sqlx.In(sql, gidList, key, key, limit, offset)

	if err := This.db.Select(userList, sql, args...); err != nil {
		panic(errors.Wrap(err, "用户列表获取失败"))
	}
}

func (This *Model) UpdateUserInfo(updateInfo map[string]any) {
	sql := "UPDATE gva.t_users set gid = :gid ,`name` = :name ,mobile = :mobile WHERE id = :id"
	if _, err := This.db.NamedExec(sql, updateInfo); err != nil {
		panic(errors.Wrap(err, "更新用户信息失败"))
	}
}

func (This *Model) GetUserListByGid(gid int64) (tUsers []TUsers, total int64) {
	sql := "SELECT COUNT(1) FROM gva.t_users WHERE gid = ?"
	if err := This.db.Get(&total, sql, gid); err != nil {
		panic(errors.Wrap(err, "列表条数获取失败"))
	}

	sql = "SELECT * FROM gva.t_users WHERE gid = ?"
	if err := This.db.Select(&tUsers, sql, gid); err != nil {
		panic(errors.Wrap(err, "用户列表获取失败"))
	}

	return tUsers, total
}

func (This *Model) UpdateUserGid(id int64, gid int64) {
	sql := "UPDATE gva.t_users set gid = ? WHERE id = ?"
	if _, err := This.db.Exec(sql, gid, id); err != nil {
		panic(errors.Wrap(err, "更新用户分组Id失败"))
	}
}

func (This *Model) DeleteUser(id int64) {
	sql := "DELETE FROM gva.t_users WHERE id = ?"

	if _, err := This.db.Exec(sql, id); err != nil {
		panic(errors.Wrap(err, "删除用户失败"))
	}
}

func (This *Model) UpdateUserPwd(id int64, pwd string) {
	sql := "UPDATE gva.t_users set password = ? WHERE id = ?"
	if _, err := This.db.Exec(sql, pwd, id); err != nil {
		panic(errors.Wrap(err, "更新密码失败"))
	}
}

func (This *Model) GetUserCountByAccount(account string) (total int64) {
	if err := This.db.Get(&total, "SELECT COUNT(1) from gva.t_users where account = ?", account); err != nil {
		return 0
	}

	return total
}

func (This *Model) GetUserCountByMobile(mobile string) (total int64) {
	if err := This.db.Get(&total, "SELECT COUNT(1) from gva.t_users where mobile = ?", mobile); err != nil {
		return 0
	}

	return total
}

func (This *Model) InsertNewUser(user TUsers) {
	var sql = `INSERT INTO gva.t_users (account, password, gid, name, head_pic, mobile, create_time, update_time, last_login_time,
                         create_by)
VALUES (:account, :password, :gid, :name, :head_pic, :mobile, :create_time, :update_time, :last_login_time, :create_by)`

	if _, err := This.db.NamedExec(sql, user); err != nil {
		panic(errors.Wrap(err, "新增用户失败"))
	}
}

func (This *Model) GetUserPwd(id int64) (pwd string) {

	if err := This.db.Get(&pwd, "SELECT password FROM gva.t_users WHERE id = ?", id); err != nil {
		return ""
	}

	return pwd
}

func (This *Model) UpdateUserHeadPic(id int64, pic int64) {
	sql := "UPDATE gva.t_users set head_pic = ? WHERE id = ?"

	if _, err := This.db.Exec(sql, pic, id); err != nil {
		panic(errors.Wrap(err, "更新头像失败"))
	}
}
