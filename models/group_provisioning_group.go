package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupProvisioningGroup 
type GroupProvisioningGroup struct {
    Linkable
    // The activationRequired property
    activationRequired *bool
    // The additionalObjects property
    additionalObjects GroupProvisioningGroup_additionalObjectsable
    // The group property
    group GroupGroupPrimerable
    // The groupOnSystem property
    groupOnSystem ProvisioningGroupOnSystemable
}
// NewGroupProvisioningGroup instantiates a new groupProvisioningGroup and sets the default values.
func NewGroupProvisioningGroup()(*GroupProvisioningGroup) {
    m := &GroupProvisioningGroup{
        Linkable: *NewLinkable(),
    }
    typeEscapedValue := "group.ProvisioningGroup"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateGroupProvisioningGroupFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupProvisioningGroupFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupProvisioningGroup(), nil
}
// GetActivationRequired gets the activationRequired property value. The activationRequired property
func (m *GroupProvisioningGroup) GetActivationRequired()(*bool) {
    return m.activationRequired
}
// GetAdditionalObjects gets the additionalObjects property value. The additionalObjects property
func (m *GroupProvisioningGroup) GetAdditionalObjects()(GroupProvisioningGroup_additionalObjectsable) {
    return m.additionalObjects
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupProvisioningGroup) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Linkable.GetFieldDeserializers()
    res["activationRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivationRequired(val)
        }
        return nil
    }
    res["additionalObjects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupProvisioningGroup_additionalObjectsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalObjects(val.(GroupProvisioningGroup_additionalObjectsable))
        }
        return nil
    }
    res["group"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupGroupPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroup(val.(GroupGroupPrimerable))
        }
        return nil
    }
    res["groupOnSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProvisioningGroupOnSystemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupOnSystem(val.(ProvisioningGroupOnSystemable))
        }
        return nil
    }
    return res
}
// GetGroup gets the group property value. The group property
func (m *GroupProvisioningGroup) GetGroup()(GroupGroupPrimerable) {
    return m.group
}
// GetGroupOnSystem gets the groupOnSystem property value. The groupOnSystem property
func (m *GroupProvisioningGroup) GetGroupOnSystem()(ProvisioningGroupOnSystemable) {
    return m.groupOnSystem
}
// Serialize serializes information the current object
func (m *GroupProvisioningGroup) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Linkable.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("activationRequired", m.GetActivationRequired())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("additionalObjects", m.GetAdditionalObjects())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("group", m.GetGroup())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("groupOnSystem", m.GetGroupOnSystem())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActivationRequired sets the activationRequired property value. The activationRequired property
func (m *GroupProvisioningGroup) SetActivationRequired(value *bool)() {
    m.activationRequired = value
}
// SetAdditionalObjects sets the additionalObjects property value. The additionalObjects property
func (m *GroupProvisioningGroup) SetAdditionalObjects(value GroupProvisioningGroup_additionalObjectsable)() {
    m.additionalObjects = value
}
// SetGroup sets the group property value. The group property
func (m *GroupProvisioningGroup) SetGroup(value GroupGroupPrimerable)() {
    m.group = value
}
// SetGroupOnSystem sets the groupOnSystem property value. The groupOnSystem property
func (m *GroupProvisioningGroup) SetGroupOnSystem(value ProvisioningGroupOnSystemable)() {
    m.groupOnSystem = value
}
// GroupProvisioningGroupable 
type GroupProvisioningGroupable interface {
    Linkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActivationRequired()(*bool)
    GetAdditionalObjects()(GroupProvisioningGroup_additionalObjectsable)
    GetGroup()(GroupGroupPrimerable)
    GetGroupOnSystem()(ProvisioningGroupOnSystemable)
    SetActivationRequired(value *bool)()
    SetAdditionalObjects(value GroupProvisioningGroup_additionalObjectsable)()
    SetGroup(value GroupGroupPrimerable)()
    SetGroupOnSystem(value ProvisioningGroupOnSystemable)()
}
