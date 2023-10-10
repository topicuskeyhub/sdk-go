package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VaultDeletedVaultHolder_additionalObjects 
type VaultDeletedVaultHolder_additionalObjects struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The audit property
    audit AuditInfoable
    // The vault property
    vault VaultVaultable
}
// NewVaultDeletedVaultHolder_additionalObjects instantiates a new vaultDeletedVaultHolder_additionalObjects and sets the default values.
func NewVaultDeletedVaultHolder_additionalObjects()(*VaultDeletedVaultHolder_additionalObjects) {
    m := &VaultDeletedVaultHolder_additionalObjects{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateVaultDeletedVaultHolder_additionalObjectsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVaultDeletedVaultHolder_additionalObjectsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVaultDeletedVaultHolder_additionalObjects(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *VaultDeletedVaultHolder_additionalObjects) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAudit gets the audit property value. The audit property
func (m *VaultDeletedVaultHolder_additionalObjects) GetAudit()(AuditInfoable) {
    return m.audit
}
// GetFieldDeserializers the deserialization information for the current model
func (m *VaultDeletedVaultHolder_additionalObjects) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
// GetVault gets the vault property value. The vault property
func (m *VaultDeletedVaultHolder_additionalObjects) GetVault()(VaultVaultable) {
    return m.vault
}
// Serialize serializes information the current object
func (m *VaultDeletedVaultHolder_additionalObjects) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("audit", m.GetAudit())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *VaultDeletedVaultHolder_additionalObjects) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAudit sets the audit property value. The audit property
func (m *VaultDeletedVaultHolder_additionalObjects) SetAudit(value AuditInfoable)() {
    m.audit = value
}
// SetVault sets the vault property value. The vault property
func (m *VaultDeletedVaultHolder_additionalObjects) SetVault(value VaultVaultable)() {
    m.vault = value
}
// VaultDeletedVaultHolder_additionalObjectsable 
type VaultDeletedVaultHolder_additionalObjectsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAudit()(AuditInfoable)
    GetVault()(VaultVaultable)
    SetAudit(value AuditInfoable)()
    SetVault(value VaultVaultable)()
}
