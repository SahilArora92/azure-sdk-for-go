//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhardwaresecuritymodules

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
)

// ClientFactory is a client factory used to create any client in this module.
// Don't use this type directly, use NewClientFactory instead.
type ClientFactory struct {
	subscriptionID string
	jobID          string
	internal       *arm.Client
}

// NewClientFactory creates a new instance of ClientFactory with the specified values.
// The parameter values will be propagated to any client created from this factory.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - jobID - The id returned as part of the backup request
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(subscriptionID string, jobID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	internal, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		subscriptionID: subscriptionID,
		jobID:          jobID,
		internal:       internal,
	}, nil
}

// NewCloudHsmClusterPrivateEndpointConnectionsClient creates a new instance of CloudHsmClusterPrivateEndpointConnectionsClient.
func (c *ClientFactory) NewCloudHsmClusterPrivateEndpointConnectionsClient() *CloudHsmClusterPrivateEndpointConnectionsClient {
	return &CloudHsmClusterPrivateEndpointConnectionsClient{
		subscriptionID: c.subscriptionID,
		jobID:          c.jobID,
		internal:       c.internal,
	}
}

// NewCloudHsmClusterPrivateLinkResourcesClient creates a new instance of CloudHsmClusterPrivateLinkResourcesClient.
func (c *ClientFactory) NewCloudHsmClusterPrivateLinkResourcesClient() *CloudHsmClusterPrivateLinkResourcesClient {
	return &CloudHsmClusterPrivateLinkResourcesClient{
		subscriptionID: c.subscriptionID,
		jobID:          c.jobID,
		internal:       c.internal,
	}
}

// NewCloudHsmClustersBackupStatusClient creates a new instance of CloudHsmClustersBackupStatusClient.
func (c *ClientFactory) NewCloudHsmClustersBackupStatusClient() *CloudHsmClustersBackupStatusClient {
	return &CloudHsmClustersBackupStatusClient{
		subscriptionID: c.subscriptionID,
		jobID:          c.jobID,
		internal:       c.internal,
	}
}

// NewCloudHsmClustersClient creates a new instance of CloudHsmClustersClient.
func (c *ClientFactory) NewCloudHsmClustersClient() *CloudHsmClustersClient {
	return &CloudHsmClustersClient{
		subscriptionID: c.subscriptionID,
		jobID:          c.jobID,
		internal:       c.internal,
	}
}

// NewCloudHsmClustersRestoreStatusClient creates a new instance of CloudHsmClustersRestoreStatusClient.
func (c *ClientFactory) NewCloudHsmClustersRestoreStatusClient() *CloudHsmClustersRestoreStatusClient {
	return &CloudHsmClustersRestoreStatusClient{
		subscriptionID: c.subscriptionID,
		jobID:          c.jobID,
		internal:       c.internal,
	}
}

// NewDedicatedHsmClient creates a new instance of DedicatedHsmClient.
func (c *ClientFactory) NewDedicatedHsmClient() *DedicatedHsmClient {
	return &DedicatedHsmClient{
		subscriptionID: c.subscriptionID,
		jobID:          c.jobID,
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
		jobID:          c.jobID,
		internal:       c.internal,
	}
}
