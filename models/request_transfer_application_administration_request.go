package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RequestTransferApplicationAdministrationRequest struct {
    RequestAbstractApplicationModificationRequest
}
// NewRequestTransferApplicationAdministrationRequest instantiates a new RequestTransferApplicationAdministrationRequest and sets the default values.
func NewRequestTransferApplicationAdministrationRequest()(*RequestTransferApplicationAdministrationRequest) {
    m := &RequestTransferApplicationAdministrationRequest{
        RequestAbstractApplicationModificationRequest: *NewRequestAbstractApplicationModificationRequest(),
    }
    typeEscapedValue := "request.TransferApplicationAdministrationRequest"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateRequestTransferApplicationAdministrationRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRequestTransferApplicationAdministrationRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequestTransferApplicationAdministrationRequest(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RequestTransferApplicationAdministrationRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RequestAbstractApplicationModificationRequest.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *RequestTransferApplicationAdministrationRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RequestAbstractApplicationModificationRequest.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type RequestTransferApplicationAdministrationRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RequestAbstractApplicationModificationRequestable
}
