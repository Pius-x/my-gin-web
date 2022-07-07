package feishu

import (
	"bytes"
	"encoding/json"
	"github.com/my-gin-web/global"
	"github.com/pkg/errors"
	"io"
	"net/http"
)

func (This *Model) FsUserInfo(code string, uri string) (userInfo FsUserInfo, err error) {
	info := make(map[string]string)
	info["grant_type"] = "authorization_code"
	info["client_id"] = global.Config.FsLogin.AppID
	info["client_secret"] = global.Config.FsLogin.AppSecret
	info["code"] = code
	info["redirect_uri"] = global.Config.FsLogin.RedirectUri + uri
	bytesData, _ := json.Marshal(info)

	url := "https://passport.feishu.cn/suite/passport/oauth/token"
	reader := bytes.NewReader(bytesData)
	req, err := http.NewRequest("POST", url, reader)
	defer req.Body.Close()
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	defer resp.Body.Close()
	reqsBytes, err := io.ReadAll(resp.Body)
	var acReq AccessReq
	json.Unmarshal(reqsBytes, &acReq)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return This.GetUserInfo(acReq.AccessToken)
}

func (This *Model) GetUserInfo(accessToken string) (fsUserInfo FsUserInfo, err error) {
	url := "https://passport.feishu.cn/suite/passport/oauth/userinfo"
	req, err := http.NewRequest("GET", url, nil)
	client := http.Client{}
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	resp, err := client.Do(req)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	defer resp.Body.Close()
	reqsBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	if err = json.Unmarshal(reqsBytes, &fsUserInfo); err != nil {
		return fsUserInfo, errors.WithStack(err)
	}

	return fsUserInfo, err
}
