package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RequestVerifyInternalAccountRequest struct {
    RequestModificationRequest
    // The internalAccountName property
    internalAccountName *string
}
// NewRequestVerifyInternalAccountRequest instantiates a new RequestVerifyInternalAccountRequest and sets the default values.
func NewRequestVerifyInternalAccountRequest()(*RequestVerifyInternalAccountRequest) {
    m := &RequestVerifyInternalAccountRequest{
        RequestModificationRequest: *NewRequestModificationRequest(),
    }
    typeEscapedValue := "request.VerifyInternalAccountRequest"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateRequestVerifyInternalAccountRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRequestVerifyInternalAccountRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequestVerifyInternalAccountRequest(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RequestVerifyInternalAccountRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RequestModificationRequest.GetFieldDeserializers()
    res["internalAccountName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInternalAccountName(val)
        }
        return nil
    }
    return res
}
// GetInternalAccountName gets the internalAccountName property value. The internalAccountName property
// returns a *string when successful
func (m *RequestVerifyInternalAccountRequest) GetInternalAccountName()(*string) {
    return m.internalAccountName
}
// Serialize serializes information the current object
func (m *RequestVerifyInternalAccountRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RequestModificationRequest.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("internalAccountName", m.GetInternalAccountName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetInternalAccountName sets the internalAccountName property value. The internalAccountName property
func (m *RequestVerifyInternalAccountRequest) SetInternalAccountName(value *string)() {
    m.internalAccountName = value
}
type RequestVerifyInternalAccountRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RequestModificationRequestable
    GetInternalAccountName()(*string)
    SetInternalAccountName(value *string)()
}
