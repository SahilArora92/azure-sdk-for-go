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
	"strings"
)

// ServiceSKUsClient contains the methods for the APIManagementServiceSKUs group.
// Don't use this type directly, use NewServiceSKUsClient() instead.
type ServiceSKUsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewServiceSKUsClient creates a new instance of ServiceSKUsClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewServiceSKUsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ServiceSKUsClient, error) {
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
	client := &ServiceSKUsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// ListAvailableServiceSKUs - Gets all available SKU for a given API Management service
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// options - ServiceSKUsClientListAvailableServiceSKUsOptions contains the optional parameters for the ServiceSKUsClient.ListAvailableServiceSKUs
// method.
func (client *ServiceSKUsClient) ListAvailableServiceSKUs(resourceGroupName string, serviceName string, options *ServiceSKUsClientListAvailableServiceSKUsOptions) *runtime.Pager[ServiceSKUsClientListAvailableServiceSKUsResponse] {
	return runtime.NewPager(runtime.PageProcessor[ServiceSKUsClientListAvailableServiceSKUsResponse]{
		More: func(page ServiceSKUsClientListAvailableServiceSKUsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ServiceSKUsClientListAvailableServiceSKUsResponse) (ServiceSKUsClientListAvailableServiceSKUsResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listAvailableServiceSKUsCreateRequest(ctx, resourceGroupName, serviceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ServiceSKUsClientListAvailableServiceSKUsResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ServiceSKUsClientListAvailableServiceSKUsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ServiceSKUsClientListAvailableServiceSKUsResponse{}, runtime.NewResponseError(resp)
			}
			return client.listAvailableServiceSKUsHandleResponse(resp)
		},
	})
}

// listAvailableServiceSKUsCreateRequest creates the ListAvailableServiceSKUs request.
func (client *ServiceSKUsClient) listAvailableServiceSKUsCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *ServiceSKUsClientListAvailableServiceSKUsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/skus"
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
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listAvailableServiceSKUsHandleResponse handles the ListAvailableServiceSKUs response.
func (client *ServiceSKUsClient) listAvailableServiceSKUsHandleResponse(resp *http.Response) (ServiceSKUsClientListAvailableServiceSKUsResponse, error) {
	result := ServiceSKUsClientListAvailableServiceSKUsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceSKUResults); err != nil {
		return ServiceSKUsClientListAvailableServiceSKUsResponse{}, err
	}
	return result, nil
}
