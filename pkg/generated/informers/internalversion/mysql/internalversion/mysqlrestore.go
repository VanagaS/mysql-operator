// Copyright 2018 Oracle and/or its affiliates. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This file was automatically generated by informer-gen

package internalversion

import (
	time "time"

	mysql "github.com/oracle/mysql-operator/pkg/apis/mysql"
	clientset_internalversion "github.com/oracle/mysql-operator/pkg/generated/clientset/internalversion"
	internalinterfaces "github.com/oracle/mysql-operator/pkg/generated/informers/internalversion/internalinterfaces"
	internalversion "github.com/oracle/mysql-operator/pkg/generated/listers/mysql/internalversion"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// MySQLRestoreInformer provides access to a shared informer and lister for
// MySQLRestores.
type MySQLRestoreInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.MySQLRestoreLister
}

type mySQLRestoreInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewMySQLRestoreInformer constructs a new informer for MySQLRestore type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewMySQLRestoreInformer(client clientset_internalversion.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredMySQLRestoreInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredMySQLRestoreInformer constructs a new informer for MySQLRestore type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredMySQLRestoreInformer(client clientset_internalversion.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Mysql().MySQLRestores(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Mysql().MySQLRestores(namespace).Watch(options)
			},
		},
		&mysql.MySQLRestore{},
		resyncPeriod,
		indexers,
	)
}

func (f *mySQLRestoreInformer) defaultInformer(client clientset_internalversion.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredMySQLRestoreInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *mySQLRestoreInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&mysql.MySQLRestore{}, f.defaultInformer)
}

func (f *mySQLRestoreInformer) Lister() internalversion.MySQLRestoreLister {
	return internalversion.NewMySQLRestoreLister(f.Informer().GetIndexer())
}
