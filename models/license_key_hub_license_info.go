package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LicenseKeyHubLicenseInfo 
type LicenseKeyHubLicenseInfo struct {
    NonLinkable
    // The currentLicenseStatus property
    currentLicenseStatus *LicenseKeyHubLicenseInfoLicenseStatus
    // The customerCompany property
    customerCompany *string
    // The customerContact property
    customerContact *string
    // The customerDomains property
    customerDomains []string
    // The details property
    details []string
    // The expirationTime property
    expirationTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The features property
    features []LicenseLicenseFeature
    // The issueTime property
    issueTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The keyHubDomain property
    keyHubDomain *string
    // The licenseKey property
    licenseKey *string
    // The notBeforeTime property
    notBeforeTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The usersHardLimit property
    usersHardLimit *int32
    // The usersProLimit property
    usersProLimit *int32
    // The usersSoftLimit property
    usersSoftLimit *int32
    // The uuid property
    uuid *string
    // The version property
    version *int32
}
// NewLicenseKeyHubLicenseInfo instantiates a new licenseKeyHubLicenseInfo and sets the default values.
func NewLicenseKeyHubLicenseInfo()(*LicenseKeyHubLicenseInfo) {
    m := &LicenseKeyHubLicenseInfo{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "license.KeyHubLicenseInfo"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateLicenseKeyHubLicenseInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLicenseKeyHubLicenseInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLicenseKeyHubLicenseInfo(), nil
}
// GetCurrentLicenseStatus gets the currentLicenseStatus property value. The currentLicenseStatus property
func (m *LicenseKeyHubLicenseInfo) GetCurrentLicenseStatus()(*LicenseKeyHubLicenseInfoLicenseStatus) {
    return m.currentLicenseStatus
}
// GetCustomerCompany gets the customerCompany property value. The customerCompany property
func (m *LicenseKeyHubLicenseInfo) GetCustomerCompany()(*string) {
    return m.customerCompany
}
// GetCustomerContact gets the customerContact property value. The customerContact property
func (m *LicenseKeyHubLicenseInfo) GetCustomerContact()(*string) {
    return m.customerContact
}
// GetCustomerDomains gets the customerDomains property value. The customerDomains property
func (m *LicenseKeyHubLicenseInfo) GetCustomerDomains()([]string) {
    return m.customerDomains
}
// GetDetails gets the details property value. The details property
func (m *LicenseKeyHubLicenseInfo) GetDetails()([]string) {
    return m.details
}
// GetExpirationTime gets the expirationTime property value. The expirationTime property
func (m *LicenseKeyHubLicenseInfo) GetExpirationTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.expirationTime
}
// GetFeatures gets the features property value. The features property
func (m *LicenseKeyHubLicenseInfo) GetFeatures()([]LicenseLicenseFeature) {
    return m.features
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LicenseKeyHubLicenseInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
    res["currentLicenseStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseLicenseKeyHubLicenseInfoLicenseStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentLicenseStatus(val.(*LicenseKeyHubLicenseInfoLicenseStatus))
        }
        return nil
    }
    res["customerCompany"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomerCompany(val)
        }
        return nil
    }
    res["customerContact"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomerContact(val)
        }
        return nil
    }
    res["customerDomains"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetCustomerDomains(res)
        }
        return nil
    }
    res["details"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetDetails(res)
        }
        return nil
    }
    res["expirationTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationTime(val)
        }
        return nil
    }
    res["features"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseLicenseLicenseFeature)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LicenseLicenseFeature, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*LicenseLicenseFeature))
                }
            }
            m.SetFeatures(res)
        }
        return nil
    }
    res["issueTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssueTime(val)
        }
        return nil
    }
    res["keyHubDomain"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyHubDomain(val)
        }
        return nil
    }
    res["licenseKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLicenseKey(val)
        }
        return nil
    }
    res["notBeforeTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotBeforeTime(val)
        }
        return nil
    }
    res["usersHardLimit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsersHardLimit(val)
        }
        return nil
    }
    res["usersProLimit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsersProLimit(val)
        }
        return nil
    }
    res["usersSoftLimit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsersSoftLimit(val)
        }
        return nil
    }
    res["uuid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUuid(val)
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetIssueTime gets the issueTime property value. The issueTime property
func (m *LicenseKeyHubLicenseInfo) GetIssueTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.issueTime
}
// GetKeyHubDomain gets the keyHubDomain property value. The keyHubDomain property
func (m *LicenseKeyHubLicenseInfo) GetKeyHubDomain()(*string) {
    return m.keyHubDomain
}
// GetLicenseKey gets the licenseKey property value. The licenseKey property
func (m *LicenseKeyHubLicenseInfo) GetLicenseKey()(*string) {
    return m.licenseKey
}
// GetNotBeforeTime gets the notBeforeTime property value. The notBeforeTime property
func (m *LicenseKeyHubLicenseInfo) GetNotBeforeTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.notBeforeTime
}
// GetUsersHardLimit gets the usersHardLimit property value. The usersHardLimit property
func (m *LicenseKeyHubLicenseInfo) GetUsersHardLimit()(*int32) {
    return m.usersHardLimit
}
// GetUsersProLimit gets the usersProLimit property value. The usersProLimit property
func (m *LicenseKeyHubLicenseInfo) GetUsersProLimit()(*int32) {
    return m.usersProLimit
}
// GetUsersSoftLimit gets the usersSoftLimit property value. The usersSoftLimit property
func (m *LicenseKeyHubLicenseInfo) GetUsersSoftLimit()(*int32) {
    return m.usersSoftLimit
}
// GetUuid gets the uuid property value. The uuid property
func (m *LicenseKeyHubLicenseInfo) GetUuid()(*string) {
    return m.uuid
}
// GetVersion gets the version property value. The version property
func (m *LicenseKeyHubLicenseInfo) GetVersion()(*int32) {
    return m.version
}
// Serialize serializes information the current object
func (m *LicenseKeyHubLicenseInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCurrentLicenseStatus() != nil {
        cast := (*m.GetCurrentLicenseStatus()).String()
        err = writer.WriteStringValue("currentLicenseStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customerCompany", m.GetCustomerCompany())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customerContact", m.GetCustomerContact())
        if err != nil {
            return err
        }
    }
    if m.GetCustomerDomains() != nil {
        err = writer.WriteCollectionOfStringValues("customerDomains", m.GetCustomerDomains())
        if err != nil {
            return err
        }
    }
    if m.GetDetails() != nil {
        err = writer.WriteCollectionOfStringValues("details", m.GetDetails())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("expirationTime", m.GetExpirationTime())
        if err != nil {
            return err
        }
    }
    if m.GetFeatures() != nil {
        err = writer.WriteCollectionOfStringValues("features", SerializeLicenseLicenseFeature(m.GetFeatures()))
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("issueTime", m.GetIssueTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("keyHubDomain", m.GetKeyHubDomain())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("licenseKey", m.GetLicenseKey())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("notBeforeTime", m.GetNotBeforeTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("usersHardLimit", m.GetUsersHardLimit())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("usersProLimit", m.GetUsersProLimit())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("usersSoftLimit", m.GetUsersSoftLimit())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("uuid", m.GetUuid())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCurrentLicenseStatus sets the currentLicenseStatus property value. The currentLicenseStatus property
func (m *LicenseKeyHubLicenseInfo) SetCurrentLicenseStatus(value *LicenseKeyHubLicenseInfoLicenseStatus)() {
    m.currentLicenseStatus = value
}
// SetCustomerCompany sets the customerCompany property value. The customerCompany property
func (m *LicenseKeyHubLicenseInfo) SetCustomerCompany(value *string)() {
    m.customerCompany = value
}
// SetCustomerContact sets the customerContact property value. The customerContact property
func (m *LicenseKeyHubLicenseInfo) SetCustomerContact(value *string)() {
    m.customerContact = value
}
// SetCustomerDomains sets the customerDomains property value. The customerDomains property
func (m *LicenseKeyHubLicenseInfo) SetCustomerDomains(value []string)() {
    m.customerDomains = value
}
// SetDetails sets the details property value. The details property
func (m *LicenseKeyHubLicenseInfo) SetDetails(value []string)() {
    m.details = value
}
// SetExpirationTime sets the expirationTime property value. The expirationTime property
func (m *LicenseKeyHubLicenseInfo) SetExpirationTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expirationTime = value
}
// SetFeatures sets the features property value. The features property
func (m *LicenseKeyHubLicenseInfo) SetFeatures(value []LicenseLicenseFeature)() {
    m.features = value
}
// SetIssueTime sets the issueTime property value. The issueTime property
func (m *LicenseKeyHubLicenseInfo) SetIssueTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.issueTime = value
}
// SetKeyHubDomain sets the keyHubDomain property value. The keyHubDomain property
func (m *LicenseKeyHubLicenseInfo) SetKeyHubDomain(value *string)() {
    m.keyHubDomain = value
}
// SetLicenseKey sets the licenseKey property value. The licenseKey property
func (m *LicenseKeyHubLicenseInfo) SetLicenseKey(value *string)() {
    m.licenseKey = value
}
// SetNotBeforeTime sets the notBeforeTime property value. The notBeforeTime property
func (m *LicenseKeyHubLicenseInfo) SetNotBeforeTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.notBeforeTime = value
}
// SetUsersHardLimit sets the usersHardLimit property value. The usersHardLimit property
func (m *LicenseKeyHubLicenseInfo) SetUsersHardLimit(value *int32)() {
    m.usersHardLimit = value
}
// SetUsersProLimit sets the usersProLimit property value. The usersProLimit property
func (m *LicenseKeyHubLicenseInfo) SetUsersProLimit(value *int32)() {
    m.usersProLimit = value
}
// SetUsersSoftLimit sets the usersSoftLimit property value. The usersSoftLimit property
func (m *LicenseKeyHubLicenseInfo) SetUsersSoftLimit(value *int32)() {
    m.usersSoftLimit = value
}
// SetUuid sets the uuid property value. The uuid property
func (m *LicenseKeyHubLicenseInfo) SetUuid(value *string)() {
    m.uuid = value
}
// SetVersion sets the version property value. The version property
func (m *LicenseKeyHubLicenseInfo) SetVersion(value *int32)() {
    m.version = value
}
// LicenseKeyHubLicenseInfoable 
type LicenseKeyHubLicenseInfoable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCurrentLicenseStatus()(*LicenseKeyHubLicenseInfoLicenseStatus)
    GetCustomerCompany()(*string)
    GetCustomerContact()(*string)
    GetCustomerDomains()([]string)
    GetDetails()([]string)
    GetExpirationTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetFeatures()([]LicenseLicenseFeature)
    GetIssueTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetKeyHubDomain()(*string)
    GetLicenseKey()(*string)
    GetNotBeforeTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUsersHardLimit()(*int32)
    GetUsersProLimit()(*int32)
    GetUsersSoftLimit()(*int32)
    GetUuid()(*string)
    GetVersion()(*int32)
    SetCurrentLicenseStatus(value *LicenseKeyHubLicenseInfoLicenseStatus)()
    SetCustomerCompany(value *string)()
    SetCustomerContact(value *string)()
    SetCustomerDomains(value []string)()
    SetDetails(value []string)()
    SetExpirationTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetFeatures(value []LicenseLicenseFeature)()
    SetIssueTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetKeyHubDomain(value *string)()
    SetLicenseKey(value *string)()
    SetNotBeforeTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUsersHardLimit(value *int32)()
    SetUsersProLimit(value *int32)()
    SetUsersSoftLimit(value *int32)()
    SetUuid(value *string)()
    SetVersion(value *int32)()
}
