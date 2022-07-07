package config

type FsLogin struct {
	AppID       string `mapstructure:"appID" json:"appID" yaml:"appID"`                   // 飞书应用APPID
	AppSecret   string `mapstructure:"appSecret" json:"appSecret" yaml:"appSecret"`       // 飞书应用AppSecret
	RedirectUri string `mapstructure:"redirectUri" json:"redirectUri" yaml:"redirectUri"` // 回调链接
}
