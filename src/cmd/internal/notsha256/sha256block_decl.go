// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build 386 || amd64
// +build 386 amd64

package notsha256

//go:noescape

func block(dig *digest, p []byte)
