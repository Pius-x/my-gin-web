package user

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/my-gin-web/global"
	systemReq "github.com/my-gin-web/model/system/request"
	"github.com/my-gin-web/utils/dbInstance"
	"go.uber.org/zap"
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

func (This *Model) GetOneUserInfo(userInfo *TUsers, condition systemReq.Login) error {
	var sql = "Select * From gva.t_users Where `account` = ? Or `mobile` = ? LIMIT 1 "
	return This.db.Get(userInfo, sql, condition.Username, condition.Username)
}

func (This *Model) UpdateLastLoginTime(account string) error {
	var sql = "UPDATE gva.t_users set last_login_time = ? where account = ?"
	_, err := This.db.Exec(sql, time.Now().Unix(), account)

	return err
}

func (This *Model) GetPageUserInfo(userList *[]TUsers, total *int64, gidList []string, limit int64, offset int64) (err error) {
	var sql = "SELECT COUNT(1) FROM gva.t_users WHERE gid IN (?)"
	sql, args, _ := sqlx.In(sql, gidList)

	err = This.db.Get(total, sql, args...)
	if err != nil {
		global.ZapLog.Error("", zap.Error(err))
		return err
	}

	sql = "SELECT * FROM gva.t_users WHERE gid IN (?) LIMIT ? OFFSET ?"
	sql, args, _ = sqlx.In(sql, gidList, limit, offset)

	err = This.db.Select(userList, sql, args...)

	return err
}

func (This *Model) GetPageUserInfoByKey(userList *[]TUsers, total *int64, gidList []string, key string, limit, offset int64) (err error) {

	key = fmt.Sprintf("%%%s%%", strings.TrimSpace(key))

	var sql = "SELECT COUNT(1) FROM gva.t_users WHERE gid IN (?) AND (account LIKE ? OR mobile LIKE ?)"
	sql, args, _ := sqlx.In(sql, gidList, key, key)
	err = This.db.Get(total, sql, args...)

	sql = "SELECT * FROM gva.t_users WHERE gid IN (?) AND (account LIKE ? OR mobile LIKE ?) LIMIT ? OFFSET ?"
	sql, args, _ = sqlx.In(sql, gidList, key, key, limit, offset)
	err = This.db.Select(userList, sql, args...)

	return err
}

func (This *Model) UpdateUserInfo(updateInfo map[string]interface{}) (err error) {
	sql := "UPDATE gva.t_users set gid = :gid ,`name` = :name ,mobile = :mobile WHERE id = :id"
	_, err = This.db.NamedExec(sql, updateInfo)
	return err
}

func (This *Model) GetUserListByGid(gid int64) (tUsers []TUsers, total int64, err error) {
	sql := "SELECT COUNT(1) FROM gva.t_users WHERE gid = ?"
	if err = This.db.Get(&total, sql, gid); err != nil {
		global.ZapLog.Error("err", zap.Error(err))
		return nil, 0, err
	}

	sql = "SELECT * FROM gva.t_users WHERE gid = ?"
	if err = This.db.Select(&tUsers, sql, gid); err != nil {
		global.ZapLog.Error("err", zap.Error(err))
		return nil, 0, err
	}

	return tUsers, total, err
}

func (This *Model) UpdateUserGid(id int64, gid int64) (err error) {
	sql := "UPDATE gva.t_users set gid = ? WHERE id = ?"
	_, err = This.db.Exec(sql, gid, id)
	return err
}

func (This *Model) DeleteUser(id int64) (err error) {
	sql := "DELETE FROM gva.t_users WHERE id = ?"
	_, err = This.db.Exec(sql, id)
	return err
}

func (This *Model) UpdateUserPwd(id int64, pwd string) (err error) {
	sql := "UPDATE gva.t_users set password = ? WHERE id = ?"
	_, err = This.db.Exec(sql, pwd, id)
	return err
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

func (This *Model) InsertNewUser(user TUsers) (err error) {
	var sql = `INSERT INTO gva.t_users (account, password, gid, name, head_pic, mobile, create_time, update_time, last_login_time,
                         create_by)
VALUES (:account, :password, :gid, :name, :head_pic, :mobile, :create_time, :update_time, :last_login_time, :create_by)`
	_, err = This.db.NamedExec(sql, user)

	return err
}

func (This *Model) GetUserPwd(id int64) (pwd string) {

	if err := This.db.Get(&pwd, "SELECT password FROM gva.t_users WHERE id = ?", id); err != nil {
		return ""
	}

	return pwd
}

func (This *Model) UpdateUserHeadPic(id int64, pic int64) (err error) {
	sql := "UPDATE gva.t_users set head_pic = ? WHERE id = ?"
	_, err = This.db.Exec(sql, pic, id)
	return err
}
