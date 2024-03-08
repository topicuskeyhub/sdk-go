package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RequestTransferGroupOnSystemOwnershipRequest struct {
    RequestModificationRequest
    // The groupOnSystem property
    groupOnSystem ProvisioningGroupOnSystemable
}
// NewRequestTransferGroupOnSystemOwnershipRequest instantiates a new RequestTransferGroupOnSystemOwnershipRequest and sets the default values.
func NewRequestTransferGroupOnSystemOwnershipRequest()(*RequestTransferGroupOnSystemOwnershipRequest) {
    m := &RequestTransferGroupOnSystemOwnershipRequest{
        RequestModificationRequest: *NewRequestModificationRequest(),
    }
    typeEscapedValue := "request.TransferGroupOnSystemOwnershipRequest"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateRequestTransferGroupOnSystemOwnershipRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRequestTransferGroupOnSystemOwnershipRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequestTransferGroupOnSystemOwnershipRequest(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RequestTransferGroupOnSystemOwnershipRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RequestModificationRequest.GetFieldDeserializers()
    res["groupOnSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProvisioningGroupOnSystemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupOnSystem(val.(ProvisioningGroupOnSystemable))
        }
        return nil
    }
    return res
}
// GetGroupOnSystem gets the groupOnSystem property value. The groupOnSystem property
// returns a ProvisioningGroupOnSystemable when successful
func (m *RequestTransferGroupOnSystemOwnershipRequest) GetGroupOnSystem()(ProvisioningGroupOnSystemable) {
    return m.groupOnSystem
}
// Serialize serializes information the current object
func (m *RequestTransferGroupOnSystemOwnershipRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RequestModificationRequest.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("groupOnSystem", m.GetGroupOnSystem())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetGroupOnSystem sets the groupOnSystem property value. The groupOnSystem property
func (m *RequestTransferGroupOnSystemOwnershipRequest) SetGroupOnSystem(value ProvisioningGroupOnSystemable)() {
    m.groupOnSystem = value
}
type RequestTransferGroupOnSystemOwnershipRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RequestModificationRequestable
    GetGroupOnSystem()(ProvisioningGroupOnSystemable)
    SetGroupOnSystem(value ProvisioningGroupOnSystemable)()
}
