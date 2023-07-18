package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProvisioningProvisionedAzureSyncLDAPDirectory 
type ProvisioningProvisionedAzureSyncLDAPDirectory struct {
    ProvisioningProvisionedSystem
    // The clientId property
    clientId *string
    // The clientSecret property
    clientSecret *string
    // The directory property
    directory DirectoryAccountDirectoryPrimerable
    // The tenant property
    tenant *string
}
// NewProvisioningProvisionedAzureSyncLDAPDirectory instantiates a new provisioningProvisionedAzureSyncLDAPDirectory and sets the default values.
func NewProvisioningProvisionedAzureSyncLDAPDirectory()(*ProvisioningProvisionedAzureSyncLDAPDirectory) {
    m := &ProvisioningProvisionedAzureSyncLDAPDirectory{
        ProvisioningProvisionedSystem: *NewProvisioningProvisionedSystem(),
    }
    typeEscapedValue := "provisioning.ProvisionedAzureSyncLDAPDirectory"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateProvisioningProvisionedAzureSyncLDAPDirectoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProvisioningProvisionedAzureSyncLDAPDirectoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProvisioningProvisionedAzureSyncLDAPDirectory(), nil
}
// GetClientId gets the clientId property value. The clientId property
func (m *ProvisioningProvisionedAzureSyncLDAPDirectory) GetClientId()(*string) {
    return m.clientId
}
// GetClientSecret gets the clientSecret property value. The clientSecret property
func (m *ProvisioningProvisionedAzureSyncLDAPDirectory) GetClientSecret()(*string) {
    return m.clientSecret
}
// GetDirectory gets the directory property value. The directory property
func (m *ProvisioningProvisionedAzureSyncLDAPDirectory) GetDirectory()(DirectoryAccountDirectoryPrimerable) {
    return m.directory
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProvisioningProvisionedAzureSyncLDAPDirectory) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ProvisioningProvisionedSystem.GetFieldDeserializers()
    res["clientId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientId(val)
        }
        return nil
    }
    res["clientSecret"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientSecret(val)
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
    res["tenant"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenant(val)
        }
        return nil
    }
    return res
}
// GetTenant gets the tenant property value. The tenant property
func (m *ProvisioningProvisionedAzureSyncLDAPDirectory) GetTenant()(*string) {
    return m.tenant
}
// Serialize serializes information the current object
func (m *ProvisioningProvisionedAzureSyncLDAPDirectory) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ProvisioningProvisionedSystem.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("clientId", m.GetClientId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("clientSecret", m.GetClientSecret())
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
        err = writer.WriteStringValue("tenant", m.GetTenant())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetClientId sets the clientId property value. The clientId property
func (m *ProvisioningProvisionedAzureSyncLDAPDirectory) SetClientId(value *string)() {
    m.clientId = value
}
// SetClientSecret sets the clientSecret property value. The clientSecret property
func (m *ProvisioningProvisionedAzureSyncLDAPDirectory) SetClientSecret(value *string)() {
    m.clientSecret = value
}
// SetDirectory sets the directory property value. The directory property
func (m *ProvisioningProvisionedAzureSyncLDAPDirectory) SetDirectory(value DirectoryAccountDirectoryPrimerable)() {
    m.directory = value
}
// SetTenant sets the tenant property value. The tenant property
func (m *ProvisioningProvisionedAzureSyncLDAPDirectory) SetTenant(value *string)() {
    m.tenant = value
}
// ProvisioningProvisionedAzureSyncLDAPDirectoryable 
type ProvisioningProvisionedAzureSyncLDAPDirectoryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ProvisioningProvisionedSystemable
    GetClientId()(*string)
    GetClientSecret()(*string)
    GetDirectory()(DirectoryAccountDirectoryPrimerable)
    GetTenant()(*string)
    SetClientId(value *string)()
    SetClientSecret(value *string)()
    SetDirectory(value DirectoryAccountDirectoryPrimerable)()
    SetTenant(value *string)()
}
