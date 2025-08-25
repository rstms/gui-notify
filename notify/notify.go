//go:build !windows

package notify

import (
	"log"
	"os"
	"strings"
)

func Send(message string) error {
	ViperSetDefault("id", ProgramName())
	ViperSetDefault("title", "notification")
	ViperSetDefault("mail.sender", "NOTIFY_DAEMON")
	ViperSetDefault("mail.subject_prefix", "Notification")

	hostname := ViperGetString("mail.hostname")
	port := ViperGetInt("mail.port")
	username := ViperGetString("mail.username")
	password := ViperGetString("mail.password")
	ca := ViperGetString("mail.ca")
	sendmail, err := NewSendmail(hostname, port, username, password, ca)
	if err != nil {
		return Fatal(err)
	}

	to := ViperGetString("mail.recipient")
	from := ViperGetString("mail.sender")
	if !strings.Contains(from, "@") {
		fqdn, err := os.Hostname()
		if err != nil {
			return Fatal(err)
		}
		from += "@" + fqdn
	}
	prefix := ViperGetString("id")
	if prefix != "" {
		prefix += ": "
	}
	subject := prefix + ViperGetString("title")
	body := []byte(message)
	params := make(map[string]any)
	params["to"] = to
	params["from"] = from
	params["subject"] = subject
	params["body"] = string(body)
	log.Printf("notification: %s\n", FormatJSON(params))
	if ViperGetBool("debug") {
		log.Println(HexDump(body))
	}
	err = sendmail.Send(to, from, subject, body)
	if err != nil {
		return Fatal(err)
	}
	return nil
}
