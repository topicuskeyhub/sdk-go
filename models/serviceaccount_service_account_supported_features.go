package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ServiceaccountServiceAccountSupportedFeatures struct {
    NonLinkable
    // The sshPublicKey property
    sshPublicKey *bool
}
// NewServiceaccountServiceAccountSupportedFeatures instantiates a new ServiceaccountServiceAccountSupportedFeatures and sets the default values.
func NewServiceaccountServiceAccountSupportedFeatures()(*ServiceaccountServiceAccountSupportedFeatures) {
    m := &ServiceaccountServiceAccountSupportedFeatures{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "serviceaccount.ServiceAccountSupportedFeatures"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateServiceaccountServiceAccountSupportedFeaturesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateServiceaccountServiceAccountSupportedFeaturesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServiceaccountServiceAccountSupportedFeatures(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ServiceaccountServiceAccountSupportedFeatures) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
    res["sshPublicKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSshPublicKey(val)
        }
        return nil
    }
    return res
}
// GetSshPublicKey gets the sshPublicKey property value. The sshPublicKey property
// returns a *bool when successful
func (m *ServiceaccountServiceAccountSupportedFeatures) GetSshPublicKey()(*bool) {
    return m.sshPublicKey
}
// Serialize serializes information the current object
func (m *ServiceaccountServiceAccountSupportedFeatures) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("sshPublicKey", m.GetSshPublicKey())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSshPublicKey sets the sshPublicKey property value. The sshPublicKey property
func (m *ServiceaccountServiceAccountSupportedFeatures) SetSshPublicKey(value *bool)() {
    m.sshPublicKey = value
}
type ServiceaccountServiceAccountSupportedFeaturesable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSshPublicKey()(*bool)
    SetSshPublicKey(value *bool)()
}
