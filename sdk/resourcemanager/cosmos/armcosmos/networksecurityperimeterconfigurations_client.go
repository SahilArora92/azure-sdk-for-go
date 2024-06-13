//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcosmos

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

// NetworkSecurityPerimeterConfigurationsClient contains the methods for the NetworkSecurityPerimeterConfigurations group.
// Don't use this type directly, use NewNetworkSecurityPerimeterConfigurationsClient() instead.
type NetworkSecurityPerimeterConfigurationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewNetworkSecurityPerimeterConfigurationsClient creates a new instance of NetworkSecurityPerimeterConfigurationsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewNetworkSecurityPerimeterConfigurationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*NetworkSecurityPerimeterConfigurationsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &NetworkSecurityPerimeterConfigurationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets effective Network Security Perimeter Configuration for association
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - Cosmos DB database account name.
//   - networkSecurityPerimeterConfigurationName - The name for Network Security Perimeter configuration
//   - options - NetworkSecurityPerimeterConfigurationsClientGetOptions contains the optional parameters for the NetworkSecurityPerimeterConfigurationsClient.Get
//     method.
func (client *NetworkSecurityPerimeterConfigurationsClient) Get(ctx context.Context, resourceGroupName string, accountName string, networkSecurityPerimeterConfigurationName string, options *NetworkSecurityPerimeterConfigurationsClientGetOptions) (NetworkSecurityPerimeterConfigurationsClientGetResponse, error) {
	var err error
	const operationName = "NetworkSecurityPerimeterConfigurationsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, networkSecurityPerimeterConfigurationName, options)
	if err != nil {
		return NetworkSecurityPerimeterConfigurationsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NetworkSecurityPerimeterConfigurationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return NetworkSecurityPerimeterConfigurationsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *NetworkSecurityPerimeterConfigurationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, networkSecurityPerimeterConfigurationName string, options *NetworkSecurityPerimeterConfigurationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/networkSecurityPerimeterConfigurations/{networkSecurityPerimeterConfigurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if networkSecurityPerimeterConfigurationName == "" {
		return nil, errors.New("parameter networkSecurityPerimeterConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkSecurityPerimeterConfigurationName}", url.PathEscape(networkSecurityPerimeterConfigurationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *NetworkSecurityPerimeterConfigurationsClient) getHandleResponse(resp *http.Response) (NetworkSecurityPerimeterConfigurationsClientGetResponse, error) {
	result := NetworkSecurityPerimeterConfigurationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkSecurityPerimeterConfiguration); err != nil {
		return NetworkSecurityPerimeterConfigurationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets list of effective Network Security Perimeter Configuration for cosmos db account
//
// Generated from API version 2024-05-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - Cosmos DB database account name.
//   - options - NetworkSecurityPerimeterConfigurationsClientListOptions contains the optional parameters for the NetworkSecurityPerimeterConfigurationsClient.NewListPager
//     method.
func (client *NetworkSecurityPerimeterConfigurationsClient) NewListPager(resourceGroupName string, accountName string, options *NetworkSecurityPerimeterConfigurationsClientListOptions) *runtime.Pager[NetworkSecurityPerimeterConfigurationsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[NetworkSecurityPerimeterConfigurationsClientListResponse]{
		More: func(page NetworkSecurityPerimeterConfigurationsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *NetworkSecurityPerimeterConfigurationsClientListResponse) (NetworkSecurityPerimeterConfigurationsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "NetworkSecurityPerimeterConfigurationsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, accountName, options)
			}, nil)
			if err != nil {
				return NetworkSecurityPerimeterConfigurationsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *NetworkSecurityPerimeterConfigurationsClient) listCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *NetworkSecurityPerimeterConfigurationsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/networkSecurityPerimeterConfigurations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *NetworkSecurityPerimeterConfigurationsClient) listHandleResponse(resp *http.Response) (NetworkSecurityPerimeterConfigurationsClientListResponse, error) {
	result := NetworkSecurityPerimeterConfigurationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkSecurityPerimeterConfigurationListResult); err != nil {
		return NetworkSecurityPerimeterConfigurationsClientListResponse{}, err
	}
	return result, nil
}

// BeginReconcile - Refreshes any information about the association.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - Cosmos DB database account name.
//   - networkSecurityPerimeterConfigurationName - The name for Network Security Perimeter configuration
//   - options - NetworkSecurityPerimeterConfigurationsClientBeginReconcileOptions contains the optional parameters for the NetworkSecurityPerimeterConfigurationsClient.BeginReconcile
//     method.
func (client *NetworkSecurityPerimeterConfigurationsClient) BeginReconcile(ctx context.Context, resourceGroupName string, accountName string, networkSecurityPerimeterConfigurationName string, options *NetworkSecurityPerimeterConfigurationsClientBeginReconcileOptions) (*runtime.Poller[NetworkSecurityPerimeterConfigurationsClientReconcileResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.reconcile(ctx, resourceGroupName, accountName, networkSecurityPerimeterConfigurationName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NetworkSecurityPerimeterConfigurationsClientReconcileResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[NetworkSecurityPerimeterConfigurationsClientReconcileResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Reconcile - Refreshes any information about the association.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-15-preview
func (client *NetworkSecurityPerimeterConfigurationsClient) reconcile(ctx context.Context, resourceGroupName string, accountName string, networkSecurityPerimeterConfigurationName string, options *NetworkSecurityPerimeterConfigurationsClientBeginReconcileOptions) (*http.Response, error) {
	var err error
	const operationName = "NetworkSecurityPerimeterConfigurationsClient.BeginReconcile"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.reconcileCreateRequest(ctx, resourceGroupName, accountName, networkSecurityPerimeterConfigurationName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// reconcileCreateRequest creates the Reconcile request.
func (client *NetworkSecurityPerimeterConfigurationsClient) reconcileCreateRequest(ctx context.Context, resourceGroupName string, accountName string, networkSecurityPerimeterConfigurationName string, options *NetworkSecurityPerimeterConfigurationsClientBeginReconcileOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/networkSecurityPerimeterConfigurations/{networkSecurityPerimeterConfigurationName}/reconcile"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if networkSecurityPerimeterConfigurationName == "" {
		return nil, errors.New("parameter networkSecurityPerimeterConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkSecurityPerimeterConfigurationName}", url.PathEscape(networkSecurityPerimeterConfigurationName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
