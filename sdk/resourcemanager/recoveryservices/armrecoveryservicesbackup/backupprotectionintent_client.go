//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicesbackup

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

// BackupProtectionIntentClient contains the methods for the BackupProtectionIntent group.
// Don't use this type directly, use NewBackupProtectionIntentClient() instead.
type BackupProtectionIntentClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewBackupProtectionIntentClient creates a new instance of BackupProtectionIntentClient with the specified values.
//   - subscriptionID - The subscription Id.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewBackupProtectionIntentClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*BackupProtectionIntentClient, error) {
	cl, err := arm.NewClient(moduleName+".BackupProtectionIntentClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &BackupProtectionIntentClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - Provides a pageable list of all intents that are present within a vault.
//
// Generated from API version 2023-06-01
//   - vaultName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - options - BackupProtectionIntentClientListOptions contains the optional parameters for the BackupProtectionIntentClient.NewListPager
//     method.
func (client *BackupProtectionIntentClient) NewListPager(vaultName string, resourceGroupName string, options *BackupProtectionIntentClientListOptions) *runtime.Pager[BackupProtectionIntentClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[BackupProtectionIntentClientListResponse]{
		More: func(page BackupProtectionIntentClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *BackupProtectionIntentClientListResponse) (BackupProtectionIntentClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, vaultName, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return BackupProtectionIntentClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return BackupProtectionIntentClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return BackupProtectionIntentClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *BackupProtectionIntentClient) listCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, options *BackupProtectionIntentClientListOptions) (*policy.Request, error) {
	urlPath := "/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupProtectionIntents"
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
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
	reqQP.Set("api-version", "2023-06-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *BackupProtectionIntentClient) listHandleResponse(resp *http.Response) (BackupProtectionIntentClientListResponse, error) {
	result := BackupProtectionIntentClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProtectionIntentResourceList); err != nil {
		return BackupProtectionIntentClientListResponse{}, err
	}
	return result, nil
}
