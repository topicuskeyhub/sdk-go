package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthAccountChangeLocale 
type AuthAccountChangeLocale struct {
    NonLinkable
    // The locale property
    locale *string
}
// NewAuthAccountChangeLocale instantiates a new authAccountChangeLocale and sets the default values.
func NewAuthAccountChangeLocale()(*AuthAccountChangeLocale) {
    m := &AuthAccountChangeLocale{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "auth.AccountChangeLocale"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateAuthAccountChangeLocaleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthAccountChangeLocaleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthAccountChangeLocale(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthAccountChangeLocale) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
    res["locale"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocale(val)
        }
        return nil
    }
    return res
}
// GetLocale gets the locale property value. The locale property
func (m *AuthAccountChangeLocale) GetLocale()(*string) {
    return m.locale
}
// Serialize serializes information the current object
func (m *AuthAccountChangeLocale) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("locale", m.GetLocale())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLocale sets the locale property value. The locale property
func (m *AuthAccountChangeLocale) SetLocale(value *string)() {
    m.locale = value
}
// AuthAccountChangeLocaleable 
type AuthAccountChangeLocaleable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLocale()(*string)
    SetLocale(value *string)()
}
