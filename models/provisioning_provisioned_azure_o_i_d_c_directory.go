package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ProvisioningProvisionedAzureOIDCDirectory struct {
    ProvisioningProvisionedSystem
    // The accountsWritable property
    accountsWritable *bool
    // The directory property
    directory DirectoryAccountDirectoryPrimerable
    // The tenant property
    tenant *string
}
// NewProvisioningProvisionedAzureOIDCDirectory instantiates a new ProvisioningProvisionedAzureOIDCDirectory and sets the default values.
func NewProvisioningProvisionedAzureOIDCDirectory()(*ProvisioningProvisionedAzureOIDCDirectory) {
    m := &ProvisioningProvisionedAzureOIDCDirectory{
        ProvisioningProvisionedSystem: *NewProvisioningProvisionedSystem(),
    }
    typeEscapedValue := "provisioning.ProvisionedAzureOIDCDirectory"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateProvisioningProvisionedAzureOIDCDirectoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateProvisioningProvisionedAzureOIDCDirectoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProvisioningProvisionedAzureOIDCDirectory(), nil
}
// GetAccountsWritable gets the accountsWritable property value. The accountsWritable property
// returns a *bool when successful
func (m *ProvisioningProvisionedAzureOIDCDirectory) GetAccountsWritable()(*bool) {
    return m.accountsWritable
}
// GetDirectory gets the directory property value. The directory property
// returns a DirectoryAccountDirectoryPrimerable when successful
func (m *ProvisioningProvisionedAzureOIDCDirectory) GetDirectory()(DirectoryAccountDirectoryPrimerable) {
    return m.directory
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ProvisioningProvisionedAzureOIDCDirectory) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ProvisioningProvisionedSystem.GetFieldDeserializers()
    res["accountsWritable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountsWritable(val)
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
// returns a *string when successful
func (m *ProvisioningProvisionedAzureOIDCDirectory) GetTenant()(*string) {
    return m.tenant
}
// Serialize serializes information the current object
func (m *ProvisioningProvisionedAzureOIDCDirectory) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ProvisioningProvisionedSystem.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("accountsWritable", m.GetAccountsWritable())
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
// SetAccountsWritable sets the accountsWritable property value. The accountsWritable property
func (m *ProvisioningProvisionedAzureOIDCDirectory) SetAccountsWritable(value *bool)() {
    m.accountsWritable = value
}
// SetDirectory sets the directory property value. The directory property
func (m *ProvisioningProvisionedAzureOIDCDirectory) SetDirectory(value DirectoryAccountDirectoryPrimerable)() {
    m.directory = value
}
// SetTenant sets the tenant property value. The tenant property
func (m *ProvisioningProvisionedAzureOIDCDirectory) SetTenant(value *string)() {
    m.tenant = value
}
type ProvisioningProvisionedAzureOIDCDirectoryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ProvisioningProvisionedSystemable
    GetAccountsWritable()(*bool)
    GetDirectory()(DirectoryAccountDirectoryPrimerable)
    GetTenant()(*string)
    SetAccountsWritable(value *bool)()
    SetDirectory(value DirectoryAccountDirectoryPrimerable)()
    SetTenant(value *string)()
}
