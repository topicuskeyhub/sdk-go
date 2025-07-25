// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RequestAbstractAccessProfileModificationRequest struct {
    RequestModificationRequest
    // The accessProfile property
    accessProfile ProfileAccessProfilePrimerable
}
// NewRequestAbstractAccessProfileModificationRequest instantiates a new RequestAbstractAccessProfileModificationRequest and sets the default values.
func NewRequestAbstractAccessProfileModificationRequest()(*RequestAbstractAccessProfileModificationRequest) {
    m := &RequestAbstractAccessProfileModificationRequest{
        RequestModificationRequest: *NewRequestModificationRequest(),
    }
    typeEscapedValue := "request.AbstractAccessProfileModificationRequest"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateRequestAbstractAccessProfileModificationRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRequestAbstractAccessProfileModificationRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "request.AddGroupToAccessProfileRequest":
                        return NewRequestAddGroupToAccessProfileRequest(), nil
                    case "request.GrantApplicationRequest":
                        return NewRequestGrantApplicationRequest(), nil
                    case "request.GrantGroupOnSystemRequest":
                        return NewRequestGrantGroupOnSystemRequest(), nil
                    case "request.GrantGroupOnSystemRequestRequest":
                        return NewRequestGrantGroupOnSystemRequestRequest(), nil
                    case "request.JoinGroupRequest":
                        return NewRequestJoinGroupRequest(), nil
                    case "request.LinkDirectoryToAccessProfileRequest":
                        return NewRequestLinkDirectoryToAccessProfileRequest(), nil
                    case "request.RemoveAccessProfileRequest":
                        return NewRequestRemoveAccessProfileRequest(), nil
                    case "request.TransferAccessProfileOwnershipRequest":
                        return NewRequestTransferAccessProfileOwnershipRequest(), nil
                }
            }
        }
    }
    return NewRequestAbstractAccessProfileModificationRequest(), nil
}
// GetAccessProfile gets the accessProfile property value. The accessProfile property
// returns a ProfileAccessProfilePrimerable when successful
func (m *RequestAbstractAccessProfileModificationRequest) GetAccessProfile()(ProfileAccessProfilePrimerable) {
    return m.accessProfile
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RequestAbstractAccessProfileModificationRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RequestModificationRequest.GetFieldDeserializers()
    res["accessProfile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProfileAccessProfilePrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessProfile(val.(ProfileAccessProfilePrimerable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *RequestAbstractAccessProfileModificationRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RequestModificationRequest.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("accessProfile", m.GetAccessProfile())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessProfile sets the accessProfile property value. The accessProfile property
func (m *RequestAbstractAccessProfileModificationRequest) SetAccessProfile(value ProfileAccessProfilePrimerable)() {
    m.accessProfile = value
}
type RequestAbstractAccessProfileModificationRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RequestModificationRequestable
    GetAccessProfile()(ProfileAccessProfilePrimerable)
    SetAccessProfile(value ProfileAccessProfilePrimerable)()
}
