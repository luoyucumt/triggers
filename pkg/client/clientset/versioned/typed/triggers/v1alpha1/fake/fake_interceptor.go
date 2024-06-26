/*
Copyright 2019 The Tekton Authors

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

	v1alpha1 "github.com/tektoncd/triggers/pkg/apis/triggers/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeInterceptors implements InterceptorInterface
type FakeInterceptors struct {
	Fake *FakeTriggersV1alpha1
	ns   string
}

var interceptorsResource = v1alpha1.SchemeGroupVersion.WithResource("interceptors")

var interceptorsKind = v1alpha1.SchemeGroupVersion.WithKind("Interceptor")

// Get takes name of the interceptor, and returns the corresponding interceptor object, and an error if there is any.
func (c *FakeInterceptors) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Interceptor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(interceptorsResource, c.ns, name), &v1alpha1.Interceptor{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Interceptor), err
}

// List takes label and field selectors, and returns the list of Interceptors that match those selectors.
func (c *FakeInterceptors) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.InterceptorList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(interceptorsResource, interceptorsKind, c.ns, opts), &v1alpha1.InterceptorList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.InterceptorList{ListMeta: obj.(*v1alpha1.InterceptorList).ListMeta}
	for _, item := range obj.(*v1alpha1.InterceptorList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested interceptors.
func (c *FakeInterceptors) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(interceptorsResource, c.ns, opts))

}

// Create takes the representation of a interceptor and creates it.  Returns the server's representation of the interceptor, and an error, if there is any.
func (c *FakeInterceptors) Create(ctx context.Context, interceptor *v1alpha1.Interceptor, opts v1.CreateOptions) (result *v1alpha1.Interceptor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(interceptorsResource, c.ns, interceptor), &v1alpha1.Interceptor{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Interceptor), err
}

// Update takes the representation of a interceptor and updates it. Returns the server's representation of the interceptor, and an error, if there is any.
func (c *FakeInterceptors) Update(ctx context.Context, interceptor *v1alpha1.Interceptor, opts v1.UpdateOptions) (result *v1alpha1.Interceptor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(interceptorsResource, c.ns, interceptor), &v1alpha1.Interceptor{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Interceptor), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeInterceptors) UpdateStatus(ctx context.Context, interceptor *v1alpha1.Interceptor, opts v1.UpdateOptions) (*v1alpha1.Interceptor, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(interceptorsResource, "status", c.ns, interceptor), &v1alpha1.Interceptor{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Interceptor), err
}

// Delete takes name of the interceptor and deletes it. Returns an error if one occurs.
func (c *FakeInterceptors) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(interceptorsResource, c.ns, name, opts), &v1alpha1.Interceptor{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeInterceptors) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(interceptorsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.InterceptorList{})
	return err
}

// Patch applies the patch and returns the patched interceptor.
func (c *FakeInterceptors) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Interceptor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(interceptorsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Interceptor{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Interceptor), err
}
