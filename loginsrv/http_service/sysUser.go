package http_service

import (
	"errors"
	"github.com/go-resty/resty/v2"
	"log"
	"loginsrv/models"
	"loginsrv/models/sysUser"
	"os"
)

type PagedUser struct {
	models.PagedResponse
	Data []sysUser.SysUser `json:"data"`
}

func getUserApiPrefix() string {
	prefix := os.Getenv("PREFIX_USER_SERVICE")
	if prefix == "" {
		prefix = "http://127.0.0.1:18081"
	}
	return prefix
}

func FindUser(userName string) (user sysUser.SysUser, err error) {
	client := resty.New().SetRetryCount(3)
	resp, err := client.R().
		SetQueryParams(map[string]string{
			"username": userName,
		}).
		SetResult(&PagedUser{}).
		Get(getUserApiPrefix() + "/users")
	if err != nil {
		log.Panicln("FindUser err: ", err.Error())
	}
	response := resp.Result().(*PagedUser)
	if response.Total == 1 {
		user = response.Data[0]
		return
	}
	err = errors.New("用户不存在")
	return
}
