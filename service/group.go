package service

import (
	"encoding/json"
	"github.com/my-gin-web/model/group"
	"github.com/pkg/errors"
	"strconv"
	"time"
)

type GroupService struct{}

// GetGroupInfoList 获取分组列表
func (This *GroupService) GetGroupInfoList(gid int64) (groupSlice []group.GroupInfo) {

	var groupList = groupModel.GetAllGroups()

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
			This.findChildrenGroup(groupList, &groupSlice[k])
		}
	}
	return groupSlice
}

// findChildrenGroup 找寻子分组
func (This *GroupService) findChildrenGroup(allGroups []group.GroupInfo, authority *group.GroupInfo) {

	for _, oneGroup := range allGroups {
		if oneGroup.ParentGid == authority.Gid {
			authority.Children = append(authority.Children, oneGroup)
		}
	}
	if len(authority.Children) > 0 {
		for k := range authority.Children {
			This.findChildrenGroup(allGroups, &authority.Children[k])
		}
	}
}

// CreateGroup 新增分组
func (This *GroupService) CreateGroup(auth group.TGroups) {

	var curTime = time.Now().Unix()
	bytes, _ := json.Marshal([]group.RouteInfo{{Path: "/welcome", Readonly: 0}})

	auth.RouterList = string(bytes)
	auth.CreateTime = curTime
	auth.UpdateTime = curTime

	groupModel.CreateNewGroup(auth)
}

// ParentGroupRouterListVerify 更新分组的权限 以及其子分组的权限
func (This *GroupService) ParentGroupRouterListVerify(auth group.GroupInfo) {

	parentGroupInfo := groupModel.GetOneGroupInfo(auth.ParentGid)

	var parentRouterMap = map[string]int64{}
	for _, oneRouter := range parentGroupInfo.RouterList {
		parentRouterMap[oneRouter.Path] = oneRouter.Readonly
	}

	for _, routeInfo := range auth.RouterList {
		if readOnly, OK := parentRouterMap[routeInfo.Path]; !OK || readOnly > routeInfo.Readonly {
			panic(errors.New("权限不能超过其父分组的权限"))
		}
	}
}

// UpdateGroupRouterList 更新分组的权限 以及其子分组的权限
func (This *GroupService) UpdateGroupRouterList(auth group.GroupInfo) {

	//更新此分组的权限
	groupModel.UpdateGroupRouter(auth.Gid, auth.RouterList)

	//更新此分组的子分组权限
	groupList := This.GetGroupInfoList(auth.Gid)

	if len(groupList) > 0 {
		var routerMap = map[string]int64{}
		for _, oneRouter := range auth.RouterList {
			routerMap[oneRouter.Path] = oneRouter.Readonly
		}
		for _, item := range groupList {
			This.UpdateChildrenGid(auth.Gid, item, routerMap)
		}
	}
}

// UpdateGroup 更新分组信息
func (This *GroupService) UpdateGroup(auth group.TGroups) {

	routerList := groupModel.GetOneGroupRouterList(auth.Gid)
	parentRouterList := groupModel.GetOneGroupRouterList(auth.ParentGid)

	var routerMap = map[string]int64{}
	for _, oneRouter := range routerList {
		routerMap[oneRouter.Path] = oneRouter.Readonly
	}

	marshal, err := json.Marshal(This.GetNewRouterByParentGid(parentRouterList, routerMap))
	if err != nil {
		panic(errors.WithStack(err))
	}

	auth.RouterList = string(marshal)

	groupModel.UpdateGroupInfo(auth)
}

// DeleteGroup 删除分组
func (This *GroupService) DeleteGroup(auth *group.TGroups) (err error) {

	if total := groupModel.GetGroupCountByGid(auth.Gid); total == 0 {
		return errors.New("该分组不存在")
	}

	if _, total := userModel.GetUserListByGid(auth.Gid); total != 0 {
		return errors.New("此分组有用户正在使用，禁止删除")
	}

	if total := groupModel.GetGroupCountByParentGid(auth.Gid); total != 0 {
		return errors.New("此分组存在子角色不允许删除")
	}

	// 删除分组
	groupModel.DeleteOneGroup(auth.Gid)

	return err
}

// UpdateChildrenGid 更新所属子分组的权限
func (This *GroupService) UpdateChildrenGid(gid int64, item group.GroupInfo, routerMap map[string]int64) {

	if item.Gid != gid {
		var routerList = This.GetNewRouterByParentGid(item.RouterList, routerMap)
		groupModel.UpdateGroupRouter(item.Gid, routerList)
	}

	if len(item.Children) > 0 {
		for _, childrenInfo := range item.Children {
			This.UpdateChildrenGid(gid, childrenInfo, routerMap)
		}
	}
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

	groupList := This.GetGroupInfoList(gid)

	var gidList []int64
	if len(groupList) > 0 {
		for _, item := range groupList {
			This.findChildrenGid(&gidList, item)
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
func (This *GroupService) findChildrenGid(gidList *[]int64, item group.GroupInfo) {

	*gidList = append(*gidList, item.Gid)
	if len(item.Children) <= 0 {
		return
	}

	for _, item2 := range item.Children {
		This.findChildrenGid(gidList, item2)
	}
}
