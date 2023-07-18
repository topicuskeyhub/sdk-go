package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WebhookWebhookLinkableWrapper 
type WebhookWebhookLinkableWrapper struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The items property
    items []WebhookWebhookable
}
// NewWebhookWebhookLinkableWrapper instantiates a new webhookWebhookLinkableWrapper and sets the default values.
func NewWebhookWebhookLinkableWrapper()(*WebhookWebhookLinkableWrapper) {
    m := &WebhookWebhookLinkableWrapper{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWebhookWebhookLinkableWrapperFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWebhookWebhookLinkableWrapperFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWebhookWebhookLinkableWrapper(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WebhookWebhookLinkableWrapper) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WebhookWebhookLinkableWrapper) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["items"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWebhookWebhookFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WebhookWebhookable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(WebhookWebhookable)
                }
            }
            m.SetItems(res)
        }
        return nil
    }
    return res
}
// GetItems gets the items property value. The items property
func (m *WebhookWebhookLinkableWrapper) GetItems()([]WebhookWebhookable) {
    return m.items
}
// Serialize serializes information the current object
func (m *WebhookWebhookLinkableWrapper) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetItems() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetItems()))
        for i, v := range m.GetItems() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("items", cast)
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WebhookWebhookLinkableWrapper) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetItems sets the items property value. The items property
func (m *WebhookWebhookLinkableWrapper) SetItems(value []WebhookWebhookable)() {
    m.items = value
}
// WebhookWebhookLinkableWrapperable 
type WebhookWebhookLinkableWrapperable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetItems()([]WebhookWebhookable)
    SetItems(value []WebhookWebhookable)()
}
