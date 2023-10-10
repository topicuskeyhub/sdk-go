package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupGroup_additionalObjects 
type GroupGroup_additionalObjects struct {
    // The accounts property
    accounts GroupGroupAccountLinkableWrapperable
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The administeredClients property
    administeredClients ClientClientApplicationLinkableWrapperable
    // The administeredSystems property
    administeredSystems ProvisioningProvisionedSystemLinkableWrapperable
    // The admins property
    admins GroupGroupAccountLinkableWrapperable
    // The audit property
    audit AuditInfoable
    // The authorizedGroups property
    authorizedGroups GroupAuthorizedGroupsWrapperable
    // The clientPermissions property
    clientPermissions ClientOAuth2ClientPermissionWithClientLinkableWrapperable
    // The clients property
    clients GroupGroupClientLinkableWrapperable
    // The contentAdministeredSystems property
    contentAdministeredSystems ProvisioningProvisionedSystemLinkableWrapperable
    // The groupauditinginfo property
    groupauditinginfo GroupGroupAuditingInfoable
    // The groupinfo property
    groupinfo GroupGroupInfoable
    // The helpdesk property
    helpdesk DirectoryAccountDirectorySummaryLinkableWrapperable
    // The markers property
    markers MarkItemMarkersable
    // The myaccount property
    myaccount GroupGroupAccountable
    // The mydelegatedaccount property
    mydelegatedaccount GroupGroupAccountable
    // The nestedGroups property
    nestedGroups GroupGroupPrimerLinkableWrapperable
    // The ownedClients property
    ownedClients ClientClientApplicationLinkableWrapperable
    // The ownedDirectories property
    ownedDirectories DirectoryAccountDirectoryLinkableWrapperable
    // The ownedGroupsOnSystem property
    ownedGroupsOnSystem ProvisioningOwnedGroupOnSystemsWrapperable
    // The ownedOrganizationalUnits property
    ownedOrganizationalUnits OrganizationOrganizationalUnitLinkableWrapperable
    // The ownedSystems property
    ownedSystems ProvisioningProvisionedSystemLinkableWrapperable
    // The recentAudits property
    recentAudits GroupGroupAuditLinkableWrapperable
    // The requeststatus property
    requeststatus *GroupGroupRequestStatus
    // The serviceAccounts property
    serviceAccounts ServiceaccountServiceAccountLinkableWrapperable
    // The systems property
    systems GroupProvisioningGroupLinkableWrapperable
    // The vault property
    vault VaultVaultable
    // The webhooks property
    webhooks WebhookWebhookLinkableWrapperable
}
// NewGroupGroup_additionalObjects instantiates a new groupGroup_additionalObjects and sets the default values.
func NewGroupGroup_additionalObjects()(*GroupGroup_additionalObjects) {
    m := &GroupGroup_additionalObjects{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGroupGroup_additionalObjectsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupGroup_additionalObjectsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupGroup_additionalObjects(), nil
}
// GetAccounts gets the accounts property value. The accounts property
func (m *GroupGroup_additionalObjects) GetAccounts()(GroupGroupAccountLinkableWrapperable) {
    return m.accounts
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GroupGroup_additionalObjects) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAdministeredClients gets the administeredClients property value. The administeredClients property
func (m *GroupGroup_additionalObjects) GetAdministeredClients()(ClientClientApplicationLinkableWrapperable) {
    return m.administeredClients
}
// GetAdministeredSystems gets the administeredSystems property value. The administeredSystems property
func (m *GroupGroup_additionalObjects) GetAdministeredSystems()(ProvisioningProvisionedSystemLinkableWrapperable) {
    return m.administeredSystems
}
// GetAdmins gets the admins property value. The admins property
func (m *GroupGroup_additionalObjects) GetAdmins()(GroupGroupAccountLinkableWrapperable) {
    return m.admins
}
// GetAudit gets the audit property value. The audit property
func (m *GroupGroup_additionalObjects) GetAudit()(AuditInfoable) {
    return m.audit
}
// GetAuthorizedGroups gets the authorizedGroups property value. The authorizedGroups property
func (m *GroupGroup_additionalObjects) GetAuthorizedGroups()(GroupAuthorizedGroupsWrapperable) {
    return m.authorizedGroups
}
// GetClientPermissions gets the clientPermissions property value. The clientPermissions property
func (m *GroupGroup_additionalObjects) GetClientPermissions()(ClientOAuth2ClientPermissionWithClientLinkableWrapperable) {
    return m.clientPermissions
}
// GetClients gets the clients property value. The clients property
func (m *GroupGroup_additionalObjects) GetClients()(GroupGroupClientLinkableWrapperable) {
    return m.clients
}
// GetContentAdministeredSystems gets the contentAdministeredSystems property value. The contentAdministeredSystems property
func (m *GroupGroup_additionalObjects) GetContentAdministeredSystems()(ProvisioningProvisionedSystemLinkableWrapperable) {
    return m.contentAdministeredSystems
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupGroup_additionalObjects) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["accounts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupGroupAccountLinkableWrapperFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccounts(val.(GroupGroupAccountLinkableWrapperable))
        }
        return nil
    }
    res["administeredClients"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateClientClientApplicationLinkableWrapperFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdministeredClients(val.(ClientClientApplicationLinkableWrapperable))
        }
        return nil
    }
    res["administeredSystems"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProvisioningProvisionedSystemLinkableWrapperFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdministeredSystems(val.(ProvisioningProvisionedSystemLinkableWrapperable))
        }
        return nil
    }
    res["admins"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupGroupAccountLinkableWrapperFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdmins(val.(GroupGroupAccountLinkableWrapperable))
        }
        return nil
    }
    res["audit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuditInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAudit(val.(AuditInfoable))
        }
        return nil
    }
    res["authorizedGroups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupAuthorizedGroupsWrapperFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthorizedGroups(val.(GroupAuthorizedGroupsWrapperable))
        }
        return nil
    }
    res["clientPermissions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateClientOAuth2ClientPermissionWithClientLinkableWrapperFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientPermissions(val.(ClientOAuth2ClientPermissionWithClientLinkableWrapperable))
        }
        return nil
    }
    res["clients"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupGroupClientLinkableWrapperFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClients(val.(GroupGroupClientLinkableWrapperable))
        }
        return nil
    }
    res["contentAdministeredSystems"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProvisioningProvisionedSystemLinkableWrapperFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentAdministeredSystems(val.(ProvisioningProvisionedSystemLinkableWrapperable))
        }
        return nil
    }
    res["groupauditinginfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupGroupAuditingInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupauditinginfo(val.(GroupGroupAuditingInfoable))
        }
        return nil
    }
    res["groupinfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupGroupInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupinfo(val.(GroupGroupInfoable))
        }
        return nil
    }
    res["helpdesk"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDirectoryAccountDirectorySummaryLinkableWrapperFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHelpdesk(val.(DirectoryAccountDirectorySummaryLinkableWrapperable))
        }
        return nil
    }
    res["markers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMarkItemMarkersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMarkers(val.(MarkItemMarkersable))
        }
        return nil
    }
    res["myaccount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupGroupAccountFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMyaccount(val.(GroupGroupAccountable))
        }
        return nil
    }
    res["mydelegatedaccount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupGroupAccountFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMydelegatedaccount(val.(GroupGroupAccountable))
        }
        return nil
    }
    res["nestedGroups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupGroupPrimerLinkableWrapperFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNestedGroups(val.(GroupGroupPrimerLinkableWrapperable))
        }
        return nil
    }
    res["ownedClients"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateClientClientApplicationLinkableWrapperFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnedClients(val.(ClientClientApplicationLinkableWrapperable))
        }
        return nil
    }
    res["ownedDirectories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDirectoryAccountDirectoryLinkableWrapperFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnedDirectories(val.(DirectoryAccountDirectoryLinkableWrapperable))
        }
        return nil
    }
    res["ownedGroupsOnSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProvisioningOwnedGroupOnSystemsWrapperFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnedGroupsOnSystem(val.(ProvisioningOwnedGroupOnSystemsWrapperable))
        }
        return nil
    }
    res["ownedOrganizationalUnits"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOrganizationOrganizationalUnitLinkableWrapperFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnedOrganizationalUnits(val.(OrganizationOrganizationalUnitLinkableWrapperable))
        }
        return nil
    }
    res["ownedSystems"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProvisioningProvisionedSystemLinkableWrapperFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnedSystems(val.(ProvisioningProvisionedSystemLinkableWrapperable))
        }
        return nil
    }
    res["recentAudits"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupGroupAuditLinkableWrapperFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecentAudits(val.(GroupGroupAuditLinkableWrapperable))
        }
        return nil
    }
    res["requeststatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseGroupGroupRequestStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequeststatus(val.(*GroupGroupRequestStatus))
        }
        return nil
    }
    res["serviceAccounts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateServiceaccountServiceAccountLinkableWrapperFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServiceAccounts(val.(ServiceaccountServiceAccountLinkableWrapperable))
        }
        return nil
    }
    res["systems"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupProvisioningGroupLinkableWrapperFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystems(val.(GroupProvisioningGroupLinkableWrapperable))
        }
        return nil
    }
    res["vault"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateVaultVaultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVault(val.(VaultVaultable))
        }
        return nil
    }
    res["webhooks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWebhookWebhookLinkableWrapperFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebhooks(val.(WebhookWebhookLinkableWrapperable))
        }
        return nil
    }
    return res
}
// GetGroupauditinginfo gets the groupauditinginfo property value. The groupauditinginfo property
func (m *GroupGroup_additionalObjects) GetGroupauditinginfo()(GroupGroupAuditingInfoable) {
    return m.groupauditinginfo
}
// GetGroupinfo gets the groupinfo property value. The groupinfo property
func (m *GroupGroup_additionalObjects) GetGroupinfo()(GroupGroupInfoable) {
    return m.groupinfo
}
// GetHelpdesk gets the helpdesk property value. The helpdesk property
func (m *GroupGroup_additionalObjects) GetHelpdesk()(DirectoryAccountDirectorySummaryLinkableWrapperable) {
    return m.helpdesk
}
// GetMarkers gets the markers property value. The markers property
func (m *GroupGroup_additionalObjects) GetMarkers()(MarkItemMarkersable) {
    return m.markers
}
// GetMyaccount gets the myaccount property value. The myaccount property
func (m *GroupGroup_additionalObjects) GetMyaccount()(GroupGroupAccountable) {
    return m.myaccount
}
// GetMydelegatedaccount gets the mydelegatedaccount property value. The mydelegatedaccount property
func (m *GroupGroup_additionalObjects) GetMydelegatedaccount()(GroupGroupAccountable) {
    return m.mydelegatedaccount
}
// GetNestedGroups gets the nestedGroups property value. The nestedGroups property
func (m *GroupGroup_additionalObjects) GetNestedGroups()(GroupGroupPrimerLinkableWrapperable) {
    return m.nestedGroups
}
// GetOwnedClients gets the ownedClients property value. The ownedClients property
func (m *GroupGroup_additionalObjects) GetOwnedClients()(ClientClientApplicationLinkableWrapperable) {
    return m.ownedClients
}
// GetOwnedDirectories gets the ownedDirectories property value. The ownedDirectories property
func (m *GroupGroup_additionalObjects) GetOwnedDirectories()(DirectoryAccountDirectoryLinkableWrapperable) {
    return m.ownedDirectories
}
// GetOwnedGroupsOnSystem gets the ownedGroupsOnSystem property value. The ownedGroupsOnSystem property
func (m *GroupGroup_additionalObjects) GetOwnedGroupsOnSystem()(ProvisioningOwnedGroupOnSystemsWrapperable) {
    return m.ownedGroupsOnSystem
}
// GetOwnedOrganizationalUnits gets the ownedOrganizationalUnits property value. The ownedOrganizationalUnits property
func (m *GroupGroup_additionalObjects) GetOwnedOrganizationalUnits()(OrganizationOrganizationalUnitLinkableWrapperable) {
    return m.ownedOrganizationalUnits
}
// GetOwnedSystems gets the ownedSystems property value. The ownedSystems property
func (m *GroupGroup_additionalObjects) GetOwnedSystems()(ProvisioningProvisionedSystemLinkableWrapperable) {
    return m.ownedSystems
}
// GetRecentAudits gets the recentAudits property value. The recentAudits property
func (m *GroupGroup_additionalObjects) GetRecentAudits()(GroupGroupAuditLinkableWrapperable) {
    return m.recentAudits
}
// GetRequeststatus gets the requeststatus property value. The requeststatus property
func (m *GroupGroup_additionalObjects) GetRequeststatus()(*GroupGroupRequestStatus) {
    return m.requeststatus
}
// GetServiceAccounts gets the serviceAccounts property value. The serviceAccounts property
func (m *GroupGroup_additionalObjects) GetServiceAccounts()(ServiceaccountServiceAccountLinkableWrapperable) {
    return m.serviceAccounts
}
// GetSystems gets the systems property value. The systems property
func (m *GroupGroup_additionalObjects) GetSystems()(GroupProvisioningGroupLinkableWrapperable) {
    return m.systems
}
// GetVault gets the vault property value. The vault property
func (m *GroupGroup_additionalObjects) GetVault()(VaultVaultable) {
    return m.vault
}
// GetWebhooks gets the webhooks property value. The webhooks property
func (m *GroupGroup_additionalObjects) GetWebhooks()(WebhookWebhookLinkableWrapperable) {
    return m.webhooks
}
// Serialize serializes information the current object
func (m *GroupGroup_additionalObjects) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("accounts", m.GetAccounts())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("administeredClients", m.GetAdministeredClients())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("administeredSystems", m.GetAdministeredSystems())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("admins", m.GetAdmins())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("audit", m.GetAudit())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("authorizedGroups", m.GetAuthorizedGroups())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("clientPermissions", m.GetClientPermissions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("clients", m.GetClients())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("contentAdministeredSystems", m.GetContentAdministeredSystems())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("groupauditinginfo", m.GetGroupauditinginfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("groupinfo", m.GetGroupinfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("helpdesk", m.GetHelpdesk())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("markers", m.GetMarkers())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("myaccount", m.GetMyaccount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("mydelegatedaccount", m.GetMydelegatedaccount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("nestedGroups", m.GetNestedGroups())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("ownedClients", m.GetOwnedClients())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("ownedDirectories", m.GetOwnedDirectories())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("ownedGroupsOnSystem", m.GetOwnedGroupsOnSystem())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("ownedOrganizationalUnits", m.GetOwnedOrganizationalUnits())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("ownedSystems", m.GetOwnedSystems())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("recentAudits", m.GetRecentAudits())
        if err != nil {
            return err
        }
    }
    if m.GetRequeststatus() != nil {
        cast := (*m.GetRequeststatus()).String()
        err := writer.WriteStringValue("requeststatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("serviceAccounts", m.GetServiceAccounts())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("systems", m.GetSystems())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("vault", m.GetVault())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("webhooks", m.GetWebhooks())
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
// SetAccounts sets the accounts property value. The accounts property
func (m *GroupGroup_additionalObjects) SetAccounts(value GroupGroupAccountLinkableWrapperable)() {
    m.accounts = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GroupGroup_additionalObjects) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAdministeredClients sets the administeredClients property value. The administeredClients property
func (m *GroupGroup_additionalObjects) SetAdministeredClients(value ClientClientApplicationLinkableWrapperable)() {
    m.administeredClients = value
}
// SetAdministeredSystems sets the administeredSystems property value. The administeredSystems property
func (m *GroupGroup_additionalObjects) SetAdministeredSystems(value ProvisioningProvisionedSystemLinkableWrapperable)() {
    m.administeredSystems = value
}
// SetAdmins sets the admins property value. The admins property
func (m *GroupGroup_additionalObjects) SetAdmins(value GroupGroupAccountLinkableWrapperable)() {
    m.admins = value
}
// SetAudit sets the audit property value. The audit property
func (m *GroupGroup_additionalObjects) SetAudit(value AuditInfoable)() {
    m.audit = value
}
// SetAuthorizedGroups sets the authorizedGroups property value. The authorizedGroups property
func (m *GroupGroup_additionalObjects) SetAuthorizedGroups(value GroupAuthorizedGroupsWrapperable)() {
    m.authorizedGroups = value
}
// SetClientPermissions sets the clientPermissions property value. The clientPermissions property
func (m *GroupGroup_additionalObjects) SetClientPermissions(value ClientOAuth2ClientPermissionWithClientLinkableWrapperable)() {
    m.clientPermissions = value
}
// SetClients sets the clients property value. The clients property
func (m *GroupGroup_additionalObjects) SetClients(value GroupGroupClientLinkableWrapperable)() {
    m.clients = value
}
// SetContentAdministeredSystems sets the contentAdministeredSystems property value. The contentAdministeredSystems property
func (m *GroupGroup_additionalObjects) SetContentAdministeredSystems(value ProvisioningProvisionedSystemLinkableWrapperable)() {
    m.contentAdministeredSystems = value
}
// SetGroupauditinginfo sets the groupauditinginfo property value. The groupauditinginfo property
func (m *GroupGroup_additionalObjects) SetGroupauditinginfo(value GroupGroupAuditingInfoable)() {
    m.groupauditinginfo = value
}
// SetGroupinfo sets the groupinfo property value. The groupinfo property
func (m *GroupGroup_additionalObjects) SetGroupinfo(value GroupGroupInfoable)() {
    m.groupinfo = value
}
// SetHelpdesk sets the helpdesk property value. The helpdesk property
func (m *GroupGroup_additionalObjects) SetHelpdesk(value DirectoryAccountDirectorySummaryLinkableWrapperable)() {
    m.helpdesk = value
}
// SetMarkers sets the markers property value. The markers property
func (m *GroupGroup_additionalObjects) SetMarkers(value MarkItemMarkersable)() {
    m.markers = value
}
// SetMyaccount sets the myaccount property value. The myaccount property
func (m *GroupGroup_additionalObjects) SetMyaccount(value GroupGroupAccountable)() {
    m.myaccount = value
}
// SetMydelegatedaccount sets the mydelegatedaccount property value. The mydelegatedaccount property
func (m *GroupGroup_additionalObjects) SetMydelegatedaccount(value GroupGroupAccountable)() {
    m.mydelegatedaccount = value
}
// SetNestedGroups sets the nestedGroups property value. The nestedGroups property
func (m *GroupGroup_additionalObjects) SetNestedGroups(value GroupGroupPrimerLinkableWrapperable)() {
    m.nestedGroups = value
}
// SetOwnedClients sets the ownedClients property value. The ownedClients property
func (m *GroupGroup_additionalObjects) SetOwnedClients(value ClientClientApplicationLinkableWrapperable)() {
    m.ownedClients = value
}
// SetOwnedDirectories sets the ownedDirectories property value. The ownedDirectories property
func (m *GroupGroup_additionalObjects) SetOwnedDirectories(value DirectoryAccountDirectoryLinkableWrapperable)() {
    m.ownedDirectories = value
}
// SetOwnedGroupsOnSystem sets the ownedGroupsOnSystem property value. The ownedGroupsOnSystem property
func (m *GroupGroup_additionalObjects) SetOwnedGroupsOnSystem(value ProvisioningOwnedGroupOnSystemsWrapperable)() {
    m.ownedGroupsOnSystem = value
}
// SetOwnedOrganizationalUnits sets the ownedOrganizationalUnits property value. The ownedOrganizationalUnits property
func (m *GroupGroup_additionalObjects) SetOwnedOrganizationalUnits(value OrganizationOrganizationalUnitLinkableWrapperable)() {
    m.ownedOrganizationalUnits = value
}
// SetOwnedSystems sets the ownedSystems property value. The ownedSystems property
func (m *GroupGroup_additionalObjects) SetOwnedSystems(value ProvisioningProvisionedSystemLinkableWrapperable)() {
    m.ownedSystems = value
}
// SetRecentAudits sets the recentAudits property value. The recentAudits property
func (m *GroupGroup_additionalObjects) SetRecentAudits(value GroupGroupAuditLinkableWrapperable)() {
    m.recentAudits = value
}
// SetRequeststatus sets the requeststatus property value. The requeststatus property
func (m *GroupGroup_additionalObjects) SetRequeststatus(value *GroupGroupRequestStatus)() {
    m.requeststatus = value
}
// SetServiceAccounts sets the serviceAccounts property value. The serviceAccounts property
func (m *GroupGroup_additionalObjects) SetServiceAccounts(value ServiceaccountServiceAccountLinkableWrapperable)() {
    m.serviceAccounts = value
}
// SetSystems sets the systems property value. The systems property
func (m *GroupGroup_additionalObjects) SetSystems(value GroupProvisioningGroupLinkableWrapperable)() {
    m.systems = value
}
// SetVault sets the vault property value. The vault property
func (m *GroupGroup_additionalObjects) SetVault(value VaultVaultable)() {
    m.vault = value
}
// SetWebhooks sets the webhooks property value. The webhooks property
func (m *GroupGroup_additionalObjects) SetWebhooks(value WebhookWebhookLinkableWrapperable)() {
    m.webhooks = value
}
// GroupGroup_additionalObjectsable 
type GroupGroup_additionalObjectsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccounts()(GroupGroupAccountLinkableWrapperable)
    GetAdministeredClients()(ClientClientApplicationLinkableWrapperable)
    GetAdministeredSystems()(ProvisioningProvisionedSystemLinkableWrapperable)
    GetAdmins()(GroupGroupAccountLinkableWrapperable)
    GetAudit()(AuditInfoable)
    GetAuthorizedGroups()(GroupAuthorizedGroupsWrapperable)
    GetClientPermissions()(ClientOAuth2ClientPermissionWithClientLinkableWrapperable)
    GetClients()(GroupGroupClientLinkableWrapperable)
    GetContentAdministeredSystems()(ProvisioningProvisionedSystemLinkableWrapperable)
    GetGroupauditinginfo()(GroupGroupAuditingInfoable)
    GetGroupinfo()(GroupGroupInfoable)
    GetHelpdesk()(DirectoryAccountDirectorySummaryLinkableWrapperable)
    GetMarkers()(MarkItemMarkersable)
    GetMyaccount()(GroupGroupAccountable)
    GetMydelegatedaccount()(GroupGroupAccountable)
    GetNestedGroups()(GroupGroupPrimerLinkableWrapperable)
    GetOwnedClients()(ClientClientApplicationLinkableWrapperable)
    GetOwnedDirectories()(DirectoryAccountDirectoryLinkableWrapperable)
    GetOwnedGroupsOnSystem()(ProvisioningOwnedGroupOnSystemsWrapperable)
    GetOwnedOrganizationalUnits()(OrganizationOrganizationalUnitLinkableWrapperable)
    GetOwnedSystems()(ProvisioningProvisionedSystemLinkableWrapperable)
    GetRecentAudits()(GroupGroupAuditLinkableWrapperable)
    GetRequeststatus()(*GroupGroupRequestStatus)
    GetServiceAccounts()(ServiceaccountServiceAccountLinkableWrapperable)
    GetSystems()(GroupProvisioningGroupLinkableWrapperable)
    GetVault()(VaultVaultable)
    GetWebhooks()(WebhookWebhookLinkableWrapperable)
    SetAccounts(value GroupGroupAccountLinkableWrapperable)()
    SetAdministeredClients(value ClientClientApplicationLinkableWrapperable)()
    SetAdministeredSystems(value ProvisioningProvisionedSystemLinkableWrapperable)()
    SetAdmins(value GroupGroupAccountLinkableWrapperable)()
    SetAudit(value AuditInfoable)()
    SetAuthorizedGroups(value GroupAuthorizedGroupsWrapperable)()
    SetClientPermissions(value ClientOAuth2ClientPermissionWithClientLinkableWrapperable)()
    SetClients(value GroupGroupClientLinkableWrapperable)()
    SetContentAdministeredSystems(value ProvisioningProvisionedSystemLinkableWrapperable)()
    SetGroupauditinginfo(value GroupGroupAuditingInfoable)()
    SetGroupinfo(value GroupGroupInfoable)()
    SetHelpdesk(value DirectoryAccountDirectorySummaryLinkableWrapperable)()
    SetMarkers(value MarkItemMarkersable)()
    SetMyaccount(value GroupGroupAccountable)()
    SetMydelegatedaccount(value GroupGroupAccountable)()
    SetNestedGroups(value GroupGroupPrimerLinkableWrapperable)()
    SetOwnedClients(value ClientClientApplicationLinkableWrapperable)()
    SetOwnedDirectories(value DirectoryAccountDirectoryLinkableWrapperable)()
    SetOwnedGroupsOnSystem(value ProvisioningOwnedGroupOnSystemsWrapperable)()
    SetOwnedOrganizationalUnits(value OrganizationOrganizationalUnitLinkableWrapperable)()
    SetOwnedSystems(value ProvisioningProvisionedSystemLinkableWrapperable)()
    SetRecentAudits(value GroupGroupAuditLinkableWrapperable)()
    SetRequeststatus(value *GroupGroupRequestStatus)()
    SetServiceAccounts(value ServiceaccountServiceAccountLinkableWrapperable)()
    SetSystems(value GroupProvisioningGroupLinkableWrapperable)()
    SetVault(value VaultVaultable)()
    SetWebhooks(value WebhookWebhookLinkableWrapperable)()
}
