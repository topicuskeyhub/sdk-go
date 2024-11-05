package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type NonLinkable struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The Type property
    typeEscaped *string
}
// NewNonLinkable instantiates a new NonLinkable and sets the default values.
func NewNonLinkable()(*NonLinkable) {
    m := &NonLinkable{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateNonLinkableFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNonLinkableFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("$type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "AuditInfo":
                        return NewAuditInfo(), nil
                    case "auth.AccountBulkUpdate":
                        return NewAuthAccountBulkUpdate(), nil
                    case "auth.AccountChangeLocale":
                        return NewAuthAccountChangeLocale(), nil
                    case "auth.AccountRecoveryStatus":
                        return NewAuthAccountRecoveryStatus(), nil
                    case "auth.AccountsAuditStats":
                        return NewAuthAccountsAuditStats(), nil
                    case "auth.AccountSettings":
                        return NewAuthAccountSettings(), nil
                    case "auth.AccountStatus":
                        return NewAuthAccountStatus(), nil
                    case "auth.InternalAccountStatusObject":
                        return NewAuthInternalAccountStatusObject(), nil
                    case "auth.MoveInternalAccount":
                        return NewAuthMoveInternalAccount(), nil
                    case "auth.StoredAccountAttribute":
                        return NewAuthStoredAccountAttribute(), nil
                    case "auth.StoredAccountAttributes":
                        return NewAuthStoredAccountAttributes(), nil
                    case "directory.AccountDirectoryStatusReport":
                        return NewDirectoryAccountDirectoryStatusReport(), nil
                    case "GeneratedSecret":
                        return NewGeneratedSecret(), nil
                    case "group.GroupAccessInfo":
                        return NewGroupGroupAccessInfo(), nil
                    case "group.GroupAccountNesting":
                        return NewGroupGroupAccountNesting(), nil
                    case "group.GroupAccountStatus":
                        return NewGroupGroupAccountStatus(), nil
                    case "group.GroupAdminMail":
                        return NewGroupGroupAdminMail(), nil
                    case "group.GroupAuditingInfo":
                        return NewGroupGroupAuditingInfo(), nil
                    case "group.GroupClassificationInfo":
                        return NewGroupGroupClassificationInfo(), nil
                    case "group.GroupClassificationUpdate":
                        return NewGroupGroupClassificationUpdate(), nil
                    case "group.GroupGlobalRoleInfo":
                        return NewGroupGroupGlobalRoleInfo(), nil
                    case "group.GroupInfo":
                        return NewGroupGroupInfo(), nil
                    case "group.GroupsAuditStats":
                        return NewGroupGroupsAuditStats(), nil
                    case "launchpad.DisplayedLaunchpadTiles":
                        return NewLaunchpadDisplayedLaunchpadTiles(), nil
                    case "license.KeyHubLicenseInfo":
                        return NewLicenseKeyHubLicenseInfo(), nil
                    case "mark.ItemMarker":
                        return NewMarkItemMarker(), nil
                    case "mark.ItemMarkers":
                        return NewMarkItemMarkers(), nil
                    case "notification.CertificateExpiredNotification":
                        return NewNotificationCertificateExpiredNotification(), nil
                    case "notification.GroupAuditRequiredNotification":
                        return NewNotificationGroupAuditRequiredNotification(), nil
                    case "notification.GroupEditRequiredNotification":
                        return NewNotificationGroupEditRequiredNotification(), nil
                    case "notification.InvalidSignaturesDetectedNotification":
                        return NewNotificationInvalidSignaturesDetectedNotification(), nil
                    case "notification.LicenseStatusNotification":
                        return NewNotificationLicenseStatusNotification(), nil
                    case "notification.ModificationRequestNotification":
                        return NewNotificationModificationRequestNotification(), nil
                    case "notification.Notification":
                        return NewNotificationNotification(), nil
                    case "notification.Notifications":
                        return NewNotificationNotifications(), nil
                    case "notification.OldApiVersionUsageNotification":
                        return NewNotificationOldApiVersionUsageNotification(), nil
                    case "notification.ProvisionConfigRequiredNotification":
                        return NewNotificationProvisionConfigRequiredNotification(), nil
                    case "notification.RotatingPasswordRequiredNotification":
                        return NewNotificationRotatingPasswordRequiredNotification(), nil
                    case "notification.UpdateAvailableNotification":
                        return NewNotificationUpdateAvailableNotification(), nil
                    case "notification.VaultRecordExpiredNotification":
                        return NewNotificationVaultRecordExpiredNotification(), nil
                    case "organization.OrganizationalUnitSettings":
                        return NewOrganizationOrganizationalUnitSettings(), nil
                    case "profile.AccessProfileAccountAttributeRuleStatus":
                        return NewProfileAccessProfileAccountAttributeRuleStatus(), nil
                    case "provisioning.AccountProvisioningStatus":
                        return NewProvisioningAccountProvisioningStatus(), nil
                    case "provisioning.AccountProvisioningStatusReport":
                        return NewProvisioningAccountProvisioningStatusReport(), nil
                    case "provisioning.CircuitBreakerStatistics":
                        return NewProvisioningCircuitBreakerStatistics(), nil
                    case "provisioning.GroupOnSystemTypes":
                        return NewProvisioningGroupOnSystemTypes(), nil
                    case "provisioning.GroupProvisioningStatus":
                        return NewProvisioningGroupProvisioningStatus(), nil
                    case "provisioning.ProvisioningManagementPermissions":
                        return NewProvisioningProvisioningManagementPermissions(), nil
                    case "provisioning.ProvisioningStatus":
                        return NewProvisioningProvisioningStatus(), nil
                    case "provisioning.TokenPassword":
                        return NewProvisioningTokenPassword(), nil
                    case "request.AcceptCreateGroupOnSystemRequestParameters":
                        return NewRequestAcceptCreateGroupOnSystemRequestParameters(), nil
                    case "request.AcceptCreateGroupRequestParameters":
                        return NewRequestAcceptCreateGroupRequestParameters(), nil
                    case "request.AcceptCreateProvisionedNamespaceRequestParameters":
                        return NewRequestAcceptCreateProvisionedNamespaceRequestParameters(), nil
                    case "request.AcceptCreateServiceAccountRequestParameters":
                        return NewRequestAcceptCreateServiceAccountRequestParameters(), nil
                    case "request.AcceptGrantAccessRequestParameters":
                        return NewRequestAcceptGrantAccessRequestParameters(), nil
                    case "request.AcceptJoinGroupRequestParameters":
                        return NewRequestAcceptJoinGroupRequestParameters(), nil
                    case "request.AcceptModificationRequestParameters":
                        return NewRequestAcceptModificationRequestParameters(), nil
                    case "request.ModificationRequestReport":
                        return NewRequestModificationRequestReport(), nil
                    case "request.ModificationRequestReportChangeDetails":
                        return NewRequestModificationRequestReportChangeDetails(), nil
                    case "request.ModificationRequestReportErrorDetails":
                        return NewRequestModificationRequestReportErrorDetails(), nil
                    case "request.ModificationRequestReportObjectChangeDetails":
                        return NewRequestModificationRequestReportObjectChangeDetails(), nil
                    case "request.PasswordResetRequestStatus":
                        return NewRequestPasswordResetRequestStatus(), nil
                    case "SegmentCount":
                        return NewSegmentCount(), nil
                    case "serviceaccount.ServiceAccountsAuditStats":
                        return NewServiceaccountServiceAccountsAuditStats(), nil
                    case "serviceaccount.ServiceAccountStatus":
                        return NewServiceaccountServiceAccountStatus(), nil
                    case "serviceaccount.ServiceAccountSupportedFeatures":
                        return NewServiceaccountServiceAccountSupportedFeatures(), nil
                    case "simple.StringValue":
                        return NewSimpleStringValue(), nil
                    case "simple.VersionInfo":
                        return NewSimpleVersionInfo(), nil
                    case "vault.DeletedVaultRecovery":
                        return NewVaultDeletedVaultRecovery(), nil
                    case "vault.MoveVaultRecord":
                        return NewVaultMoveVaultRecord(), nil
                    case "vault.PasswordMetadata":
                        return NewVaultPasswordMetadata(), nil
                    case "vault.VaultActivationStatus":
                        return NewVaultVaultActivationStatus(), nil
                    case "vault.VaultRecordSecrets":
                        return NewVaultVaultRecordSecrets(), nil
                    case "vault.VaultRecordShare":
                        return NewVaultVaultRecordShare(), nil
                    case "vault.VaultRecordShareSummary":
                        return NewVaultVaultRecordShareSummary(), nil
                    case "vault.VaultRecovery":
                        return NewVaultVaultRecovery(), nil
                    case "vault.VaultUnlock":
                        return NewVaultVaultUnlock(), nil
                    case "vault.VaultUnlockResponse":
                        return NewVaultVaultUnlockResponse(), nil
                    case "webhook.WebhookNameUuid":
                        return NewWebhookWebhookNameUuid(), nil
                    case "webhook.WebhookPush":
                        return NewWebhookWebhookPush(), nil
                }
            }
        }
    }
    return NewNonLinkable(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *NonLinkable) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NonLinkable) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["$type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val)
        }
        return nil
    }
    return res
}
// GetTypeEscaped gets the $type property value. The Type property
// returns a *string when successful
func (m *NonLinkable) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *NonLinkable) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("$type", m.GetTypeEscaped())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *NonLinkable) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetTypeEscaped sets the $type property value. The Type property
func (m *NonLinkable) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
type NonLinkableable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetTypeEscaped()(*string)
    SetTypeEscaped(value *string)()
}
