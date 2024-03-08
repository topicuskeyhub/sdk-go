package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type VaultDeletedVaultHolder struct {
    Linkable
    // The additionalObjects property
    additionalObjects VaultDeletedVaultHolder_additionalObjectsable
    // The name property
    name *string
    // The recordCount property
    recordCount *int32
    // The type property
    vaultDeletedVaultHolderType *VaultVaultHolderType
}
// NewVaultDeletedVaultHolder instantiates a new VaultDeletedVaultHolder and sets the default values.
func NewVaultDeletedVaultHolder()(*VaultDeletedVaultHolder) {
    m := &VaultDeletedVaultHolder{
        Linkable: *NewLinkable(),
    }
    typeEscapedValue := "vault.DeletedVaultHolder"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateVaultDeletedVaultHolderFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateVaultDeletedVaultHolderFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVaultDeletedVaultHolder(), nil
}
// GetAdditionalObjects gets the additionalObjects property value. The additionalObjects property
// returns a VaultDeletedVaultHolder_additionalObjectsable when successful
func (m *VaultDeletedVaultHolder) GetAdditionalObjects()(VaultDeletedVaultHolder_additionalObjectsable) {
    return m.additionalObjects
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *VaultDeletedVaultHolder) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Linkable.GetFieldDeserializers()
    res["additionalObjects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateVaultDeletedVaultHolder_additionalObjectsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalObjects(val.(VaultDeletedVaultHolder_additionalObjectsable))
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
    res["recordCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecordCount(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVaultVaultHolderType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVaultDeletedVaultHolderType(val.(*VaultVaultHolderType))
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *VaultDeletedVaultHolder) GetName()(*string) {
    return m.name
}
// GetRecordCount gets the recordCount property value. The recordCount property
// returns a *int32 when successful
func (m *VaultDeletedVaultHolder) GetRecordCount()(*int32) {
    return m.recordCount
}
// GetVaultDeletedVaultHolderType gets the type property value. The type property
// returns a *VaultVaultHolderType when successful
func (m *VaultDeletedVaultHolder) GetVaultDeletedVaultHolderType()(*VaultVaultHolderType) {
    return m.vaultDeletedVaultHolderType
}
// Serialize serializes information the current object
func (m *VaultDeletedVaultHolder) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Linkable.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("additionalObjects", m.GetAdditionalObjects())
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
    if m.GetVaultDeletedVaultHolderType() != nil {
        cast := (*m.GetVaultDeletedVaultHolderType()).String()
        err = writer.WriteStringValue("type", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalObjects sets the additionalObjects property value. The additionalObjects property
func (m *VaultDeletedVaultHolder) SetAdditionalObjects(value VaultDeletedVaultHolder_additionalObjectsable)() {
    m.additionalObjects = value
}
// SetName sets the name property value. The name property
func (m *VaultDeletedVaultHolder) SetName(value *string)() {
    m.name = value
}
// SetRecordCount sets the recordCount property value. The recordCount property
func (m *VaultDeletedVaultHolder) SetRecordCount(value *int32)() {
    m.recordCount = value
}
// SetVaultDeletedVaultHolderType sets the type property value. The type property
func (m *VaultDeletedVaultHolder) SetVaultDeletedVaultHolderType(value *VaultVaultHolderType)() {
    m.vaultDeletedVaultHolderType = value
}
type VaultDeletedVaultHolderable interface {
    Linkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdditionalObjects()(VaultDeletedVaultHolder_additionalObjectsable)
    GetName()(*string)
    GetRecordCount()(*int32)
    GetVaultDeletedVaultHolderType()(*VaultVaultHolderType)
    SetAdditionalObjects(value VaultDeletedVaultHolder_additionalObjectsable)()
    SetName(value *string)()
    SetRecordCount(value *int32)()
    SetVaultDeletedVaultHolderType(value *VaultVaultHolderType)()
}
