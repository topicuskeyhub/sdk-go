package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ProfileAccessProfile struct {
    ProfileAccessProfilePrimer
    // The activateRuleScript property
    activateRuleScript *string
    // The additionalObjects property
    additionalObjects ProfileAccessProfile_additionalObjectsable
    // The description property
    description *string
    // The directory property
    directory DirectoryAccountDirectoryPrimerable
    // The matchRuleScript property
    matchRuleScript *string
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
// GetActivateRuleScript gets the activateRuleScript property value. The activateRuleScript property
// returns a *string when successful
func (m *ProfileAccessProfile) GetActivateRuleScript()(*string) {
    return m.activateRuleScript
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
// GetDirectory gets the directory property value. The directory property
// returns a DirectoryAccountDirectoryPrimerable when successful
func (m *ProfileAccessProfile) GetDirectory()(DirectoryAccountDirectoryPrimerable) {
    return m.directory
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ProfileAccessProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ProfileAccessProfilePrimer.GetFieldDeserializers()
    res["activateRuleScript"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivateRuleScript(val)
        }
        return nil
    }
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
    res["directory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDirectoryAccountDirectoryPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDirectory(val.(DirectoryAccountDirectoryPrimerable))
        }
        return nil
    }
    res["matchRuleScript"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMatchRuleScript(val)
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
// GetMatchRuleScript gets the matchRuleScript property value. The matchRuleScript property
// returns a *string when successful
func (m *ProfileAccessProfile) GetMatchRuleScript()(*string) {
    return m.matchRuleScript
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
        err = writer.WriteStringValue("activateRuleScript", m.GetActivateRuleScript())
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
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("directory", m.GetDirectory())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("matchRuleScript", m.GetMatchRuleScript())
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
// SetActivateRuleScript sets the activateRuleScript property value. The activateRuleScript property
func (m *ProfileAccessProfile) SetActivateRuleScript(value *string)() {
    m.activateRuleScript = value
}
// SetAdditionalObjects sets the additionalObjects property value. The additionalObjects property
func (m *ProfileAccessProfile) SetAdditionalObjects(value ProfileAccessProfile_additionalObjectsable)() {
    m.additionalObjects = value
}
// SetDescription sets the description property value. The description property
func (m *ProfileAccessProfile) SetDescription(value *string)() {
    m.description = value
}
// SetDirectory sets the directory property value. The directory property
func (m *ProfileAccessProfile) SetDirectory(value DirectoryAccountDirectoryPrimerable)() {
    m.directory = value
}
// SetMatchRuleScript sets the matchRuleScript property value. The matchRuleScript property
func (m *ProfileAccessProfile) SetMatchRuleScript(value *string)() {
    m.matchRuleScript = value
}
// SetOwner sets the owner property value. The owner property
func (m *ProfileAccessProfile) SetOwner(value GroupGroupPrimerable)() {
    m.owner = value
}
type ProfileAccessProfileable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ProfileAccessProfilePrimerable
    GetActivateRuleScript()(*string)
    GetAdditionalObjects()(ProfileAccessProfile_additionalObjectsable)
    GetDescription()(*string)
    GetDirectory()(DirectoryAccountDirectoryPrimerable)
    GetMatchRuleScript()(*string)
    GetOwner()(GroupGroupPrimerable)
    SetActivateRuleScript(value *string)()
    SetAdditionalObjects(value ProfileAccessProfile_additionalObjectsable)()
    SetDescription(value *string)()
    SetDirectory(value DirectoryAccountDirectoryPrimerable)()
    SetMatchRuleScript(value *string)()
    SetOwner(value GroupGroupPrimerable)()
}
