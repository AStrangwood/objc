// Copyright (c) 2012 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foundation

import (
	"github.com/mkrautz/objc"
)

type NSDictionary struct {
	objc.Object
}

func (dict NSDictionary) ObjectForKey(key objc.Object) objc.Object {
	return dict.SendMsg("objectForKey:", key)
}