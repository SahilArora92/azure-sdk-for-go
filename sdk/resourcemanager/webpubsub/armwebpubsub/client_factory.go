//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armwebpubsub

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
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
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

// NewClient creates a new instance of Client.
func (c *ClientFactory) NewClient() *Client {
	return &Client{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewCustomCertificatesClient creates a new instance of CustomCertificatesClient.
func (c *ClientFactory) NewCustomCertificatesClient() *CustomCertificatesClient {
	return &CustomCertificatesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewCustomDomainsClient creates a new instance of CustomDomainsClient.
func (c *ClientFactory) NewCustomDomainsClient() *CustomDomainsClient {
	return &CustomDomainsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewHubsClient creates a new instance of HubsClient.
func (c *ClientFactory) NewHubsClient() *HubsClient {
	return &HubsClient{
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

// NewReplicaSharedPrivateLinkResourcesClient creates a new instance of ReplicaSharedPrivateLinkResourcesClient.
func (c *ClientFactory) NewReplicaSharedPrivateLinkResourcesClient() *ReplicaSharedPrivateLinkResourcesClient {
	return &ReplicaSharedPrivateLinkResourcesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewReplicasClient creates a new instance of ReplicasClient.
func (c *ClientFactory) NewReplicasClient() *ReplicasClient {
	return &ReplicasClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSharedPrivateLinkResourcesClient creates a new instance of SharedPrivateLinkResourcesClient.
func (c *ClientFactory) NewSharedPrivateLinkResourcesClient() *SharedPrivateLinkResourcesClient {
	return &SharedPrivateLinkResourcesClient{
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
