package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ProfileAccessProfileAccount struct {
    AuthAccountPrimer
    // The additionalObjects property
    additionalObjects ProfileAccessProfileAccount_additionalObjectsable
    // The manual property
    manual *bool
}
// NewProfileAccessProfileAccount instantiates a new ProfileAccessProfileAccount and sets the default values.
func NewProfileAccessProfileAccount()(*ProfileAccessProfileAccount) {
    m := &ProfileAccessProfileAccount{
        AuthAccountPrimer: *NewAuthAccountPrimer(),
    }
    typeEscapedValue := "profile.AccessProfileAccount"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateProfileAccessProfileAccountFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateProfileAccessProfileAccountFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProfileAccessProfileAccount(), nil
}
// GetAdditionalObjects gets the additionalObjects property value. The additionalObjects property
// returns a ProfileAccessProfileAccount_additionalObjectsable when successful
func (m *ProfileAccessProfileAccount) GetAdditionalObjects()(ProfileAccessProfileAccount_additionalObjectsable) {
    return m.additionalObjects
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ProfileAccessProfileAccount) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuthAccountPrimer.GetFieldDeserializers()
    res["additionalObjects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProfileAccessProfileAccount_additionalObjectsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalObjects(val.(ProfileAccessProfileAccount_additionalObjectsable))
        }
        return nil
    }
    res["manual"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManual(val)
        }
        return nil
    }
    return res
}
// GetManual gets the manual property value. The manual property
// returns a *bool when successful
func (m *ProfileAccessProfileAccount) GetManual()(*bool) {
    return m.manual
}
// Serialize serializes information the current object
func (m *ProfileAccessProfileAccount) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AuthAccountPrimer.Serialize(writer)
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
        err = writer.WriteBoolValue("manual", m.GetManual())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalObjects sets the additionalObjects property value. The additionalObjects property
func (m *ProfileAccessProfileAccount) SetAdditionalObjects(value ProfileAccessProfileAccount_additionalObjectsable)() {
    m.additionalObjects = value
}
// SetManual sets the manual property value. The manual property
func (m *ProfileAccessProfileAccount) SetManual(value *bool)() {
    m.manual = value
}
type ProfileAccessProfileAccountable interface {
    AuthAccountPrimerable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdditionalObjects()(ProfileAccessProfileAccount_additionalObjectsable)
    GetManual()(*bool)
    SetAdditionalObjects(value ProfileAccessProfileAccount_additionalObjectsable)()
    SetManual(value *bool)()
}
