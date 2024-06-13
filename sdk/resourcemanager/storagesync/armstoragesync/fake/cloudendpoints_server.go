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

// CloudEndpointsServer is a fake server for instances of the armstoragesync.CloudEndpointsClient type.
type CloudEndpointsServer struct {
	// AfsShareMetadataCertificatePublicKeys is the fake for method CloudEndpointsClient.AfsShareMetadataCertificatePublicKeys
	// HTTP status codes to indicate success: http.StatusOK
	AfsShareMetadataCertificatePublicKeys func(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, cloudEndpointName string, options *armstoragesync.CloudEndpointsClientAfsShareMetadataCertificatePublicKeysOptions) (resp azfake.Responder[armstoragesync.CloudEndpointsClientAfsShareMetadataCertificatePublicKeysResponse], errResp azfake.ErrorResponder)

	// BeginCreate is the fake for method CloudEndpointsClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginCreate func(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, cloudEndpointName string, parameters armstoragesync.CloudEndpointCreateParameters, options *armstoragesync.CloudEndpointsClientBeginCreateOptions) (resp azfake.PollerResponder[armstoragesync.CloudEndpointsClientCreateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method CloudEndpointsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, cloudEndpointName string, options *armstoragesync.CloudEndpointsClientBeginDeleteOptions) (resp azfake.PollerResponder[armstoragesync.CloudEndpointsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method CloudEndpointsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, cloudEndpointName string, options *armstoragesync.CloudEndpointsClientGetOptions) (resp azfake.Responder[armstoragesync.CloudEndpointsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListBySyncGroupPager is the fake for method CloudEndpointsClient.NewListBySyncGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySyncGroupPager func(resourceGroupName string, storageSyncServiceName string, syncGroupName string, options *armstoragesync.CloudEndpointsClientListBySyncGroupOptions) (resp azfake.PagerResponder[armstoragesync.CloudEndpointsClientListBySyncGroupResponse])

	// BeginPostBackup is the fake for method CloudEndpointsClient.BeginPostBackup
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginPostBackup func(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, cloudEndpointName string, parameters armstoragesync.BackupRequest, options *armstoragesync.CloudEndpointsClientBeginPostBackupOptions) (resp azfake.PollerResponder[armstoragesync.CloudEndpointsClientPostBackupResponse], errResp azfake.ErrorResponder)

	// BeginPostRestore is the fake for method CloudEndpointsClient.BeginPostRestore
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginPostRestore func(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, cloudEndpointName string, parameters armstoragesync.PostRestoreRequest, options *armstoragesync.CloudEndpointsClientBeginPostRestoreOptions) (resp azfake.PollerResponder[armstoragesync.CloudEndpointsClientPostRestoreResponse], errResp azfake.ErrorResponder)

	// BeginPreBackup is the fake for method CloudEndpointsClient.BeginPreBackup
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginPreBackup func(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, cloudEndpointName string, parameters armstoragesync.BackupRequest, options *armstoragesync.CloudEndpointsClientBeginPreBackupOptions) (resp azfake.PollerResponder[armstoragesync.CloudEndpointsClientPreBackupResponse], errResp azfake.ErrorResponder)

	// BeginPreRestore is the fake for method CloudEndpointsClient.BeginPreRestore
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginPreRestore func(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, cloudEndpointName string, parameters armstoragesync.PreRestoreRequest, options *armstoragesync.CloudEndpointsClientBeginPreRestoreOptions) (resp azfake.PollerResponder[armstoragesync.CloudEndpointsClientPreRestoreResponse], errResp azfake.ErrorResponder)

	// Restoreheartbeat is the fake for method CloudEndpointsClient.Restoreheartbeat
	// HTTP status codes to indicate success: http.StatusOK
	Restoreheartbeat func(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, cloudEndpointName string, options *armstoragesync.CloudEndpointsClientRestoreheartbeatOptions) (resp azfake.Responder[armstoragesync.CloudEndpointsClientRestoreheartbeatResponse], errResp azfake.ErrorResponder)

	// BeginTriggerChangeDetection is the fake for method CloudEndpointsClient.BeginTriggerChangeDetection
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginTriggerChangeDetection func(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, cloudEndpointName string, parameters armstoragesync.TriggerChangeDetectionParameters, options *armstoragesync.CloudEndpointsClientBeginTriggerChangeDetectionOptions) (resp azfake.PollerResponder[armstoragesync.CloudEndpointsClientTriggerChangeDetectionResponse], errResp azfake.ErrorResponder)
}

// NewCloudEndpointsServerTransport creates a new instance of CloudEndpointsServerTransport with the provided implementation.
// The returned CloudEndpointsServerTransport instance is connected to an instance of armstoragesync.CloudEndpointsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewCloudEndpointsServerTransport(srv *CloudEndpointsServer) *CloudEndpointsServerTransport {
	return &CloudEndpointsServerTransport{
		srv:                         srv,
		beginCreate:                 newTracker[azfake.PollerResponder[armstoragesync.CloudEndpointsClientCreateResponse]](),
		beginDelete:                 newTracker[azfake.PollerResponder[armstoragesync.CloudEndpointsClientDeleteResponse]](),
		newListBySyncGroupPager:     newTracker[azfake.PagerResponder[armstoragesync.CloudEndpointsClientListBySyncGroupResponse]](),
		beginPostBackup:             newTracker[azfake.PollerResponder[armstoragesync.CloudEndpointsClientPostBackupResponse]](),
		beginPostRestore:            newTracker[azfake.PollerResponder[armstoragesync.CloudEndpointsClientPostRestoreResponse]](),
		beginPreBackup:              newTracker[azfake.PollerResponder[armstoragesync.CloudEndpointsClientPreBackupResponse]](),
		beginPreRestore:             newTracker[azfake.PollerResponder[armstoragesync.CloudEndpointsClientPreRestoreResponse]](),
		beginTriggerChangeDetection: newTracker[azfake.PollerResponder[armstoragesync.CloudEndpointsClientTriggerChangeDetectionResponse]](),
	}
}

// CloudEndpointsServerTransport connects instances of armstoragesync.CloudEndpointsClient to instances of CloudEndpointsServer.
// Don't use this type directly, use NewCloudEndpointsServerTransport instead.
type CloudEndpointsServerTransport struct {
	srv                         *CloudEndpointsServer
	beginCreate                 *tracker[azfake.PollerResponder[armstoragesync.CloudEndpointsClientCreateResponse]]
	beginDelete                 *tracker[azfake.PollerResponder[armstoragesync.CloudEndpointsClientDeleteResponse]]
	newListBySyncGroupPager     *tracker[azfake.PagerResponder[armstoragesync.CloudEndpointsClientListBySyncGroupResponse]]
	beginPostBackup             *tracker[azfake.PollerResponder[armstoragesync.CloudEndpointsClientPostBackupResponse]]
	beginPostRestore            *tracker[azfake.PollerResponder[armstoragesync.CloudEndpointsClientPostRestoreResponse]]
	beginPreBackup              *tracker[azfake.PollerResponder[armstoragesync.CloudEndpointsClientPreBackupResponse]]
	beginPreRestore             *tracker[azfake.PollerResponder[armstoragesync.CloudEndpointsClientPreRestoreResponse]]
	beginTriggerChangeDetection *tracker[azfake.PollerResponder[armstoragesync.CloudEndpointsClientTriggerChangeDetectionResponse]]
}

// Do implements the policy.Transporter interface for CloudEndpointsServerTransport.
func (c *CloudEndpointsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "CloudEndpointsClient.AfsShareMetadataCertificatePublicKeys":
		resp, err = c.dispatchAfsShareMetadataCertificatePublicKeys(req)
	case "CloudEndpointsClient.BeginCreate":
		resp, err = c.dispatchBeginCreate(req)
	case "CloudEndpointsClient.BeginDelete":
		resp, err = c.dispatchBeginDelete(req)
	case "CloudEndpointsClient.Get":
		resp, err = c.dispatchGet(req)
	case "CloudEndpointsClient.NewListBySyncGroupPager":
		resp, err = c.dispatchNewListBySyncGroupPager(req)
	case "CloudEndpointsClient.BeginPostBackup":
		resp, err = c.dispatchBeginPostBackup(req)
	case "CloudEndpointsClient.BeginPostRestore":
		resp, err = c.dispatchBeginPostRestore(req)
	case "CloudEndpointsClient.BeginPreBackup":
		resp, err = c.dispatchBeginPreBackup(req)
	case "CloudEndpointsClient.BeginPreRestore":
		resp, err = c.dispatchBeginPreRestore(req)
	case "CloudEndpointsClient.Restoreheartbeat":
		resp, err = c.dispatchRestoreheartbeat(req)
	case "CloudEndpointsClient.BeginTriggerChangeDetection":
		resp, err = c.dispatchBeginTriggerChangeDetection(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *CloudEndpointsServerTransport) dispatchAfsShareMetadataCertificatePublicKeys(req *http.Request) (*http.Response, error) {
	if c.srv.AfsShareMetadataCertificatePublicKeys == nil {
		return nil, &nonRetriableError{errors.New("fake for method AfsShareMetadataCertificatePublicKeys not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorageSync/storageSyncServices/(?P<storageSyncServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/syncGroups/(?P<syncGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/cloudEndpoints/(?P<cloudEndpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/afsShareMetadataCertificatePublicKeys`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
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
	cloudEndpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("cloudEndpointName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.AfsShareMetadataCertificatePublicKeys(req.Context(), resourceGroupNameParam, storageSyncServiceNameParam, syncGroupNameParam, cloudEndpointNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).CloudEndpointAfsShareMetadataCertificatePublicKeys, req)
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

func (c *CloudEndpointsServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if c.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := c.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorageSync/storageSyncServices/(?P<storageSyncServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/syncGroups/(?P<syncGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/cloudEndpoints/(?P<cloudEndpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armstoragesync.CloudEndpointCreateParameters](req)
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
		cloudEndpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("cloudEndpointName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginCreate(req.Context(), resourceGroupNameParam, storageSyncServiceNameParam, syncGroupNameParam, cloudEndpointNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreate = &respr
		c.beginCreate.add(req, beginCreate)
	}

	resp, err := server.PollerResponderNext(beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		c.beginCreate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreate) {
		c.beginCreate.remove(req)
	}

	return resp, nil
}

func (c *CloudEndpointsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if c.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := c.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorageSync/storageSyncServices/(?P<storageSyncServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/syncGroups/(?P<syncGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/cloudEndpoints/(?P<cloudEndpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
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
		cloudEndpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("cloudEndpointName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginDelete(req.Context(), resourceGroupNameParam, storageSyncServiceNameParam, syncGroupNameParam, cloudEndpointNameParam, nil)
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

func (c *CloudEndpointsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if c.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorageSync/storageSyncServices/(?P<storageSyncServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/syncGroups/(?P<syncGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/cloudEndpoints/(?P<cloudEndpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
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
	cloudEndpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("cloudEndpointName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Get(req.Context(), resourceGroupNameParam, storageSyncServiceNameParam, syncGroupNameParam, cloudEndpointNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).CloudEndpoint, req)
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

func (c *CloudEndpointsServerTransport) dispatchNewListBySyncGroupPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListBySyncGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySyncGroupPager not implemented")}
	}
	newListBySyncGroupPager := c.newListBySyncGroupPager.get(req)
	if newListBySyncGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorageSync/storageSyncServices/(?P<storageSyncServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/syncGroups/(?P<syncGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/cloudEndpoints`
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
		resp := c.srv.NewListBySyncGroupPager(resourceGroupNameParam, storageSyncServiceNameParam, syncGroupNameParam, nil)
		newListBySyncGroupPager = &resp
		c.newListBySyncGroupPager.add(req, newListBySyncGroupPager)
	}
	resp, err := server.PagerResponderNext(newListBySyncGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListBySyncGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySyncGroupPager) {
		c.newListBySyncGroupPager.remove(req)
	}
	return resp, nil
}

func (c *CloudEndpointsServerTransport) dispatchBeginPostBackup(req *http.Request) (*http.Response, error) {
	if c.srv.BeginPostBackup == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginPostBackup not implemented")}
	}
	beginPostBackup := c.beginPostBackup.get(req)
	if beginPostBackup == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorageSync/storageSyncServices/(?P<storageSyncServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/syncGroups/(?P<syncGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/cloudEndpoints/(?P<cloudEndpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/postbackup`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armstoragesync.BackupRequest](req)
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
		cloudEndpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("cloudEndpointName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginPostBackup(req.Context(), resourceGroupNameParam, storageSyncServiceNameParam, syncGroupNameParam, cloudEndpointNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginPostBackup = &respr
		c.beginPostBackup.add(req, beginPostBackup)
	}

	resp, err := server.PollerResponderNext(beginPostBackup, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		c.beginPostBackup.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginPostBackup) {
		c.beginPostBackup.remove(req)
	}

	return resp, nil
}

func (c *CloudEndpointsServerTransport) dispatchBeginPostRestore(req *http.Request) (*http.Response, error) {
	if c.srv.BeginPostRestore == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginPostRestore not implemented")}
	}
	beginPostRestore := c.beginPostRestore.get(req)
	if beginPostRestore == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorageSync/storageSyncServices/(?P<storageSyncServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/syncGroups/(?P<syncGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/cloudEndpoints/(?P<cloudEndpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/postrestore`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armstoragesync.PostRestoreRequest](req)
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
		cloudEndpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("cloudEndpointName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginPostRestore(req.Context(), resourceGroupNameParam, storageSyncServiceNameParam, syncGroupNameParam, cloudEndpointNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginPostRestore = &respr
		c.beginPostRestore.add(req, beginPostRestore)
	}

	resp, err := server.PollerResponderNext(beginPostRestore, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		c.beginPostRestore.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginPostRestore) {
		c.beginPostRestore.remove(req)
	}

	return resp, nil
}

func (c *CloudEndpointsServerTransport) dispatchBeginPreBackup(req *http.Request) (*http.Response, error) {
	if c.srv.BeginPreBackup == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginPreBackup not implemented")}
	}
	beginPreBackup := c.beginPreBackup.get(req)
	if beginPreBackup == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorageSync/storageSyncServices/(?P<storageSyncServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/syncGroups/(?P<syncGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/cloudEndpoints/(?P<cloudEndpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/prebackup`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armstoragesync.BackupRequest](req)
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
		cloudEndpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("cloudEndpointName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginPreBackup(req.Context(), resourceGroupNameParam, storageSyncServiceNameParam, syncGroupNameParam, cloudEndpointNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginPreBackup = &respr
		c.beginPreBackup.add(req, beginPreBackup)
	}

	resp, err := server.PollerResponderNext(beginPreBackup, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		c.beginPreBackup.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginPreBackup) {
		c.beginPreBackup.remove(req)
	}

	return resp, nil
}

func (c *CloudEndpointsServerTransport) dispatchBeginPreRestore(req *http.Request) (*http.Response, error) {
	if c.srv.BeginPreRestore == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginPreRestore not implemented")}
	}
	beginPreRestore := c.beginPreRestore.get(req)
	if beginPreRestore == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorageSync/storageSyncServices/(?P<storageSyncServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/syncGroups/(?P<syncGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/cloudEndpoints/(?P<cloudEndpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/prerestore`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armstoragesync.PreRestoreRequest](req)
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
		cloudEndpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("cloudEndpointName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginPreRestore(req.Context(), resourceGroupNameParam, storageSyncServiceNameParam, syncGroupNameParam, cloudEndpointNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginPreRestore = &respr
		c.beginPreRestore.add(req, beginPreRestore)
	}

	resp, err := server.PollerResponderNext(beginPreRestore, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		c.beginPreRestore.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginPreRestore) {
		c.beginPreRestore.remove(req)
	}

	return resp, nil
}

func (c *CloudEndpointsServerTransport) dispatchRestoreheartbeat(req *http.Request) (*http.Response, error) {
	if c.srv.Restoreheartbeat == nil {
		return nil, &nonRetriableError{errors.New("fake for method Restoreheartbeat not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorageSync/storageSyncServices/(?P<storageSyncServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/syncGroups/(?P<syncGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/cloudEndpoints/(?P<cloudEndpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/restoreheartbeat`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
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
	cloudEndpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("cloudEndpointName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Restoreheartbeat(req.Context(), resourceGroupNameParam, storageSyncServiceNameParam, syncGroupNameParam, cloudEndpointNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
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

func (c *CloudEndpointsServerTransport) dispatchBeginTriggerChangeDetection(req *http.Request) (*http.Response, error) {
	if c.srv.BeginTriggerChangeDetection == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginTriggerChangeDetection not implemented")}
	}
	beginTriggerChangeDetection := c.beginTriggerChangeDetection.get(req)
	if beginTriggerChangeDetection == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorageSync/storageSyncServices/(?P<storageSyncServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/syncGroups/(?P<syncGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/cloudEndpoints/(?P<cloudEndpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/triggerChangeDetection`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armstoragesync.TriggerChangeDetectionParameters](req)
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
		cloudEndpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("cloudEndpointName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginTriggerChangeDetection(req.Context(), resourceGroupNameParam, storageSyncServiceNameParam, syncGroupNameParam, cloudEndpointNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginTriggerChangeDetection = &respr
		c.beginTriggerChangeDetection.add(req, beginTriggerChangeDetection)
	}

	resp, err := server.PollerResponderNext(beginTriggerChangeDetection, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		c.beginTriggerChangeDetection.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginTriggerChangeDetection) {
		c.beginTriggerChangeDetection.remove(req)
	}

	return resp, nil
}
