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

package v1alpha1

import (
	v1alpha1 "github.com/4ltieres/cert-manager/pkg/apis/certmanager/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClusterIssuerLister helps list ClusterIssuers.
type ClusterIssuerLister interface {
	// List lists all ClusterIssuers in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ClusterIssuer, err error)
	// Get retrieves the ClusterIssuer from the index for a given name.
	Get(name string) (*v1alpha1.ClusterIssuer, error)
	ClusterIssuerListerExpansion
}

// clusterIssuerLister implements the ClusterIssuerLister interface.
type clusterIssuerLister struct {
	indexer cache.Indexer
}

// NewClusterIssuerLister returns a new ClusterIssuerLister.
func NewClusterIssuerLister(indexer cache.Indexer) ClusterIssuerLister {
	return &clusterIssuerLister{indexer: indexer}
}

// List lists all ClusterIssuers in the indexer.
func (s *clusterIssuerLister) List(selector labels.Selector) (ret []*v1alpha1.ClusterIssuer, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClusterIssuer))
	})
	return ret, err
}

// Get retrieves the ClusterIssuer from the index for a given name.
func (s *clusterIssuerLister) Get(name string) (*v1alpha1.ClusterIssuer, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("clusterissuer"), name)
	}
	return obj.(*v1alpha1.ClusterIssuer), nil
}
