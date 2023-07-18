package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RequestTransferOrganizationalUnitOwnershipRequest 
type RequestTransferOrganizationalUnitOwnershipRequest struct {
    RequestAbstractOrganizationalUnitModificationRequest
}
// NewRequestTransferOrganizationalUnitOwnershipRequest instantiates a new requestTransferOrganizationalUnitOwnershipRequest and sets the default values.
func NewRequestTransferOrganizationalUnitOwnershipRequest()(*RequestTransferOrganizationalUnitOwnershipRequest) {
    m := &RequestTransferOrganizationalUnitOwnershipRequest{
        RequestAbstractOrganizationalUnitModificationRequest: *NewRequestAbstractOrganizationalUnitModificationRequest(),
    }
    typeEscapedValue := "request.TransferOrganizationalUnitOwnershipRequest"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateRequestTransferOrganizationalUnitOwnershipRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRequestTransferOrganizationalUnitOwnershipRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequestTransferOrganizationalUnitOwnershipRequest(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RequestTransferOrganizationalUnitOwnershipRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RequestAbstractOrganizationalUnitModificationRequest.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *RequestTransferOrganizationalUnitOwnershipRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RequestAbstractOrganizationalUnitModificationRequest.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// RequestTransferOrganizationalUnitOwnershipRequestable 
type RequestTransferOrganizationalUnitOwnershipRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RequestAbstractOrganizationalUnitModificationRequestable
}
