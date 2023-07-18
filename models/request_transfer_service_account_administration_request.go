package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RequestTransferServiceAccountAdministrationRequest 
type RequestTransferServiceAccountAdministrationRequest struct {
    RequestModificationRequest
    // The serviceAccount property
    serviceAccount ServiceaccountServiceAccountPrimerable
}
// NewRequestTransferServiceAccountAdministrationRequest instantiates a new requestTransferServiceAccountAdministrationRequest and sets the default values.
func NewRequestTransferServiceAccountAdministrationRequest()(*RequestTransferServiceAccountAdministrationRequest) {
    m := &RequestTransferServiceAccountAdministrationRequest{
        RequestModificationRequest: *NewRequestModificationRequest(),
    }
    typeEscapedValue := "request.TransferServiceAccountAdministrationRequest"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateRequestTransferServiceAccountAdministrationRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRequestTransferServiceAccountAdministrationRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequestTransferServiceAccountAdministrationRequest(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RequestTransferServiceAccountAdministrationRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RequestModificationRequest.GetFieldDeserializers()
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
// GetServiceAccount gets the serviceAccount property value. The serviceAccount property
func (m *RequestTransferServiceAccountAdministrationRequest) GetServiceAccount()(ServiceaccountServiceAccountPrimerable) {
    return m.serviceAccount
}
// Serialize serializes information the current object
func (m *RequestTransferServiceAccountAdministrationRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RequestModificationRequest.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("serviceAccount", m.GetServiceAccount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetServiceAccount sets the serviceAccount property value. The serviceAccount property
func (m *RequestTransferServiceAccountAdministrationRequest) SetServiceAccount(value ServiceaccountServiceAccountPrimerable)() {
    m.serviceAccount = value
}
// RequestTransferServiceAccountAdministrationRequestable 
type RequestTransferServiceAccountAdministrationRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RequestModificationRequestable
    GetServiceAccount()(ServiceaccountServiceAccountPrimerable)
    SetServiceAccount(value ServiceaccountServiceAccountPrimerable)()
}
