package group

import (
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/my-gin-web/utils/dbInstance"
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

	sql := "select router_list from gva.t_groups where gid = ?"
	if err := This.db.Get(&routerString, sql, gid); err != nil {
		return ""
	}

	return routerString
}

func (This *Model) GetAllGroups() (allGroups []GroupInfo, err error) {
	var groups []TGroups
	if err = This.db.Select(&groups, "SELECT * from gva.t_groups"); err != nil {
		return nil, err
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
		_ = json.Unmarshal([]byte(oneGroup.RouterList), &one.RouterList)

		allGroups = append(allGroups, one)
	}

	return allGroups, err
}

func (This *Model) GetGroupCountByGid(gid int64) (total int64) {
	if err := This.db.Get(&total, "SELECT COUNT(1) from gva.t_groups where gid = ?", gid); err != nil {
		fmt.Printf("err %+v \n", err)
		return 0
	}

	return total
}

func (This *Model) GetGroupCountByParentGid(gid int64) (total int64) {
	if err := This.db.Get(&total, "SELECT COUNT(1) from gva.t_groups where parent_gid = ?", gid); err != nil {
		return 0
	}

	return total
}

func (This *Model) GetOneGroupInfo(gid int64) (groupInfo GroupInfo, err error) {
	var tGroup TGroups

	if err = This.db.Get(&tGroup, "select * from gva.t_groups where gid = ?", gid); err != nil {
		return GroupInfo{}, err
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

	return groupInfo, err
}

func (This *Model) GetOneGroupRouterList(gid int64) (routerList []RouteInfo, err error) {

	var routerInfo string

	if err = This.db.Get(&routerInfo, "select router_list from gva.t_groups where gid = ?", gid); err != nil {
		return nil, err
	}

	_ = json.Unmarshal([]byte(routerInfo), &routerList)
	return routerList, err
}

func (This *Model) UpdateGroupRouter(gid int64, routerList []RouteInfo) (err error) {

	routers, _ := json.Marshal(routerList)

	sql := "UPDATE gva.t_groups SET router_list = ? WHERE gid = ?"
	_, err = This.db.Exec(sql, routers, gid)

	return err
}

func (This *Model) UpdateGroupInfo(auth TGroups) (err error) {
	sql := "UPDATE gva.t_groups SET gname = ?,parent_gid = ?,router_list = ? WHERE gid = ?"
	_, err = This.db.Exec(sql, auth.Gname, auth.ParentGid, auth.RouterList, auth.Gid)

	return err
}

func (This *Model) CreateNewGroup(auth TGroups) (err error) {
	sql := `insert into gva.t_groups (gname, parent_gid, router_list, create_time, update_time)
values (:gname, :parent_gid, :router_list, :create_time, :update_time)`

	_, err = This.db.NamedExec(sql, auth)
	return err
}

func (This *Model) DeleteOneGroup(gid int64) (err error) {
	sql := `DELETE FROM gva.t_groups WHERE gid = ?`
	_, err = This.db.Exec(sql, gid)
	return err
}
