//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armfabric

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/fabric/armfabric"
	moduleVersion = "v1.0.0"
)

// ActionType - Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
type ActionType string

const (
	ActionTypeInternal ActionType = "Internal"
)

// PossibleActionTypeValues returns the possible values for the ActionType const type.
func PossibleActionTypeValues() []ActionType {
	return []ActionType{
		ActionTypeInternal,
	}
}

// CheckNameAvailabilityReason - The reason why the given name is not available.
type CheckNameAvailabilityReason string

const (
	CheckNameAvailabilityReasonAlreadyExists CheckNameAvailabilityReason = "AlreadyExists"
	CheckNameAvailabilityReasonInvalid       CheckNameAvailabilityReason = "Invalid"
)

// PossibleCheckNameAvailabilityReasonValues returns the possible values for the CheckNameAvailabilityReason const type.
func PossibleCheckNameAvailabilityReasonValues() []CheckNameAvailabilityReason {
	return []CheckNameAvailabilityReason{
		CheckNameAvailabilityReasonAlreadyExists,
		CheckNameAvailabilityReasonInvalid,
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

// Origin - The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
// value is "user,system"
type Origin string

const (
	OriginSystem     Origin = "system"
	OriginUser       Origin = "user"
	OriginUserSystem Origin = "user,system"
)

// PossibleOriginValues returns the possible values for the Origin const type.
func PossibleOriginValues() []Origin {
	return []Origin{
		OriginSystem,
		OriginUser,
		OriginUserSystem,
	}
}

// ProvisioningState - The resource provisioning state.
type ProvisioningState string

const (
	// ProvisioningStateCanceled - Resource creation was canceled.
	ProvisioningStateCanceled ProvisioningState = "Canceled"
	// ProvisioningStateDeleting - Resource is deleting
	ProvisioningStateDeleting ProvisioningState = "Deleting"
	// ProvisioningStateFailed - Resource creation failed.
	ProvisioningStateFailed ProvisioningState = "Failed"
	// ProvisioningStateProvisioning - Resource is provisioning
	ProvisioningStateProvisioning ProvisioningState = "Provisioning"
	// ProvisioningStateSucceeded - Resource has been created.
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	// ProvisioningStateUpdating - Resource is updating
	ProvisioningStateUpdating ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateCanceled,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateProvisioning,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
	}
}

// RpSKUTier - The name of the Azure pricing tier to which the SKU applies.
type RpSKUTier string

const (
	// RpSKUTierFabric - Fabric tier
	RpSKUTierFabric RpSKUTier = "Fabric"
)

// PossibleRpSKUTierValues returns the possible values for the RpSKUTier const type.
func PossibleRpSKUTierValues() []RpSKUTier {
	return []RpSKUTier{
		RpSKUTierFabric,
	}
}

// State - The resource provisioning state.
type State string

const (
	// StateActive - Resource is active
	StateActive State = "Active"
	// StateCanceled - Resource creation was canceled.
	StateCanceled State = "Canceled"
	// StateDeleting - Resource is deleting
	StateDeleting State = "Deleting"
	// StateFailed - Resource creation failed.
	StateFailed State = "Failed"
	// StatePaused - Resource is paused
	StatePaused State = "Paused"
	// StatePausing - Resource is pausing
	StatePausing State = "Pausing"
	// StatePreparing - Resource is preparing
	StatePreparing State = "Preparing"
	// StateProvisioning - Resource is provisioning
	StateProvisioning State = "Provisioning"
	// StateResuming - Resource is resuming
	StateResuming State = "Resuming"
	// StateScaling - Resource is scaling
	StateScaling State = "Scaling"
	// StateSucceeded - Resource has been created.
	StateSucceeded State = "Succeeded"
	// StateSuspended - Resource is suspended
	StateSuspended State = "Suspended"
	// StateSuspending - Resource is suspending
	StateSuspending State = "Suspending"
	// StateUpdating - Resource is updating
	StateUpdating State = "Updating"
)

// PossibleStateValues returns the possible values for the State const type.
func PossibleStateValues() []State {
	return []State{
		StateActive,
		StateCanceled,
		StateDeleting,
		StateFailed,
		StatePaused,
		StatePausing,
		StatePreparing,
		StateProvisioning,
		StateResuming,
		StateScaling,
		StateSucceeded,
		StateSuspended,
		StateSuspending,
		StateUpdating,
	}
}
