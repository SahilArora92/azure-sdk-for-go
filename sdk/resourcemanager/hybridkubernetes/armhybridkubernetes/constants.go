//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybridkubernetes

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridkubernetes/armhybridkubernetes"
	moduleVersion = "v2.0.0"
)

// AuthenticationMethod - The mode of client authentication.
type AuthenticationMethod string

const (
	AuthenticationMethodAAD   AuthenticationMethod = "AAD"
	AuthenticationMethodToken AuthenticationMethod = "Token"
)

// PossibleAuthenticationMethodValues returns the possible values for the AuthenticationMethod const type.
func PossibleAuthenticationMethodValues() []AuthenticationMethod {
	return []AuthenticationMethod{
		AuthenticationMethodAAD,
		AuthenticationMethodToken,
	}
}

// AutoUpgradeOptions - Indicates whether the Arc agents on the be upgraded automatically to the latest version. Defaults
// to Enabled.
type AutoUpgradeOptions string

const (
	AutoUpgradeOptionsDisabled AutoUpgradeOptions = "Disabled"
	AutoUpgradeOptionsEnabled  AutoUpgradeOptions = "Enabled"
)

// PossibleAutoUpgradeOptionsValues returns the possible values for the AutoUpgradeOptions const type.
func PossibleAutoUpgradeOptionsValues() []AutoUpgradeOptions {
	return []AutoUpgradeOptions{
		AutoUpgradeOptionsDisabled,
		AutoUpgradeOptionsEnabled,
	}
}

// AzureHybridBenefit - Indicates whether Azure Hybrid Benefit is opted in
type AzureHybridBenefit string

const (
	AzureHybridBenefitFalse         AzureHybridBenefit = "False"
	AzureHybridBenefitNotApplicable AzureHybridBenefit = "NotApplicable"
	AzureHybridBenefitTrue          AzureHybridBenefit = "True"
)

// PossibleAzureHybridBenefitValues returns the possible values for the AzureHybridBenefit const type.
func PossibleAzureHybridBenefitValues() []AzureHybridBenefit {
	return []AzureHybridBenefit{
		AzureHybridBenefitFalse,
		AzureHybridBenefitNotApplicable,
		AzureHybridBenefitTrue,
	}
}

// ConnectedClusterKind - Indicates the kind of Arc connected cluster based on host infrastructure.
type ConnectedClusterKind string

const (
	ConnectedClusterKindProvisionedCluster ConnectedClusterKind = "ProvisionedCluster"
)

// PossibleConnectedClusterKindValues returns the possible values for the ConnectedClusterKind const type.
func PossibleConnectedClusterKindValues() []ConnectedClusterKind {
	return []ConnectedClusterKind{
		ConnectedClusterKindProvisionedCluster,
	}
}

// ConnectivityStatus - Represents the connectivity status of the connected cluster.
type ConnectivityStatus string

const (
	ConnectivityStatusConnected  ConnectivityStatus = "Connected"
	ConnectivityStatusConnecting ConnectivityStatus = "Connecting"
	ConnectivityStatusExpired    ConnectivityStatus = "Expired"
	ConnectivityStatusOffline    ConnectivityStatus = "Offline"
)

// PossibleConnectivityStatusValues returns the possible values for the ConnectivityStatus const type.
func PossibleConnectivityStatusValues() []ConnectivityStatus {
	return []ConnectivityStatus{
		ConnectivityStatusConnected,
		ConnectivityStatusConnecting,
		ConnectivityStatusExpired,
		ConnectivityStatusOffline,
	}
}

// CreatedByType - The type of identity that created the resource.
type CreatedByType string

const (
	CreatedByTypeApplication     CreatedByType = "Application"
	CreatedByTypeKey             CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser            CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// LastModifiedByType - The type of identity that last modified the resource.
type LastModifiedByType string

const (
	LastModifiedByTypeApplication     LastModifiedByType = "Application"
	LastModifiedByTypeKey             LastModifiedByType = "Key"
	LastModifiedByTypeManagedIdentity LastModifiedByType = "ManagedIdentity"
	LastModifiedByTypeUser            LastModifiedByType = "User"
)

// PossibleLastModifiedByTypeValues returns the possible values for the LastModifiedByType const type.
func PossibleLastModifiedByTypeValues() []LastModifiedByType {
	return []LastModifiedByType{
		LastModifiedByTypeApplication,
		LastModifiedByTypeKey,
		LastModifiedByTypeManagedIdentity,
		LastModifiedByTypeUser,
	}
}

// PrivateLinkState - Property which describes the state of private link on a connected cluster resource.
type PrivateLinkState string

const (
	PrivateLinkStateDisabled PrivateLinkState = "Disabled"
	PrivateLinkStateEnabled  PrivateLinkState = "Enabled"
)

// PossiblePrivateLinkStateValues returns the possible values for the PrivateLinkState const type.
func PossiblePrivateLinkStateValues() []PrivateLinkState {
	return []PrivateLinkState{
		PrivateLinkStateDisabled,
		PrivateLinkStateEnabled,
	}
}

// ProvisioningState - The current deployment state of connectedClusters.
type ProvisioningState string

const (
	ProvisioningStateAccepted     ProvisioningState = "Accepted"
	ProvisioningStateCanceled     ProvisioningState = "Canceled"
	ProvisioningStateDeleting     ProvisioningState = "Deleting"
	ProvisioningStateFailed       ProvisioningState = "Failed"
	ProvisioningStateProvisioning ProvisioningState = "Provisioning"
	ProvisioningStateSucceeded    ProvisioningState = "Succeeded"
	ProvisioningStateUpdating     ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateAccepted,
		ProvisioningStateCanceled,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateProvisioning,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
	}
}

// ResourceIdentityType - The type of identity used for the connected cluster. The type 'SystemAssigned, includes a system
// created identity. The type 'None' means no identity is assigned to the connected cluster.
type ResourceIdentityType string

const (
	ResourceIdentityTypeNone           ResourceIdentityType = "None"
	ResourceIdentityTypeSystemAssigned ResourceIdentityType = "SystemAssigned"
)

// PossibleResourceIdentityTypeValues returns the possible values for the ResourceIdentityType const type.
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return []ResourceIdentityType{
		ResourceIdentityTypeNone,
		ResourceIdentityTypeSystemAssigned,
	}
}
