package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupGroupPrimer 
type GroupGroupPrimer struct {
    Linkable
    // The admin property
    admin *bool
    // The name property
    name *string
    // The uuid property
    uuid *string
}
// NewGroupGroupPrimer instantiates a new groupGroupPrimer and sets the default values.
func NewGroupGroupPrimer()(*GroupGroupPrimer) {
    m := &GroupGroupPrimer{
        Linkable: *NewLinkable(),
    }
    typeEscapedValue := "group.GroupPrimer"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateGroupGroupPrimerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupGroupPrimerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("$type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "group.AccountGroup":
                        return NewGroupAccountGroup(), nil
                    case "group.Group":
                        return NewGroupGroup(), nil
                }
            }
        }
    }
    return NewGroupGroupPrimer(), nil
}
// GetAdmin gets the admin property value. The admin property
func (m *GroupGroupPrimer) GetAdmin()(*bool) {
    return m.admin
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupGroupPrimer) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Linkable.GetFieldDeserializers()
    res["admin"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdmin(val)
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
func (m *GroupGroupPrimer) GetName()(*string) {
    return m.name
}
// GetUuid gets the uuid property value. The uuid property
func (m *GroupGroupPrimer) GetUuid()(*string) {
    return m.uuid
}
// Serialize serializes information the current object
func (m *GroupGroupPrimer) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Linkable.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdmin sets the admin property value. The admin property
func (m *GroupGroupPrimer) SetAdmin(value *bool)() {
    m.admin = value
}
// SetName sets the name property value. The name property
func (m *GroupGroupPrimer) SetName(value *string)() {
    m.name = value
}
// SetUuid sets the uuid property value. The uuid property
func (m *GroupGroupPrimer) SetUuid(value *string)() {
    m.uuid = value
}
// GroupGroupPrimerable 
type GroupGroupPrimerable interface {
    Linkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdmin()(*bool)
    GetName()(*string)
    GetUuid()(*string)
    SetAdmin(value *bool)()
    SetName(value *string)()
    SetUuid(value *string)()
}
