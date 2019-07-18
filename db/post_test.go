// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package db

import (
	"fmt"
	"testing"
)

func TestGetPostsByName(t *testing.T) {
	var post Post
	t.Fatal(post.GetPostsByName(`点道为止`))
}

func TestGetPostsByNameFn(t *testing.T) {
	// 没有成功组装...
	var s = fmt.Sprintf(`%v`, `点道`)
	t.Fatal("%" + s + "%")
}
