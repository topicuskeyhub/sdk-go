package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ClientClientApplication_additionalObjects 
type ClientClientApplication_additionalObjects struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The audit property
    audit AuditInfoable
    // The deleteTile property
    deleteTile *bool
    // The groupclients property
    groupclients GroupGroupClientLinkableWrapperable
    // The groups property
    groups GroupGroupLinkableWrapperable
    // The secret property
    secret GeneratedSecretable
    // The tile property
    tile LaunchpadSsoApplicationLaunchpadTileable
    // The vaultRecordCount property
    vaultRecordCount *int32
}
// NewClientClientApplication_additionalObjects instantiates a new clientClientApplication_additionalObjects and sets the default values.
func NewClientClientApplication_additionalObjects()(*ClientClientApplication_additionalObjects) {
    m := &ClientClientApplication_additionalObjects{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateClientClientApplication_additionalObjectsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateClientClientApplication_additionalObjectsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewClientClientApplication_additionalObjects(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ClientClientApplication_additionalObjects) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAudit gets the audit property value. The audit property
func (m *ClientClientApplication_additionalObjects) GetAudit()(AuditInfoable) {
    return m.audit
}
// GetDeleteTile gets the deleteTile property value. The deleteTile property
func (m *ClientClientApplication_additionalObjects) GetDeleteTile()(*bool) {
    return m.deleteTile
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ClientClientApplication_additionalObjects) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["deleteTile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeleteTile(val)
        }
        return nil
    }
    res["groupclients"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupGroupClientLinkableWrapperFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupclients(val.(GroupGroupClientLinkableWrapperable))
        }
        return nil
    }
    res["groups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupGroupLinkableWrapperFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroups(val.(GroupGroupLinkableWrapperable))
        }
        return nil
    }
    res["secret"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGeneratedSecretFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecret(val.(GeneratedSecretable))
        }
        return nil
    }
    res["tile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLaunchpadSsoApplicationLaunchpadTileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTile(val.(LaunchpadSsoApplicationLaunchpadTileable))
        }
        return nil
    }
    res["vaultRecordCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVaultRecordCount(val)
        }
        return nil
    }
    return res
}
// GetGroupclients gets the groupclients property value. The groupclients property
func (m *ClientClientApplication_additionalObjects) GetGroupclients()(GroupGroupClientLinkableWrapperable) {
    return m.groupclients
}
// GetGroups gets the groups property value. The groups property
func (m *ClientClientApplication_additionalObjects) GetGroups()(GroupGroupLinkableWrapperable) {
    return m.groups
}
// GetSecret gets the secret property value. The secret property
func (m *ClientClientApplication_additionalObjects) GetSecret()(GeneratedSecretable) {
    return m.secret
}
// GetTile gets the tile property value. The tile property
func (m *ClientClientApplication_additionalObjects) GetTile()(LaunchpadSsoApplicationLaunchpadTileable) {
    return m.tile
}
// GetVaultRecordCount gets the vaultRecordCount property value. The vaultRecordCount property
func (m *ClientClientApplication_additionalObjects) GetVaultRecordCount()(*int32) {
    return m.vaultRecordCount
}
// Serialize serializes information the current object
func (m *ClientClientApplication_additionalObjects) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("audit", m.GetAudit())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("deleteTile", m.GetDeleteTile())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("groupclients", m.GetGroupclients())
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
        err := writer.WriteObjectValue("secret", m.GetSecret())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("tile", m.GetTile())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("vaultRecordCount", m.GetVaultRecordCount())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ClientClientApplication_additionalObjects) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAudit sets the audit property value. The audit property
func (m *ClientClientApplication_additionalObjects) SetAudit(value AuditInfoable)() {
    m.audit = value
}
// SetDeleteTile sets the deleteTile property value. The deleteTile property
func (m *ClientClientApplication_additionalObjects) SetDeleteTile(value *bool)() {
    m.deleteTile = value
}
// SetGroupclients sets the groupclients property value. The groupclients property
func (m *ClientClientApplication_additionalObjects) SetGroupclients(value GroupGroupClientLinkableWrapperable)() {
    m.groupclients = value
}
// SetGroups sets the groups property value. The groups property
func (m *ClientClientApplication_additionalObjects) SetGroups(value GroupGroupLinkableWrapperable)() {
    m.groups = value
}
// SetSecret sets the secret property value. The secret property
func (m *ClientClientApplication_additionalObjects) SetSecret(value GeneratedSecretable)() {
    m.secret = value
}
// SetTile sets the tile property value. The tile property
func (m *ClientClientApplication_additionalObjects) SetTile(value LaunchpadSsoApplicationLaunchpadTileable)() {
    m.tile = value
}
// SetVaultRecordCount sets the vaultRecordCount property value. The vaultRecordCount property
func (m *ClientClientApplication_additionalObjects) SetVaultRecordCount(value *int32)() {
    m.vaultRecordCount = value
}
// ClientClientApplication_additionalObjectsable 
type ClientClientApplication_additionalObjectsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAudit()(AuditInfoable)
    GetDeleteTile()(*bool)
    GetGroupclients()(GroupGroupClientLinkableWrapperable)
    GetGroups()(GroupGroupLinkableWrapperable)
    GetSecret()(GeneratedSecretable)
    GetTile()(LaunchpadSsoApplicationLaunchpadTileable)
    GetVaultRecordCount()(*int32)
    SetAudit(value AuditInfoable)()
    SetDeleteTile(value *bool)()
    SetGroupclients(value GroupGroupClientLinkableWrapperable)()
    SetGroups(value GroupGroupLinkableWrapperable)()
    SetSecret(value GeneratedSecretable)()
    SetTile(value LaunchpadSsoApplicationLaunchpadTileable)()
    SetVaultRecordCount(value *int32)()
}
