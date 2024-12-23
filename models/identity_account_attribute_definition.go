package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type IdentityAccountAttributeDefinition struct {
    Linkable
    // The additionalObjects property
    additionalObjects IdentityAccountAttributeDefinition_additionalObjectsable
    // The format property
    format *IdentityAccountAttributeFormat
    // The name property
    name *string
    // The systemDefinition property
    systemDefinition *IdentityAccountAttributeSystemDefinition
}
// NewIdentityAccountAttributeDefinition instantiates a new IdentityAccountAttributeDefinition and sets the default values.
func NewIdentityAccountAttributeDefinition()(*IdentityAccountAttributeDefinition) {
    m := &IdentityAccountAttributeDefinition{
        Linkable: *NewLinkable(),
    }
    typeEscapedValue := "identity.AccountAttributeDefinition"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateIdentityAccountAttributeDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIdentityAccountAttributeDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIdentityAccountAttributeDefinition(), nil
}
// GetAdditionalObjects gets the additionalObjects property value. The additionalObjects property
// returns a IdentityAccountAttributeDefinition_additionalObjectsable when successful
func (m *IdentityAccountAttributeDefinition) GetAdditionalObjects()(IdentityAccountAttributeDefinition_additionalObjectsable) {
    return m.additionalObjects
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *IdentityAccountAttributeDefinition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Linkable.GetFieldDeserializers()
    res["additionalObjects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentityAccountAttributeDefinition_additionalObjectsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalObjects(val.(IdentityAccountAttributeDefinition_additionalObjectsable))
        }
        return nil
    }
    res["format"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseIdentityAccountAttributeFormat)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormat(val.(*IdentityAccountAttributeFormat))
        }
        return nil
    }
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
    res["systemDefinition"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseIdentityAccountAttributeSystemDefinition)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystemDefinition(val.(*IdentityAccountAttributeSystemDefinition))
        }
        return nil
    }
    return res
}
// GetFormat gets the format property value. The format property
// returns a *IdentityAccountAttributeFormat when successful
func (m *IdentityAccountAttributeDefinition) GetFormat()(*IdentityAccountAttributeFormat) {
    return m.format
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *IdentityAccountAttributeDefinition) GetName()(*string) {
    return m.name
}
// GetSystemDefinition gets the systemDefinition property value. The systemDefinition property
// returns a *IdentityAccountAttributeSystemDefinition when successful
func (m *IdentityAccountAttributeDefinition) GetSystemDefinition()(*IdentityAccountAttributeSystemDefinition) {
    return m.systemDefinition
}
// Serialize serializes information the current object
func (m *IdentityAccountAttributeDefinition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetFormat() != nil {
        cast := (*m.GetFormat()).String()
        err = writer.WriteStringValue("format", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    if m.GetSystemDefinition() != nil {
        cast := (*m.GetSystemDefinition()).String()
        err = writer.WriteStringValue("systemDefinition", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalObjects sets the additionalObjects property value. The additionalObjects property
func (m *IdentityAccountAttributeDefinition) SetAdditionalObjects(value IdentityAccountAttributeDefinition_additionalObjectsable)() {
    m.additionalObjects = value
}
// SetFormat sets the format property value. The format property
func (m *IdentityAccountAttributeDefinition) SetFormat(value *IdentityAccountAttributeFormat)() {
    m.format = value
}
// SetName sets the name property value. The name property
func (m *IdentityAccountAttributeDefinition) SetName(value *string)() {
    m.name = value
}
// SetSystemDefinition sets the systemDefinition property value. The systemDefinition property
func (m *IdentityAccountAttributeDefinition) SetSystemDefinition(value *IdentityAccountAttributeSystemDefinition)() {
    m.systemDefinition = value
}
type IdentityAccountAttributeDefinitionable interface {
    Linkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdditionalObjects()(IdentityAccountAttributeDefinition_additionalObjectsable)
    GetFormat()(*IdentityAccountAttributeFormat)
    GetName()(*string)
    GetSystemDefinition()(*IdentityAccountAttributeSystemDefinition)
    SetAdditionalObjects(value IdentityAccountAttributeDefinition_additionalObjectsable)()
    SetFormat(value *IdentityAccountAttributeFormat)()
    SetName(value *string)()
    SetSystemDefinition(value *IdentityAccountAttributeSystemDefinition)()
}
