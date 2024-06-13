//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappcontainers

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

// ContainerAppsPatchesClient contains the methods for the ContainerAppsPatches group.
// Don't use this type directly, use NewContainerAppsPatchesClient() instead.
type ContainerAppsPatchesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewContainerAppsPatchesClient creates a new instance of ContainerAppsPatchesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewContainerAppsPatchesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ContainerAppsPatchesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ContainerAppsPatchesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginApply - Apply a Container Apps Patch resource with patch name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - containerAppName - Name of the Container App the Patch is associated.
//   - patchName - The name of the patch
//   - options - ContainerAppsPatchesClientBeginApplyOptions contains the optional parameters for the ContainerAppsPatchesClient.BeginApply
//     method.
func (client *ContainerAppsPatchesClient) BeginApply(ctx context.Context, resourceGroupName string, containerAppName string, patchName string, options *ContainerAppsPatchesClientBeginApplyOptions) (*runtime.Poller[ContainerAppsPatchesClientApplyResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.apply(ctx, resourceGroupName, containerAppName, patchName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ContainerAppsPatchesClientApplyResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ContainerAppsPatchesClientApplyResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Apply - Apply a Container Apps Patch resource with patch name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-02-preview
func (client *ContainerAppsPatchesClient) apply(ctx context.Context, resourceGroupName string, containerAppName string, patchName string, options *ContainerAppsPatchesClientBeginApplyOptions) (*http.Response, error) {
	var err error
	const operationName = "ContainerAppsPatchesClient.BeginApply"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.applyCreateRequest(ctx, resourceGroupName, containerAppName, patchName, options)
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

// applyCreateRequest creates the Apply request.
func (client *ContainerAppsPatchesClient) applyCreateRequest(ctx context.Context, resourceGroupName string, containerAppName string, patchName string, options *ContainerAppsPatchesClientBeginApplyOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/containerApps/{containerAppName}/patches/{patchName}/apply"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerAppName == "" {
		return nil, errors.New("parameter containerAppName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerAppName}", url.PathEscape(containerAppName))
	if patchName == "" {
		return nil, errors.New("parameter patchName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{patchName}", url.PathEscape(patchName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginDelete - Delete specific Container Apps Patch by patch name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - containerAppName - Name of the Container App the Patch is associated.
//   - patchName - The name of the patch
//   - options - ContainerAppsPatchesClientBeginDeleteOptions contains the optional parameters for the ContainerAppsPatchesClient.BeginDelete
//     method.
func (client *ContainerAppsPatchesClient) BeginDelete(ctx context.Context, resourceGroupName string, containerAppName string, patchName string, options *ContainerAppsPatchesClientBeginDeleteOptions) (*runtime.Poller[ContainerAppsPatchesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, containerAppName, patchName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ContainerAppsPatchesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ContainerAppsPatchesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete specific Container Apps Patch by patch name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-02-preview
func (client *ContainerAppsPatchesClient) deleteOperation(ctx context.Context, resourceGroupName string, containerAppName string, patchName string, options *ContainerAppsPatchesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "ContainerAppsPatchesClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, containerAppName, patchName, options)
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
func (client *ContainerAppsPatchesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, containerAppName string, patchName string, options *ContainerAppsPatchesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/containerApps/{containerAppName}/patches/{patchName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerAppName == "" {
		return nil, errors.New("parameter containerAppName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerAppName}", url.PathEscape(containerAppName))
	if patchName == "" {
		return nil, errors.New("parameter patchName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{patchName}", url.PathEscape(patchName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get details for specific Container Apps Patch by patch name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - containerAppName - Name of the Container App the Patch is associated.
//   - patchName - The name of the patch
//   - options - ContainerAppsPatchesClientGetOptions contains the optional parameters for the ContainerAppsPatchesClient.Get
//     method.
func (client *ContainerAppsPatchesClient) Get(ctx context.Context, resourceGroupName string, containerAppName string, patchName string, options *ContainerAppsPatchesClientGetOptions) (ContainerAppsPatchesClientGetResponse, error) {
	var err error
	const operationName = "ContainerAppsPatchesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, containerAppName, patchName, options)
	if err != nil {
		return ContainerAppsPatchesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ContainerAppsPatchesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ContainerAppsPatchesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ContainerAppsPatchesClient) getCreateRequest(ctx context.Context, resourceGroupName string, containerAppName string, patchName string, options *ContainerAppsPatchesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/containerApps/{containerAppName}/patches/{patchName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerAppName == "" {
		return nil, errors.New("parameter containerAppName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerAppName}", url.PathEscape(containerAppName))
	if patchName == "" {
		return nil, errors.New("parameter patchName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{patchName}", url.PathEscape(patchName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ContainerAppsPatchesClient) getHandleResponse(resp *http.Response) (ContainerAppsPatchesClientGetResponse, error) {
	result := ContainerAppsPatchesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ContainerAppsPatchResource); err != nil {
		return ContainerAppsPatchesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByContainerAppPager - List Container Apps Patch resources by ContainerApp.
//
// Generated from API version 2024-02-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - containerAppName - Name of the Container App the Patch is associated.
//   - options - ContainerAppsPatchesClientListByContainerAppOptions contains the optional parameters for the ContainerAppsPatchesClient.NewListByContainerAppPager
//     method.
func (client *ContainerAppsPatchesClient) NewListByContainerAppPager(resourceGroupName string, containerAppName string, options *ContainerAppsPatchesClientListByContainerAppOptions) *runtime.Pager[ContainerAppsPatchesClientListByContainerAppResponse] {
	return runtime.NewPager(runtime.PagingHandler[ContainerAppsPatchesClientListByContainerAppResponse]{
		More: func(page ContainerAppsPatchesClientListByContainerAppResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ContainerAppsPatchesClientListByContainerAppResponse) (ContainerAppsPatchesClientListByContainerAppResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ContainerAppsPatchesClient.NewListByContainerAppPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByContainerAppCreateRequest(ctx, resourceGroupName, containerAppName, options)
			}, nil)
			if err != nil {
				return ContainerAppsPatchesClientListByContainerAppResponse{}, err
			}
			return client.listByContainerAppHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByContainerAppCreateRequest creates the ListByContainerApp request.
func (client *ContainerAppsPatchesClient) listByContainerAppCreateRequest(ctx context.Context, resourceGroupName string, containerAppName string, options *ContainerAppsPatchesClientListByContainerAppOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/containerApps/{containerAppName}/patches"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerAppName == "" {
		return nil, errors.New("parameter containerAppName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerAppName}", url.PathEscape(containerAppName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2024-02-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByContainerAppHandleResponse handles the ListByContainerApp response.
func (client *ContainerAppsPatchesClient) listByContainerAppHandleResponse(resp *http.Response) (ContainerAppsPatchesClientListByContainerAppResponse, error) {
	result := ContainerAppsPatchesClientListByContainerAppResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PatchCollection); err != nil {
		return ContainerAppsPatchesClientListByContainerAppResponse{}, err
	}
	return result, nil
}

// BeginSkipConfigure - Configure the Container Apps Patch skip option by patch name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - containerAppName - Name of the Container App the Patch is associated.
//   - patchName - The name of the patch
//   - patchSkipConfig - Configure patcher to skip a patch or not.
//   - options - ContainerAppsPatchesClientBeginSkipConfigureOptions contains the optional parameters for the ContainerAppsPatchesClient.BeginSkipConfigure
//     method.
func (client *ContainerAppsPatchesClient) BeginSkipConfigure(ctx context.Context, resourceGroupName string, containerAppName string, patchName string, patchSkipConfig PatchSkipConfig, options *ContainerAppsPatchesClientBeginSkipConfigureOptions) (*runtime.Poller[ContainerAppsPatchesClientSkipConfigureResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.skipConfigure(ctx, resourceGroupName, containerAppName, patchName, patchSkipConfig, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ContainerAppsPatchesClientSkipConfigureResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ContainerAppsPatchesClientSkipConfigureResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// SkipConfigure - Configure the Container Apps Patch skip option by patch name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-02-preview
func (client *ContainerAppsPatchesClient) skipConfigure(ctx context.Context, resourceGroupName string, containerAppName string, patchName string, patchSkipConfig PatchSkipConfig, options *ContainerAppsPatchesClientBeginSkipConfigureOptions) (*http.Response, error) {
	var err error
	const operationName = "ContainerAppsPatchesClient.BeginSkipConfigure"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.skipConfigureCreateRequest(ctx, resourceGroupName, containerAppName, patchName, patchSkipConfig, options)
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

// skipConfigureCreateRequest creates the SkipConfigure request.
func (client *ContainerAppsPatchesClient) skipConfigureCreateRequest(ctx context.Context, resourceGroupName string, containerAppName string, patchName string, patchSkipConfig PatchSkipConfig, options *ContainerAppsPatchesClientBeginSkipConfigureOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/containerApps/{containerAppName}/patches/{patchName}/skipConfig"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerAppName == "" {
		return nil, errors.New("parameter containerAppName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerAppName}", url.PathEscape(containerAppName))
	if patchName == "" {
		return nil, errors.New("parameter patchName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{patchName}", url.PathEscape(patchName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, patchSkipConfig); err != nil {
		return nil, err
	}
	return req, nil
}
