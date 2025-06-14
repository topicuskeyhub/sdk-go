// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AuthInternalAccountStatusObject struct {
    NonLinkable
    // The status property
    status *AuthInternalAccountStatus
}
// NewAuthInternalAccountStatusObject instantiates a new AuthInternalAccountStatusObject and sets the default values.
func NewAuthInternalAccountStatusObject()(*AuthInternalAccountStatusObject) {
    m := &AuthInternalAccountStatusObject{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "auth.InternalAccountStatusObject"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateAuthInternalAccountStatusObjectFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAuthInternalAccountStatusObjectFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthInternalAccountStatusObject(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AuthInternalAccountStatusObject) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAuthInternalAccountStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*AuthInternalAccountStatus))
        }
        return nil
    }
    return res
}
// GetStatus gets the status property value. The status property
// returns a *AuthInternalAccountStatus when successful
func (m *AuthInternalAccountStatusObject) GetStatus()(*AuthInternalAccountStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *AuthInternalAccountStatusObject) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetStatus sets the status property value. The status property
func (m *AuthInternalAccountStatusObject) SetStatus(value *AuthInternalAccountStatus)() {
    m.status = value
}
type AuthInternalAccountStatusObjectable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetStatus()(*AuthInternalAccountStatus)
    SetStatus(value *AuthInternalAccountStatus)()
}
