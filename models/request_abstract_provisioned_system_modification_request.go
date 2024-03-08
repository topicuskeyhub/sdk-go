package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RequestAbstractProvisionedSystemModificationRequest struct {
    RequestModificationRequest
    // The system property
    system ProvisioningProvisionedSystemPrimerable
}
// NewRequestAbstractProvisionedSystemModificationRequest instantiates a new RequestAbstractProvisionedSystemModificationRequest and sets the default values.
func NewRequestAbstractProvisionedSystemModificationRequest()(*RequestAbstractProvisionedSystemModificationRequest) {
    m := &RequestAbstractProvisionedSystemModificationRequest{
        RequestModificationRequest: *NewRequestModificationRequest(),
    }
    typeEscapedValue := "request.AbstractProvisionedSystemModificationRequest"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateRequestAbstractProvisionedSystemModificationRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRequestAbstractProvisionedSystemModificationRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("$type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "request.CreateGroupOnSystemRequest":
                        return NewRequestCreateGroupOnSystemRequest(), nil
                    case "request.CreateProvisionedNamespaceRequest":
                        return NewRequestCreateProvisionedNamespaceRequest(), nil
                    case "request.CreateServiceAccountRequest":
                        return NewRequestCreateServiceAccountRequest(), nil
                    case "request.RemoveProvisionedSystemRequest":
                        return NewRequestRemoveProvisionedSystemRequest(), nil
                    case "request.TransferProvisionedSystemAdministrationRequest":
                        return NewRequestTransferProvisionedSystemAdministrationRequest(), nil
                    case "request.TransferProvisionedSystemContentAdministrationRequest":
                        return NewRequestTransferProvisionedSystemContentAdministrationRequest(), nil
                    case "request.TransferProvisionedSystemOwnershipRequest":
                        return NewRequestTransferProvisionedSystemOwnershipRequest(), nil
                }
            }
        }
    }
    return NewRequestAbstractProvisionedSystemModificationRequest(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RequestAbstractProvisionedSystemModificationRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RequestModificationRequest.GetFieldDeserializers()
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
// GetSystem gets the system property value. The system property
// returns a ProvisioningProvisionedSystemPrimerable when successful
func (m *RequestAbstractProvisionedSystemModificationRequest) GetSystem()(ProvisioningProvisionedSystemPrimerable) {
    return m.system
}
// Serialize serializes information the current object
func (m *RequestAbstractProvisionedSystemModificationRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RequestModificationRequest.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("system", m.GetSystem())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSystem sets the system property value. The system property
func (m *RequestAbstractProvisionedSystemModificationRequest) SetSystem(value ProvisioningProvisionedSystemPrimerable)() {
    m.system = value
}
type RequestAbstractProvisionedSystemModificationRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RequestModificationRequestable
    GetSystem()(ProvisioningProvisionedSystemPrimerable)
    SetSystem(value ProvisioningProvisionedSystemPrimerable)()
}
