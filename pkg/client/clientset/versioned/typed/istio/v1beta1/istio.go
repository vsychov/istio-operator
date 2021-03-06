/*
Copyright 2019 Banzai Cloud.

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

package v1beta1

import (
	"context"
	"time"

	v1beta1 "github.com/banzaicloud/istio-operator/pkg/apis/istio/v1beta1"
	scheme "github.com/banzaicloud/istio-operator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// IstiosGetter has a method to return a IstioInterface.
// A group's client should implement this interface.
type IstiosGetter interface {
	Istios(namespace string) IstioInterface
}

// IstioInterface has methods to work with Istio resources.
type IstioInterface interface {
	Create(ctx context.Context, istio *v1beta1.Istio, opts v1.CreateOptions) (*v1beta1.Istio, error)
	Update(ctx context.Context, istio *v1beta1.Istio, opts v1.UpdateOptions) (*v1beta1.Istio, error)
	UpdateStatus(ctx context.Context, istio *v1beta1.Istio, opts v1.UpdateOptions) (*v1beta1.Istio, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.Istio, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.IstioList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Istio, err error)
	IstioExpansion
}

// istios implements IstioInterface
type istios struct {
	client rest.Interface
	ns     string
}

// newIstios returns a Istios
func newIstios(c *IstioV1beta1Client, namespace string) *istios {
	return &istios{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the istio, and returns the corresponding istio object, and an error if there is any.
func (c *istios) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.Istio, err error) {
	result = &v1beta1.Istio{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("istios").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Istios that match those selectors.
func (c *istios) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.IstioList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.IstioList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("istios").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested istios.
func (c *istios) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("istios").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a istio and creates it.  Returns the server's representation of the istio, and an error, if there is any.
func (c *istios) Create(ctx context.Context, istio *v1beta1.Istio, opts v1.CreateOptions) (result *v1beta1.Istio, err error) {
	result = &v1beta1.Istio{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("istios").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(istio).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a istio and updates it. Returns the server's representation of the istio, and an error, if there is any.
func (c *istios) Update(ctx context.Context, istio *v1beta1.Istio, opts v1.UpdateOptions) (result *v1beta1.Istio, err error) {
	result = &v1beta1.Istio{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("istios").
		Name(istio.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(istio).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *istios) UpdateStatus(ctx context.Context, istio *v1beta1.Istio, opts v1.UpdateOptions) (result *v1beta1.Istio, err error) {
	result = &v1beta1.Istio{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("istios").
		Name(istio.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(istio).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the istio and deletes it. Returns an error if one occurs.
func (c *istios) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("istios").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *istios) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("istios").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched istio.
func (c *istios) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Istio, err error) {
	result = &v1beta1.Istio{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("istios").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
