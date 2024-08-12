package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RequestTransferAccessProfileOwnershipRequest struct {
    RequestAbstractAccessProfileModificationRequest
}
// NewRequestTransferAccessProfileOwnershipRequest instantiates a new RequestTransferAccessProfileOwnershipRequest and sets the default values.
func NewRequestTransferAccessProfileOwnershipRequest()(*RequestTransferAccessProfileOwnershipRequest) {
    m := &RequestTransferAccessProfileOwnershipRequest{
        RequestAbstractAccessProfileModificationRequest: *NewRequestAbstractAccessProfileModificationRequest(),
    }
    typeEscapedValue := "request.TransferAccessProfileOwnershipRequest"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateRequestTransferAccessProfileOwnershipRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRequestTransferAccessProfileOwnershipRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequestTransferAccessProfileOwnershipRequest(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RequestTransferAccessProfileOwnershipRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RequestAbstractAccessProfileModificationRequest.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *RequestTransferAccessProfileOwnershipRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RequestAbstractAccessProfileModificationRequest.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type RequestTransferAccessProfileOwnershipRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RequestAbstractAccessProfileModificationRequestable
}
