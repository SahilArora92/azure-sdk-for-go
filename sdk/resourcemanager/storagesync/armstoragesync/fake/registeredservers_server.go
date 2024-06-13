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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagesync/armstoragesync/v2"
	"net/http"
	"net/url"
	"regexp"
)

// RegisteredServersServer is a fake server for instances of the armstoragesync.RegisteredServersClient type.
type RegisteredServersServer struct {
	// BeginCreate is the fake for method RegisteredServersClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginCreate func(ctx context.Context, resourceGroupName string, storageSyncServiceName string, serverID string, parameters armstoragesync.RegisteredServerCreateParameters, options *armstoragesync.RegisteredServersClientBeginCreateOptions) (resp azfake.PollerResponder[armstoragesync.RegisteredServersClientCreateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method RegisteredServersClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, storageSyncServiceName string, serverID string, options *armstoragesync.RegisteredServersClientBeginDeleteOptions) (resp azfake.PollerResponder[armstoragesync.RegisteredServersClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method RegisteredServersClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, storageSyncServiceName string, serverID string, options *armstoragesync.RegisteredServersClientGetOptions) (resp azfake.Responder[armstoragesync.RegisteredServersClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByStorageSyncServicePager is the fake for method RegisteredServersClient.NewListByStorageSyncServicePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByStorageSyncServicePager func(resourceGroupName string, storageSyncServiceName string, options *armstoragesync.RegisteredServersClientListByStorageSyncServiceOptions) (resp azfake.PagerResponder[armstoragesync.RegisteredServersClientListByStorageSyncServiceResponse])

	// BeginTriggerRollover is the fake for method RegisteredServersClient.BeginTriggerRollover
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginTriggerRollover func(ctx context.Context, resourceGroupName string, storageSyncServiceName string, serverID string, parameters armstoragesync.TriggerRolloverRequest, options *armstoragesync.RegisteredServersClientBeginTriggerRolloverOptions) (resp azfake.PollerResponder[armstoragesync.RegisteredServersClientTriggerRolloverResponse], errResp azfake.ErrorResponder)

	// BeginUpdate is the fake for method RegisteredServersClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, storageSyncServiceName string, serverID string, parameters armstoragesync.RegisteredServerUpdateParameters, options *armstoragesync.RegisteredServersClientBeginUpdateOptions) (resp azfake.PollerResponder[armstoragesync.RegisteredServersClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewRegisteredServersServerTransport creates a new instance of RegisteredServersServerTransport with the provided implementation.
// The returned RegisteredServersServerTransport instance is connected to an instance of armstoragesync.RegisteredServersClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewRegisteredServersServerTransport(srv *RegisteredServersServer) *RegisteredServersServerTransport {
	return &RegisteredServersServerTransport{
		srv:                              srv,
		beginCreate:                      newTracker[azfake.PollerResponder[armstoragesync.RegisteredServersClientCreateResponse]](),
		beginDelete:                      newTracker[azfake.PollerResponder[armstoragesync.RegisteredServersClientDeleteResponse]](),
		newListByStorageSyncServicePager: newTracker[azfake.PagerResponder[armstoragesync.RegisteredServersClientListByStorageSyncServiceResponse]](),
		beginTriggerRollover:             newTracker[azfake.PollerResponder[armstoragesync.RegisteredServersClientTriggerRolloverResponse]](),
		beginUpdate:                      newTracker[azfake.PollerResponder[armstoragesync.RegisteredServersClientUpdateResponse]](),
	}
}

// RegisteredServersServerTransport connects instances of armstoragesync.RegisteredServersClient to instances of RegisteredServersServer.
// Don't use this type directly, use NewRegisteredServersServerTransport instead.
type RegisteredServersServerTransport struct {
	srv                              *RegisteredServersServer
	beginCreate                      *tracker[azfake.PollerResponder[armstoragesync.RegisteredServersClientCreateResponse]]
	beginDelete                      *tracker[azfake.PollerResponder[armstoragesync.RegisteredServersClientDeleteResponse]]
	newListByStorageSyncServicePager *tracker[azfake.PagerResponder[armstoragesync.RegisteredServersClientListByStorageSyncServiceResponse]]
	beginTriggerRollover             *tracker[azfake.PollerResponder[armstoragesync.RegisteredServersClientTriggerRolloverResponse]]
	beginUpdate                      *tracker[azfake.PollerResponder[armstoragesync.RegisteredServersClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for RegisteredServersServerTransport.
func (r *RegisteredServersServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "RegisteredServersClient.BeginCreate":
		resp, err = r.dispatchBeginCreate(req)
	case "RegisteredServersClient.BeginDelete":
		resp, err = r.dispatchBeginDelete(req)
	case "RegisteredServersClient.Get":
		resp, err = r.dispatchGet(req)
	case "RegisteredServersClient.NewListByStorageSyncServicePager":
		resp, err = r.dispatchNewListByStorageSyncServicePager(req)
	case "RegisteredServersClient.BeginTriggerRollover":
		resp, err = r.dispatchBeginTriggerRollover(req)
	case "RegisteredServersClient.BeginUpdate":
		resp, err = r.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *RegisteredServersServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if r.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := r.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorageSync/storageSyncServices/(?P<storageSyncServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/registeredServers/(?P<serverId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armstoragesync.RegisteredServerCreateParameters](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		storageSyncServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("storageSyncServiceName")])
		if err != nil {
			return nil, err
		}
		serverIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverId")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginCreate(req.Context(), resourceGroupNameParam, storageSyncServiceNameParam, serverIDParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreate = &respr
		r.beginCreate.add(req, beginCreate)
	}

	resp, err := server.PollerResponderNext(beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		r.beginCreate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreate) {
		r.beginCreate.remove(req)
	}

	return resp, nil
}

func (r *RegisteredServersServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if r.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := r.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorageSync/storageSyncServices/(?P<storageSyncServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/registeredServers/(?P<serverId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		storageSyncServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("storageSyncServiceName")])
		if err != nil {
			return nil, err
		}
		serverIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverId")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginDelete(req.Context(), resourceGroupNameParam, storageSyncServiceNameParam, serverIDParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		r.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		r.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		r.beginDelete.remove(req)
	}

	return resp, nil
}

func (r *RegisteredServersServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if r.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorageSync/storageSyncServices/(?P<storageSyncServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/registeredServers/(?P<serverId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	storageSyncServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("storageSyncServiceName")])
	if err != nil {
		return nil, err
	}
	serverIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Get(req.Context(), resourceGroupNameParam, storageSyncServiceNameParam, serverIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RegisteredServer, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).XMSCorrelationRequestID; val != nil {
		resp.Header.Set("x-ms-correlation-request-id", *val)
	}
	if val := server.GetResponse(respr).XMSRequestID; val != nil {
		resp.Header.Set("x-ms-request-id", *val)
	}
	return resp, nil
}

func (r *RegisteredServersServerTransport) dispatchNewListByStorageSyncServicePager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListByStorageSyncServicePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByStorageSyncServicePager not implemented")}
	}
	newListByStorageSyncServicePager := r.newListByStorageSyncServicePager.get(req)
	if newListByStorageSyncServicePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorageSync/storageSyncServices/(?P<storageSyncServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/registeredServers`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		storageSyncServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("storageSyncServiceName")])
		if err != nil {
			return nil, err
		}
		resp := r.srv.NewListByStorageSyncServicePager(resourceGroupNameParam, storageSyncServiceNameParam, nil)
		newListByStorageSyncServicePager = &resp
		r.newListByStorageSyncServicePager.add(req, newListByStorageSyncServicePager)
	}
	resp, err := server.PagerResponderNext(newListByStorageSyncServicePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		r.newListByStorageSyncServicePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByStorageSyncServicePager) {
		r.newListByStorageSyncServicePager.remove(req)
	}
	return resp, nil
}

func (r *RegisteredServersServerTransport) dispatchBeginTriggerRollover(req *http.Request) (*http.Response, error) {
	if r.srv.BeginTriggerRollover == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginTriggerRollover not implemented")}
	}
	beginTriggerRollover := r.beginTriggerRollover.get(req)
	if beginTriggerRollover == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorageSync/storageSyncServices/(?P<storageSyncServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/registeredServers/(?P<serverId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/triggerRollover`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armstoragesync.TriggerRolloverRequest](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		storageSyncServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("storageSyncServiceName")])
		if err != nil {
			return nil, err
		}
		serverIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverId")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginTriggerRollover(req.Context(), resourceGroupNameParam, storageSyncServiceNameParam, serverIDParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginTriggerRollover = &respr
		r.beginTriggerRollover.add(req, beginTriggerRollover)
	}

	resp, err := server.PollerResponderNext(beginTriggerRollover, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		r.beginTriggerRollover.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginTriggerRollover) {
		r.beginTriggerRollover.remove(req)
	}

	return resp, nil
}

func (r *RegisteredServersServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if r.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := r.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorageSync/storageSyncServices/(?P<storageSyncServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/registeredServers/(?P<serverId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armstoragesync.RegisteredServerUpdateParameters](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		storageSyncServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("storageSyncServiceName")])
		if err != nil {
			return nil, err
		}
		serverIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverId")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginUpdate(req.Context(), resourceGroupNameParam, storageSyncServiceNameParam, serverIDParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		r.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		r.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		r.beginUpdate.remove(req)
	}

	return resp, nil
}
