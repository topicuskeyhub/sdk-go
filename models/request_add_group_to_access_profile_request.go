// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RequestAddGroupToAccessProfileRequest struct {
    RequestAbstractAccessProfileModificationRequest
}
// NewRequestAddGroupToAccessProfileRequest instantiates a new RequestAddGroupToAccessProfileRequest and sets the default values.
func NewRequestAddGroupToAccessProfileRequest()(*RequestAddGroupToAccessProfileRequest) {
    m := &RequestAddGroupToAccessProfileRequest{
        RequestAbstractAccessProfileModificationRequest: *NewRequestAbstractAccessProfileModificationRequest(),
    }
    typeEscapedValue := "request.AddGroupToAccessProfileRequest"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateRequestAddGroupToAccessProfileRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRequestAddGroupToAccessProfileRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequestAddGroupToAccessProfileRequest(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RequestAddGroupToAccessProfileRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RequestAbstractAccessProfileModificationRequest.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *RequestAddGroupToAccessProfileRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RequestAbstractAccessProfileModificationRequest.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type RequestAddGroupToAccessProfileRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RequestAbstractAccessProfileModificationRequestable
}
