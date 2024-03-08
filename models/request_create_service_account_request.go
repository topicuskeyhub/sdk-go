package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RequestCreateServiceAccountRequest struct {
    RequestAbstractProvisionedSystemModificationRequest
    // The username property
    username *string
}
// NewRequestCreateServiceAccountRequest instantiates a new RequestCreateServiceAccountRequest and sets the default values.
func NewRequestCreateServiceAccountRequest()(*RequestCreateServiceAccountRequest) {
    m := &RequestCreateServiceAccountRequest{
        RequestAbstractProvisionedSystemModificationRequest: *NewRequestAbstractProvisionedSystemModificationRequest(),
    }
    typeEscapedValue := "request.CreateServiceAccountRequest"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateRequestCreateServiceAccountRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRequestCreateServiceAccountRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequestCreateServiceAccountRequest(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RequestCreateServiceAccountRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RequestAbstractProvisionedSystemModificationRequest.GetFieldDeserializers()
    res["username"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsername(val)
        }
        return nil
    }
    return res
}
// GetUsername gets the username property value. The username property
// returns a *string when successful
func (m *RequestCreateServiceAccountRequest) GetUsername()(*string) {
    return m.username
}
// Serialize serializes information the current object
func (m *RequestCreateServiceAccountRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RequestAbstractProvisionedSystemModificationRequest.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("username", m.GetUsername())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetUsername sets the username property value. The username property
func (m *RequestCreateServiceAccountRequest) SetUsername(value *string)() {
    m.username = value
}
type RequestCreateServiceAccountRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RequestAbstractProvisionedSystemModificationRequestable
    GetUsername()(*string)
    SetUsername(value *string)()
}
