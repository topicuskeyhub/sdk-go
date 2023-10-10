package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RequestAcceptGrantAccessRequestParameters_accessDuration 
type RequestAcceptGrantAccessRequestParameters_accessDuration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The nanos property
    nanos *int32
    // The seconds property
    seconds *int64
}
// NewRequestAcceptGrantAccessRequestParameters_accessDuration instantiates a new requestAcceptGrantAccessRequestParameters_accessDuration and sets the default values.
func NewRequestAcceptGrantAccessRequestParameters_accessDuration()(*RequestAcceptGrantAccessRequestParameters_accessDuration) {
    m := &RequestAcceptGrantAccessRequestParameters_accessDuration{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRequestAcceptGrantAccessRequestParameters_accessDurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRequestAcceptGrantAccessRequestParameters_accessDurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequestAcceptGrantAccessRequestParameters_accessDuration(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RequestAcceptGrantAccessRequestParameters_accessDuration) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RequestAcceptGrantAccessRequestParameters_accessDuration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["nanos"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNanos(val)
        }
        return nil
    }
    res["seconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSeconds(val)
        }
        return nil
    }
    return res
}
// GetNanos gets the nanos property value. The nanos property
func (m *RequestAcceptGrantAccessRequestParameters_accessDuration) GetNanos()(*int32) {
    return m.nanos
}
// GetSeconds gets the seconds property value. The seconds property
func (m *RequestAcceptGrantAccessRequestParameters_accessDuration) GetSeconds()(*int64) {
    return m.seconds
}
// Serialize serializes information the current object
func (m *RequestAcceptGrantAccessRequestParameters_accessDuration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("nanos", m.GetNanos())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("seconds", m.GetSeconds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RequestAcceptGrantAccessRequestParameters_accessDuration) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetNanos sets the nanos property value. The nanos property
func (m *RequestAcceptGrantAccessRequestParameters_accessDuration) SetNanos(value *int32)() {
    m.nanos = value
}
// SetSeconds sets the seconds property value. The seconds property
func (m *RequestAcceptGrantAccessRequestParameters_accessDuration) SetSeconds(value *int64)() {
    m.seconds = value
}
// RequestAcceptGrantAccessRequestParameters_accessDurationable 
type RequestAcceptGrantAccessRequestParameters_accessDurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetNanos()(*int32)
    GetSeconds()(*int64)
    SetNanos(value *int32)()
    SetSeconds(value *int64)()
}
