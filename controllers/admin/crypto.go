package admin

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
)

func Md5Pass(s string) string {
	md5Password := md5.New()
	io.WriteString(md5Password, s)
	buffer := bytes.NewBuffer(nil)
	fmt.Fprintf(buffer, "%x", md5Password.Sum(nil))
	newPass := buffer.String()

	return newPass
}
