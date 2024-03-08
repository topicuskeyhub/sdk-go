package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RequestRemoveProvisionedSystemRequest struct {
    RequestAbstractProvisionedSystemModificationRequest
    // The systemName property
    systemName *string
}
// NewRequestRemoveProvisionedSystemRequest instantiates a new RequestRemoveProvisionedSystemRequest and sets the default values.
func NewRequestRemoveProvisionedSystemRequest()(*RequestRemoveProvisionedSystemRequest) {
    m := &RequestRemoveProvisionedSystemRequest{
        RequestAbstractProvisionedSystemModificationRequest: *NewRequestAbstractProvisionedSystemModificationRequest(),
    }
    typeEscapedValue := "request.RemoveProvisionedSystemRequest"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateRequestRemoveProvisionedSystemRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRequestRemoveProvisionedSystemRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequestRemoveProvisionedSystemRequest(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RequestRemoveProvisionedSystemRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RequestAbstractProvisionedSystemModificationRequest.GetFieldDeserializers()
    res["systemName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystemName(val)
        }
        return nil
    }
    return res
}
// GetSystemName gets the systemName property value. The systemName property
// returns a *string when successful
func (m *RequestRemoveProvisionedSystemRequest) GetSystemName()(*string) {
    return m.systemName
}
// Serialize serializes information the current object
func (m *RequestRemoveProvisionedSystemRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RequestAbstractProvisionedSystemModificationRequest.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("systemName", m.GetSystemName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSystemName sets the systemName property value. The systemName property
func (m *RequestRemoveProvisionedSystemRequest) SetSystemName(value *string)() {
    m.systemName = value
}
type RequestRemoveProvisionedSystemRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RequestAbstractProvisionedSystemModificationRequestable
    GetSystemName()(*string)
    SetSystemName(value *string)()
}
