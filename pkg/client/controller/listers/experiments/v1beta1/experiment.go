/*
Copyright 2022 The Kubeflow Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/kubeflow/katib/pkg/apis/controller/experiments/v1beta1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// ExperimentLister helps list Experiments.
// All objects returned here must be treated as read-only.
type ExperimentLister interface {
	// List lists all Experiments in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.Experiment, err error)
	// Experiments returns an object that can list and get Experiments.
	Experiments(namespace string) ExperimentNamespaceLister
	ExperimentListerExpansion
}

// experimentLister implements the ExperimentLister interface.
type experimentLister struct {
	listers.ResourceIndexer[*v1beta1.Experiment]
}

// NewExperimentLister returns a new ExperimentLister.
func NewExperimentLister(indexer cache.Indexer) ExperimentLister {
	return &experimentLister{listers.New[*v1beta1.Experiment](indexer, v1beta1.Resource("experiment"))}
}

// Experiments returns an object that can list and get Experiments.
func (s *experimentLister) Experiments(namespace string) ExperimentNamespaceLister {
	return experimentNamespaceLister{listers.NewNamespaced[*v1beta1.Experiment](s.ResourceIndexer, namespace)}
}

// ExperimentNamespaceLister helps list and get Experiments.
// All objects returned here must be treated as read-only.
type ExperimentNamespaceLister interface {
	// List lists all Experiments in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.Experiment, err error)
	// Get retrieves the Experiment from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.Experiment, error)
	ExperimentNamespaceListerExpansion
}

// experimentNamespaceLister implements the ExperimentNamespaceLister
// interface.
type experimentNamespaceLister struct {
	listers.ResourceIndexer[*v1beta1.Experiment]
}
