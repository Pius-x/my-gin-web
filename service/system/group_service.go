package system

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"gorm.io/gorm"
)

type GroupService struct{}

// GetGroupInfoList 获取分组列表
func (This *GroupService) GetGroupInfoList(gid uint64) (err error, list []system.TGroups) {
	db := global.GVA_DB.Model(&system.TGroups{})
	var groups []system.TGroups
	var allGroups []system.TGroups
	if err = db.Find(&allGroups).Error; err != nil {
		return err, nil
	}

	for _, oneGroup := range allGroups {
		if gid == 0 {
			if oneGroup.ParentGid == 0 {
				groups = append(groups, oneGroup)
			}
		} else {
			if oneGroup.Gid == gid {
				groups = append(groups, oneGroup)
			}
		}
	}
	fmt.Printf("groups %+v \n", groups)

	if len(groups) > 0 {
		for k := range groups {
			err = This.findChildrenGroup(allGroups, &groups[k])
		}
	}
	return err, groups
}

// findChildrenGroup 找寻子分组
func (This *GroupService) findChildrenGroup(allGroups []system.TGroups, authority *system.TGroups) (err error) {
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
func (This *GroupService) CreateGroup(auth system.TGroups) (err error, authority system.TGroups) {
	auth.RouterList = append(auth.RouterList, system.RouteInfo{Path: "/welcome", Readonly: 0})

	err = global.GVA_DB.Omit("Gid").Create(&auth).Error
	return err, auth
}

// ParentGroupRouterListVerify 更新分组的权限 以及其子分组的权限
func (This *GroupService) ParentGroupRouterListVerify(auth system.TGroups) (err error) {
	var parentGroupInfo system.TGroups
	if err = global.GVA_DB.Where("gid = ?", auth.ParentGid).First(&parentGroupInfo).Error; err != nil {
		return err
	}
	var parentRouterMap = map[string]uint64{}
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
func (This *GroupService) UpdateGroupRouterList(auth system.TGroups) (err error) {
	//更新此分组的权限
	routers, _ := json.Marshal(auth.RouterList)
	updateFields := map[string]interface{}{
		"router_list": routers,
	}
	if err = global.GVA_DB.Where("gid = ?", auth.Gid).First(&system.TGroups{}).Updates(&updateFields).Error; err != nil {
		return err
	}

	//更新此分组的子分组权限
	err, groupList := This.GetGroupInfoList(auth.Gid)
	if err != nil {
		return err
	}
	if len(groupList) > 0 {
		var routerMap = map[string]uint64{}
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
func (This *GroupService) UpdateGroup(auth system.TGroups) (err error, authority system.TGroups) {
	updateFields := map[string]interface{}{
		"gname":      auth.Gname,
		"parent_gid": auth.ParentGid,
	}
	err = global.GVA_DB.Where("gid = ?", auth.Gid).First(&system.TGroups{}).Updates(&updateFields).Error
	return err, auth
}

// DeleteGroup 删除分组
func (This *GroupService) DeleteGroup(auth *system.TGroups) (err error) {
	if errors.Is(global.GVA_DB.Where("gid = ?", auth.Gid).First(&auth).Error, gorm.ErrRecordNotFound) {
		return errors.New("该角色不存在")
	}

	if !errors.Is(global.GVA_DB.Where("gid = ?", auth.Gid).First(&system.TUsers{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色有用户正在使用禁止删除")
	}
	if !errors.Is(global.GVA_DB.Where("parent_gid = ?", auth.Gid).First(&system.TGroups{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色存在子角色不允许删除")
	}

	err = global.GVA_DB.Delete(auth, "gid = ?", auth.Gid).Error
	if err != nil {
		return
	}

	return err
}

// UpdateChildrenGid 更新所属子分组的权限
func (This *GroupService) UpdateChildrenGid(gid uint64, item system.TGroups, routerMap map[string]uint64) (err error) {
	if item.Gid != gid {
		var newRouterList []system.RouteInfo
		for _, routeInfo := range item.RouterList {
			if readOnly, OK := routerMap[routeInfo.Path]; OK {
				if readOnly > routeInfo.Readonly {
					routeInfo.Readonly = readOnly
				}
				newRouterList = append(newRouterList, routeInfo)
			}
		}
		fmt.Printf("newRouterList: %+v \n\n", newRouterList)

		routers, _ := json.Marshal(newRouterList)
		updateFields := map[string]interface{}{
			"router_list": routers,
		}
		if err = global.GVA_DB.Where("gid = ?", item.Gid).First(&system.TGroups{}).Updates(&updateFields).Error; err != nil {
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

// GetChildrenIdListByGid 获取子分组Id列表通过Gid
func (This *GroupService) GetChildrenIdListByGid(gid uint64) (gidArr []uint64) {
	err, groupList := This.GetGroupInfoList(gid)
	if err != nil {
		return nil
	}

	var gidList []uint64
	if len(groupList) > 0 {
		for _, item := range groupList {
			err = This.findChildrenGid(&gidList, item)
		}
	}

	if gid == 0 {
		gidList = append(gidList, 0)
	}

	return gidList
}

// findChildrenGid  查询子分组ID
func (This *GroupService) findChildrenGid(gidList *[]uint64, item system.TGroups) (err error) {
	*gidList = append(*gidList, item.Gid)
	if len(item.Children) > 0 {
		for _, item2 := range item.Children {
			err = This.findChildrenGid(gidList, item2)
		}
	}
	return err
}
