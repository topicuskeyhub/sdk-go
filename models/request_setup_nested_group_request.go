package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RequestSetupNestedGroupRequest 
type RequestSetupNestedGroupRequest struct {
    RequestModificationRequest
    // The accountAction property
    accountAction *RequestGroupNestAccountAction
    // The connect property
    connect *bool
    // The requestingGroup property
    requestingGroup GroupGroupPrimerable
}
// NewRequestSetupNestedGroupRequest instantiates a new requestSetupNestedGroupRequest and sets the default values.
func NewRequestSetupNestedGroupRequest()(*RequestSetupNestedGroupRequest) {
    m := &RequestSetupNestedGroupRequest{
        RequestModificationRequest: *NewRequestModificationRequest(),
    }
    typeEscapedValue := "request.SetupNestedGroupRequest"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateRequestSetupNestedGroupRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRequestSetupNestedGroupRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequestSetupNestedGroupRequest(), nil
}
// GetAccountAction gets the accountAction property value. The accountAction property
func (m *RequestSetupNestedGroupRequest) GetAccountAction()(*RequestGroupNestAccountAction) {
    return m.accountAction
}
// GetConnect gets the connect property value. The connect property
func (m *RequestSetupNestedGroupRequest) GetConnect()(*bool) {
    return m.connect
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RequestSetupNestedGroupRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RequestModificationRequest.GetFieldDeserializers()
    res["accountAction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRequestGroupNestAccountAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountAction(val.(*RequestGroupNestAccountAction))
        }
        return nil
    }
    res["connect"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnect(val)
        }
        return nil
    }
    res["requestingGroup"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupGroupPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestingGroup(val.(GroupGroupPrimerable))
        }
        return nil
    }
    return res
}
// GetRequestingGroup gets the requestingGroup property value. The requestingGroup property
func (m *RequestSetupNestedGroupRequest) GetRequestingGroup()(GroupGroupPrimerable) {
    return m.requestingGroup
}
// Serialize serializes information the current object
func (m *RequestSetupNestedGroupRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RequestModificationRequest.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAccountAction() != nil {
        cast := (*m.GetAccountAction()).String()
        err = writer.WriteStringValue("accountAction", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("connect", m.GetConnect())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("requestingGroup", m.GetRequestingGroup())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccountAction sets the accountAction property value. The accountAction property
func (m *RequestSetupNestedGroupRequest) SetAccountAction(value *RequestGroupNestAccountAction)() {
    m.accountAction = value
}
// SetConnect sets the connect property value. The connect property
func (m *RequestSetupNestedGroupRequest) SetConnect(value *bool)() {
    m.connect = value
}
// SetRequestingGroup sets the requestingGroup property value. The requestingGroup property
func (m *RequestSetupNestedGroupRequest) SetRequestingGroup(value GroupGroupPrimerable)() {
    m.requestingGroup = value
}
// RequestSetupNestedGroupRequestable 
type RequestSetupNestedGroupRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RequestModificationRequestable
    GetAccountAction()(*RequestGroupNestAccountAction)
    GetConnect()(*bool)
    GetRequestingGroup()(GroupGroupPrimerable)
    SetAccountAction(value *RequestGroupNestAccountAction)()
    SetConnect(value *bool)()
    SetRequestingGroup(value GroupGroupPrimerable)()
}
