//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappservice

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

// KubeEnvironmentsClient contains the methods for the KubeEnvironments group.
// Don't use this type directly, use NewKubeEnvironmentsClient() instead.
type KubeEnvironmentsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewKubeEnvironmentsClient creates a new instance of KubeEnvironmentsClient with the specified values.
//   - subscriptionID - Your Azure subscription ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000).
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewKubeEnvironmentsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*KubeEnvironmentsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &KubeEnvironmentsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Description for Creates or updates a Kubernetes Environment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - resourceGroupName - Name of the resource group to which the resource belongs.
//   - name - Name of the Kubernetes Environment.
//   - kubeEnvironmentEnvelope - Configuration details of the Kubernetes Environment.
//   - options - KubeEnvironmentsClientBeginCreateOrUpdateOptions contains the optional parameters for the KubeEnvironmentsClient.BeginCreateOrUpdate
//     method.
func (client *KubeEnvironmentsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, name string, kubeEnvironmentEnvelope KubeEnvironment, options *KubeEnvironmentsClientBeginCreateOrUpdateOptions) (*runtime.Poller[KubeEnvironmentsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, name, kubeEnvironmentEnvelope, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[KubeEnvironmentsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[KubeEnvironmentsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Description for Creates or updates a Kubernetes Environment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
func (client *KubeEnvironmentsClient) createOrUpdate(ctx context.Context, resourceGroupName string, name string, kubeEnvironmentEnvelope KubeEnvironment, options *KubeEnvironmentsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "KubeEnvironmentsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, name, kubeEnvironmentEnvelope, options)
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
func (client *KubeEnvironmentsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, name string, kubeEnvironmentEnvelope KubeEnvironment, options *KubeEnvironmentsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/kubeEnvironments/{name}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, kubeEnvironmentEnvelope); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Description for Delete a Kubernetes Environment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - resourceGroupName - Name of the resource group to which the resource belongs.
//   - name - Name of the Kubernetes Environment.
//   - options - KubeEnvironmentsClientBeginDeleteOptions contains the optional parameters for the KubeEnvironmentsClient.BeginDelete
//     method.
func (client *KubeEnvironmentsClient) BeginDelete(ctx context.Context, resourceGroupName string, name string, options *KubeEnvironmentsClientBeginDeleteOptions) (*runtime.Poller[KubeEnvironmentsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, name, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[KubeEnvironmentsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[KubeEnvironmentsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Description for Delete a Kubernetes Environment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
func (client *KubeEnvironmentsClient) deleteOperation(ctx context.Context, resourceGroupName string, name string, options *KubeEnvironmentsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "KubeEnvironmentsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, name, options)
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
func (client *KubeEnvironmentsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, name string, options *KubeEnvironmentsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/kubeEnvironments/{name}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Description for Get the properties of a Kubernetes Environment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - resourceGroupName - Name of the resource group to which the resource belongs.
//   - name - Name of the Kubernetes Environment.
//   - options - KubeEnvironmentsClientGetOptions contains the optional parameters for the KubeEnvironmentsClient.Get method.
func (client *KubeEnvironmentsClient) Get(ctx context.Context, resourceGroupName string, name string, options *KubeEnvironmentsClientGetOptions) (KubeEnvironmentsClientGetResponse, error) {
	var err error
	const operationName = "KubeEnvironmentsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, name, options)
	if err != nil {
		return KubeEnvironmentsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return KubeEnvironmentsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return KubeEnvironmentsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *KubeEnvironmentsClient) getCreateRequest(ctx context.Context, resourceGroupName string, name string, options *KubeEnvironmentsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/kubeEnvironments/{name}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
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

// getHandleResponse handles the Get response.
func (client *KubeEnvironmentsClient) getHandleResponse(resp *http.Response) (KubeEnvironmentsClientGetResponse, error) {
	result := KubeEnvironmentsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.KubeEnvironment); err != nil {
		return KubeEnvironmentsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Description for Get all the Kubernetes Environments in a resource group.
//
// Generated from API version 2024-04-01
//   - resourceGroupName - Name of the resource group to which the resource belongs.
//   - options - KubeEnvironmentsClientListByResourceGroupOptions contains the optional parameters for the KubeEnvironmentsClient.NewListByResourceGroupPager
//     method.
func (client *KubeEnvironmentsClient) NewListByResourceGroupPager(resourceGroupName string, options *KubeEnvironmentsClientListByResourceGroupOptions) *runtime.Pager[KubeEnvironmentsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[KubeEnvironmentsClientListByResourceGroupResponse]{
		More: func(page KubeEnvironmentsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *KubeEnvironmentsClientListByResourceGroupResponse) (KubeEnvironmentsClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "KubeEnvironmentsClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return KubeEnvironmentsClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *KubeEnvironmentsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *KubeEnvironmentsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/kubeEnvironments"
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

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *KubeEnvironmentsClient) listByResourceGroupHandleResponse(resp *http.Response) (KubeEnvironmentsClientListByResourceGroupResponse, error) {
	result := KubeEnvironmentsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.KubeEnvironmentCollection); err != nil {
		return KubeEnvironmentsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Description for Get all Kubernetes Environments for a subscription.
//
// Generated from API version 2024-04-01
//   - options - KubeEnvironmentsClientListBySubscriptionOptions contains the optional parameters for the KubeEnvironmentsClient.NewListBySubscriptionPager
//     method.
func (client *KubeEnvironmentsClient) NewListBySubscriptionPager(options *KubeEnvironmentsClientListBySubscriptionOptions) *runtime.Pager[KubeEnvironmentsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[KubeEnvironmentsClientListBySubscriptionResponse]{
		More: func(page KubeEnvironmentsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *KubeEnvironmentsClientListBySubscriptionResponse) (KubeEnvironmentsClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "KubeEnvironmentsClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return KubeEnvironmentsClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *KubeEnvironmentsClient) listBySubscriptionCreateRequest(ctx context.Context, options *KubeEnvironmentsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Web/kubeEnvironments"
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

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *KubeEnvironmentsClient) listBySubscriptionHandleResponse(resp *http.Response) (KubeEnvironmentsClientListBySubscriptionResponse, error) {
	result := KubeEnvironmentsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.KubeEnvironmentCollection); err != nil {
		return KubeEnvironmentsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Description for Creates or updates a Kubernetes Environment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - resourceGroupName - Name of the resource group to which the resource belongs.
//   - name - Name of the Kubernetes Environment.
//   - kubeEnvironmentEnvelope - Configuration details of the Kubernetes Environment.
//   - options - KubeEnvironmentsClientUpdateOptions contains the optional parameters for the KubeEnvironmentsClient.Update method.
func (client *KubeEnvironmentsClient) Update(ctx context.Context, resourceGroupName string, name string, kubeEnvironmentEnvelope KubeEnvironmentPatchResource, options *KubeEnvironmentsClientUpdateOptions) (KubeEnvironmentsClientUpdateResponse, error) {
	var err error
	const operationName = "KubeEnvironmentsClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, name, kubeEnvironmentEnvelope, options)
	if err != nil {
		return KubeEnvironmentsClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return KubeEnvironmentsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return KubeEnvironmentsClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *KubeEnvironmentsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, name string, kubeEnvironmentEnvelope KubeEnvironmentPatchResource, options *KubeEnvironmentsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/kubeEnvironments/{name}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, kubeEnvironmentEnvelope); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *KubeEnvironmentsClient) updateHandleResponse(resp *http.Response) (KubeEnvironmentsClientUpdateResponse, error) {
	result := KubeEnvironmentsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.KubeEnvironment); err != nil {
		return KubeEnvironmentsClientUpdateResponse{}, err
	}
	return result, nil
}
