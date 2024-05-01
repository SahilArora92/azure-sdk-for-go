//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstoragemover

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

// ProjectsClient contains the methods for the Projects group.
// Don't use this type directly, use NewProjectsClient() instead.
type ProjectsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewProjectsClient creates a new instance of ProjectsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewProjectsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ProjectsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ProjectsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a Project resource, which is a logical grouping of related jobs.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - storageMoverName - The name of the Storage Mover resource.
//   - projectName - The name of the Project resource.
//   - options - ProjectsClientCreateOrUpdateOptions contains the optional parameters for the ProjectsClient.CreateOrUpdate method.
func (client *ProjectsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, storageMoverName string, projectName string, project Project, options *ProjectsClientCreateOrUpdateOptions) (ProjectsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "ProjectsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, storageMoverName, projectName, project, options)
	if err != nil {
		return ProjectsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProjectsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ProjectsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ProjectsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, storageMoverName string, projectName string, project Project, options *ProjectsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageMover/storageMovers/{storageMoverName}/projects/{projectName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageMoverName == "" {
		return nil, errors.New("parameter storageMoverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageMoverName}", url.PathEscape(storageMoverName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, project); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ProjectsClient) createOrUpdateHandleResponse(resp *http.Response) (ProjectsClientCreateOrUpdateResponse, error) {
	result := ProjectsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Project); err != nil {
		return ProjectsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// BeginDelete - Deletes a Project resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - storageMoverName - The name of the Storage Mover resource.
//   - projectName - The name of the Project resource.
//   - options - ProjectsClientBeginDeleteOptions contains the optional parameters for the ProjectsClient.BeginDelete method.
func (client *ProjectsClient) BeginDelete(ctx context.Context, resourceGroupName string, storageMoverName string, projectName string, options *ProjectsClientBeginDeleteOptions) (*runtime.Poller[ProjectsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, storageMoverName, projectName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ProjectsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ProjectsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes a Project resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01-preview
func (client *ProjectsClient) deleteOperation(ctx context.Context, resourceGroupName string, storageMoverName string, projectName string, options *ProjectsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "ProjectsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, storageMoverName, projectName, options)
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
func (client *ProjectsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, storageMoverName string, projectName string, options *ProjectsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageMover/storageMovers/{storageMoverName}/projects/{projectName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageMoverName == "" {
		return nil, errors.New("parameter storageMoverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageMoverName}", url.PathEscape(storageMoverName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a Project resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - storageMoverName - The name of the Storage Mover resource.
//   - projectName - The name of the Project resource.
//   - options - ProjectsClientGetOptions contains the optional parameters for the ProjectsClient.Get method.
func (client *ProjectsClient) Get(ctx context.Context, resourceGroupName string, storageMoverName string, projectName string, options *ProjectsClientGetOptions) (ProjectsClientGetResponse, error) {
	var err error
	const operationName = "ProjectsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, storageMoverName, projectName, options)
	if err != nil {
		return ProjectsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProjectsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ProjectsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ProjectsClient) getCreateRequest(ctx context.Context, resourceGroupName string, storageMoverName string, projectName string, options *ProjectsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageMover/storageMovers/{storageMoverName}/projects/{projectName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageMoverName == "" {
		return nil, errors.New("parameter storageMoverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageMoverName}", url.PathEscape(storageMoverName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ProjectsClient) getHandleResponse(resp *http.Response) (ProjectsClientGetResponse, error) {
	result := ProjectsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Project); err != nil {
		return ProjectsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all Projects in a Storage Mover.
//
// Generated from API version 2024-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - storageMoverName - The name of the Storage Mover resource.
//   - options - ProjectsClientListOptions contains the optional parameters for the ProjectsClient.NewListPager method.
func (client *ProjectsClient) NewListPager(resourceGroupName string, storageMoverName string, options *ProjectsClientListOptions) *runtime.Pager[ProjectsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ProjectsClientListResponse]{
		More: func(page ProjectsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ProjectsClientListResponse) (ProjectsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ProjectsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, storageMoverName, options)
			}, nil)
			if err != nil {
				return ProjectsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ProjectsClient) listCreateRequest(ctx context.Context, resourceGroupName string, storageMoverName string, options *ProjectsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageMover/storageMovers/{storageMoverName}/projects"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageMoverName == "" {
		return nil, errors.New("parameter storageMoverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageMoverName}", url.PathEscape(storageMoverName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ProjectsClient) listHandleResponse(resp *http.Response) (ProjectsClientListResponse, error) {
	result := ProjectsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProjectList); err != nil {
		return ProjectsClientListResponse{}, err
	}
	return result, nil
}

// Update - Updates properties for a Project resource. Properties not specified in the request body will be unchanged.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - storageMoverName - The name of the Storage Mover resource.
//   - projectName - The name of the Project resource.
//   - options - ProjectsClientUpdateOptions contains the optional parameters for the ProjectsClient.Update method.
func (client *ProjectsClient) Update(ctx context.Context, resourceGroupName string, storageMoverName string, projectName string, project ProjectUpdateParameters, options *ProjectsClientUpdateOptions) (ProjectsClientUpdateResponse, error) {
	var err error
	const operationName = "ProjectsClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, storageMoverName, projectName, project, options)
	if err != nil {
		return ProjectsClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProjectsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ProjectsClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *ProjectsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, storageMoverName string, projectName string, project ProjectUpdateParameters, options *ProjectsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageMover/storageMovers/{storageMoverName}/projects/{projectName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageMoverName == "" {
		return nil, errors.New("parameter storageMoverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageMoverName}", url.PathEscape(storageMoverName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, project); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *ProjectsClient) updateHandleResponse(resp *http.Response) (ProjectsClientUpdateResponse, error) {
	result := ProjectsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Project); err != nil {
		return ProjectsClientUpdateResponse{}, err
	}
	return result, nil
}
