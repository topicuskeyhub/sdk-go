// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type VaultVaultUnlockResponse struct {
    NonLinkable
    // The expiresAt property
    expiresAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The sessionPassword property
    sessionPassword *string
}
// NewVaultVaultUnlockResponse instantiates a new VaultVaultUnlockResponse and sets the default values.
func NewVaultVaultUnlockResponse()(*VaultVaultUnlockResponse) {
    m := &VaultVaultUnlockResponse{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "vault.VaultUnlockResponse"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateVaultVaultUnlockResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateVaultVaultUnlockResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVaultVaultUnlockResponse(), nil
}
// GetExpiresAt gets the expiresAt property value. The expiresAt property
// returns a *Time when successful
func (m *VaultVaultUnlockResponse) GetExpiresAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.expiresAt
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *VaultVaultUnlockResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
    res["expiresAt"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpiresAt(val)
        }
        return nil
    }
    res["sessionPassword"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSessionPassword(val)
        }
        return nil
    }
    return res
}
// GetSessionPassword gets the sessionPassword property value. The sessionPassword property
// returns a *string when successful
func (m *VaultVaultUnlockResponse) GetSessionPassword()(*string) {
    return m.sessionPassword
}
// Serialize serializes information the current object
func (m *VaultVaultUnlockResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("expiresAt", m.GetExpiresAt())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("sessionPassword", m.GetSessionPassword())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetExpiresAt sets the expiresAt property value. The expiresAt property
func (m *VaultVaultUnlockResponse) SetExpiresAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expiresAt = value
}
// SetSessionPassword sets the sessionPassword property value. The sessionPassword property
func (m *VaultVaultUnlockResponse) SetSessionPassword(value *string)() {
    m.sessionPassword = value
}
type VaultVaultUnlockResponseable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExpiresAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSessionPassword()(*string)
    SetExpiresAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSessionPassword(value *string)()
}
