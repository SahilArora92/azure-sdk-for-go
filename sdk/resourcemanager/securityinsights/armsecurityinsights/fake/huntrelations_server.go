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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// HuntRelationsServer is a fake server for instances of the armsecurityinsights.HuntRelationsClient type.
type HuntRelationsServer struct {
	// CreateOrUpdate is the fake for method HuntRelationsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, workspaceName string, huntID string, huntRelationID string, huntRelation armsecurityinsights.HuntRelation, options *armsecurityinsights.HuntRelationsClientCreateOrUpdateOptions) (resp azfake.Responder[armsecurityinsights.HuntRelationsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method HuntRelationsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, workspaceName string, huntID string, huntRelationID string, options *armsecurityinsights.HuntRelationsClientDeleteOptions) (resp azfake.Responder[armsecurityinsights.HuntRelationsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method HuntRelationsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, workspaceName string, huntID string, huntRelationID string, options *armsecurityinsights.HuntRelationsClientGetOptions) (resp azfake.Responder[armsecurityinsights.HuntRelationsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method HuntRelationsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, workspaceName string, huntID string, options *armsecurityinsights.HuntRelationsClientListOptions) (resp azfake.PagerResponder[armsecurityinsights.HuntRelationsClientListResponse])
}

// NewHuntRelationsServerTransport creates a new instance of HuntRelationsServerTransport with the provided implementation.
// The returned HuntRelationsServerTransport instance is connected to an instance of armsecurityinsights.HuntRelationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewHuntRelationsServerTransport(srv *HuntRelationsServer) *HuntRelationsServerTransport {
	return &HuntRelationsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armsecurityinsights.HuntRelationsClientListResponse]](),
	}
}

// HuntRelationsServerTransport connects instances of armsecurityinsights.HuntRelationsClient to instances of HuntRelationsServer.
// Don't use this type directly, use NewHuntRelationsServerTransport instead.
type HuntRelationsServerTransport struct {
	srv          *HuntRelationsServer
	newListPager *tracker[azfake.PagerResponder[armsecurityinsights.HuntRelationsClientListResponse]]
}

// Do implements the policy.Transporter interface for HuntRelationsServerTransport.
func (h *HuntRelationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "HuntRelationsClient.CreateOrUpdate":
		resp, err = h.dispatchCreateOrUpdate(req)
	case "HuntRelationsClient.Delete":
		resp, err = h.dispatchDelete(req)
	case "HuntRelationsClient.Get":
		resp, err = h.dispatchGet(req)
	case "HuntRelationsClient.NewListPager":
		resp, err = h.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h *HuntRelationsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if h.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.SecurityInsights/hunts/(?P<huntId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/relations/(?P<huntRelationId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armsecurityinsights.HuntRelation](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	huntIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("huntId")])
	if err != nil {
		return nil, err
	}
	huntRelationIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("huntRelationId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := h.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, workspaceNameParam, huntIDParam, huntRelationIDParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).HuntRelation, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HuntRelationsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if h.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.SecurityInsights/hunts/(?P<huntId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/relations/(?P<huntRelationId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	huntIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("huntId")])
	if err != nil {
		return nil, err
	}
	huntRelationIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("huntRelationId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := h.srv.Delete(req.Context(), resourceGroupNameParam, workspaceNameParam, huntIDParam, huntRelationIDParam, nil)
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

func (h *HuntRelationsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if h.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.SecurityInsights/hunts/(?P<huntId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/relations/(?P<huntRelationId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	huntIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("huntId")])
	if err != nil {
		return nil, err
	}
	huntRelationIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("huntRelationId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := h.srv.Get(req.Context(), resourceGroupNameParam, workspaceNameParam, huntIDParam, huntRelationIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).HuntRelation, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HuntRelationsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if h.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := h.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.SecurityInsights/hunts/(?P<huntId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/relations`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
		if err != nil {
			return nil, err
		}
		huntIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("huntId")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		orderbyUnescaped, err := url.QueryUnescape(qp.Get("$orderby"))
		if err != nil {
			return nil, err
		}
		orderbyParam := getOptional(orderbyUnescaped)
		topUnescaped, err := url.QueryUnescape(qp.Get("$top"))
		if err != nil {
			return nil, err
		}
		topParam, err := parseOptional(topUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		skipTokenUnescaped, err := url.QueryUnescape(qp.Get("$skipToken"))
		if err != nil {
			return nil, err
		}
		skipTokenParam := getOptional(skipTokenUnescaped)
		var options *armsecurityinsights.HuntRelationsClientListOptions
		if filterParam != nil || orderbyParam != nil || topParam != nil || skipTokenParam != nil {
			options = &armsecurityinsights.HuntRelationsClientListOptions{
				Filter:    filterParam,
				Orderby:   orderbyParam,
				Top:       topParam,
				SkipToken: skipTokenParam,
			}
		}
		resp := h.srv.NewListPager(resourceGroupNameParam, workspaceNameParam, huntIDParam, options)
		newListPager = &resp
		h.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armsecurityinsights.HuntRelationsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		h.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		h.newListPager.remove(req)
	}
	return resp, nil
}
