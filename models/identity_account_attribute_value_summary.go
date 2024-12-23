package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type IdentityAccountAttributeValueSummary struct {
    NonLinkable
    // The attribute property
    attribute IdentityAccountAttributeDefinitionable
    // The context property
    context *string
    // The currentValue property
    currentValue *string
    // The date property
    date *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The expectedValue property
    expectedValue *string
    // The source property
    source *IdentityAccountAttributeSource
    // The status property
    status *IdentityAccountAttributeValueSelectionStatus
}
// NewIdentityAccountAttributeValueSummary instantiates a new IdentityAccountAttributeValueSummary and sets the default values.
func NewIdentityAccountAttributeValueSummary()(*IdentityAccountAttributeValueSummary) {
    m := &IdentityAccountAttributeValueSummary{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "identity.AccountAttributeValueSummary"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateIdentityAccountAttributeValueSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIdentityAccountAttributeValueSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIdentityAccountAttributeValueSummary(), nil
}
// GetAttribute gets the attribute property value. The attribute property
// returns a IdentityAccountAttributeDefinitionable when successful
func (m *IdentityAccountAttributeValueSummary) GetAttribute()(IdentityAccountAttributeDefinitionable) {
    return m.attribute
}
// GetContext gets the context property value. The context property
// returns a *string when successful
func (m *IdentityAccountAttributeValueSummary) GetContext()(*string) {
    return m.context
}
// GetCurrentValue gets the currentValue property value. The currentValue property
// returns a *string when successful
func (m *IdentityAccountAttributeValueSummary) GetCurrentValue()(*string) {
    return m.currentValue
}
// GetDate gets the date property value. The date property
// returns a *Time when successful
func (m *IdentityAccountAttributeValueSummary) GetDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.date
}
// GetExpectedValue gets the expectedValue property value. The expectedValue property
// returns a *string when successful
func (m *IdentityAccountAttributeValueSummary) GetExpectedValue()(*string) {
    return m.expectedValue
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *IdentityAccountAttributeValueSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
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
    res["currentValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentValue(val)
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
    res["expectedValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpectedValue(val)
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
        val, err := n.GetEnumValue(ParseIdentityAccountAttributeValueSelectionStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*IdentityAccountAttributeValueSelectionStatus))
        }
        return nil
    }
    return res
}
// GetSource gets the source property value. The source property
// returns a *IdentityAccountAttributeSource when successful
func (m *IdentityAccountAttributeValueSummary) GetSource()(*IdentityAccountAttributeSource) {
    return m.source
}
// GetStatus gets the status property value. The status property
// returns a *IdentityAccountAttributeValueSelectionStatus when successful
func (m *IdentityAccountAttributeValueSummary) GetStatus()(*IdentityAccountAttributeValueSelectionStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *IdentityAccountAttributeValueSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
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
        err = writer.WriteStringValue("currentValue", m.GetCurrentValue())
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
        err = writer.WriteStringValue("expectedValue", m.GetExpectedValue())
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
    return nil
}
// SetAttribute sets the attribute property value. The attribute property
func (m *IdentityAccountAttributeValueSummary) SetAttribute(value IdentityAccountAttributeDefinitionable)() {
    m.attribute = value
}
// SetContext sets the context property value. The context property
func (m *IdentityAccountAttributeValueSummary) SetContext(value *string)() {
    m.context = value
}
// SetCurrentValue sets the currentValue property value. The currentValue property
func (m *IdentityAccountAttributeValueSummary) SetCurrentValue(value *string)() {
    m.currentValue = value
}
// SetDate sets the date property value. The date property
func (m *IdentityAccountAttributeValueSummary) SetDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.date = value
}
// SetExpectedValue sets the expectedValue property value. The expectedValue property
func (m *IdentityAccountAttributeValueSummary) SetExpectedValue(value *string)() {
    m.expectedValue = value
}
// SetSource sets the source property value. The source property
func (m *IdentityAccountAttributeValueSummary) SetSource(value *IdentityAccountAttributeSource)() {
    m.source = value
}
// SetStatus sets the status property value. The status property
func (m *IdentityAccountAttributeValueSummary) SetStatus(value *IdentityAccountAttributeValueSelectionStatus)() {
    m.status = value
}
type IdentityAccountAttributeValueSummaryable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAttribute()(IdentityAccountAttributeDefinitionable)
    GetContext()(*string)
    GetCurrentValue()(*string)
    GetDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetExpectedValue()(*string)
    GetSource()(*IdentityAccountAttributeSource)
    GetStatus()(*IdentityAccountAttributeValueSelectionStatus)
    SetAttribute(value IdentityAccountAttributeDefinitionable)()
    SetContext(value *string)()
    SetCurrentValue(value *string)()
    SetDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetExpectedValue(value *string)()
    SetSource(value *IdentityAccountAttributeSource)()
    SetStatus(value *IdentityAccountAttributeValueSelectionStatus)()
}
