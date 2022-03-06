// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

//go:build darwin || freebsd || netbsd || openbsd || solaris || dragonfly

package prompt

import "syscall"

const ioctlReadTermios = syscall.TIOCGETA
const ioctlWriteTermios = syscall.TIOCSETA
