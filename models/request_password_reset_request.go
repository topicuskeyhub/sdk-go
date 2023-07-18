package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RequestPasswordResetRequest 
type RequestPasswordResetRequest struct {
    RequestModificationRequest
    // The newVaultAndDirectoryPassword property
    newVaultAndDirectoryPassword *string
    // The newVaultPassword property
    newVaultPassword *string
    // The unsyncedPassword property
    unsyncedPassword *bool
}
// NewRequestPasswordResetRequest instantiates a new requestPasswordResetRequest and sets the default values.
func NewRequestPasswordResetRequest()(*RequestPasswordResetRequest) {
    m := &RequestPasswordResetRequest{
        RequestModificationRequest: *NewRequestModificationRequest(),
    }
    typeEscapedValue := "request.PasswordResetRequest"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateRequestPasswordResetRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRequestPasswordResetRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequestPasswordResetRequest(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RequestPasswordResetRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RequestModificationRequest.GetFieldDeserializers()
    res["newVaultAndDirectoryPassword"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNewVaultAndDirectoryPassword(val)
        }
        return nil
    }
    res["newVaultPassword"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNewVaultPassword(val)
        }
        return nil
    }
    res["unsyncedPassword"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnsyncedPassword(val)
        }
        return nil
    }
    return res
}
// GetNewVaultAndDirectoryPassword gets the newVaultAndDirectoryPassword property value. The newVaultAndDirectoryPassword property
func (m *RequestPasswordResetRequest) GetNewVaultAndDirectoryPassword()(*string) {
    return m.newVaultAndDirectoryPassword
}
// GetNewVaultPassword gets the newVaultPassword property value. The newVaultPassword property
func (m *RequestPasswordResetRequest) GetNewVaultPassword()(*string) {
    return m.newVaultPassword
}
// GetUnsyncedPassword gets the unsyncedPassword property value. The unsyncedPassword property
func (m *RequestPasswordResetRequest) GetUnsyncedPassword()(*bool) {
    return m.unsyncedPassword
}
// Serialize serializes information the current object
func (m *RequestPasswordResetRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RequestModificationRequest.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("newVaultAndDirectoryPassword", m.GetNewVaultAndDirectoryPassword())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("newVaultPassword", m.GetNewVaultPassword())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("unsyncedPassword", m.GetUnsyncedPassword())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetNewVaultAndDirectoryPassword sets the newVaultAndDirectoryPassword property value. The newVaultAndDirectoryPassword property
func (m *RequestPasswordResetRequest) SetNewVaultAndDirectoryPassword(value *string)() {
    m.newVaultAndDirectoryPassword = value
}
// SetNewVaultPassword sets the newVaultPassword property value. The newVaultPassword property
func (m *RequestPasswordResetRequest) SetNewVaultPassword(value *string)() {
    m.newVaultPassword = value
}
// SetUnsyncedPassword sets the unsyncedPassword property value. The unsyncedPassword property
func (m *RequestPasswordResetRequest) SetUnsyncedPassword(value *bool)() {
    m.unsyncedPassword = value
}
// RequestPasswordResetRequestable 
type RequestPasswordResetRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RequestModificationRequestable
    GetNewVaultAndDirectoryPassword()(*string)
    GetNewVaultPassword()(*string)
    GetUnsyncedPassword()(*bool)
    SetNewVaultAndDirectoryPassword(value *string)()
    SetNewVaultPassword(value *string)()
    SetUnsyncedPassword(value *bool)()
}
