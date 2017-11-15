/*
Copyright 2017 The KubeDB Authors.

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
	v1alpha1 "github.com/k8sdb/apimachinery/apis/kubedb/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRedises implements RedisInterface
type FakeRedises struct {
	Fake *FakeKubedbV1alpha1
	ns   string
}

var redisesResource = schema.GroupVersionResource{Group: "kubedb.com", Version: "v1alpha1", Resource: "redises"}

var redisesKind = schema.GroupVersionKind{Group: "kubedb.com", Version: "v1alpha1", Kind: "Redis"}

// Get takes name of the redis, and returns the corresponding redis object, and an error if there is any.
func (c *FakeRedises) Get(name string, options v1.GetOptions) (result *v1alpha1.Redis, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(redisesResource, c.ns, name), &v1alpha1.Redis{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Redis), err
}

// List takes label and field selectors, and returns the list of Redises that match those selectors.
func (c *FakeRedises) List(opts v1.ListOptions) (result *v1alpha1.RedisList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(redisesResource, redisesKind, c.ns, opts), &v1alpha1.RedisList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RedisList{}
	for _, item := range obj.(*v1alpha1.RedisList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested redises.
func (c *FakeRedises) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(redisesResource, c.ns, opts))

}

// Create takes the representation of a redis and creates it.  Returns the server's representation of the redis, and an error, if there is any.
func (c *FakeRedises) Create(redis *v1alpha1.Redis) (result *v1alpha1.Redis, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(redisesResource, c.ns, redis), &v1alpha1.Redis{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Redis), err
}

// Update takes the representation of a redis and updates it. Returns the server's representation of the redis, and an error, if there is any.
func (c *FakeRedises) Update(redis *v1alpha1.Redis) (result *v1alpha1.Redis, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(redisesResource, c.ns, redis), &v1alpha1.Redis{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Redis), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRedises) UpdateStatus(redis *v1alpha1.Redis) (*v1alpha1.Redis, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(redisesResource, "status", c.ns, redis), &v1alpha1.Redis{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Redis), err
}

// Delete takes name of the redis and deletes it. Returns an error if one occurs.
func (c *FakeRedises) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(redisesResource, c.ns, name), &v1alpha1.Redis{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRedises) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(redisesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.RedisList{})
	return err
}

// Patch applies the patch and returns the patched redis.
func (c *FakeRedises) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Redis, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(redisesResource, c.ns, name, data, subresources...), &v1alpha1.Redis{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Redis), err
}
