//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armguestconfiguration

const (
	moduleName    = "armguestconfiguration"
	moduleVersion = "v0.4.0"
)

// ActionAfterReboot - Specifies what happens after a reboot during the application of a configuration. The possible values
// are ContinueConfiguration and StopConfiguration
type ActionAfterReboot string

const (
	ActionAfterRebootContinueConfiguration ActionAfterReboot = "ContinueConfiguration"
	ActionAfterRebootStopConfiguration     ActionAfterReboot = "StopConfiguration"
)

// PossibleActionAfterRebootValues returns the possible values for the ActionAfterReboot const type.
func PossibleActionAfterRebootValues() []ActionAfterReboot {
	return []ActionAfterReboot{
		ActionAfterRebootContinueConfiguration,
		ActionAfterRebootStopConfiguration,
	}
}

// AssignmentType - Specifies the assignment type and execution of the configuration. Possible values are Audit, DeployAndAutoCorrect,
// ApplyAndAutoCorrect and ApplyAndMonitor.
type AssignmentType string

const (
	AssignmentTypeApplyAndAutoCorrect  AssignmentType = "ApplyAndAutoCorrect"
	AssignmentTypeApplyAndMonitor      AssignmentType = "ApplyAndMonitor"
	AssignmentTypeAudit                AssignmentType = "Audit"
	AssignmentTypeDeployAndAutoCorrect AssignmentType = "DeployAndAutoCorrect"
)

// PossibleAssignmentTypeValues returns the possible values for the AssignmentType const type.
func PossibleAssignmentTypeValues() []AssignmentType {
	return []AssignmentType{
		AssignmentTypeApplyAndAutoCorrect,
		AssignmentTypeApplyAndMonitor,
		AssignmentTypeAudit,
		AssignmentTypeDeployAndAutoCorrect,
	}
}

// ComplianceStatus - A value indicating compliance status of the machine for the assigned guest configuration.
type ComplianceStatus string

const (
	ComplianceStatusCompliant    ComplianceStatus = "Compliant"
	ComplianceStatusNonCompliant ComplianceStatus = "NonCompliant"
	ComplianceStatusPending      ComplianceStatus = "Pending"
)

// PossibleComplianceStatusValues returns the possible values for the ComplianceStatus const type.
func PossibleComplianceStatusValues() []ComplianceStatus {
	return []ComplianceStatus{
		ComplianceStatusCompliant,
		ComplianceStatusNonCompliant,
		ComplianceStatusPending,
	}
}

// ConfigurationMode - Specifies how the LCM(Local Configuration Manager) actually applies the configuration to the target
// nodes. Possible values are ApplyOnly, ApplyAndMonitor, and ApplyAndAutoCorrect.
type ConfigurationMode string

const (
	ConfigurationModeApplyAndAutoCorrect ConfigurationMode = "ApplyAndAutoCorrect"
	ConfigurationModeApplyAndMonitor     ConfigurationMode = "ApplyAndMonitor"
	ConfigurationModeApplyOnly           ConfigurationMode = "ApplyOnly"
)

// PossibleConfigurationModeValues returns the possible values for the ConfigurationMode const type.
func PossibleConfigurationModeValues() []ConfigurationMode {
	return []ConfigurationMode{
		ConfigurationModeApplyAndAutoCorrect,
		ConfigurationModeApplyAndMonitor,
		ConfigurationModeApplyOnly,
	}
}

// Kind - Kind of the guest configuration. For example:DSC
type Kind string

const (
	KindDSC Kind = "DSC"
)

// PossibleKindValues returns the possible values for the Kind const type.
func PossibleKindValues() []Kind {
	return []Kind{
		KindDSC,
	}
}

// ProvisioningState - The provisioning state, which only appears in the response.
type ProvisioningState string

const (
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateCreated   ProvisioningState = "Created"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateCanceled,
		ProvisioningStateCreated,
		ProvisioningStateFailed,
		ProvisioningStateSucceeded,
	}
}

// Type - Type of report, Consistency or Initial
type Type string

const (
	TypeConsistency Type = "Consistency"
	TypeInitial     Type = "Initial"
)

// PossibleTypeValues returns the possible values for the Type const type.
func PossibleTypeValues() []Type {
	return []Type{
		TypeConsistency,
		TypeInitial,
	}
}
