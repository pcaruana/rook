/*
Copyright 2018 The Kubernetes Authors.

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

package fake

import (
	v1alpha1 "github.com/rook/rook/pkg/apis/nfs.rook.io/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNFSServers implements NFSServerInterface
type FakeNFSServers struct {
	Fake *FakeNfsV1alpha1
	ns   string
}

var nfsserversResource = schema.GroupVersionResource{Group: "nfs.rook.io", Version: "v1alpha1", Resource: "nfsservers"}

var nfsserversKind = schema.GroupVersionKind{Group: "nfs.rook.io", Version: "v1alpha1", Kind: "NFSServer"}

// Get takes name of the nFSServer, and returns the corresponding nFSServer object, and an error if there is any.
func (c *FakeNFSServers) Get(name string, options v1.GetOptions) (result *v1alpha1.NFSServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(nfsserversResource, c.ns, name), &v1alpha1.NFSServer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NFSServer), err
}

// List takes label and field selectors, and returns the list of NFSServers that match those selectors.
func (c *FakeNFSServers) List(opts v1.ListOptions) (result *v1alpha1.NFSServerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(nfsserversResource, nfsserversKind, c.ns, opts), &v1alpha1.NFSServerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NFSServerList{}
	for _, item := range obj.(*v1alpha1.NFSServerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested nFSServers.
func (c *FakeNFSServers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(nfsserversResource, c.ns, opts))

}

// Create takes the representation of a nFSServer and creates it.  Returns the server's representation of the nFSServer, and an error, if there is any.
func (c *FakeNFSServers) Create(nFSServer *v1alpha1.NFSServer) (result *v1alpha1.NFSServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(nfsserversResource, c.ns, nFSServer), &v1alpha1.NFSServer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NFSServer), err
}

// Update takes the representation of a nFSServer and updates it. Returns the server's representation of the nFSServer, and an error, if there is any.
func (c *FakeNFSServers) Update(nFSServer *v1alpha1.NFSServer) (result *v1alpha1.NFSServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(nfsserversResource, c.ns, nFSServer), &v1alpha1.NFSServer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NFSServer), err
}

// Delete takes name of the nFSServer and deletes it. Returns an error if one occurs.
func (c *FakeNFSServers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(nfsserversResource, c.ns, name), &v1alpha1.NFSServer{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNFSServers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(nfsserversResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.NFSServerList{})
	return err
}

// Patch applies the patch and returns the patched nFSServer.
func (c *FakeNFSServers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NFSServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(nfsserversResource, c.ns, name, data, subresources...), &v1alpha1.NFSServer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NFSServer), err
}
