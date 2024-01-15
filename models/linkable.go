package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Linkable 
type Linkable struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The links property
    links []RestLinkable
    // The permissions property
    permissions []AuthPermissionable
    // The Type property
    typeEscaped *string
}
// NewLinkable instantiates a new Linkable and sets the default values.
func NewLinkable()(*Linkable) {
    m := &Linkable{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateLinkableFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLinkableFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "audit.AuditRecord":
                        return NewAuditAuditRecord(), nil
                    case "audit.GroupAudit":
                        return NewAuditGroupAudit(), nil
                    case "audit.GroupAuditAccount":
                        return NewAuditGroupAuditAccount(), nil
                    case "audit.NestedGroupAudit":
                        return NewAuditNestedGroupAudit(), nil
                    case "auth.Account":
                        return NewAuthAccount(), nil
                    case "auth.AccountPrimer":
                        return NewAuthAccountPrimer(), nil
                    case "auth.InternalAccount":
                        return NewAuthInternalAccount(), nil
                    case "certificate.Certificate":
                        return NewCertificateCertificate(), nil
                    case "certificate.CertificatePrimer":
                        return NewCertificateCertificatePrimer(), nil
                    case "client.ClientApplication":
                        return NewClientClientApplication(), nil
                    case "client.ClientApplicationPrimer":
                        return NewClientClientApplicationPrimer(), nil
                    case "client.LdapClient":
                        return NewClientLdapClient(), nil
                    case "client.OAuth2Client":
                        return NewClientOAuth2Client(), nil
                    case "client.OAuth2ClientPermission":
                        return NewClientOAuth2ClientPermission(), nil
                    case "client.OAuth2ClientPermissionWithClient":
                        return NewClientOAuth2ClientPermissionWithClient(), nil
                    case "client.Saml2Client":
                        return NewClientSaml2Client(), nil
                    case "directory.AccountDirectory":
                        return NewDirectoryAccountDirectory(), nil
                    case "directory.AccountDirectoryPrimer":
                        return NewDirectoryAccountDirectoryPrimer(), nil
                    case "directory.AccountDirectorySummary":
                        return NewDirectoryAccountDirectorySummary(), nil
                    case "directory.InternalDirectory":
                        return NewDirectoryInternalDirectory(), nil
                    case "directory.LDAPDirectory":
                        return NewDirectoryLDAPDirectory(), nil
                    case "directory.MaintenanceDirectory":
                        return NewDirectoryMaintenanceDirectory(), nil
                    case "directory.OIDCDirectory":
                        return NewDirectoryOIDCDirectory(), nil
                    case "group.AccountGroup":
                        return NewGroupAccountGroup(), nil
                    case "group.Group":
                        return NewGroupGroup(), nil
                    case "group.GroupAccount":
                        return NewGroupGroupAccount(), nil
                    case "group.GroupAuditConfig":
                        return NewGroupGroupAuditConfig(), nil
                    case "group.GroupClassification":
                        return NewGroupGroupClassification(), nil
                    case "group.GroupClassificationPrimer":
                        return NewGroupGroupClassificationPrimer(), nil
                    case "group.GroupClient":
                        return NewGroupGroupClient(), nil
                    case "group.GroupFolder":
                        return NewGroupGroupFolder(), nil
                    case "group.GroupPrimer":
                        return NewGroupGroupPrimer(), nil
                    case "group.ProvisioningGroup":
                        return NewGroupProvisioningGroup(), nil
                    case "launchpad.DisplayedLaunchpadTile":
                        return NewLaunchpadDisplayedLaunchpadTile(), nil
                    case "launchpad.LaunchpadTile":
                        return NewLaunchpadLaunchpadTile(), nil
                    case "launchpad.LaunchpadTilePrimer":
                        return NewLaunchpadLaunchpadTilePrimer(), nil
                    case "launchpad.ManualLaunchpadTile":
                        return NewLaunchpadManualLaunchpadTile(), nil
                    case "launchpad.SsoApplicationLaunchpadTile":
                        return NewLaunchpadSsoApplicationLaunchpadTile(), nil
                    case "launchpad.VaultRecordLaunchpadTile":
                        return NewLaunchpadVaultRecordLaunchpadTile(), nil
                    case "organization.AccountOrganizationalUnit":
                        return NewOrganizationAccountOrganizationalUnit(), nil
                    case "organization.OrganizationalUnit":
                        return NewOrganizationOrganizationalUnit(), nil
                    case "organization.OrganizationalUnitAccount":
                        return NewOrganizationOrganizationalUnitAccount(), nil
                    case "organization.OrganizationalUnitPrimer":
                        return NewOrganizationOrganizationalUnitPrimer(), nil
                    case "provisioning.AbstractProvisionedLDAP":
                        return NewProvisioningAbstractProvisionedLDAP(), nil
                    case "provisioning.GroupOnSystem":
                        return NewProvisioningGroupOnSystem(), nil
                    case "provisioning.GroupOnSystemPrimer":
                        return NewProvisioningGroupOnSystemPrimer(), nil
                    case "provisioning.ProvisionedAccount":
                        return NewProvisioningProvisionedAccount(), nil
                    case "provisioning.ProvisionedAD":
                        return NewProvisioningProvisionedAD(), nil
                    case "provisioning.ProvisionedAzureOIDCDirectory":
                        return NewProvisioningProvisionedAzureOIDCDirectory(), nil
                    case "provisioning.ProvisionedAzureSyncLDAPDirectory":
                        return NewProvisioningProvisionedAzureSyncLDAPDirectory(), nil
                    case "provisioning.ProvisionedAzureTenant":
                        return NewProvisioningProvisionedAzureTenant(), nil
                    case "provisioning.ProvisionedInternalLDAP":
                        return NewProvisioningProvisionedInternalLDAP(), nil
                    case "provisioning.ProvisionedLDAP":
                        return NewProvisioningProvisionedLDAP(), nil
                    case "provisioning.ProvisionedLDAPDirectory":
                        return NewProvisioningProvisionedLDAPDirectory(), nil
                    case "provisioning.ProvisionedNamespace":
                        return NewProvisioningProvisionedNamespace(), nil
                    case "provisioning.ProvisionedSCIM":
                        return NewProvisioningProvisionedSCIM(), nil
                    case "provisioning.ProvisionedSystem":
                        return NewProvisioningProvisionedSystem(), nil
                    case "provisioning.ProvisionedSystemPrimer":
                        return NewProvisioningProvisionedSystemPrimer(), nil
                    case "provisioning.ProvisionedSystemSyncLog":
                        return NewProvisioningProvisionedSystemSyncLog(), nil
                    case "provisioning.ProvisionNumberSequence":
                        return NewProvisioningProvisionNumberSequence(), nil
                    case "request.AbstractApplicationModificationRequest":
                        return NewRequestAbstractApplicationModificationRequest(), nil
                    case "request.AbstractOrganizationalUnitModificationRequest":
                        return NewRequestAbstractOrganizationalUnitModificationRequest(), nil
                    case "request.AbstractProvisionedSystemModificationRequest":
                        return NewRequestAbstractProvisionedSystemModificationRequest(), nil
                    case "request.AddGroupAdminRequest":
                        return NewRequestAddGroupAdminRequest(), nil
                    case "request.CreateGroupOnSystemRequest":
                        return NewRequestCreateGroupOnSystemRequest(), nil
                    case "request.CreateGroupRequest":
                        return NewRequestCreateGroupRequest(), nil
                    case "request.CreateProvisionedNamespaceRequest":
                        return NewRequestCreateProvisionedNamespaceRequest(), nil
                    case "request.CreateServiceAccountRequest":
                        return NewRequestCreateServiceAccountRequest(), nil
                    case "request.Disable2FARequest":
                        return NewRequestDisable2FARequest(), nil
                    case "request.EnableTechnicalAdministrationRequest":
                        return NewRequestEnableTechnicalAdministrationRequest(), nil
                    case "request.ExtendAccessRequest":
                        return NewRequestExtendAccessRequest(), nil
                    case "request.GrantAccessRequest":
                        return NewRequestGrantAccessRequest(), nil
                    case "request.GrantApplicationRequest":
                        return NewRequestGrantApplicationRequest(), nil
                    case "request.GrantClientPermissionRequest":
                        return NewRequestGrantClientPermissionRequest(), nil
                    case "request.GrantGroupOnSystemRequest":
                        return NewRequestGrantGroupOnSystemRequest(), nil
                    case "request.GrantGroupOnSystemRequestRequest":
                        return NewRequestGrantGroupOnSystemRequestRequest(), nil
                    case "request.GrantServiceAccountGroupRequest":
                        return NewRequestGrantServiceAccountGroupRequest(), nil
                    case "request.JoinGroupRequest":
                        return NewRequestJoinGroupRequest(), nil
                    case "request.JoinVaultRequest":
                        return NewRequestJoinVaultRequest(), nil
                    case "request.ModificationRequest":
                        return NewRequestModificationRequest(), nil
                    case "request.PasswordResetRequest":
                        return NewRequestPasswordResetRequest(), nil
                    case "request.RemoveGroupRequest":
                        return NewRequestRemoveGroupRequest(), nil
                    case "request.RemoveOrganizationalUnitRequest":
                        return NewRequestRemoveOrganizationalUnitRequest(), nil
                    case "request.RemoveProvisionedSystemRequest":
                        return NewRequestRemoveProvisionedSystemRequest(), nil
                    case "request.ReviewAuditRequest":
                        return NewRequestReviewAuditRequest(), nil
                    case "request.RevokeAdminRequest":
                        return NewRequestRevokeAdminRequest(), nil
                    case "request.SetupAuthorizingGroupRequest":
                        return NewRequestSetupAuthorizingGroupRequest(), nil
                    case "request.SetupNestedGroupRequest":
                        return NewRequestSetupNestedGroupRequest(), nil
                    case "request.TransferApplicationAdministrationRequest":
                        return NewRequestTransferApplicationAdministrationRequest(), nil
                    case "request.TransferApplicationOwnershipRequest":
                        return NewRequestTransferApplicationOwnershipRequest(), nil
                    case "request.TransferAuditorGroupRequest":
                        return NewRequestTransferAuditorGroupRequest(), nil
                    case "request.TransferGroupOnSystemOwnershipRequest":
                        return NewRequestTransferGroupOnSystemOwnershipRequest(), nil
                    case "request.TransferOrganizationalUnitOwnershipRequest":
                        return NewRequestTransferOrganizationalUnitOwnershipRequest(), nil
                    case "request.TransferProvisionedSystemAdministrationRequest":
                        return NewRequestTransferProvisionedSystemAdministrationRequest(), nil
                    case "request.TransferProvisionedSystemContentAdministrationRequest":
                        return NewRequestTransferProvisionedSystemContentAdministrationRequest(), nil
                    case "request.TransferProvisionedSystemOwnershipRequest":
                        return NewRequestTransferProvisionedSystemOwnershipRequest(), nil
                    case "request.TransferServiceAccountAdministrationRequest":
                        return NewRequestTransferServiceAccountAdministrationRequest(), nil
                    case "request.UpdateGroupMembershipRequest":
                        return NewRequestUpdateGroupMembershipRequest(), nil
                    case "request.VerifyInternalAccountRequest":
                        return NewRequestVerifyInternalAccountRequest(), nil
                    case "serviceaccount.ServiceAccount":
                        return NewServiceaccountServiceAccount(), nil
                    case "serviceaccount.ServiceAccountGroup":
                        return NewServiceaccountServiceAccountGroup(), nil
                    case "serviceaccount.ServiceAccountPrimer":
                        return NewServiceaccountServiceAccountPrimer(), nil
                    case "vault.DeletedVaultHolder":
                        return NewVaultDeletedVaultHolder(), nil
                    case "vault.Vault":
                        return NewVaultVault(), nil
                    case "vault.VaultRecord":
                        return NewVaultVaultRecord(), nil
                    case "vault.VaultRecordPrimer":
                        return NewVaultVaultRecordPrimer(), nil
                    case "webhook.Webhook":
                        return NewWebhookWebhook(), nil
                    case "webhook.WebhookDelivery":
                        return NewWebhookWebhookDelivery(), nil
                }
            }
        }
    }
    return NewLinkable(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Linkable) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Linkable) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["links"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRestLinkFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RestLinkable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(RestLinkable)
                }
            }
            m.SetLinks(res)
        }
        return nil
    }
    res["permissions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuthPermissionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuthPermissionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AuthPermissionable)
                }
            }
            m.SetPermissions(res)
        }
        return nil
    }
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
// GetLinks gets the links property value. The links property
func (m *Linkable) GetLinks()([]RestLinkable) {
    return m.links
}
// GetPermissions gets the permissions property value. The permissions property
func (m *Linkable) GetPermissions()([]AuthPermissionable) {
    return m.permissions
}
// GetTypeEscaped gets the $type property value. The Type property
func (m *Linkable) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *Linkable) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetLinks() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLinks()))
        for i, v := range m.GetLinks() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("links", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPermissions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPermissions()))
        for i, v := range m.GetPermissions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("permissions", cast)
        if err != nil {
            return err
        }
    }
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
func (m *Linkable) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetLinks sets the links property value. The links property
func (m *Linkable) SetLinks(value []RestLinkable)() {
    m.links = value
}
// SetPermissions sets the permissions property value. The permissions property
func (m *Linkable) SetPermissions(value []AuthPermissionable)() {
    m.permissions = value
}
// SetTypeEscaped sets the $type property value. The Type property
func (m *Linkable) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// Linkableable 
type Linkableable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLinks()([]RestLinkable)
    GetPermissions()([]AuthPermissionable)
    GetTypeEscaped()(*string)
    SetLinks(value []RestLinkable)()
    SetPermissions(value []AuthPermissionable)()
    SetTypeEscaped(value *string)()
}
