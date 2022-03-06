// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

//go:build linux

package prompt

import "syscall"

const ioctlReadTermios = syscall.TCGETS
const ioctlWriteTermios = syscall.TCSETS
