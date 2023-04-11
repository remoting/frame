// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// Package gcrc32 provides useful API for CRC32 encryption algorithms.
package crypto

import (
	"hash/crc32"
)

// Encrypt encrypts any type of variable using CRC32 algorithms.
// It uses gconv package to convert `v` to its bytes type.
func Crc32Encrypt(bytes []byte) uint32 {
	return crc32.ChecksumIEEE(bytes)
}
