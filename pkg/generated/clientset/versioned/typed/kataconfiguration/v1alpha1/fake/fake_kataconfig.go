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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/openshift/kata-operator/pkg/apis/kataconfiguration/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeKataConfigs implements KataConfigInterface
type FakeKataConfigs struct {
	Fake *FakeKataconfigurationV1alpha1
	ns   string
}

var kataconfigsResource = schema.GroupVersionResource{Group: "kataconfiguration.openshift.io", Version: "v1alpha1", Resource: "kataconfigs"}

var kataconfigsKind = schema.GroupVersionKind{Group: "kataconfiguration.openshift.io", Version: "v1alpha1", Kind: "KataConfig"}

// Get takes name of the kataConfig, and returns the corresponding kataConfig object, and an error if there is any.
func (c *FakeKataConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.KataConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(kataconfigsResource, c.ns, name), &v1alpha1.KataConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KataConfig), err
}

// List takes label and field selectors, and returns the list of KataConfigs that match those selectors.
func (c *FakeKataConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.KataConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(kataconfigsResource, kataconfigsKind, c.ns, opts), &v1alpha1.KataConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.KataConfigList{ListMeta: obj.(*v1alpha1.KataConfigList).ListMeta}
	for _, item := range obj.(*v1alpha1.KataConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kataConfigs.
func (c *FakeKataConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(kataconfigsResource, c.ns, opts))

}

// Create takes the representation of a kataConfig and creates it.  Returns the server's representation of the kataConfig, and an error, if there is any.
func (c *FakeKataConfigs) Create(ctx context.Context, kataConfig *v1alpha1.KataConfig, opts v1.CreateOptions) (result *v1alpha1.KataConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(kataconfigsResource, c.ns, kataConfig), &v1alpha1.KataConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KataConfig), err
}

// Update takes the representation of a kataConfig and updates it. Returns the server's representation of the kataConfig, and an error, if there is any.
func (c *FakeKataConfigs) Update(ctx context.Context, kataConfig *v1alpha1.KataConfig, opts v1.UpdateOptions) (result *v1alpha1.KataConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(kataconfigsResource, c.ns, kataConfig), &v1alpha1.KataConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KataConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKataConfigs) UpdateStatus(ctx context.Context, kataConfig *v1alpha1.KataConfig, opts v1.UpdateOptions) (*v1alpha1.KataConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(kataconfigsResource, "status", c.ns, kataConfig), &v1alpha1.KataConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KataConfig), err
}

// Delete takes name of the kataConfig and deletes it. Returns an error if one occurs.
func (c *FakeKataConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(kataconfigsResource, c.ns, name), &v1alpha1.KataConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKataConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(kataconfigsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.KataConfigList{})
	return err
}

// Patch applies the patch and returns the patched kataConfig.
func (c *FakeKataConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.KataConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(kataconfigsResource, c.ns, name, pt, data, subresources...), &v1alpha1.KataConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KataConfig), err
}
