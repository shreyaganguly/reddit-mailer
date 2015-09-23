package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/mail"
	"net/smtp"
)

type FeedMailer struct {
	Sender     mail.Address
	Server     string
	PortNumber int
	Auth       smtp.Auth
}

const tpl = `
<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8">
	<title>Go News</title>
</head>
<body>
	<ul>
	{{range .}}
		<li><a href="{{ .URL }}">{{ .Title }}</a></li>
	{{end}}
	</ul>
</body>
</html>`

func NewMailer(server string, port int, sender mail.Address, userName, secret string) FeedMailer {
	return FeedMailer{
		sender,
		server,
		port,
		smtp.PlainAuth("", userName, secret, server),
	}
}

var header map[string]string

func toString(m map[string]string) string {
	var concat string
	for k, v := range m {
		concat += fmt.Sprintf("%s:%s\r\n", k, v)
	}
	return concat
}

func (m *FeedMailer) MakeHeader(recipient mail.Address) string {
	header = make(map[string]string)
	header["MIME-Version"] = "1.0"
	header["From"] = m.Sender.String()
	header["To"] = recipient.String()
	header["Subject"] = "Go lang updates"
	header["Content-type"] = "text/html"
	return toString(header)
}

func (m *FeedMailer) MailBody(feeds []Feed, receipient mail.Address) []byte {

	t, err := template.New("webpage").Parse(tpl)

	var buff bytes.Buffer
	err = t.Execute(&buff, feeds)
	if err != nil {
		return make([]byte, 0)
	}

	return []byte(m.MakeHeader(receipient) + buff.String())
}

func (m *FeedMailer) ServerName() string {
	return fmt.Sprintf("%s:%d", m.Server, m.PortNumber)
}

func (m *FeedMailer) SendTo(recipient mail.Address, feeds []Feed) {
	smtp.SendMail(m.ServerName(), m.Auth, m.Sender.Address, []string{recipient.Address}, m.MailBody(feeds, recipient))
}

func (m *FeedMailer) Send(receipients []mail.Address, feeds []Feed) {
	for i := 0; i < len(receipients); i++ {
		m.SendTo(receipients[i], feeds)
	}
}
