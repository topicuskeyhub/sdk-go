package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ProfileAccessProfile struct {
    ProfileAccessProfilePrimer
    // The additionalObjects property
    additionalObjects ProfileAccessProfile_additionalObjectsable
    // The description property
    description *string
    // The owner property
    owner GroupGroupPrimerable
}
// NewProfileAccessProfile instantiates a new ProfileAccessProfile and sets the default values.
func NewProfileAccessProfile()(*ProfileAccessProfile) {
    m := &ProfileAccessProfile{
        ProfileAccessProfilePrimer: *NewProfileAccessProfilePrimer(),
    }
    typeEscapedValue := "profile.AccessProfile"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateProfileAccessProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateProfileAccessProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProfileAccessProfile(), nil
}
// GetAdditionalObjects gets the additionalObjects property value. The additionalObjects property
// returns a ProfileAccessProfile_additionalObjectsable when successful
func (m *ProfileAccessProfile) GetAdditionalObjects()(ProfileAccessProfile_additionalObjectsable) {
    return m.additionalObjects
}
// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *ProfileAccessProfile) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ProfileAccessProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ProfileAccessProfilePrimer.GetFieldDeserializers()
    res["additionalObjects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProfileAccessProfile_additionalObjectsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalObjects(val.(ProfileAccessProfile_additionalObjectsable))
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["owner"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupGroupPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwner(val.(GroupGroupPrimerable))
        }
        return nil
    }
    return res
}
// GetOwner gets the owner property value. The owner property
// returns a GroupGroupPrimerable when successful
func (m *ProfileAccessProfile) GetOwner()(GroupGroupPrimerable) {
    return m.owner
}
// Serialize serializes information the current object
func (m *ProfileAccessProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ProfileAccessProfilePrimer.Serialize(writer)
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
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("owner", m.GetOwner())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalObjects sets the additionalObjects property value. The additionalObjects property
func (m *ProfileAccessProfile) SetAdditionalObjects(value ProfileAccessProfile_additionalObjectsable)() {
    m.additionalObjects = value
}
// SetDescription sets the description property value. The description property
func (m *ProfileAccessProfile) SetDescription(value *string)() {
    m.description = value
}
// SetOwner sets the owner property value. The owner property
func (m *ProfileAccessProfile) SetOwner(value GroupGroupPrimerable)() {
    m.owner = value
}
type ProfileAccessProfileable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ProfileAccessProfilePrimerable
    GetAdditionalObjects()(ProfileAccessProfile_additionalObjectsable)
    GetDescription()(*string)
    GetOwner()(GroupGroupPrimerable)
    SetAdditionalObjects(value ProfileAccessProfile_additionalObjectsable)()
    SetDescription(value *string)()
    SetOwner(value GroupGroupPrimerable)()
}
