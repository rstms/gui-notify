//go:build !windows

package notify

import (
	"log"
	"os"
	"strings"
)

func Send(title, message string) error {
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
	prefix := ViperGetString("mail.subject_prefix")
	if prefix != "" {
		prefix += ": "
	}
	subject := prefix + title
	body := []byte(message)
	log.Printf("to=%s\n", to)
	log.Printf("from=%s\n", from)
	log.Printf("subject=%s\n", subject)
	log.Printf("body:\n%s\n", HexDump(body))
	err = sendmail.Send(to, from, subject, body)
	if err != nil {
		return Fatal(err)
	}
	return nil
}
