//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorage

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
)

// ClientFactory is a client factory used to create any client in this module.
// Don't use this type directly, use NewClientFactory instead.
type ClientFactory struct {
	subscriptionID string
	internal       *arm.Client
}

// NewClientFactory creates a new instance of ClientFactory with the specified values.
// The parameter values will be propagated to any client created from this factory.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	internal, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		subscriptionID: subscriptionID,
		internal:       internal,
	}, nil
}

// NewAccountsClient creates a new instance of AccountsClient.
func (c *ClientFactory) NewAccountsClient() *AccountsClient {
	return &AccountsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewBlobContainersClient creates a new instance of BlobContainersClient.
func (c *ClientFactory) NewBlobContainersClient() *BlobContainersClient {
	return &BlobContainersClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewBlobInventoryPoliciesClient creates a new instance of BlobInventoryPoliciesClient.
func (c *ClientFactory) NewBlobInventoryPoliciesClient() *BlobInventoryPoliciesClient {
	return &BlobInventoryPoliciesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewBlobServicesClient creates a new instance of BlobServicesClient.
func (c *ClientFactory) NewBlobServicesClient() *BlobServicesClient {
	return &BlobServicesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewDeletedAccountsClient creates a new instance of DeletedAccountsClient.
func (c *ClientFactory) NewDeletedAccountsClient() *DeletedAccountsClient {
	return &DeletedAccountsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewEncryptionScopesClient creates a new instance of EncryptionScopesClient.
func (c *ClientFactory) NewEncryptionScopesClient() *EncryptionScopesClient {
	return &EncryptionScopesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewFileServicesClient creates a new instance of FileServicesClient.
func (c *ClientFactory) NewFileServicesClient() *FileServicesClient {
	return &FileServicesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewFileSharesClient creates a new instance of FileSharesClient.
func (c *ClientFactory) NewFileSharesClient() *FileSharesClient {
	return &FileSharesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewLocalUsersClient creates a new instance of LocalUsersClient.
func (c *ClientFactory) NewLocalUsersClient() *LocalUsersClient {
	return &LocalUsersClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewManagementPoliciesClient creates a new instance of ManagementPoliciesClient.
func (c *ClientFactory) NewManagementPoliciesClient() *ManagementPoliciesClient {
	return &ManagementPoliciesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewNetworkSecurityPerimeterConfigurationsClient creates a new instance of NetworkSecurityPerimeterConfigurationsClient.
func (c *ClientFactory) NewNetworkSecurityPerimeterConfigurationsClient() *NetworkSecurityPerimeterConfigurationsClient {
	return &NetworkSecurityPerimeterConfigurationsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewObjectReplicationPoliciesClient creates a new instance of ObjectReplicationPoliciesClient.
func (c *ClientFactory) NewObjectReplicationPoliciesClient() *ObjectReplicationPoliciesClient {
	return &ObjectReplicationPoliciesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewOperationsClient creates a new instance of OperationsClient.
func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	return &OperationsClient{
		internal: c.internal,
	}
}

// NewPrivateEndpointConnectionsClient creates a new instance of PrivateEndpointConnectionsClient.
func (c *ClientFactory) NewPrivateEndpointConnectionsClient() *PrivateEndpointConnectionsClient {
	return &PrivateEndpointConnectionsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewPrivateLinkResourcesClient creates a new instance of PrivateLinkResourcesClient.
func (c *ClientFactory) NewPrivateLinkResourcesClient() *PrivateLinkResourcesClient {
	return &PrivateLinkResourcesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewQueueClient creates a new instance of QueueClient.
func (c *ClientFactory) NewQueueClient() *QueueClient {
	return &QueueClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewQueueServicesClient creates a new instance of QueueServicesClient.
func (c *ClientFactory) NewQueueServicesClient() *QueueServicesClient {
	return &QueueServicesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSKUsClient creates a new instance of SKUsClient.
func (c *ClientFactory) NewSKUsClient() *SKUsClient {
	return &SKUsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewTableClient creates a new instance of TableClient.
func (c *ClientFactory) NewTableClient() *TableClient {
	return &TableClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewTableServicesClient creates a new instance of TableServicesClient.
func (c *ClientFactory) NewTableServicesClient() *TableServicesClient {
	return &TableServicesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewUsagesClient creates a new instance of UsagesClient.
func (c *ClientFactory) NewUsagesClient() *UsagesClient {
	return &UsagesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}
