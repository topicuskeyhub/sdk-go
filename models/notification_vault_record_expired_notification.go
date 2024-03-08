package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type NotificationVaultRecordExpiredNotification struct {
    NotificationNotification
    // The record property
    record VaultVaultRecordable
}
// NewNotificationVaultRecordExpiredNotification instantiates a new NotificationVaultRecordExpiredNotification and sets the default values.
func NewNotificationVaultRecordExpiredNotification()(*NotificationVaultRecordExpiredNotification) {
    m := &NotificationVaultRecordExpiredNotification{
        NotificationNotification: *NewNotificationNotification(),
    }
    typeEscapedValue := "notification.VaultRecordExpiredNotification"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateNotificationVaultRecordExpiredNotificationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNotificationVaultRecordExpiredNotificationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNotificationVaultRecordExpiredNotification(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NotificationVaultRecordExpiredNotification) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NotificationNotification.GetFieldDeserializers()
    res["record"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateVaultVaultRecordFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecord(val.(VaultVaultRecordable))
        }
        return nil
    }
    return res
}
// GetRecord gets the record property value. The record property
// returns a VaultVaultRecordable when successful
func (m *NotificationVaultRecordExpiredNotification) GetRecord()(VaultVaultRecordable) {
    return m.record
}
// Serialize serializes information the current object
func (m *NotificationVaultRecordExpiredNotification) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NotificationNotification.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("record", m.GetRecord())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetRecord sets the record property value. The record property
func (m *NotificationVaultRecordExpiredNotification) SetRecord(value VaultVaultRecordable)() {
    m.record = value
}
type NotificationVaultRecordExpiredNotificationable interface {
    NotificationNotificationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRecord()(VaultVaultRecordable)
    SetRecord(value VaultVaultRecordable)()
}
