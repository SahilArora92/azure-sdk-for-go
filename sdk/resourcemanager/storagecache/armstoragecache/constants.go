//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstoragecache

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagecache/armstoragecache"
	moduleVersion = "v4.1.0"
)

// AmlFilesystemHealthStateType - List of AML file system health states.
type AmlFilesystemHealthStateType string

const (
	AmlFilesystemHealthStateTypeAvailable     AmlFilesystemHealthStateType = "Available"
	AmlFilesystemHealthStateTypeDegraded      AmlFilesystemHealthStateType = "Degraded"
	AmlFilesystemHealthStateTypeMaintenance   AmlFilesystemHealthStateType = "Maintenance"
	AmlFilesystemHealthStateTypeTransitioning AmlFilesystemHealthStateType = "Transitioning"
	AmlFilesystemHealthStateTypeUnavailable   AmlFilesystemHealthStateType = "Unavailable"
)

// PossibleAmlFilesystemHealthStateTypeValues returns the possible values for the AmlFilesystemHealthStateType const type.
func PossibleAmlFilesystemHealthStateTypeValues() []AmlFilesystemHealthStateType {
	return []AmlFilesystemHealthStateType{
		AmlFilesystemHealthStateTypeAvailable,
		AmlFilesystemHealthStateTypeDegraded,
		AmlFilesystemHealthStateTypeMaintenance,
		AmlFilesystemHealthStateTypeTransitioning,
		AmlFilesystemHealthStateTypeUnavailable,
	}
}

// AmlFilesystemIdentityType - The type of identity used for the resource.
type AmlFilesystemIdentityType string

const (
	AmlFilesystemIdentityTypeNone         AmlFilesystemIdentityType = "None"
	AmlFilesystemIdentityTypeUserAssigned AmlFilesystemIdentityType = "UserAssigned"
)

// PossibleAmlFilesystemIdentityTypeValues returns the possible values for the AmlFilesystemIdentityType const type.
func PossibleAmlFilesystemIdentityTypeValues() []AmlFilesystemIdentityType {
	return []AmlFilesystemIdentityType{
		AmlFilesystemIdentityTypeNone,
		AmlFilesystemIdentityTypeUserAssigned,
	}
}

// AmlFilesystemProvisioningStateType - ARM provisioning state.
type AmlFilesystemProvisioningStateType string

const (
	AmlFilesystemProvisioningStateTypeCanceled  AmlFilesystemProvisioningStateType = "Canceled"
	AmlFilesystemProvisioningStateTypeCreating  AmlFilesystemProvisioningStateType = "Creating"
	AmlFilesystemProvisioningStateTypeDeleting  AmlFilesystemProvisioningStateType = "Deleting"
	AmlFilesystemProvisioningStateTypeFailed    AmlFilesystemProvisioningStateType = "Failed"
	AmlFilesystemProvisioningStateTypeSucceeded AmlFilesystemProvisioningStateType = "Succeeded"
	AmlFilesystemProvisioningStateTypeUpdating  AmlFilesystemProvisioningStateType = "Updating"
)

// PossibleAmlFilesystemProvisioningStateTypeValues returns the possible values for the AmlFilesystemProvisioningStateType const type.
func PossibleAmlFilesystemProvisioningStateTypeValues() []AmlFilesystemProvisioningStateType {
	return []AmlFilesystemProvisioningStateType{
		AmlFilesystemProvisioningStateTypeCanceled,
		AmlFilesystemProvisioningStateTypeCreating,
		AmlFilesystemProvisioningStateTypeDeleting,
		AmlFilesystemProvisioningStateTypeFailed,
		AmlFilesystemProvisioningStateTypeSucceeded,
		AmlFilesystemProvisioningStateTypeUpdating,
	}
}

// AmlFilesystemSquashMode - Squash mode of the AML file system. 'All': User and Group IDs on files will be squashed to the
// provided values for all users on non-trusted systems. 'RootOnly': User and Group IDs on files will be
// squashed to provided values for solely the root user on non-trusted systems. 'None': No squashing of User and Group IDs
// is performed for any users on any systems.
type AmlFilesystemSquashMode string

const (
	AmlFilesystemSquashModeAll      AmlFilesystemSquashMode = "All"
	AmlFilesystemSquashModeNone     AmlFilesystemSquashMode = "None"
	AmlFilesystemSquashModeRootOnly AmlFilesystemSquashMode = "RootOnly"
)

// PossibleAmlFilesystemSquashModeValues returns the possible values for the AmlFilesystemSquashMode const type.
func PossibleAmlFilesystemSquashModeValues() []AmlFilesystemSquashMode {
	return []AmlFilesystemSquashMode{
		AmlFilesystemSquashModeAll,
		AmlFilesystemSquashModeNone,
		AmlFilesystemSquashModeRootOnly,
	}
}

// ArchiveStatusType - The state of the archive operation
type ArchiveStatusType string

const (
	ArchiveStatusTypeCanceled         ArchiveStatusType = "Canceled"
	ArchiveStatusTypeCancelling       ArchiveStatusType = "Cancelling"
	ArchiveStatusTypeCompleted        ArchiveStatusType = "Completed"
	ArchiveStatusTypeFSScanInProgress ArchiveStatusType = "FSScanInProgress"
	ArchiveStatusTypeFailed           ArchiveStatusType = "Failed"
	ArchiveStatusTypeIdle             ArchiveStatusType = "Idle"
	ArchiveStatusTypeInProgress       ArchiveStatusType = "InProgress"
	ArchiveStatusTypeNotConfigured    ArchiveStatusType = "NotConfigured"
)

// PossibleArchiveStatusTypeValues returns the possible values for the ArchiveStatusType const type.
func PossibleArchiveStatusTypeValues() []ArchiveStatusType {
	return []ArchiveStatusType{
		ArchiveStatusTypeCanceled,
		ArchiveStatusTypeCancelling,
		ArchiveStatusTypeCompleted,
		ArchiveStatusTypeFSScanInProgress,
		ArchiveStatusTypeFailed,
		ArchiveStatusTypeIdle,
		ArchiveStatusTypeInProgress,
		ArchiveStatusTypeNotConfigured,
	}
}

// AutoExportJobAdminStatus - The administrative status of the auto export job. Possible values: 'Enable', 'Disable'. Passing
// in a value of 'Disable' will disable the current active auto export job. By default it is set to
// 'Enable'.
type AutoExportJobAdminStatus string

const (
	AutoExportJobAdminStatusDisable AutoExportJobAdminStatus = "Disable"
	AutoExportJobAdminStatusEnable  AutoExportJobAdminStatus = "Enable"
)

// PossibleAutoExportJobAdminStatusValues returns the possible values for the AutoExportJobAdminStatus const type.
func PossibleAutoExportJobAdminStatusValues() []AutoExportJobAdminStatus {
	return []AutoExportJobAdminStatus{
		AutoExportJobAdminStatusDisable,
		AutoExportJobAdminStatusEnable,
	}
}

// AutoExportJobProvisioningStateType - ARM provisioning state.
type AutoExportJobProvisioningStateType string

const (
	AutoExportJobProvisioningStateTypeCanceled  AutoExportJobProvisioningStateType = "Canceled"
	AutoExportJobProvisioningStateTypeCreating  AutoExportJobProvisioningStateType = "Creating"
	AutoExportJobProvisioningStateTypeDeleting  AutoExportJobProvisioningStateType = "Deleting"
	AutoExportJobProvisioningStateTypeFailed    AutoExportJobProvisioningStateType = "Failed"
	AutoExportJobProvisioningStateTypeSucceeded AutoExportJobProvisioningStateType = "Succeeded"
	AutoExportJobProvisioningStateTypeUpdating  AutoExportJobProvisioningStateType = "Updating"
)

// PossibleAutoExportJobProvisioningStateTypeValues returns the possible values for the AutoExportJobProvisioningStateType const type.
func PossibleAutoExportJobProvisioningStateTypeValues() []AutoExportJobProvisioningStateType {
	return []AutoExportJobProvisioningStateType{
		AutoExportJobProvisioningStateTypeCanceled,
		AutoExportJobProvisioningStateTypeCreating,
		AutoExportJobProvisioningStateTypeDeleting,
		AutoExportJobProvisioningStateTypeFailed,
		AutoExportJobProvisioningStateTypeSucceeded,
		AutoExportJobProvisioningStateTypeUpdating,
	}
}

// AutoExportStatusType - The operational state of auto export. InProgress indicates the export is running. Disabling indicates
// the user has requested to disable the export but the disabling is still in progress. Disabled
// indicates auto export has been disabled. DisableFailed indicates the disabling has failed. Failed means the export was
// unable to continue, due to a fatal error.
type AutoExportStatusType string

const (
	AutoExportStatusTypeDisableFailed AutoExportStatusType = "DisableFailed"
	AutoExportStatusTypeDisabled      AutoExportStatusType = "Disabled"
	AutoExportStatusTypeDisabling     AutoExportStatusType = "Disabling"
	AutoExportStatusTypeFailed        AutoExportStatusType = "Failed"
	AutoExportStatusTypeInProgress    AutoExportStatusType = "InProgress"
)

// PossibleAutoExportStatusTypeValues returns the possible values for the AutoExportStatusType const type.
func PossibleAutoExportStatusTypeValues() []AutoExportStatusType {
	return []AutoExportStatusType{
		AutoExportStatusTypeDisableFailed,
		AutoExportStatusTypeDisabled,
		AutoExportStatusTypeDisabling,
		AutoExportStatusTypeFailed,
		AutoExportStatusTypeInProgress,
	}
}

// CacheIdentityType - The type of identity used for the cache
type CacheIdentityType string

const (
	CacheIdentityTypeNone                       CacheIdentityType = "None"
	CacheIdentityTypeSystemAssigned             CacheIdentityType = "SystemAssigned"
	CacheIdentityTypeSystemAssignedUserAssigned CacheIdentityType = "SystemAssigned, UserAssigned"
	CacheIdentityTypeUserAssigned               CacheIdentityType = "UserAssigned"
)

// PossibleCacheIdentityTypeValues returns the possible values for the CacheIdentityType const type.
func PossibleCacheIdentityTypeValues() []CacheIdentityType {
	return []CacheIdentityType{
		CacheIdentityTypeNone,
		CacheIdentityTypeSystemAssigned,
		CacheIdentityTypeSystemAssignedUserAssigned,
		CacheIdentityTypeUserAssigned,
	}
}

// ConflictResolutionMode - How the import job will handle conflicts. For example, if the import job is trying to bring in
// a directory, but a file is at that path, how it handles it. Fail indicates that the import job should
// stop immediately and not do anything with the conflict. Skip indicates that it should pass over the conflict. OverwriteIfDirty
// causes the import job to delete and re-import the file or directory if it
// is a conflicting type, is dirty, or was not previously imported. OverwriteAlways extends OverwriteIfDirty to include releasing
// files that had been restored but were not dirty. Please reference
// https://learn.microsoft.com/en-us/azure/azure-managed-lustre/ for a thorough explanation of these resolution modes.
type ConflictResolutionMode string

const (
	ConflictResolutionModeFail             ConflictResolutionMode = "Fail"
	ConflictResolutionModeOverwriteAlways  ConflictResolutionMode = "OverwriteAlways"
	ConflictResolutionModeOverwriteIfDirty ConflictResolutionMode = "OverwriteIfDirty"
	ConflictResolutionModeSkip             ConflictResolutionMode = "Skip"
)

// PossibleConflictResolutionModeValues returns the possible values for the ConflictResolutionMode const type.
func PossibleConflictResolutionModeValues() []ConflictResolutionMode {
	return []ConflictResolutionMode{
		ConflictResolutionModeFail,
		ConflictResolutionModeOverwriteAlways,
		ConflictResolutionModeOverwriteIfDirty,
		ConflictResolutionModeSkip,
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

// DomainJoinedType - True if the HPC Cache is joined to the Active Directory domain.
type DomainJoinedType string

const (
	DomainJoinedTypeError DomainJoinedType = "Error"
	DomainJoinedTypeNo    DomainJoinedType = "No"
	DomainJoinedTypeYes   DomainJoinedType = "Yes"
)

// PossibleDomainJoinedTypeValues returns the possible values for the DomainJoinedType const type.
func PossibleDomainJoinedTypeValues() []DomainJoinedType {
	return []DomainJoinedType{
		DomainJoinedTypeError,
		DomainJoinedTypeNo,
		DomainJoinedTypeYes,
	}
}

// FirmwareStatusType - True if there is a firmware update ready to install on this cache. The firmware will automatically
// be installed after firmwareUpdateDeadline if not triggered earlier via the upgrade operation.
type FirmwareStatusType string

const (
	FirmwareStatusTypeAvailable   FirmwareStatusType = "available"
	FirmwareStatusTypeUnavailable FirmwareStatusType = "unavailable"
)

// PossibleFirmwareStatusTypeValues returns the possible values for the FirmwareStatusType const type.
func PossibleFirmwareStatusTypeValues() []FirmwareStatusType {
	return []FirmwareStatusType{
		FirmwareStatusTypeAvailable,
		FirmwareStatusTypeUnavailable,
	}
}

// HealthStateType - List of cache health states. Down is when the cluster is not responding. Degraded is when its functioning
// but has some alerts. Transitioning when it is creating or deleting. Unknown will be returned
// in old api versions when a new value is added in future versions. WaitingForKey is when the create is waiting for the system
// assigned identity to be given access to the encryption key in the
// encryption settings.
type HealthStateType string

const (
	HealthStateTypeDegraded      HealthStateType = "Degraded"
	HealthStateTypeDown          HealthStateType = "Down"
	HealthStateTypeFlushing      HealthStateType = "Flushing"
	HealthStateTypeHealthy       HealthStateType = "Healthy"
	HealthStateTypeStartFailed   HealthStateType = "StartFailed"
	HealthStateTypeStopped       HealthStateType = "Stopped"
	HealthStateTypeStopping      HealthStateType = "Stopping"
	HealthStateTypeTransitioning HealthStateType = "Transitioning"
	HealthStateTypeUnknown       HealthStateType = "Unknown"
	HealthStateTypeUpgradeFailed HealthStateType = "UpgradeFailed"
	HealthStateTypeUpgrading     HealthStateType = "Upgrading"
	HealthStateTypeWaitingForKey HealthStateType = "WaitingForKey"
)

// PossibleHealthStateTypeValues returns the possible values for the HealthStateType const type.
func PossibleHealthStateTypeValues() []HealthStateType {
	return []HealthStateType{
		HealthStateTypeDegraded,
		HealthStateTypeDown,
		HealthStateTypeFlushing,
		HealthStateTypeHealthy,
		HealthStateTypeStartFailed,
		HealthStateTypeStopped,
		HealthStateTypeStopping,
		HealthStateTypeTransitioning,
		HealthStateTypeUnknown,
		HealthStateTypeUpgradeFailed,
		HealthStateTypeUpgrading,
		HealthStateTypeWaitingForKey,
	}
}

// ImportJobAdminStatus - The administrative status of the import job. Possible values: 'Enable', 'Disable'. Passing in a
// value of 'Disable' will cancel the current active import job. By default it is set to 'Enable'.
type ImportJobAdminStatus string

const (
	ImportJobAdminStatusDisable ImportJobAdminStatus = "Disable"
	ImportJobAdminStatusEnable  ImportJobAdminStatus = "Enable"
)

// PossibleImportJobAdminStatusValues returns the possible values for the ImportJobAdminStatus const type.
func PossibleImportJobAdminStatusValues() []ImportJobAdminStatus {
	return []ImportJobAdminStatus{
		ImportJobAdminStatusDisable,
		ImportJobAdminStatusEnable,
	}
}

// ImportJobProvisioningStateType - ARM provisioning state.
type ImportJobProvisioningStateType string

const (
	ImportJobProvisioningStateTypeCanceled  ImportJobProvisioningStateType = "Canceled"
	ImportJobProvisioningStateTypeCreating  ImportJobProvisioningStateType = "Creating"
	ImportJobProvisioningStateTypeDeleting  ImportJobProvisioningStateType = "Deleting"
	ImportJobProvisioningStateTypeFailed    ImportJobProvisioningStateType = "Failed"
	ImportJobProvisioningStateTypeSucceeded ImportJobProvisioningStateType = "Succeeded"
	ImportJobProvisioningStateTypeUpdating  ImportJobProvisioningStateType = "Updating"
)

// PossibleImportJobProvisioningStateTypeValues returns the possible values for the ImportJobProvisioningStateType const type.
func PossibleImportJobProvisioningStateTypeValues() []ImportJobProvisioningStateType {
	return []ImportJobProvisioningStateType{
		ImportJobProvisioningStateTypeCanceled,
		ImportJobProvisioningStateTypeCreating,
		ImportJobProvisioningStateTypeDeleting,
		ImportJobProvisioningStateTypeFailed,
		ImportJobProvisioningStateTypeSucceeded,
		ImportJobProvisioningStateTypeUpdating,
	}
}

// ImportStatusType - The operational state of the import job. InProgress indicates the import is still running. Canceled
// indicates it has been canceled by the user. Completed indicates import finished, successfully
// importing all discovered blobs into the Lustre namespace. CompletedPartial indicates the import finished but some blobs
// either were found to be conflicting and could not be imported or other errors
// were encountered. Failed means the import was unable to complete due to a fatal error.
type ImportStatusType string

const (
	ImportStatusTypeCanceled         ImportStatusType = "Canceled"
	ImportStatusTypeCancelling       ImportStatusType = "Cancelling"
	ImportStatusTypeCompleted        ImportStatusType = "Completed"
	ImportStatusTypeCompletedPartial ImportStatusType = "CompletedPartial"
	ImportStatusTypeFailed           ImportStatusType = "Failed"
	ImportStatusTypeInProgress       ImportStatusType = "InProgress"
)

// PossibleImportStatusTypeValues returns the possible values for the ImportStatusType const type.
func PossibleImportStatusTypeValues() []ImportStatusType {
	return []ImportStatusType{
		ImportStatusTypeCanceled,
		ImportStatusTypeCancelling,
		ImportStatusTypeCompleted,
		ImportStatusTypeCompletedPartial,
		ImportStatusTypeFailed,
		ImportStatusTypeInProgress,
	}
}

// MaintenanceDayOfWeekType - Day of the week on which the maintenance window will occur.
type MaintenanceDayOfWeekType string

const (
	MaintenanceDayOfWeekTypeFriday    MaintenanceDayOfWeekType = "Friday"
	MaintenanceDayOfWeekTypeMonday    MaintenanceDayOfWeekType = "Monday"
	MaintenanceDayOfWeekTypeSaturday  MaintenanceDayOfWeekType = "Saturday"
	MaintenanceDayOfWeekTypeSunday    MaintenanceDayOfWeekType = "Sunday"
	MaintenanceDayOfWeekTypeThursday  MaintenanceDayOfWeekType = "Thursday"
	MaintenanceDayOfWeekTypeTuesday   MaintenanceDayOfWeekType = "Tuesday"
	MaintenanceDayOfWeekTypeWednesday MaintenanceDayOfWeekType = "Wednesday"
)

// PossibleMaintenanceDayOfWeekTypeValues returns the possible values for the MaintenanceDayOfWeekType const type.
func PossibleMaintenanceDayOfWeekTypeValues() []MaintenanceDayOfWeekType {
	return []MaintenanceDayOfWeekType{
		MaintenanceDayOfWeekTypeFriday,
		MaintenanceDayOfWeekTypeMonday,
		MaintenanceDayOfWeekTypeSaturday,
		MaintenanceDayOfWeekTypeSunday,
		MaintenanceDayOfWeekTypeThursday,
		MaintenanceDayOfWeekTypeTuesday,
		MaintenanceDayOfWeekTypeWednesday,
	}
}

type MetricAggregationType string

const (
	MetricAggregationTypeAverage      MetricAggregationType = "Average"
	MetricAggregationTypeCount        MetricAggregationType = "Count"
	MetricAggregationTypeMaximum      MetricAggregationType = "Maximum"
	MetricAggregationTypeMinimum      MetricAggregationType = "Minimum"
	MetricAggregationTypeNone         MetricAggregationType = "None"
	MetricAggregationTypeNotSpecified MetricAggregationType = "NotSpecified"
	MetricAggregationTypeTotal        MetricAggregationType = "Total"
)

// PossibleMetricAggregationTypeValues returns the possible values for the MetricAggregationType const type.
func PossibleMetricAggregationTypeValues() []MetricAggregationType {
	return []MetricAggregationType{
		MetricAggregationTypeAverage,
		MetricAggregationTypeCount,
		MetricAggregationTypeMaximum,
		MetricAggregationTypeMinimum,
		MetricAggregationTypeNone,
		MetricAggregationTypeNotSpecified,
		MetricAggregationTypeTotal,
	}
}

// NfsAccessRuleAccess - Access allowed by this rule.
type NfsAccessRuleAccess string

const (
	NfsAccessRuleAccessNo NfsAccessRuleAccess = "no"
	NfsAccessRuleAccessRo NfsAccessRuleAccess = "ro"
	NfsAccessRuleAccessRw NfsAccessRuleAccess = "rw"
)

// PossibleNfsAccessRuleAccessValues returns the possible values for the NfsAccessRuleAccess const type.
func PossibleNfsAccessRuleAccessValues() []NfsAccessRuleAccess {
	return []NfsAccessRuleAccess{
		NfsAccessRuleAccessNo,
		NfsAccessRuleAccessRo,
		NfsAccessRuleAccessRw,
	}
}

// NfsAccessRuleScope - Scope for this rule. The scope and filter determine which clients match the rule.
type NfsAccessRuleScope string

const (
	NfsAccessRuleScopeDefault NfsAccessRuleScope = "default"
	NfsAccessRuleScopeHost    NfsAccessRuleScope = "host"
	NfsAccessRuleScopeNetwork NfsAccessRuleScope = "network"
)

// PossibleNfsAccessRuleScopeValues returns the possible values for the NfsAccessRuleScope const type.
func PossibleNfsAccessRuleScopeValues() []NfsAccessRuleScope {
	return []NfsAccessRuleScope{
		NfsAccessRuleScopeDefault,
		NfsAccessRuleScopeHost,
		NfsAccessRuleScopeNetwork,
	}
}

// OperationalStateType - Storage target operational state.
type OperationalStateType string

const (
	OperationalStateTypeBusy      OperationalStateType = "Busy"
	OperationalStateTypeFlushing  OperationalStateType = "Flushing"
	OperationalStateTypeReady     OperationalStateType = "Ready"
	OperationalStateTypeSuspended OperationalStateType = "Suspended"
)

// PossibleOperationalStateTypeValues returns the possible values for the OperationalStateType const type.
func PossibleOperationalStateTypeValues() []OperationalStateType {
	return []OperationalStateType{
		OperationalStateTypeBusy,
		OperationalStateTypeFlushing,
		OperationalStateTypeReady,
		OperationalStateTypeSuspended,
	}
}

// PrimingJobState - The state of the priming operation.
type PrimingJobState string

const (
	PrimingJobStateComplete PrimingJobState = "Complete"
	PrimingJobStatePaused   PrimingJobState = "Paused"
	PrimingJobStateQueued   PrimingJobState = "Queued"
	PrimingJobStateRunning  PrimingJobState = "Running"
)

// PossiblePrimingJobStateValues returns the possible values for the PrimingJobState const type.
func PossiblePrimingJobStateValues() []PrimingJobState {
	return []PrimingJobState{
		PrimingJobStateComplete,
		PrimingJobStatePaused,
		PrimingJobStateQueued,
		PrimingJobStateRunning,
	}
}

// ProvisioningStateType - ARM provisioning state, see https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/Addendum.md#provisioningstate-property
type ProvisioningStateType string

const (
	ProvisioningStateTypeCancelled ProvisioningStateType = "Canceled"
	ProvisioningStateTypeCreating  ProvisioningStateType = "Creating"
	ProvisioningStateTypeDeleting  ProvisioningStateType = "Deleting"
	ProvisioningStateTypeFailed    ProvisioningStateType = "Failed"
	ProvisioningStateTypeSucceeded ProvisioningStateType = "Succeeded"
	ProvisioningStateTypeUpdating  ProvisioningStateType = "Updating"
)

// PossibleProvisioningStateTypeValues returns the possible values for the ProvisioningStateType const type.
func PossibleProvisioningStateTypeValues() []ProvisioningStateType {
	return []ProvisioningStateType{
		ProvisioningStateTypeCancelled,
		ProvisioningStateTypeCreating,
		ProvisioningStateTypeDeleting,
		ProvisioningStateTypeFailed,
		ProvisioningStateTypeSucceeded,
		ProvisioningStateTypeUpdating,
	}
}

// ReasonCode - The reason for the restriction. As of now this can be "QuotaId" or "NotAvailableForSubscription". "QuotaId"
// is set when the SKU has requiredQuotas parameter as the subscription does not belong to that
// quota. "NotAvailableForSubscription" is related to capacity at the datacenter.
type ReasonCode string

const (
	ReasonCodeNotAvailableForSubscription ReasonCode = "NotAvailableForSubscription"
	ReasonCodeQuotaID                     ReasonCode = "QuotaId"
)

// PossibleReasonCodeValues returns the possible values for the ReasonCode const type.
func PossibleReasonCodeValues() []ReasonCode {
	return []ReasonCode{
		ReasonCodeNotAvailableForSubscription,
		ReasonCodeQuotaID,
	}
}

// StorageTargetType - Type of the Storage Target.
type StorageTargetType string

const (
	StorageTargetTypeBlobNfs StorageTargetType = "blobNfs"
	StorageTargetTypeClfs    StorageTargetType = "clfs"
	StorageTargetTypeNfs3    StorageTargetType = "nfs3"
	StorageTargetTypeUnknown StorageTargetType = "unknown"
)

// PossibleStorageTargetTypeValues returns the possible values for the StorageTargetType const type.
func PossibleStorageTargetTypeValues() []StorageTargetType {
	return []StorageTargetType{
		StorageTargetTypeBlobNfs,
		StorageTargetTypeClfs,
		StorageTargetTypeNfs3,
		StorageTargetTypeUnknown,
	}
}

// UsernameDownloadedType - Indicates whether or not the HPC Cache has performed the username download successfully.
type UsernameDownloadedType string

const (
	UsernameDownloadedTypeError UsernameDownloadedType = "Error"
	UsernameDownloadedTypeNo    UsernameDownloadedType = "No"
	UsernameDownloadedTypeYes   UsernameDownloadedType = "Yes"
)

// PossibleUsernameDownloadedTypeValues returns the possible values for the UsernameDownloadedType const type.
func PossibleUsernameDownloadedTypeValues() []UsernameDownloadedType {
	return []UsernameDownloadedType{
		UsernameDownloadedTypeError,
		UsernameDownloadedTypeNo,
		UsernameDownloadedTypeYes,
	}
}

// UsernameSource - This setting determines how the cache gets username and group names for clients.
type UsernameSource string

const (
	UsernameSourceAD   UsernameSource = "AD"
	UsernameSourceFile UsernameSource = "File"
	UsernameSourceLDAP UsernameSource = "LDAP"
	UsernameSourceNone UsernameSource = "None"
)

// PossibleUsernameSourceValues returns the possible values for the UsernameSource const type.
func PossibleUsernameSourceValues() []UsernameSource {
	return []UsernameSource{
		UsernameSourceAD,
		UsernameSourceFile,
		UsernameSourceLDAP,
		UsernameSourceNone,
	}
}
