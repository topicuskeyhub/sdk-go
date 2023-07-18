package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WebhookWebhookNameUuid 
type WebhookWebhookNameUuid struct {
    NonLinkable
    // The name property
    name *string
    // The object property
    object Linkableable
    // The uuid property
    uuid *string
}
// NewWebhookWebhookNameUuid instantiates a new webhookWebhookNameUuid and sets the default values.
func NewWebhookWebhookNameUuid()(*WebhookWebhookNameUuid) {
    m := &WebhookWebhookNameUuid{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "webhook.WebhookNameUuid"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateWebhookWebhookNameUuidFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWebhookWebhookNameUuidFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWebhookWebhookNameUuid(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WebhookWebhookNameUuid) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["object"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLinkableFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetObject(val.(Linkableable))
        }
        return nil
    }
    res["uuid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUuid(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name property
func (m *WebhookWebhookNameUuid) GetName()(*string) {
    return m.name
}
// GetObject gets the object property value. The object property
func (m *WebhookWebhookNameUuid) GetObject()(Linkableable) {
    return m.object
}
// GetUuid gets the uuid property value. The uuid property
func (m *WebhookWebhookNameUuid) GetUuid()(*string) {
    return m.uuid
}
// Serialize serializes information the current object
func (m *WebhookWebhookNameUuid) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("object", m.GetObject())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("uuid", m.GetUuid())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetName sets the name property value. The name property
func (m *WebhookWebhookNameUuid) SetName(value *string)() {
    m.name = value
}
// SetObject sets the object property value. The object property
func (m *WebhookWebhookNameUuid) SetObject(value Linkableable)() {
    m.object = value
}
// SetUuid sets the uuid property value. The uuid property
func (m *WebhookWebhookNameUuid) SetUuid(value *string)() {
    m.uuid = value
}
// WebhookWebhookNameUuidable 
type WebhookWebhookNameUuidable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetName()(*string)
    GetObject()(Linkableable)
    GetUuid()(*string)
    SetName(value *string)()
    SetObject(value Linkableable)()
    SetUuid(value *string)()
}
