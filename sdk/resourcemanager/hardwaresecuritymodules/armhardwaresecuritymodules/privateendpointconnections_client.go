//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhardwaresecuritymodules

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

// PrivateEndpointConnectionsClient contains the methods for the PrivateEndpointConnections group.
// Don't use this type directly, use NewPrivateEndpointConnectionsClient() instead.
type PrivateEndpointConnectionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewPrivateEndpointConnectionsClient creates a new instance of PrivateEndpointConnectionsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPrivateEndpointConnectionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PrivateEndpointConnectionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PrivateEndpointConnectionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListByCloudHsmClusterPager - The List operation gets information about the private endpoint connections associated with
// the Cloud HSM Cluster
//
// Generated from API version 2024-06-30
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - cloudHsmClusterName - The name of the Cloud HSM Cluster within the specified resource group. Cloud HSM Cluster names must
//     be between 3 and 24 characters in length.
//   - options - PrivateEndpointConnectionsClientListByCloudHsmClusterOptions contains the optional parameters for the PrivateEndpointConnectionsClient.NewListByCloudHsmClusterPager
//     method.
func (client *PrivateEndpointConnectionsClient) NewListByCloudHsmClusterPager(resourceGroupName string, cloudHsmClusterName string, options *PrivateEndpointConnectionsClientListByCloudHsmClusterOptions) *runtime.Pager[PrivateEndpointConnectionsClientListByCloudHsmClusterResponse] {
	return runtime.NewPager(runtime.PagingHandler[PrivateEndpointConnectionsClientListByCloudHsmClusterResponse]{
		More: func(page PrivateEndpointConnectionsClientListByCloudHsmClusterResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PrivateEndpointConnectionsClientListByCloudHsmClusterResponse) (PrivateEndpointConnectionsClientListByCloudHsmClusterResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "PrivateEndpointConnectionsClient.NewListByCloudHsmClusterPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByCloudHsmClusterCreateRequest(ctx, resourceGroupName, cloudHsmClusterName, options)
			}, nil)
			if err != nil {
				return PrivateEndpointConnectionsClientListByCloudHsmClusterResponse{}, err
			}
			return client.listByCloudHsmClusterHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByCloudHsmClusterCreateRequest creates the ListByCloudHsmCluster request.
func (client *PrivateEndpointConnectionsClient) listByCloudHsmClusterCreateRequest(ctx context.Context, resourceGroupName string, cloudHsmClusterName string, options *PrivateEndpointConnectionsClientListByCloudHsmClusterOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HardwareSecurityModules/cloudHsmClusters/{cloudHsmClusterName}/privateEndpointConnections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cloudHsmClusterName == "" {
		return nil, errors.New("parameter cloudHsmClusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cloudHsmClusterName}", url.PathEscape(cloudHsmClusterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-30")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByCloudHsmClusterHandleResponse handles the ListByCloudHsmCluster response.
func (client *PrivateEndpointConnectionsClient) listByCloudHsmClusterHandleResponse(resp *http.Response) (PrivateEndpointConnectionsClientListByCloudHsmClusterResponse, error) {
	result := PrivateEndpointConnectionsClientListByCloudHsmClusterResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateEndpointConnectionListResult); err != nil {
		return PrivateEndpointConnectionsClientListByCloudHsmClusterResponse{}, err
	}
	return result, nil
}
