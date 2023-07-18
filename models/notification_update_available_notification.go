package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// NotificationUpdateAvailableNotification 
type NotificationUpdateAvailableNotification struct {
    NotificationNotification
    // The releasedAt property
    releasedAt *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
}
// NewNotificationUpdateAvailableNotification instantiates a new notificationUpdateAvailableNotification and sets the default values.
func NewNotificationUpdateAvailableNotification()(*NotificationUpdateAvailableNotification) {
    m := &NotificationUpdateAvailableNotification{
        NotificationNotification: *NewNotificationNotification(),
    }
    typeEscapedValue := "notification.UpdateAvailableNotification"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateNotificationUpdateAvailableNotificationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateNotificationUpdateAvailableNotificationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNotificationUpdateAvailableNotification(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *NotificationUpdateAvailableNotification) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NotificationNotification.GetFieldDeserializers()
    res["releasedAt"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReleasedAt(val)
        }
        return nil
    }
    return res
}
// GetReleasedAt gets the releasedAt property value. The releasedAt property
func (m *NotificationUpdateAvailableNotification) GetReleasedAt()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.releasedAt
}
// Serialize serializes information the current object
func (m *NotificationUpdateAvailableNotification) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NotificationNotification.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteDateOnlyValue("releasedAt", m.GetReleasedAt())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetReleasedAt sets the releasedAt property value. The releasedAt property
func (m *NotificationUpdateAvailableNotification) SetReleasedAt(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.releasedAt = value
}
// NotificationUpdateAvailableNotificationable 
type NotificationUpdateAvailableNotificationable interface {
    NotificationNotificationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetReleasedAt()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    SetReleasedAt(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
}
