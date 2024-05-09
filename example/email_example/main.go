package main

import (
	"github.com/penguinn/go-sdk/email"
	"github.com/penguinn/go-sdk/log"
)

func main() {
	es := []*email.Email{
		{
			From:     "852198764@qq.com",
			Password: "qq code",
			Addr:     "smtp.qq.com:25",
			To:       []string{"songjiayu02@baidu.com"},
			Subject:  "主题",
			Body:     "内容",
		},
		{
			From:      "852198764@qq.com",
			FromAlias: "宋家宇",
			Password:  "qq code",
			Addr:      "smtp.qq.com:25",
			To:        []string{"songjiayu02@baidu.com"},
			Subject:   "主题",
			Body:      "内容",
		},
		{
			From:     "songjiayu02@baidu.com",
			Password: "密码",
			AuthType: email.LoginAuthType,
			Addr:     "email.baidu.com:25",
			To:       []string{"852198764@qq.com"},
			Subject:  "主题to8521",
			Body:     "内容to8521",
		},
	}

	for index, e := range es {
		err := email.Send(e)
		if err != nil {
			log.Errorf("index(%v) error: %s", index, err.Error())
		}
	}
}
