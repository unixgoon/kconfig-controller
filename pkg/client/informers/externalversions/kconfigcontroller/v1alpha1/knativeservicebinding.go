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
	time "time"

	kconfigcontrollerv1alpha1 "github.com/att-cloudnative-labs/kconfig-controller/pkg/apis/kconfigcontroller/v1alpha1"
	versioned "github.com/att-cloudnative-labs/kconfig-controller/pkg/client/clientset/versioned"
	internalinterfaces "github.com/att-cloudnative-labs/kconfig-controller/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/att-cloudnative-labs/kconfig-controller/pkg/client/listers/kconfigcontroller/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// KnativeServiceBindingInformer provides access to a shared informer and lister for
// KnativeServiceBindings.
type KnativeServiceBindingInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.KnativeServiceBindingLister
}

type knativeServiceBindingInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewKnativeServiceBindingInformer constructs a new informer for KnativeServiceBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewKnativeServiceBindingInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredKnativeServiceBindingInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredKnativeServiceBindingInformer constructs a new informer for KnativeServiceBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredKnativeServiceBindingInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KconfigcontrollerV1alpha1().KnativeServiceBindings(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KconfigcontrollerV1alpha1().KnativeServiceBindings(namespace).Watch(options)
			},
		},
		&kconfigcontrollerv1alpha1.KnativeServiceBinding{},
		resyncPeriod,
		indexers,
	)
}

func (f *knativeServiceBindingInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredKnativeServiceBindingInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *knativeServiceBindingInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&kconfigcontrollerv1alpha1.KnativeServiceBinding{}, f.defaultInformer)
}

func (f *knativeServiceBindingInformer) Lister() v1alpha1.KnativeServiceBindingLister {
	return v1alpha1.NewKnativeServiceBindingLister(f.Informer().GetIndexer())
}
