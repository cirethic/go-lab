// Copyright (c) 2012 The GoTamer Authors. All rights reserved.
// Use of this source code is governed by the MIT license.
// The license can be found at http://www.robotamer.com

package cfg

import (
	"os"
	"testing"
)

type Sqlite struct {
	Path string
	Port int
	Name string
	Pass string
}

var filename = os.TempDir() + "/testcfg.json"

var dbpath = "/tmp/testcfg.db"
var port int = 8000
var name string = "TaMeR"
var pass string = ":-)"
var sql = &Sqlite{Path: dbpath, Port: port, Name: name, Pass: pass}

func TestJson(t *testing.T) {
	if err := Save(filename, sql); err != nil {
		t.Errorf("Config Save Error: %s", err.Error())
	}
	sqllocal := new(Sqlite)
	if err := Load(filename, sqllocal); err != nil {
		t.Errorf("Config Load Error: %s", err.Error())
	}
	if sqllocal.Path != dbpath {
		t.Errorf("Config Path doesn't match: %s", sqllocal.Path)
	}
	if sqllocal.Port != port {
		t.Errorf("Config Path doesn't match: %s", sqllocal.Port)
	}
}
