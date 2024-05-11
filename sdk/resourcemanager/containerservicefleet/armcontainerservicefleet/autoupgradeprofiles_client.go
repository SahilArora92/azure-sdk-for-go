//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcontainerservicefleet

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

// AutoUpgradeProfilesClient contains the methods for the AutoUpgradeProfiles group.
// Don't use this type directly, use NewAutoUpgradeProfilesClient() instead.
type AutoUpgradeProfilesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAutoUpgradeProfilesClient creates a new instance of AutoUpgradeProfilesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAutoUpgradeProfilesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AutoUpgradeProfilesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AutoUpgradeProfilesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create a AutoUpgradeProfile
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - fleetName - The name of the Fleet resource.
//   - autoUpgradeProfileName - The name of the AutoUpgradeProfile resource.
//   - resource - Resource create parameters.
//   - options - AutoUpgradeProfilesClientBeginCreateOrUpdateOptions contains the optional parameters for the AutoUpgradeProfilesClient.BeginCreateOrUpdate
//     method.
func (client *AutoUpgradeProfilesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, fleetName string, autoUpgradeProfileName string, resource AutoUpgradeProfile, options *AutoUpgradeProfilesClientBeginCreateOrUpdateOptions) (*runtime.Poller[AutoUpgradeProfilesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, fleetName, autoUpgradeProfileName, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AutoUpgradeProfilesClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[AutoUpgradeProfilesClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create a AutoUpgradeProfile
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-02-preview
func (client *AutoUpgradeProfilesClient) createOrUpdate(ctx context.Context, resourceGroupName string, fleetName string, autoUpgradeProfileName string, resource AutoUpgradeProfile, options *AutoUpgradeProfilesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "AutoUpgradeProfilesClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, fleetName, autoUpgradeProfileName, resource, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *AutoUpgradeProfilesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, fleetName string, autoUpgradeProfileName string, resource AutoUpgradeProfile, options *AutoUpgradeProfilesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/fleets/{fleetName}/autoUpgradeProfiles/{autoUpgradeProfileName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if fleetName == "" {
		return nil, errors.New("parameter fleetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fleetName}", url.PathEscape(fleetName))
	if autoUpgradeProfileName == "" {
		return nil, errors.New("parameter autoUpgradeProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{autoUpgradeProfileName}", url.PathEscape(autoUpgradeProfileName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header["If-None-Match"] = []string{*options.IfNoneMatch}
	}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete a AutoUpgradeProfile
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - fleetName - The name of the Fleet resource.
//   - autoUpgradeProfileName - The name of the AutoUpgradeProfile resource.
//   - options - AutoUpgradeProfilesClientBeginDeleteOptions contains the optional parameters for the AutoUpgradeProfilesClient.BeginDelete
//     method.
func (client *AutoUpgradeProfilesClient) BeginDelete(ctx context.Context, resourceGroupName string, fleetName string, autoUpgradeProfileName string, options *AutoUpgradeProfilesClientBeginDeleteOptions) (*runtime.Poller[AutoUpgradeProfilesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, fleetName, autoUpgradeProfileName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AutoUpgradeProfilesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[AutoUpgradeProfilesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete a AutoUpgradeProfile
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-02-preview
func (client *AutoUpgradeProfilesClient) deleteOperation(ctx context.Context, resourceGroupName string, fleetName string, autoUpgradeProfileName string, options *AutoUpgradeProfilesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "AutoUpgradeProfilesClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, fleetName, autoUpgradeProfileName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *AutoUpgradeProfilesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, fleetName string, autoUpgradeProfileName string, options *AutoUpgradeProfilesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/fleets/{fleetName}/autoUpgradeProfiles/{autoUpgradeProfileName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if fleetName == "" {
		return nil, errors.New("parameter fleetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fleetName}", url.PathEscape(fleetName))
	if autoUpgradeProfileName == "" {
		return nil, errors.New("parameter autoUpgradeProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{autoUpgradeProfileName}", url.PathEscape(autoUpgradeProfileName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	return req, nil
}

// Get - Get a AutoUpgradeProfile
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - fleetName - The name of the Fleet resource.
//   - autoUpgradeProfileName - The name of the AutoUpgradeProfile resource.
//   - options - AutoUpgradeProfilesClientGetOptions contains the optional parameters for the AutoUpgradeProfilesClient.Get method.
func (client *AutoUpgradeProfilesClient) Get(ctx context.Context, resourceGroupName string, fleetName string, autoUpgradeProfileName string, options *AutoUpgradeProfilesClientGetOptions) (AutoUpgradeProfilesClientGetResponse, error) {
	var err error
	const operationName = "AutoUpgradeProfilesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, fleetName, autoUpgradeProfileName, options)
	if err != nil {
		return AutoUpgradeProfilesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AutoUpgradeProfilesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AutoUpgradeProfilesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *AutoUpgradeProfilesClient) getCreateRequest(ctx context.Context, resourceGroupName string, fleetName string, autoUpgradeProfileName string, options *AutoUpgradeProfilesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/fleets/{fleetName}/autoUpgradeProfiles/{autoUpgradeProfileName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if fleetName == "" {
		return nil, errors.New("parameter fleetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fleetName}", url.PathEscape(fleetName))
	if autoUpgradeProfileName == "" {
		return nil, errors.New("parameter autoUpgradeProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{autoUpgradeProfileName}", url.PathEscape(autoUpgradeProfileName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AutoUpgradeProfilesClient) getHandleResponse(resp *http.Response) (AutoUpgradeProfilesClientGetResponse, error) {
	result := AutoUpgradeProfilesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AutoUpgradeProfile); err != nil {
		return AutoUpgradeProfilesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByFleetPager - List AutoUpgradeProfile resources by Fleet
//
// Generated from API version 2024-05-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - fleetName - The name of the Fleet resource.
//   - options - AutoUpgradeProfilesClientListByFleetOptions contains the optional parameters for the AutoUpgradeProfilesClient.NewListByFleetPager
//     method.
func (client *AutoUpgradeProfilesClient) NewListByFleetPager(resourceGroupName string, fleetName string, options *AutoUpgradeProfilesClientListByFleetOptions) *runtime.Pager[AutoUpgradeProfilesClientListByFleetResponse] {
	return runtime.NewPager(runtime.PagingHandler[AutoUpgradeProfilesClientListByFleetResponse]{
		More: func(page AutoUpgradeProfilesClientListByFleetResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AutoUpgradeProfilesClientListByFleetResponse) (AutoUpgradeProfilesClientListByFleetResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AutoUpgradeProfilesClient.NewListByFleetPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByFleetCreateRequest(ctx, resourceGroupName, fleetName, options)
			}, nil)
			if err != nil {
				return AutoUpgradeProfilesClientListByFleetResponse{}, err
			}
			return client.listByFleetHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByFleetCreateRequest creates the ListByFleet request.
func (client *AutoUpgradeProfilesClient) listByFleetCreateRequest(ctx context.Context, resourceGroupName string, fleetName string, options *AutoUpgradeProfilesClientListByFleetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/fleets/{fleetName}/autoUpgradeProfiles"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if fleetName == "" {
		return nil, errors.New("parameter fleetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fleetName}", url.PathEscape(fleetName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByFleetHandleResponse handles the ListByFleet response.
func (client *AutoUpgradeProfilesClient) listByFleetHandleResponse(resp *http.Response) (AutoUpgradeProfilesClientListByFleetResponse, error) {
	result := AutoUpgradeProfilesClientListByFleetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AutoUpgradeProfileListResult); err != nil {
		return AutoUpgradeProfilesClientListByFleetResponse{}, err
	}
	return result, nil
}
