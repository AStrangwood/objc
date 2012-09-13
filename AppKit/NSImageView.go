// Copyright (c) 2012 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package appkit

import (
	"github.com/mkrautz/objc"
)

type NSImageView struct {
	objc.Object
}

func NewNSImageView() NSImageView {
	return NSImageView{objc.GetClass("NSImageView").Alloc().Init()}
}

func (imgView NSImageView) SetImage(img objc.Object) {
	imgView.SendMsg("setImage:", img)
}

func (imgView NSImageView) Image() NSImage {
	return NSImage{imgView.SendMsg("image")}
}