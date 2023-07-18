package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RequestRemoveOrganizationalUnitRequest 
type RequestRemoveOrganizationalUnitRequest struct {
    RequestAbstractOrganizationalUnitModificationRequest
    // The organizationalUnitName property
    organizationalUnitName *string
}
// NewRequestRemoveOrganizationalUnitRequest instantiates a new requestRemoveOrganizationalUnitRequest and sets the default values.
func NewRequestRemoveOrganizationalUnitRequest()(*RequestRemoveOrganizationalUnitRequest) {
    m := &RequestRemoveOrganizationalUnitRequest{
        RequestAbstractOrganizationalUnitModificationRequest: *NewRequestAbstractOrganizationalUnitModificationRequest(),
    }
    typeEscapedValue := "request.RemoveOrganizationalUnitRequest"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateRequestRemoveOrganizationalUnitRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRequestRemoveOrganizationalUnitRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequestRemoveOrganizationalUnitRequest(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RequestRemoveOrganizationalUnitRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RequestAbstractOrganizationalUnitModificationRequest.GetFieldDeserializers()
    res["organizationalUnitName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationalUnitName(val)
        }
        return nil
    }
    return res
}
// GetOrganizationalUnitName gets the organizationalUnitName property value. The organizationalUnitName property
func (m *RequestRemoveOrganizationalUnitRequest) GetOrganizationalUnitName()(*string) {
    return m.organizationalUnitName
}
// Serialize serializes information the current object
func (m *RequestRemoveOrganizationalUnitRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RequestAbstractOrganizationalUnitModificationRequest.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("organizationalUnitName", m.GetOrganizationalUnitName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOrganizationalUnitName sets the organizationalUnitName property value. The organizationalUnitName property
func (m *RequestRemoveOrganizationalUnitRequest) SetOrganizationalUnitName(value *string)() {
    m.organizationalUnitName = value
}
// RequestRemoveOrganizationalUnitRequestable 
type RequestRemoveOrganizationalUnitRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RequestAbstractOrganizationalUnitModificationRequestable
    GetOrganizationalUnitName()(*string)
    SetOrganizationalUnitName(value *string)()
}
