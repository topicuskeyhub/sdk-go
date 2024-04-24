package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type VaultVaultActivationStatus struct {
    NonLinkable
    // The activated property
    activated *bool
    // The activationRequired property
    activationRequired *bool
}
// NewVaultVaultActivationStatus instantiates a new VaultVaultActivationStatus and sets the default values.
func NewVaultVaultActivationStatus()(*VaultVaultActivationStatus) {
    m := &VaultVaultActivationStatus{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "vault.VaultActivationStatus"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateVaultVaultActivationStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateVaultVaultActivationStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVaultVaultActivationStatus(), nil
}
// GetActivated gets the activated property value. The activated property
// returns a *bool when successful
func (m *VaultVaultActivationStatus) GetActivated()(*bool) {
    return m.activated
}
// GetActivationRequired gets the activationRequired property value. The activationRequired property
// returns a *bool when successful
func (m *VaultVaultActivationStatus) GetActivationRequired()(*bool) {
    return m.activationRequired
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *VaultVaultActivationStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
    res["activated"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivated(val)
        }
        return nil
    }
    res["activationRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivationRequired(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *VaultVaultActivationStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("activated", m.GetActivated())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("activationRequired", m.GetActivationRequired())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActivated sets the activated property value. The activated property
func (m *VaultVaultActivationStatus) SetActivated(value *bool)() {
    m.activated = value
}
// SetActivationRequired sets the activationRequired property value. The activationRequired property
func (m *VaultVaultActivationStatus) SetActivationRequired(value *bool)() {
    m.activationRequired = value
}
type VaultVaultActivationStatusable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActivated()(*bool)
    GetActivationRequired()(*bool)
    SetActivated(value *bool)()
    SetActivationRequired(value *bool)()
}
