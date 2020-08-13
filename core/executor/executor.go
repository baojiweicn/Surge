// Copyright 2020 Surge Project. rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Author : baojiwei@live.com.
package executor

/*
exector is the operator for any command.
*/

type LangType uint16

var (
	LangUnknonw LangType = 0
	LangGolang           = 1
	LangNode             = 2
	LangPython           = 3
	LangRust             = 4
	LangRuby             = 5
)

// Executor : is the exector for a language command executor
type Executor interface {
}
