package service

import (
	"encoding/json"
	"errors"
	"github.com/my-gin-web/global"
	"github.com/my-gin-web/model/group"
	"go.uber.org/zap"
	"strconv"
	"time"
)

type GroupService struct{}

// GetGroupInfoList 获取分组列表
func (This *GroupService) GetGroupInfoList(gid int64) (err error, list []group.GroupInfo) {
	var groupSlice []group.GroupInfo
	var groupList []group.GroupInfo

	groupList, err = groupModel.GetAllGroups()
	if err != nil {
		global.ZapLog.Error("get all group list err", zap.Error(err))
		return err, nil
	}

	for _, oneGroup := range groupList {
		if gid == 0 {
			if oneGroup.ParentGid == 0 {
				groupSlice = append(groupSlice, oneGroup)
			}
		} else {
			if oneGroup.Gid == gid {
				groupSlice = append(groupSlice, oneGroup)
			}
		}
	}

	if len(groupSlice) > 0 {
		for k := range groupSlice {
			err = This.findChildrenGroup(groupList, &groupSlice[k])
		}
	}
	return err, groupSlice
}

// findChildrenGroup 找寻子分组
func (This *GroupService) findChildrenGroup(allGroups []group.GroupInfo, authority *group.GroupInfo) (err error) {
	for _, oneGroup := range allGroups {
		if oneGroup.ParentGid == authority.Gid {
			authority.Children = append(authority.Children, oneGroup)
		}
	}
	if len(authority.Children) > 0 {
		for k := range authority.Children {
			err = This.findChildrenGroup(allGroups, &authority.Children[k])
		}
	}
	return err
}

// CreateGroup 新增分组
func (This *GroupService) CreateGroup(auth group.TGroups) (err error) {
	var curTime = time.Now().Unix()
	bytes, _ := json.Marshal([]group.RouteInfo{{Path: "/welcome", Readonly: 0}})

	auth.RouterList = string(bytes)
	auth.CreateTime = curTime
	auth.UpdateTime = curTime

	return groupModel.CreateNewGroup(auth)
}

// ParentGroupRouterListVerify 更新分组的权限 以及其子分组的权限
func (This *GroupService) ParentGroupRouterListVerify(auth group.GroupInfo) (err error) {
	parentGroupInfo, err := groupModel.GetOneGroupInfo(auth.ParentGid)
	if err != nil {
		return err
	}

	var parentRouterMap = map[string]int64{}
	for _, oneRouter := range parentGroupInfo.RouterList {
		parentRouterMap[oneRouter.Path] = oneRouter.Readonly
	}

	for _, routeInfo := range auth.RouterList {
		if readOnly, OK := parentRouterMap[routeInfo.Path]; !OK || readOnly > routeInfo.Readonly {
			return errors.New("权限不能超过其父分组的权限")
		}
	}

	return err
}

// UpdateGroupRouterList 更新分组的权限 以及其子分组的权限
func (This *GroupService) UpdateGroupRouterList(auth group.GroupInfo) (err error) {
	//更新此分组的权限
	if err = groupModel.UpdateGroupRouter(auth.Gid, auth.RouterList); err != nil {
		return err
	}

	//更新此分组的子分组权限
	err, groupList := This.GetGroupInfoList(auth.Gid)
	if err != nil {
		return err
	}
	if len(groupList) > 0 {
		var routerMap = map[string]int64{}
		for _, oneRouter := range auth.RouterList {
			routerMap[oneRouter.Path] = oneRouter.Readonly
		}
		for _, item := range groupList {
			err = This.UpdateChildrenGid(auth.Gid, item, routerMap)
		}
	}
	return err
}

// UpdateGroup 更新分组信息
func (This *GroupService) UpdateGroup(auth group.TGroups) (err error) {

	parentRouterList, err := groupModel.GetOneGroupRouterList(auth.ParentGid)
	if err != nil {
		return err
	}

	routerList, err := groupModel.GetOneGroupRouterList(auth.Gid)
	if err != nil {
		return err
	}

	var routerMap = map[string]int64{}
	for _, oneRouter := range routerList {
		routerMap[oneRouter.Path] = oneRouter.Readonly
	}

	marshal, err := json.Marshal(This.GetNewRouterByParentGid(parentRouterList, routerMap))
	if err != nil {
		return err
	}

	auth.RouterList = string(marshal)

	return groupModel.UpdateGroupInfo(auth)
}

// DeleteGroup 删除分组
func (This *GroupService) DeleteGroup(auth *group.TGroups) (err error) {
	if total := groupModel.GetGroupCountByGid(auth.Gid); total == 0 {
		return errors.New("该分组不存在")
	}

	if _, total, _ := userModel.GetUserListByGid(auth.Gid); total != 0 {
		return errors.New("此分组有用户正在使用，禁止删除")
	}

	if total := groupModel.GetGroupCountByParentGid(auth.Gid); total != 0 {
		return errors.New("此分组存在子角色不允许删除")
	}

	if err = groupModel.DeleteOneGroup(auth.Gid); err != nil {
		return err
	}

	return err
}

// UpdateChildrenGid 更新所属子分组的权限
func (This *GroupService) UpdateChildrenGid(gid int64, item group.GroupInfo, routerMap map[string]int64) (err error) {
	if item.Gid != gid {
		var routerList = This.GetNewRouterByParentGid(item.RouterList, routerMap)

		if err = groupModel.UpdateGroupRouter(item.Gid, routerList); err != nil {
			return err
		}
	}

	if len(item.Children) > 0 {
		for _, childrenInfo := range item.Children {
			err = This.UpdateChildrenGid(gid, childrenInfo, routerMap)
		}
	}
	return err
}

func (This GroupService) GetNewRouterByParentGid(parentRouter []group.RouteInfo, routerMap map[string]int64) (routerList []group.RouteInfo) {

	for _, routeInfo := range parentRouter {
		if readOnly, OK := routerMap[routeInfo.Path]; OK {
			if readOnly > routeInfo.Readonly {
				routeInfo.Readonly = readOnly
			}
			routerList = append(routerList, routeInfo)
		}
	}

	return routerList
}

// GetChildrenIdListByGid 获取子分组Id列表通过Gid
func (This *GroupService) GetChildrenIdListByGid(gid int64) (gidArr []string) {
	err, groupList := This.GetGroupInfoList(gid)
	if err != nil {
		return nil
	}

	var gidList []int64
	if len(groupList) > 0 {
		for _, item := range groupList {
			err = This.findChildrenGid(&gidList, item)
		}
	}

	if gid == 0 {
		gidList = append(gidList, 0)
	}

	for i := 0; i < len(gidList); i++ {
		gidArr = append(gidArr, strconv.Itoa(int(gidList[i])))
	}

	return gidArr
}

// findChildrenGid  查询子分组ID
func (This *GroupService) findChildrenGid(gidList *[]int64, item group.GroupInfo) (err error) {
	*gidList = append(*gidList, item.Gid)
	if len(item.Children) > 0 {
		for _, item2 := range item.Children {
			err = This.findChildrenGid(gidList, item2)
		}
	}
	return err
}
