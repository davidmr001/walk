// Copyright 2010 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gui

import (
	"os"
)

import (
	. "walk/winapi/user32"
)

type Settings interface {
	Get(key string) (string, bool)
	Put(key, value string) os.Error
	Load() os.Error
	Save() os.Error
}

type Persistable interface {
	Persistent() bool
	SetPersistent(value bool)
	SaveState() os.Error
	RestoreState() os.Error
}

type Application struct {
	organizationName string
	productName      string
	settings         Settings
}

var appSingleton *Application = &Application{}

func App() *Application {
	return appSingleton
}

func (app *Application) OrganizationName() string {
	return app.organizationName
}

func (app *Application) SetOrganizationName(value string) {
	app.organizationName = value
}

func (app *Application) ProductName() string {
	return app.productName
}

func (app *Application) SetProductName(value string) {
	app.productName = value
}

func (app *Application) Settings() Settings {
	return app.settings
}

func (app *Application) SetSettings(value Settings) {
	app.settings = value
}

func (app *Application) Exit(exitCode int) {
	PostQuitMessage(exitCode)
}
