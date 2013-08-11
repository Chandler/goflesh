package utils

import (
	"encoding/base64"
	"fmt"
	"github.com/robfig/revel"
	"net/mail"
	"net/smtp"
	"os"
	"strings"
)

/*
Return a bcrypt hash for a pre-salted password string
*/
func SendEmail(
	subject string,
	body string,
	from_name string,
	from_address string,
	to_name string,
	to_address string,
) error {
	isDisabled := revel.Config.BoolDefault("noemail", false)
	if isDisabled {
		return nil
	}

	smtp_host, found := revel.Config.String("smtp.host")
	if !found {
		revel.ERROR.Fatal("No SMTP host configured")
	}
	smtp_port, found := revel.Config.String("smtp.port")
	if !found {
		revel.ERROR.Fatal("No SMTP port configured")
	}
	smtp_username, found := revel.Config.String("smtp.username")
	if !found {
		revel.ERROR.Fatal("No SMTP username configured")
	}
	smtp_key := os.Getenv("FLESH_MANDRILL_KEY")
	if smtp_key == "" {
		revel.ERROR.Fatal("OS ENV var FLESH_MANDRILL_KEY not set!")
	}

	auth := smtp.PlainAuth(
		"",
		smtp_username,
		smtp_key,
		smtp_host,
	)

	from := mail.Address{from_name, from_address}
	to := mail.Address{to_name, to_address}
	to_actual := to // the address to actually send to, not just put in the header

	isDev := revel.Config.BoolDefault("mode.dev", true)
	if isDev {
		overrideEmail := os.Getenv("FLESH_EMAIL_OVERRIDE")
		if overrideEmail == "" {
			overrideEmail = "fleshuser@mailinator.com"
		}
		revel.WARN.Print(isDev, overrideEmail)
		to_actual = mail.Address{to_name, overrideEmail}
	}

	header := make(map[string]string)
	header["From"] = from.String()
	header["To"] = to.String()
	header["Subject"] = encodeRFC2047(subject)
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/html; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(body))

	return smtp.SendMail(
		smtp_host+":"+smtp_port,
		auth,
		from.String(),
		[]string{to_actual.String()},
		[]byte(message),
	)
}

func encodeRFC2047(String string) string {
	// use mail's rfc2047 to encode any string
	addr := mail.Address{String, ""}
	return strings.Trim(addr.String(), " <>")
}
