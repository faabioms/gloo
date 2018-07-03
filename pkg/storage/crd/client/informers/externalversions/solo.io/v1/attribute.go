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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	time "time"

	versioned "github.com/solo-io/gloo/pkg/storage/crd/client/clientset/versioned"
	internalinterfaces "github.com/solo-io/gloo/pkg/storage/crd/client/informers/externalversions/internalinterfaces"
	v1 "github.com/solo-io/gloo/pkg/storage/crd/client/listers/solo.io/v1"
	solo_io_v1 "github.com/solo-io/gloo/pkg/storage/crd/solo.io/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// AttributeInformer provides access to a shared informer and lister for
// Attributes.
type AttributeInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.AttributeLister
}

type attributeInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewAttributeInformer constructs a new informer for Attribute type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAttributeInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredAttributeInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredAttributeInformer constructs a new informer for Attribute type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAttributeInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.GlooV1().Attributes(namespace).List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.GlooV1().Attributes(namespace).Watch(options)
			},
		},
		&solo_io_v1.Attribute{},
		resyncPeriod,
		indexers,
	)
}

func (f *attributeInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredAttributeInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *attributeInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&solo_io_v1.Attribute{}, f.defaultInformer)
}

func (f *attributeInformer) Lister() v1.AttributeLister {
	return v1.NewAttributeLister(f.Informer().GetIndexer())
}
