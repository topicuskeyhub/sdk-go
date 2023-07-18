package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RequestTransferProvisionedSystemAdministrationRequest 
type RequestTransferProvisionedSystemAdministrationRequest struct {
    RequestAbstractProvisionedSystemModificationRequest
}
// NewRequestTransferProvisionedSystemAdministrationRequest instantiates a new requestTransferProvisionedSystemAdministrationRequest and sets the default values.
func NewRequestTransferProvisionedSystemAdministrationRequest()(*RequestTransferProvisionedSystemAdministrationRequest) {
    m := &RequestTransferProvisionedSystemAdministrationRequest{
        RequestAbstractProvisionedSystemModificationRequest: *NewRequestAbstractProvisionedSystemModificationRequest(),
    }
    typeEscapedValue := "request.TransferProvisionedSystemAdministrationRequest"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateRequestTransferProvisionedSystemAdministrationRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRequestTransferProvisionedSystemAdministrationRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequestTransferProvisionedSystemAdministrationRequest(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RequestTransferProvisionedSystemAdministrationRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RequestAbstractProvisionedSystemModificationRequest.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *RequestTransferProvisionedSystemAdministrationRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RequestAbstractProvisionedSystemModificationRequest.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// RequestTransferProvisionedSystemAdministrationRequestable 
type RequestTransferProvisionedSystemAdministrationRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RequestAbstractProvisionedSystemModificationRequestable
}
