package common

import (
	mail "github.com/xhit/go-simple-mail/v2"
	"log"
)

type SendEmailProvider interface {
	SendVerifyCode(code, mailAddress string, emailType int) error
}

type emailProvider struct {
	username string
	password string
}

func NewEmailProvider(username, password string) *emailProvider {
	return &emailProvider{username: username, password: password}
}

func (sep *emailProvider) SendVerifyCode(code, mailAddress string, emailType int) error {
	server := mail.NewSMTPClient()
	server.Host = "smtp.gmail.com"
	server.Port = 587
	server.Username = sep.username
	server.Password = sep.password
	server.Encryption = mail.EncryptionTLS

	smtpClient, err := server.Connect()
	if err != nil {
		log.Fatal(err)
	}

	// Create email
	email := mail.NewMSG()
	email.SetFrom("Instagram <" + sep.username + ">")
	email.AddTo(mailAddress)
	//email.AddCc("redryo1505@gmail.com")
	if emailType == 1 {
		SendMailVerify(email, code)
	}
	if emailType == 2 {
		SendMailForgotPassword(email, code)
	}
	//email.AddAttachment("super_cool_file.png")

	// Send email
	err = email.Send(smtpClient)
	if err != nil {
		return err
	}
	return nil
}
func SendMailVerify(email *mail.Email, code string) {
	email.SetSubject("Mã kích hoạt Email")

	email.SetBody(mail.TextHTML, code)
}
func SendMailForgotPassword(email *mail.Email, code string) {
	email.SetSubject("Xác thực email khôi phục mật khẩu")
	code = "http://localhost:80/v1/users/reset_password/?code=" + code
	email.SetBody(mail.TextHTML, code)
}
