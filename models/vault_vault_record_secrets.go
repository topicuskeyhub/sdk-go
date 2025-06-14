// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

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
    // The totpKey property
    totpKey *string
}
// NewVaultVaultRecordSecrets instantiates a new VaultVaultRecordSecrets and sets the default values.
func NewVaultVaultRecordSecrets()(*VaultVaultRecordSecrets) {
    m := &VaultVaultRecordSecrets{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "vault.VaultRecordSecrets"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateVaultVaultRecordSecretsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateVaultVaultRecordSecretsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVaultVaultRecordSecrets(), nil
}
// GetComment gets the comment property value. The comment property
// returns a *string when successful
func (m *VaultVaultRecordSecrets) GetComment()(*string) {
    return m.comment
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
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
    res["totpKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotpKey(val)
        }
        return nil
    }
    return res
}
// GetFile gets the file property value. The file property
// returns a *string when successful
func (m *VaultVaultRecordSecrets) GetFile()(*string) {
    return m.file
}
// GetPassword gets the password property value. The password property
// returns a *string when successful
func (m *VaultVaultRecordSecrets) GetPassword()(*string) {
    return m.password
}
// GetTotp gets the totp property value. The totp property
// returns a *string when successful
func (m *VaultVaultRecordSecrets) GetTotp()(*string) {
    return m.totp
}
// GetTotpKey gets the totpKey property value. The totpKey property
// returns a *string when successful
func (m *VaultVaultRecordSecrets) GetTotpKey()(*string) {
    return m.totpKey
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
        err = writer.WriteStringValue("totpKey", m.GetTotpKey())
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
// SetTotpKey sets the totpKey property value. The totpKey property
func (m *VaultVaultRecordSecrets) SetTotpKey(value *string)() {
    m.totpKey = value
}
type VaultVaultRecordSecretsable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetComment()(*string)
    GetFile()(*string)
    GetPassword()(*string)
    GetTotp()(*string)
    GetTotpKey()(*string)
    SetComment(value *string)()
    SetFile(value *string)()
    SetPassword(value *string)()
    SetTotp(value *string)()
    SetTotpKey(value *string)()
}
