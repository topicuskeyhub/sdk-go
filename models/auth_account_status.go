package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthAccountStatus 
type AuthAccountStatus struct {
    NonLinkable
    // The accountEnabled property
    accountEnabled *bool
    // The canRequestGroups property
    canRequestGroups *bool
    // The declineRecoveryRequests property
    declineRecoveryRequests *bool
    // The licenseRole property
    licenseRole *AuthAccountLicenseRole
    // The reregistrationRequired property
    reregistrationRequired *bool
    // The twoFactorAuthenticationEnabled property
    twoFactorAuthenticationEnabled *bool
}
// NewAuthAccountStatus instantiates a new authAccountStatus and sets the default values.
func NewAuthAccountStatus()(*AuthAccountStatus) {
    m := &AuthAccountStatus{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "auth.AccountStatus"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateAuthAccountStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthAccountStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthAccountStatus(), nil
}
// GetAccountEnabled gets the accountEnabled property value. The accountEnabled property
func (m *AuthAccountStatus) GetAccountEnabled()(*bool) {
    return m.accountEnabled
}
// GetCanRequestGroups gets the canRequestGroups property value. The canRequestGroups property
func (m *AuthAccountStatus) GetCanRequestGroups()(*bool) {
    return m.canRequestGroups
}
// GetDeclineRecoveryRequests gets the declineRecoveryRequests property value. The declineRecoveryRequests property
func (m *AuthAccountStatus) GetDeclineRecoveryRequests()(*bool) {
    return m.declineRecoveryRequests
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthAccountStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
    res["accountEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountEnabled(val)
        }
        return nil
    }
    res["canRequestGroups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCanRequestGroups(val)
        }
        return nil
    }
    res["declineRecoveryRequests"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeclineRecoveryRequests(val)
        }
        return nil
    }
    res["licenseRole"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAuthAccountLicenseRole)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLicenseRole(val.(*AuthAccountLicenseRole))
        }
        return nil
    }
    res["reregistrationRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReregistrationRequired(val)
        }
        return nil
    }
    res["twoFactorAuthenticationEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTwoFactorAuthenticationEnabled(val)
        }
        return nil
    }
    return res
}
// GetLicenseRole gets the licenseRole property value. The licenseRole property
func (m *AuthAccountStatus) GetLicenseRole()(*AuthAccountLicenseRole) {
    return m.licenseRole
}
// GetReregistrationRequired gets the reregistrationRequired property value. The reregistrationRequired property
func (m *AuthAccountStatus) GetReregistrationRequired()(*bool) {
    return m.reregistrationRequired
}
// GetTwoFactorAuthenticationEnabled gets the twoFactorAuthenticationEnabled property value. The twoFactorAuthenticationEnabled property
func (m *AuthAccountStatus) GetTwoFactorAuthenticationEnabled()(*bool) {
    return m.twoFactorAuthenticationEnabled
}
// Serialize serializes information the current object
func (m *AuthAccountStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("accountEnabled", m.GetAccountEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("canRequestGroups", m.GetCanRequestGroups())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("declineRecoveryRequests", m.GetDeclineRecoveryRequests())
        if err != nil {
            return err
        }
    }
    if m.GetLicenseRole() != nil {
        cast := (*m.GetLicenseRole()).String()
        err = writer.WriteStringValue("licenseRole", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("reregistrationRequired", m.GetReregistrationRequired())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("twoFactorAuthenticationEnabled", m.GetTwoFactorAuthenticationEnabled())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccountEnabled sets the accountEnabled property value. The accountEnabled property
func (m *AuthAccountStatus) SetAccountEnabled(value *bool)() {
    m.accountEnabled = value
}
// SetCanRequestGroups sets the canRequestGroups property value. The canRequestGroups property
func (m *AuthAccountStatus) SetCanRequestGroups(value *bool)() {
    m.canRequestGroups = value
}
// SetDeclineRecoveryRequests sets the declineRecoveryRequests property value. The declineRecoveryRequests property
func (m *AuthAccountStatus) SetDeclineRecoveryRequests(value *bool)() {
    m.declineRecoveryRequests = value
}
// SetLicenseRole sets the licenseRole property value. The licenseRole property
func (m *AuthAccountStatus) SetLicenseRole(value *AuthAccountLicenseRole)() {
    m.licenseRole = value
}
// SetReregistrationRequired sets the reregistrationRequired property value. The reregistrationRequired property
func (m *AuthAccountStatus) SetReregistrationRequired(value *bool)() {
    m.reregistrationRequired = value
}
// SetTwoFactorAuthenticationEnabled sets the twoFactorAuthenticationEnabled property value. The twoFactorAuthenticationEnabled property
func (m *AuthAccountStatus) SetTwoFactorAuthenticationEnabled(value *bool)() {
    m.twoFactorAuthenticationEnabled = value
}
// AuthAccountStatusable 
type AuthAccountStatusable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccountEnabled()(*bool)
    GetCanRequestGroups()(*bool)
    GetDeclineRecoveryRequests()(*bool)
    GetLicenseRole()(*AuthAccountLicenseRole)
    GetReregistrationRequired()(*bool)
    GetTwoFactorAuthenticationEnabled()(*bool)
    SetAccountEnabled(value *bool)()
    SetCanRequestGroups(value *bool)()
    SetDeclineRecoveryRequests(value *bool)()
    SetLicenseRole(value *AuthAccountLicenseRole)()
    SetReregistrationRequired(value *bool)()
    SetTwoFactorAuthenticationEnabled(value *bool)()
}
