package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthAccount_additionalObjects 
type AuthAccount_additionalObjects struct {
    // The activeLogin property
    activeLogin *bool
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The audit property
    audit AuditInfoable
    // The groups property
    groups GroupAccountGroupLinkableWrapperable
    // The pendingRecoveryRequests property
    pendingRecoveryRequests AuthAccountRecoveryStatusable
    // The settings property
    settings AuthAccountSettingsable
    // The storedAttributes property
    storedAttributes AuthStoredAccountAttributesable
    // The vault property
    vault VaultVaultable
}
// NewAuthAccount_additionalObjects instantiates a new authAccount_additionalObjects and sets the default values.
func NewAuthAccount_additionalObjects()(*AuthAccount_additionalObjects) {
    m := &AuthAccount_additionalObjects{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAuthAccount_additionalObjectsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthAccount_additionalObjectsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthAccount_additionalObjects(), nil
}
// GetActiveLogin gets the activeLogin property value. The activeLogin property
func (m *AuthAccount_additionalObjects) GetActiveLogin()(*bool) {
    return m.activeLogin
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AuthAccount_additionalObjects) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAudit gets the audit property value. The audit property
func (m *AuthAccount_additionalObjects) GetAudit()(AuditInfoable) {
    return m.audit
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthAccount_additionalObjects) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["activeLogin"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActiveLogin(val)
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
    res["groups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupAccountGroupLinkableWrapperFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroups(val.(GroupAccountGroupLinkableWrapperable))
        }
        return nil
    }
    res["pendingRecoveryRequests"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthAccountRecoveryStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPendingRecoveryRequests(val.(AuthAccountRecoveryStatusable))
        }
        return nil
    }
    res["settings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthAccountSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettings(val.(AuthAccountSettingsable))
        }
        return nil
    }
    res["storedAttributes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthStoredAccountAttributesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStoredAttributes(val.(AuthStoredAccountAttributesable))
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
    return res
}
// GetGroups gets the groups property value. The groups property
func (m *AuthAccount_additionalObjects) GetGroups()(GroupAccountGroupLinkableWrapperable) {
    return m.groups
}
// GetPendingRecoveryRequests gets the pendingRecoveryRequests property value. The pendingRecoveryRequests property
func (m *AuthAccount_additionalObjects) GetPendingRecoveryRequests()(AuthAccountRecoveryStatusable) {
    return m.pendingRecoveryRequests
}
// GetSettings gets the settings property value. The settings property
func (m *AuthAccount_additionalObjects) GetSettings()(AuthAccountSettingsable) {
    return m.settings
}
// GetStoredAttributes gets the storedAttributes property value. The storedAttributes property
func (m *AuthAccount_additionalObjects) GetStoredAttributes()(AuthStoredAccountAttributesable) {
    return m.storedAttributes
}
// GetVault gets the vault property value. The vault property
func (m *AuthAccount_additionalObjects) GetVault()(VaultVaultable) {
    return m.vault
}
// Serialize serializes information the current object
func (m *AuthAccount_additionalObjects) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("audit", m.GetAudit())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("groups", m.GetGroups())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("pendingRecoveryRequests", m.GetPendingRecoveryRequests())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("settings", m.GetSettings())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("storedAttributes", m.GetStoredAttributes())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActiveLogin sets the activeLogin property value. The activeLogin property
func (m *AuthAccount_additionalObjects) SetActiveLogin(value *bool)() {
    m.activeLogin = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AuthAccount_additionalObjects) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAudit sets the audit property value. The audit property
func (m *AuthAccount_additionalObjects) SetAudit(value AuditInfoable)() {
    m.audit = value
}
// SetGroups sets the groups property value. The groups property
func (m *AuthAccount_additionalObjects) SetGroups(value GroupAccountGroupLinkableWrapperable)() {
    m.groups = value
}
// SetPendingRecoveryRequests sets the pendingRecoveryRequests property value. The pendingRecoveryRequests property
func (m *AuthAccount_additionalObjects) SetPendingRecoveryRequests(value AuthAccountRecoveryStatusable)() {
    m.pendingRecoveryRequests = value
}
// SetSettings sets the settings property value. The settings property
func (m *AuthAccount_additionalObjects) SetSettings(value AuthAccountSettingsable)() {
    m.settings = value
}
// SetStoredAttributes sets the storedAttributes property value. The storedAttributes property
func (m *AuthAccount_additionalObjects) SetStoredAttributes(value AuthStoredAccountAttributesable)() {
    m.storedAttributes = value
}
// SetVault sets the vault property value. The vault property
func (m *AuthAccount_additionalObjects) SetVault(value VaultVaultable)() {
    m.vault = value
}
// AuthAccount_additionalObjectsable 
type AuthAccount_additionalObjectsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActiveLogin()(*bool)
    GetAudit()(AuditInfoable)
    GetGroups()(GroupAccountGroupLinkableWrapperable)
    GetPendingRecoveryRequests()(AuthAccountRecoveryStatusable)
    GetSettings()(AuthAccountSettingsable)
    GetStoredAttributes()(AuthStoredAccountAttributesable)
    GetVault()(VaultVaultable)
    SetActiveLogin(value *bool)()
    SetAudit(value AuditInfoable)()
    SetGroups(value GroupAccountGroupLinkableWrapperable)()
    SetPendingRecoveryRequests(value AuthAccountRecoveryStatusable)()
    SetSettings(value AuthAccountSettingsable)()
    SetStoredAttributes(value AuthStoredAccountAttributesable)()
    SetVault(value VaultVaultable)()
}
