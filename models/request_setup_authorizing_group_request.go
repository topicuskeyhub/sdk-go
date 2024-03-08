package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RequestSetupAuthorizingGroupRequest struct {
    RequestModificationRequest
    // The authorizingGroupType property
    authorizingGroupType *RequestAuthorizingGroupType
    // The connect property
    connect *bool
    // The requestingGroup property
    requestingGroup GroupGroupPrimerable
}
// NewRequestSetupAuthorizingGroupRequest instantiates a new RequestSetupAuthorizingGroupRequest and sets the default values.
func NewRequestSetupAuthorizingGroupRequest()(*RequestSetupAuthorizingGroupRequest) {
    m := &RequestSetupAuthorizingGroupRequest{
        RequestModificationRequest: *NewRequestModificationRequest(),
    }
    typeEscapedValue := "request.SetupAuthorizingGroupRequest"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateRequestSetupAuthorizingGroupRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRequestSetupAuthorizingGroupRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequestSetupAuthorizingGroupRequest(), nil
}
// GetAuthorizingGroupType gets the authorizingGroupType property value. The authorizingGroupType property
// returns a *RequestAuthorizingGroupType when successful
func (m *RequestSetupAuthorizingGroupRequest) GetAuthorizingGroupType()(*RequestAuthorizingGroupType) {
    return m.authorizingGroupType
}
// GetConnect gets the connect property value. The connect property
// returns a *bool when successful
func (m *RequestSetupAuthorizingGroupRequest) GetConnect()(*bool) {
    return m.connect
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RequestSetupAuthorizingGroupRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RequestModificationRequest.GetFieldDeserializers()
    res["authorizingGroupType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRequestAuthorizingGroupType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthorizingGroupType(val.(*RequestAuthorizingGroupType))
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
// returns a GroupGroupPrimerable when successful
func (m *RequestSetupAuthorizingGroupRequest) GetRequestingGroup()(GroupGroupPrimerable) {
    return m.requestingGroup
}
// Serialize serializes information the current object
func (m *RequestSetupAuthorizingGroupRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RequestModificationRequest.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAuthorizingGroupType() != nil {
        cast := (*m.GetAuthorizingGroupType()).String()
        err = writer.WriteStringValue("authorizingGroupType", &cast)
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
// SetAuthorizingGroupType sets the authorizingGroupType property value. The authorizingGroupType property
func (m *RequestSetupAuthorizingGroupRequest) SetAuthorizingGroupType(value *RequestAuthorizingGroupType)() {
    m.authorizingGroupType = value
}
// SetConnect sets the connect property value. The connect property
func (m *RequestSetupAuthorizingGroupRequest) SetConnect(value *bool)() {
    m.connect = value
}
// SetRequestingGroup sets the requestingGroup property value. The requestingGroup property
func (m *RequestSetupAuthorizingGroupRequest) SetRequestingGroup(value GroupGroupPrimerable)() {
    m.requestingGroup = value
}
type RequestSetupAuthorizingGroupRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RequestModificationRequestable
    GetAuthorizingGroupType()(*RequestAuthorizingGroupType)
    GetConnect()(*bool)
    GetRequestingGroup()(GroupGroupPrimerable)
    SetAuthorizingGroupType(value *RequestAuthorizingGroupType)()
    SetConnect(value *bool)()
    SetRequestingGroup(value GroupGroupPrimerable)()
}
