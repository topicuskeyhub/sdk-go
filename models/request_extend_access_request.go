package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RequestExtendAccessRequest 
type RequestExtendAccessRequest struct {
    RequestModificationRequest
    // The extendUntil property
    extendUntil *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewRequestExtendAccessRequest instantiates a new requestExtendAccessRequest and sets the default values.
func NewRequestExtendAccessRequest()(*RequestExtendAccessRequest) {
    m := &RequestExtendAccessRequest{
        RequestModificationRequest: *NewRequestModificationRequest(),
    }
    typeEscapedValue := "request.ExtendAccessRequest"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateRequestExtendAccessRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRequestExtendAccessRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequestExtendAccessRequest(), nil
}
// GetExtendUntil gets the extendUntil property value. The extendUntil property
func (m *RequestExtendAccessRequest) GetExtendUntil()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.extendUntil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RequestExtendAccessRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RequestModificationRequest.GetFieldDeserializers()
    res["extendUntil"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExtendUntil(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *RequestExtendAccessRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RequestModificationRequest.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("extendUntil", m.GetExtendUntil())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetExtendUntil sets the extendUntil property value. The extendUntil property
func (m *RequestExtendAccessRequest) SetExtendUntil(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.extendUntil = value
}
// RequestExtendAccessRequestable 
type RequestExtendAccessRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RequestModificationRequestable
    GetExtendUntil()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetExtendUntil(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
