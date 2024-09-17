package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type IdentityIdentity struct {
    Linkable
    // The firstName property
    firstName *string
    // The lastName property
    lastName *string
    // The telephone property
    telephone *string
}
// NewIdentityIdentity instantiates a new IdentityIdentity and sets the default values.
func NewIdentityIdentity()(*IdentityIdentity) {
    m := &IdentityIdentity{
        Linkable: *NewLinkable(),
    }
    typeEscapedValue := "identity.Identity"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateIdentityIdentityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIdentityIdentityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIdentityIdentity(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *IdentityIdentity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Linkable.GetFieldDeserializers()
    res["firstName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirstName(val)
        }
        return nil
    }
    res["lastName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastName(val)
        }
        return nil
    }
    res["telephone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTelephone(val)
        }
        return nil
    }
    return res
}
// GetFirstName gets the firstName property value. The firstName property
// returns a *string when successful
func (m *IdentityIdentity) GetFirstName()(*string) {
    return m.firstName
}
// GetLastName gets the lastName property value. The lastName property
// returns a *string when successful
func (m *IdentityIdentity) GetLastName()(*string) {
    return m.lastName
}
// GetTelephone gets the telephone property value. The telephone property
// returns a *string when successful
func (m *IdentityIdentity) GetTelephone()(*string) {
    return m.telephone
}
// Serialize serializes information the current object
func (m *IdentityIdentity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Linkable.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("firstName", m.GetFirstName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("lastName", m.GetLastName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("telephone", m.GetTelephone())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFirstName sets the firstName property value. The firstName property
func (m *IdentityIdentity) SetFirstName(value *string)() {
    m.firstName = value
}
// SetLastName sets the lastName property value. The lastName property
func (m *IdentityIdentity) SetLastName(value *string)() {
    m.lastName = value
}
// SetTelephone sets the telephone property value. The telephone property
func (m *IdentityIdentity) SetTelephone(value *string)() {
    m.telephone = value
}
type IdentityIdentityable interface {
    Linkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFirstName()(*string)
    GetLastName()(*string)
    GetTelephone()(*string)
    SetFirstName(value *string)()
    SetLastName(value *string)()
    SetTelephone(value *string)()
}
