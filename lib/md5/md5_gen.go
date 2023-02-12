package md5

import (
	"crypto/md5"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"io"
	"strings"
)

func genMd5(code string) string {
	md5Code := md5.New()
	_, _ = io.WriteString(md5Code, code)
	return hex.EncodeToString(md5Code.Sum(nil))
}

func encodePwd(rawPwd string) string {

	// Using custom options
	options := &password.Options{
		SaltLen:      16,
		Iterations:   100,
		KeyLen:       32,
		HashFunction: sha512.New,
	}
	salt, encodedPwd := password.Encode(rawPwd, options)

	return fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
}

func verify(rawPwd, encodedPwd, salt string) bool {
	// Using custom options
	options := &password.Options{
		SaltLen:      16,
		Iterations:   100,
		KeyLen:       32,
		HashFunction: sha512.New,
	}
	passwordInfo := strings.Split(encodedPwd, "$")
	
	return password.Verify(rawPwd, passwordInfo[2], passwordInfo[3], options)
}
