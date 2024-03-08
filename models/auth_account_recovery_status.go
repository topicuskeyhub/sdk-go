package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AuthAccountRecoveryStatus struct {
    NonLinkable
    // The pending2FARecoveryRequest property
    pending2FARecoveryRequest *bool
    // The pendingPasswordRecoveryRequest property
    pendingPasswordRecoveryRequest *bool
}
// NewAuthAccountRecoveryStatus instantiates a new AuthAccountRecoveryStatus and sets the default values.
func NewAuthAccountRecoveryStatus()(*AuthAccountRecoveryStatus) {
    m := &AuthAccountRecoveryStatus{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "auth.AccountRecoveryStatus"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateAuthAccountRecoveryStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAuthAccountRecoveryStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthAccountRecoveryStatus(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AuthAccountRecoveryStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
    res["pending2FARecoveryRequest"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPending2FARecoveryRequest(val)
        }
        return nil
    }
    res["pendingPasswordRecoveryRequest"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPendingPasswordRecoveryRequest(val)
        }
        return nil
    }
    return res
}
// GetPending2FARecoveryRequest gets the pending2FARecoveryRequest property value. The pending2FARecoveryRequest property
// returns a *bool when successful
func (m *AuthAccountRecoveryStatus) GetPending2FARecoveryRequest()(*bool) {
    return m.pending2FARecoveryRequest
}
// GetPendingPasswordRecoveryRequest gets the pendingPasswordRecoveryRequest property value. The pendingPasswordRecoveryRequest property
// returns a *bool when successful
func (m *AuthAccountRecoveryStatus) GetPendingPasswordRecoveryRequest()(*bool) {
    return m.pendingPasswordRecoveryRequest
}
// Serialize serializes information the current object
func (m *AuthAccountRecoveryStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("pending2FARecoveryRequest", m.GetPending2FARecoveryRequest())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("pendingPasswordRecoveryRequest", m.GetPendingPasswordRecoveryRequest())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetPending2FARecoveryRequest sets the pending2FARecoveryRequest property value. The pending2FARecoveryRequest property
func (m *AuthAccountRecoveryStatus) SetPending2FARecoveryRequest(value *bool)() {
    m.pending2FARecoveryRequest = value
}
// SetPendingPasswordRecoveryRequest sets the pendingPasswordRecoveryRequest property value. The pendingPasswordRecoveryRequest property
func (m *AuthAccountRecoveryStatus) SetPendingPasswordRecoveryRequest(value *bool)() {
    m.pendingPasswordRecoveryRequest = value
}
type AuthAccountRecoveryStatusable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPending2FARecoveryRequest()(*bool)
    GetPendingPasswordRecoveryRequest()(*bool)
    SetPending2FARecoveryRequest(value *bool)()
    SetPendingPasswordRecoveryRequest(value *bool)()
}
