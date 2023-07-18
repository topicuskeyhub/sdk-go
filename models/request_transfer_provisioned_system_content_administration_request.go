package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RequestTransferProvisionedSystemContentAdministrationRequest 
type RequestTransferProvisionedSystemContentAdministrationRequest struct {
    RequestAbstractProvisionedSystemModificationRequest
}
// NewRequestTransferProvisionedSystemContentAdministrationRequest instantiates a new requestTransferProvisionedSystemContentAdministrationRequest and sets the default values.
func NewRequestTransferProvisionedSystemContentAdministrationRequest()(*RequestTransferProvisionedSystemContentAdministrationRequest) {
    m := &RequestTransferProvisionedSystemContentAdministrationRequest{
        RequestAbstractProvisionedSystemModificationRequest: *NewRequestAbstractProvisionedSystemModificationRequest(),
    }
    typeEscapedValue := "request.TransferProvisionedSystemContentAdministrationRequest"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateRequestTransferProvisionedSystemContentAdministrationRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRequestTransferProvisionedSystemContentAdministrationRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequestTransferProvisionedSystemContentAdministrationRequest(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RequestTransferProvisionedSystemContentAdministrationRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RequestAbstractProvisionedSystemModificationRequest.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *RequestTransferProvisionedSystemContentAdministrationRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RequestAbstractProvisionedSystemModificationRequest.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// RequestTransferProvisionedSystemContentAdministrationRequestable 
type RequestTransferProvisionedSystemContentAdministrationRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RequestAbstractProvisionedSystemModificationRequestable
}
