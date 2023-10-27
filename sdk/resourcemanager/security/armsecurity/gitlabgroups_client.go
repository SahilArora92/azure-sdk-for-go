//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurity

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// GitLabGroupsClient contains the methods for the GitLabGroups group.
// Don't use this type directly, use NewGitLabGroupsClient() instead.
type GitLabGroupsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewGitLabGroupsClient creates a new instance of GitLabGroupsClient with the specified values.
//   - subscriptionID - Azure subscription ID
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewGitLabGroupsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*GitLabGroupsClient, error) {
	cl, err := arm.NewClient(moduleName+".GitLabGroupsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &GitLabGroupsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Returns a monitored GitLab Group resource for a given fully-qualified name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - securityConnectorName - The security connector name.
//   - groupFQName - The GitLab group fully-qualified name.
//   - options - GitLabGroupsClientGetOptions contains the optional parameters for the GitLabGroupsClient.Get method.
func (client *GitLabGroupsClient) Get(ctx context.Context, resourceGroupName string, securityConnectorName string, groupFQName string, options *GitLabGroupsClientGetOptions) (GitLabGroupsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, securityConnectorName, groupFQName, options)
	if err != nil {
		return GitLabGroupsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return GitLabGroupsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return GitLabGroupsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *GitLabGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, securityConnectorName string, groupFQName string, options *GitLabGroupsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/securityConnectors/{securityConnectorName}/devops/default/gitLabGroups/{groupFQName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if securityConnectorName == "" {
		return nil, errors.New("parameter securityConnectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securityConnectorName}", url.PathEscape(securityConnectorName))
	if groupFQName == "" {
		return nil, errors.New("parameter groupFQName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupFQName}", url.PathEscape(groupFQName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *GitLabGroupsClient) getHandleResponse(resp *http.Response) (GitLabGroupsClientGetResponse, error) {
	result := GitLabGroupsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GitLabGroup); err != nil {
		return GitLabGroupsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Returns a list of GitLab groups onboarded to the connector.
//
// Generated from API version 2023-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - securityConnectorName - The security connector name.
//   - options - GitLabGroupsClientListOptions contains the optional parameters for the GitLabGroupsClient.NewListPager method.
func (client *GitLabGroupsClient) NewListPager(resourceGroupName string, securityConnectorName string, options *GitLabGroupsClientListOptions) *runtime.Pager[GitLabGroupsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[GitLabGroupsClientListResponse]{
		More: func(page GitLabGroupsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *GitLabGroupsClientListResponse) (GitLabGroupsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, securityConnectorName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return GitLabGroupsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return GitLabGroupsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return GitLabGroupsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *GitLabGroupsClient) listCreateRequest(ctx context.Context, resourceGroupName string, securityConnectorName string, options *GitLabGroupsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/securityConnectors/{securityConnectorName}/devops/default/gitLabGroups"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if securityConnectorName == "" {
		return nil, errors.New("parameter securityConnectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securityConnectorName}", url.PathEscape(securityConnectorName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *GitLabGroupsClient) listHandleResponse(resp *http.Response) (GitLabGroupsClientListResponse, error) {
	result := GitLabGroupsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GitLabGroupListResponse); err != nil {
		return GitLabGroupsClientListResponse{}, err
	}
	return result, nil
}

// ListAvailable - Returns a list of all GitLab groups accessible by the user token consumed by the connector.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - securityConnectorName - The security connector name.
//   - options - GitLabGroupsClientListAvailableOptions contains the optional parameters for the GitLabGroupsClient.ListAvailable
//     method.
func (client *GitLabGroupsClient) ListAvailable(ctx context.Context, resourceGroupName string, securityConnectorName string, options *GitLabGroupsClientListAvailableOptions) (GitLabGroupsClientListAvailableResponse, error) {
	var err error
	req, err := client.listAvailableCreateRequest(ctx, resourceGroupName, securityConnectorName, options)
	if err != nil {
		return GitLabGroupsClientListAvailableResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return GitLabGroupsClientListAvailableResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return GitLabGroupsClientListAvailableResponse{}, err
	}
	resp, err := client.listAvailableHandleResponse(httpResp)
	return resp, err
}

// listAvailableCreateRequest creates the ListAvailable request.
func (client *GitLabGroupsClient) listAvailableCreateRequest(ctx context.Context, resourceGroupName string, securityConnectorName string, options *GitLabGroupsClientListAvailableOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/securityConnectors/{securityConnectorName}/devops/default/listAvailableGitLabGroups"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if securityConnectorName == "" {
		return nil, errors.New("parameter securityConnectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securityConnectorName}", url.PathEscape(securityConnectorName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listAvailableHandleResponse handles the ListAvailable response.
func (client *GitLabGroupsClient) listAvailableHandleResponse(resp *http.Response) (GitLabGroupsClientListAvailableResponse, error) {
	result := GitLabGroupsClientListAvailableResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GitLabGroupListResponse); err != nil {
		return GitLabGroupsClientListAvailableResponse{}, err
	}
	return result, nil
}
