package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SimpleStringValue 
type SimpleStringValue struct {
    NonLinkable
    // The value property
    value *string
}
// NewSimpleStringValue instantiates a new simpleStringValue and sets the default values.
func NewSimpleStringValue()(*SimpleStringValue) {
    m := &SimpleStringValue{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "simple.StringValue"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateSimpleStringValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSimpleStringValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSimpleStringValue(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SimpleStringValue) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
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
// GetValue gets the value property value. The value property
func (m *SimpleStringValue) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *SimpleStringValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *SimpleStringValue) SetValue(value *string)() {
    m.value = value
}
// SimpleStringValueable 
type SimpleStringValueable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()(*string)
    SetValue(value *string)()
}
