// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AuthWebAuthnKey struct {
    Linkable
    // The displayName property
    displayName *string
    // The keyId property
    keyId *string
    // The prfSupport property
    prfSupport *AuthWebAuthnPRFSupported
    // The publicKey property
    publicKey *string
}
// NewAuthWebAuthnKey instantiates a new AuthWebAuthnKey and sets the default values.
func NewAuthWebAuthnKey()(*AuthWebAuthnKey) {
    m := &AuthWebAuthnKey{
        Linkable: *NewLinkable(),
    }
    typeEscapedValue := "auth.WebAuthnKey"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateAuthWebAuthnKeyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAuthWebAuthnKeyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthWebAuthnKey(), nil
}
// GetDisplayName gets the displayName property value. The displayName property
// returns a *string when successful
func (m *AuthWebAuthnKey) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AuthWebAuthnKey) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["keyId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyId(val)
        }
        return nil
    }
    res["prfSupport"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAuthWebAuthnPRFSupported)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrfSupport(val.(*AuthWebAuthnPRFSupported))
        }
        return nil
    }
    res["publicKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublicKey(val)
        }
        return nil
    }
    return res
}
// GetKeyId gets the keyId property value. The keyId property
// returns a *string when successful
func (m *AuthWebAuthnKey) GetKeyId()(*string) {
    return m.keyId
}
// GetPrfSupport gets the prfSupport property value. The prfSupport property
// returns a *AuthWebAuthnPRFSupported when successful
func (m *AuthWebAuthnKey) GetPrfSupport()(*AuthWebAuthnPRFSupported) {
    return m.prfSupport
}
// GetPublicKey gets the publicKey property value. The publicKey property
// returns a *string when successful
func (m *AuthWebAuthnKey) GetPublicKey()(*string) {
    return m.publicKey
}
// Serialize serializes information the current object
func (m *AuthWebAuthnKey) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetPrfSupport() != nil {
        cast := (*m.GetPrfSupport()).String()
        err = writer.WriteStringValue("prfSupport", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *AuthWebAuthnKey) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetKeyId sets the keyId property value. The keyId property
func (m *AuthWebAuthnKey) SetKeyId(value *string)() {
    m.keyId = value
}
// SetPrfSupport sets the prfSupport property value. The prfSupport property
func (m *AuthWebAuthnKey) SetPrfSupport(value *AuthWebAuthnPRFSupported)() {
    m.prfSupport = value
}
// SetPublicKey sets the publicKey property value. The publicKey property
func (m *AuthWebAuthnKey) SetPublicKey(value *string)() {
    m.publicKey = value
}
type AuthWebAuthnKeyable interface {
    Linkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetKeyId()(*string)
    GetPrfSupport()(*AuthWebAuthnPRFSupported)
    GetPublicKey()(*string)
    SetDisplayName(value *string)()
    SetKeyId(value *string)()
    SetPrfSupport(value *AuthWebAuthnPRFSupported)()
    SetPublicKey(value *string)()
}
