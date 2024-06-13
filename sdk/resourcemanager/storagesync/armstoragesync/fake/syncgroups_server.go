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

// SyncGroupsServer is a fake server for instances of the armstoragesync.SyncGroupsClient type.
type SyncGroupsServer struct {
	// Create is the fake for method SyncGroupsClient.Create
	// HTTP status codes to indicate success: http.StatusOK
	Create func(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, parameters armstoragesync.SyncGroupCreateParameters, options *armstoragesync.SyncGroupsClientCreateOptions) (resp azfake.Responder[armstoragesync.SyncGroupsClientCreateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method SyncGroupsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, options *armstoragesync.SyncGroupsClientDeleteOptions) (resp azfake.Responder[armstoragesync.SyncGroupsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method SyncGroupsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, options *armstoragesync.SyncGroupsClientGetOptions) (resp azfake.Responder[armstoragesync.SyncGroupsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByStorageSyncServicePager is the fake for method SyncGroupsClient.NewListByStorageSyncServicePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByStorageSyncServicePager func(resourceGroupName string, storageSyncServiceName string, options *armstoragesync.SyncGroupsClientListByStorageSyncServiceOptions) (resp azfake.PagerResponder[armstoragesync.SyncGroupsClientListByStorageSyncServiceResponse])
}

// NewSyncGroupsServerTransport creates a new instance of SyncGroupsServerTransport with the provided implementation.
// The returned SyncGroupsServerTransport instance is connected to an instance of armstoragesync.SyncGroupsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewSyncGroupsServerTransport(srv *SyncGroupsServer) *SyncGroupsServerTransport {
	return &SyncGroupsServerTransport{
		srv:                              srv,
		newListByStorageSyncServicePager: newTracker[azfake.PagerResponder[armstoragesync.SyncGroupsClientListByStorageSyncServiceResponse]](),
	}
}

// SyncGroupsServerTransport connects instances of armstoragesync.SyncGroupsClient to instances of SyncGroupsServer.
// Don't use this type directly, use NewSyncGroupsServerTransport instead.
type SyncGroupsServerTransport struct {
	srv                              *SyncGroupsServer
	newListByStorageSyncServicePager *tracker[azfake.PagerResponder[armstoragesync.SyncGroupsClientListByStorageSyncServiceResponse]]
}

// Do implements the policy.Transporter interface for SyncGroupsServerTransport.
func (s *SyncGroupsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "SyncGroupsClient.Create":
		resp, err = s.dispatchCreate(req)
	case "SyncGroupsClient.Delete":
		resp, err = s.dispatchDelete(req)
	case "SyncGroupsClient.Get":
		resp, err = s.dispatchGet(req)
	case "SyncGroupsClient.NewListByStorageSyncServicePager":
		resp, err = s.dispatchNewListByStorageSyncServicePager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *SyncGroupsServerTransport) dispatchCreate(req *http.Request) (*http.Response, error) {
	if s.srv.Create == nil {
		return nil, &nonRetriableError{errors.New("fake for method Create not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorageSync/storageSyncServices/(?P<storageSyncServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/syncGroups/(?P<syncGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armstoragesync.SyncGroupCreateParameters](req)
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
	syncGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("syncGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Create(req.Context(), resourceGroupNameParam, storageSyncServiceNameParam, syncGroupNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SyncGroup, req)
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

func (s *SyncGroupsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if s.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorageSync/storageSyncServices/(?P<storageSyncServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/syncGroups/(?P<syncGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	syncGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("syncGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Delete(req.Context(), resourceGroupNameParam, storageSyncServiceNameParam, syncGroupNameParam, nil)
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
	if val := server.GetResponse(respr).XMSCorrelationRequestID; val != nil {
		resp.Header.Set("x-ms-correlation-request-id", *val)
	}
	if val := server.GetResponse(respr).XMSRequestID; val != nil {
		resp.Header.Set("x-ms-request-id", *val)
	}
	return resp, nil
}

func (s *SyncGroupsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorageSync/storageSyncServices/(?P<storageSyncServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/syncGroups/(?P<syncGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	syncGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("syncGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Get(req.Context(), resourceGroupNameParam, storageSyncServiceNameParam, syncGroupNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SyncGroup, req)
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

func (s *SyncGroupsServerTransport) dispatchNewListByStorageSyncServicePager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListByStorageSyncServicePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByStorageSyncServicePager not implemented")}
	}
	newListByStorageSyncServicePager := s.newListByStorageSyncServicePager.get(req)
	if newListByStorageSyncServicePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorageSync/storageSyncServices/(?P<storageSyncServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/syncGroups`
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
		resp := s.srv.NewListByStorageSyncServicePager(resourceGroupNameParam, storageSyncServiceNameParam, nil)
		newListByStorageSyncServicePager = &resp
		s.newListByStorageSyncServicePager.add(req, newListByStorageSyncServicePager)
	}
	resp, err := server.PagerResponderNext(newListByStorageSyncServicePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListByStorageSyncServicePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByStorageSyncServicePager) {
		s.newListByStorageSyncServicePager.remove(req)
	}
	return resp, nil
}
