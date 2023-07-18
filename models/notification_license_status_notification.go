package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// NotificationLicenseStatusNotification 
type NotificationLicenseStatusNotification struct {
    NotificationNotification
    // The activeAccounts property
    activeAccounts *int32
    // The licenseInfo property
    licenseInfo LicenseKeyHubLicenseInfoable
}
// NewNotificationLicenseStatusNotification instantiates a new notificationLicenseStatusNotification and sets the default values.
func NewNotificationLicenseStatusNotification()(*NotificationLicenseStatusNotification) {
    m := &NotificationLicenseStatusNotification{
        NotificationNotification: *NewNotificationNotification(),
    }
    typeEscapedValue := "notification.LicenseStatusNotification"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateNotificationLicenseStatusNotificationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateNotificationLicenseStatusNotificationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNotificationLicenseStatusNotification(), nil
}
// GetActiveAccounts gets the activeAccounts property value. The activeAccounts property
func (m *NotificationLicenseStatusNotification) GetActiveAccounts()(*int32) {
    return m.activeAccounts
}
// GetFieldDeserializers the deserialization information for the current model
func (m *NotificationLicenseStatusNotification) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NotificationNotification.GetFieldDeserializers()
    res["activeAccounts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActiveAccounts(val)
        }
        return nil
    }
    res["licenseInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLicenseKeyHubLicenseInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLicenseInfo(val.(LicenseKeyHubLicenseInfoable))
        }
        return nil
    }
    return res
}
// GetLicenseInfo gets the licenseInfo property value. The licenseInfo property
func (m *NotificationLicenseStatusNotification) GetLicenseInfo()(LicenseKeyHubLicenseInfoable) {
    return m.licenseInfo
}
// Serialize serializes information the current object
func (m *NotificationLicenseStatusNotification) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NotificationNotification.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("activeAccounts", m.GetActiveAccounts())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("licenseInfo", m.GetLicenseInfo())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActiveAccounts sets the activeAccounts property value. The activeAccounts property
func (m *NotificationLicenseStatusNotification) SetActiveAccounts(value *int32)() {
    m.activeAccounts = value
}
// SetLicenseInfo sets the licenseInfo property value. The licenseInfo property
func (m *NotificationLicenseStatusNotification) SetLicenseInfo(value LicenseKeyHubLicenseInfoable)() {
    m.licenseInfo = value
}
// NotificationLicenseStatusNotificationable 
type NotificationLicenseStatusNotificationable interface {
    NotificationNotificationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActiveAccounts()(*int32)
    GetLicenseInfo()(LicenseKeyHubLicenseInfoable)
    SetActiveAccounts(value *int32)()
    SetLicenseInfo(value LicenseKeyHubLicenseInfoable)()
}
