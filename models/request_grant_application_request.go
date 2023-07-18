package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RequestGrantApplicationRequest 
type RequestGrantApplicationRequest struct {
    RequestAbstractApplicationModificationRequest
}
// NewRequestGrantApplicationRequest instantiates a new requestGrantApplicationRequest and sets the default values.
func NewRequestGrantApplicationRequest()(*RequestGrantApplicationRequest) {
    m := &RequestGrantApplicationRequest{
        RequestAbstractApplicationModificationRequest: *NewRequestAbstractApplicationModificationRequest(),
    }
    typeEscapedValue := "request.GrantApplicationRequest"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateRequestGrantApplicationRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRequestGrantApplicationRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequestGrantApplicationRequest(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RequestGrantApplicationRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RequestAbstractApplicationModificationRequest.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *RequestGrantApplicationRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RequestAbstractApplicationModificationRequest.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// RequestGrantApplicationRequestable 
type RequestGrantApplicationRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RequestAbstractApplicationModificationRequestable
}
