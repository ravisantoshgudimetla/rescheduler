// +build go1.9

// Copyright 2019 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package eventhub

import original "github.com/Azure/azure-sdk-for-go/services/eventhub/mgmt/2017-04-01/eventhub"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient
type ConsumerGroupsClient = original.ConsumerGroupsClient
type DisasterRecoveryConfigsClient = original.DisasterRecoveryConfigsClient
type EventHubsClient = original.EventHubsClient
type AccessRights = original.AccessRights

const (
	Listen AccessRights = original.Listen
	Manage AccessRights = original.Manage
	Send   AccessRights = original.Send
)

type EncodingCaptureDescription = original.EncodingCaptureDescription

const (
	Avro        EncodingCaptureDescription = original.Avro
	AvroDeflate EncodingCaptureDescription = original.AvroDeflate
)

type EntityStatus = original.EntityStatus

const (
	Active          EntityStatus = original.Active
	Creating        EntityStatus = original.Creating
	Deleting        EntityStatus = original.Deleting
	Disabled        EntityStatus = original.Disabled
	ReceiveDisabled EntityStatus = original.ReceiveDisabled
	Renaming        EntityStatus = original.Renaming
	Restoring       EntityStatus = original.Restoring
	SendDisabled    EntityStatus = original.SendDisabled
	Unknown         EntityStatus = original.Unknown
)

type KeyType = original.KeyType

const (
	PrimaryKey   KeyType = original.PrimaryKey
	SecondaryKey KeyType = original.SecondaryKey
)

type ProvisioningStateDR = original.ProvisioningStateDR

const (
	Accepted  ProvisioningStateDR = original.Accepted
	Failed    ProvisioningStateDR = original.Failed
	Succeeded ProvisioningStateDR = original.Succeeded
)

type RoleDisasterRecovery = original.RoleDisasterRecovery

const (
	Primary               RoleDisasterRecovery = original.Primary
	PrimaryNotReplicating RoleDisasterRecovery = original.PrimaryNotReplicating
	Secondary             RoleDisasterRecovery = original.Secondary
)

type SkuName = original.SkuName

const (
	Basic    SkuName = original.Basic
	Standard SkuName = original.Standard
)

type SkuTier = original.SkuTier

const (
	SkuTierBasic    SkuTier = original.SkuTierBasic
	SkuTierStandard SkuTier = original.SkuTierStandard
)

type UnavailableReason = original.UnavailableReason

const (
	InvalidName                           UnavailableReason = original.InvalidName
	NameInLockdown                        UnavailableReason = original.NameInLockdown
	NameInUse                             UnavailableReason = original.NameInUse
	None                                  UnavailableReason = original.None
	SubscriptionIsDisabled                UnavailableReason = original.SubscriptionIsDisabled
	TooManyNamespaceInCurrentSubscription UnavailableReason = original.TooManyNamespaceInCurrentSubscription
)

type AccessKeys = original.AccessKeys
type ArmDisasterRecovery = original.ArmDisasterRecovery
type ArmDisasterRecoveryListResult = original.ArmDisasterRecoveryListResult
type ArmDisasterRecoveryListResultIterator = original.ArmDisasterRecoveryListResultIterator
type ArmDisasterRecoveryListResultPage = original.ArmDisasterRecoveryListResultPage
type ArmDisasterRecoveryProperties = original.ArmDisasterRecoveryProperties
type AuthorizationRule = original.AuthorizationRule
type AuthorizationRuleListResult = original.AuthorizationRuleListResult
type AuthorizationRuleListResultIterator = original.AuthorizationRuleListResultIterator
type AuthorizationRuleListResultPage = original.AuthorizationRuleListResultPage
type AuthorizationRuleProperties = original.AuthorizationRuleProperties
type CaptureDescription = original.CaptureDescription
type CheckNameAvailabilityParameter = original.CheckNameAvailabilityParameter
type CheckNameAvailabilityResult = original.CheckNameAvailabilityResult
type ConsumerGroup = original.ConsumerGroup
type ConsumerGroupListResult = original.ConsumerGroupListResult
type ConsumerGroupListResultIterator = original.ConsumerGroupListResultIterator
type ConsumerGroupListResultPage = original.ConsumerGroupListResultPage
type ConsumerGroupProperties = original.ConsumerGroupProperties
type Destination = original.Destination
type DestinationProperties = original.DestinationProperties
type EHNamespace = original.EHNamespace
type EHNamespaceListResult = original.EHNamespaceListResult
type EHNamespaceListResultIterator = original.EHNamespaceListResultIterator
type EHNamespaceListResultPage = original.EHNamespaceListResultPage
type EHNamespaceProperties = original.EHNamespaceProperties
type ErrorResponse = original.ErrorResponse
type ListResult = original.ListResult
type ListResultIterator = original.ListResultIterator
type ListResultPage = original.ListResultPage
type MessagingPlan = original.MessagingPlan
type MessagingPlanProperties = original.MessagingPlanProperties
type MessagingRegions = original.MessagingRegions
type MessagingRegionsListResult = original.MessagingRegionsListResult
type MessagingRegionsListResultIterator = original.MessagingRegionsListResultIterator
type MessagingRegionsListResultPage = original.MessagingRegionsListResultPage
type MessagingRegionsProperties = original.MessagingRegionsProperties
type Model = original.Model
type NamespacesCreateOrUpdateFuture = original.NamespacesCreateOrUpdateFuture
type NamespacesDeleteFuture = original.NamespacesDeleteFuture
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type Properties = original.Properties
type RegenerateAccessKeyParameters = original.RegenerateAccessKeyParameters
type Resource = original.Resource
type Sku = original.Sku
type TrackedResource = original.TrackedResource
type NamespacesClient = original.NamespacesClient
type OperationsClient = original.OperationsClient
type RegionsClient = original.RegionsClient

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewConsumerGroupsClient(subscriptionID string) ConsumerGroupsClient {
	return original.NewConsumerGroupsClient(subscriptionID)
}
func NewConsumerGroupsClientWithBaseURI(baseURI string, subscriptionID string) ConsumerGroupsClient {
	return original.NewConsumerGroupsClientWithBaseURI(baseURI, subscriptionID)
}
func NewDisasterRecoveryConfigsClient(subscriptionID string) DisasterRecoveryConfigsClient {
	return original.NewDisasterRecoveryConfigsClient(subscriptionID)
}
func NewDisasterRecoveryConfigsClientWithBaseURI(baseURI string, subscriptionID string) DisasterRecoveryConfigsClient {
	return original.NewDisasterRecoveryConfigsClientWithBaseURI(baseURI, subscriptionID)
}
func NewEventHubsClient(subscriptionID string) EventHubsClient {
	return original.NewEventHubsClient(subscriptionID)
}
func NewEventHubsClientWithBaseURI(baseURI string, subscriptionID string) EventHubsClient {
	return original.NewEventHubsClientWithBaseURI(baseURI, subscriptionID)
}
func PossibleAccessRightsValues() []AccessRights {
	return original.PossibleAccessRightsValues()
}
func PossibleEncodingCaptureDescriptionValues() []EncodingCaptureDescription {
	return original.PossibleEncodingCaptureDescriptionValues()
}
func PossibleEntityStatusValues() []EntityStatus {
	return original.PossibleEntityStatusValues()
}
func PossibleKeyTypeValues() []KeyType {
	return original.PossibleKeyTypeValues()
}
func PossibleProvisioningStateDRValues() []ProvisioningStateDR {
	return original.PossibleProvisioningStateDRValues()
}
func PossibleRoleDisasterRecoveryValues() []RoleDisasterRecovery {
	return original.PossibleRoleDisasterRecoveryValues()
}
func PossibleSkuNameValues() []SkuName {
	return original.PossibleSkuNameValues()
}
func PossibleSkuTierValues() []SkuTier {
	return original.PossibleSkuTierValues()
}
func PossibleUnavailableReasonValues() []UnavailableReason {
	return original.PossibleUnavailableReasonValues()
}
func NewNamespacesClient(subscriptionID string) NamespacesClient {
	return original.NewNamespacesClient(subscriptionID)
}
func NewNamespacesClientWithBaseURI(baseURI string, subscriptionID string) NamespacesClient {
	return original.NewNamespacesClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewRegionsClient(subscriptionID string) RegionsClient {
	return original.NewRegionsClient(subscriptionID)
}
func NewRegionsClientWithBaseURI(baseURI string, subscriptionID string) RegionsClient {
	return original.NewRegionsClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
