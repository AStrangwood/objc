// Copyright (c) 2013 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/mkrautz/objc"
	. "github.com/mkrautz/objc/UIKit"
	. "github.com/mkrautz/objc/Foundation"
	"runtime"
	"log"
)

type AppDelegate struct {
	objc.Object `objc:"GOAppDelegate : NSObject"`
}

func (delegate *AppDelegate) ApplicationDidFinishLaunchingWithOptions(app objc.Object, options objc.Object) {
	log.Printf("applicationDidFinishLaunching! %v", app)
}

func init() {
	defer runtime.LockOSThread()

	c := objc.NewClass(AppDelegate{})
	c.AddMethod("application:didFinishLaunchingWithOptions:", (*AppDelegate).ApplicationDidFinishLaunchingWithOptions)
	objc.RegisterClass(c)
}

func main() {
	pool := NewNSAutoreleasePool()			
	defer pool.Release()

	UIApplicationMain("UIApplication", "GOAppDelegate")
}
