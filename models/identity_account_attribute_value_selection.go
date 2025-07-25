// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type IdentityAccountAttributeValueSelection struct {
    NonLinkable
    // The context property
    context *string
    // The currentValue property
    currentValue *string
    // The status property
    status *IdentityAccountAttributeValueSelectionStatus
}
// NewIdentityAccountAttributeValueSelection instantiates a new IdentityAccountAttributeValueSelection and sets the default values.
func NewIdentityAccountAttributeValueSelection()(*IdentityAccountAttributeValueSelection) {
    m := &IdentityAccountAttributeValueSelection{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "identity.AccountAttributeValueSelection"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateIdentityAccountAttributeValueSelectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIdentityAccountAttributeValueSelectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIdentityAccountAttributeValueSelection(), nil
}
// GetContext gets the context property value. The context property
// returns a *string when successful
func (m *IdentityAccountAttributeValueSelection) GetContext()(*string) {
    return m.context
}
// GetCurrentValue gets the currentValue property value. The currentValue property
// returns a *string when successful
func (m *IdentityAccountAttributeValueSelection) GetCurrentValue()(*string) {
    return m.currentValue
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *IdentityAccountAttributeValueSelection) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
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
// GetStatus gets the status property value. The status property
// returns a *IdentityAccountAttributeValueSelectionStatus when successful
func (m *IdentityAccountAttributeValueSelection) GetStatus()(*IdentityAccountAttributeValueSelectionStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *IdentityAccountAttributeValueSelection) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
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
// SetContext sets the context property value. The context property
func (m *IdentityAccountAttributeValueSelection) SetContext(value *string)() {
    m.context = value
}
// SetCurrentValue sets the currentValue property value. The currentValue property
func (m *IdentityAccountAttributeValueSelection) SetCurrentValue(value *string)() {
    m.currentValue = value
}
// SetStatus sets the status property value. The status property
func (m *IdentityAccountAttributeValueSelection) SetStatus(value *IdentityAccountAttributeValueSelectionStatus)() {
    m.status = value
}
type IdentityAccountAttributeValueSelectionable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContext()(*string)
    GetCurrentValue()(*string)
    GetStatus()(*IdentityAccountAttributeValueSelectionStatus)
    SetContext(value *string)()
    SetCurrentValue(value *string)()
    SetStatus(value *IdentityAccountAttributeValueSelectionStatus)()
}
