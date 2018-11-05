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
	"fmt"

	"github.com/4ltieres/cert-manager/pkg/apis/certmanager/v1alpha1"
	"github.com/4ltieres/cert-manager/pkg/issuer"
)

const (
	// IssuerACME is the name of the ACME issuer
	IssuerACME string = "acme"
	// IssuerCA is the name of the simple issuer
	IssuerCA string = "ca"
	// IssuerVault is the name of the Vault issuer
	IssuerVault string = "vault"
	// IssuerSelfSigned is a self signing issuer
	IssuerSelfSigned string = "selfsigned"
)

// IssuerFactory is an interface that can be used to obtain Issuer implementations.
// It determines which issuer implementation to use by introspecting the
// given Issuer resource.
type IssuerFactory interface {
	IssuerFor(v1alpha1.GenericIssuer) (issuer.Interface, error)
}

// NewIssuerFactory returns a new issuer factory with the given issuer context.
// The context will be injected into each Issuer upon creation.
func NewIssuerFactory(ctx *Context) IssuerFactory {
	return &factory{ctx: ctx}
}

// factory is the default Factory implementation
type factory struct {
	ctx *Context
}

// IssuerFor will return an Issuer interface for the given Issuer. If the
// requested Issuer is not registered, an error will be returned.
// A new instance of the Issuer will be returned for each call to IssuerFor,
// however this is an inexpensive operation and so, Issuers should not need
// to be cached and reused.
func (f *factory) IssuerFor(issuer v1alpha1.GenericIssuer) (issuer.Interface, error) {
	issuerType, err := NameForIssuer(issuer)
	if err != nil {
		return nil, fmt.Errorf("could not get issuer type: %s", err.Error())
	}

	constructorsLock.RLock()
	defer constructorsLock.RUnlock()
	if constructor, ok := constructors[issuerType]; ok {
		return constructor(f.ctx, issuer)
	}

	return nil, fmt.Errorf("issuer '%s' not registered", issuerType)
}

// nameForIssuer determines the name of the issuer implementation given an
// Issuer resource.
func NameForIssuer(i v1alpha1.GenericIssuer) (string, error) {
	switch {
	case i.GetSpec().ACME != nil:
		return IssuerACME, nil
	case i.GetSpec().CA != nil:
		return IssuerCA, nil
	case i.GetSpec().Vault != nil:
		return IssuerVault, nil
	case i.GetSpec().SelfSigned != nil:
		return IssuerSelfSigned, nil
	}
	return "", fmt.Errorf("no issuer specified for Issuer '%s/%s'", i.GetObjectMeta().Namespace, i.GetObjectMeta().Name)
}
