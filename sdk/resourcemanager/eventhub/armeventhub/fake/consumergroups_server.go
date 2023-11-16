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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// ConsumerGroupsServer is a fake server for instances of the armeventhub.ConsumerGroupsClient type.
type ConsumerGroupsServer struct {
	// CreateOrUpdate is the fake for method ConsumerGroupsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, consumerGroupName string, parameters armeventhub.ConsumerGroup, options *armeventhub.ConsumerGroupsClientCreateOrUpdateOptions) (resp azfake.Responder[armeventhub.ConsumerGroupsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method ConsumerGroupsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, consumerGroupName string, options *armeventhub.ConsumerGroupsClientDeleteOptions) (resp azfake.Responder[armeventhub.ConsumerGroupsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ConsumerGroupsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, consumerGroupName string, options *armeventhub.ConsumerGroupsClientGetOptions) (resp azfake.Responder[armeventhub.ConsumerGroupsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByEventHubPager is the fake for method ConsumerGroupsClient.NewListByEventHubPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByEventHubPager func(resourceGroupName string, namespaceName string, eventHubName string, options *armeventhub.ConsumerGroupsClientListByEventHubOptions) (resp azfake.PagerResponder[armeventhub.ConsumerGroupsClientListByEventHubResponse])
}

// NewConsumerGroupsServerTransport creates a new instance of ConsumerGroupsServerTransport with the provided implementation.
// The returned ConsumerGroupsServerTransport instance is connected to an instance of armeventhub.ConsumerGroupsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewConsumerGroupsServerTransport(srv *ConsumerGroupsServer) *ConsumerGroupsServerTransport {
	return &ConsumerGroupsServerTransport{
		srv:                    srv,
		newListByEventHubPager: newTracker[azfake.PagerResponder[armeventhub.ConsumerGroupsClientListByEventHubResponse]](),
	}
}

// ConsumerGroupsServerTransport connects instances of armeventhub.ConsumerGroupsClient to instances of ConsumerGroupsServer.
// Don't use this type directly, use NewConsumerGroupsServerTransport instead.
type ConsumerGroupsServerTransport struct {
	srv                    *ConsumerGroupsServer
	newListByEventHubPager *tracker[azfake.PagerResponder[armeventhub.ConsumerGroupsClientListByEventHubResponse]]
}

// Do implements the policy.Transporter interface for ConsumerGroupsServerTransport.
func (c *ConsumerGroupsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ConsumerGroupsClient.CreateOrUpdate":
		resp, err = c.dispatchCreateOrUpdate(req)
	case "ConsumerGroupsClient.Delete":
		resp, err = c.dispatchDelete(req)
	case "ConsumerGroupsClient.Get":
		resp, err = c.dispatchGet(req)
	case "ConsumerGroupsClient.NewListByEventHubPager":
		resp, err = c.dispatchNewListByEventHubPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *ConsumerGroupsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EventHub/namespaces/(?P<namespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/eventhubs/(?P<eventHubName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/consumergroups/(?P<consumerGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armeventhub.ConsumerGroup](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	namespaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("namespaceName")])
	if err != nil {
		return nil, err
	}
	eventHubNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("eventHubName")])
	if err != nil {
		return nil, err
	}
	consumerGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("consumerGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, namespaceNameParam, eventHubNameParam, consumerGroupNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ConsumerGroup, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ConsumerGroupsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if c.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EventHub/namespaces/(?P<namespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/eventhubs/(?P<eventHubName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/consumergroups/(?P<consumerGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	namespaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("namespaceName")])
	if err != nil {
		return nil, err
	}
	eventHubNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("eventHubName")])
	if err != nil {
		return nil, err
	}
	consumerGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("consumerGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Delete(req.Context(), resourceGroupNameParam, namespaceNameParam, eventHubNameParam, consumerGroupNameParam, nil)
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

func (c *ConsumerGroupsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if c.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EventHub/namespaces/(?P<namespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/eventhubs/(?P<eventHubName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/consumergroups/(?P<consumerGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	namespaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("namespaceName")])
	if err != nil {
		return nil, err
	}
	eventHubNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("eventHubName")])
	if err != nil {
		return nil, err
	}
	consumerGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("consumerGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Get(req.Context(), resourceGroupNameParam, namespaceNameParam, eventHubNameParam, consumerGroupNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ConsumerGroup, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ConsumerGroupsServerTransport) dispatchNewListByEventHubPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListByEventHubPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByEventHubPager not implemented")}
	}
	newListByEventHubPager := c.newListByEventHubPager.get(req)
	if newListByEventHubPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EventHub/namespaces/(?P<namespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/eventhubs/(?P<eventHubName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/consumergroups`
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
		namespaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("namespaceName")])
		if err != nil {
			return nil, err
		}
		eventHubNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("eventHubName")])
		if err != nil {
			return nil, err
		}
		skipUnescaped, err := url.QueryUnescape(qp.Get("$skip"))
		if err != nil {
			return nil, err
		}
		skipParam, err := parseOptional(skipUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
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
		var options *armeventhub.ConsumerGroupsClientListByEventHubOptions
		if skipParam != nil || topParam != nil {
			options = &armeventhub.ConsumerGroupsClientListByEventHubOptions{
				Skip: skipParam,
				Top:  topParam,
			}
		}
		resp := c.srv.NewListByEventHubPager(resourceGroupNameParam, namespaceNameParam, eventHubNameParam, options)
		newListByEventHubPager = &resp
		c.newListByEventHubPager.add(req, newListByEventHubPager)
		server.PagerResponderInjectNextLinks(newListByEventHubPager, req, func(page *armeventhub.ConsumerGroupsClientListByEventHubResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByEventHubPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListByEventHubPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByEventHubPager) {
		c.newListByEventHubPager.remove(req)
	}
	return resp, nil
}
