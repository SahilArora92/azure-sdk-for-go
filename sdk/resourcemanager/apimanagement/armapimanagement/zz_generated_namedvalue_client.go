//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement

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
	"strconv"
	"strings"
)

// NamedValueClient contains the methods for the NamedValue group.
// Don't use this type directly, use NewNamedValueClient() instead.
type NamedValueClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewNamedValueClient creates a new instance of NamedValueClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewNamedValueClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*NamedValueClient, error) {
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
	client := &NamedValueClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates named value.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// namedValueID - Identifier of the NamedValue.
// parameters - Create parameters.
// options - NamedValueClientBeginCreateOrUpdateOptions contains the optional parameters for the NamedValueClient.BeginCreateOrUpdate
// method.
func (client *NamedValueClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, namedValueID string, parameters NamedValueCreateContract, options *NamedValueClientBeginCreateOrUpdateOptions) (*armruntime.Poller[NamedValueClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, serviceName, namedValueID, parameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[NamedValueClientCreateOrUpdateResponse]{
			FinalStateVia: armruntime.FinalStateViaLocation,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[NamedValueClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates named value.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *NamedValueClient) createOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, namedValueID string, parameters NamedValueCreateContract, options *NamedValueClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, namedValueID, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *NamedValueClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, namedValueID string, parameters NamedValueCreateContract, options *NamedValueClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/namedValues/{namedValueId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if namedValueID == "" {
		return nil, errors.New("parameter namedValueID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namedValueId}", url.PathEscape(namedValueID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header.Set("If-Match", *options.IfMatch)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// Delete - Deletes specific named value from the API Management service instance.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// namedValueID - Identifier of the NamedValue.
// ifMatch - ETag of the Entity. ETag should match the current entity state from the header response of the GET request or
// it should be * for unconditional update.
// options - NamedValueClientDeleteOptions contains the optional parameters for the NamedValueClient.Delete method.
func (client *NamedValueClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, namedValueID string, ifMatch string, options *NamedValueClientDeleteOptions) (NamedValueClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, namedValueID, ifMatch, options)
	if err != nil {
		return NamedValueClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return NamedValueClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return NamedValueClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return NamedValueClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *NamedValueClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, namedValueID string, ifMatch string, options *NamedValueClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/namedValues/{namedValueId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if namedValueID == "" {
		return nil, errors.New("parameter namedValueID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namedValueId}", url.PathEscape(namedValueID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("If-Match", ifMatch)
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets the details of the named value specified by its identifier.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// namedValueID - Identifier of the NamedValue.
// options - NamedValueClientGetOptions contains the optional parameters for the NamedValueClient.Get method.
func (client *NamedValueClient) Get(ctx context.Context, resourceGroupName string, serviceName string, namedValueID string, options *NamedValueClientGetOptions) (NamedValueClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, namedValueID, options)
	if err != nil {
		return NamedValueClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return NamedValueClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return NamedValueClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *NamedValueClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, namedValueID string, options *NamedValueClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/namedValues/{namedValueId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if namedValueID == "" {
		return nil, errors.New("parameter namedValueID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namedValueId}", url.PathEscape(namedValueID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *NamedValueClient) getHandleResponse(resp *http.Response) (NamedValueClientGetResponse, error) {
	result := NamedValueClientGetResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.NamedValueContract); err != nil {
		return NamedValueClientGetResponse{}, err
	}
	return result, nil
}

// GetEntityTag - Gets the entity state (Etag) version of the named value specified by its identifier.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// namedValueID - Identifier of the NamedValue.
// options - NamedValueClientGetEntityTagOptions contains the optional parameters for the NamedValueClient.GetEntityTag method.
func (client *NamedValueClient) GetEntityTag(ctx context.Context, resourceGroupName string, serviceName string, namedValueID string, options *NamedValueClientGetEntityTagOptions) (NamedValueClientGetEntityTagResponse, error) {
	req, err := client.getEntityTagCreateRequest(ctx, resourceGroupName, serviceName, namedValueID, options)
	if err != nil {
		return NamedValueClientGetEntityTagResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return NamedValueClientGetEntityTagResponse{}, err
	}
	return client.getEntityTagHandleResponse(resp)
}

// getEntityTagCreateRequest creates the GetEntityTag request.
func (client *NamedValueClient) getEntityTagCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, namedValueID string, options *NamedValueClientGetEntityTagOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/namedValues/{namedValueId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if namedValueID == "" {
		return nil, errors.New("parameter namedValueID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namedValueId}", url.PathEscape(namedValueID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getEntityTagHandleResponse handles the GetEntityTag response.
func (client *NamedValueClient) getEntityTagHandleResponse(resp *http.Response) (NamedValueClientGetEntityTagResponse, error) {
	result := NamedValueClientGetEntityTagResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		result.Success = true
	}
	return result, nil
}

// ListByService - Lists a collection of named values defined within a service instance.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// options - NamedValueClientListByServiceOptions contains the optional parameters for the NamedValueClient.ListByService
// method.
func (client *NamedValueClient) ListByService(resourceGroupName string, serviceName string, options *NamedValueClientListByServiceOptions) *runtime.Pager[NamedValueClientListByServiceResponse] {
	return runtime.NewPager(runtime.PageProcessor[NamedValueClientListByServiceResponse]{
		More: func(page NamedValueClientListByServiceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *NamedValueClientListByServiceResponse) (NamedValueClientListByServiceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByServiceCreateRequest(ctx, resourceGroupName, serviceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return NamedValueClientListByServiceResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return NamedValueClientListByServiceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return NamedValueClientListByServiceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByServiceHandleResponse(resp)
		},
	})
}

// listByServiceCreateRequest creates the ListByService request.
func (client *NamedValueClient) listByServiceCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *NamedValueClientListByServiceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/namedValues"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	if options != nil && options.IsKeyVaultRefreshFailed != nil {
		reqQP.Set("isKeyVaultRefreshFailed", strconv.FormatBool(*options.IsKeyVaultRefreshFailed))
	}
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByServiceHandleResponse handles the ListByService response.
func (client *NamedValueClient) listByServiceHandleResponse(resp *http.Response) (NamedValueClientListByServiceResponse, error) {
	result := NamedValueClientListByServiceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NamedValueCollection); err != nil {
		return NamedValueClientListByServiceResponse{}, err
	}
	return result, nil
}

// ListValue - Gets the secret of the named value specified by its identifier.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// namedValueID - Identifier of the NamedValue.
// options - NamedValueClientListValueOptions contains the optional parameters for the NamedValueClient.ListValue method.
func (client *NamedValueClient) ListValue(ctx context.Context, resourceGroupName string, serviceName string, namedValueID string, options *NamedValueClientListValueOptions) (NamedValueClientListValueResponse, error) {
	req, err := client.listValueCreateRequest(ctx, resourceGroupName, serviceName, namedValueID, options)
	if err != nil {
		return NamedValueClientListValueResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return NamedValueClientListValueResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return NamedValueClientListValueResponse{}, runtime.NewResponseError(resp)
	}
	return client.listValueHandleResponse(resp)
}

// listValueCreateRequest creates the ListValue request.
func (client *NamedValueClient) listValueCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, namedValueID string, options *NamedValueClientListValueOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/namedValues/{namedValueId}/listValue"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if namedValueID == "" {
		return nil, errors.New("parameter namedValueID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namedValueId}", url.PathEscape(namedValueID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listValueHandleResponse handles the ListValue response.
func (client *NamedValueClient) listValueHandleResponse(resp *http.Response) (NamedValueClientListValueResponse, error) {
	result := NamedValueClientListValueResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.NamedValueSecretContract); err != nil {
		return NamedValueClientListValueResponse{}, err
	}
	return result, nil
}

// BeginRefreshSecret - Refresh the secret of the named value specified by its identifier.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// namedValueID - Identifier of the NamedValue.
// options - NamedValueClientBeginRefreshSecretOptions contains the optional parameters for the NamedValueClient.BeginRefreshSecret
// method.
func (client *NamedValueClient) BeginRefreshSecret(ctx context.Context, resourceGroupName string, serviceName string, namedValueID string, options *NamedValueClientBeginRefreshSecretOptions) (*armruntime.Poller[NamedValueClientRefreshSecretResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.refreshSecret(ctx, resourceGroupName, serviceName, namedValueID, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[NamedValueClientRefreshSecretResponse]{
			FinalStateVia: armruntime.FinalStateViaLocation,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[NamedValueClientRefreshSecretResponse](options.ResumeToken, client.pl, nil)
	}
}

// RefreshSecret - Refresh the secret of the named value specified by its identifier.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *NamedValueClient) refreshSecret(ctx context.Context, resourceGroupName string, serviceName string, namedValueID string, options *NamedValueClientBeginRefreshSecretOptions) (*http.Response, error) {
	req, err := client.refreshSecretCreateRequest(ctx, resourceGroupName, serviceName, namedValueID, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// refreshSecretCreateRequest creates the RefreshSecret request.
func (client *NamedValueClient) refreshSecretCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, namedValueID string, options *NamedValueClientBeginRefreshSecretOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/namedValues/{namedValueId}/refreshSecret"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if namedValueID == "" {
		return nil, errors.New("parameter namedValueID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namedValueId}", url.PathEscape(namedValueID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// BeginUpdate - Updates the specific named value.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// namedValueID - Identifier of the NamedValue.
// ifMatch - ETag of the Entity. ETag should match the current entity state from the header response of the GET request or
// it should be * for unconditional update.
// parameters - Update parameters.
// options - NamedValueClientBeginUpdateOptions contains the optional parameters for the NamedValueClient.BeginUpdate method.
func (client *NamedValueClient) BeginUpdate(ctx context.Context, resourceGroupName string, serviceName string, namedValueID string, ifMatch string, parameters NamedValueUpdateParameters, options *NamedValueClientBeginUpdateOptions) (*armruntime.Poller[NamedValueClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, serviceName, namedValueID, ifMatch, parameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[NamedValueClientUpdateResponse]{
			FinalStateVia: armruntime.FinalStateViaLocation,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[NamedValueClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Updates the specific named value.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *NamedValueClient) update(ctx context.Context, resourceGroupName string, serviceName string, namedValueID string, ifMatch string, parameters NamedValueUpdateParameters, options *NamedValueClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, serviceName, namedValueID, ifMatch, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *NamedValueClient) updateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, namedValueID string, ifMatch string, parameters NamedValueUpdateParameters, options *NamedValueClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/namedValues/{namedValueId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if namedValueID == "" {
		return nil, errors.New("parameter namedValueID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namedValueId}", url.PathEscape(namedValueID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("If-Match", ifMatch)
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}
