package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RequestCreateGroupOnSystemRequest struct {
    RequestAbstractProvisionedSystemModificationRequest
    // The activationRequired property
    activationRequired *bool
    // The groupOnSystemType property
    groupOnSystemType *ProvisioningGroupOnSystemType
    // The nameInSystem property
    nameInSystem *string
    // The owner property
    owner GroupGroupPrimerable
}
// NewRequestCreateGroupOnSystemRequest instantiates a new RequestCreateGroupOnSystemRequest and sets the default values.
func NewRequestCreateGroupOnSystemRequest()(*RequestCreateGroupOnSystemRequest) {
    m := &RequestCreateGroupOnSystemRequest{
        RequestAbstractProvisionedSystemModificationRequest: *NewRequestAbstractProvisionedSystemModificationRequest(),
    }
    typeEscapedValue := "request.CreateGroupOnSystemRequest"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateRequestCreateGroupOnSystemRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRequestCreateGroupOnSystemRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequestCreateGroupOnSystemRequest(), nil
}
// GetActivationRequired gets the activationRequired property value. The activationRequired property
// returns a *bool when successful
func (m *RequestCreateGroupOnSystemRequest) GetActivationRequired()(*bool) {
    return m.activationRequired
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RequestCreateGroupOnSystemRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RequestAbstractProvisionedSystemModificationRequest.GetFieldDeserializers()
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
    res["groupOnSystemType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseProvisioningGroupOnSystemType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupOnSystemType(val.(*ProvisioningGroupOnSystemType))
        }
        return nil
    }
    res["nameInSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNameInSystem(val)
        }
        return nil
    }
    res["owner"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupGroupPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwner(val.(GroupGroupPrimerable))
        }
        return nil
    }
    return res
}
// GetGroupOnSystemType gets the groupOnSystemType property value. The groupOnSystemType property
// returns a *ProvisioningGroupOnSystemType when successful
func (m *RequestCreateGroupOnSystemRequest) GetGroupOnSystemType()(*ProvisioningGroupOnSystemType) {
    return m.groupOnSystemType
}
// GetNameInSystem gets the nameInSystem property value. The nameInSystem property
// returns a *string when successful
func (m *RequestCreateGroupOnSystemRequest) GetNameInSystem()(*string) {
    return m.nameInSystem
}
// GetOwner gets the owner property value. The owner property
// returns a GroupGroupPrimerable when successful
func (m *RequestCreateGroupOnSystemRequest) GetOwner()(GroupGroupPrimerable) {
    return m.owner
}
// Serialize serializes information the current object
func (m *RequestCreateGroupOnSystemRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RequestAbstractProvisionedSystemModificationRequest.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("activationRequired", m.GetActivationRequired())
        if err != nil {
            return err
        }
    }
    if m.GetGroupOnSystemType() != nil {
        cast := (*m.GetGroupOnSystemType()).String()
        err = writer.WriteStringValue("groupOnSystemType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("nameInSystem", m.GetNameInSystem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("owner", m.GetOwner())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActivationRequired sets the activationRequired property value. The activationRequired property
func (m *RequestCreateGroupOnSystemRequest) SetActivationRequired(value *bool)() {
    m.activationRequired = value
}
// SetGroupOnSystemType sets the groupOnSystemType property value. The groupOnSystemType property
func (m *RequestCreateGroupOnSystemRequest) SetGroupOnSystemType(value *ProvisioningGroupOnSystemType)() {
    m.groupOnSystemType = value
}
// SetNameInSystem sets the nameInSystem property value. The nameInSystem property
func (m *RequestCreateGroupOnSystemRequest) SetNameInSystem(value *string)() {
    m.nameInSystem = value
}
// SetOwner sets the owner property value. The owner property
func (m *RequestCreateGroupOnSystemRequest) SetOwner(value GroupGroupPrimerable)() {
    m.owner = value
}
type RequestCreateGroupOnSystemRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RequestAbstractProvisionedSystemModificationRequestable
    GetActivationRequired()(*bool)
    GetGroupOnSystemType()(*ProvisioningGroupOnSystemType)
    GetNameInSystem()(*string)
    GetOwner()(GroupGroupPrimerable)
    SetActivationRequired(value *bool)()
    SetGroupOnSystemType(value *ProvisioningGroupOnSystemType)()
    SetNameInSystem(value *string)()
    SetOwner(value GroupGroupPrimerable)()
}
