package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ServiceaccountServiceAccountStatus struct {
    NonLinkable
    // The accountEnabled property
    accountEnabled *bool
}
// NewServiceaccountServiceAccountStatus instantiates a new ServiceaccountServiceAccountStatus and sets the default values.
func NewServiceaccountServiceAccountStatus()(*ServiceaccountServiceAccountStatus) {
    m := &ServiceaccountServiceAccountStatus{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "serviceaccount.ServiceAccountStatus"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateServiceaccountServiceAccountStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateServiceaccountServiceAccountStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServiceaccountServiceAccountStatus(), nil
}
// GetAccountEnabled gets the accountEnabled property value. The accountEnabled property
// returns a *bool when successful
func (m *ServiceaccountServiceAccountStatus) GetAccountEnabled()(*bool) {
    return m.accountEnabled
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ServiceaccountServiceAccountStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
    res["accountEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountEnabled(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ServiceaccountServiceAccountStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("accountEnabled", m.GetAccountEnabled())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccountEnabled sets the accountEnabled property value. The accountEnabled property
func (m *ServiceaccountServiceAccountStatus) SetAccountEnabled(value *bool)() {
    m.accountEnabled = value
}
type ServiceaccountServiceAccountStatusable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccountEnabled()(*bool)
    SetAccountEnabled(value *bool)()
}
