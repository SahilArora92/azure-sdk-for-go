//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdevcenter

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

// CheckScopedNameAvailabilityClient contains the methods for the CheckScopedNameAvailability group.
// Don't use this type directly, use NewCheckScopedNameAvailabilityClient() instead.
type CheckScopedNameAvailabilityClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewCheckScopedNameAvailabilityClient creates a new instance of CheckScopedNameAvailabilityClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewCheckScopedNameAvailabilityClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CheckScopedNameAvailabilityClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &CheckScopedNameAvailabilityClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Execute - Check the availability of name for resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
//   - nameAvailabilityRequest - The required parameters for checking if resource name is available.
//   - options - CheckScopedNameAvailabilityClientExecuteOptions contains the optional parameters for the CheckScopedNameAvailabilityClient.Execute
//     method.
func (client *CheckScopedNameAvailabilityClient) Execute(ctx context.Context, nameAvailabilityRequest CheckScopedNameAvailabilityRequest, options *CheckScopedNameAvailabilityClientExecuteOptions) (CheckScopedNameAvailabilityClientExecuteResponse, error) {
	var err error
	const operationName = "CheckScopedNameAvailabilityClient.Execute"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.executeCreateRequest(ctx, nameAvailabilityRequest, options)
	if err != nil {
		return CheckScopedNameAvailabilityClientExecuteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CheckScopedNameAvailabilityClientExecuteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return CheckScopedNameAvailabilityClientExecuteResponse{}, err
	}
	resp, err := client.executeHandleResponse(httpResp)
	return resp, err
}

// executeCreateRequest creates the Execute request.
func (client *CheckScopedNameAvailabilityClient) executeCreateRequest(ctx context.Context, nameAvailabilityRequest CheckScopedNameAvailabilityRequest, options *CheckScopedNameAvailabilityClientExecuteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DevCenter/checkScopedNameAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, nameAvailabilityRequest); err != nil {
		return nil, err
	}
	return req, nil
}

// executeHandleResponse handles the Execute response.
func (client *CheckScopedNameAvailabilityClient) executeHandleResponse(resp *http.Response) (CheckScopedNameAvailabilityClientExecuteResponse, error) {
	result := CheckScopedNameAvailabilityClientExecuteResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckNameAvailabilityResponse); err != nil {
		return CheckScopedNameAvailabilityClientExecuteResponse{}, err
	}
	return result, nil
}
