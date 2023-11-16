//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeventhub

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub"
	moduleVersion = "v1.2.0"
)

type AccessRights string

const (
	AccessRightsListen AccessRights = "Listen"
	AccessRightsManage AccessRights = "Manage"
	AccessRightsSend   AccessRights = "Send"
)

// PossibleAccessRightsValues returns the possible values for the AccessRights const type.
func PossibleAccessRightsValues() []AccessRights {
	return []AccessRights{
		AccessRightsListen,
		AccessRightsManage,
		AccessRightsSend,
	}
}

// ApplicationGroupPolicyType - Application Group Policy types
type ApplicationGroupPolicyType string

const (
	ApplicationGroupPolicyTypeThrottlingPolicy ApplicationGroupPolicyType = "ThrottlingPolicy"
)

// PossibleApplicationGroupPolicyTypeValues returns the possible values for the ApplicationGroupPolicyType const type.
func PossibleApplicationGroupPolicyTypeValues() []ApplicationGroupPolicyType {
	return []ApplicationGroupPolicyType{
		ApplicationGroupPolicyTypeThrottlingPolicy,
	}
}

// CaptureIdentityType - Type of Azure Active Directory Managed Identity.
type CaptureIdentityType string

const (
	CaptureIdentityTypeSystemAssigned CaptureIdentityType = "SystemAssigned"
	CaptureIdentityTypeUserAssigned   CaptureIdentityType = "UserAssigned"
)

// PossibleCaptureIdentityTypeValues returns the possible values for the CaptureIdentityType const type.
func PossibleCaptureIdentityTypeValues() []CaptureIdentityType {
	return []CaptureIdentityType{
		CaptureIdentityTypeSystemAssigned,
		CaptureIdentityTypeUserAssigned,
	}
}

// CleanupPolicyRetentionDescription - Enumerates the possible values for cleanup policy
type CleanupPolicyRetentionDescription string

const (
	CleanupPolicyRetentionDescriptionCompact CleanupPolicyRetentionDescription = "Compact"
	CleanupPolicyRetentionDescriptionDelete  CleanupPolicyRetentionDescription = "Delete"
)

// PossibleCleanupPolicyRetentionDescriptionValues returns the possible values for the CleanupPolicyRetentionDescription const type.
func PossibleCleanupPolicyRetentionDescriptionValues() []CleanupPolicyRetentionDescription {
	return []CleanupPolicyRetentionDescription{
		CleanupPolicyRetentionDescriptionCompact,
		CleanupPolicyRetentionDescriptionDelete,
	}
}

// ClusterSKUName - Name of this SKU.
type ClusterSKUName string

const (
	ClusterSKUNameDedicated ClusterSKUName = "Dedicated"
)

// PossibleClusterSKUNameValues returns the possible values for the ClusterSKUName const type.
func PossibleClusterSKUNameValues() []ClusterSKUName {
	return []ClusterSKUName{
		ClusterSKUNameDedicated,
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

// DefaultAction - Default Action for Network Rule Set
type DefaultAction string

const (
	DefaultActionAllow DefaultAction = "Allow"
	DefaultActionDeny  DefaultAction = "Deny"
)

// PossibleDefaultActionValues returns the possible values for the DefaultAction const type.
func PossibleDefaultActionValues() []DefaultAction {
	return []DefaultAction{
		DefaultActionAllow,
		DefaultActionDeny,
	}
}

// EncodingCaptureDescription - Enumerates the possible values for the encoding format of capture description. Note: 'AvroDeflate'
// will be deprecated in New API Version
type EncodingCaptureDescription string

const (
	EncodingCaptureDescriptionAvro        EncodingCaptureDescription = "Avro"
	EncodingCaptureDescriptionAvroDeflate EncodingCaptureDescription = "AvroDeflate"
)

// PossibleEncodingCaptureDescriptionValues returns the possible values for the EncodingCaptureDescription const type.
func PossibleEncodingCaptureDescriptionValues() []EncodingCaptureDescription {
	return []EncodingCaptureDescription{
		EncodingCaptureDescriptionAvro,
		EncodingCaptureDescriptionAvroDeflate,
	}
}

// EndPointProvisioningState - Provisioning state of the Private Endpoint Connection.
type EndPointProvisioningState string

const (
	EndPointProvisioningStateCanceled  EndPointProvisioningState = "Canceled"
	EndPointProvisioningStateCreating  EndPointProvisioningState = "Creating"
	EndPointProvisioningStateDeleting  EndPointProvisioningState = "Deleting"
	EndPointProvisioningStateFailed    EndPointProvisioningState = "Failed"
	EndPointProvisioningStateSucceeded EndPointProvisioningState = "Succeeded"
	EndPointProvisioningStateUpdating  EndPointProvisioningState = "Updating"
)

// PossibleEndPointProvisioningStateValues returns the possible values for the EndPointProvisioningState const type.
func PossibleEndPointProvisioningStateValues() []EndPointProvisioningState {
	return []EndPointProvisioningState{
		EndPointProvisioningStateCanceled,
		EndPointProvisioningStateCreating,
		EndPointProvisioningStateDeleting,
		EndPointProvisioningStateFailed,
		EndPointProvisioningStateSucceeded,
		EndPointProvisioningStateUpdating,
	}
}

// EntityStatus - Enumerates the possible values for the status of the Event Hub.
type EntityStatus string

const (
	EntityStatusActive          EntityStatus = "Active"
	EntityStatusCreating        EntityStatus = "Creating"
	EntityStatusDeleting        EntityStatus = "Deleting"
	EntityStatusDisabled        EntityStatus = "Disabled"
	EntityStatusReceiveDisabled EntityStatus = "ReceiveDisabled"
	EntityStatusRenaming        EntityStatus = "Renaming"
	EntityStatusRestoring       EntityStatus = "Restoring"
	EntityStatusSendDisabled    EntityStatus = "SendDisabled"
	EntityStatusUnknown         EntityStatus = "Unknown"
)

// PossibleEntityStatusValues returns the possible values for the EntityStatus const type.
func PossibleEntityStatusValues() []EntityStatus {
	return []EntityStatus{
		EntityStatusActive,
		EntityStatusCreating,
		EntityStatusDeleting,
		EntityStatusDisabled,
		EntityStatusReceiveDisabled,
		EntityStatusRenaming,
		EntityStatusRestoring,
		EntityStatusSendDisabled,
		EntityStatusUnknown,
	}
}

// KeyType - The access key to regenerate.
type KeyType string

const (
	KeyTypePrimaryKey   KeyType = "PrimaryKey"
	KeyTypeSecondaryKey KeyType = "SecondaryKey"
)

// PossibleKeyTypeValues returns the possible values for the KeyType const type.
func PossibleKeyTypeValues() []KeyType {
	return []KeyType{
		KeyTypePrimaryKey,
		KeyTypeSecondaryKey,
	}
}

// ManagedServiceIdentityType - Type of managed service identity.
type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeNone                       ManagedServiceIdentityType = "None"
	ManagedServiceIdentityTypeSystemAssigned             ManagedServiceIdentityType = "SystemAssigned"
	ManagedServiceIdentityTypeSystemAssignedUserAssigned ManagedServiceIdentityType = "SystemAssigned, UserAssigned"
	ManagedServiceIdentityTypeUserAssigned               ManagedServiceIdentityType = "UserAssigned"
)

// PossibleManagedServiceIdentityTypeValues returns the possible values for the ManagedServiceIdentityType const type.
func PossibleManagedServiceIdentityTypeValues() []ManagedServiceIdentityType {
	return []ManagedServiceIdentityType{
		ManagedServiceIdentityTypeNone,
		ManagedServiceIdentityTypeSystemAssigned,
		ManagedServiceIdentityTypeSystemAssignedUserAssigned,
		ManagedServiceIdentityTypeUserAssigned,
	}
}

// MetricID - Metric Id on which the throttle limit should be set, MetricId can be discovered by hovering over Metric in the
// Metrics section of Event Hub Namespace inside Azure Portal
type MetricID string

const (
	MetricIDIncomingBytes    MetricID = "IncomingBytes"
	MetricIDIncomingMessages MetricID = "IncomingMessages"
	MetricIDOutgoingBytes    MetricID = "OutgoingBytes"
	MetricIDOutgoingMessages MetricID = "OutgoingMessages"
)

// PossibleMetricIDValues returns the possible values for the MetricID const type.
func PossibleMetricIDValues() []MetricID {
	return []MetricID{
		MetricIDIncomingBytes,
		MetricIDIncomingMessages,
		MetricIDOutgoingBytes,
		MetricIDOutgoingMessages,
	}
}

// NetworkRuleIPAction - The IP Filter Action
type NetworkRuleIPAction string

const (
	NetworkRuleIPActionAllow NetworkRuleIPAction = "Allow"
)

// PossibleNetworkRuleIPActionValues returns the possible values for the NetworkRuleIPAction const type.
func PossibleNetworkRuleIPActionValues() []NetworkRuleIPAction {
	return []NetworkRuleIPAction{
		NetworkRuleIPActionAllow,
	}
}

// NetworkSecurityPerimeterConfigurationProvisioningState - Provisioning state of NetworkSecurityPerimeter configuration propagation
type NetworkSecurityPerimeterConfigurationProvisioningState string

const (
	NetworkSecurityPerimeterConfigurationProvisioningStateAccepted            NetworkSecurityPerimeterConfigurationProvisioningState = "Accepted"
	NetworkSecurityPerimeterConfigurationProvisioningStateCanceled            NetworkSecurityPerimeterConfigurationProvisioningState = "Canceled"
	NetworkSecurityPerimeterConfigurationProvisioningStateCreating            NetworkSecurityPerimeterConfigurationProvisioningState = "Creating"
	NetworkSecurityPerimeterConfigurationProvisioningStateDeleted             NetworkSecurityPerimeterConfigurationProvisioningState = "Deleted"
	NetworkSecurityPerimeterConfigurationProvisioningStateDeleting            NetworkSecurityPerimeterConfigurationProvisioningState = "Deleting"
	NetworkSecurityPerimeterConfigurationProvisioningStateFailed              NetworkSecurityPerimeterConfigurationProvisioningState = "Failed"
	NetworkSecurityPerimeterConfigurationProvisioningStateInvalidResponse     NetworkSecurityPerimeterConfigurationProvisioningState = "InvalidResponse"
	NetworkSecurityPerimeterConfigurationProvisioningStateSucceeded           NetworkSecurityPerimeterConfigurationProvisioningState = "Succeeded"
	NetworkSecurityPerimeterConfigurationProvisioningStateSucceededWithIssues NetworkSecurityPerimeterConfigurationProvisioningState = "SucceededWithIssues"
	NetworkSecurityPerimeterConfigurationProvisioningStateUnknown             NetworkSecurityPerimeterConfigurationProvisioningState = "Unknown"
	NetworkSecurityPerimeterConfigurationProvisioningStateUpdating            NetworkSecurityPerimeterConfigurationProvisioningState = "Updating"
)

// PossibleNetworkSecurityPerimeterConfigurationProvisioningStateValues returns the possible values for the NetworkSecurityPerimeterConfigurationProvisioningState const type.
func PossibleNetworkSecurityPerimeterConfigurationProvisioningStateValues() []NetworkSecurityPerimeterConfigurationProvisioningState {
	return []NetworkSecurityPerimeterConfigurationProvisioningState{
		NetworkSecurityPerimeterConfigurationProvisioningStateAccepted,
		NetworkSecurityPerimeterConfigurationProvisioningStateCanceled,
		NetworkSecurityPerimeterConfigurationProvisioningStateCreating,
		NetworkSecurityPerimeterConfigurationProvisioningStateDeleted,
		NetworkSecurityPerimeterConfigurationProvisioningStateDeleting,
		NetworkSecurityPerimeterConfigurationProvisioningStateFailed,
		NetworkSecurityPerimeterConfigurationProvisioningStateInvalidResponse,
		NetworkSecurityPerimeterConfigurationProvisioningStateSucceeded,
		NetworkSecurityPerimeterConfigurationProvisioningStateSucceededWithIssues,
		NetworkSecurityPerimeterConfigurationProvisioningStateUnknown,
		NetworkSecurityPerimeterConfigurationProvisioningStateUpdating,
	}
}

// NspAccessRuleDirection - Direction of Access Rule
type NspAccessRuleDirection string

const (
	NspAccessRuleDirectionInbound  NspAccessRuleDirection = "Inbound"
	NspAccessRuleDirectionOutbound NspAccessRuleDirection = "Outbound"
)

// PossibleNspAccessRuleDirectionValues returns the possible values for the NspAccessRuleDirection const type.
func PossibleNspAccessRuleDirectionValues() []NspAccessRuleDirection {
	return []NspAccessRuleDirection{
		NspAccessRuleDirectionInbound,
		NspAccessRuleDirectionOutbound,
	}
}

// PrivateLinkConnectionStatus - Status of the connection.
type PrivateLinkConnectionStatus string

const (
	PrivateLinkConnectionStatusApproved     PrivateLinkConnectionStatus = "Approved"
	PrivateLinkConnectionStatusDisconnected PrivateLinkConnectionStatus = "Disconnected"
	PrivateLinkConnectionStatusPending      PrivateLinkConnectionStatus = "Pending"
	PrivateLinkConnectionStatusRejected     PrivateLinkConnectionStatus = "Rejected"
)

// PossiblePrivateLinkConnectionStatusValues returns the possible values for the PrivateLinkConnectionStatus const type.
func PossiblePrivateLinkConnectionStatusValues() []PrivateLinkConnectionStatus {
	return []PrivateLinkConnectionStatus{
		PrivateLinkConnectionStatusApproved,
		PrivateLinkConnectionStatusDisconnected,
		PrivateLinkConnectionStatusPending,
		PrivateLinkConnectionStatusRejected,
	}
}

// ProvisioningState - Provisioning state of the Cluster.
type ProvisioningState string

const (
	ProvisioningStateActive    ProvisioningState = "Active"
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateScaling   ProvisioningState = "Scaling"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUnknown   ProvisioningState = "Unknown"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateActive,
		ProvisioningStateCanceled,
		ProvisioningStateCreating,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateScaling,
		ProvisioningStateSucceeded,
		ProvisioningStateUnknown,
	}
}

// ProvisioningStateDR - Provisioning state of the Alias(Disaster Recovery configuration) - possible values 'Accepted' or
// 'Succeeded' or 'Failed'
type ProvisioningStateDR string

const (
	ProvisioningStateDRAccepted  ProvisioningStateDR = "Accepted"
	ProvisioningStateDRFailed    ProvisioningStateDR = "Failed"
	ProvisioningStateDRSucceeded ProvisioningStateDR = "Succeeded"
)

// PossibleProvisioningStateDRValues returns the possible values for the ProvisioningStateDR const type.
func PossibleProvisioningStateDRValues() []ProvisioningStateDR {
	return []ProvisioningStateDR{
		ProvisioningStateDRAccepted,
		ProvisioningStateDRFailed,
		ProvisioningStateDRSucceeded,
	}
}

// PublicNetworkAccess - This determines if traffic is allowed over public network. By default it is enabled.
type PublicNetworkAccess string

const (
	PublicNetworkAccessDisabled           PublicNetworkAccess = "Disabled"
	PublicNetworkAccessEnabled            PublicNetworkAccess = "Enabled"
	PublicNetworkAccessSecuredByPerimeter PublicNetworkAccess = "SecuredByPerimeter"
)

// PossiblePublicNetworkAccessValues returns the possible values for the PublicNetworkAccess const type.
func PossiblePublicNetworkAccessValues() []PublicNetworkAccess {
	return []PublicNetworkAccess{
		PublicNetworkAccessDisabled,
		PublicNetworkAccessEnabled,
		PublicNetworkAccessSecuredByPerimeter,
	}
}

// PublicNetworkAccessFlag - This determines if traffic is allowed over public network. By default it is enabled. If value
// is SecuredByPerimeter then Inbound and Outbound communication is controlled by the network security
// perimeter and profile's access rules.
type PublicNetworkAccessFlag string

const (
	PublicNetworkAccessFlagDisabled           PublicNetworkAccessFlag = "Disabled"
	PublicNetworkAccessFlagEnabled            PublicNetworkAccessFlag = "Enabled"
	PublicNetworkAccessFlagSecuredByPerimeter PublicNetworkAccessFlag = "SecuredByPerimeter"
)

// PossiblePublicNetworkAccessFlagValues returns the possible values for the PublicNetworkAccessFlag const type.
func PossiblePublicNetworkAccessFlagValues() []PublicNetworkAccessFlag {
	return []PublicNetworkAccessFlag{
		PublicNetworkAccessFlagDisabled,
		PublicNetworkAccessFlagEnabled,
		PublicNetworkAccessFlagSecuredByPerimeter,
	}
}

// ResourceAssociationAccessMode - Access Mode of the resource association
type ResourceAssociationAccessMode string

const (
	ResourceAssociationAccessModeAuditMode         ResourceAssociationAccessMode = "AuditMode"
	ResourceAssociationAccessModeEnforcedMode      ResourceAssociationAccessMode = "EnforcedMode"
	ResourceAssociationAccessModeLearningMode      ResourceAssociationAccessMode = "LearningMode"
	ResourceAssociationAccessModeNoAssociationMode ResourceAssociationAccessMode = "NoAssociationMode"
	ResourceAssociationAccessModeUnspecifiedMode   ResourceAssociationAccessMode = "UnspecifiedMode"
)

// PossibleResourceAssociationAccessModeValues returns the possible values for the ResourceAssociationAccessMode const type.
func PossibleResourceAssociationAccessModeValues() []ResourceAssociationAccessMode {
	return []ResourceAssociationAccessMode{
		ResourceAssociationAccessModeAuditMode,
		ResourceAssociationAccessModeEnforcedMode,
		ResourceAssociationAccessModeLearningMode,
		ResourceAssociationAccessModeNoAssociationMode,
		ResourceAssociationAccessModeUnspecifiedMode,
	}
}

// RoleDisasterRecovery - role of namespace in GEO DR - possible values 'Primary' or 'PrimaryNotReplicating' or 'Secondary'
type RoleDisasterRecovery string

const (
	RoleDisasterRecoveryPrimary               RoleDisasterRecovery = "Primary"
	RoleDisasterRecoveryPrimaryNotReplicating RoleDisasterRecovery = "PrimaryNotReplicating"
	RoleDisasterRecoverySecondary             RoleDisasterRecovery = "Secondary"
)

// PossibleRoleDisasterRecoveryValues returns the possible values for the RoleDisasterRecovery const type.
func PossibleRoleDisasterRecoveryValues() []RoleDisasterRecovery {
	return []RoleDisasterRecovery{
		RoleDisasterRecoveryPrimary,
		RoleDisasterRecoveryPrimaryNotReplicating,
		RoleDisasterRecoverySecondary,
	}
}

// SKUName - Name of this SKU.
type SKUName string

const (
	SKUNameBasic    SKUName = "Basic"
	SKUNamePremium  SKUName = "Premium"
	SKUNameStandard SKUName = "Standard"
)

// PossibleSKUNameValues returns the possible values for the SKUName const type.
func PossibleSKUNameValues() []SKUName {
	return []SKUName{
		SKUNameBasic,
		SKUNamePremium,
		SKUNameStandard,
	}
}

// SKUTier - The billing tier of this particular SKU.
type SKUTier string

const (
	SKUTierBasic    SKUTier = "Basic"
	SKUTierPremium  SKUTier = "Premium"
	SKUTierStandard SKUTier = "Standard"
)

// PossibleSKUTierValues returns the possible values for the SKUTier const type.
func PossibleSKUTierValues() []SKUTier {
	return []SKUTier{
		SKUTierBasic,
		SKUTierPremium,
		SKUTierStandard,
	}
}

type SchemaCompatibility string

const (
	SchemaCompatibilityBackward SchemaCompatibility = "Backward"
	SchemaCompatibilityForward  SchemaCompatibility = "Forward"
	SchemaCompatibilityNone     SchemaCompatibility = "None"
)

// PossibleSchemaCompatibilityValues returns the possible values for the SchemaCompatibility const type.
func PossibleSchemaCompatibilityValues() []SchemaCompatibility {
	return []SchemaCompatibility{
		SchemaCompatibilityBackward,
		SchemaCompatibilityForward,
		SchemaCompatibilityNone,
	}
}

type SchemaType string

const (
	SchemaTypeAvro    SchemaType = "Avro"
	SchemaTypeUnknown SchemaType = "Unknown"
)

// PossibleSchemaTypeValues returns the possible values for the SchemaType const type.
func PossibleSchemaTypeValues() []SchemaType {
	return []SchemaType{
		SchemaTypeAvro,
		SchemaTypeUnknown,
	}
}

// StartDayOfWeek - Preferred day of the week in UTC time to begin an upgrade. If 'Any' is selected, upgrade will proceed
// at any given weekday
type StartDayOfWeek string

const (
	StartDayOfWeekAny       StartDayOfWeek = "Any"
	StartDayOfWeekFriday    StartDayOfWeek = "Friday"
	StartDayOfWeekMonday    StartDayOfWeek = "Monday"
	StartDayOfWeekSaturday  StartDayOfWeek = "Saturday"
	StartDayOfWeekSunday    StartDayOfWeek = "Sunday"
	StartDayOfWeekThursday  StartDayOfWeek = "Thursday"
	StartDayOfWeekTuesday   StartDayOfWeek = "Tuesday"
	StartDayOfWeekWednesday StartDayOfWeek = "Wednesday"
)

// PossibleStartDayOfWeekValues returns the possible values for the StartDayOfWeek const type.
func PossibleStartDayOfWeekValues() []StartDayOfWeek {
	return []StartDayOfWeek{
		StartDayOfWeekAny,
		StartDayOfWeekFriday,
		StartDayOfWeekMonday,
		StartDayOfWeekSaturday,
		StartDayOfWeekSunday,
		StartDayOfWeekThursday,
		StartDayOfWeekTuesday,
		StartDayOfWeekWednesday,
	}
}

// TLSVersion - The minimum TLS version for the cluster to support, e.g. '1.2'
type TLSVersion string

const (
	TLSVersionOne0 TLSVersion = "1.0"
	TLSVersionOne1 TLSVersion = "1.1"
	TLSVersionOne2 TLSVersion = "1.2"
)

// PossibleTLSVersionValues returns the possible values for the TLSVersion const type.
func PossibleTLSVersionValues() []TLSVersion {
	return []TLSVersion{
		TLSVersionOne0,
		TLSVersionOne1,
		TLSVersionOne2,
	}
}

// UnavailableReason - Specifies the reason for the unavailability of the service.
type UnavailableReason string

const (
	UnavailableReasonInvalidName                           UnavailableReason = "InvalidName"
	UnavailableReasonNameInLockdown                        UnavailableReason = "NameInLockdown"
	UnavailableReasonNameInUse                             UnavailableReason = "NameInUse"
	UnavailableReasonNone                                  UnavailableReason = "None"
	UnavailableReasonSubscriptionIsDisabled                UnavailableReason = "SubscriptionIsDisabled"
	UnavailableReasonTooManyNamespaceInCurrentSubscription UnavailableReason = "TooManyNamespaceInCurrentSubscription"
)

// PossibleUnavailableReasonValues returns the possible values for the UnavailableReason const type.
func PossibleUnavailableReasonValues() []UnavailableReason {
	return []UnavailableReason{
		UnavailableReasonInvalidName,
		UnavailableReasonNameInLockdown,
		UnavailableReasonNameInUse,
		UnavailableReasonNone,
		UnavailableReasonSubscriptionIsDisabled,
		UnavailableReasonTooManyNamespaceInCurrentSubscription,
	}
}
