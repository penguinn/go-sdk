package email

import (
	"crypto/tls"
	"encoding/base64"
	"errors"
	"fmt"
	"net"
	"net/smtp"
	"strings"

	"github.com/penguinn/go-sdk/log"
)

type Email struct {
	From      string   // 发件人邮箱，如：songjiayu02@baidu.com
	FromAlias string   // 发件人别名（选填），如：宋家宇
	Password  string   // 发件人邮箱密码
	AuthType  string   // smtp认证方式，默认PLAIN
	Addr      string   // 邮箱地址(一般端口为25)，如：email.baidu.com:25
	To        []string // 收件人邮箱
	Subject   string   // 主题
	Body      string   // 内容
}

func Send(e *Email) error {
	hp := strings.Split(e.Addr, ":")
	if len(hp) < 1 {
		return errors.New("invalid host")
	}
	if e.AuthType == "" {
		e.AuthType = PlainAuthType
	}

	err := SendMail(e.Addr, e.From, e.FromAlias, e.Password, e.AuthType, e.To, e.Subject, []byte(e.Body))
	if err != nil {
		return err
	}

	return nil
}

type LoginAuth struct {
	username string
	password string
}

func NewLoginAuth(username, password string) smtp.Auth {
	return &LoginAuth{username, password}
}

func (a *LoginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte{}, nil
}

func (a *LoginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("unknown from server")
		}
	}
	return nil, nil
}

func SendMail(addr, from, fromAlias, password, authType string, to []string, subject string, msg []byte) error {
	c, err := smtp.Dial(addr)
	host, _, _ := net.SplitHostPort(addr)
	if err != nil {
		log.Error("call dial")
		return err
	}
	defer func(c *smtp.Client) {
		_ = c.Quit()
	}(c)

	if ok, _ := c.Extension("STARTTLS"); ok {
		config := &tls.Config{ServerName: host}
		if err = c.StartTLS(config); err != nil {
			log.Error("call start tls")
			return err
		}
	}

	if ok, authStr := c.Extension("AUTH"); ok {
		var a smtp.Auth
		if strings.Contains(authStr, "PLAIN") && authType == PlainAuthType {
			a = smtp.PlainAuth("", from, password, host)
		} else if strings.Contains(authStr, "LOGIN") && authType == LoginAuthType {
			a = NewLoginAuth(from, password)
		} else if strings.Contains(authStr, "") && authType == CramMd5AuthType {
			a = smtp.CRAMMD5Auth(from, password)
		} else {
			log.Errorf("dont support this authType: %s", authType)
			return fmt.Errorf("dont support this authType: %s", authType)
		}
		if err = c.Auth(a); err != nil {
			log.Error("check auth with err:", err)
			return err
		}
	}

	if err = c.Mail(from); err != nil {
		log.Error(err)
		return err
	}
	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			log.Error(err)
			return err
		}
	}
	w, err := c.Data()
	if err != nil {
		log.Error(err)
		return err
	}
	defer w.Close()

	header := make(map[string]string)
	header["Subject"] = subject
	if fromAlias == "" {
		header["From"] = from
	} else {
		header["From"] = fmt.Sprintf("%s<%s>", fromAlias, from)
	}
	header["To"] = strings.Join(to, ";")
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"
	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString(msg)
	_, err = w.Write([]byte(message))
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
