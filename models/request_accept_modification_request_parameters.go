package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RequestAcceptModificationRequestParameters struct {
    NonLinkable
}
// NewRequestAcceptModificationRequestParameters instantiates a new RequestAcceptModificationRequestParameters and sets the default values.
func NewRequestAcceptModificationRequestParameters()(*RequestAcceptModificationRequestParameters) {
    m := &RequestAcceptModificationRequestParameters{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "request.AcceptModificationRequestParameters"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateRequestAcceptModificationRequestParametersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRequestAcceptModificationRequestParametersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "request.AcceptCreateGroupOnSystemRequestParameters":
                        return NewRequestAcceptCreateGroupOnSystemRequestParameters(), nil
                    case "request.AcceptCreateGroupRequestParameters":
                        return NewRequestAcceptCreateGroupRequestParameters(), nil
                    case "request.AcceptCreateProvisionedNamespaceRequestParameters":
                        return NewRequestAcceptCreateProvisionedNamespaceRequestParameters(), nil
                    case "request.AcceptCreateServiceAccountRequestParameters":
                        return NewRequestAcceptCreateServiceAccountRequestParameters(), nil
                    case "request.AcceptGrantAccessRequestParameters":
                        return NewRequestAcceptGrantAccessRequestParameters(), nil
                    case "request.AcceptJoinGroupRequestParameters":
                        return NewRequestAcceptJoinGroupRequestParameters(), nil
                }
            }
        }
    }
    return NewRequestAcceptModificationRequestParameters(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RequestAcceptModificationRequestParameters) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *RequestAcceptModificationRequestParameters) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type RequestAcceptModificationRequestParametersable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
