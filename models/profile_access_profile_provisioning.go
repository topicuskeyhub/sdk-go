package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ProfileAccessProfileProvisioning struct {
    Linkable
    // The accessProfile property
    accessProfile ProfileAccessProfilePrimerable
    // The additionalObjects property
    additionalObjects ProfileAccessProfileProvisioning_additionalObjectsable
    // The groupOnSystem property
    groupOnSystem ProvisioningGroupOnSystemable
}
// NewProfileAccessProfileProvisioning instantiates a new ProfileAccessProfileProvisioning and sets the default values.
func NewProfileAccessProfileProvisioning()(*ProfileAccessProfileProvisioning) {
    m := &ProfileAccessProfileProvisioning{
        Linkable: *NewLinkable(),
    }
    typeEscapedValue := "profile.AccessProfileProvisioning"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateProfileAccessProfileProvisioningFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateProfileAccessProfileProvisioningFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProfileAccessProfileProvisioning(), nil
}
// GetAccessProfile gets the accessProfile property value. The accessProfile property
// returns a ProfileAccessProfilePrimerable when successful
func (m *ProfileAccessProfileProvisioning) GetAccessProfile()(ProfileAccessProfilePrimerable) {
    return m.accessProfile
}
// GetAdditionalObjects gets the additionalObjects property value. The additionalObjects property
// returns a ProfileAccessProfileProvisioning_additionalObjectsable when successful
func (m *ProfileAccessProfileProvisioning) GetAdditionalObjects()(ProfileAccessProfileProvisioning_additionalObjectsable) {
    return m.additionalObjects
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ProfileAccessProfileProvisioning) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
        val, err := n.GetObjectValue(CreateProfileAccessProfileProvisioning_additionalObjectsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalObjects(val.(ProfileAccessProfileProvisioning_additionalObjectsable))
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
// GetGroupOnSystem gets the groupOnSystem property value. The groupOnSystem property
// returns a ProvisioningGroupOnSystemable when successful
func (m *ProfileAccessProfileProvisioning) GetGroupOnSystem()(ProvisioningGroupOnSystemable) {
    return m.groupOnSystem
}
// Serialize serializes information the current object
func (m *ProfileAccessProfileProvisioning) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteObjectValue("groupOnSystem", m.GetGroupOnSystem())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessProfile sets the accessProfile property value. The accessProfile property
func (m *ProfileAccessProfileProvisioning) SetAccessProfile(value ProfileAccessProfilePrimerable)() {
    m.accessProfile = value
}
// SetAdditionalObjects sets the additionalObjects property value. The additionalObjects property
func (m *ProfileAccessProfileProvisioning) SetAdditionalObjects(value ProfileAccessProfileProvisioning_additionalObjectsable)() {
    m.additionalObjects = value
}
// SetGroupOnSystem sets the groupOnSystem property value. The groupOnSystem property
func (m *ProfileAccessProfileProvisioning) SetGroupOnSystem(value ProvisioningGroupOnSystemable)() {
    m.groupOnSystem = value
}
type ProfileAccessProfileProvisioningable interface {
    Linkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessProfile()(ProfileAccessProfilePrimerable)
    GetAdditionalObjects()(ProfileAccessProfileProvisioning_additionalObjectsable)
    GetGroupOnSystem()(ProvisioningGroupOnSystemable)
    SetAccessProfile(value ProfileAccessProfilePrimerable)()
    SetAdditionalObjects(value ProfileAccessProfileProvisioning_additionalObjectsable)()
    SetGroupOnSystem(value ProvisioningGroupOnSystemable)()
}