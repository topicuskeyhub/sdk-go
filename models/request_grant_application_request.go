package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RequestGrantApplicationRequest struct {
    RequestAbstractAccessProfileModificationRequest
    // The application property
    application ClientClientApplicationPrimerable
}
// NewRequestGrantApplicationRequest instantiates a new RequestGrantApplicationRequest and sets the default values.
func NewRequestGrantApplicationRequest()(*RequestGrantApplicationRequest) {
    m := &RequestGrantApplicationRequest{
        RequestAbstractAccessProfileModificationRequest: *NewRequestAbstractAccessProfileModificationRequest(),
    }
    typeEscapedValue := "request.GrantApplicationRequest"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateRequestGrantApplicationRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRequestGrantApplicationRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequestGrantApplicationRequest(), nil
}
// GetApplication gets the application property value. The application property
// returns a ClientClientApplicationPrimerable when successful
func (m *RequestGrantApplicationRequest) GetApplication()(ClientClientApplicationPrimerable) {
    return m.application
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RequestGrantApplicationRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RequestAbstractAccessProfileModificationRequest.GetFieldDeserializers()
    res["application"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateClientClientApplicationPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplication(val.(ClientClientApplicationPrimerable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *RequestGrantApplicationRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RequestAbstractAccessProfileModificationRequest.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("application", m.GetApplication())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApplication sets the application property value. The application property
func (m *RequestGrantApplicationRequest) SetApplication(value ClientClientApplicationPrimerable)() {
    m.application = value
}
type RequestGrantApplicationRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RequestAbstractAccessProfileModificationRequestable
    GetApplication()(ClientClientApplicationPrimerable)
    SetApplication(value ClientClientApplicationPrimerable)()
}
