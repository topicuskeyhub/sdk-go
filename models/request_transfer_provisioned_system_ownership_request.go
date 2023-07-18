package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RequestTransferProvisionedSystemOwnershipRequest 
type RequestTransferProvisionedSystemOwnershipRequest struct {
    RequestAbstractProvisionedSystemModificationRequest
}
// NewRequestTransferProvisionedSystemOwnershipRequest instantiates a new requestTransferProvisionedSystemOwnershipRequest and sets the default values.
func NewRequestTransferProvisionedSystemOwnershipRequest()(*RequestTransferProvisionedSystemOwnershipRequest) {
    m := &RequestTransferProvisionedSystemOwnershipRequest{
        RequestAbstractProvisionedSystemModificationRequest: *NewRequestAbstractProvisionedSystemModificationRequest(),
    }
    typeEscapedValue := "request.TransferProvisionedSystemOwnershipRequest"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateRequestTransferProvisionedSystemOwnershipRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRequestTransferProvisionedSystemOwnershipRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequestTransferProvisionedSystemOwnershipRequest(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RequestTransferProvisionedSystemOwnershipRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RequestAbstractProvisionedSystemModificationRequest.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *RequestTransferProvisionedSystemOwnershipRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RequestAbstractProvisionedSystemModificationRequest.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// RequestTransferProvisionedSystemOwnershipRequestable 
type RequestTransferProvisionedSystemOwnershipRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RequestAbstractProvisionedSystemModificationRequestable
}
