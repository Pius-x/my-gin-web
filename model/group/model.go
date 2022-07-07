package group

import (
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/my-gin-web/utils/dbInstance"
	"github.com/pkg/errors"
)

func NewModel() *Model {
	return &Model{
		db: dbInstance.SelectConn(),
	}
}

type Model struct {
	db *sqlx.DB
}

func (This *Model) GetRouterList(gid int64) string {
	var routerString string

	sql := "select router_list from crazy_cms.cms_group where gid = ?"
	if err := This.db.Get(&routerString, sql, gid); err != nil {
		return "[]"
	}

	return routerString
}

func (This *Model) GetAllGroups() (allGroups []GroupInfo) {
	var groups []TGroups
	if err := This.db.Select(&groups, "SELECT * from crazy_cms.cms_group"); err != nil {
		panic(errors.WithStack(err))
	}

	for i := 0; i < len(groups); i++ {
		oneGroup := groups[i]
		one := GroupInfo{
			Gid:        oneGroup.Gid,
			Gname:      oneGroup.Gname,
			ParentGid:  oneGroup.ParentGid,
			RouterList: nil,
			CreateTime: oneGroup.CreateTime,
			UpdateTime: oneGroup.UpdateTime,
			Children:   nil,
		}

		if err := json.Unmarshal([]byte(oneGroup.RouterList), &one.RouterList); err != nil {
			panic(errors.Wrap(err, "权限列表解析失败"))
		}

		allGroups = append(allGroups, one)
	}

	return allGroups
}

func (This *Model) GetGroupCountByGid(gid int64) (total int64) {
	if err := This.db.Get(&total, "SELECT COUNT(1) from crazy_cms.cms_group where gid = ?", gid); err != nil {
		return 0
	}

	return total
}

func (This *Model) GetGroupCountByParentGid(gid int64) (total int64) {
	if err := This.db.Get(&total, "SELECT COUNT(1) from crazy_cms.cms_group where parent_gid = ?", gid); err != nil {
		return 0
	}

	return total
}

func (This *Model) GetOneGroupInfo(gid int64) (groupInfo GroupInfo) {
	var tGroup TGroups

	if err := This.db.Get(&tGroup, "select * from crazy_cms.cms_group where gid = ?", gid); err != nil {
		panic(errors.Wrap(err, fmt.Sprintf("未找到分组信息，gid：%d", gid)))
	}

	groupInfo = GroupInfo{
		Gid:        tGroup.Gid,
		Gname:      tGroup.Gname,
		ParentGid:  tGroup.ParentGid,
		RouterList: nil,
		CreateTime: tGroup.CreateTime,
		UpdateTime: tGroup.UpdateTime,
		Children:   nil,
	}
	_ = json.Unmarshal([]byte(tGroup.RouterList), &groupInfo.RouterList)

	return groupInfo
}

func (This *Model) GetOneGroupRouterList(gid int64) (routerList []RouteInfo) {

	var routerInfo string
	if err := This.db.Get(&routerInfo, "select router_list from crazy_cms.cms_group where gid = ?", gid); err != nil {
		panic(errors.Wrap(err, fmt.Sprintf("权限列表为空，gid：%d", gid)))
	}

	if err := json.Unmarshal([]byte(routerInfo), &routerList); err != nil {
		panic(errors.WithStack(err))
	}

	return routerList
}

func (This *Model) UpdateGroupRouter(gid int64, routerList []RouteInfo) {

	routers, _ := json.Marshal(routerList)

	sql := "UPDATE crazy_cms.cms_group SET router_list = ? WHERE gid = ?"
	_, err := This.db.Exec(sql, routers, gid)
	if err != nil {
		panic(errors.Wrap(err, "更新分组路由权限失败"))
	}
}

func (This *Model) UpdateGroupInfo(auth TGroups) {

	sql := "UPDATE crazy_cms.cms_group SET gname = ?,parent_gid = ?,router_list = ? WHERE gid = ?"
	if _, err := This.db.Exec(sql, auth.Gname, auth.ParentGid, auth.RouterList, auth.Gid); err != nil {
		panic(errors.Wrap(err, "更新分组信息失败"))
	}
}

func (This *Model) CreateNewGroup(auth TGroups) {
	sql := `insert into crazy_cms.cms_group (gname, parent_gid, router_list, create_time, update_time)
values (:gname, :parent_gid, :router_list, :create_time, :update_time)`

	if _, err := This.db.NamedExec(sql, auth); err != nil {
		panic(errors.Wrap(err, "创建新分组失败"))
	}
}

func (This *Model) DeleteOneGroup(gid int64) {
	sql := `DELETE FROM crazy_cms.cms_group WHERE gid = ?`
	if _, err := This.db.Exec(sql, gid); err != nil {
		panic(errors.Wrap(err, "删除分组失败"))
	}
}
