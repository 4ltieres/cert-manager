/*
Copyright 2018 The Jetstack cert-manager contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"sync"

	"github.com/4ltieres/cert-manager/pkg/apis/certmanager/v1alpha1"
	"github.com/4ltieres/cert-manager/pkg/issuer"
)

// This file defines types for controllers to register themselves with the
// controller package.

// Interface represents a controller that can run. 'workers' should be the
// number of independent goroutines for this controller in question that
// are to be run, and the workers should shut down upon a signal on stopCh.
// This method should block until all workers have exited cleanly, thus
// allowing for graceful shutdown of control loops.
type Interface func(workers int, stopCh <-chan struct{}) error

// Constructor is a function that creates a new control loop given a
// controller Context.
type Constructor func(ctx *Context) Interface

var (
	known = make(map[string]Constructor, 0)
)

// Known returns a map of the registered controller Constructors
func Known() map[string]Constructor {
	return known
}

// Register registers a controller constructor with the controller package
func Register(name string, fn Constructor) {
	known[name] = fn
}

// issuerConstructor constructs an issuer given an Issuer resource and a Context.
// An error will be returned if the appropriate issuer is not registered.
type IssuerConstructor func(*Context, v1alpha1.GenericIssuer) (issuer.Interface, error)

var (
	constructors     = make(map[string]IssuerConstructor)
	constructorsLock sync.RWMutex
)

// Register will register an issuer constructor so it can be used within the
// application. 'name' should be unique, and should be used to identify this
// issuer.
// TODO: move this method to be on Factory, and invent a way to obtain a
// SharedFactory. This will make testing easier.
func RegisterIssuer(name string, c IssuerConstructor) {
	constructorsLock.Lock()
	defer constructorsLock.Unlock()
	constructors[name] = c
}
