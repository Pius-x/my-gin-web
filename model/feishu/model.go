package feishu

import (
	"github.com/jmoiron/sqlx"
	"github.com/my-gin-web/utils/dbInstance"
	"github.com/pkg/errors"
)

type Model struct {
	db *sqlx.DB
}

func NewModel() *Model {
	return &Model{
		db: dbInstance.SelectConn(),
	}
}

// GetFsUserInfoByUnionId 获取飞书用户信息通过unionId
func (This *Model) GetFsUserInfoByUnionId(unionId string) (fsUserInfo FsUserInfo, err error) {
	var sql = "Select * From crazy_cms.cms_fs_user Where union_id = ? LIMIT 1 "
	err = This.db.Get(&fsUserInfo, sql, unionId)

	return fsUserInfo, err
}

// GetFsUserInfoByCmsUserId 获取飞书用户信息通过Cms用户id
func (This *Model) GetFsUserInfoByCmsUserId(userId int64) (fsUserInfo FsUserInfo, err error) {
	var sql = "Select * From crazy_cms.cms_fs_user Where cms_user_id = ? LIMIT 1 "
	err = This.db.Get(&fsUserInfo, sql, userId)

	return fsUserInfo, err
}

// InsertNewFsUser 新增飞书用户信息
func (This *Model) InsertNewFsUser(fsUserInfo FsUserInfo) (err error) {
	var sql = `INSERT INTO crazy_cms.cms_fs_user (sub, name, picture, open_id, union_id, en_name,
                                   tenant_key, avatar_url, avatar_thumb, avatar_middle, avatar_big, email, user_id,
                                   mobile, cms_user_id)
VALUES (:sub, :name, :picture, :open_id, :union_id, :en_name, :tenant_key, :avatar_url, :avatar_thumb,
        :avatar_middle, :avatar_big, :email, :user_id, :mobile, :cms_user_id)`
	_, err = This.db.NamedExec(sql, fsUserInfo)

	return err
}

// UpdateFsUserInfo 更新飞书用户信息
func (This *Model) UpdateFsUserInfo(fsUserInfo FsUserInfo) (err error) {
	var sql = `UPDATE crazy_cms.cms_fs_user
SET sub=:sub,
    name=:name,
    picture=:picture,
    open_id=:open_id,
    union_id=:union_id,
    en_name=:en_name,
    tenant_key=:tenant_key,
    avatar_url=:avatar_url,
    avatar_thumb=:avatar_thumb,
    avatar_middle=:avatar_middle,
    avatar_big=:avatar_big,
    email=:email,
    user_id=:user_id,
    mobile=:mobile
where fs_id = :fs_id`
	_, err = This.db.NamedExec(sql, fsUserInfo)

	return err
}

// DelFsUserByCmsUserId 删除飞书用户信息通过Cms用户id
func (This *Model) DelFsUserByCmsUserId(id int64) {
	sql := "DELETE FROM crazy_cms.cms_fs_user WHERE cms_user_id = ?"

	if _, err := This.db.Exec(sql, id); err != nil {
		panic(errors.Wrap(err, "删除飞书用户信息失败"))
	}
}
