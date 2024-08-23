package initialize

import (
	"net/smtp"

	"github.com/flipped-aurora/gin-vue-admin/server/global"

	"github.com/jordan-wright/email"
)

// InitEmailPool 创建一个邮件池，每秒最多发送 MaxClient 封邮件
func InitEmailPool() *email.Pool {
	pool, err := email.NewPool(global.GVA_CONFIG.EmailConfig.SmtpHost+global.GVA_CONFIG.EmailConfig.SmtpPort, global.GVA_CONFIG.EmailConfig.MaxClient,
		smtp.PlainAuth("", global.GVA_CONFIG.EmailConfig.SmtpEmail, global.GVA_CONFIG.EmailConfig.SmtpPassword, global.GVA_CONFIG.EmailConfig.SmtpHost))
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		panic(err)
	}
	global.GVA_LOG.Info("init email pool successful")
	return pool
}
