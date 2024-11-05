package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RequestAcceptCreateGroupRequestParameters struct {
    RequestAcceptModificationRequestParameters
    // The updatedName property
    updatedName *string
}
// NewRequestAcceptCreateGroupRequestParameters instantiates a new RequestAcceptCreateGroupRequestParameters and sets the default values.
func NewRequestAcceptCreateGroupRequestParameters()(*RequestAcceptCreateGroupRequestParameters) {
    m := &RequestAcceptCreateGroupRequestParameters{
        RequestAcceptModificationRequestParameters: *NewRequestAcceptModificationRequestParameters(),
    }
    typeEscapedValue := "request.AcceptCreateGroupRequestParameters"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateRequestAcceptCreateGroupRequestParametersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRequestAcceptCreateGroupRequestParametersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequestAcceptCreateGroupRequestParameters(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RequestAcceptCreateGroupRequestParameters) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RequestAcceptModificationRequestParameters.GetFieldDeserializers()
    res["updatedName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdatedName(val)
        }
        return nil
    }
    return res
}
// GetUpdatedName gets the updatedName property value. The updatedName property
// returns a *string when successful
func (m *RequestAcceptCreateGroupRequestParameters) GetUpdatedName()(*string) {
    return m.updatedName
}
// Serialize serializes information the current object
func (m *RequestAcceptCreateGroupRequestParameters) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RequestAcceptModificationRequestParameters.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("updatedName", m.GetUpdatedName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetUpdatedName sets the updatedName property value. The updatedName property
func (m *RequestAcceptCreateGroupRequestParameters) SetUpdatedName(value *string)() {
    m.updatedName = value
}
type RequestAcceptCreateGroupRequestParametersable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RequestAcceptModificationRequestParametersable
    GetUpdatedName()(*string)
    SetUpdatedName(value *string)()
}