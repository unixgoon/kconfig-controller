/*
Copyright The Kubernetes Authors.

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

// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/att-cloudnative-labs/kconfig-controller/pkg/apis/kconfigcontroller/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// StatefulSetBindingLister helps list StatefulSetBindings.
type StatefulSetBindingLister interface {
	// List lists all StatefulSetBindings in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.StatefulSetBinding, err error)
	// StatefulSetBindings returns an object that can list and get StatefulSetBindings.
	StatefulSetBindings(namespace string) StatefulSetBindingNamespaceLister
	StatefulSetBindingListerExpansion
}

// statefulSetBindingLister implements the StatefulSetBindingLister interface.
type statefulSetBindingLister struct {
	indexer cache.Indexer
}

// NewStatefulSetBindingLister returns a new StatefulSetBindingLister.
func NewStatefulSetBindingLister(indexer cache.Indexer) StatefulSetBindingLister {
	return &statefulSetBindingLister{indexer: indexer}
}

// List lists all StatefulSetBindings in the indexer.
func (s *statefulSetBindingLister) List(selector labels.Selector) (ret []*v1alpha1.StatefulSetBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StatefulSetBinding))
	})
	return ret, err
}

// StatefulSetBindings returns an object that can list and get StatefulSetBindings.
func (s *statefulSetBindingLister) StatefulSetBindings(namespace string) StatefulSetBindingNamespaceLister {
	return statefulSetBindingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// StatefulSetBindingNamespaceLister helps list and get StatefulSetBindings.
type StatefulSetBindingNamespaceLister interface {
	// List lists all StatefulSetBindings in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.StatefulSetBinding, err error)
	// Get retrieves the StatefulSetBinding from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.StatefulSetBinding, error)
	StatefulSetBindingNamespaceListerExpansion
}

// statefulSetBindingNamespaceLister implements the StatefulSetBindingNamespaceLister
// interface.
type statefulSetBindingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all StatefulSetBindings in the indexer for a given namespace.
func (s statefulSetBindingNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.StatefulSetBinding, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StatefulSetBinding))
	})
	return ret, err
}

// Get retrieves the StatefulSetBinding from the indexer for a given namespace and name.
func (s statefulSetBindingNamespaceLister) Get(name string) (*v1alpha1.StatefulSetBinding, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("statefulsetbinding"), name)
	}
	return obj.(*v1alpha1.StatefulSetBinding), nil
}
