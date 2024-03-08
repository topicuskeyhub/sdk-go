package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ProvisioningProvisionedLDAPDirectory struct {
    ProvisioningProvisionedSystem
    // The directory property
    directory DirectoryAccountDirectoryPrimerable
    // The groupDN property
    groupDN *string
}
// NewProvisioningProvisionedLDAPDirectory instantiates a new ProvisioningProvisionedLDAPDirectory and sets the default values.
func NewProvisioningProvisionedLDAPDirectory()(*ProvisioningProvisionedLDAPDirectory) {
    m := &ProvisioningProvisionedLDAPDirectory{
        ProvisioningProvisionedSystem: *NewProvisioningProvisionedSystem(),
    }
    typeEscapedValue := "provisioning.ProvisionedLDAPDirectory"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateProvisioningProvisionedLDAPDirectoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateProvisioningProvisionedLDAPDirectoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProvisioningProvisionedLDAPDirectory(), nil
}
// GetDirectory gets the directory property value. The directory property
// returns a DirectoryAccountDirectoryPrimerable when successful
func (m *ProvisioningProvisionedLDAPDirectory) GetDirectory()(DirectoryAccountDirectoryPrimerable) {
    return m.directory
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ProvisioningProvisionedLDAPDirectory) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ProvisioningProvisionedSystem.GetFieldDeserializers()
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
    res["groupDN"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupDN(val)
        }
        return nil
    }
    return res
}
// GetGroupDN gets the groupDN property value. The groupDN property
// returns a *string when successful
func (m *ProvisioningProvisionedLDAPDirectory) GetGroupDN()(*string) {
    return m.groupDN
}
// Serialize serializes information the current object
func (m *ProvisioningProvisionedLDAPDirectory) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ProvisioningProvisionedSystem.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("directory", m.GetDirectory())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("groupDN", m.GetGroupDN())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDirectory sets the directory property value. The directory property
func (m *ProvisioningProvisionedLDAPDirectory) SetDirectory(value DirectoryAccountDirectoryPrimerable)() {
    m.directory = value
}
// SetGroupDN sets the groupDN property value. The groupDN property
func (m *ProvisioningProvisionedLDAPDirectory) SetGroupDN(value *string)() {
    m.groupDN = value
}
type ProvisioningProvisionedLDAPDirectoryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ProvisioningProvisionedSystemable
    GetDirectory()(DirectoryAccountDirectoryPrimerable)
    GetGroupDN()(*string)
    SetDirectory(value DirectoryAccountDirectoryPrimerable)()
    SetGroupDN(value *string)()
}
