package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AuthAccountSettings struct {
    NonLinkable
    // The defaultOrganizationalUnit property
    defaultOrganizationalUnit OrganizationOrganizationalUnitPrimerable
    // The directoryName property
    directoryName *string
    // The directoryType property
    directoryType *DirectoryAccountDirectoryType
    // The inGroups property
    inGroups *bool
    // The inMultipleOrganizationalUnits property
    inMultipleOrganizationalUnits *bool
    // The keyHubAdmin property
    keyHubAdmin *bool
    // The multipleOrganizationalUnitsExist property
    multipleOrganizationalUnitsExist *bool
    // The passwordMode property
    passwordMode *AuthPasswordMode
    // The sshPublicKey property
    sshPublicKey *string
    // The twoFactorAuthentication property
    twoFactorAuthentication *AuthTwoFactorAuthenticationStatus
    // The useTokenPassword property
    useTokenPassword *bool
    // The vaultStatus property
    vaultStatus *VaultAccountVaultStatus
}
// NewAuthAccountSettings instantiates a new AuthAccountSettings and sets the default values.
func NewAuthAccountSettings()(*AuthAccountSettings) {
    m := &AuthAccountSettings{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "auth.AccountSettings"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateAuthAccountSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAuthAccountSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthAccountSettings(), nil
}
// GetDefaultOrganizationalUnit gets the defaultOrganizationalUnit property value. The defaultOrganizationalUnit property
// returns a OrganizationOrganizationalUnitPrimerable when successful
func (m *AuthAccountSettings) GetDefaultOrganizationalUnit()(OrganizationOrganizationalUnitPrimerable) {
    return m.defaultOrganizationalUnit
}
// GetDirectoryName gets the directoryName property value. The directoryName property
// returns a *string when successful
func (m *AuthAccountSettings) GetDirectoryName()(*string) {
    return m.directoryName
}
// GetDirectoryType gets the directoryType property value. The directoryType property
// returns a *DirectoryAccountDirectoryType when successful
func (m *AuthAccountSettings) GetDirectoryType()(*DirectoryAccountDirectoryType) {
    return m.directoryType
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AuthAccountSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
    res["defaultOrganizationalUnit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOrganizationOrganizationalUnitPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultOrganizationalUnit(val.(OrganizationOrganizationalUnitPrimerable))
        }
        return nil
    }
    res["directoryName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDirectoryName(val)
        }
        return nil
    }
    res["directoryType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDirectoryAccountDirectoryType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDirectoryType(val.(*DirectoryAccountDirectoryType))
        }
        return nil
    }
    res["inGroups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInGroups(val)
        }
        return nil
    }
    res["inMultipleOrganizationalUnits"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInMultipleOrganizationalUnits(val)
        }
        return nil
    }
    res["keyHubAdmin"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyHubAdmin(val)
        }
        return nil
    }
    res["multipleOrganizationalUnitsExist"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMultipleOrganizationalUnitsExist(val)
        }
        return nil
    }
    res["passwordMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAuthPasswordMode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordMode(val.(*AuthPasswordMode))
        }
        return nil
    }
    res["sshPublicKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSshPublicKey(val)
        }
        return nil
    }
    res["twoFactorAuthentication"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAuthTwoFactorAuthenticationStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTwoFactorAuthentication(val.(*AuthTwoFactorAuthenticationStatus))
        }
        return nil
    }
    res["useTokenPassword"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUseTokenPassword(val)
        }
        return nil
    }
    res["vaultStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVaultAccountVaultStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVaultStatus(val.(*VaultAccountVaultStatus))
        }
        return nil
    }
    return res
}
// GetInGroups gets the inGroups property value. The inGroups property
// returns a *bool when successful
func (m *AuthAccountSettings) GetInGroups()(*bool) {
    return m.inGroups
}
// GetInMultipleOrganizationalUnits gets the inMultipleOrganizationalUnits property value. The inMultipleOrganizationalUnits property
// returns a *bool when successful
func (m *AuthAccountSettings) GetInMultipleOrganizationalUnits()(*bool) {
    return m.inMultipleOrganizationalUnits
}
// GetKeyHubAdmin gets the keyHubAdmin property value. The keyHubAdmin property
// returns a *bool when successful
func (m *AuthAccountSettings) GetKeyHubAdmin()(*bool) {
    return m.keyHubAdmin
}
// GetMultipleOrganizationalUnitsExist gets the multipleOrganizationalUnitsExist property value. The multipleOrganizationalUnitsExist property
// returns a *bool when successful
func (m *AuthAccountSettings) GetMultipleOrganizationalUnitsExist()(*bool) {
    return m.multipleOrganizationalUnitsExist
}
// GetPasswordMode gets the passwordMode property value. The passwordMode property
// returns a *AuthPasswordMode when successful
func (m *AuthAccountSettings) GetPasswordMode()(*AuthPasswordMode) {
    return m.passwordMode
}
// GetSshPublicKey gets the sshPublicKey property value. The sshPublicKey property
// returns a *string when successful
func (m *AuthAccountSettings) GetSshPublicKey()(*string) {
    return m.sshPublicKey
}
// GetTwoFactorAuthentication gets the twoFactorAuthentication property value. The twoFactorAuthentication property
// returns a *AuthTwoFactorAuthenticationStatus when successful
func (m *AuthAccountSettings) GetTwoFactorAuthentication()(*AuthTwoFactorAuthenticationStatus) {
    return m.twoFactorAuthentication
}
// GetUseTokenPassword gets the useTokenPassword property value. The useTokenPassword property
// returns a *bool when successful
func (m *AuthAccountSettings) GetUseTokenPassword()(*bool) {
    return m.useTokenPassword
}
// GetVaultStatus gets the vaultStatus property value. The vaultStatus property
// returns a *VaultAccountVaultStatus when successful
func (m *AuthAccountSettings) GetVaultStatus()(*VaultAccountVaultStatus) {
    return m.vaultStatus
}
// Serialize serializes information the current object
func (m *AuthAccountSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("defaultOrganizationalUnit", m.GetDefaultOrganizationalUnit())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("directoryName", m.GetDirectoryName())
        if err != nil {
            return err
        }
    }
    if m.GetDirectoryType() != nil {
        cast := (*m.GetDirectoryType()).String()
        err = writer.WriteStringValue("directoryType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("inGroups", m.GetInGroups())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("inMultipleOrganizationalUnits", m.GetInMultipleOrganizationalUnits())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("keyHubAdmin", m.GetKeyHubAdmin())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("multipleOrganizationalUnitsExist", m.GetMultipleOrganizationalUnitsExist())
        if err != nil {
            return err
        }
    }
    if m.GetPasswordMode() != nil {
        cast := (*m.GetPasswordMode()).String()
        err = writer.WriteStringValue("passwordMode", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("sshPublicKey", m.GetSshPublicKey())
        if err != nil {
            return err
        }
    }
    if m.GetTwoFactorAuthentication() != nil {
        cast := (*m.GetTwoFactorAuthentication()).String()
        err = writer.WriteStringValue("twoFactorAuthentication", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("useTokenPassword", m.GetUseTokenPassword())
        if err != nil {
            return err
        }
    }
    if m.GetVaultStatus() != nil {
        cast := (*m.GetVaultStatus()).String()
        err = writer.WriteStringValue("vaultStatus", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDefaultOrganizationalUnit sets the defaultOrganizationalUnit property value. The defaultOrganizationalUnit property
func (m *AuthAccountSettings) SetDefaultOrganizationalUnit(value OrganizationOrganizationalUnitPrimerable)() {
    m.defaultOrganizationalUnit = value
}
// SetDirectoryName sets the directoryName property value. The directoryName property
func (m *AuthAccountSettings) SetDirectoryName(value *string)() {
    m.directoryName = value
}
// SetDirectoryType sets the directoryType property value. The directoryType property
func (m *AuthAccountSettings) SetDirectoryType(value *DirectoryAccountDirectoryType)() {
    m.directoryType = value
}
// SetInGroups sets the inGroups property value. The inGroups property
func (m *AuthAccountSettings) SetInGroups(value *bool)() {
    m.inGroups = value
}
// SetInMultipleOrganizationalUnits sets the inMultipleOrganizationalUnits property value. The inMultipleOrganizationalUnits property
func (m *AuthAccountSettings) SetInMultipleOrganizationalUnits(value *bool)() {
    m.inMultipleOrganizationalUnits = value
}
// SetKeyHubAdmin sets the keyHubAdmin property value. The keyHubAdmin property
func (m *AuthAccountSettings) SetKeyHubAdmin(value *bool)() {
    m.keyHubAdmin = value
}
// SetMultipleOrganizationalUnitsExist sets the multipleOrganizationalUnitsExist property value. The multipleOrganizationalUnitsExist property
func (m *AuthAccountSettings) SetMultipleOrganizationalUnitsExist(value *bool)() {
    m.multipleOrganizationalUnitsExist = value
}
// SetPasswordMode sets the passwordMode property value. The passwordMode property
func (m *AuthAccountSettings) SetPasswordMode(value *AuthPasswordMode)() {
    m.passwordMode = value
}
// SetSshPublicKey sets the sshPublicKey property value. The sshPublicKey property
func (m *AuthAccountSettings) SetSshPublicKey(value *string)() {
    m.sshPublicKey = value
}
// SetTwoFactorAuthentication sets the twoFactorAuthentication property value. The twoFactorAuthentication property
func (m *AuthAccountSettings) SetTwoFactorAuthentication(value *AuthTwoFactorAuthenticationStatus)() {
    m.twoFactorAuthentication = value
}
// SetUseTokenPassword sets the useTokenPassword property value. The useTokenPassword property
func (m *AuthAccountSettings) SetUseTokenPassword(value *bool)() {
    m.useTokenPassword = value
}
// SetVaultStatus sets the vaultStatus property value. The vaultStatus property
func (m *AuthAccountSettings) SetVaultStatus(value *VaultAccountVaultStatus)() {
    m.vaultStatus = value
}
type AuthAccountSettingsable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDefaultOrganizationalUnit()(OrganizationOrganizationalUnitPrimerable)
    GetDirectoryName()(*string)
    GetDirectoryType()(*DirectoryAccountDirectoryType)
    GetInGroups()(*bool)
    GetInMultipleOrganizationalUnits()(*bool)
    GetKeyHubAdmin()(*bool)
    GetMultipleOrganizationalUnitsExist()(*bool)
    GetPasswordMode()(*AuthPasswordMode)
    GetSshPublicKey()(*string)
    GetTwoFactorAuthentication()(*AuthTwoFactorAuthenticationStatus)
    GetUseTokenPassword()(*bool)
    GetVaultStatus()(*VaultAccountVaultStatus)
    SetDefaultOrganizationalUnit(value OrganizationOrganizationalUnitPrimerable)()
    SetDirectoryName(value *string)()
    SetDirectoryType(value *DirectoryAccountDirectoryType)()
    SetInGroups(value *bool)()
    SetInMultipleOrganizationalUnits(value *bool)()
    SetKeyHubAdmin(value *bool)()
    SetMultipleOrganizationalUnitsExist(value *bool)()
    SetPasswordMode(value *AuthPasswordMode)()
    SetSshPublicKey(value *string)()
    SetTwoFactorAuthentication(value *AuthTwoFactorAuthenticationStatus)()
    SetUseTokenPassword(value *bool)()
    SetVaultStatus(value *VaultAccountVaultStatus)()
}
