package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProvisioningTokenPassword 
type ProvisioningTokenPassword struct {
    NonLinkable
    // The password property
    password *string
}
// NewProvisioningTokenPassword instantiates a new provisioningTokenPassword and sets the default values.
func NewProvisioningTokenPassword()(*ProvisioningTokenPassword) {
    m := &ProvisioningTokenPassword{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "provisioning.TokenPassword"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateProvisioningTokenPasswordFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProvisioningTokenPasswordFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProvisioningTokenPassword(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProvisioningTokenPassword) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
    res["password"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPassword(val)
        }
        return nil
    }
    return res
}
// GetPassword gets the password property value. The password property
func (m *ProvisioningTokenPassword) GetPassword()(*string) {
    return m.password
}
// Serialize serializes information the current object
func (m *ProvisioningTokenPassword) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("password", m.GetPassword())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetPassword sets the password property value. The password property
func (m *ProvisioningTokenPassword) SetPassword(value *string)() {
    m.password = value
}
// ProvisioningTokenPasswordable 
type ProvisioningTokenPasswordable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPassword()(*string)
    SetPassword(value *string)()
}
