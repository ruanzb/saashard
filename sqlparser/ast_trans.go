// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The MIT License (MIT)

// Copyright (c) 2016 Jerry Bai

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package sqlparser

// TransactionStatement Transaction Statement
type TransactionStatement interface {
	IStatement()
	ITransactionStatement()
	SQLNode
}

// Begin statement
type Begin struct {
}

func (node *Begin) Format(buf *TrackedBuffer) {
	buf.Fprintf("begin")
}

func (node *Begin) IStatement()            {}
func (node *Begin) ITransactionStatement() {}

// Commit statement
type Commit struct {
}

func (node *Commit) Format(buf *TrackedBuffer) {
	buf.Fprintf("commit")
}

func (node *Commit) IStatement()            {}
func (node *Commit) ITransactionStatement() {}

// Rollback statement
type Rollback struct {
}

func (node *Rollback) Format(buf *TrackedBuffer) {
	buf.Fprintf("rollback")
}

func (node *Rollback) IStatement()            {}
func (node *Rollback) ITransactionStatement() {}
