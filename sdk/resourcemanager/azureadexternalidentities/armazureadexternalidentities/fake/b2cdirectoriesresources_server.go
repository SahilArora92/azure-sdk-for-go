//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azureadexternalidentities/armazureadexternalidentities"
	"net/http"
	"net/url"
	"regexp"
)

// B2CDirectoriesResourcesServer is a fake server for instances of the armazureadexternalidentities.B2CDirectoriesResourcesClient type.
type B2CDirectoriesResourcesServer struct {
	// BeginCreateOrUpdate is the fake for method B2CDirectoriesResourcesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, resourceName string, resource armazureadexternalidentities.B2CDirectoryResource, options *armazureadexternalidentities.B2CDirectoriesResourcesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armazureadexternalidentities.B2CDirectoriesResourcesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method B2CDirectoriesResourcesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, resourceName string, options *armazureadexternalidentities.B2CDirectoriesResourcesClientBeginDeleteOptions) (resp azfake.PollerResponder[armazureadexternalidentities.B2CDirectoriesResourcesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method B2CDirectoriesResourcesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, resourceName string, options *armazureadexternalidentities.B2CDirectoriesResourcesClientGetOptions) (resp azfake.Responder[armazureadexternalidentities.B2CDirectoriesResourcesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method B2CDirectoriesResourcesClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armazureadexternalidentities.B2CDirectoriesResourcesClientListByResourceGroupOptions) (resp azfake.PagerResponder[armazureadexternalidentities.B2CDirectoriesResourcesClientListByResourceGroupResponse])

	// NewListBySubscriptionPager is the fake for method B2CDirectoriesResourcesClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armazureadexternalidentities.B2CDirectoriesResourcesClientListBySubscriptionOptions) (resp azfake.PagerResponder[armazureadexternalidentities.B2CDirectoriesResourcesClientListBySubscriptionResponse])

	// Update is the fake for method B2CDirectoriesResourcesClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, resourceName string, properties armazureadexternalidentities.B2CDirectoryResourceUpdate, options *armazureadexternalidentities.B2CDirectoriesResourcesClientUpdateOptions) (resp azfake.Responder[armazureadexternalidentities.B2CDirectoriesResourcesClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewB2CDirectoriesResourcesServerTransport creates a new instance of B2CDirectoriesResourcesServerTransport with the provided implementation.
// The returned B2CDirectoriesResourcesServerTransport instance is connected to an instance of armazureadexternalidentities.B2CDirectoriesResourcesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewB2CDirectoriesResourcesServerTransport(srv *B2CDirectoriesResourcesServer) *B2CDirectoriesResourcesServerTransport {
	return &B2CDirectoriesResourcesServerTransport{
		srv:                         srv,
		beginCreateOrUpdate:         newTracker[azfake.PollerResponder[armazureadexternalidentities.B2CDirectoriesResourcesClientCreateOrUpdateResponse]](),
		beginDelete:                 newTracker[azfake.PollerResponder[armazureadexternalidentities.B2CDirectoriesResourcesClientDeleteResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armazureadexternalidentities.B2CDirectoriesResourcesClientListByResourceGroupResponse]](),
		newListBySubscriptionPager:  newTracker[azfake.PagerResponder[armazureadexternalidentities.B2CDirectoriesResourcesClientListBySubscriptionResponse]](),
	}
}

// B2CDirectoriesResourcesServerTransport connects instances of armazureadexternalidentities.B2CDirectoriesResourcesClient to instances of B2CDirectoriesResourcesServer.
// Don't use this type directly, use NewB2CDirectoriesResourcesServerTransport instead.
type B2CDirectoriesResourcesServerTransport struct {
	srv                         *B2CDirectoriesResourcesServer
	beginCreateOrUpdate         *tracker[azfake.PollerResponder[armazureadexternalidentities.B2CDirectoriesResourcesClientCreateOrUpdateResponse]]
	beginDelete                 *tracker[azfake.PollerResponder[armazureadexternalidentities.B2CDirectoriesResourcesClientDeleteResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armazureadexternalidentities.B2CDirectoriesResourcesClientListByResourceGroupResponse]]
	newListBySubscriptionPager  *tracker[azfake.PagerResponder[armazureadexternalidentities.B2CDirectoriesResourcesClientListBySubscriptionResponse]]
}

// Do implements the policy.Transporter interface for B2CDirectoriesResourcesServerTransport.
func (b *B2CDirectoriesResourcesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "B2CDirectoriesResourcesClient.BeginCreateOrUpdate":
		resp, err = b.dispatchBeginCreateOrUpdate(req)
	case "B2CDirectoriesResourcesClient.BeginDelete":
		resp, err = b.dispatchBeginDelete(req)
	case "B2CDirectoriesResourcesClient.Get":
		resp, err = b.dispatchGet(req)
	case "B2CDirectoriesResourcesClient.NewListByResourceGroupPager":
		resp, err = b.dispatchNewListByResourceGroupPager(req)
	case "B2CDirectoriesResourcesClient.NewListBySubscriptionPager":
		resp, err = b.dispatchNewListBySubscriptionPager(req)
	case "B2CDirectoriesResourcesClient.Update":
		resp, err = b.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (b *B2CDirectoriesResourcesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if b.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := b.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureActiveDirectory/b2cDirectories/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armazureadexternalidentities.B2CDirectoryResource](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := b.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, resourceNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		b.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		b.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		b.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (b *B2CDirectoriesResourcesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if b.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := b.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureActiveDirectory/b2cDirectories/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := b.srv.BeginDelete(req.Context(), resourceGroupNameParam, resourceNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		b.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		b.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		b.beginDelete.remove(req)
	}

	return resp, nil
}

func (b *B2CDirectoriesResourcesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if b.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureActiveDirectory/b2cDirectories/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.Get(req.Context(), resourceGroupNameParam, resourceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).B2CDirectoryResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (b *B2CDirectoriesResourcesServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if b.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := b.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureActiveDirectory/b2cDirectories`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := b.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		b.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armazureadexternalidentities.B2CDirectoriesResourcesClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		b.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		b.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (b *B2CDirectoriesResourcesServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if b.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := b.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureActiveDirectory/b2cDirectories`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := b.srv.NewListBySubscriptionPager(nil)
		newListBySubscriptionPager = &resp
		b.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armazureadexternalidentities.B2CDirectoriesResourcesClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		b.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		b.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}

func (b *B2CDirectoriesResourcesServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if b.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureActiveDirectory/b2cDirectories/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armazureadexternalidentities.B2CDirectoryResourceUpdate](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.Update(req.Context(), resourceGroupNameParam, resourceNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).B2CDirectoryResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
