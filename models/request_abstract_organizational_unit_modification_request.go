package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RequestAbstractOrganizationalUnitModificationRequest struct {
    RequestModificationRequest
    // The organizationalUnit property
    organizationalUnit OrganizationOrganizationalUnitPrimerable
}
// NewRequestAbstractOrganizationalUnitModificationRequest instantiates a new RequestAbstractOrganizationalUnitModificationRequest and sets the default values.
func NewRequestAbstractOrganizationalUnitModificationRequest()(*RequestAbstractOrganizationalUnitModificationRequest) {
    m := &RequestAbstractOrganizationalUnitModificationRequest{
        RequestModificationRequest: *NewRequestModificationRequest(),
    }
    typeEscapedValue := "request.AbstractOrganizationalUnitModificationRequest"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateRequestAbstractOrganizationalUnitModificationRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRequestAbstractOrganizationalUnitModificationRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "request.CreateGroupRequest":
                        return NewRequestCreateGroupRequest(), nil
                    case "request.MoveGroupsRequest":
                        return NewRequestMoveGroupsRequest(), nil
                    case "request.RemoveOrganizationalUnitRequest":
                        return NewRequestRemoveOrganizationalUnitRequest(), nil
                    case "request.TransferOrganizationalUnitOwnershipRequest":
                        return NewRequestTransferOrganizationalUnitOwnershipRequest(), nil
                }
            }
        }
    }
    return NewRequestAbstractOrganizationalUnitModificationRequest(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RequestAbstractOrganizationalUnitModificationRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RequestModificationRequest.GetFieldDeserializers()
    res["organizationalUnit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOrganizationOrganizationalUnitPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationalUnit(val.(OrganizationOrganizationalUnitPrimerable))
        }
        return nil
    }
    return res
}
// GetOrganizationalUnit gets the organizationalUnit property value. The organizationalUnit property
// returns a OrganizationOrganizationalUnitPrimerable when successful
func (m *RequestAbstractOrganizationalUnitModificationRequest) GetOrganizationalUnit()(OrganizationOrganizationalUnitPrimerable) {
    return m.organizationalUnit
}
// Serialize serializes information the current object
func (m *RequestAbstractOrganizationalUnitModificationRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RequestModificationRequest.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("organizationalUnit", m.GetOrganizationalUnit())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOrganizationalUnit sets the organizationalUnit property value. The organizationalUnit property
func (m *RequestAbstractOrganizationalUnitModificationRequest) SetOrganizationalUnit(value OrganizationOrganizationalUnitPrimerable)() {
    m.organizationalUnit = value
}
type RequestAbstractOrganizationalUnitModificationRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RequestModificationRequestable
    GetOrganizationalUnit()(OrganizationOrganizationalUnitPrimerable)
    SetOrganizationalUnit(value OrganizationOrganizationalUnitPrimerable)()
}
