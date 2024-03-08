package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RequestGrantGroupOnSystemRequestRequest struct {
    RequestModificationRequest
    // The activationRequired property
    activationRequired *bool
    // The groupOnSystem property
    groupOnSystem ProvisioningGroupOnSystemable
}
// NewRequestGrantGroupOnSystemRequestRequest instantiates a new RequestGrantGroupOnSystemRequestRequest and sets the default values.
func NewRequestGrantGroupOnSystemRequestRequest()(*RequestGrantGroupOnSystemRequestRequest) {
    m := &RequestGrantGroupOnSystemRequestRequest{
        RequestModificationRequest: *NewRequestModificationRequest(),
    }
    typeEscapedValue := "request.GrantGroupOnSystemRequestRequest"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateRequestGrantGroupOnSystemRequestRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRequestGrantGroupOnSystemRequestRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequestGrantGroupOnSystemRequestRequest(), nil
}
// GetActivationRequired gets the activationRequired property value. The activationRequired property
// returns a *bool when successful
func (m *RequestGrantGroupOnSystemRequestRequest) GetActivationRequired()(*bool) {
    return m.activationRequired
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RequestGrantGroupOnSystemRequestRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RequestModificationRequest.GetFieldDeserializers()
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
func (m *RequestGrantGroupOnSystemRequestRequest) GetGroupOnSystem()(ProvisioningGroupOnSystemable) {
    return m.groupOnSystem
}
// Serialize serializes information the current object
func (m *RequestGrantGroupOnSystemRequestRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RequestModificationRequest.Serialize(writer)
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
func (m *RequestGrantGroupOnSystemRequestRequest) SetActivationRequired(value *bool)() {
    m.activationRequired = value
}
// SetGroupOnSystem sets the groupOnSystem property value. The groupOnSystem property
func (m *RequestGrantGroupOnSystemRequestRequest) SetGroupOnSystem(value ProvisioningGroupOnSystemable)() {
    m.groupOnSystem = value
}
type RequestGrantGroupOnSystemRequestRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RequestModificationRequestable
    GetActivationRequired()(*bool)
    GetGroupOnSystem()(ProvisioningGroupOnSystemable)
    SetActivationRequired(value *bool)()
    SetGroupOnSystem(value ProvisioningGroupOnSystemable)()
}
