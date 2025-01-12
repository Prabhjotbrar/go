// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !amd64 && !386
// +build !amd64,!386

package notsha256

func block(dig *digest, p []byte) {
	blockGeneric(dig, p)
}
