package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VaultVaultRecordSecrets 
type VaultVaultRecordSecrets struct {
    NonLinkable
    // The comment property
    comment *string
    // The file property
    file *string
    // The password property
    password *string
    // The totp property
    totp *string
    // The writeTotp property
    writeTotp *bool
}
// NewVaultVaultRecordSecrets instantiates a new vaultVaultRecordSecrets and sets the default values.
func NewVaultVaultRecordSecrets()(*VaultVaultRecordSecrets) {
    m := &VaultVaultRecordSecrets{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "vault.VaultRecordSecrets"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateVaultVaultRecordSecretsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVaultVaultRecordSecretsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVaultVaultRecordSecrets(), nil
}
// GetComment gets the comment property value. The comment property
func (m *VaultVaultRecordSecrets) GetComment()(*string) {
    return m.comment
}
// GetFieldDeserializers the deserialization information for the current model
func (m *VaultVaultRecordSecrets) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
    res["comment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComment(val)
        }
        return nil
    }
    res["file"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFile(val)
        }
        return nil
    }
    res["password"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPassword(val)
        }
        return nil
    }
    res["totp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotp(val)
        }
        return nil
    }
    res["writeTotp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWriteTotp(val)
        }
        return nil
    }
    return res
}
// GetFile gets the file property value. The file property
func (m *VaultVaultRecordSecrets) GetFile()(*string) {
    return m.file
}
// GetPassword gets the password property value. The password property
func (m *VaultVaultRecordSecrets) GetPassword()(*string) {
    return m.password
}
// GetTotp gets the totp property value. The totp property
func (m *VaultVaultRecordSecrets) GetTotp()(*string) {
    return m.totp
}
// GetWriteTotp gets the writeTotp property value. The writeTotp property
func (m *VaultVaultRecordSecrets) GetWriteTotp()(*bool) {
    return m.writeTotp
}
// Serialize serializes information the current object
func (m *VaultVaultRecordSecrets) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("comment", m.GetComment())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("file", m.GetFile())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("password", m.GetPassword())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("totp", m.GetTotp())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("writeTotp", m.GetWriteTotp())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetComment sets the comment property value. The comment property
func (m *VaultVaultRecordSecrets) SetComment(value *string)() {
    m.comment = value
}
// SetFile sets the file property value. The file property
func (m *VaultVaultRecordSecrets) SetFile(value *string)() {
    m.file = value
}
// SetPassword sets the password property value. The password property
func (m *VaultVaultRecordSecrets) SetPassword(value *string)() {
    m.password = value
}
// SetTotp sets the totp property value. The totp property
func (m *VaultVaultRecordSecrets) SetTotp(value *string)() {
    m.totp = value
}
// SetWriteTotp sets the writeTotp property value. The writeTotp property
func (m *VaultVaultRecordSecrets) SetWriteTotp(value *bool)() {
    m.writeTotp = value
}
// VaultVaultRecordSecretsable 
type VaultVaultRecordSecretsable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetComment()(*string)
    GetFile()(*string)
    GetPassword()(*string)
    GetTotp()(*string)
    GetWriteTotp()(*bool)
    SetComment(value *string)()
    SetFile(value *string)()
    SetPassword(value *string)()
    SetTotp(value *string)()
    SetWriteTotp(value *bool)()
}
