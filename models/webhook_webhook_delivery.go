package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WebhookWebhookDelivery 
type WebhookWebhookDelivery struct {
    Linkable
    // The additionalObjects property
    additionalObjects WebhookWebhookDelivery_additionalObjectsable
    // The payload property
    payload WebhookWebhookPushable
    // The reponseHeaders property
    reponseHeaders *string
    // The requestHeaders property
    requestHeaders *string
    // The response property
    response *string
    // The status property
    status *int32
    // The time property
    time *int32
    // The triggerTime property
    triggerTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewWebhookWebhookDelivery instantiates a new webhookWebhookDelivery and sets the default values.
func NewWebhookWebhookDelivery()(*WebhookWebhookDelivery) {
    m := &WebhookWebhookDelivery{
        Linkable: *NewLinkable(),
    }
    typeEscapedValue := "webhook.WebhookDelivery"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateWebhookWebhookDeliveryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWebhookWebhookDeliveryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWebhookWebhookDelivery(), nil
}
// GetAdditionalObjects gets the additionalObjects property value. The additionalObjects property
func (m *WebhookWebhookDelivery) GetAdditionalObjects()(WebhookWebhookDelivery_additionalObjectsable) {
    return m.additionalObjects
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WebhookWebhookDelivery) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Linkable.GetFieldDeserializers()
    res["additionalObjects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWebhookWebhookDelivery_additionalObjectsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalObjects(val.(WebhookWebhookDelivery_additionalObjectsable))
        }
        return nil
    }
    res["payload"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWebhookWebhookPushFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayload(val.(WebhookWebhookPushable))
        }
        return nil
    }
    res["reponseHeaders"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReponseHeaders(val)
        }
        return nil
    }
    res["requestHeaders"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestHeaders(val)
        }
        return nil
    }
    res["response"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResponse(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    res["time"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTime(val)
        }
        return nil
    }
    res["triggerTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTriggerTime(val)
        }
        return nil
    }
    return res
}
// GetPayload gets the payload property value. The payload property
func (m *WebhookWebhookDelivery) GetPayload()(WebhookWebhookPushable) {
    return m.payload
}
// GetReponseHeaders gets the reponseHeaders property value. The reponseHeaders property
func (m *WebhookWebhookDelivery) GetReponseHeaders()(*string) {
    return m.reponseHeaders
}
// GetRequestHeaders gets the requestHeaders property value. The requestHeaders property
func (m *WebhookWebhookDelivery) GetRequestHeaders()(*string) {
    return m.requestHeaders
}
// GetResponse gets the response property value. The response property
func (m *WebhookWebhookDelivery) GetResponse()(*string) {
    return m.response
}
// GetStatus gets the status property value. The status property
func (m *WebhookWebhookDelivery) GetStatus()(*int32) {
    return m.status
}
// GetTime gets the time property value. The time property
func (m *WebhookWebhookDelivery) GetTime()(*int32) {
    return m.time
}
// GetTriggerTime gets the triggerTime property value. The triggerTime property
func (m *WebhookWebhookDelivery) GetTriggerTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.triggerTime
}
// Serialize serializes information the current object
func (m *WebhookWebhookDelivery) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Linkable.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("additionalObjects", m.GetAdditionalObjects())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("payload", m.GetPayload())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("reponseHeaders", m.GetReponseHeaders())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("requestHeaders", m.GetRequestHeaders())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("response", m.GetResponse())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("time", m.GetTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("triggerTime", m.GetTriggerTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalObjects sets the additionalObjects property value. The additionalObjects property
func (m *WebhookWebhookDelivery) SetAdditionalObjects(value WebhookWebhookDelivery_additionalObjectsable)() {
    m.additionalObjects = value
}
// SetPayload sets the payload property value. The payload property
func (m *WebhookWebhookDelivery) SetPayload(value WebhookWebhookPushable)() {
    m.payload = value
}
// SetReponseHeaders sets the reponseHeaders property value. The reponseHeaders property
func (m *WebhookWebhookDelivery) SetReponseHeaders(value *string)() {
    m.reponseHeaders = value
}
// SetRequestHeaders sets the requestHeaders property value. The requestHeaders property
func (m *WebhookWebhookDelivery) SetRequestHeaders(value *string)() {
    m.requestHeaders = value
}
// SetResponse sets the response property value. The response property
func (m *WebhookWebhookDelivery) SetResponse(value *string)() {
    m.response = value
}
// SetStatus sets the status property value. The status property
func (m *WebhookWebhookDelivery) SetStatus(value *int32)() {
    m.status = value
}
// SetTime sets the time property value. The time property
func (m *WebhookWebhookDelivery) SetTime(value *int32)() {
    m.time = value
}
// SetTriggerTime sets the triggerTime property value. The triggerTime property
func (m *WebhookWebhookDelivery) SetTriggerTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.triggerTime = value
}
// WebhookWebhookDeliveryable 
type WebhookWebhookDeliveryable interface {
    Linkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdditionalObjects()(WebhookWebhookDelivery_additionalObjectsable)
    GetPayload()(WebhookWebhookPushable)
    GetReponseHeaders()(*string)
    GetRequestHeaders()(*string)
    GetResponse()(*string)
    GetStatus()(*int32)
    GetTime()(*int32)
    GetTriggerTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetAdditionalObjects(value WebhookWebhookDelivery_additionalObjectsable)()
    SetPayload(value WebhookWebhookPushable)()
    SetReponseHeaders(value *string)()
    SetRequestHeaders(value *string)()
    SetResponse(value *string)()
    SetStatus(value *int32)()
    SetTime(value *int32)()
    SetTriggerTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
