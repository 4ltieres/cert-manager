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

package middleware

import (
	"context"

	"github.com/golang/glog"

	"github.com/4ltieres/cert-manager/pkg/acme/client"
	"github.com/4ltieres/cert-manager/third_party/crypto/acme"
)

func NewLogger(baseCl client.Interface) client.Interface {
	return &Logger{baseCl: baseCl}
}

// Logger is a glog based logging middleware for an ACME client
type Logger struct {
	baseCl client.Interface
}

func (l *Logger) CreateOrder(ctx context.Context, order *acme.Order) (*acme.Order, error) {
	glog.Infof("Calling CreateOrder")
	return l.baseCl.CreateOrder(ctx, order)
}

func (l *Logger) GetOrder(ctx context.Context, url string) (*acme.Order, error) {
	glog.Infof("Calling GetOrder")
	return l.baseCl.GetOrder(ctx, url)
}

func (l *Logger) GetCertificate(ctx context.Context, url string) ([][]byte, error) {
	glog.Infof("Calling GetCertificate")
	return l.baseCl.GetCertificate(ctx, url)
}

func (l *Logger) WaitOrder(ctx context.Context, url string) (*acme.Order, error) {
	glog.Infof("Calling WaitOrder")
	return l.baseCl.WaitOrder(ctx, url)
}

func (l *Logger) FinalizeOrder(ctx context.Context, finalizeURL string, csr []byte) (der [][]byte, err error) {
	glog.Infof("Calling FinalizeOrder")
	return l.baseCl.FinalizeOrder(ctx, finalizeURL, csr)
}

func (l *Logger) AcceptChallenge(ctx context.Context, chal *acme.Challenge) (*acme.Challenge, error) {
	glog.Infof("Calling AcceptChallenge")
	return l.baseCl.AcceptChallenge(ctx, chal)
}

func (l *Logger) GetChallenge(ctx context.Context, url string) (*acme.Challenge, error) {
	glog.Infof("Calling GetChallenge")
	return l.baseCl.GetChallenge(ctx, url)
}

func (l *Logger) GetAuthorization(ctx context.Context, url string) (*acme.Authorization, error) {
	glog.Infof("Calling GetAuthorization")
	return l.baseCl.GetAuthorization(ctx, url)
}

func (l *Logger) WaitAuthorization(ctx context.Context, url string) (*acme.Authorization, error) {
	glog.Infof("Calling WaitAuthorization")
	return l.baseCl.WaitAuthorization(ctx, url)
}

func (l *Logger) CreateAccount(ctx context.Context, a *acme.Account) (*acme.Account, error) {
	glog.Infof("Calling CreateAccount")
	return l.baseCl.CreateAccount(ctx, a)
}

func (l *Logger) GetAccount(ctx context.Context) (*acme.Account, error) {
	glog.Infof("Calling GetAccount")
	return l.baseCl.GetAccount(ctx)
}

func (l *Logger) HTTP01ChallengeResponse(token string) (string, error) {
	glog.Infof("Calling HTTP01ChallengeResponse")
	return l.baseCl.HTTP01ChallengeResponse(token)
}

func (l *Logger) DNS01ChallengeRecord(token string) (string, error) {
	glog.Infof("Calling DNS01ChallengeRecord")
	return l.baseCl.DNS01ChallengeRecord(token)
}
