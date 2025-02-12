package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type NotificationProvisionConfigRequiredNotification struct {
    NotificationNotification
    // The keyhubPassword property
    keyhubPassword *bool
    // The systems property
    systems []ProvisioningProvisionedSystemPrimerable
}
// NewNotificationProvisionConfigRequiredNotification instantiates a new NotificationProvisionConfigRequiredNotification and sets the default values.
func NewNotificationProvisionConfigRequiredNotification()(*NotificationProvisionConfigRequiredNotification) {
    m := &NotificationProvisionConfigRequiredNotification{
        NotificationNotification: *NewNotificationNotification(),
    }
    typeEscapedValue := "notification.ProvisionConfigRequiredNotification"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateNotificationProvisionConfigRequiredNotificationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNotificationProvisionConfigRequiredNotificationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNotificationProvisionConfigRequiredNotification(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NotificationProvisionConfigRequiredNotification) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NotificationNotification.GetFieldDeserializers()
    res["keyhubPassword"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyhubPassword(val)
        }
        return nil
    }
    res["systems"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateProvisioningProvisionedSystemPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ProvisioningProvisionedSystemPrimerable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ProvisioningProvisionedSystemPrimerable)
                }
            }
            m.SetSystems(res)
        }
        return nil
    }
    return res
}
// GetKeyhubPassword gets the keyhubPassword property value. The keyhubPassword property
// returns a *bool when successful
func (m *NotificationProvisionConfigRequiredNotification) GetKeyhubPassword()(*bool) {
    return m.keyhubPassword
}
// GetSystems gets the systems property value. The systems property
// returns a []ProvisioningProvisionedSystemPrimerable when successful
func (m *NotificationProvisionConfigRequiredNotification) GetSystems()([]ProvisioningProvisionedSystemPrimerable) {
    return m.systems
}
// Serialize serializes information the current object
func (m *NotificationProvisionConfigRequiredNotification) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NotificationNotification.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("keyhubPassword", m.GetKeyhubPassword())
        if err != nil {
            return err
        }
    }
    if m.GetSystems() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSystems()))
        for i, v := range m.GetSystems() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("systems", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetKeyhubPassword sets the keyhubPassword property value. The keyhubPassword property
func (m *NotificationProvisionConfigRequiredNotification) SetKeyhubPassword(value *bool)() {
    m.keyhubPassword = value
}
// SetSystems sets the systems property value. The systems property
func (m *NotificationProvisionConfigRequiredNotification) SetSystems(value []ProvisioningProvisionedSystemPrimerable)() {
    m.systems = value
}
type NotificationProvisionConfigRequiredNotificationable interface {
    NotificationNotificationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetKeyhubPassword()(*bool)
    GetSystems()([]ProvisioningProvisionedSystemPrimerable)
    SetKeyhubPassword(value *bool)()
    SetSystems(value []ProvisioningProvisionedSystemPrimerable)()
}
