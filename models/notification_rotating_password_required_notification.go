package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type NotificationRotatingPasswordRequiredNotification struct {
    NotificationNotification
    // The groups property
    groups []GroupGroupPrimerable
}
// NewNotificationRotatingPasswordRequiredNotification instantiates a new NotificationRotatingPasswordRequiredNotification and sets the default values.
func NewNotificationRotatingPasswordRequiredNotification()(*NotificationRotatingPasswordRequiredNotification) {
    m := &NotificationRotatingPasswordRequiredNotification{
        NotificationNotification: *NewNotificationNotification(),
    }
    typeEscapedValue := "notification.RotatingPasswordRequiredNotification"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateNotificationRotatingPasswordRequiredNotificationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNotificationRotatingPasswordRequiredNotificationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNotificationRotatingPasswordRequiredNotification(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NotificationRotatingPasswordRequiredNotification) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NotificationNotification.GetFieldDeserializers()
    res["groups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGroupGroupPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GroupGroupPrimerable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(GroupGroupPrimerable)
                }
            }
            m.SetGroups(res)
        }
        return nil
    }
    return res
}
// GetGroups gets the groups property value. The groups property
// returns a []GroupGroupPrimerable when successful
func (m *NotificationRotatingPasswordRequiredNotification) GetGroups()([]GroupGroupPrimerable) {
    return m.groups
}
// Serialize serializes information the current object
func (m *NotificationRotatingPasswordRequiredNotification) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NotificationNotification.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetGroups() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetGroups()))
        for i, v := range m.GetGroups() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("groups", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetGroups sets the groups property value. The groups property
func (m *NotificationRotatingPasswordRequiredNotification) SetGroups(value []GroupGroupPrimerable)() {
    m.groups = value
}
type NotificationRotatingPasswordRequiredNotificationable interface {
    NotificationNotificationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetGroups()([]GroupGroupPrimerable)
    SetGroups(value []GroupGroupPrimerable)()
}
