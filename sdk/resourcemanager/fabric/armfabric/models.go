//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armfabric

import "time"

// Capacity - Fabric Capacity resource
type Capacity struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// REQUIRED; The SKU details
	SKU *RpSKU

	// The resource-specific properties for this resource.
	Properties *CapacityProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// CapacityAdministration - The administration properties of the Fabric capacity resource
type CapacityAdministration struct {
	// REQUIRED; An array of administrator user identities.
	Members []*string
}

// CapacityAdministrationUpdate - The administration properties of the Fabric capacity resource
type CapacityAdministrationUpdate struct {
	// An array of administrator user identities.
	Members []*string
}

// CapacityListResult - The response of a FabricCapacity list operation.
type CapacityListResult struct {
	// REQUIRED; The FabricCapacity items on this page
	Value []*Capacity

	// The link to the next page of items
	NextLink *string
}

// CapacityProperties - The Microsoft Fabric capacity properties.
type CapacityProperties struct {
	// REQUIRED; The capacity administration
	Administration *CapacityAdministration

	// READ-ONLY; The current state of Microsoft Fabric resource. The state is to indicate more states outside of resource provisioning.
	State *State

	// READ-ONLY; The current deployment state of Microsoft Fabric resource. The provisioningState is to indicate states for resource
	// provisioning.
	ProvisioningState *ProvisioningState
}

// CapacityUpdate - The type used for update operations of the FabricCapacity.
type CapacityUpdate struct {
	// The updatable properties of the FabricCapacity.
	Properties *CapacityUpdateProperties

	// The SKU details
	SKU *RpSKUUpdate

	// Resource tags.
	Tags map[string]*string
}

// CapacityUpdateProperties - The updatable properties of the FabricCapacity.
type CapacityUpdateProperties struct {
	// The capacity administration
	Administration *CapacityAdministrationUpdate
}

// CheckNameAvailabilityRequest - The check availability request body.
type CheckNameAvailabilityRequest struct {
	// The name of the resource for which availability needs to be checked.
	Name *string

	// The resource type.
	Type *string
}

// CheckNameAvailabilityResponse - The check availability result.
type CheckNameAvailabilityResponse struct {
	// Detailed reason why the given name is available.
	Message *string

	// Indicates if the resource name is available.
	NameAvailable *bool

	// The reason why the given name is not available.
	Reason *CheckNameAvailabilityReason
}

// Operation - Details of a REST API operation, returned from the Resource Provider Operations API
type Operation struct {
	// Localized display information for this particular operation.
	Display *OperationDisplay

	// READ-ONLY; Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
	ActionType *ActionType

	// READ-ONLY; Whether the operation applies to data-plane. This is "true" for data-plane operations and "false" for ARM/control-plane
	// operations.
	IsDataAction *bool

	// READ-ONLY; The name of the operation, as per Resource-Based Access Control (RBAC). Examples: "Microsoft.Compute/virtualMachines/write",
	// "Microsoft.Compute/virtualMachines/capture/action"
	Name *string

	// READ-ONLY; The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
	// value is "user,system"
	Origin *Origin
}

// OperationDisplay - Localized display information for this particular operation.
type OperationDisplay struct {
	// READ-ONLY; The short, localized friendly description of the operation; suitable for tool tips and detailed views.
	Description *string

	// READ-ONLY; The concise, localized friendly name for the operation; suitable for dropdowns. E.g. "Create or Update Virtual
	// Machine", "Restart Virtual Machine".
	Operation *string

	// READ-ONLY; The localized friendly form of the resource provider name, e.g. "Microsoft Monitoring Insights" or "Microsoft
	// Compute".
	Provider *string

	// READ-ONLY; The localized friendly name of the resource type related to this operation. E.g. "Virtual Machines" or "Job
	// Schedule Collections".
	Resource *string
}

// OperationListResult - A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to
// get the next set of results.
type OperationListResult struct {
	// READ-ONLY; URL to get the next set of operation list results (if there are any).
	NextLink *string

	// READ-ONLY; List of operations supported by the resource provider
	Value []*Operation
}

// RpSKU - Represents the SKU name and Azure pricing tier for Microsoft Fabric capacity resource.
type RpSKU struct {
	// REQUIRED; The name of the SKU level.
	Name *string

	// REQUIRED; The name of the Azure pricing tier to which the SKU applies.
	Tier *RpSKUTier
}

// RpSKUDetailsForExistingResource - An object that represents SKU details for existing resources
type RpSKUDetailsForExistingResource struct {
	// REQUIRED; The resource type
	ResourceType *string

	// REQUIRED; The SKU details
	SKU *RpSKU
}

// RpSKUDetailsForNewResource - The SKU details
type RpSKUDetailsForNewResource struct {
	// REQUIRED; The list of available locations for the SKU
	Locations []*string

	// REQUIRED; The SKU's name
	Name *string

	// REQUIRED; The resource type
	ResourceType *string
}

// RpSKUEnumerationForExistingResourceResult - An object that represents enumerating SKUs for existing resources
type RpSKUEnumerationForExistingResourceResult struct {
	// REQUIRED; The SKU details
	Value []*RpSKUDetailsForExistingResource

	// Url for the next page. Null if no more pages available
	NextLink *string
}

// RpSKUEnumerationForNewResourceResult - An object that represents enumerating SKUs for new resources.
type RpSKUEnumerationForNewResourceResult struct {
	// REQUIRED; The collection of available SKUs for new resources
	Value []*RpSKUDetailsForNewResource

	// Url for the next page. Null if no more pages available
	NextLink *string
}

// RpSKUUpdate - Represents the SKU name and Azure pricing tier for Microsoft Fabric capacity resource.
type RpSKUUpdate struct {
	// The name of the SKU level.
	Name *string

	// The name of the Azure pricing tier to which the SKU applies.
	Tier *RpSKUTier
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time

	// The identity that created the resource.
	CreatedBy *string

	// The type of identity that created the resource.
	CreatedByType *CreatedByType

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time

	// The identity that last modified the resource.
	LastModifiedBy *string

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType
}
