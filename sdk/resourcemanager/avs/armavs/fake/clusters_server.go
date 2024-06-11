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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs/v2"
	"net/http"
	"net/url"
	"regexp"
)

// ClustersServer is a fake server for instances of the armavs.ClustersClient type.
type ClustersServer struct {
	// BeginCreateOrUpdate is the fake for method ClustersClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, privateCloudName string, clusterName string, cluster armavs.Cluster, options *armavs.ClustersClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armavs.ClustersClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method ClustersClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, privateCloudName string, clusterName string, options *armavs.ClustersClientBeginDeleteOptions) (resp azfake.PollerResponder[armavs.ClustersClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ClustersClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, privateCloudName string, clusterName string, options *armavs.ClustersClientGetOptions) (resp azfake.Responder[armavs.ClustersClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ClustersClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, privateCloudName string, options *armavs.ClustersClientListOptions) (resp azfake.PagerResponder[armavs.ClustersClientListResponse])

	// ListZones is the fake for method ClustersClient.ListZones
	// HTTP status codes to indicate success: http.StatusOK
	ListZones func(ctx context.Context, resourceGroupName string, privateCloudName string, clusterName string, options *armavs.ClustersClientListZonesOptions) (resp azfake.Responder[armavs.ClustersClientListZonesResponse], errResp azfake.ErrorResponder)

	// BeginUpdate is the fake for method ClustersClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginUpdate func(ctx context.Context, resourceGroupName string, privateCloudName string, clusterName string, clusterUpdate armavs.ClusterUpdate, options *armavs.ClustersClientBeginUpdateOptions) (resp azfake.PollerResponder[armavs.ClustersClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewClustersServerTransport creates a new instance of ClustersServerTransport with the provided implementation.
// The returned ClustersServerTransport instance is connected to an instance of armavs.ClustersClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewClustersServerTransport(srv *ClustersServer) *ClustersServerTransport {
	return &ClustersServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armavs.ClustersClientCreateOrUpdateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armavs.ClustersClientDeleteResponse]](),
		newListPager:        newTracker[azfake.PagerResponder[armavs.ClustersClientListResponse]](),
		beginUpdate:         newTracker[azfake.PollerResponder[armavs.ClustersClientUpdateResponse]](),
	}
}

// ClustersServerTransport connects instances of armavs.ClustersClient to instances of ClustersServer.
// Don't use this type directly, use NewClustersServerTransport instead.
type ClustersServerTransport struct {
	srv                 *ClustersServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armavs.ClustersClientCreateOrUpdateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armavs.ClustersClientDeleteResponse]]
	newListPager        *tracker[azfake.PagerResponder[armavs.ClustersClientListResponse]]
	beginUpdate         *tracker[azfake.PollerResponder[armavs.ClustersClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for ClustersServerTransport.
func (c *ClustersServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ClustersClient.BeginCreateOrUpdate":
		resp, err = c.dispatchBeginCreateOrUpdate(req)
	case "ClustersClient.BeginDelete":
		resp, err = c.dispatchBeginDelete(req)
	case "ClustersClient.Get":
		resp, err = c.dispatchGet(req)
	case "ClustersClient.NewListPager":
		resp, err = c.dispatchNewListPager(req)
	case "ClustersClient.ListZones":
		resp, err = c.dispatchListZones(req)
	case "ClustersClient.BeginUpdate":
		resp, err = c.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *ClustersServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := c.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AVS/privateClouds/(?P<privateCloudName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/clusters/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armavs.Cluster](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		privateCloudNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateCloudName")])
		if err != nil {
			return nil, err
		}
		clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, privateCloudNameParam, clusterNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		c.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		c.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		c.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (c *ClustersServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if c.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := c.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AVS/privateClouds/(?P<privateCloudName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/clusters/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		privateCloudNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateCloudName")])
		if err != nil {
			return nil, err
		}
		clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginDelete(req.Context(), resourceGroupNameParam, privateCloudNameParam, clusterNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		c.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		c.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		c.beginDelete.remove(req)
	}

	return resp, nil
}

func (c *ClustersServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if c.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AVS/privateClouds/(?P<privateCloudName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/clusters/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	privateCloudNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateCloudName")])
	if err != nil {
		return nil, err
	}
	clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Get(req.Context(), resourceGroupNameParam, privateCloudNameParam, clusterNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Cluster, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ClustersServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := c.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AVS/privateClouds/(?P<privateCloudName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/clusters`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		privateCloudNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateCloudName")])
		if err != nil {
			return nil, err
		}
		resp := c.srv.NewListPager(resourceGroupNameParam, privateCloudNameParam, nil)
		newListPager = &resp
		c.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armavs.ClustersClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		c.newListPager.remove(req)
	}
	return resp, nil
}

func (c *ClustersServerTransport) dispatchListZones(req *http.Request) (*http.Response, error) {
	if c.srv.ListZones == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListZones not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AVS/privateClouds/(?P<privateCloudName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/clusters/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/listZones`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	privateCloudNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateCloudName")])
	if err != nil {
		return nil, err
	}
	clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.ListZones(req.Context(), resourceGroupNameParam, privateCloudNameParam, clusterNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ClusterZoneList, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ClustersServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := c.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AVS/privateClouds/(?P<privateCloudName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/clusters/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armavs.ClusterUpdate](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		privateCloudNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateCloudName")])
		if err != nil {
			return nil, err
		}
		clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginUpdate(req.Context(), resourceGroupNameParam, privateCloudNameParam, clusterNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		c.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		c.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		c.beginUpdate.remove(req)
	}

	return resp, nil
}
