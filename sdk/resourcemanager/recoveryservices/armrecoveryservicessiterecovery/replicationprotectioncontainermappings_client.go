//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicessiterecovery

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

// ReplicationProtectionContainerMappingsClient contains the methods for the ReplicationProtectionContainerMappings group.
// Don't use this type directly, use NewReplicationProtectionContainerMappingsClient() instead.
type ReplicationProtectionContainerMappingsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewReplicationProtectionContainerMappingsClient creates a new instance of ReplicationProtectionContainerMappingsClient with the specified values.
//   - subscriptionID - The subscription Id.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewReplicationProtectionContainerMappingsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ReplicationProtectionContainerMappingsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ReplicationProtectionContainerMappingsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - The operation to create a protection container mapping.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - resourceName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - fabricName - Fabric name.
//   - protectionContainerName - Protection container name.
//   - mappingName - Protection container mapping name.
//   - creationInput - Mapping creation input.
//   - options - ReplicationProtectionContainerMappingsClientBeginCreateOptions contains the optional parameters for the ReplicationProtectionContainerMappingsClient.BeginCreate
//     method.
func (client *ReplicationProtectionContainerMappingsClient) BeginCreate(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, protectionContainerName string, mappingName string, creationInput CreateProtectionContainerMappingInput, options *ReplicationProtectionContainerMappingsClientBeginCreateOptions) (*runtime.Poller[ReplicationProtectionContainerMappingsClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceName, resourceGroupName, fabricName, protectionContainerName, mappingName, creationInput, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ReplicationProtectionContainerMappingsClientCreateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ReplicationProtectionContainerMappingsClientCreateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Create - The operation to create a protection container mapping.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
func (client *ReplicationProtectionContainerMappingsClient) create(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, protectionContainerName string, mappingName string, creationInput CreateProtectionContainerMappingInput, options *ReplicationProtectionContainerMappingsClientBeginCreateOptions) (*http.Response, error) {
	var err error
	const operationName = "ReplicationProtectionContainerMappingsClient.BeginCreate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceName, resourceGroupName, fabricName, protectionContainerName, mappingName, creationInput, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createCreateRequest creates the Create request.
func (client *ReplicationProtectionContainerMappingsClient) createCreateRequest(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, protectionContainerName string, mappingName string, creationInput CreateProtectionContainerMappingInput, options *ReplicationProtectionContainerMappingsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectionContainerMappings/{mappingName}"
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if fabricName == "" {
		return nil, errors.New("parameter fabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fabricName}", url.PathEscape(fabricName))
	if protectionContainerName == "" {
		return nil, errors.New("parameter protectionContainerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{protectionContainerName}", url.PathEscape(protectionContainerName))
	if mappingName == "" {
		return nil, errors.New("parameter mappingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mappingName}", url.PathEscape(mappingName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, creationInput); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - The operation to delete or remove a protection container mapping.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - resourceName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - fabricName - Fabric name.
//   - protectionContainerName - Protection container name.
//   - mappingName - Protection container mapping name.
//   - removalInput - Removal input.
//   - options - ReplicationProtectionContainerMappingsClientBeginDeleteOptions contains the optional parameters for the ReplicationProtectionContainerMappingsClient.BeginDelete
//     method.
func (client *ReplicationProtectionContainerMappingsClient) BeginDelete(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, protectionContainerName string, mappingName string, removalInput RemoveProtectionContainerMappingInput, options *ReplicationProtectionContainerMappingsClientBeginDeleteOptions) (*runtime.Poller[ReplicationProtectionContainerMappingsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceName, resourceGroupName, fabricName, protectionContainerName, mappingName, removalInput, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ReplicationProtectionContainerMappingsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ReplicationProtectionContainerMappingsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - The operation to delete or remove a protection container mapping.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
func (client *ReplicationProtectionContainerMappingsClient) deleteOperation(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, protectionContainerName string, mappingName string, removalInput RemoveProtectionContainerMappingInput, options *ReplicationProtectionContainerMappingsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "ReplicationProtectionContainerMappingsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceName, resourceGroupName, fabricName, protectionContainerName, mappingName, removalInput, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ReplicationProtectionContainerMappingsClient) deleteCreateRequest(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, protectionContainerName string, mappingName string, removalInput RemoveProtectionContainerMappingInput, options *ReplicationProtectionContainerMappingsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectionContainerMappings/{mappingName}/remove"
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if fabricName == "" {
		return nil, errors.New("parameter fabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fabricName}", url.PathEscape(fabricName))
	if protectionContainerName == "" {
		return nil, errors.New("parameter protectionContainerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{protectionContainerName}", url.PathEscape(protectionContainerName))
	if mappingName == "" {
		return nil, errors.New("parameter mappingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mappingName}", url.PathEscape(mappingName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if err := runtime.MarshalAsJSON(req, removalInput); err != nil {
		return nil, err
	}
	return req, nil
}

// Get - Gets the details of a protection container mapping.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - resourceName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - fabricName - Fabric name.
//   - protectionContainerName - Protection container name.
//   - mappingName - Protection Container mapping name.
//   - options - ReplicationProtectionContainerMappingsClientGetOptions contains the optional parameters for the ReplicationProtectionContainerMappingsClient.Get
//     method.
func (client *ReplicationProtectionContainerMappingsClient) Get(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, protectionContainerName string, mappingName string, options *ReplicationProtectionContainerMappingsClientGetOptions) (ReplicationProtectionContainerMappingsClientGetResponse, error) {
	var err error
	const operationName = "ReplicationProtectionContainerMappingsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceName, resourceGroupName, fabricName, protectionContainerName, mappingName, options)
	if err != nil {
		return ReplicationProtectionContainerMappingsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ReplicationProtectionContainerMappingsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ReplicationProtectionContainerMappingsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ReplicationProtectionContainerMappingsClient) getCreateRequest(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, protectionContainerName string, mappingName string, options *ReplicationProtectionContainerMappingsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectionContainerMappings/{mappingName}"
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if fabricName == "" {
		return nil, errors.New("parameter fabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fabricName}", url.PathEscape(fabricName))
	if protectionContainerName == "" {
		return nil, errors.New("parameter protectionContainerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{protectionContainerName}", url.PathEscape(protectionContainerName))
	if mappingName == "" {
		return nil, errors.New("parameter mappingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mappingName}", url.PathEscape(mappingName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ReplicationProtectionContainerMappingsClient) getHandleResponse(resp *http.Response) (ReplicationProtectionContainerMappingsClientGetResponse, error) {
	result := ReplicationProtectionContainerMappingsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProtectionContainerMapping); err != nil {
		return ReplicationProtectionContainerMappingsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists the protection container mappings in the vault.
//
// Generated from API version 2024-04-01
//   - resourceName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - options - ReplicationProtectionContainerMappingsClientListOptions contains the optional parameters for the ReplicationProtectionContainerMappingsClient.NewListPager
//     method.
func (client *ReplicationProtectionContainerMappingsClient) NewListPager(resourceName string, resourceGroupName string, options *ReplicationProtectionContainerMappingsClientListOptions) *runtime.Pager[ReplicationProtectionContainerMappingsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ReplicationProtectionContainerMappingsClientListResponse]{
		More: func(page ReplicationProtectionContainerMappingsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ReplicationProtectionContainerMappingsClientListResponse) (ReplicationProtectionContainerMappingsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ReplicationProtectionContainerMappingsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceName, resourceGroupName, options)
			}, nil)
			if err != nil {
				return ReplicationProtectionContainerMappingsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ReplicationProtectionContainerMappingsClient) listCreateRequest(ctx context.Context, resourceName string, resourceGroupName string, options *ReplicationProtectionContainerMappingsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationProtectionContainerMappings"
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ReplicationProtectionContainerMappingsClient) listHandleResponse(resp *http.Response) (ReplicationProtectionContainerMappingsClientListResponse, error) {
	result := ReplicationProtectionContainerMappingsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProtectionContainerMappingCollection); err != nil {
		return ReplicationProtectionContainerMappingsClientListResponse{}, err
	}
	return result, nil
}

// NewListByReplicationProtectionContainersPager - Lists the protection container mappings for a protection container.
//
// Generated from API version 2024-04-01
//   - resourceName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - fabricName - Fabric name.
//   - protectionContainerName - Protection container name.
//   - options - ReplicationProtectionContainerMappingsClientListByReplicationProtectionContainersOptions contains the optional
//     parameters for the ReplicationProtectionContainerMappingsClient.NewListByReplicationProtectionContainersPager method.
func (client *ReplicationProtectionContainerMappingsClient) NewListByReplicationProtectionContainersPager(resourceName string, resourceGroupName string, fabricName string, protectionContainerName string, options *ReplicationProtectionContainerMappingsClientListByReplicationProtectionContainersOptions) *runtime.Pager[ReplicationProtectionContainerMappingsClientListByReplicationProtectionContainersResponse] {
	return runtime.NewPager(runtime.PagingHandler[ReplicationProtectionContainerMappingsClientListByReplicationProtectionContainersResponse]{
		More: func(page ReplicationProtectionContainerMappingsClientListByReplicationProtectionContainersResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ReplicationProtectionContainerMappingsClientListByReplicationProtectionContainersResponse) (ReplicationProtectionContainerMappingsClientListByReplicationProtectionContainersResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ReplicationProtectionContainerMappingsClient.NewListByReplicationProtectionContainersPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByReplicationProtectionContainersCreateRequest(ctx, resourceName, resourceGroupName, fabricName, protectionContainerName, options)
			}, nil)
			if err != nil {
				return ReplicationProtectionContainerMappingsClientListByReplicationProtectionContainersResponse{}, err
			}
			return client.listByReplicationProtectionContainersHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByReplicationProtectionContainersCreateRequest creates the ListByReplicationProtectionContainers request.
func (client *ReplicationProtectionContainerMappingsClient) listByReplicationProtectionContainersCreateRequest(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, protectionContainerName string, options *ReplicationProtectionContainerMappingsClientListByReplicationProtectionContainersOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectionContainerMappings"
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if fabricName == "" {
		return nil, errors.New("parameter fabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fabricName}", url.PathEscape(fabricName))
	if protectionContainerName == "" {
		return nil, errors.New("parameter protectionContainerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{protectionContainerName}", url.PathEscape(protectionContainerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByReplicationProtectionContainersHandleResponse handles the ListByReplicationProtectionContainers response.
func (client *ReplicationProtectionContainerMappingsClient) listByReplicationProtectionContainersHandleResponse(resp *http.Response) (ReplicationProtectionContainerMappingsClientListByReplicationProtectionContainersResponse, error) {
	result := ReplicationProtectionContainerMappingsClientListByReplicationProtectionContainersResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProtectionContainerMappingCollection); err != nil {
		return ReplicationProtectionContainerMappingsClientListByReplicationProtectionContainersResponse{}, err
	}
	return result, nil
}

// BeginPurge - The operation to purge(force delete) a protection container mapping.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - resourceName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - fabricName - Fabric name.
//   - protectionContainerName - Protection container name.
//   - mappingName - Protection container mapping name.
//   - options - ReplicationProtectionContainerMappingsClientBeginPurgeOptions contains the optional parameters for the ReplicationProtectionContainerMappingsClient.BeginPurge
//     method.
func (client *ReplicationProtectionContainerMappingsClient) BeginPurge(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, protectionContainerName string, mappingName string, options *ReplicationProtectionContainerMappingsClientBeginPurgeOptions) (*runtime.Poller[ReplicationProtectionContainerMappingsClientPurgeResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.purge(ctx, resourceName, resourceGroupName, fabricName, protectionContainerName, mappingName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ReplicationProtectionContainerMappingsClientPurgeResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ReplicationProtectionContainerMappingsClientPurgeResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Purge - The operation to purge(force delete) a protection container mapping.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
func (client *ReplicationProtectionContainerMappingsClient) purge(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, protectionContainerName string, mappingName string, options *ReplicationProtectionContainerMappingsClientBeginPurgeOptions) (*http.Response, error) {
	var err error
	const operationName = "ReplicationProtectionContainerMappingsClient.BeginPurge"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.purgeCreateRequest(ctx, resourceName, resourceGroupName, fabricName, protectionContainerName, mappingName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// purgeCreateRequest creates the Purge request.
func (client *ReplicationProtectionContainerMappingsClient) purgeCreateRequest(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, protectionContainerName string, mappingName string, options *ReplicationProtectionContainerMappingsClientBeginPurgeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectionContainerMappings/{mappingName}"
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if fabricName == "" {
		return nil, errors.New("parameter fabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fabricName}", url.PathEscape(fabricName))
	if protectionContainerName == "" {
		return nil, errors.New("parameter protectionContainerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{protectionContainerName}", url.PathEscape(protectionContainerName))
	if mappingName == "" {
		return nil, errors.New("parameter mappingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mappingName}", url.PathEscape(mappingName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// BeginUpdate - The operation to update protection container mapping.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - resourceName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - fabricName - Fabric name.
//   - protectionContainerName - Protection container name.
//   - mappingName - Protection container mapping name.
//   - updateInput - Mapping update input.
//   - options - ReplicationProtectionContainerMappingsClientBeginUpdateOptions contains the optional parameters for the ReplicationProtectionContainerMappingsClient.BeginUpdate
//     method.
func (client *ReplicationProtectionContainerMappingsClient) BeginUpdate(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, protectionContainerName string, mappingName string, updateInput UpdateProtectionContainerMappingInput, options *ReplicationProtectionContainerMappingsClientBeginUpdateOptions) (*runtime.Poller[ReplicationProtectionContainerMappingsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceName, resourceGroupName, fabricName, protectionContainerName, mappingName, updateInput, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ReplicationProtectionContainerMappingsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ReplicationProtectionContainerMappingsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - The operation to update protection container mapping.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
func (client *ReplicationProtectionContainerMappingsClient) update(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, protectionContainerName string, mappingName string, updateInput UpdateProtectionContainerMappingInput, options *ReplicationProtectionContainerMappingsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "ReplicationProtectionContainerMappingsClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceName, resourceGroupName, fabricName, protectionContainerName, mappingName, updateInput, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateCreateRequest creates the Update request.
func (client *ReplicationProtectionContainerMappingsClient) updateCreateRequest(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, protectionContainerName string, mappingName string, updateInput UpdateProtectionContainerMappingInput, options *ReplicationProtectionContainerMappingsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectionContainerMappings/{mappingName}"
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if fabricName == "" {
		return nil, errors.New("parameter fabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fabricName}", url.PathEscape(fabricName))
	if protectionContainerName == "" {
		return nil, errors.New("parameter protectionContainerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{protectionContainerName}", url.PathEscape(protectionContainerName))
	if mappingName == "" {
		return nil, errors.New("parameter mappingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mappingName}", url.PathEscape(mappingName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, updateInput); err != nil {
		return nil, err
	}
	return req, nil
}
