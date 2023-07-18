package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RequestAcceptGrantAccessRequestParameters 
type RequestAcceptGrantAccessRequestParameters struct {
    RequestAcceptModificationRequestParameters
    // The accessDuration property
    accessDuration RequestAcceptGrantAccessRequestParameters_accessDurationable
}
// NewRequestAcceptGrantAccessRequestParameters instantiates a new requestAcceptGrantAccessRequestParameters and sets the default values.
func NewRequestAcceptGrantAccessRequestParameters()(*RequestAcceptGrantAccessRequestParameters) {
    m := &RequestAcceptGrantAccessRequestParameters{
        RequestAcceptModificationRequestParameters: *NewRequestAcceptModificationRequestParameters(),
    }
    typeEscapedValue := "request.AcceptGrantAccessRequestParameters"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateRequestAcceptGrantAccessRequestParametersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRequestAcceptGrantAccessRequestParametersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequestAcceptGrantAccessRequestParameters(), nil
}
// GetAccessDuration gets the accessDuration property value. The accessDuration property
func (m *RequestAcceptGrantAccessRequestParameters) GetAccessDuration()(RequestAcceptGrantAccessRequestParameters_accessDurationable) {
    return m.accessDuration
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RequestAcceptGrantAccessRequestParameters) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RequestAcceptModificationRequestParameters.GetFieldDeserializers()
    res["accessDuration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRequestAcceptGrantAccessRequestParameters_accessDurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessDuration(val.(RequestAcceptGrantAccessRequestParameters_accessDurationable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *RequestAcceptGrantAccessRequestParameters) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RequestAcceptModificationRequestParameters.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("accessDuration", m.GetAccessDuration())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessDuration sets the accessDuration property value. The accessDuration property
func (m *RequestAcceptGrantAccessRequestParameters) SetAccessDuration(value RequestAcceptGrantAccessRequestParameters_accessDurationable)() {
    m.accessDuration = value
}
// RequestAcceptGrantAccessRequestParametersable 
type RequestAcceptGrantAccessRequestParametersable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RequestAcceptModificationRequestParametersable
    GetAccessDuration()(RequestAcceptGrantAccessRequestParameters_accessDurationable)
    SetAccessDuration(value RequestAcceptGrantAccessRequestParameters_accessDurationable)()
}
