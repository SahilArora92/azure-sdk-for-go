//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armavs

// AddonPropertiesClassification provides polymorphic access to related types.
// Call the interface's GetAddonProperties() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AddonArcProperties, *AddonHcxProperties, *AddonProperties, *AddonSrmProperties, *AddonVrProperties
type AddonPropertiesClassification interface {
	// GetAddonProperties returns the AddonProperties content of the underlying type.
	GetAddonProperties() *AddonProperties
}

// PlacementPolicyPropertiesClassification provides polymorphic access to related types.
// Call the interface's GetPlacementPolicyProperties() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *PlacementPolicyProperties, *VMHostPlacementPolicyProperties, *VMPlacementPolicyProperties
type PlacementPolicyPropertiesClassification interface {
	// GetPlacementPolicyProperties returns the PlacementPolicyProperties content of the underlying type.
	GetPlacementPolicyProperties() *PlacementPolicyProperties
}

// ScriptExecutionParameterClassification provides polymorphic access to related types.
// Call the interface's GetScriptExecutionParameter() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *PSCredentialExecutionParameter, *ScriptExecutionParameter, *ScriptSecureStringExecutionParameter, *ScriptStringExecutionParameter
type ScriptExecutionParameterClassification interface {
	// GetScriptExecutionParameter returns the ScriptExecutionParameter content of the underlying type.
	GetScriptExecutionParameter() *ScriptExecutionParameter
}

// WorkloadNetworkDhcpEntityClassification provides polymorphic access to related types.
// Call the interface's GetWorkloadNetworkDhcpEntity() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *WorkloadNetworkDhcpEntity, *WorkloadNetworkDhcpRelay, *WorkloadNetworkDhcpServer
type WorkloadNetworkDhcpEntityClassification interface {
	// GetWorkloadNetworkDhcpEntity returns the WorkloadNetworkDhcpEntity content of the underlying type.
	GetWorkloadNetworkDhcpEntity() *WorkloadNetworkDhcpEntity
}

// WorkloadNetworkDhcpEntityUpdateClassification provides polymorphic access to related types.
// Call the interface's GetWorkloadNetworkDhcpEntityUpdate() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *WorkloadNetworkDhcpEntityUpdate, *WorkloadNetworkDhcpRelayUpdate, *WorkloadNetworkDhcpServerUpdate
type WorkloadNetworkDhcpEntityUpdateClassification interface {
	// GetWorkloadNetworkDhcpEntityUpdate returns the WorkloadNetworkDhcpEntityUpdate content of the underlying type.
	GetWorkloadNetworkDhcpEntityUpdate() *WorkloadNetworkDhcpEntityUpdate
}
