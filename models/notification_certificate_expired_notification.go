package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// NotificationCertificateExpiredNotification 
type NotificationCertificateExpiredNotification struct {
    NotificationNotification
    // The certificate property
    certificate CertificateCertificateable
    // The directory property
    directory DirectoryLDAPDirectoryable
    // The expiration property
    expiration *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The ldapClient property
    ldapClient ClientLdapClientable
    // The system property
    system ProvisioningProvisionedSystemable
    // The usage property
    usage *NotificationCertificateUsage
    // The webhook property
    webhook WebhookWebhookable
}
// NewNotificationCertificateExpiredNotification instantiates a new notificationCertificateExpiredNotification and sets the default values.
func NewNotificationCertificateExpiredNotification()(*NotificationCertificateExpiredNotification) {
    m := &NotificationCertificateExpiredNotification{
        NotificationNotification: *NewNotificationNotification(),
    }
    typeEscapedValue := "notification.CertificateExpiredNotification"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateNotificationCertificateExpiredNotificationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateNotificationCertificateExpiredNotificationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNotificationCertificateExpiredNotification(), nil
}
// GetCertificate gets the certificate property value. The certificate property
func (m *NotificationCertificateExpiredNotification) GetCertificate()(CertificateCertificateable) {
    return m.certificate
}
// GetDirectory gets the directory property value. The directory property
func (m *NotificationCertificateExpiredNotification) GetDirectory()(DirectoryLDAPDirectoryable) {
    return m.directory
}
// GetExpiration gets the expiration property value. The expiration property
func (m *NotificationCertificateExpiredNotification) GetExpiration()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.expiration
}
// GetFieldDeserializers the deserialization information for the current model
func (m *NotificationCertificateExpiredNotification) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NotificationNotification.GetFieldDeserializers()
    res["certificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCertificateCertificateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificate(val.(CertificateCertificateable))
        }
        return nil
    }
    res["directory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDirectoryLDAPDirectoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDirectory(val.(DirectoryLDAPDirectoryable))
        }
        return nil
    }
    res["expiration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpiration(val)
        }
        return nil
    }
    res["ldapClient"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateClientLdapClientFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLdapClient(val.(ClientLdapClientable))
        }
        return nil
    }
    res["system"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProvisioningProvisionedSystemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystem(val.(ProvisioningProvisionedSystemable))
        }
        return nil
    }
    res["usage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseNotificationCertificateUsage)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsage(val.(*NotificationCertificateUsage))
        }
        return nil
    }
    res["webhook"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWebhookWebhookFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebhook(val.(WebhookWebhookable))
        }
        return nil
    }
    return res
}
// GetLdapClient gets the ldapClient property value. The ldapClient property
func (m *NotificationCertificateExpiredNotification) GetLdapClient()(ClientLdapClientable) {
    return m.ldapClient
}
// GetSystem gets the system property value. The system property
func (m *NotificationCertificateExpiredNotification) GetSystem()(ProvisioningProvisionedSystemable) {
    return m.system
}
// GetUsage gets the usage property value. The usage property
func (m *NotificationCertificateExpiredNotification) GetUsage()(*NotificationCertificateUsage) {
    return m.usage
}
// GetWebhook gets the webhook property value. The webhook property
func (m *NotificationCertificateExpiredNotification) GetWebhook()(WebhookWebhookable) {
    return m.webhook
}
// Serialize serializes information the current object
func (m *NotificationCertificateExpiredNotification) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NotificationNotification.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("certificate", m.GetCertificate())
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
        err = writer.WriteTimeValue("expiration", m.GetExpiration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("ldapClient", m.GetLdapClient())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("system", m.GetSystem())
        if err != nil {
            return err
        }
    }
    if m.GetUsage() != nil {
        cast := (*m.GetUsage()).String()
        err = writer.WriteStringValue("usage", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("webhook", m.GetWebhook())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCertificate sets the certificate property value. The certificate property
func (m *NotificationCertificateExpiredNotification) SetCertificate(value CertificateCertificateable)() {
    m.certificate = value
}
// SetDirectory sets the directory property value. The directory property
func (m *NotificationCertificateExpiredNotification) SetDirectory(value DirectoryLDAPDirectoryable)() {
    m.directory = value
}
// SetExpiration sets the expiration property value. The expiration property
func (m *NotificationCertificateExpiredNotification) SetExpiration(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expiration = value
}
// SetLdapClient sets the ldapClient property value. The ldapClient property
func (m *NotificationCertificateExpiredNotification) SetLdapClient(value ClientLdapClientable)() {
    m.ldapClient = value
}
// SetSystem sets the system property value. The system property
func (m *NotificationCertificateExpiredNotification) SetSystem(value ProvisioningProvisionedSystemable)() {
    m.system = value
}
// SetUsage sets the usage property value. The usage property
func (m *NotificationCertificateExpiredNotification) SetUsage(value *NotificationCertificateUsage)() {
    m.usage = value
}
// SetWebhook sets the webhook property value. The webhook property
func (m *NotificationCertificateExpiredNotification) SetWebhook(value WebhookWebhookable)() {
    m.webhook = value
}
// NotificationCertificateExpiredNotificationable 
type NotificationCertificateExpiredNotificationable interface {
    NotificationNotificationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCertificate()(CertificateCertificateable)
    GetDirectory()(DirectoryLDAPDirectoryable)
    GetExpiration()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLdapClient()(ClientLdapClientable)
    GetSystem()(ProvisioningProvisionedSystemable)
    GetUsage()(*NotificationCertificateUsage)
    GetWebhook()(WebhookWebhookable)
    SetCertificate(value CertificateCertificateable)()
    SetDirectory(value DirectoryLDAPDirectoryable)()
    SetExpiration(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLdapClient(value ClientLdapClientable)()
    SetSystem(value ProvisioningProvisionedSystemable)()
    SetUsage(value *NotificationCertificateUsage)()
    SetWebhook(value WebhookWebhookable)()
}
