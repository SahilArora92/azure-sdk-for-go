//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrelay

// HybridConnectionsClientCreateOrUpdateAuthorizationRuleResponse contains the response from method HybridConnectionsClient.CreateOrUpdateAuthorizationRule.
type HybridConnectionsClientCreateOrUpdateAuthorizationRuleResponse struct {
	AuthorizationRule
}

// HybridConnectionsClientCreateOrUpdateResponse contains the response from method HybridConnectionsClient.CreateOrUpdate.
type HybridConnectionsClientCreateOrUpdateResponse struct {
	HybridConnection
}

// HybridConnectionsClientDeleteAuthorizationRuleResponse contains the response from method HybridConnectionsClient.DeleteAuthorizationRule.
type HybridConnectionsClientDeleteAuthorizationRuleResponse struct {
	// placeholder for future response values
}

// HybridConnectionsClientDeleteResponse contains the response from method HybridConnectionsClient.Delete.
type HybridConnectionsClientDeleteResponse struct {
	// placeholder for future response values
}

// HybridConnectionsClientGetAuthorizationRuleResponse contains the response from method HybridConnectionsClient.GetAuthorizationRule.
type HybridConnectionsClientGetAuthorizationRuleResponse struct {
	AuthorizationRule
}

// HybridConnectionsClientGetResponse contains the response from method HybridConnectionsClient.Get.
type HybridConnectionsClientGetResponse struct {
	HybridConnection
}

// HybridConnectionsClientListAuthorizationRulesResponse contains the response from method HybridConnectionsClient.ListAuthorizationRules.
type HybridConnectionsClientListAuthorizationRulesResponse struct {
	AuthorizationRuleListResult
}

// HybridConnectionsClientListByNamespaceResponse contains the response from method HybridConnectionsClient.ListByNamespace.
type HybridConnectionsClientListByNamespaceResponse struct {
	HybridConnectionListResult
}

// HybridConnectionsClientListKeysResponse contains the response from method HybridConnectionsClient.ListKeys.
type HybridConnectionsClientListKeysResponse struct {
	AccessKeys
}

// HybridConnectionsClientRegenerateKeysResponse contains the response from method HybridConnectionsClient.RegenerateKeys.
type HybridConnectionsClientRegenerateKeysResponse struct {
	AccessKeys
}

// NamespacesClientCheckNameAvailabilityResponse contains the response from method NamespacesClient.CheckNameAvailability.
type NamespacesClientCheckNameAvailabilityResponse struct {
	CheckNameAvailabilityResult
}

// NamespacesClientCreateOrUpdateAuthorizationRuleResponse contains the response from method NamespacesClient.CreateOrUpdateAuthorizationRule.
type NamespacesClientCreateOrUpdateAuthorizationRuleResponse struct {
	AuthorizationRule
}

// NamespacesClientCreateOrUpdateNetworkRuleSetResponse contains the response from method NamespacesClient.CreateOrUpdateNetworkRuleSet.
type NamespacesClientCreateOrUpdateNetworkRuleSetResponse struct {
	NetworkRuleSet
}

// NamespacesClientCreateOrUpdateResponse contains the response from method NamespacesClient.CreateOrUpdate.
type NamespacesClientCreateOrUpdateResponse struct {
	Namespace
}

// NamespacesClientDeleteAuthorizationRuleResponse contains the response from method NamespacesClient.DeleteAuthorizationRule.
type NamespacesClientDeleteAuthorizationRuleResponse struct {
	// placeholder for future response values
}

// NamespacesClientDeleteResponse contains the response from method NamespacesClient.Delete.
type NamespacesClientDeleteResponse struct {
	// placeholder for future response values
}

// NamespacesClientGetAuthorizationRuleResponse contains the response from method NamespacesClient.GetAuthorizationRule.
type NamespacesClientGetAuthorizationRuleResponse struct {
	AuthorizationRule
}

// NamespacesClientGetNetworkRuleSetResponse contains the response from method NamespacesClient.GetNetworkRuleSet.
type NamespacesClientGetNetworkRuleSetResponse struct {
	NetworkRuleSet
}

// NamespacesClientGetResponse contains the response from method NamespacesClient.Get.
type NamespacesClientGetResponse struct {
	Namespace
}

// NamespacesClientListAuthorizationRulesResponse contains the response from method NamespacesClient.ListAuthorizationRules.
type NamespacesClientListAuthorizationRulesResponse struct {
	AuthorizationRuleListResult
}

// NamespacesClientListByResourceGroupResponse contains the response from method NamespacesClient.ListByResourceGroup.
type NamespacesClientListByResourceGroupResponse struct {
	NamespaceListResult
}

// NamespacesClientListKeysResponse contains the response from method NamespacesClient.ListKeys.
type NamespacesClientListKeysResponse struct {
	AccessKeys
}

// NamespacesClientListResponse contains the response from method NamespacesClient.List.
type NamespacesClientListResponse struct {
	NamespaceListResult
}

// NamespacesClientRegenerateKeysResponse contains the response from method NamespacesClient.RegenerateKeys.
type NamespacesClientRegenerateKeysResponse struct {
	AccessKeys
}

// NamespacesClientUpdateResponse contains the response from method NamespacesClient.Update.
type NamespacesClientUpdateResponse struct {
	Namespace
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationListResult
}

// PrivateEndpointConnectionsClientCreateOrUpdateResponse contains the response from method PrivateEndpointConnectionsClient.CreateOrUpdate.
type PrivateEndpointConnectionsClientCreateOrUpdateResponse struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientDeleteResponse contains the response from method PrivateEndpointConnectionsClient.Delete.
type PrivateEndpointConnectionsClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateEndpointConnectionsClientGetResponse contains the response from method PrivateEndpointConnectionsClient.Get.
type PrivateEndpointConnectionsClientGetResponse struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientListResponse contains the response from method PrivateEndpointConnectionsClient.List.
type PrivateEndpointConnectionsClientListResponse struct {
	PrivateEndpointConnectionListResult
}

// PrivateLinkResourcesClientGetResponse contains the response from method PrivateLinkResourcesClient.Get.
type PrivateLinkResourcesClientGetResponse struct {
	PrivateLinkResource
}

// PrivateLinkResourcesClientListResponse contains the response from method PrivateLinkResourcesClient.List.
type PrivateLinkResourcesClientListResponse struct {
	PrivateLinkResourcesListResult
}

// WCFRelaysClientCreateOrUpdateAuthorizationRuleResponse contains the response from method WCFRelaysClient.CreateOrUpdateAuthorizationRule.
type WCFRelaysClientCreateOrUpdateAuthorizationRuleResponse struct {
	AuthorizationRule
}

// WCFRelaysClientCreateOrUpdateResponse contains the response from method WCFRelaysClient.CreateOrUpdate.
type WCFRelaysClientCreateOrUpdateResponse struct {
	WcfRelay
}

// WCFRelaysClientDeleteAuthorizationRuleResponse contains the response from method WCFRelaysClient.DeleteAuthorizationRule.
type WCFRelaysClientDeleteAuthorizationRuleResponse struct {
	// placeholder for future response values
}

// WCFRelaysClientDeleteResponse contains the response from method WCFRelaysClient.Delete.
type WCFRelaysClientDeleteResponse struct {
	// placeholder for future response values
}

// WCFRelaysClientGetAuthorizationRuleResponse contains the response from method WCFRelaysClient.GetAuthorizationRule.
type WCFRelaysClientGetAuthorizationRuleResponse struct {
	AuthorizationRule
}

// WCFRelaysClientGetResponse contains the response from method WCFRelaysClient.Get.
type WCFRelaysClientGetResponse struct {
	WcfRelay
}

// WCFRelaysClientListAuthorizationRulesResponse contains the response from method WCFRelaysClient.ListAuthorizationRules.
type WCFRelaysClientListAuthorizationRulesResponse struct {
	AuthorizationRuleListResult
}

// WCFRelaysClientListByNamespaceResponse contains the response from method WCFRelaysClient.ListByNamespace.
type WCFRelaysClientListByNamespaceResponse struct {
	WcfRelaysListResult
}

// WCFRelaysClientListKeysResponse contains the response from method WCFRelaysClient.ListKeys.
type WCFRelaysClientListKeysResponse struct {
	AccessKeys
}

// WCFRelaysClientRegenerateKeysResponse contains the response from method WCFRelaysClient.RegenerateKeys.
type WCFRelaysClientRegenerateKeysResponse struct {
	AccessKeys
}
