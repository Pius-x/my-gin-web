package utils

var (
	IdVerify                 = Rules{"Id": {NotEmpty()}}
	LoginVerify              = Rules{"Username": {NotEmpty()}, "Password": {NotEmpty()}}
	UpdateUserInfoVerify     = Rules{"Id": {NotEmpty()}, "Name": {NotEmpty()}, "Mobile": {NotEmpty()}}
	MultiUpdateUserGidVerify = Rules{"UserGidList": {NotEmpty()}}
	CreateUserInfoVerify     = Rules{"Account": {NotEmpty()}, "Name": {NotEmpty()}, "Mobile": {NotEmpty()}}
	UserInfoVerify           = Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}}
	CreateGroupVerify        = Rules{"Gname": {NotEmpty()}}
	AuthorityVerify          = Rules{"Gid": {NotEmpty()}, "Gname": {NotEmpty()}}
	UpdateGroupRouterVerify  = Rules{"Gid": {NotEmpty()}, "RouterList": {NotEmpty()}}
	AuthorityIdVerify        = Rules{"Gid": {NotEmpty()}}
	ChangePasswordVerify     = Rules{"Id": {NotEmpty()}, "Password": {NotEmpty()}, "NewPassword": {NotEmpty()}}
	UpdateHeadPicVerify      = Rules{"Id": {NotEmpty()}, "HeadPic": {NotEmpty()}}
)
