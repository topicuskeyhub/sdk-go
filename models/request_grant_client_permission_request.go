package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RequestGrantClientPermissionRequest 
type RequestGrantClientPermissionRequest struct {
    RequestAbstractApplicationModificationRequest
    // The permissionType property
    permissionType *ClientOAuth2ClientPermissionType
    // The system property
    system ProvisioningProvisionedSystemPrimerable
}
// NewRequestGrantClientPermissionRequest instantiates a new requestGrantClientPermissionRequest and sets the default values.
func NewRequestGrantClientPermissionRequest()(*RequestGrantClientPermissionRequest) {
    m := &RequestGrantClientPermissionRequest{
        RequestAbstractApplicationModificationRequest: *NewRequestAbstractApplicationModificationRequest(),
    }
    typeEscapedValue := "request.GrantClientPermissionRequest"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateRequestGrantClientPermissionRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRequestGrantClientPermissionRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequestGrantClientPermissionRequest(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RequestGrantClientPermissionRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RequestAbstractApplicationModificationRequest.GetFieldDeserializers()
    res["permissionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseClientOAuth2ClientPermissionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermissionType(val.(*ClientOAuth2ClientPermissionType))
        }
        return nil
    }
    res["system"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProvisioningProvisionedSystemPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystem(val.(ProvisioningProvisionedSystemPrimerable))
        }
        return nil
    }
    return res
}
// GetPermissionType gets the permissionType property value. The permissionType property
func (m *RequestGrantClientPermissionRequest) GetPermissionType()(*ClientOAuth2ClientPermissionType) {
    return m.permissionType
}
// GetSystem gets the system property value. The system property
func (m *RequestGrantClientPermissionRequest) GetSystem()(ProvisioningProvisionedSystemPrimerable) {
    return m.system
}
// Serialize serializes information the current object
func (m *RequestGrantClientPermissionRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RequestAbstractApplicationModificationRequest.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetPermissionType() != nil {
        cast := (*m.GetPermissionType()).String()
        err = writer.WriteStringValue("permissionType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("system", m.GetSystem())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetPermissionType sets the permissionType property value. The permissionType property
func (m *RequestGrantClientPermissionRequest) SetPermissionType(value *ClientOAuth2ClientPermissionType)() {
    m.permissionType = value
}
// SetSystem sets the system property value. The system property
func (m *RequestGrantClientPermissionRequest) SetSystem(value ProvisioningProvisionedSystemPrimerable)() {
    m.system = value
}
// RequestGrantClientPermissionRequestable 
type RequestGrantClientPermissionRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RequestAbstractApplicationModificationRequestable
    GetPermissionType()(*ClientOAuth2ClientPermissionType)
    GetSystem()(ProvisioningProvisionedSystemPrimerable)
    SetPermissionType(value *ClientOAuth2ClientPermissionType)()
    SetSystem(value ProvisioningProvisionedSystemPrimerable)()
}
