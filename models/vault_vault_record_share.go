package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VaultVaultRecordShare 
type VaultVaultRecordShare struct {
    NonLinkable
    // The name property
    name *string
    // The type property
    vaultVaultRecordShareType *VaultVaultHolderType
}
// NewVaultVaultRecordShare instantiates a new vaultVaultRecordShare and sets the default values.
func NewVaultVaultRecordShare()(*VaultVaultRecordShare) {
    m := &VaultVaultRecordShare{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "vault.VaultRecordShare"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateVaultVaultRecordShareFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVaultVaultRecordShareFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVaultVaultRecordShare(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *VaultVaultRecordShare) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
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
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVaultVaultHolderType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVaultVaultRecordShareType(val.(*VaultVaultHolderType))
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name property
func (m *VaultVaultRecordShare) GetName()(*string) {
    return m.name
}
// GetVaultVaultRecordShareType gets the type property value. The type property
func (m *VaultVaultRecordShare) GetVaultVaultRecordShareType()(*VaultVaultHolderType) {
    return m.vaultVaultRecordShareType
}
// Serialize serializes information the current object
func (m *VaultVaultRecordShare) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    if m.GetVaultVaultRecordShareType() != nil {
        cast := (*m.GetVaultVaultRecordShareType()).String()
        err = writer.WriteStringValue("type", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetName sets the name property value. The name property
func (m *VaultVaultRecordShare) SetName(value *string)() {
    m.name = value
}
// SetVaultVaultRecordShareType sets the type property value. The type property
func (m *VaultVaultRecordShare) SetVaultVaultRecordShareType(value *VaultVaultHolderType)() {
    m.vaultVaultRecordShareType = value
}
// VaultVaultRecordShareable 
type VaultVaultRecordShareable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetName()(*string)
    GetVaultVaultRecordShareType()(*VaultVaultHolderType)
    SetName(value *string)()
    SetVaultVaultRecordShareType(value *VaultVaultHolderType)()
}
