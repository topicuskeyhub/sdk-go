package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RequestCreateGroupRequest 
type RequestCreateGroupRequest struct {
    RequestAbstractOrganizationalUnitModificationRequest
    // The groupName property
    groupName *string
}
// NewRequestCreateGroupRequest instantiates a new requestCreateGroupRequest and sets the default values.
func NewRequestCreateGroupRequest()(*RequestCreateGroupRequest) {
    m := &RequestCreateGroupRequest{
        RequestAbstractOrganizationalUnitModificationRequest: *NewRequestAbstractOrganizationalUnitModificationRequest(),
    }
    typeEscapedValue := "request.CreateGroupRequest"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateRequestCreateGroupRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRequestCreateGroupRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequestCreateGroupRequest(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RequestCreateGroupRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RequestAbstractOrganizationalUnitModificationRequest.GetFieldDeserializers()
    res["groupName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupName(val)
        }
        return nil
    }
    return res
}
// GetGroupName gets the groupName property value. The groupName property
func (m *RequestCreateGroupRequest) GetGroupName()(*string) {
    return m.groupName
}
// Serialize serializes information the current object
func (m *RequestCreateGroupRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RequestAbstractOrganizationalUnitModificationRequest.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("groupName", m.GetGroupName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetGroupName sets the groupName property value. The groupName property
func (m *RequestCreateGroupRequest) SetGroupName(value *string)() {
    m.groupName = value
}
// RequestCreateGroupRequestable 
type RequestCreateGroupRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RequestAbstractOrganizationalUnitModificationRequestable
    GetGroupName()(*string)
    SetGroupName(value *string)()
}
