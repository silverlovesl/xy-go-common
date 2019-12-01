package utils

import (
	"fmt"
	"os/exec"

	"bitbucket.org/beecomb-grid/grid-ai/gridai_common/go/errors"
)

// SendMail メール送信
func SendMail(from, to, title, body string) error {
	return sendMailMock(from, to, title, body)
}

func sendMailMock(from, to, title, body string) error {
	sendmail := exec.Command("/usr/sbin/sendmail", "-f", from, to)
	fmt.Println(body)
	stdin, err := sendmail.StdinPipe()
	if err != nil {
		return &errors.ErrSendMailFailed
	}
	_, err = sendmail.StdoutPipe()
	if err != nil {
		return &errors.ErrSendMailFailed
	}
	err = sendmail.Start()
	if err != nil {
		return &errors.ErrSendMailFailed
	}

	_, err = stdin.Write([]byte(body))
	if err != nil {
		return &errors.ErrSendMailFailed
	}

	err = stdin.Close()
	if err != nil {
		return &errors.ErrSendMailFailed
	}

	err = sendmail.Wait()
	if err != nil {
		return &errors.ErrSendMailFailed
	}
	return nil
}
