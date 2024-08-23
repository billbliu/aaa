package utils

import (
	"errors"
	"fmt"

	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/jordan-wright/email"
)

// example: https://cloud.tencent.com/developer/article/1730821

const (
	SMTPHost     = "smtp.gmail.com"
	SMTPPort     = ":587"
	SMTPUsername = "billliu201919@gmail.com" // 邮箱账号
	SMTPPassword = "eoyq wxiw vtcv wodj"     // 邮箱授权码
	MaxClient    = 10
)

func SendEmail(appName string, receiver string, subject string, content string) error {
	e := &email.Email{
		From:    fmt.Sprintf("%s<%s>", appName, SMTPUsername),
		To:      []string{receiver},
		Subject: subject,
		HTML:    []byte(content),
	}
	global.GVA_LOG.Sugar().Info(e)
	if err := global.EmailPool.Send(e, 5*time.Second); err != nil {
		global.GVA_LOG.Error(err.Error())
		return errors.New("send email failed, please try again")
	}
	return nil
}
