package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type VaultVaultRecord_additionalObjects struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The audit property
    audit AuditInfoable
    // The deleteTile property
    deleteTile *bool
    // The parent property
    parent VaultVaultRecordPrimerable
    // The passwordMetadata property
    passwordMetadata VaultPasswordMetadataable
    // The secret property
    secret VaultVaultRecordSecretsable
    // The shares property
    shares VaultVaultRecordPrimerLinkableWrapperable
    // The shareSummary property
    shareSummary VaultVaultRecordShareSummaryable
    // The tile property
    tile LaunchpadVaultRecordLaunchpadTileable
    // The vaultholder property
    vaultholder VaultVaultHolderable
}
// NewVaultVaultRecord_additionalObjects instantiates a new VaultVaultRecord_additionalObjects and sets the default values.
func NewVaultVaultRecord_additionalObjects()(*VaultVaultRecord_additionalObjects) {
    m := &VaultVaultRecord_additionalObjects{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateVaultVaultRecord_additionalObjectsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateVaultVaultRecord_additionalObjectsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVaultVaultRecord_additionalObjects(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *VaultVaultRecord_additionalObjects) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAudit gets the audit property value. The audit property
// returns a AuditInfoable when successful
func (m *VaultVaultRecord_additionalObjects) GetAudit()(AuditInfoable) {
    return m.audit
}
// GetDeleteTile gets the deleteTile property value. The deleteTile property
// returns a *bool when successful
func (m *VaultVaultRecord_additionalObjects) GetDeleteTile()(*bool) {
    return m.deleteTile
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *VaultVaultRecord_additionalObjects) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["parent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateVaultVaultRecordPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParent(val.(VaultVaultRecordPrimerable))
        }
        return nil
    }
    res["passwordMetadata"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateVaultPasswordMetadataFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordMetadata(val.(VaultPasswordMetadataable))
        }
        return nil
    }
    res["secret"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateVaultVaultRecordSecretsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecret(val.(VaultVaultRecordSecretsable))
        }
        return nil
    }
    res["shares"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateVaultVaultRecordPrimerLinkableWrapperFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShares(val.(VaultVaultRecordPrimerLinkableWrapperable))
        }
        return nil
    }
    res["shareSummary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateVaultVaultRecordShareSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShareSummary(val.(VaultVaultRecordShareSummaryable))
        }
        return nil
    }
    res["tile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLaunchpadVaultRecordLaunchpadTileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTile(val.(LaunchpadVaultRecordLaunchpadTileable))
        }
        return nil
    }
    res["vaultholder"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateVaultVaultHolderFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVaultholder(val.(VaultVaultHolderable))
        }
        return nil
    }
    return res
}
// GetParent gets the parent property value. The parent property
// returns a VaultVaultRecordPrimerable when successful
func (m *VaultVaultRecord_additionalObjects) GetParent()(VaultVaultRecordPrimerable) {
    return m.parent
}
// GetPasswordMetadata gets the passwordMetadata property value. The passwordMetadata property
// returns a VaultPasswordMetadataable when successful
func (m *VaultVaultRecord_additionalObjects) GetPasswordMetadata()(VaultPasswordMetadataable) {
    return m.passwordMetadata
}
// GetSecret gets the secret property value. The secret property
// returns a VaultVaultRecordSecretsable when successful
func (m *VaultVaultRecord_additionalObjects) GetSecret()(VaultVaultRecordSecretsable) {
    return m.secret
}
// GetShares gets the shares property value. The shares property
// returns a VaultVaultRecordPrimerLinkableWrapperable when successful
func (m *VaultVaultRecord_additionalObjects) GetShares()(VaultVaultRecordPrimerLinkableWrapperable) {
    return m.shares
}
// GetShareSummary gets the shareSummary property value. The shareSummary property
// returns a VaultVaultRecordShareSummaryable when successful
func (m *VaultVaultRecord_additionalObjects) GetShareSummary()(VaultVaultRecordShareSummaryable) {
    return m.shareSummary
}
// GetTile gets the tile property value. The tile property
// returns a LaunchpadVaultRecordLaunchpadTileable when successful
func (m *VaultVaultRecord_additionalObjects) GetTile()(LaunchpadVaultRecordLaunchpadTileable) {
    return m.tile
}
// GetVaultholder gets the vaultholder property value. The vaultholder property
// returns a VaultVaultHolderable when successful
func (m *VaultVaultRecord_additionalObjects) GetVaultholder()(VaultVaultHolderable) {
    return m.vaultholder
}
// Serialize serializes information the current object
func (m *VaultVaultRecord_additionalObjects) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteObjectValue("parent", m.GetParent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("passwordMetadata", m.GetPasswordMetadata())
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
        err := writer.WriteObjectValue("shares", m.GetShares())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("shareSummary", m.GetShareSummary())
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
        err := writer.WriteObjectValue("vaultholder", m.GetVaultholder())
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
func (m *VaultVaultRecord_additionalObjects) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAudit sets the audit property value. The audit property
func (m *VaultVaultRecord_additionalObjects) SetAudit(value AuditInfoable)() {
    m.audit = value
}
// SetDeleteTile sets the deleteTile property value. The deleteTile property
func (m *VaultVaultRecord_additionalObjects) SetDeleteTile(value *bool)() {
    m.deleteTile = value
}
// SetParent sets the parent property value. The parent property
func (m *VaultVaultRecord_additionalObjects) SetParent(value VaultVaultRecordPrimerable)() {
    m.parent = value
}
// SetPasswordMetadata sets the passwordMetadata property value. The passwordMetadata property
func (m *VaultVaultRecord_additionalObjects) SetPasswordMetadata(value VaultPasswordMetadataable)() {
    m.passwordMetadata = value
}
// SetSecret sets the secret property value. The secret property
func (m *VaultVaultRecord_additionalObjects) SetSecret(value VaultVaultRecordSecretsable)() {
    m.secret = value
}
// SetShares sets the shares property value. The shares property
func (m *VaultVaultRecord_additionalObjects) SetShares(value VaultVaultRecordPrimerLinkableWrapperable)() {
    m.shares = value
}
// SetShareSummary sets the shareSummary property value. The shareSummary property
func (m *VaultVaultRecord_additionalObjects) SetShareSummary(value VaultVaultRecordShareSummaryable)() {
    m.shareSummary = value
}
// SetTile sets the tile property value. The tile property
func (m *VaultVaultRecord_additionalObjects) SetTile(value LaunchpadVaultRecordLaunchpadTileable)() {
    m.tile = value
}
// SetVaultholder sets the vaultholder property value. The vaultholder property
func (m *VaultVaultRecord_additionalObjects) SetVaultholder(value VaultVaultHolderable)() {
    m.vaultholder = value
}
type VaultVaultRecord_additionalObjectsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAudit()(AuditInfoable)
    GetDeleteTile()(*bool)
    GetParent()(VaultVaultRecordPrimerable)
    GetPasswordMetadata()(VaultPasswordMetadataable)
    GetSecret()(VaultVaultRecordSecretsable)
    GetShares()(VaultVaultRecordPrimerLinkableWrapperable)
    GetShareSummary()(VaultVaultRecordShareSummaryable)
    GetTile()(LaunchpadVaultRecordLaunchpadTileable)
    GetVaultholder()(VaultVaultHolderable)
    SetAudit(value AuditInfoable)()
    SetDeleteTile(value *bool)()
    SetParent(value VaultVaultRecordPrimerable)()
    SetPasswordMetadata(value VaultPasswordMetadataable)()
    SetSecret(value VaultVaultRecordSecretsable)()
    SetShares(value VaultVaultRecordPrimerLinkableWrapperable)()
    SetShareSummary(value VaultVaultRecordShareSummaryable)()
    SetTile(value LaunchpadVaultRecordLaunchpadTileable)()
    SetVaultholder(value VaultVaultHolderable)()
}
