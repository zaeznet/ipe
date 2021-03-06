// Copyright 2016 Claudemiro Alves Feitosa Neto. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ipe

import (
	"errors"
	"sync"
)

// db represents a app database
// For now it there is only one memory database implementation
// but in the future I can write a sql implementation
type db interface {
	GetAppByAppID(appID string) (*app, error)
	GetAppByKey(key string) (*app, error)
	AddApp(*app) error
}

// memdb is a in memory implementation of db interface
type memdb struct {
	sync.Mutex
	Apps []*app
}

func newMemdb() *memdb {
	return &memdb{}
}

func (db *memdb) AddApp(a *app) error {
	db.Lock()
	defer db.Unlock()

	db.Apps = append(db.Apps, a)

	return nil
}

// GetAppByAppID returns an App with by appID
func (db *memdb) GetAppByAppID(appID string) (*app, error) {
	for _, a := range db.Apps {
		if a.AppID == appID {
			return a, nil
		}
	}
	return nil, errors.New("App not found")
}

// GetAppByKey returns an App with by key
func (db *memdb) GetAppByKey(key string) (*app, error) {
	for _, a := range db.Apps {
		if a.Key == key {
			return a, nil
		}
	}
	return nil, errors.New("App not found")
}
