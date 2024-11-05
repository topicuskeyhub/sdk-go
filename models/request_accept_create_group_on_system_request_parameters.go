package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RequestAcceptCreateGroupOnSystemRequestParameters struct {
    RequestAcceptModificationRequestParameters
    // The updatedName property
    updatedName *string
}
// NewRequestAcceptCreateGroupOnSystemRequestParameters instantiates a new RequestAcceptCreateGroupOnSystemRequestParameters and sets the default values.
func NewRequestAcceptCreateGroupOnSystemRequestParameters()(*RequestAcceptCreateGroupOnSystemRequestParameters) {
    m := &RequestAcceptCreateGroupOnSystemRequestParameters{
        RequestAcceptModificationRequestParameters: *NewRequestAcceptModificationRequestParameters(),
    }
    typeEscapedValue := "request.AcceptCreateGroupOnSystemRequestParameters"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateRequestAcceptCreateGroupOnSystemRequestParametersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRequestAcceptCreateGroupOnSystemRequestParametersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequestAcceptCreateGroupOnSystemRequestParameters(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RequestAcceptCreateGroupOnSystemRequestParameters) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *RequestAcceptCreateGroupOnSystemRequestParameters) GetUpdatedName()(*string) {
    return m.updatedName
}
// Serialize serializes information the current object
func (m *RequestAcceptCreateGroupOnSystemRequestParameters) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *RequestAcceptCreateGroupOnSystemRequestParameters) SetUpdatedName(value *string)() {
    m.updatedName = value
}
type RequestAcceptCreateGroupOnSystemRequestParametersable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RequestAcceptModificationRequestParametersable
    GetUpdatedName()(*string)
    SetUpdatedName(value *string)()
}
