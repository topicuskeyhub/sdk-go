package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RequestAddGroupAdminRequest 
type RequestAddGroupAdminRequest struct {
    RequestModificationRequest
    // The newAdmin property
    newAdmin AuthAccountPrimerable
    // The privateKey property
    privateKey *string
}
// NewRequestAddGroupAdminRequest instantiates a new requestAddGroupAdminRequest and sets the default values.
func NewRequestAddGroupAdminRequest()(*RequestAddGroupAdminRequest) {
    m := &RequestAddGroupAdminRequest{
        RequestModificationRequest: *NewRequestModificationRequest(),
    }
    typeEscapedValue := "request.AddGroupAdminRequest"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateRequestAddGroupAdminRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRequestAddGroupAdminRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequestAddGroupAdminRequest(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RequestAddGroupAdminRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RequestModificationRequest.GetFieldDeserializers()
    res["newAdmin"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthAccountPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNewAdmin(val.(AuthAccountPrimerable))
        }
        return nil
    }
    res["privateKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivateKey(val)
        }
        return nil
    }
    return res
}
// GetNewAdmin gets the newAdmin property value. The newAdmin property
func (m *RequestAddGroupAdminRequest) GetNewAdmin()(AuthAccountPrimerable) {
    return m.newAdmin
}
// GetPrivateKey gets the privateKey property value. The privateKey property
func (m *RequestAddGroupAdminRequest) GetPrivateKey()(*string) {
    return m.privateKey
}
// Serialize serializes information the current object
func (m *RequestAddGroupAdminRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RequestModificationRequest.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("newAdmin", m.GetNewAdmin())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("privateKey", m.GetPrivateKey())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetNewAdmin sets the newAdmin property value. The newAdmin property
func (m *RequestAddGroupAdminRequest) SetNewAdmin(value AuthAccountPrimerable)() {
    m.newAdmin = value
}
// SetPrivateKey sets the privateKey property value. The privateKey property
func (m *RequestAddGroupAdminRequest) SetPrivateKey(value *string)() {
    m.privateKey = value
}
// RequestAddGroupAdminRequestable 
type RequestAddGroupAdminRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RequestModificationRequestable
    GetNewAdmin()(AuthAccountPrimerable)
    GetPrivateKey()(*string)
    SetNewAdmin(value AuthAccountPrimerable)()
    SetPrivateKey(value *string)()
}
