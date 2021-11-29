// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

// Some of the code in this file is taken from
// https://code.google.com/p/go/source/browse/?repo=crypto#hg%2Fssh%2Fterminal
// and then has been modified by Jochen Voss.
//
// The original code is distributed under the following license:
//
//	   Copyright 2011 The Go Authors. All rights reserved.
//	   Use of this source code is governed by a BSD-style
//	   license that can be found in the LICENSE file.
//
// All changes to the original code are distributed under the
// following license:
//
//	   Copyright 2013 Jochen Voss. All rights reserved.
//	   Use of this source code is governed by a BSD-style
//	   license that can be found in the LICENSE file.

//go:build ((linux && !appengine) || darwin) && linux
// +build linux,!appengine darwin
// +build linux

// Package password provides a function to read passwords on the
// command line on Linux and BSD Unix (including MacOS X) systems.

// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package prompt

import (
	"fmt"
	"io"
	"os"
	"strings"
	"syscall"
	"time"
	"unsafe"

	"github.com/pkg/errors"
)

const ioctlReadTermios = syscall.TCGETS
const ioctlWriteTermios = syscall.TCSETS

var ErrKeyboardInterrupt = errors.New("keyboard interrupt")

func Prompt(msg string, hidden bool) (string, error) {
	bytes, err := Read(msg, hidden)
	if err != nil {
		if err == ErrKeyboardInterrupt {
			//lint:ignore faillint for prompts can't use logs.
			fmt.Println(err)
			os.Exit(0)
		}
		return "", errors.Wrap(err, "read input")
	}

	result := string(bytes)
	return strings.TrimSpace(result), nil
}

func PromptFileName() string {
	for {
		filename, err := Prompt("Enter output filename: ", false)
		if err != nil {
			//lint:ignore faillint for prompts can't use logs.
			fmt.Println("getting input from terminal:", err)
			continue
		}
		return filename
	}
}

// ReadWithTimings prints the given prompt to standard output and then
// reads a line of input from standard input with echoing of input
// disabled.  This is commonly used for inputting passwords and other
// sensitive data.  The byte slice returned does not include the
// terminating "\n".
//
// The time of every keypress during password input is written into
// the channel 'timings'.  This allows to use password input to gather
// entropy for a random number generator.  Care must be taken to not
// disclose these timings to an attacker: there is correlation between
// keys pressed and the times between key presses.
func ReadWithTimings(prompt string, timings chan<- time.Time, hidden bool) ([]byte, error) {
	if timings != nil {
		timings <- time.Now()
	}

	_, err := os.Stdout.Write([]byte(prompt))
	if err != nil {
		return nil, err
	}

	fd := 0
	var oldState syscall.Termios
	_, _, rc := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(fd),
		ioctlReadTermios, uintptr(unsafe.Pointer(&oldState)), 0, 0, 0)
	if rc != 0 {
		return nil, err
	}

	restore := func() {
		// nolint: errcheck
		syscall.Syscall6(syscall.SYS_IOCTL, uintptr(fd),
			ioctlWriteTermios, uintptr(unsafe.Pointer(&oldState)), 0, 0, 0)
		os.Stdout.Write([]byte("\n"))
	}
	defer restore()

	// Go does not allow to interrupt the syscall.Read() on interrupt;
	// either the whole program is aborted, or the .Read() call keeps
	// running.  On the other hand, we need to catch interrupts in
	// order to restore the terminal settings before exiting.  To get
	// the best of both worlds, we switch the terminal to raw mode and
	// interpret control characters manually in the switch statement,
	// below.  On interrupt, instead of sending a signal, we return
	// with error code ErrKeyboardInterrupt.
	newState := oldState

	if hidden {
		newState.Lflag &^= syscall.ECHO
	}
	newState.Lflag &^= syscall.ISIG | syscall.ICANON |
		syscall.IEXTEN
	newState.Iflag &^= syscall.IXON | syscall.IXOFF
	newState.Iflag |= syscall.ICRNL
	newState.Cc[syscall.VMIN] = 1
	_, _, rc = syscall.Syscall6(syscall.SYS_IOCTL, uintptr(fd),
		ioctlWriteTermios, uintptr(unsafe.Pointer(&newState)), 0, 0, 0)
	if rc != 0 {
		return nil, err
	}

	var ret []byte
	quote := false
inputLoop:
	for {
		var buf [1]byte
		n, err := syscall.Read(fd, buf[:])
		if timings != nil {
			timings <- time.Now()
		}
		if err != nil {
			return nil, err
		}
		if n == 0 {
			if len(ret) == 0 {
				return nil, io.EOF
			}
			break
		}
		if quote {
			quote = false
			ret = append(ret, buf[0])
			continue
		}

		switch buf[0] {
		case '\n', newState.Cc[syscall.VEOF], newState.Cc[syscall.VEOL],
			newState.Cc[syscall.VEOL2]:
			break inputLoop
		case newState.Cc[syscall.VERASE]:
			k := len(ret)
			if k > 0 {
				ret = ret[:k-1]
			}
		case newState.Cc[syscall.VINTR], newState.Cc[syscall.VQUIT]:
			return nil, ErrKeyboardInterrupt
		case newState.Cc[syscall.VKILL]:
			ret = []byte{}
		case newState.Cc[syscall.VLNEXT]:
			quote = true
		case newState.Cc[syscall.VWERASE]:
			for len(ret) > 0 && ret[len(ret)-1] != ' ' {
				ret = ret[:len(ret)-1]
			}
		case newState.Cc[syscall.VSTART], newState.Cc[syscall.VSTOP]:
			// ignore
		default:
			ret = append(ret, buf[0])
		}
	}
	return ret, nil
}

// Read prints the given prompt to standard output and then reads a
// line of input from standard input with echoing of input disabled.
// This is commonly used for inputting passwords and other sensitive
// data.  The byte slice returned does not include the terminating
// "\n".
func Read(prompt string, hidden bool) ([]byte, error) {
	return ReadWithTimings(prompt, nil, hidden)
}
