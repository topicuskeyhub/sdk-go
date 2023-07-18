package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RequestRemoveGroupRequest 
type RequestRemoveGroupRequest struct {
    RequestModificationRequest
    // The groupName property
    groupName *string
}
// NewRequestRemoveGroupRequest instantiates a new requestRemoveGroupRequest and sets the default values.
func NewRequestRemoveGroupRequest()(*RequestRemoveGroupRequest) {
    m := &RequestRemoveGroupRequest{
        RequestModificationRequest: *NewRequestModificationRequest(),
    }
    typeEscapedValue := "request.RemoveGroupRequest"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateRequestRemoveGroupRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRequestRemoveGroupRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequestRemoveGroupRequest(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RequestRemoveGroupRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RequestModificationRequest.GetFieldDeserializers()
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
func (m *RequestRemoveGroupRequest) GetGroupName()(*string) {
    return m.groupName
}
// Serialize serializes information the current object
func (m *RequestRemoveGroupRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RequestModificationRequest.Serialize(writer)
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
func (m *RequestRemoveGroupRequest) SetGroupName(value *string)() {
    m.groupName = value
}
// RequestRemoveGroupRequestable 
type RequestRemoveGroupRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RequestModificationRequestable
    GetGroupName()(*string)
    SetGroupName(value *string)()
}
