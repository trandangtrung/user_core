package initialize

import (
	"github.com/quannv/strongbody-api/global"
	"github.com/quannv/strongbody-api/internal/config"
	mail "github.com/quannv/strongbody-api/utility/gmail"
	"github.com/quannv/strongbody-api/utility/template"
)

func InitGmailAndTemplate() {
	gmailConfig := config.GetConfig().MailConfig
	nameEmail := gmailConfig.NameEmail
	accountEmail := gmailConfig.AccountEmail
	passwordEmail := gmailConfig.PasswordEmail
	global.Gmail = mail.NewGmailSender(nameEmail, accountEmail, passwordEmail)
	global.Template = template.NewTemplate()
}
