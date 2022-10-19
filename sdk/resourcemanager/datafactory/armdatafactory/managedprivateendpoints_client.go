//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdatafactory

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ManagedPrivateEndpointsClient contains the methods for the ManagedPrivateEndpoints group.
// Don't use this type directly, use NewManagedPrivateEndpointsClient() instead.
type ManagedPrivateEndpointsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewManagedPrivateEndpointsClient creates a new instance of ManagedPrivateEndpointsClient with the specified values.
// subscriptionID - The subscription identifier.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewManagedPrivateEndpointsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ManagedPrivateEndpointsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &ManagedPrivateEndpointsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a managed private endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-06-01
// resourceGroupName - The resource group name.
// factoryName - The factory name.
// managedVirtualNetworkName - Managed virtual network name
// managedPrivateEndpointName - Managed private endpoint name
// managedPrivateEndpoint - Managed private endpoint resource definition.
// options - ManagedPrivateEndpointsClientCreateOrUpdateOptions contains the optional parameters for the ManagedPrivateEndpointsClient.CreateOrUpdate
// method.
func (client *ManagedPrivateEndpointsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, factoryName string, managedVirtualNetworkName string, managedPrivateEndpointName string, managedPrivateEndpoint ManagedPrivateEndpointResource, options *ManagedPrivateEndpointsClientCreateOrUpdateOptions) (ManagedPrivateEndpointsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, factoryName, managedVirtualNetworkName, managedPrivateEndpointName, managedPrivateEndpoint, options)
	if err != nil {
		return ManagedPrivateEndpointsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagedPrivateEndpointsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagedPrivateEndpointsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ManagedPrivateEndpointsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, managedVirtualNetworkName string, managedPrivateEndpointName string, managedPrivateEndpoint ManagedPrivateEndpointResource, options *ManagedPrivateEndpointsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/managedVirtualNetworks/{managedVirtualNetworkName}/managedPrivateEndpoints/{managedPrivateEndpointName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	if managedVirtualNetworkName == "" {
		return nil, errors.New("parameter managedVirtualNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedVirtualNetworkName}", url.PathEscape(managedVirtualNetworkName))
	if managedPrivateEndpointName == "" {
		return nil, errors.New("parameter managedPrivateEndpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedPrivateEndpointName}", url.PathEscape(managedPrivateEndpointName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, managedPrivateEndpoint)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ManagedPrivateEndpointsClient) createOrUpdateHandleResponse(resp *http.Response) (ManagedPrivateEndpointsClientCreateOrUpdateResponse, error) {
	result := ManagedPrivateEndpointsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedPrivateEndpointResource); err != nil {
		return ManagedPrivateEndpointsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a managed private endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-06-01
// resourceGroupName - The resource group name.
// factoryName - The factory name.
// managedVirtualNetworkName - Managed virtual network name
// managedPrivateEndpointName - Managed private endpoint name
// options - ManagedPrivateEndpointsClientDeleteOptions contains the optional parameters for the ManagedPrivateEndpointsClient.Delete
// method.
func (client *ManagedPrivateEndpointsClient) Delete(ctx context.Context, resourceGroupName string, factoryName string, managedVirtualNetworkName string, managedPrivateEndpointName string, options *ManagedPrivateEndpointsClientDeleteOptions) (ManagedPrivateEndpointsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, factoryName, managedVirtualNetworkName, managedPrivateEndpointName, options)
	if err != nil {
		return ManagedPrivateEndpointsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagedPrivateEndpointsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ManagedPrivateEndpointsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return ManagedPrivateEndpointsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ManagedPrivateEndpointsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, managedVirtualNetworkName string, managedPrivateEndpointName string, options *ManagedPrivateEndpointsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/managedVirtualNetworks/{managedVirtualNetworkName}/managedPrivateEndpoints/{managedPrivateEndpointName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	if managedVirtualNetworkName == "" {
		return nil, errors.New("parameter managedVirtualNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedVirtualNetworkName}", url.PathEscape(managedVirtualNetworkName))
	if managedPrivateEndpointName == "" {
		return nil, errors.New("parameter managedPrivateEndpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedPrivateEndpointName}", url.PathEscape(managedPrivateEndpointName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a managed private endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-06-01
// resourceGroupName - The resource group name.
// factoryName - The factory name.
// managedVirtualNetworkName - Managed virtual network name
// managedPrivateEndpointName - Managed private endpoint name
// options - ManagedPrivateEndpointsClientGetOptions contains the optional parameters for the ManagedPrivateEndpointsClient.Get
// method.
func (client *ManagedPrivateEndpointsClient) Get(ctx context.Context, resourceGroupName string, factoryName string, managedVirtualNetworkName string, managedPrivateEndpointName string, options *ManagedPrivateEndpointsClientGetOptions) (ManagedPrivateEndpointsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, factoryName, managedVirtualNetworkName, managedPrivateEndpointName, options)
	if err != nil {
		return ManagedPrivateEndpointsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagedPrivateEndpointsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagedPrivateEndpointsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ManagedPrivateEndpointsClient) getCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, managedVirtualNetworkName string, managedPrivateEndpointName string, options *ManagedPrivateEndpointsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/managedVirtualNetworks/{managedVirtualNetworkName}/managedPrivateEndpoints/{managedPrivateEndpointName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	if managedVirtualNetworkName == "" {
		return nil, errors.New("parameter managedVirtualNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedVirtualNetworkName}", url.PathEscape(managedVirtualNetworkName))
	if managedPrivateEndpointName == "" {
		return nil, errors.New("parameter managedPrivateEndpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedPrivateEndpointName}", url.PathEscape(managedPrivateEndpointName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header["If-None-Match"] = []string{*options.IfNoneMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ManagedPrivateEndpointsClient) getHandleResponse(resp *http.Response) (ManagedPrivateEndpointsClientGetResponse, error) {
	result := ManagedPrivateEndpointsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedPrivateEndpointResource); err != nil {
		return ManagedPrivateEndpointsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByFactoryPager - Lists managed private endpoints.
// Generated from API version 2018-06-01
// resourceGroupName - The resource group name.
// factoryName - The factory name.
// managedVirtualNetworkName - Managed virtual network name
// options - ManagedPrivateEndpointsClientListByFactoryOptions contains the optional parameters for the ManagedPrivateEndpointsClient.ListByFactory
// method.
func (client *ManagedPrivateEndpointsClient) NewListByFactoryPager(resourceGroupName string, factoryName string, managedVirtualNetworkName string, options *ManagedPrivateEndpointsClientListByFactoryOptions) *runtime.Pager[ManagedPrivateEndpointsClientListByFactoryResponse] {
	return runtime.NewPager(runtime.PagingHandler[ManagedPrivateEndpointsClientListByFactoryResponse]{
		More: func(page ManagedPrivateEndpointsClientListByFactoryResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ManagedPrivateEndpointsClientListByFactoryResponse) (ManagedPrivateEndpointsClientListByFactoryResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByFactoryCreateRequest(ctx, resourceGroupName, factoryName, managedVirtualNetworkName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ManagedPrivateEndpointsClientListByFactoryResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ManagedPrivateEndpointsClientListByFactoryResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ManagedPrivateEndpointsClientListByFactoryResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByFactoryHandleResponse(resp)
		},
	})
}

// listByFactoryCreateRequest creates the ListByFactory request.
func (client *ManagedPrivateEndpointsClient) listByFactoryCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, managedVirtualNetworkName string, options *ManagedPrivateEndpointsClientListByFactoryOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/managedVirtualNetworks/{managedVirtualNetworkName}/managedPrivateEndpoints"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	if managedVirtualNetworkName == "" {
		return nil, errors.New("parameter managedVirtualNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedVirtualNetworkName}", url.PathEscape(managedVirtualNetworkName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByFactoryHandleResponse handles the ListByFactory response.
func (client *ManagedPrivateEndpointsClient) listByFactoryHandleResponse(resp *http.Response) (ManagedPrivateEndpointsClientListByFactoryResponse, error) {
	result := ManagedPrivateEndpointsClientListByFactoryResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedPrivateEndpointListResponse); err != nil {
		return ManagedPrivateEndpointsClientListByFactoryResponse{}, err
	}
	return result, nil
}
