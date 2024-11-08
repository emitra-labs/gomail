package gomail_test

import (
	"testing"

	"github.com/emitra-labs/gomail"
)

func TestSend_Success(t *testing.T) {
	gomail.Open()

	err := gomail.Send(&gomail.Mail{
		From:    "foo@example.com",
		To:      "bar@example.com",
		Subject: "test",
		Body:    "<p>test</p>",
	})

	if err != nil {
		t.Error("error sending email:", err)
	}

	gomail.Close()
}
