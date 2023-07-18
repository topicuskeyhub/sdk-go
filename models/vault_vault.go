package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VaultVault 
type VaultVault struct {
    Linkable
    // The accessAvailable property
    accessAvailable *bool
    // The name property
    name *string
    // The records property
    records []VaultVaultRecordable
}
// NewVaultVault instantiates a new vaultVault and sets the default values.
func NewVaultVault()(*VaultVault) {
    m := &VaultVault{
        Linkable: *NewLinkable(),
    }
    typeEscapedValue := "vault.Vault"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateVaultVaultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVaultVaultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVaultVault(), nil
}
// GetAccessAvailable gets the accessAvailable property value. The accessAvailable property
func (m *VaultVault) GetAccessAvailable()(*bool) {
    return m.accessAvailable
}
// GetFieldDeserializers the deserialization information for the current model
func (m *VaultVault) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Linkable.GetFieldDeserializers()
    res["accessAvailable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessAvailable(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["records"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateVaultVaultRecordFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]VaultVaultRecordable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(VaultVaultRecordable)
                }
            }
            m.SetRecords(res)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name property
func (m *VaultVault) GetName()(*string) {
    return m.name
}
// GetRecords gets the records property value. The records property
func (m *VaultVault) GetRecords()([]VaultVaultRecordable) {
    return m.records
}
// Serialize serializes information the current object
func (m *VaultVault) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Linkable.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("accessAvailable", m.GetAccessAvailable())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    if m.GetRecords() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRecords()))
        for i, v := range m.GetRecords() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("records", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessAvailable sets the accessAvailable property value. The accessAvailable property
func (m *VaultVault) SetAccessAvailable(value *bool)() {
    m.accessAvailable = value
}
// SetName sets the name property value. The name property
func (m *VaultVault) SetName(value *string)() {
    m.name = value
}
// SetRecords sets the records property value. The records property
func (m *VaultVault) SetRecords(value []VaultVaultRecordable)() {
    m.records = value
}
// VaultVaultable 
type VaultVaultable interface {
    Linkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessAvailable()(*bool)
    GetName()(*string)
    GetRecords()([]VaultVaultRecordable)
    SetAccessAvailable(value *bool)()
    SetName(value *string)()
    SetRecords(value []VaultVaultRecordable)()
}
