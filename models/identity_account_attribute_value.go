package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type IdentityAccountAttributeValue struct {
    Linkable
    // The additionalObjects property
    additionalObjects IdentityAccountAttributeValue_additionalObjectsable
    // The attribute property
    attribute IdentityAccountAttributeDefinitionable
    // The context property
    context *string
    // The date property
    date *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The latestForSource property
    latestForSource *bool
    // The source property
    source *IdentityAccountAttributeSource
    // The status property
    status *IdentityAccountAttributeValueStatus
    // The value property
    value *string
}
// NewIdentityAccountAttributeValue instantiates a new IdentityAccountAttributeValue and sets the default values.
func NewIdentityAccountAttributeValue()(*IdentityAccountAttributeValue) {
    m := &IdentityAccountAttributeValue{
        Linkable: *NewLinkable(),
    }
    typeEscapedValue := "identity.AccountAttributeValue"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateIdentityAccountAttributeValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIdentityAccountAttributeValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIdentityAccountAttributeValue(), nil
}
// GetAdditionalObjects gets the additionalObjects property value. The additionalObjects property
// returns a IdentityAccountAttributeValue_additionalObjectsable when successful
func (m *IdentityAccountAttributeValue) GetAdditionalObjects()(IdentityAccountAttributeValue_additionalObjectsable) {
    return m.additionalObjects
}
// GetAttribute gets the attribute property value. The attribute property
// returns a IdentityAccountAttributeDefinitionable when successful
func (m *IdentityAccountAttributeValue) GetAttribute()(IdentityAccountAttributeDefinitionable) {
    return m.attribute
}
// GetContext gets the context property value. The context property
// returns a *string when successful
func (m *IdentityAccountAttributeValue) GetContext()(*string) {
    return m.context
}
// GetDate gets the date property value. The date property
// returns a *Time when successful
func (m *IdentityAccountAttributeValue) GetDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.date
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *IdentityAccountAttributeValue) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Linkable.GetFieldDeserializers()
    res["additionalObjects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentityAccountAttributeValue_additionalObjectsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalObjects(val.(IdentityAccountAttributeValue_additionalObjectsable))
        }
        return nil
    }
    res["attribute"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentityAccountAttributeDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttribute(val.(IdentityAccountAttributeDefinitionable))
        }
        return nil
    }
    res["context"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContext(val)
        }
        return nil
    }
    res["date"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDate(val)
        }
        return nil
    }
    res["latestForSource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLatestForSource(val)
        }
        return nil
    }
    res["source"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseIdentityAccountAttributeSource)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSource(val.(*IdentityAccountAttributeSource))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseIdentityAccountAttributeValueStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*IdentityAccountAttributeValueStatus))
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetLatestForSource gets the latestForSource property value. The latestForSource property
// returns a *bool when successful
func (m *IdentityAccountAttributeValue) GetLatestForSource()(*bool) {
    return m.latestForSource
}
// GetSource gets the source property value. The source property
// returns a *IdentityAccountAttributeSource when successful
func (m *IdentityAccountAttributeValue) GetSource()(*IdentityAccountAttributeSource) {
    return m.source
}
// GetStatus gets the status property value. The status property
// returns a *IdentityAccountAttributeValueStatus when successful
func (m *IdentityAccountAttributeValue) GetStatus()(*IdentityAccountAttributeValueStatus) {
    return m.status
}
// GetValue gets the value property value. The value property
// returns a *string when successful
func (m *IdentityAccountAttributeValue) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *IdentityAccountAttributeValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteObjectValue("attribute", m.GetAttribute())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("context", m.GetContext())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("date", m.GetDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("latestForSource", m.GetLatestForSource())
        if err != nil {
            return err
        }
    }
    if m.GetSource() != nil {
        cast := (*m.GetSource()).String()
        err = writer.WriteStringValue("source", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalObjects sets the additionalObjects property value. The additionalObjects property
func (m *IdentityAccountAttributeValue) SetAdditionalObjects(value IdentityAccountAttributeValue_additionalObjectsable)() {
    m.additionalObjects = value
}
// SetAttribute sets the attribute property value. The attribute property
func (m *IdentityAccountAttributeValue) SetAttribute(value IdentityAccountAttributeDefinitionable)() {
    m.attribute = value
}
// SetContext sets the context property value. The context property
func (m *IdentityAccountAttributeValue) SetContext(value *string)() {
    m.context = value
}
// SetDate sets the date property value. The date property
func (m *IdentityAccountAttributeValue) SetDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.date = value
}
// SetLatestForSource sets the latestForSource property value. The latestForSource property
func (m *IdentityAccountAttributeValue) SetLatestForSource(value *bool)() {
    m.latestForSource = value
}
// SetSource sets the source property value. The source property
func (m *IdentityAccountAttributeValue) SetSource(value *IdentityAccountAttributeSource)() {
    m.source = value
}
// SetStatus sets the status property value. The status property
func (m *IdentityAccountAttributeValue) SetStatus(value *IdentityAccountAttributeValueStatus)() {
    m.status = value
}
// SetValue sets the value property value. The value property
func (m *IdentityAccountAttributeValue) SetValue(value *string)() {
    m.value = value
}
type IdentityAccountAttributeValueable interface {
    Linkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdditionalObjects()(IdentityAccountAttributeValue_additionalObjectsable)
    GetAttribute()(IdentityAccountAttributeDefinitionable)
    GetContext()(*string)
    GetDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLatestForSource()(*bool)
    GetSource()(*IdentityAccountAttributeSource)
    GetStatus()(*IdentityAccountAttributeValueStatus)
    GetValue()(*string)
    SetAdditionalObjects(value IdentityAccountAttributeValue_additionalObjectsable)()
    SetAttribute(value IdentityAccountAttributeDefinitionable)()
    SetContext(value *string)()
    SetDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLatestForSource(value *bool)()
    SetSource(value *IdentityAccountAttributeSource)()
    SetStatus(value *IdentityAccountAttributeValueStatus)()
    SetValue(value *string)()
}
