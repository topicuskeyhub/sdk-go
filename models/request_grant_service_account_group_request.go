package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RequestGrantServiceAccountGroupRequest struct {
    RequestModificationRequest
    // The groupOnSystem property
    groupOnSystem ProvisioningGroupOnSystemable
    // The serviceAccount property
    serviceAccount ServiceaccountServiceAccountPrimerable
}
// NewRequestGrantServiceAccountGroupRequest instantiates a new RequestGrantServiceAccountGroupRequest and sets the default values.
func NewRequestGrantServiceAccountGroupRequest()(*RequestGrantServiceAccountGroupRequest) {
    m := &RequestGrantServiceAccountGroupRequest{
        RequestModificationRequest: *NewRequestModificationRequest(),
    }
    typeEscapedValue := "request.GrantServiceAccountGroupRequest"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateRequestGrantServiceAccountGroupRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRequestGrantServiceAccountGroupRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequestGrantServiceAccountGroupRequest(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RequestGrantServiceAccountGroupRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["serviceAccount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateServiceaccountServiceAccountPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServiceAccount(val.(ServiceaccountServiceAccountPrimerable))
        }
        return nil
    }
    return res
}
// GetGroupOnSystem gets the groupOnSystem property value. The groupOnSystem property
// returns a ProvisioningGroupOnSystemable when successful
func (m *RequestGrantServiceAccountGroupRequest) GetGroupOnSystem()(ProvisioningGroupOnSystemable) {
    return m.groupOnSystem
}
// GetServiceAccount gets the serviceAccount property value. The serviceAccount property
// returns a ServiceaccountServiceAccountPrimerable when successful
func (m *RequestGrantServiceAccountGroupRequest) GetServiceAccount()(ServiceaccountServiceAccountPrimerable) {
    return m.serviceAccount
}
// Serialize serializes information the current object
func (m *RequestGrantServiceAccountGroupRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err = writer.WriteObjectValue("serviceAccount", m.GetServiceAccount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetGroupOnSystem sets the groupOnSystem property value. The groupOnSystem property
func (m *RequestGrantServiceAccountGroupRequest) SetGroupOnSystem(value ProvisioningGroupOnSystemable)() {
    m.groupOnSystem = value
}
// SetServiceAccount sets the serviceAccount property value. The serviceAccount property
func (m *RequestGrantServiceAccountGroupRequest) SetServiceAccount(value ServiceaccountServiceAccountPrimerable)() {
    m.serviceAccount = value
}
type RequestGrantServiceAccountGroupRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RequestModificationRequestable
    GetGroupOnSystem()(ProvisioningGroupOnSystemable)
    GetServiceAccount()(ServiceaccountServiceAccountPrimerable)
    SetGroupOnSystem(value ProvisioningGroupOnSystemable)()
    SetServiceAccount(value ServiceaccountServiceAccountPrimerable)()
}
