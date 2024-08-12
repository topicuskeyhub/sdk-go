package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RequestGrantGroupOnSystemRequest struct {
    RequestAbstractAccessProfileModificationRequest
    // The activationRequired property
    activationRequired *bool
    // The groupOnSystem property
    groupOnSystem ProvisioningGroupOnSystemable
}
// NewRequestGrantGroupOnSystemRequest instantiates a new RequestGrantGroupOnSystemRequest and sets the default values.
func NewRequestGrantGroupOnSystemRequest()(*RequestGrantGroupOnSystemRequest) {
    m := &RequestGrantGroupOnSystemRequest{
        RequestAbstractAccessProfileModificationRequest: *NewRequestAbstractAccessProfileModificationRequest(),
    }
    typeEscapedValue := "request.GrantGroupOnSystemRequest"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateRequestGrantGroupOnSystemRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRequestGrantGroupOnSystemRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequestGrantGroupOnSystemRequest(), nil
}
// GetActivationRequired gets the activationRequired property value. The activationRequired property
// returns a *bool when successful
func (m *RequestGrantGroupOnSystemRequest) GetActivationRequired()(*bool) {
    return m.activationRequired
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RequestGrantGroupOnSystemRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RequestAbstractAccessProfileModificationRequest.GetFieldDeserializers()
    res["activationRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivationRequired(val)
        }
        return nil
    }
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
func (m *RequestGrantGroupOnSystemRequest) GetGroupOnSystem()(ProvisioningGroupOnSystemable) {
    return m.groupOnSystem
}
// Serialize serializes information the current object
func (m *RequestGrantGroupOnSystemRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RequestAbstractAccessProfileModificationRequest.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("activationRequired", m.GetActivationRequired())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("groupOnSystem", m.GetGroupOnSystem())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActivationRequired sets the activationRequired property value. The activationRequired property
func (m *RequestGrantGroupOnSystemRequest) SetActivationRequired(value *bool)() {
    m.activationRequired = value
}
// SetGroupOnSystem sets the groupOnSystem property value. The groupOnSystem property
func (m *RequestGrantGroupOnSystemRequest) SetGroupOnSystem(value ProvisioningGroupOnSystemable)() {
    m.groupOnSystem = value
}
type RequestGrantGroupOnSystemRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RequestAbstractAccessProfileModificationRequestable
    GetActivationRequired()(*bool)
    GetGroupOnSystem()(ProvisioningGroupOnSystemable)
    SetActivationRequired(value *bool)()
    SetGroupOnSystem(value ProvisioningGroupOnSystemable)()
}
