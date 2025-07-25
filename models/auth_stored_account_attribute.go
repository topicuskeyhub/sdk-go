// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AuthStoredAccountAttribute struct {
    NonLinkable
    // The name property
    name *string
    // The value property
    value *string
}
// NewAuthStoredAccountAttribute instantiates a new AuthStoredAccountAttribute and sets the default values.
func NewAuthStoredAccountAttribute()(*AuthStoredAccountAttribute) {
    m := &AuthStoredAccountAttribute{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "auth.StoredAccountAttribute"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateAuthStoredAccountAttributeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAuthStoredAccountAttributeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthStoredAccountAttribute(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AuthStoredAccountAttribute) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *AuthStoredAccountAttribute) GetName()(*string) {
    return m.name
}
// GetValue gets the value property value. The value property
// returns a *string when successful
func (m *AuthStoredAccountAttribute) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *AuthStoredAccountAttribute) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err = writer.WriteStringValue("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetName sets the name property value. The name property
func (m *AuthStoredAccountAttribute) SetName(value *string)() {
    m.name = value
}
// SetValue sets the value property value. The value property
func (m *AuthStoredAccountAttribute) SetValue(value *string)() {
    m.value = value
}
type AuthStoredAccountAttributeable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetName()(*string)
    GetValue()(*string)
    SetName(value *string)()
    SetValue(value *string)()
}
