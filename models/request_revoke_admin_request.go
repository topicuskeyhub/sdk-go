package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RequestRevokeAdminRequest struct {
    RequestModificationRequest
    // The admin property
    admin AuthAccountPrimerable
}
// NewRequestRevokeAdminRequest instantiates a new RequestRevokeAdminRequest and sets the default values.
func NewRequestRevokeAdminRequest()(*RequestRevokeAdminRequest) {
    m := &RequestRevokeAdminRequest{
        RequestModificationRequest: *NewRequestModificationRequest(),
    }
    typeEscapedValue := "request.RevokeAdminRequest"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateRequestRevokeAdminRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRequestRevokeAdminRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequestRevokeAdminRequest(), nil
}
// GetAdmin gets the admin property value. The admin property
// returns a AuthAccountPrimerable when successful
func (m *RequestRevokeAdminRequest) GetAdmin()(AuthAccountPrimerable) {
    return m.admin
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RequestRevokeAdminRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RequestModificationRequest.GetFieldDeserializers()
    res["admin"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthAccountPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdmin(val.(AuthAccountPrimerable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *RequestRevokeAdminRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RequestModificationRequest.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("admin", m.GetAdmin())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdmin sets the admin property value. The admin property
func (m *RequestRevokeAdminRequest) SetAdmin(value AuthAccountPrimerable)() {
    m.admin = value
}
type RequestRevokeAdminRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RequestModificationRequestable
    GetAdmin()(AuthAccountPrimerable)
    SetAdmin(value AuthAccountPrimerable)()
}
