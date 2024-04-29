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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/providerhub/armproviderhub/v2"
	"net/http"
	"net/url"
	"regexp"
)

// NotificationRegistrationsServer is a fake server for instances of the armproviderhub.NotificationRegistrationsClient type.
type NotificationRegistrationsServer struct {
	// CreateOrUpdate is the fake for method NotificationRegistrationsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK
	CreateOrUpdate func(ctx context.Context, providerNamespace string, notificationRegistrationName string, properties armproviderhub.NotificationRegistration, options *armproviderhub.NotificationRegistrationsClientCreateOrUpdateOptions) (resp azfake.Responder[armproviderhub.NotificationRegistrationsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method NotificationRegistrationsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, providerNamespace string, notificationRegistrationName string, options *armproviderhub.NotificationRegistrationsClientDeleteOptions) (resp azfake.Responder[armproviderhub.NotificationRegistrationsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method NotificationRegistrationsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, providerNamespace string, notificationRegistrationName string, options *armproviderhub.NotificationRegistrationsClientGetOptions) (resp azfake.Responder[armproviderhub.NotificationRegistrationsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByProviderRegistrationPager is the fake for method NotificationRegistrationsClient.NewListByProviderRegistrationPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByProviderRegistrationPager func(providerNamespace string, options *armproviderhub.NotificationRegistrationsClientListByProviderRegistrationOptions) (resp azfake.PagerResponder[armproviderhub.NotificationRegistrationsClientListByProviderRegistrationResponse])
}

// NewNotificationRegistrationsServerTransport creates a new instance of NotificationRegistrationsServerTransport with the provided implementation.
// The returned NotificationRegistrationsServerTransport instance is connected to an instance of armproviderhub.NotificationRegistrationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewNotificationRegistrationsServerTransport(srv *NotificationRegistrationsServer) *NotificationRegistrationsServerTransport {
	return &NotificationRegistrationsServerTransport{
		srv:                                srv,
		newListByProviderRegistrationPager: newTracker[azfake.PagerResponder[armproviderhub.NotificationRegistrationsClientListByProviderRegistrationResponse]](),
	}
}

// NotificationRegistrationsServerTransport connects instances of armproviderhub.NotificationRegistrationsClient to instances of NotificationRegistrationsServer.
// Don't use this type directly, use NewNotificationRegistrationsServerTransport instead.
type NotificationRegistrationsServerTransport struct {
	srv                                *NotificationRegistrationsServer
	newListByProviderRegistrationPager *tracker[azfake.PagerResponder[armproviderhub.NotificationRegistrationsClientListByProviderRegistrationResponse]]
}

// Do implements the policy.Transporter interface for NotificationRegistrationsServerTransport.
func (n *NotificationRegistrationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "NotificationRegistrationsClient.CreateOrUpdate":
		resp, err = n.dispatchCreateOrUpdate(req)
	case "NotificationRegistrationsClient.Delete":
		resp, err = n.dispatchDelete(req)
	case "NotificationRegistrationsClient.Get":
		resp, err = n.dispatchGet(req)
	case "NotificationRegistrationsClient.NewListByProviderRegistrationPager":
		resp, err = n.dispatchNewListByProviderRegistrationPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (n *NotificationRegistrationsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if n.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ProviderHub/providerRegistrations/(?P<providerNamespace>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/notificationRegistrations/(?P<notificationRegistrationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armproviderhub.NotificationRegistration](req)
	if err != nil {
		return nil, err
	}
	providerNamespaceParam, err := url.PathUnescape(matches[regex.SubexpIndex("providerNamespace")])
	if err != nil {
		return nil, err
	}
	notificationRegistrationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("notificationRegistrationName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := n.srv.CreateOrUpdate(req.Context(), providerNamespaceParam, notificationRegistrationNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).NotificationRegistration, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (n *NotificationRegistrationsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if n.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ProviderHub/providerRegistrations/(?P<providerNamespace>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/notificationRegistrations/(?P<notificationRegistrationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	providerNamespaceParam, err := url.PathUnescape(matches[regex.SubexpIndex("providerNamespace")])
	if err != nil {
		return nil, err
	}
	notificationRegistrationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("notificationRegistrationName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := n.srv.Delete(req.Context(), providerNamespaceParam, notificationRegistrationNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (n *NotificationRegistrationsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if n.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ProviderHub/providerRegistrations/(?P<providerNamespace>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/notificationRegistrations/(?P<notificationRegistrationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	providerNamespaceParam, err := url.PathUnescape(matches[regex.SubexpIndex("providerNamespace")])
	if err != nil {
		return nil, err
	}
	notificationRegistrationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("notificationRegistrationName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := n.srv.Get(req.Context(), providerNamespaceParam, notificationRegistrationNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).NotificationRegistration, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (n *NotificationRegistrationsServerTransport) dispatchNewListByProviderRegistrationPager(req *http.Request) (*http.Response, error) {
	if n.srv.NewListByProviderRegistrationPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByProviderRegistrationPager not implemented")}
	}
	newListByProviderRegistrationPager := n.newListByProviderRegistrationPager.get(req)
	if newListByProviderRegistrationPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ProviderHub/providerRegistrations/(?P<providerNamespace>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/notificationRegistrations`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		providerNamespaceParam, err := url.PathUnescape(matches[regex.SubexpIndex("providerNamespace")])
		if err != nil {
			return nil, err
		}
		resp := n.srv.NewListByProviderRegistrationPager(providerNamespaceParam, nil)
		newListByProviderRegistrationPager = &resp
		n.newListByProviderRegistrationPager.add(req, newListByProviderRegistrationPager)
		server.PagerResponderInjectNextLinks(newListByProviderRegistrationPager, req, func(page *armproviderhub.NotificationRegistrationsClientListByProviderRegistrationResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByProviderRegistrationPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		n.newListByProviderRegistrationPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByProviderRegistrationPager) {
		n.newListByProviderRegistrationPager.remove(req)
	}
	return resp, nil
}
