package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RestLink 
type RestLink struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The href property
    href *string
    // The id property
    id *int64
    // The rel property
    rel *string
    // The type property
    typeEscaped *string
}
// NewRestLink instantiates a new RestLink and sets the default values.
func NewRestLink()(*RestLink) {
    m := &RestLink{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRestLinkFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRestLinkFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRestLink(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RestLink) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RestLink) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["href"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHref(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["rel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRel(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val)
        }
        return nil
    }
    return res
}
// GetHref gets the href property value. The href property
func (m *RestLink) GetHref()(*string) {
    return m.href
}
// GetId gets the id property value. The id property
func (m *RestLink) GetId()(*int64) {
    return m.id
}
// GetRel gets the rel property value. The rel property
func (m *RestLink) GetRel()(*string) {
    return m.rel
}
// GetTypeEscaped gets the type property value. The type property
func (m *RestLink) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *RestLink) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("href", m.GetHref())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("rel", m.GetRel())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type", m.GetTypeEscaped())
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
func (m *RestLink) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetHref sets the href property value. The href property
func (m *RestLink) SetHref(value *string)() {
    m.href = value
}
// SetId sets the id property value. The id property
func (m *RestLink) SetId(value *int64)() {
    m.id = value
}
// SetRel sets the rel property value. The rel property
func (m *RestLink) SetRel(value *string)() {
    m.rel = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *RestLink) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// RestLinkable 
type RestLinkable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHref()(*string)
    GetId()(*int64)
    GetRel()(*string)
    GetTypeEscaped()(*string)
    SetHref(value *string)()
    SetId(value *int64)()
    SetRel(value *string)()
    SetTypeEscaped(value *string)()
}
