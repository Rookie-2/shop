package main

import (
	"crypto/sha512"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
)

func main() {
	// Using the default options
	//salt, encodedPwd := password.Encode("generic password", nil)
	//check := password.Verify("generic password", salt, encodedPwd, nil)
	//fmt.Println(check) // true

	// Using custom options
	// 建议算法SHA512
	options := &password.Options{10, 100, 32, sha512.New}
	salt, encodedPwd := password.Encode("generic password", options)
	pwd := fmt.Sprintf("$pdkdf2-sha512$%s$%s", salt, encodedPwd)
	fmt.Println(salt)
	fmt.Println(encodedPwd)
	fmt.Println(pwd)
	check := password.Verify("generic password", salt, encodedPwd, options)
	fmt.Println(check) // true
}
