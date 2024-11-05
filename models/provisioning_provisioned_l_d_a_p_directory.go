package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ProvisioningProvisionedLDAPDirectory struct {
    ProvisioningProvisionedSystem
    // The accountsWritable property
    accountsWritable *bool
    // The directory property
    directory DirectoryAccountDirectoryPrimerable
    // The gid property
    gid *int64
    // The groupDN property
    groupDN *string
    // The hashingScheme property
    hashingScheme *ProvisioningLDAPPasswordHashingScheme
    // The numbering property
    numbering ProvisioningProvisionNumberSequenceable
    // The samAccountNameScheme property
    samAccountNameScheme *ProvisioningADSamAccountNameScheme
    // The sshPublicKeySupport property
    sshPublicKeySupport *ProvisioningLDAPSshPublicKeySupport
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
// GetAccountsWritable gets the accountsWritable property value. The accountsWritable property
// returns a *bool when successful
func (m *ProvisioningProvisionedLDAPDirectory) GetAccountsWritable()(*bool) {
    return m.accountsWritable
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
    res["gid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGid(val)
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
    res["hashingScheme"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseProvisioningLDAPPasswordHashingScheme)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHashingScheme(val.(*ProvisioningLDAPPasswordHashingScheme))
        }
        return nil
    }
    res["numbering"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProvisioningProvisionNumberSequenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumbering(val.(ProvisioningProvisionNumberSequenceable))
        }
        return nil
    }
    res["samAccountNameScheme"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseProvisioningADSamAccountNameScheme)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSamAccountNameScheme(val.(*ProvisioningADSamAccountNameScheme))
        }
        return nil
    }
    res["sshPublicKeySupport"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseProvisioningLDAPSshPublicKeySupport)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSshPublicKeySupport(val.(*ProvisioningLDAPSshPublicKeySupport))
        }
        return nil
    }
    return res
}
// GetGid gets the gid property value. The gid property
// returns a *int64 when successful
func (m *ProvisioningProvisionedLDAPDirectory) GetGid()(*int64) {
    return m.gid
}
// GetGroupDN gets the groupDN property value. The groupDN property
// returns a *string when successful
func (m *ProvisioningProvisionedLDAPDirectory) GetGroupDN()(*string) {
    return m.groupDN
}
// GetHashingScheme gets the hashingScheme property value. The hashingScheme property
// returns a *ProvisioningLDAPPasswordHashingScheme when successful
func (m *ProvisioningProvisionedLDAPDirectory) GetHashingScheme()(*ProvisioningLDAPPasswordHashingScheme) {
    return m.hashingScheme
}
// GetNumbering gets the numbering property value. The numbering property
// returns a ProvisioningProvisionNumberSequenceable when successful
func (m *ProvisioningProvisionedLDAPDirectory) GetNumbering()(ProvisioningProvisionNumberSequenceable) {
    return m.numbering
}
// GetSamAccountNameScheme gets the samAccountNameScheme property value. The samAccountNameScheme property
// returns a *ProvisioningADSamAccountNameScheme when successful
func (m *ProvisioningProvisionedLDAPDirectory) GetSamAccountNameScheme()(*ProvisioningADSamAccountNameScheme) {
    return m.samAccountNameScheme
}
// GetSshPublicKeySupport gets the sshPublicKeySupport property value. The sshPublicKeySupport property
// returns a *ProvisioningLDAPSshPublicKeySupport when successful
func (m *ProvisioningProvisionedLDAPDirectory) GetSshPublicKeySupport()(*ProvisioningLDAPSshPublicKeySupport) {
    return m.sshPublicKeySupport
}
// Serialize serializes information the current object
func (m *ProvisioningProvisionedLDAPDirectory) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteInt64Value("gid", m.GetGid())
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
    if m.GetHashingScheme() != nil {
        cast := (*m.GetHashingScheme()).String()
        err = writer.WriteStringValue("hashingScheme", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("numbering", m.GetNumbering())
        if err != nil {
            return err
        }
    }
    if m.GetSamAccountNameScheme() != nil {
        cast := (*m.GetSamAccountNameScheme()).String()
        err = writer.WriteStringValue("samAccountNameScheme", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSshPublicKeySupport() != nil {
        cast := (*m.GetSshPublicKeySupport()).String()
        err = writer.WriteStringValue("sshPublicKeySupport", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccountsWritable sets the accountsWritable property value. The accountsWritable property
func (m *ProvisioningProvisionedLDAPDirectory) SetAccountsWritable(value *bool)() {
    m.accountsWritable = value
}
// SetDirectory sets the directory property value. The directory property
func (m *ProvisioningProvisionedLDAPDirectory) SetDirectory(value DirectoryAccountDirectoryPrimerable)() {
    m.directory = value
}
// SetGid sets the gid property value. The gid property
func (m *ProvisioningProvisionedLDAPDirectory) SetGid(value *int64)() {
    m.gid = value
}
// SetGroupDN sets the groupDN property value. The groupDN property
func (m *ProvisioningProvisionedLDAPDirectory) SetGroupDN(value *string)() {
    m.groupDN = value
}
// SetHashingScheme sets the hashingScheme property value. The hashingScheme property
func (m *ProvisioningProvisionedLDAPDirectory) SetHashingScheme(value *ProvisioningLDAPPasswordHashingScheme)() {
    m.hashingScheme = value
}
// SetNumbering sets the numbering property value. The numbering property
func (m *ProvisioningProvisionedLDAPDirectory) SetNumbering(value ProvisioningProvisionNumberSequenceable)() {
    m.numbering = value
}
// SetSamAccountNameScheme sets the samAccountNameScheme property value. The samAccountNameScheme property
func (m *ProvisioningProvisionedLDAPDirectory) SetSamAccountNameScheme(value *ProvisioningADSamAccountNameScheme)() {
    m.samAccountNameScheme = value
}
// SetSshPublicKeySupport sets the sshPublicKeySupport property value. The sshPublicKeySupport property
func (m *ProvisioningProvisionedLDAPDirectory) SetSshPublicKeySupport(value *ProvisioningLDAPSshPublicKeySupport)() {
    m.sshPublicKeySupport = value
}
type ProvisioningProvisionedLDAPDirectoryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ProvisioningProvisionedSystemable
    GetAccountsWritable()(*bool)
    GetDirectory()(DirectoryAccountDirectoryPrimerable)
    GetGid()(*int64)
    GetGroupDN()(*string)
    GetHashingScheme()(*ProvisioningLDAPPasswordHashingScheme)
    GetNumbering()(ProvisioningProvisionNumberSequenceable)
    GetSamAccountNameScheme()(*ProvisioningADSamAccountNameScheme)
    GetSshPublicKeySupport()(*ProvisioningLDAPSshPublicKeySupport)
    SetAccountsWritable(value *bool)()
    SetDirectory(value DirectoryAccountDirectoryPrimerable)()
    SetGid(value *int64)()
    SetGroupDN(value *string)()
    SetHashingScheme(value *ProvisioningLDAPPasswordHashingScheme)()
    SetNumbering(value ProvisioningProvisionNumberSequenceable)()
    SetSamAccountNameScheme(value *ProvisioningADSamAccountNameScheme)()
    SetSshPublicKeySupport(value *ProvisioningLDAPSshPublicKeySupport)()
}
