package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type NotificationInvalidSignaturesDetectedNotification struct {
    NotificationNotification
}
// NewNotificationInvalidSignaturesDetectedNotification instantiates a new NotificationInvalidSignaturesDetectedNotification and sets the default values.
func NewNotificationInvalidSignaturesDetectedNotification()(*NotificationInvalidSignaturesDetectedNotification) {
    m := &NotificationInvalidSignaturesDetectedNotification{
        NotificationNotification: *NewNotificationNotification(),
    }
    typeEscapedValue := "notification.InvalidSignaturesDetectedNotification"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateNotificationInvalidSignaturesDetectedNotificationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNotificationInvalidSignaturesDetectedNotificationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNotificationInvalidSignaturesDetectedNotification(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NotificationInvalidSignaturesDetectedNotification) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NotificationNotification.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *NotificationInvalidSignaturesDetectedNotification) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NotificationNotification.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type NotificationInvalidSignaturesDetectedNotificationable interface {
    NotificationNotificationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
