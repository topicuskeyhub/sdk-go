package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthAccount 
type AuthAccount struct {
    AuthAccountPrimer
    // The accountPermissions property
    accountPermissions []AuthPermissionable
    // The active property
    active *bool
    // The additionalObjects property
    additionalObjects AuthAccount_additionalObjectsable
    // The canRequestGroups property
    canRequestGroups *bool
    // The directory property
    directory DirectoryAccountDirectoryPrimerable
    // The directoryName property
    directoryName *string
    // The directoryPasswordChangeRequired property
    directoryPasswordChangeRequired *bool
    // The directoryRotatingPassword property
    directoryRotatingPassword *DirectoryDirectoryRotatingPassword
    // The directoryType property
    directoryType *DirectoryAccountDirectoryType
    // The email property
    email *string
    // The idInDirectory property
    idInDirectory *string
    // The keyHubPasswordChangeRequired property
    keyHubPasswordChangeRequired *bool
    // The lastModifiedAt property
    lastModifiedAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The licenseRole property
    licenseRole *AuthAccountLicenseRole
    // The locale property
    locale *string
    // The reregistrationRequired property
    reregistrationRequired *bool
    // The tokenPasswordEnabled property
    tokenPasswordEnabled *bool
    // The twoFactorStatus property
    twoFactorStatus *AuthTwoFactorAuthenticationStatus
    // The validInDirectory property
    validInDirectory *bool
}
// NewAuthAccount instantiates a new authAccount and sets the default values.
func NewAuthAccount()(*AuthAccount) {
    m := &AuthAccount{
        AuthAccountPrimer: *NewAuthAccountPrimer(),
    }
    typeEscapedValue := "auth.Account"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateAuthAccountFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthAccountFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("$type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "auth.InternalAccount":
                        return NewAuthInternalAccount(), nil
                }
            }
        }
    }
    return NewAuthAccount(), nil
}
// GetAccountPermissions gets the accountPermissions property value. The accountPermissions property
func (m *AuthAccount) GetAccountPermissions()([]AuthPermissionable) {
    return m.accountPermissions
}
// GetActive gets the active property value. The active property
func (m *AuthAccount) GetActive()(*bool) {
    return m.active
}
// GetAdditionalObjects gets the additionalObjects property value. The additionalObjects property
func (m *AuthAccount) GetAdditionalObjects()(AuthAccount_additionalObjectsable) {
    return m.additionalObjects
}
// GetCanRequestGroups gets the canRequestGroups property value. The canRequestGroups property
func (m *AuthAccount) GetCanRequestGroups()(*bool) {
    return m.canRequestGroups
}
// GetDirectory gets the directory property value. The directory property
func (m *AuthAccount) GetDirectory()(DirectoryAccountDirectoryPrimerable) {
    return m.directory
}
// GetDirectoryName gets the directoryName property value. The directoryName property
func (m *AuthAccount) GetDirectoryName()(*string) {
    return m.directoryName
}
// GetDirectoryPasswordChangeRequired gets the directoryPasswordChangeRequired property value. The directoryPasswordChangeRequired property
func (m *AuthAccount) GetDirectoryPasswordChangeRequired()(*bool) {
    return m.directoryPasswordChangeRequired
}
// GetDirectoryRotatingPassword gets the directoryRotatingPassword property value. The directoryRotatingPassword property
func (m *AuthAccount) GetDirectoryRotatingPassword()(*DirectoryDirectoryRotatingPassword) {
    return m.directoryRotatingPassword
}
// GetDirectoryType gets the directoryType property value. The directoryType property
func (m *AuthAccount) GetDirectoryType()(*DirectoryAccountDirectoryType) {
    return m.directoryType
}
// GetEmail gets the email property value. The email property
func (m *AuthAccount) GetEmail()(*string) {
    return m.email
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthAccount) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuthAccountPrimer.GetFieldDeserializers()
    res["accountPermissions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuthPermissionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuthPermissionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AuthPermissionable)
                }
            }
            m.SetAccountPermissions(res)
        }
        return nil
    }
    res["active"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActive(val)
        }
        return nil
    }
    res["additionalObjects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthAccount_additionalObjectsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalObjects(val.(AuthAccount_additionalObjectsable))
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
    res["directoryPasswordChangeRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDirectoryPasswordChangeRequired(val)
        }
        return nil
    }
    res["directoryRotatingPassword"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDirectoryDirectoryRotatingPassword)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDirectoryRotatingPassword(val.(*DirectoryDirectoryRotatingPassword))
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
    res["email"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmail(val)
        }
        return nil
    }
    res["idInDirectory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdInDirectory(val)
        }
        return nil
    }
    res["keyHubPasswordChangeRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyHubPasswordChangeRequired(val)
        }
        return nil
    }
    res["lastModifiedAt"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedAt(val)
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
    res["locale"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocale(val)
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
    res["tokenPasswordEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTokenPasswordEnabled(val)
        }
        return nil
    }
    res["twoFactorStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAuthTwoFactorAuthenticationStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTwoFactorStatus(val.(*AuthTwoFactorAuthenticationStatus))
        }
        return nil
    }
    res["validInDirectory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValidInDirectory(val)
        }
        return nil
    }
    return res
}
// GetIdInDirectory gets the idInDirectory property value. The idInDirectory property
func (m *AuthAccount) GetIdInDirectory()(*string) {
    return m.idInDirectory
}
// GetKeyHubPasswordChangeRequired gets the keyHubPasswordChangeRequired property value. The keyHubPasswordChangeRequired property
func (m *AuthAccount) GetKeyHubPasswordChangeRequired()(*bool) {
    return m.keyHubPasswordChangeRequired
}
// GetLastModifiedAt gets the lastModifiedAt property value. The lastModifiedAt property
func (m *AuthAccount) GetLastModifiedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedAt
}
// GetLicenseRole gets the licenseRole property value. The licenseRole property
func (m *AuthAccount) GetLicenseRole()(*AuthAccountLicenseRole) {
    return m.licenseRole
}
// GetLocale gets the locale property value. The locale property
func (m *AuthAccount) GetLocale()(*string) {
    return m.locale
}
// GetReregistrationRequired gets the reregistrationRequired property value. The reregistrationRequired property
func (m *AuthAccount) GetReregistrationRequired()(*bool) {
    return m.reregistrationRequired
}
// GetTokenPasswordEnabled gets the tokenPasswordEnabled property value. The tokenPasswordEnabled property
func (m *AuthAccount) GetTokenPasswordEnabled()(*bool) {
    return m.tokenPasswordEnabled
}
// GetTwoFactorStatus gets the twoFactorStatus property value. The twoFactorStatus property
func (m *AuthAccount) GetTwoFactorStatus()(*AuthTwoFactorAuthenticationStatus) {
    return m.twoFactorStatus
}
// GetValidInDirectory gets the validInDirectory property value. The validInDirectory property
func (m *AuthAccount) GetValidInDirectory()(*bool) {
    return m.validInDirectory
}
// Serialize serializes information the current object
func (m *AuthAccount) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AuthAccountPrimer.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("additionalObjects", m.GetAdditionalObjects())
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
    if m.GetDirectoryRotatingPassword() != nil {
        cast := (*m.GetDirectoryRotatingPassword()).String()
        err = writer.WriteStringValue("directoryRotatingPassword", &cast)
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
    if m.GetLicenseRole() != nil {
        cast := (*m.GetLicenseRole()).String()
        err = writer.WriteStringValue("licenseRole", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetTwoFactorStatus() != nil {
        cast := (*m.GetTwoFactorStatus()).String()
        err = writer.WriteStringValue("twoFactorStatus", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccountPermissions sets the accountPermissions property value. The accountPermissions property
func (m *AuthAccount) SetAccountPermissions(value []AuthPermissionable)() {
    m.accountPermissions = value
}
// SetActive sets the active property value. The active property
func (m *AuthAccount) SetActive(value *bool)() {
    m.active = value
}
// SetAdditionalObjects sets the additionalObjects property value. The additionalObjects property
func (m *AuthAccount) SetAdditionalObjects(value AuthAccount_additionalObjectsable)() {
    m.additionalObjects = value
}
// SetCanRequestGroups sets the canRequestGroups property value. The canRequestGroups property
func (m *AuthAccount) SetCanRequestGroups(value *bool)() {
    m.canRequestGroups = value
}
// SetDirectory sets the directory property value. The directory property
func (m *AuthAccount) SetDirectory(value DirectoryAccountDirectoryPrimerable)() {
    m.directory = value
}
// SetDirectoryName sets the directoryName property value. The directoryName property
func (m *AuthAccount) SetDirectoryName(value *string)() {
    m.directoryName = value
}
// SetDirectoryPasswordChangeRequired sets the directoryPasswordChangeRequired property value. The directoryPasswordChangeRequired property
func (m *AuthAccount) SetDirectoryPasswordChangeRequired(value *bool)() {
    m.directoryPasswordChangeRequired = value
}
// SetDirectoryRotatingPassword sets the directoryRotatingPassword property value. The directoryRotatingPassword property
func (m *AuthAccount) SetDirectoryRotatingPassword(value *DirectoryDirectoryRotatingPassword)() {
    m.directoryRotatingPassword = value
}
// SetDirectoryType sets the directoryType property value. The directoryType property
func (m *AuthAccount) SetDirectoryType(value *DirectoryAccountDirectoryType)() {
    m.directoryType = value
}
// SetEmail sets the email property value. The email property
func (m *AuthAccount) SetEmail(value *string)() {
    m.email = value
}
// SetIdInDirectory sets the idInDirectory property value. The idInDirectory property
func (m *AuthAccount) SetIdInDirectory(value *string)() {
    m.idInDirectory = value
}
// SetKeyHubPasswordChangeRequired sets the keyHubPasswordChangeRequired property value. The keyHubPasswordChangeRequired property
func (m *AuthAccount) SetKeyHubPasswordChangeRequired(value *bool)() {
    m.keyHubPasswordChangeRequired = value
}
// SetLastModifiedAt sets the lastModifiedAt property value. The lastModifiedAt property
func (m *AuthAccount) SetLastModifiedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedAt = value
}
// SetLicenseRole sets the licenseRole property value. The licenseRole property
func (m *AuthAccount) SetLicenseRole(value *AuthAccountLicenseRole)() {
    m.licenseRole = value
}
// SetLocale sets the locale property value. The locale property
func (m *AuthAccount) SetLocale(value *string)() {
    m.locale = value
}
// SetReregistrationRequired sets the reregistrationRequired property value. The reregistrationRequired property
func (m *AuthAccount) SetReregistrationRequired(value *bool)() {
    m.reregistrationRequired = value
}
// SetTokenPasswordEnabled sets the tokenPasswordEnabled property value. The tokenPasswordEnabled property
func (m *AuthAccount) SetTokenPasswordEnabled(value *bool)() {
    m.tokenPasswordEnabled = value
}
// SetTwoFactorStatus sets the twoFactorStatus property value. The twoFactorStatus property
func (m *AuthAccount) SetTwoFactorStatus(value *AuthTwoFactorAuthenticationStatus)() {
    m.twoFactorStatus = value
}
// SetValidInDirectory sets the validInDirectory property value. The validInDirectory property
func (m *AuthAccount) SetValidInDirectory(value *bool)() {
    m.validInDirectory = value
}
// AuthAccountable 
type AuthAccountable interface {
    AuthAccountPrimerable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccountPermissions()([]AuthPermissionable)
    GetActive()(*bool)
    GetAdditionalObjects()(AuthAccount_additionalObjectsable)
    GetCanRequestGroups()(*bool)
    GetDirectory()(DirectoryAccountDirectoryPrimerable)
    GetDirectoryName()(*string)
    GetDirectoryPasswordChangeRequired()(*bool)
    GetDirectoryRotatingPassword()(*DirectoryDirectoryRotatingPassword)
    GetDirectoryType()(*DirectoryAccountDirectoryType)
    GetEmail()(*string)
    GetIdInDirectory()(*string)
    GetKeyHubPasswordChangeRequired()(*bool)
    GetLastModifiedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLicenseRole()(*AuthAccountLicenseRole)
    GetLocale()(*string)
    GetReregistrationRequired()(*bool)
    GetTokenPasswordEnabled()(*bool)
    GetTwoFactorStatus()(*AuthTwoFactorAuthenticationStatus)
    GetValidInDirectory()(*bool)
    SetAccountPermissions(value []AuthPermissionable)()
    SetActive(value *bool)()
    SetAdditionalObjects(value AuthAccount_additionalObjectsable)()
    SetCanRequestGroups(value *bool)()
    SetDirectory(value DirectoryAccountDirectoryPrimerable)()
    SetDirectoryName(value *string)()
    SetDirectoryPasswordChangeRequired(value *bool)()
    SetDirectoryRotatingPassword(value *DirectoryDirectoryRotatingPassword)()
    SetDirectoryType(value *DirectoryAccountDirectoryType)()
    SetEmail(value *string)()
    SetIdInDirectory(value *string)()
    SetKeyHubPasswordChangeRequired(value *bool)()
    SetLastModifiedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLicenseRole(value *AuthAccountLicenseRole)()
    SetLocale(value *string)()
    SetReregistrationRequired(value *bool)()
    SetTokenPasswordEnabled(value *bool)()
    SetTwoFactorStatus(value *AuthTwoFactorAuthenticationStatus)()
    SetValidInDirectory(value *bool)()
}
