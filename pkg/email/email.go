package email

import (
	"crypto/tls"
	"fmt"
	"gopkg.in/gomail.v2"
	"mark3/global"
)

func SendEmail(to []string, cc []string, bcc []string, subject string, body string, attach string) error {
	if global.GVA_CONFIG.Smtp.IsOpen != "yes" {
		return fmt.Errorf("邮件服务器未开启")
	}

	host := global.GVA_CONFIG.Smtp.Host
	port := global.GVA_CONFIG.Smtp.Port
	username := global.GVA_CONFIG.Smtp.Username
	password := global.GVA_CONFIG.Smtp.Password
	from := global.GVA_CONFIG.Smtp.From
	to = append(to, global.GVA_CONFIG.Smtp.From)
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	if len(cc) > 0 {
		//抄送
		m.SetHeader("Cc", cc...)
	}
	if len(bcc) > 0 {
		// 密送
		m.SetHeader("Bcc", bcc...)
	}
	if len(attach) > 0 {
		//添加附件
		m.Attach(attach)
	}

	d := gomail.NewDialer(host, port, username, password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		global.GVA_LOG.Error(err.Error())
		return err
	}
	return nil
}
