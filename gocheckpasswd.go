package gocheckpasswd

import (
	"strings"
)

// Copyright 2019 Ninja Software. All rights reserved.
// Use of this source code is governed by a
// license that can be found in the LICENSE file.

// Package go-check-passwd checks the password for commonly used password

// IsCommon check is the password a common password
func IsCommon(password string) bool {
	commonPasswords := strings.Split(strCommonPass100k, "\n")

	for _, commonPassword := range commonPasswords {
		if password == commonPassword {
			return true
		}
	}

	return false
}
