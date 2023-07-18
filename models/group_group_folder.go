package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupGroupFolder 
type GroupGroupFolder struct {
    Linkable
    // The additionalObjects property
    additionalObjects GroupGroupFolder_additionalObjectsable
    // The name property
    name *string
}
// NewGroupGroupFolder instantiates a new groupGroupFolder and sets the default values.
func NewGroupGroupFolder()(*GroupGroupFolder) {
    m := &GroupGroupFolder{
        Linkable: *NewLinkable(),
    }
    typeEscapedValue := "group.GroupFolder"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateGroupGroupFolderFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupGroupFolderFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupGroupFolder(), nil
}
// GetAdditionalObjects gets the additionalObjects property value. The additionalObjects property
func (m *GroupGroupFolder) GetAdditionalObjects()(GroupGroupFolder_additionalObjectsable) {
    return m.additionalObjects
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupGroupFolder) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Linkable.GetFieldDeserializers()
    res["additionalObjects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupGroupFolder_additionalObjectsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalObjects(val.(GroupGroupFolder_additionalObjectsable))
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
    return res
}
// GetName gets the name property value. The name property
func (m *GroupGroupFolder) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *GroupGroupFolder) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalObjects sets the additionalObjects property value. The additionalObjects property
func (m *GroupGroupFolder) SetAdditionalObjects(value GroupGroupFolder_additionalObjectsable)() {
    m.additionalObjects = value
}
// SetName sets the name property value. The name property
func (m *GroupGroupFolder) SetName(value *string)() {
    m.name = value
}
// GroupGroupFolderable 
type GroupGroupFolderable interface {
    Linkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdditionalObjects()(GroupGroupFolder_additionalObjectsable)
    GetName()(*string)
    SetAdditionalObjects(value GroupGroupFolder_additionalObjectsable)()
    SetName(value *string)()
}
