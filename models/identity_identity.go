package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type IdentityIdentity struct {
    Linkable
    // The displayName property
    displayName *string
    // The familyName property
    familyName *string
    // The givenName property
    givenName *string
    // The middleName property
    middleName *string
    // The privateEmail property
    privateEmail *string
    // The privateTelephone property
    privateTelephone *string
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
// GetDisplayName gets the displayName property value. The displayName property
// returns a *string when successful
func (m *IdentityIdentity) GetDisplayName()(*string) {
    return m.displayName
}
// GetFamilyName gets the familyName property value. The familyName property
// returns a *string when successful
func (m *IdentityIdentity) GetFamilyName()(*string) {
    return m.familyName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *IdentityIdentity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Linkable.GetFieldDeserializers()
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["familyName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFamilyName(val)
        }
        return nil
    }
    res["givenName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGivenName(val)
        }
        return nil
    }
    res["middleName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMiddleName(val)
        }
        return nil
    }
    res["privateEmail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivateEmail(val)
        }
        return nil
    }
    res["privateTelephone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivateTelephone(val)
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
// GetGivenName gets the givenName property value. The givenName property
// returns a *string when successful
func (m *IdentityIdentity) GetGivenName()(*string) {
    return m.givenName
}
// GetMiddleName gets the middleName property value. The middleName property
// returns a *string when successful
func (m *IdentityIdentity) GetMiddleName()(*string) {
    return m.middleName
}
// GetPrivateEmail gets the privateEmail property value. The privateEmail property
// returns a *string when successful
func (m *IdentityIdentity) GetPrivateEmail()(*string) {
    return m.privateEmail
}
// GetPrivateTelephone gets the privateTelephone property value. The privateTelephone property
// returns a *string when successful
func (m *IdentityIdentity) GetPrivateTelephone()(*string) {
    return m.privateTelephone
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
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("familyName", m.GetFamilyName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("givenName", m.GetGivenName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("middleName", m.GetMiddleName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("privateEmail", m.GetPrivateEmail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("privateTelephone", m.GetPrivateTelephone())
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
// SetDisplayName sets the displayName property value. The displayName property
func (m *IdentityIdentity) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetFamilyName sets the familyName property value. The familyName property
func (m *IdentityIdentity) SetFamilyName(value *string)() {
    m.familyName = value
}
// SetGivenName sets the givenName property value. The givenName property
func (m *IdentityIdentity) SetGivenName(value *string)() {
    m.givenName = value
}
// SetMiddleName sets the middleName property value. The middleName property
func (m *IdentityIdentity) SetMiddleName(value *string)() {
    m.middleName = value
}
// SetPrivateEmail sets the privateEmail property value. The privateEmail property
func (m *IdentityIdentity) SetPrivateEmail(value *string)() {
    m.privateEmail = value
}
// SetPrivateTelephone sets the privateTelephone property value. The privateTelephone property
func (m *IdentityIdentity) SetPrivateTelephone(value *string)() {
    m.privateTelephone = value
}
// SetTelephone sets the telephone property value. The telephone property
func (m *IdentityIdentity) SetTelephone(value *string)() {
    m.telephone = value
}
type IdentityIdentityable interface {
    Linkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetFamilyName()(*string)
    GetGivenName()(*string)
    GetMiddleName()(*string)
    GetPrivateEmail()(*string)
    GetPrivateTelephone()(*string)
    GetTelephone()(*string)
    SetDisplayName(value *string)()
    SetFamilyName(value *string)()
    SetGivenName(value *string)()
    SetMiddleName(value *string)()
    SetPrivateEmail(value *string)()
    SetPrivateTelephone(value *string)()
    SetTelephone(value *string)()
}
