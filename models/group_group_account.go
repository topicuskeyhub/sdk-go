package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupGroupAccount 
type GroupGroupAccount struct {
    AuthAccountPrimer
    // The additionalObjects property
    additionalObjects GroupGroupAccount_additionalObjectsable
    // The directory property
    directory DirectoryAccountDirectoryPrimerable
    // The disconnectedNested property
    disconnectedNested *bool
    // The endDate property
    endDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The lastUsed property
    lastUsed *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The nested property
    nested *bool
    // The provisioningEndTime property
    provisioningEndTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The rights property
    rights *GroupGroupRights
    // The twoFactorStatus property
    twoFactorStatus *AuthTwoFactorAuthenticationStatus
    // The visibleForProvisioning property
    visibleForProvisioning *bool
}
// NewGroupGroupAccount instantiates a new groupGroupAccount and sets the default values.
func NewGroupGroupAccount()(*GroupGroupAccount) {
    m := &GroupGroupAccount{
        AuthAccountPrimer: *NewAuthAccountPrimer(),
    }
    typeEscapedValue := "group.GroupAccount"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateGroupGroupAccountFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupGroupAccountFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupGroupAccount(), nil
}
// GetAdditionalObjects gets the additionalObjects property value. The additionalObjects property
func (m *GroupGroupAccount) GetAdditionalObjects()(GroupGroupAccount_additionalObjectsable) {
    return m.additionalObjects
}
// GetDirectory gets the directory property value. The directory property
func (m *GroupGroupAccount) GetDirectory()(DirectoryAccountDirectoryPrimerable) {
    return m.directory
}
// GetDisconnectedNested gets the disconnectedNested property value. The disconnectedNested property
func (m *GroupGroupAccount) GetDisconnectedNested()(*bool) {
    return m.disconnectedNested
}
// GetEndDate gets the endDate property value. The endDate property
func (m *GroupGroupAccount) GetEndDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.endDate
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupGroupAccount) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuthAccountPrimer.GetFieldDeserializers()
    res["additionalObjects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupGroupAccount_additionalObjectsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalObjects(val.(GroupGroupAccount_additionalObjectsable))
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
    res["disconnectedNested"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisconnectedNested(val)
        }
        return nil
    }
    res["endDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDate(val)
        }
        return nil
    }
    res["lastUsed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastUsed(val)
        }
        return nil
    }
    res["nested"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNested(val)
        }
        return nil
    }
    res["provisioningEndTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProvisioningEndTime(val)
        }
        return nil
    }
    res["rights"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseGroupGroupRights)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRights(val.(*GroupGroupRights))
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
    res["visibleForProvisioning"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVisibleForProvisioning(val)
        }
        return nil
    }
    return res
}
// GetLastUsed gets the lastUsed property value. The lastUsed property
func (m *GroupGroupAccount) GetLastUsed()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.lastUsed
}
// GetNested gets the nested property value. The nested property
func (m *GroupGroupAccount) GetNested()(*bool) {
    return m.nested
}
// GetProvisioningEndTime gets the provisioningEndTime property value. The provisioningEndTime property
func (m *GroupGroupAccount) GetProvisioningEndTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.provisioningEndTime
}
// GetRights gets the rights property value. The rights property
func (m *GroupGroupAccount) GetRights()(*GroupGroupRights) {
    return m.rights
}
// GetTwoFactorStatus gets the twoFactorStatus property value. The twoFactorStatus property
func (m *GroupGroupAccount) GetTwoFactorStatus()(*AuthTwoFactorAuthenticationStatus) {
    return m.twoFactorStatus
}
// GetVisibleForProvisioning gets the visibleForProvisioning property value. The visibleForProvisioning property
func (m *GroupGroupAccount) GetVisibleForProvisioning()(*bool) {
    return m.visibleForProvisioning
}
// Serialize serializes information the current object
func (m *GroupGroupAccount) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err = writer.WriteDateOnlyValue("endDate", m.GetEndDate())
        if err != nil {
            return err
        }
    }
    if m.GetRights() != nil {
        cast := (*m.GetRights()).String()
        err = writer.WriteStringValue("rights", &cast)
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
// SetAdditionalObjects sets the additionalObjects property value. The additionalObjects property
func (m *GroupGroupAccount) SetAdditionalObjects(value GroupGroupAccount_additionalObjectsable)() {
    m.additionalObjects = value
}
// SetDirectory sets the directory property value. The directory property
func (m *GroupGroupAccount) SetDirectory(value DirectoryAccountDirectoryPrimerable)() {
    m.directory = value
}
// SetDisconnectedNested sets the disconnectedNested property value. The disconnectedNested property
func (m *GroupGroupAccount) SetDisconnectedNested(value *bool)() {
    m.disconnectedNested = value
}
// SetEndDate sets the endDate property value. The endDate property
func (m *GroupGroupAccount) SetEndDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.endDate = value
}
// SetLastUsed sets the lastUsed property value. The lastUsed property
func (m *GroupGroupAccount) SetLastUsed(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.lastUsed = value
}
// SetNested sets the nested property value. The nested property
func (m *GroupGroupAccount) SetNested(value *bool)() {
    m.nested = value
}
// SetProvisioningEndTime sets the provisioningEndTime property value. The provisioningEndTime property
func (m *GroupGroupAccount) SetProvisioningEndTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.provisioningEndTime = value
}
// SetRights sets the rights property value. The rights property
func (m *GroupGroupAccount) SetRights(value *GroupGroupRights)() {
    m.rights = value
}
// SetTwoFactorStatus sets the twoFactorStatus property value. The twoFactorStatus property
func (m *GroupGroupAccount) SetTwoFactorStatus(value *AuthTwoFactorAuthenticationStatus)() {
    m.twoFactorStatus = value
}
// SetVisibleForProvisioning sets the visibleForProvisioning property value. The visibleForProvisioning property
func (m *GroupGroupAccount) SetVisibleForProvisioning(value *bool)() {
    m.visibleForProvisioning = value
}
// GroupGroupAccountable 
type GroupGroupAccountable interface {
    AuthAccountPrimerable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdditionalObjects()(GroupGroupAccount_additionalObjectsable)
    GetDirectory()(DirectoryAccountDirectoryPrimerable)
    GetDisconnectedNested()(*bool)
    GetEndDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetLastUsed()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetNested()(*bool)
    GetProvisioningEndTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRights()(*GroupGroupRights)
    GetTwoFactorStatus()(*AuthTwoFactorAuthenticationStatus)
    GetVisibleForProvisioning()(*bool)
    SetAdditionalObjects(value GroupGroupAccount_additionalObjectsable)()
    SetDirectory(value DirectoryAccountDirectoryPrimerable)()
    SetDisconnectedNested(value *bool)()
    SetEndDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetLastUsed(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetNested(value *bool)()
    SetProvisioningEndTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRights(value *GroupGroupRights)()
    SetTwoFactorStatus(value *AuthTwoFactorAuthenticationStatus)()
    SetVisibleForProvisioning(value *bool)()
}
