package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProvisioningProvisionedLDAP 
type ProvisioningProvisionedLDAP struct {
    ProvisioningAbstractProvisionedLDAP
    // The gid property
    gid *int64
    // The hashingScheme property
    hashingScheme *ProvisioningLDAPPasswordHashingScheme
    // The numbering property
    numbering ProvisioningProvisionNumberSequenceable
}
// NewProvisioningProvisionedLDAP instantiates a new provisioningProvisionedLDAP and sets the default values.
func NewProvisioningProvisionedLDAP()(*ProvisioningProvisionedLDAP) {
    m := &ProvisioningProvisionedLDAP{
        ProvisioningAbstractProvisionedLDAP: *NewProvisioningAbstractProvisionedLDAP(),
    }
    typeEscapedValue := "provisioning.ProvisionedLDAP"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateProvisioningProvisionedLDAPFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProvisioningProvisionedLDAPFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProvisioningProvisionedLDAP(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProvisioningProvisionedLDAP) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ProvisioningAbstractProvisionedLDAP.GetFieldDeserializers()
    res["gid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGid(val)
        }
        return nil
    }
    res["hashingScheme"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseProvisioningLDAPPasswordHashingScheme)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHashingScheme(val.(*ProvisioningLDAPPasswordHashingScheme))
        }
        return nil
    }
    res["numbering"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProvisioningProvisionNumberSequenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumbering(val.(ProvisioningProvisionNumberSequenceable))
        }
        return nil
    }
    return res
}
// GetGid gets the gid property value. The gid property
func (m *ProvisioningProvisionedLDAP) GetGid()(*int64) {
    return m.gid
}
// GetHashingScheme gets the hashingScheme property value. The hashingScheme property
func (m *ProvisioningProvisionedLDAP) GetHashingScheme()(*ProvisioningLDAPPasswordHashingScheme) {
    return m.hashingScheme
}
// GetNumbering gets the numbering property value. The numbering property
func (m *ProvisioningProvisionedLDAP) GetNumbering()(ProvisioningProvisionNumberSequenceable) {
    return m.numbering
}
// Serialize serializes information the current object
func (m *ProvisioningProvisionedLDAP) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ProvisioningAbstractProvisionedLDAP.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("gid", m.GetGid())
        if err != nil {
            return err
        }
    }
    if m.GetHashingScheme() != nil {
        cast := (*m.GetHashingScheme()).String()
        err = writer.WriteStringValue("hashingScheme", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("numbering", m.GetNumbering())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetGid sets the gid property value. The gid property
func (m *ProvisioningProvisionedLDAP) SetGid(value *int64)() {
    m.gid = value
}
// SetHashingScheme sets the hashingScheme property value. The hashingScheme property
func (m *ProvisioningProvisionedLDAP) SetHashingScheme(value *ProvisioningLDAPPasswordHashingScheme)() {
    m.hashingScheme = value
}
// SetNumbering sets the numbering property value. The numbering property
func (m *ProvisioningProvisionedLDAP) SetNumbering(value ProvisioningProvisionNumberSequenceable)() {
    m.numbering = value
}
// ProvisioningProvisionedLDAPable 
type ProvisioningProvisionedLDAPable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ProvisioningAbstractProvisionedLDAPable
    GetGid()(*int64)
    GetHashingScheme()(*ProvisioningLDAPPasswordHashingScheme)
    GetNumbering()(ProvisioningProvisionNumberSequenceable)
    SetGid(value *int64)()
    SetHashingScheme(value *ProvisioningLDAPPasswordHashingScheme)()
    SetNumbering(value ProvisioningProvisionNumberSequenceable)()
}
