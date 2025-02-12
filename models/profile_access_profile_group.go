package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ProfileAccessProfileGroup struct {
    Linkable
    // The accessProfile property
    accessProfile ProfileAccessProfilePrimerable
    // The additionalObjects property
    additionalObjects ProfileAccessProfileGroup_additionalObjectsable
    // The group property
    group GroupGroupPrimerable
}
// NewProfileAccessProfileGroup instantiates a new ProfileAccessProfileGroup and sets the default values.
func NewProfileAccessProfileGroup()(*ProfileAccessProfileGroup) {
    m := &ProfileAccessProfileGroup{
        Linkable: *NewLinkable(),
    }
    typeEscapedValue := "profile.AccessProfileGroup"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateProfileAccessProfileGroupFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateProfileAccessProfileGroupFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProfileAccessProfileGroup(), nil
}
// GetAccessProfile gets the accessProfile property value. The accessProfile property
// returns a ProfileAccessProfilePrimerable when successful
func (m *ProfileAccessProfileGroup) GetAccessProfile()(ProfileAccessProfilePrimerable) {
    return m.accessProfile
}
// GetAdditionalObjects gets the additionalObjects property value. The additionalObjects property
// returns a ProfileAccessProfileGroup_additionalObjectsable when successful
func (m *ProfileAccessProfileGroup) GetAdditionalObjects()(ProfileAccessProfileGroup_additionalObjectsable) {
    return m.additionalObjects
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ProfileAccessProfileGroup) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Linkable.GetFieldDeserializers()
    res["accessProfile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProfileAccessProfilePrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessProfile(val.(ProfileAccessProfilePrimerable))
        }
        return nil
    }
    res["additionalObjects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProfileAccessProfileGroup_additionalObjectsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalObjects(val.(ProfileAccessProfileGroup_additionalObjectsable))
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
    return res
}
// GetGroup gets the group property value. The group property
// returns a GroupGroupPrimerable when successful
func (m *ProfileAccessProfileGroup) GetGroup()(GroupGroupPrimerable) {
    return m.group
}
// Serialize serializes information the current object
func (m *ProfileAccessProfileGroup) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Linkable.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("accessProfile", m.GetAccessProfile())
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
    return nil
}
// SetAccessProfile sets the accessProfile property value. The accessProfile property
func (m *ProfileAccessProfileGroup) SetAccessProfile(value ProfileAccessProfilePrimerable)() {
    m.accessProfile = value
}
// SetAdditionalObjects sets the additionalObjects property value. The additionalObjects property
func (m *ProfileAccessProfileGroup) SetAdditionalObjects(value ProfileAccessProfileGroup_additionalObjectsable)() {
    m.additionalObjects = value
}
// SetGroup sets the group property value. The group property
func (m *ProfileAccessProfileGroup) SetGroup(value GroupGroupPrimerable)() {
    m.group = value
}
type ProfileAccessProfileGroupable interface {
    Linkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessProfile()(ProfileAccessProfilePrimerable)
    GetAdditionalObjects()(ProfileAccessProfileGroup_additionalObjectsable)
    GetGroup()(GroupGroupPrimerable)
    SetAccessProfile(value ProfileAccessProfilePrimerable)()
    SetAdditionalObjects(value ProfileAccessProfileGroup_additionalObjectsable)()
    SetGroup(value GroupGroupPrimerable)()
}
