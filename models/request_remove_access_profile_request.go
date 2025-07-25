// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RequestRemoveAccessProfileRequest struct {
    RequestAbstractAccessProfileModificationRequest
    // The accessProfileName property
    accessProfileName *string
}
// NewRequestRemoveAccessProfileRequest instantiates a new RequestRemoveAccessProfileRequest and sets the default values.
func NewRequestRemoveAccessProfileRequest()(*RequestRemoveAccessProfileRequest) {
    m := &RequestRemoveAccessProfileRequest{
        RequestAbstractAccessProfileModificationRequest: *NewRequestAbstractAccessProfileModificationRequest(),
    }
    typeEscapedValue := "request.RemoveAccessProfileRequest"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateRequestRemoveAccessProfileRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRequestRemoveAccessProfileRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequestRemoveAccessProfileRequest(), nil
}
// GetAccessProfileName gets the accessProfileName property value. The accessProfileName property
// returns a *string when successful
func (m *RequestRemoveAccessProfileRequest) GetAccessProfileName()(*string) {
    return m.accessProfileName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RequestRemoveAccessProfileRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RequestAbstractAccessProfileModificationRequest.GetFieldDeserializers()
    res["accessProfileName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessProfileName(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *RequestRemoveAccessProfileRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RequestAbstractAccessProfileModificationRequest.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("accessProfileName", m.GetAccessProfileName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessProfileName sets the accessProfileName property value. The accessProfileName property
func (m *RequestRemoveAccessProfileRequest) SetAccessProfileName(value *string)() {
    m.accessProfileName = value
}
type RequestRemoveAccessProfileRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RequestAbstractAccessProfileModificationRequestable
    GetAccessProfileName()(*string)
    SetAccessProfileName(value *string)()
}
