// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// Package gsha1 provides useful API for SHA1 encryption algorithms.
package crypto

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"os"
)

// Encrypt encrypts any type of variable using SHA1 algorithms.
// It uses package gconv to convert `v` to its bytes type.
func Sha1Encrypt(data []byte) string {
	r := sha1.Sum(data)
	return hex.EncodeToString(r[:])
}

// EncryptFile encrypts file content of `path` using SHA1 algorithms.
func Sha1EncryptFile(path string) (encrypt string, err error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	h := sha1.New()
	_, err = io.Copy(h, f)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

// MustEncryptFile encrypts file content of `path` using SHA1 algorithms.
// It panics if any error occurs.
func Sha1MustEncryptFile(path string) string {
	result, err := Sha1EncryptFile(path)
	if err != nil {
		panic(err)
	}
	return result
}
