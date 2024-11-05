package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ProfileAccessProfileAccountAttributeRuleStatus struct {
    NonLinkable
    // The attribute property
    attribute *ProfileAccessProfileComputedAttribute
    // The context property
    context *string
    // The currentValue property
    currentValue *string
    // The status property
    status *ProfileAccessProfileComputedAttributeStatus
}
// NewProfileAccessProfileAccountAttributeRuleStatus instantiates a new ProfileAccessProfileAccountAttributeRuleStatus and sets the default values.
func NewProfileAccessProfileAccountAttributeRuleStatus()(*ProfileAccessProfileAccountAttributeRuleStatus) {
    m := &ProfileAccessProfileAccountAttributeRuleStatus{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "profile.AccessProfileAccountAttributeRuleStatus"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateProfileAccessProfileAccountAttributeRuleStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateProfileAccessProfileAccountAttributeRuleStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProfileAccessProfileAccountAttributeRuleStatus(), nil
}
// GetAttribute gets the attribute property value. The attribute property
// returns a *ProfileAccessProfileComputedAttribute when successful
func (m *ProfileAccessProfileAccountAttributeRuleStatus) GetAttribute()(*ProfileAccessProfileComputedAttribute) {
    return m.attribute
}
// GetContext gets the context property value. The context property
// returns a *string when successful
func (m *ProfileAccessProfileAccountAttributeRuleStatus) GetContext()(*string) {
    return m.context
}
// GetCurrentValue gets the currentValue property value. The currentValue property
// returns a *string when successful
func (m *ProfileAccessProfileAccountAttributeRuleStatus) GetCurrentValue()(*string) {
    return m.currentValue
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ProfileAccessProfileAccountAttributeRuleStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
    res["attribute"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseProfileAccessProfileComputedAttribute)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttribute(val.(*ProfileAccessProfileComputedAttribute))
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
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseProfileAccessProfileComputedAttributeStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*ProfileAccessProfileComputedAttributeStatus))
        }
        return nil
    }
    return res
}
// GetStatus gets the status property value. The status property
// returns a *ProfileAccessProfileComputedAttributeStatus when successful
func (m *ProfileAccessProfileAccountAttributeRuleStatus) GetStatus()(*ProfileAccessProfileComputedAttributeStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *ProfileAccessProfileAccountAttributeRuleStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAttribute() != nil {
        cast := (*m.GetAttribute()).String()
        err = writer.WriteStringValue("attribute", &cast)
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
func (m *ProfileAccessProfileAccountAttributeRuleStatus) SetAttribute(value *ProfileAccessProfileComputedAttribute)() {
    m.attribute = value
}
// SetContext sets the context property value. The context property
func (m *ProfileAccessProfileAccountAttributeRuleStatus) SetContext(value *string)() {
    m.context = value
}
// SetCurrentValue sets the currentValue property value. The currentValue property
func (m *ProfileAccessProfileAccountAttributeRuleStatus) SetCurrentValue(value *string)() {
    m.currentValue = value
}
// SetStatus sets the status property value. The status property
func (m *ProfileAccessProfileAccountAttributeRuleStatus) SetStatus(value *ProfileAccessProfileComputedAttributeStatus)() {
    m.status = value
}
type ProfileAccessProfileAccountAttributeRuleStatusable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAttribute()(*ProfileAccessProfileComputedAttribute)
    GetContext()(*string)
    GetCurrentValue()(*string)
    GetStatus()(*ProfileAccessProfileComputedAttributeStatus)
    SetAttribute(value *ProfileAccessProfileComputedAttribute)()
    SetContext(value *string)()
    SetCurrentValue(value *string)()
    SetStatus(value *ProfileAccessProfileComputedAttributeStatus)()
}
