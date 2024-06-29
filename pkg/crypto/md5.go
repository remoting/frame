// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// Package gmd5 provides useful API for MD5 encryption algorithms.
package crypto

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

// Encrypt encrypts any type of variable using MD5 algorithms.
// It uses gconv package to convert `v` to its bytes type.
func Md5Encrypt(data []byte) (encrypt string, err error) {
	return Md5EncryptBytes(data)
}

// MustEncrypt encrypts any type of variable using MD5 algorithms.
// It uses gconv package to convert `v` to its bytes type.
// It panics if any error occurs.
func Md5MustEncrypt(data []byte) string {
	result, err := Md5Encrypt(data)
	if err != nil {
		panic(err)
	}
	return result
}

// EncryptBytes encrypts `data` using MD5 algorithms.
func Md5EncryptBytes(data []byte) (encrypt string, err error) {
	h := md5.New()
	if _, err = h.Write([]byte(data)); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

// MustEncryptBytes encrypts `data` using MD5 algorithms.
// It panics if any error occurs.
func Md5MustEncryptBytes(data []byte) string {
	result, err := Md5EncryptBytes(data)
	if err != nil {
		panic(err)
	}
	return result
}

// EncryptBytes encrypts string `data` using MD5 algorithms.
func Md5EncryptString(data string) (encrypt string, err error) {
	return Md5EncryptBytes([]byte(data))
}

// MustEncryptString encrypts string `data` using MD5 algorithms.
// It panics if any error occurs.
func Md5MustEncryptString(data string) string {
	result, err := Md5EncryptString(data)
	if err != nil {
		panic(err)
	}
	return result
}

// EncryptFile encrypts file content of `path` using MD5 algorithms.
func Md5EncryptFile(path string) (encrypt string, err error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	h := md5.New()
	_, err = io.Copy(h, f)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

// MustEncryptFile encrypts file content of `path` using MD5 algorithms.
// It panics if any error occurs.
func Md5MustEncryptFile(path string) string {
	result, err := Md5EncryptFile(path)
	if err != nil {
		panic(err)
	}
	return result
}
