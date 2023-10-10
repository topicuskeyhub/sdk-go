package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VaultPasswordMetadata 
type VaultPasswordMetadata struct {
    NonLinkable
    // The dictionary property
    dictionary *bool
    // The duplicate property
    duplicate *bool
    // The hash property
    hash *string
    // The length property
    length *int32
    // The lowerCount property
    lowerCount *int32
    // The numberCount property
    numberCount *int32
    // The specialCount property
    specialCount *int32
    // The strength property
    strength *int32
    // The upperCount property
    upperCount *int32
}
// NewVaultPasswordMetadata instantiates a new vaultPasswordMetadata and sets the default values.
func NewVaultPasswordMetadata()(*VaultPasswordMetadata) {
    m := &VaultPasswordMetadata{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "vault.PasswordMetadata"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateVaultPasswordMetadataFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVaultPasswordMetadataFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVaultPasswordMetadata(), nil
}
// GetDictionary gets the dictionary property value. The dictionary property
func (m *VaultPasswordMetadata) GetDictionary()(*bool) {
    return m.dictionary
}
// GetDuplicate gets the duplicate property value. The duplicate property
func (m *VaultPasswordMetadata) GetDuplicate()(*bool) {
    return m.duplicate
}
// GetFieldDeserializers the deserialization information for the current model
func (m *VaultPasswordMetadata) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
    res["dictionary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDictionary(val)
        }
        return nil
    }
    res["duplicate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDuplicate(val)
        }
        return nil
    }
    res["hash"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHash(val)
        }
        return nil
    }
    res["length"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLength(val)
        }
        return nil
    }
    res["lowerCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLowerCount(val)
        }
        return nil
    }
    res["numberCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberCount(val)
        }
        return nil
    }
    res["specialCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSpecialCount(val)
        }
        return nil
    }
    res["strength"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStrength(val)
        }
        return nil
    }
    res["upperCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpperCount(val)
        }
        return nil
    }
    return res
}
// GetHash gets the hash property value. The hash property
func (m *VaultPasswordMetadata) GetHash()(*string) {
    return m.hash
}
// GetLength gets the length property value. The length property
func (m *VaultPasswordMetadata) GetLength()(*int32) {
    return m.length
}
// GetLowerCount gets the lowerCount property value. The lowerCount property
func (m *VaultPasswordMetadata) GetLowerCount()(*int32) {
    return m.lowerCount
}
// GetNumberCount gets the numberCount property value. The numberCount property
func (m *VaultPasswordMetadata) GetNumberCount()(*int32) {
    return m.numberCount
}
// GetSpecialCount gets the specialCount property value. The specialCount property
func (m *VaultPasswordMetadata) GetSpecialCount()(*int32) {
    return m.specialCount
}
// GetStrength gets the strength property value. The strength property
func (m *VaultPasswordMetadata) GetStrength()(*int32) {
    return m.strength
}
// GetUpperCount gets the upperCount property value. The upperCount property
func (m *VaultPasswordMetadata) GetUpperCount()(*int32) {
    return m.upperCount
}
// Serialize serializes information the current object
func (m *VaultPasswordMetadata) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("dictionary", m.GetDictionary())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("duplicate", m.GetDuplicate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("hash", m.GetHash())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("length", m.GetLength())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("lowerCount", m.GetLowerCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("numberCount", m.GetNumberCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("specialCount", m.GetSpecialCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("strength", m.GetStrength())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("upperCount", m.GetUpperCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDictionary sets the dictionary property value. The dictionary property
func (m *VaultPasswordMetadata) SetDictionary(value *bool)() {
    m.dictionary = value
}
// SetDuplicate sets the duplicate property value. The duplicate property
func (m *VaultPasswordMetadata) SetDuplicate(value *bool)() {
    m.duplicate = value
}
// SetHash sets the hash property value. The hash property
func (m *VaultPasswordMetadata) SetHash(value *string)() {
    m.hash = value
}
// SetLength sets the length property value. The length property
func (m *VaultPasswordMetadata) SetLength(value *int32)() {
    m.length = value
}
// SetLowerCount sets the lowerCount property value. The lowerCount property
func (m *VaultPasswordMetadata) SetLowerCount(value *int32)() {
    m.lowerCount = value
}
// SetNumberCount sets the numberCount property value. The numberCount property
func (m *VaultPasswordMetadata) SetNumberCount(value *int32)() {
    m.numberCount = value
}
// SetSpecialCount sets the specialCount property value. The specialCount property
func (m *VaultPasswordMetadata) SetSpecialCount(value *int32)() {
    m.specialCount = value
}
// SetStrength sets the strength property value. The strength property
func (m *VaultPasswordMetadata) SetStrength(value *int32)() {
    m.strength = value
}
// SetUpperCount sets the upperCount property value. The upperCount property
func (m *VaultPasswordMetadata) SetUpperCount(value *int32)() {
    m.upperCount = value
}
// VaultPasswordMetadataable 
type VaultPasswordMetadataable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDictionary()(*bool)
    GetDuplicate()(*bool)
    GetHash()(*string)
    GetLength()(*int32)
    GetLowerCount()(*int32)
    GetNumberCount()(*int32)
    GetSpecialCount()(*int32)
    GetStrength()(*int32)
    GetUpperCount()(*int32)
    SetDictionary(value *bool)()
    SetDuplicate(value *bool)()
    SetHash(value *string)()
    SetLength(value *int32)()
    SetLowerCount(value *int32)()
    SetNumberCount(value *int32)()
    SetSpecialCount(value *int32)()
    SetStrength(value *int32)()
    SetUpperCount(value *int32)()
}
