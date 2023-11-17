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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v3"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
)

// ServersServer is a fake server for instances of the armpostgresqlflexibleservers.ServersClient type.
type ServersServer struct {
	// BeginCreate is the fake for method ServersClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted
	BeginCreate func(ctx context.Context, resourceGroupName string, serverName string, parameters armpostgresqlflexibleservers.Server, options *armpostgresqlflexibleservers.ServersClientBeginCreateOptions) (resp azfake.PollerResponder[armpostgresqlflexibleservers.ServersClientCreateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method ServersClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, serverName string, options *armpostgresqlflexibleservers.ServersClientBeginDeleteOptions) (resp azfake.PollerResponder[armpostgresqlflexibleservers.ServersClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ServersClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, serverName string, options *armpostgresqlflexibleservers.ServersClientGetOptions) (resp azfake.Responder[armpostgresqlflexibleservers.ServersClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ServersClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armpostgresqlflexibleservers.ServersClientListOptions) (resp azfake.PagerResponder[armpostgresqlflexibleservers.ServersClientListResponse])

	// NewListByResourceGroupPager is the fake for method ServersClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armpostgresqlflexibleservers.ServersClientListByResourceGroupOptions) (resp azfake.PagerResponder[armpostgresqlflexibleservers.ServersClientListByResourceGroupResponse])

	// BeginRestart is the fake for method ServersClient.BeginRestart
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginRestart func(ctx context.Context, resourceGroupName string, serverName string, options *armpostgresqlflexibleservers.ServersClientBeginRestartOptions) (resp azfake.PollerResponder[armpostgresqlflexibleservers.ServersClientRestartResponse], errResp azfake.ErrorResponder)

	// BeginStart is the fake for method ServersClient.BeginStart
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginStart func(ctx context.Context, resourceGroupName string, serverName string, options *armpostgresqlflexibleservers.ServersClientBeginStartOptions) (resp azfake.PollerResponder[armpostgresqlflexibleservers.ServersClientStartResponse], errResp azfake.ErrorResponder)

	// BeginStop is the fake for method ServersClient.BeginStop
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginStop func(ctx context.Context, resourceGroupName string, serverName string, options *armpostgresqlflexibleservers.ServersClientBeginStopOptions) (resp azfake.PollerResponder[armpostgresqlflexibleservers.ServersClientStopResponse], errResp azfake.ErrorResponder)

	// BeginUpdate is the fake for method ServersClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, serverName string, parameters armpostgresqlflexibleservers.ServerForUpdate, options *armpostgresqlflexibleservers.ServersClientBeginUpdateOptions) (resp azfake.PollerResponder[armpostgresqlflexibleservers.ServersClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewServersServerTransport creates a new instance of ServersServerTransport with the provided implementation.
// The returned ServersServerTransport instance is connected to an instance of armpostgresqlflexibleservers.ServersClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServersServerTransport(srv *ServersServer) *ServersServerTransport {
	return &ServersServerTransport{
		srv:                         srv,
		beginCreate:                 newTracker[azfake.PollerResponder[armpostgresqlflexibleservers.ServersClientCreateResponse]](),
		beginDelete:                 newTracker[azfake.PollerResponder[armpostgresqlflexibleservers.ServersClientDeleteResponse]](),
		newListPager:                newTracker[azfake.PagerResponder[armpostgresqlflexibleservers.ServersClientListResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armpostgresqlflexibleservers.ServersClientListByResourceGroupResponse]](),
		beginRestart:                newTracker[azfake.PollerResponder[armpostgresqlflexibleservers.ServersClientRestartResponse]](),
		beginStart:                  newTracker[azfake.PollerResponder[armpostgresqlflexibleservers.ServersClientStartResponse]](),
		beginStop:                   newTracker[azfake.PollerResponder[armpostgresqlflexibleservers.ServersClientStopResponse]](),
		beginUpdate:                 newTracker[azfake.PollerResponder[armpostgresqlflexibleservers.ServersClientUpdateResponse]](),
	}
}

// ServersServerTransport connects instances of armpostgresqlflexibleservers.ServersClient to instances of ServersServer.
// Don't use this type directly, use NewServersServerTransport instead.
type ServersServerTransport struct {
	srv                         *ServersServer
	beginCreate                 *tracker[azfake.PollerResponder[armpostgresqlflexibleservers.ServersClientCreateResponse]]
	beginDelete                 *tracker[azfake.PollerResponder[armpostgresqlflexibleservers.ServersClientDeleteResponse]]
	newListPager                *tracker[azfake.PagerResponder[armpostgresqlflexibleservers.ServersClientListResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armpostgresqlflexibleservers.ServersClientListByResourceGroupResponse]]
	beginRestart                *tracker[azfake.PollerResponder[armpostgresqlflexibleservers.ServersClientRestartResponse]]
	beginStart                  *tracker[azfake.PollerResponder[armpostgresqlflexibleservers.ServersClientStartResponse]]
	beginStop                   *tracker[azfake.PollerResponder[armpostgresqlflexibleservers.ServersClientStopResponse]]
	beginUpdate                 *tracker[azfake.PollerResponder[armpostgresqlflexibleservers.ServersClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for ServersServerTransport.
func (s *ServersServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ServersClient.BeginCreate":
		resp, err = s.dispatchBeginCreate(req)
	case "ServersClient.BeginDelete":
		resp, err = s.dispatchBeginDelete(req)
	case "ServersClient.Get":
		resp, err = s.dispatchGet(req)
	case "ServersClient.NewListPager":
		resp, err = s.dispatchNewListPager(req)
	case "ServersClient.NewListByResourceGroupPager":
		resp, err = s.dispatchNewListByResourceGroupPager(req)
	case "ServersClient.BeginRestart":
		resp, err = s.dispatchBeginRestart(req)
	case "ServersClient.BeginStart":
		resp, err = s.dispatchBeginStart(req)
	case "ServersClient.BeginStop":
		resp, err = s.dispatchBeginStop(req)
	case "ServersClient.BeginUpdate":
		resp, err = s.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *ServersServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if s.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := s.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DBforPostgreSQL/flexibleServers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armpostgresqlflexibleservers.Server](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginCreate(req.Context(), resourceGroupNameParam, serverNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreate = &respr
		s.beginCreate.add(req, beginCreate)
	}

	resp, err := server.PollerResponderNext(beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted}, resp.StatusCode) {
		s.beginCreate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreate) {
		s.beginCreate.remove(req)
	}

	return resp, nil
}

func (s *ServersServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if s.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := s.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DBforPostgreSQL/flexibleServers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginDelete(req.Context(), resourceGroupNameParam, serverNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		s.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		s.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		s.beginDelete.remove(req)
	}

	return resp, nil
}

func (s *ServersServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DBforPostgreSQL/flexibleServers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Get(req.Context(), resourceGroupNameParam, serverNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Server, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ServersServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := s.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DBforPostgreSQL/flexibleServers`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := s.srv.NewListPager(nil)
		newListPager = &resp
		s.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armpostgresqlflexibleservers.ServersClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		s.newListPager.remove(req)
	}
	return resp, nil
}

func (s *ServersServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := s.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DBforPostgreSQL/flexibleServers`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := s.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		s.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armpostgresqlflexibleservers.ServersClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		s.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (s *ServersServerTransport) dispatchBeginRestart(req *http.Request) (*http.Response, error) {
	if s.srv.BeginRestart == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginRestart not implemented")}
	}
	beginRestart := s.beginRestart.get(req)
	if beginRestart == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DBforPostgreSQL/flexibleServers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/restart`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armpostgresqlflexibleservers.RestartParameter](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
		if err != nil {
			return nil, err
		}
		var options *armpostgresqlflexibleservers.ServersClientBeginRestartOptions
		if !reflect.ValueOf(body).IsZero() {
			options = &armpostgresqlflexibleservers.ServersClientBeginRestartOptions{
				Parameters: &body,
			}
		}
		respr, errRespr := s.srv.BeginRestart(req.Context(), resourceGroupNameParam, serverNameParam, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginRestart = &respr
		s.beginRestart.add(req, beginRestart)
	}

	resp, err := server.PollerResponderNext(beginRestart, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		s.beginRestart.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginRestart) {
		s.beginRestart.remove(req)
	}

	return resp, nil
}

func (s *ServersServerTransport) dispatchBeginStart(req *http.Request) (*http.Response, error) {
	if s.srv.BeginStart == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginStart not implemented")}
	}
	beginStart := s.beginStart.get(req)
	if beginStart == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DBforPostgreSQL/flexibleServers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/start`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginStart(req.Context(), resourceGroupNameParam, serverNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginStart = &respr
		s.beginStart.add(req, beginStart)
	}

	resp, err := server.PollerResponderNext(beginStart, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		s.beginStart.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginStart) {
		s.beginStart.remove(req)
	}

	return resp, nil
}

func (s *ServersServerTransport) dispatchBeginStop(req *http.Request) (*http.Response, error) {
	if s.srv.BeginStop == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginStop not implemented")}
	}
	beginStop := s.beginStop.get(req)
	if beginStop == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DBforPostgreSQL/flexibleServers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/stop`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginStop(req.Context(), resourceGroupNameParam, serverNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginStop = &respr
		s.beginStop.add(req, beginStop)
	}

	resp, err := server.PollerResponderNext(beginStop, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		s.beginStop.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginStop) {
		s.beginStop.remove(req)
	}

	return resp, nil
}

func (s *ServersServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if s.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := s.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DBforPostgreSQL/flexibleServers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armpostgresqlflexibleservers.ServerForUpdate](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginUpdate(req.Context(), resourceGroupNameParam, serverNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		s.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		s.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		s.beginUpdate.remove(req)
	}

	return resp, nil
}
