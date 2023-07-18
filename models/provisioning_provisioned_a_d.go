package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProvisioningProvisionedAD 
type ProvisioningProvisionedAD struct {
    ProvisioningAbstractProvisionedLDAP
    // The samAccountNameScheme property
    samAccountNameScheme *ProvisioningADSamAccountNameScheme
}
// NewProvisioningProvisionedAD instantiates a new provisioningProvisionedAD and sets the default values.
func NewProvisioningProvisionedAD()(*ProvisioningProvisionedAD) {
    m := &ProvisioningProvisionedAD{
        ProvisioningAbstractProvisionedLDAP: *NewProvisioningAbstractProvisionedLDAP(),
    }
    typeEscapedValue := "provisioning.ProvisionedAD"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateProvisioningProvisionedADFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProvisioningProvisionedADFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProvisioningProvisionedAD(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProvisioningProvisionedAD) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ProvisioningAbstractProvisionedLDAP.GetFieldDeserializers()
    res["samAccountNameScheme"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseProvisioningADSamAccountNameScheme)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSamAccountNameScheme(val.(*ProvisioningADSamAccountNameScheme))
        }
        return nil
    }
    return res
}
// GetSamAccountNameScheme gets the samAccountNameScheme property value. The samAccountNameScheme property
func (m *ProvisioningProvisionedAD) GetSamAccountNameScheme()(*ProvisioningADSamAccountNameScheme) {
    return m.samAccountNameScheme
}
// Serialize serializes information the current object
func (m *ProvisioningProvisionedAD) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ProvisioningAbstractProvisionedLDAP.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetSamAccountNameScheme() != nil {
        cast := (*m.GetSamAccountNameScheme()).String()
        err = writer.WriteStringValue("samAccountNameScheme", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSamAccountNameScheme sets the samAccountNameScheme property value. The samAccountNameScheme property
func (m *ProvisioningProvisionedAD) SetSamAccountNameScheme(value *ProvisioningADSamAccountNameScheme)() {
    m.samAccountNameScheme = value
}
// ProvisioningProvisionedADable 
type ProvisioningProvisionedADable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ProvisioningAbstractProvisionedLDAPable
    GetSamAccountNameScheme()(*ProvisioningADSamAccountNameScheme)
    SetSamAccountNameScheme(value *ProvisioningADSamAccountNameScheme)()
}
