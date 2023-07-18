package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RequestAcceptJoinGroupRequestParameters 
type RequestAcceptJoinGroupRequestParameters struct {
    RequestAcceptModificationRequestParameters
    // The endDate property
    endDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The makeManager property
    makeManager *bool
}
// NewRequestAcceptJoinGroupRequestParameters instantiates a new requestAcceptJoinGroupRequestParameters and sets the default values.
func NewRequestAcceptJoinGroupRequestParameters()(*RequestAcceptJoinGroupRequestParameters) {
    m := &RequestAcceptJoinGroupRequestParameters{
        RequestAcceptModificationRequestParameters: *NewRequestAcceptModificationRequestParameters(),
    }
    typeEscapedValue := "request.AcceptJoinGroupRequestParameters"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateRequestAcceptJoinGroupRequestParametersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRequestAcceptJoinGroupRequestParametersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequestAcceptJoinGroupRequestParameters(), nil
}
// GetEndDate gets the endDate property value. The endDate property
func (m *RequestAcceptJoinGroupRequestParameters) GetEndDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.endDate
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RequestAcceptJoinGroupRequestParameters) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RequestAcceptModificationRequestParameters.GetFieldDeserializers()
    res["endDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDate(val)
        }
        return nil
    }
    res["makeManager"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMakeManager(val)
        }
        return nil
    }
    return res
}
// GetMakeManager gets the makeManager property value. The makeManager property
func (m *RequestAcceptJoinGroupRequestParameters) GetMakeManager()(*bool) {
    return m.makeManager
}
// Serialize serializes information the current object
func (m *RequestAcceptJoinGroupRequestParameters) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RequestAcceptModificationRequestParameters.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteDateOnlyValue("endDate", m.GetEndDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("makeManager", m.GetMakeManager())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEndDate sets the endDate property value. The endDate property
func (m *RequestAcceptJoinGroupRequestParameters) SetEndDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.endDate = value
}
// SetMakeManager sets the makeManager property value. The makeManager property
func (m *RequestAcceptJoinGroupRequestParameters) SetMakeManager(value *bool)() {
    m.makeManager = value
}
// RequestAcceptJoinGroupRequestParametersable 
type RequestAcceptJoinGroupRequestParametersable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RequestAcceptModificationRequestParametersable
    GetEndDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetMakeManager()(*bool)
    SetEndDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetMakeManager(value *bool)()
}
