//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdeploymentstacks

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armdeploymentstacks"
	moduleVersion = "v1.0.0"
)

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

// DenySettingsMode - denySettings Mode that defines denied actions.
type DenySettingsMode string

const (
	// DenySettingsModeDenyDelete - Authorized users are able to read and modify the resources, but cannot delete.
	DenySettingsModeDenyDelete DenySettingsMode = "denyDelete"
	// DenySettingsModeDenyWriteAndDelete - Authorized users can read from a resource, but cannot modify or delete it.
	DenySettingsModeDenyWriteAndDelete DenySettingsMode = "denyWriteAndDelete"
	// DenySettingsModeNone - No denyAssignments have been applied.
	DenySettingsModeNone DenySettingsMode = "none"
)

// PossibleDenySettingsModeValues returns the possible values for the DenySettingsMode const type.
func PossibleDenySettingsModeValues() []DenySettingsMode {
	return []DenySettingsMode{
		DenySettingsModeDenyDelete,
		DenySettingsModeDenyWriteAndDelete,
		DenySettingsModeNone,
	}
}

// DenyStatusMode - denyAssignment settings applied to the resource.
type DenyStatusMode string

const (
	// DenyStatusModeDenyDelete - Authorized users are able to read and modify the resources, but cannot delete.
	DenyStatusModeDenyDelete DenyStatusMode = "denyDelete"
	// DenyStatusModeDenyWriteAndDelete - Authorized users can only read from a resource, but cannot modify or delete it.
	DenyStatusModeDenyWriteAndDelete DenyStatusMode = "denyWriteAndDelete"
	// DenyStatusModeInapplicable - denyAssignments are not supported on resources outside the scope of the deployment stack.
	DenyStatusModeInapplicable DenyStatusMode = "inapplicable"
	// DenyStatusModeNone - No denyAssignments have been applied.
	DenyStatusModeNone DenyStatusMode = "none"
	// DenyStatusModeNotSupported - Resource type does not support denyAssignments.
	DenyStatusModeNotSupported DenyStatusMode = "notSupported"
	// DenyStatusModeRemovedBySystem - Deny assignment has been removed by Azure due to a resource management change (management
	// group move, etc.)
	DenyStatusModeRemovedBySystem DenyStatusMode = "removedBySystem"
)

// PossibleDenyStatusModeValues returns the possible values for the DenyStatusMode const type.
func PossibleDenyStatusModeValues() []DenyStatusMode {
	return []DenyStatusMode{
		DenyStatusModeDenyDelete,
		DenyStatusModeDenyWriteAndDelete,
		DenyStatusModeInapplicable,
		DenyStatusModeNone,
		DenyStatusModeNotSupported,
		DenyStatusModeRemovedBySystem,
	}
}

// DeploymentStackProvisioningState - State of the deployment stack.
type DeploymentStackProvisioningState string

const (
	DeploymentStackProvisioningStateCanceled                DeploymentStackProvisioningState = "canceled"
	DeploymentStackProvisioningStateCanceling               DeploymentStackProvisioningState = "canceling"
	DeploymentStackProvisioningStateCreating                DeploymentStackProvisioningState = "creating"
	DeploymentStackProvisioningStateDeleting                DeploymentStackProvisioningState = "deleting"
	DeploymentStackProvisioningStateDeletingResources       DeploymentStackProvisioningState = "deletingResources"
	DeploymentStackProvisioningStateDeploying               DeploymentStackProvisioningState = "deploying"
	DeploymentStackProvisioningStateFailed                  DeploymentStackProvisioningState = "failed"
	DeploymentStackProvisioningStateSucceeded               DeploymentStackProvisioningState = "succeeded"
	DeploymentStackProvisioningStateUpdatingDenyAssignments DeploymentStackProvisioningState = "updatingDenyAssignments"
	DeploymentStackProvisioningStateValidating              DeploymentStackProvisioningState = "validating"
	DeploymentStackProvisioningStateWaiting                 DeploymentStackProvisioningState = "waiting"
)

// PossibleDeploymentStackProvisioningStateValues returns the possible values for the DeploymentStackProvisioningState const type.
func PossibleDeploymentStackProvisioningStateValues() []DeploymentStackProvisioningState {
	return []DeploymentStackProvisioningState{
		DeploymentStackProvisioningStateCanceled,
		DeploymentStackProvisioningStateCanceling,
		DeploymentStackProvisioningStateCreating,
		DeploymentStackProvisioningStateDeleting,
		DeploymentStackProvisioningStateDeletingResources,
		DeploymentStackProvisioningStateDeploying,
		DeploymentStackProvisioningStateFailed,
		DeploymentStackProvisioningStateSucceeded,
		DeploymentStackProvisioningStateUpdatingDenyAssignments,
		DeploymentStackProvisioningStateValidating,
		DeploymentStackProvisioningStateWaiting,
	}
}

// DeploymentStacksDeleteDetachEnum - Specifies an action for a newly unmanaged resource. Delete will attempt to delete the
// resource from Azure. Detach will leave the resource in it's current state.
type DeploymentStacksDeleteDetachEnum string

const (
	DeploymentStacksDeleteDetachEnumDelete DeploymentStacksDeleteDetachEnum = "delete"
	DeploymentStacksDeleteDetachEnumDetach DeploymentStacksDeleteDetachEnum = "detach"
)

// PossibleDeploymentStacksDeleteDetachEnumValues returns the possible values for the DeploymentStacksDeleteDetachEnum const type.
func PossibleDeploymentStacksDeleteDetachEnumValues() []DeploymentStacksDeleteDetachEnum {
	return []DeploymentStacksDeleteDetachEnum{
		DeploymentStacksDeleteDetachEnumDelete,
		DeploymentStacksDeleteDetachEnumDetach,
	}
}

// ResourceStatusMode - Current management state of the resource in the deployment stack.
type ResourceStatusMode string

const (
	// ResourceStatusModeDeleteFailed - Unable to delete the resource from Azure. The delete will be retried on the next stack
	// deployment, or can be deleted manually.
	ResourceStatusModeDeleteFailed ResourceStatusMode = "deleteFailed"
	// ResourceStatusModeManaged - This resource is managed by the deployment stack.
	ResourceStatusModeManaged ResourceStatusMode = "managed"
	// ResourceStatusModeRemoveDenyFailed - Unable to remove the deny assignment on resource.
	ResourceStatusModeRemoveDenyFailed ResourceStatusMode = "removeDenyFailed"
)

// PossibleResourceStatusModeValues returns the possible values for the ResourceStatusMode const type.
func PossibleResourceStatusModeValues() []ResourceStatusMode {
	return []ResourceStatusMode{
		ResourceStatusModeDeleteFailed,
		ResourceStatusModeManaged,
		ResourceStatusModeRemoveDenyFailed,
	}
}

type UnmanageActionManagementGroupMode string

const (
	UnmanageActionManagementGroupModeDelete UnmanageActionManagementGroupMode = "delete"
	UnmanageActionManagementGroupModeDetach UnmanageActionManagementGroupMode = "detach"
)

// PossibleUnmanageActionManagementGroupModeValues returns the possible values for the UnmanageActionManagementGroupMode const type.
func PossibleUnmanageActionManagementGroupModeValues() []UnmanageActionManagementGroupMode {
	return []UnmanageActionManagementGroupMode{
		UnmanageActionManagementGroupModeDelete,
		UnmanageActionManagementGroupModeDetach,
	}
}

type UnmanageActionResourceGroupMode string

const (
	UnmanageActionResourceGroupModeDelete UnmanageActionResourceGroupMode = "delete"
	UnmanageActionResourceGroupModeDetach UnmanageActionResourceGroupMode = "detach"
)

// PossibleUnmanageActionResourceGroupModeValues returns the possible values for the UnmanageActionResourceGroupMode const type.
func PossibleUnmanageActionResourceGroupModeValues() []UnmanageActionResourceGroupMode {
	return []UnmanageActionResourceGroupMode{
		UnmanageActionResourceGroupModeDelete,
		UnmanageActionResourceGroupModeDetach,
	}
}

type UnmanageActionResourceMode string

const (
	UnmanageActionResourceModeDelete UnmanageActionResourceMode = "delete"
	UnmanageActionResourceModeDetach UnmanageActionResourceMode = "detach"
)

// PossibleUnmanageActionResourceModeValues returns the possible values for the UnmanageActionResourceMode const type.
func PossibleUnmanageActionResourceModeValues() []UnmanageActionResourceMode {
	return []UnmanageActionResourceMode{
		UnmanageActionResourceModeDelete,
		UnmanageActionResourceModeDetach,
	}
}
