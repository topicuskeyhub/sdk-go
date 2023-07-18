package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// NotificationNotification 
type NotificationNotification struct {
    NonLinkable
}
// NewNotificationNotification instantiates a new notificationNotification and sets the default values.
func NewNotificationNotification()(*NotificationNotification) {
    m := &NotificationNotification{
        NonLinkable: *NewNonLinkable(),
    }
    return m
}
// CreateNotificationNotificationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateNotificationNotificationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("$type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "notification.CertificateExpiredNotification":
                        return NewNotificationCertificateExpiredNotification(), nil
                    case "notification.GroupAuditRequiredNotification":
                        return NewNotificationGroupAuditRequiredNotification(), nil
                    case "notification.GroupEditRequiredNotification":
                        return NewNotificationGroupEditRequiredNotification(), nil
                    case "notification.InvalidSignaturesDetectedNotification":
                        return NewNotificationInvalidSignaturesDetectedNotification(), nil
                    case "notification.LicenseStatusNotification":
                        return NewNotificationLicenseStatusNotification(), nil
                    case "notification.ModificationRequestNotification":
                        return NewNotificationModificationRequestNotification(), nil
                    case "notification.OldApiVersionUsageNotification":
                        return NewNotificationOldApiVersionUsageNotification(), nil
                    case "notification.ProvisionConfigRequiredNotification":
                        return NewNotificationProvisionConfigRequiredNotification(), nil
                    case "notification.RotatingPasswordRequiredNotification":
                        return NewNotificationRotatingPasswordRequiredNotification(), nil
                    case "notification.UpdateAvailableNotification":
                        return NewNotificationUpdateAvailableNotification(), nil
                    case "notification.VaultRecordExpiredNotification":
                        return NewNotificationVaultRecordExpiredNotification(), nil
                }
            }
        }
    }
    return NewNotificationNotification(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *NotificationNotification) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *NotificationNotification) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// NotificationNotificationable 
type NotificationNotificationable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
