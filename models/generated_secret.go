package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GeneratedSecret 
type GeneratedSecret struct {
    NonLinkable
    // The generatedSecret property
    generatedSecret *string
    // The oldSecret property
    oldSecret *string
    // The regenerate property
    regenerate *bool
}
// NewGeneratedSecret instantiates a new GeneratedSecret and sets the default values.
func NewGeneratedSecret()(*GeneratedSecret) {
    m := &GeneratedSecret{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "GeneratedSecret"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateGeneratedSecretFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGeneratedSecretFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGeneratedSecret(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GeneratedSecret) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
    res["generatedSecret"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGeneratedSecret(val)
        }
        return nil
    }
    res["oldSecret"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOldSecret(val)
        }
        return nil
    }
    res["regenerate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegenerate(val)
        }
        return nil
    }
    return res
}
// GetGeneratedSecret gets the generatedSecret property value. The generatedSecret property
func (m *GeneratedSecret) GetGeneratedSecret()(*string) {
    return m.generatedSecret
}
// GetOldSecret gets the oldSecret property value. The oldSecret property
func (m *GeneratedSecret) GetOldSecret()(*string) {
    return m.oldSecret
}
// GetRegenerate gets the regenerate property value. The regenerate property
func (m *GeneratedSecret) GetRegenerate()(*bool) {
    return m.regenerate
}
// Serialize serializes information the current object
func (m *GeneratedSecret) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("generatedSecret", m.GetGeneratedSecret())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("oldSecret", m.GetOldSecret())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("regenerate", m.GetRegenerate())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetGeneratedSecret sets the generatedSecret property value. The generatedSecret property
func (m *GeneratedSecret) SetGeneratedSecret(value *string)() {
    m.generatedSecret = value
}
// SetOldSecret sets the oldSecret property value. The oldSecret property
func (m *GeneratedSecret) SetOldSecret(value *string)() {
    m.oldSecret = value
}
// SetRegenerate sets the regenerate property value. The regenerate property
func (m *GeneratedSecret) SetRegenerate(value *bool)() {
    m.regenerate = value
}
// GeneratedSecretable 
type GeneratedSecretable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetGeneratedSecret()(*string)
    GetOldSecret()(*string)
    GetRegenerate()(*bool)
    SetGeneratedSecret(value *string)()
    SetOldSecret(value *string)()
    SetRegenerate(value *bool)()
}
