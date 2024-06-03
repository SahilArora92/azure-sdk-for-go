//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpostgresqlflexibleservers

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

// QuotaUsagesClient contains the methods for the QuotaUsages group.
// Don't use this type directly, use NewQuotaUsagesClient() instead.
type QuotaUsagesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewQuotaUsagesClient creates a new instance of QuotaUsagesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewQuotaUsagesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*QuotaUsagesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &QuotaUsagesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - Get quota usages at specified location in a given subscription.
//
// Generated from API version 2024-03-01-privatepreview
//   - locationName - The name of the location.
//   - options - QuotaUsagesClientListOptions contains the optional parameters for the QuotaUsagesClient.NewListPager method.
func (client *QuotaUsagesClient) NewListPager(locationName string, options *QuotaUsagesClientListOptions) *runtime.Pager[QuotaUsagesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[QuotaUsagesClientListResponse]{
		More: func(page QuotaUsagesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *QuotaUsagesClientListResponse) (QuotaUsagesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "QuotaUsagesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, locationName, options)
			}, nil)
			if err != nil {
				return QuotaUsagesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *QuotaUsagesClient) listCreateRequest(ctx context.Context, locationName string, options *QuotaUsagesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DBforPostgreSQL/locations/{locationName}/resourceType/flexibleServers/usages"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *QuotaUsagesClient) listHandleResponse(resp *http.Response) (QuotaUsagesClientListResponse, error) {
	result := QuotaUsagesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.QuotaUsagesListResult); err != nil {
		return QuotaUsagesClientListResponse{}, err
	}
	return result, nil
}
