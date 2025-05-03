package initialize

import (
	"github.com/quannv/strongbody-api/global"
	"github.com/quannv/strongbody-api/internal/config"
	mail "github.com/quannv/strongbody-api/utility/gmail"
)

func InitGmail() {
	name := config.GetConfig().MailCfg.SmtpSenderName
	fromEmailAddress := config.GetConfig().MailCfg.SmtpAuthEmail
	fromEmailPassword := config.GetConfig().MailCfg.SmtpAuthPassword

	global.Gmail = mail.NewGmailSender(name, fromEmailAddress, fromEmailPassword)
}
