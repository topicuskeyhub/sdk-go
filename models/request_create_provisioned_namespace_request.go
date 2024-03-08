package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RequestCreateProvisionedNamespaceRequest struct {
    RequestAbstractProvisionedSystemModificationRequest
    // The namespaceName property
    namespaceName *string
}
// NewRequestCreateProvisionedNamespaceRequest instantiates a new RequestCreateProvisionedNamespaceRequest and sets the default values.
func NewRequestCreateProvisionedNamespaceRequest()(*RequestCreateProvisionedNamespaceRequest) {
    m := &RequestCreateProvisionedNamespaceRequest{
        RequestAbstractProvisionedSystemModificationRequest: *NewRequestAbstractProvisionedSystemModificationRequest(),
    }
    typeEscapedValue := "request.CreateProvisionedNamespaceRequest"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateRequestCreateProvisionedNamespaceRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRequestCreateProvisionedNamespaceRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequestCreateProvisionedNamespaceRequest(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RequestCreateProvisionedNamespaceRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RequestAbstractProvisionedSystemModificationRequest.GetFieldDeserializers()
    res["namespaceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNamespaceName(val)
        }
        return nil
    }
    return res
}
// GetNamespaceName gets the namespaceName property value. The namespaceName property
// returns a *string when successful
func (m *RequestCreateProvisionedNamespaceRequest) GetNamespaceName()(*string) {
    return m.namespaceName
}
// Serialize serializes information the current object
func (m *RequestCreateProvisionedNamespaceRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RequestAbstractProvisionedSystemModificationRequest.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("namespaceName", m.GetNamespaceName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetNamespaceName sets the namespaceName property value. The namespaceName property
func (m *RequestCreateProvisionedNamespaceRequest) SetNamespaceName(value *string)() {
    m.namespaceName = value
}
type RequestCreateProvisionedNamespaceRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RequestAbstractProvisionedSystemModificationRequestable
    GetNamespaceName()(*string)
    SetNamespaceName(value *string)()
}
