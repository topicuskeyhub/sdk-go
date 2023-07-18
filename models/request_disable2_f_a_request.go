package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RequestDisable2FARequest 
type RequestDisable2FARequest struct {
    RequestModificationRequest
    // The subject property
    subject *string
}
// NewRequestDisable2FARequest instantiates a new requestDisable2FARequest and sets the default values.
func NewRequestDisable2FARequest()(*RequestDisable2FARequest) {
    m := &RequestDisable2FARequest{
        RequestModificationRequest: *NewRequestModificationRequest(),
    }
    typeEscapedValue := "request.Disable2FARequest"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateRequestDisable2FARequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRequestDisable2FARequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequestDisable2FARequest(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RequestDisable2FARequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RequestModificationRequest.GetFieldDeserializers()
    res["subject"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubject(val)
        }
        return nil
    }
    return res
}
// GetSubject gets the subject property value. The subject property
func (m *RequestDisable2FARequest) GetSubject()(*string) {
    return m.subject
}
// Serialize serializes information the current object
func (m *RequestDisable2FARequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RequestModificationRequest.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("subject", m.GetSubject())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSubject sets the subject property value. The subject property
func (m *RequestDisable2FARequest) SetSubject(value *string)() {
    m.subject = value
}
// RequestDisable2FARequestable 
type RequestDisable2FARequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RequestModificationRequestable
    GetSubject()(*string)
    SetSubject(value *string)()
}
