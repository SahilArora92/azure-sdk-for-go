//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapplicationinsights

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

// APIKeysClient contains the methods for the APIKeys group.
// Don't use this type directly, use NewAPIKeysClient() instead.
type APIKeysClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAPIKeysClient creates a new instance of APIKeysClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAPIKeysClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*APIKeysClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &APIKeysClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Create - Create an API Key of an Application Insights component.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the Application Insights component resource.
// apiKeyProperties - Properties that need to be specified to create an API key of a Application Insights component.
// options - APIKeysClientCreateOptions contains the optional parameters for the APIKeysClient.Create method.
func (client *APIKeysClient) Create(ctx context.Context, resourceGroupName string, resourceName string, apiKeyProperties APIKeyRequest, options *APIKeysClientCreateOptions) (APIKeysClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, resourceName, apiKeyProperties, options)
	if err != nil {
		return APIKeysClientCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return APIKeysClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return APIKeysClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *APIKeysClient) createCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, apiKeyProperties APIKeyRequest, options *APIKeysClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/ApiKeys"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, apiKeyProperties)
}

// createHandleResponse handles the Create response.
func (client *APIKeysClient) createHandleResponse(resp *http.Response) (APIKeysClientCreateResponse, error) {
	result := APIKeysClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ComponentAPIKey); err != nil {
		return APIKeysClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Delete an API Key of an Application Insights component.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the Application Insights component resource.
// keyID - The API Key ID. This is unique within a Application Insights component.
// options - APIKeysClientDeleteOptions contains the optional parameters for the APIKeysClient.Delete method.
func (client *APIKeysClient) Delete(ctx context.Context, resourceGroupName string, resourceName string, keyID string, options *APIKeysClientDeleteOptions) (APIKeysClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, resourceName, keyID, options)
	if err != nil {
		return APIKeysClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return APIKeysClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return APIKeysClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return client.deleteHandleResponse(resp)
}

// deleteCreateRequest creates the Delete request.
func (client *APIKeysClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, keyID string, options *APIKeysClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/APIKeys/{keyId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if keyID == "" {
		return nil, errors.New("parameter keyID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{keyId}", url.PathEscape(keyID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *APIKeysClient) deleteHandleResponse(resp *http.Response) (APIKeysClientDeleteResponse, error) {
	result := APIKeysClientDeleteResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ComponentAPIKey); err != nil {
		return APIKeysClientDeleteResponse{}, err
	}
	return result, nil
}

// Get - Get the API Key for this key id.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the Application Insights component resource.
// keyID - The API Key ID. This is unique within a Application Insights component.
// options - APIKeysClientGetOptions contains the optional parameters for the APIKeysClient.Get method.
func (client *APIKeysClient) Get(ctx context.Context, resourceGroupName string, resourceName string, keyID string, options *APIKeysClientGetOptions) (APIKeysClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, keyID, options)
	if err != nil {
		return APIKeysClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return APIKeysClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return APIKeysClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *APIKeysClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, keyID string, options *APIKeysClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/APIKeys/{keyId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if keyID == "" {
		return nil, errors.New("parameter keyID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{keyId}", url.PathEscape(keyID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *APIKeysClient) getHandleResponse(resp *http.Response) (APIKeysClientGetResponse, error) {
	result := APIKeysClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ComponentAPIKey); err != nil {
		return APIKeysClientGetResponse{}, err
	}
	return result, nil
}

// List - Gets a list of API keys of an Application Insights component.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the Application Insights component resource.
// options - APIKeysClientListOptions contains the optional parameters for the APIKeysClient.List method.
func (client *APIKeysClient) List(resourceGroupName string, resourceName string, options *APIKeysClientListOptions) *runtime.Pager[APIKeysClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[APIKeysClientListResponse]{
		More: func(page APIKeysClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *APIKeysClientListResponse) (APIKeysClientListResponse, error) {
			req, err := client.listCreateRequest(ctx, resourceGroupName, resourceName, options)
			if err != nil {
				return APIKeysClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return APIKeysClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return APIKeysClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *APIKeysClient) listCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *APIKeysClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/ApiKeys"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *APIKeysClient) listHandleResponse(resp *http.Response) (APIKeysClientListResponse, error) {
	result := APIKeysClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ComponentAPIKeyListResult); err != nil {
		return APIKeysClientListResponse{}, err
	}
	return result, nil
}
