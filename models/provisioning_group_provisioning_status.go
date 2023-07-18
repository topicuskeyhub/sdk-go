package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProvisioningGroupProvisioningStatus 
type ProvisioningGroupProvisioningStatus struct {
    NonLinkable
    // The folder property
    folder GroupGroupFolderable
    // The group property
    group GroupGroupable
    // The provisioningDuration property
    provisioningDuration ProvisioningGroupProvisioningStatus_provisioningDurationable
    // The provisioningEndTime property
    provisioningEndTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The provisioningPermissionEndTime property
    provisioningPermissionEndTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The statusReport property
    statusReport ProvisioningAccountProvisioningStatusReportable
    // The visibleOnDashboard property
    visibleOnDashboard *bool
}
// NewProvisioningGroupProvisioningStatus instantiates a new provisioningGroupProvisioningStatus and sets the default values.
func NewProvisioningGroupProvisioningStatus()(*ProvisioningGroupProvisioningStatus) {
    m := &ProvisioningGroupProvisioningStatus{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "provisioning.GroupProvisioningStatus"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateProvisioningGroupProvisioningStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProvisioningGroupProvisioningStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProvisioningGroupProvisioningStatus(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProvisioningGroupProvisioningStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
    res["folder"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupGroupFolderFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFolder(val.(GroupGroupFolderable))
        }
        return nil
    }
    res["group"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroup(val.(GroupGroupable))
        }
        return nil
    }
    res["provisioningDuration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProvisioningGroupProvisioningStatus_provisioningDurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProvisioningDuration(val.(ProvisioningGroupProvisioningStatus_provisioningDurationable))
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
    res["provisioningPermissionEndTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProvisioningPermissionEndTime(val)
        }
        return nil
    }
    res["statusReport"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProvisioningAccountProvisioningStatusReportFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatusReport(val.(ProvisioningAccountProvisioningStatusReportable))
        }
        return nil
    }
    res["visibleOnDashboard"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVisibleOnDashboard(val)
        }
        return nil
    }
    return res
}
// GetFolder gets the folder property value. The folder property
func (m *ProvisioningGroupProvisioningStatus) GetFolder()(GroupGroupFolderable) {
    return m.folder
}
// GetGroup gets the group property value. The group property
func (m *ProvisioningGroupProvisioningStatus) GetGroup()(GroupGroupable) {
    return m.group
}
// GetProvisioningDuration gets the provisioningDuration property value. The provisioningDuration property
func (m *ProvisioningGroupProvisioningStatus) GetProvisioningDuration()(ProvisioningGroupProvisioningStatus_provisioningDurationable) {
    return m.provisioningDuration
}
// GetProvisioningEndTime gets the provisioningEndTime property value. The provisioningEndTime property
func (m *ProvisioningGroupProvisioningStatus) GetProvisioningEndTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.provisioningEndTime
}
// GetProvisioningPermissionEndTime gets the provisioningPermissionEndTime property value. The provisioningPermissionEndTime property
func (m *ProvisioningGroupProvisioningStatus) GetProvisioningPermissionEndTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.provisioningPermissionEndTime
}
// GetStatusReport gets the statusReport property value. The statusReport property
func (m *ProvisioningGroupProvisioningStatus) GetStatusReport()(ProvisioningAccountProvisioningStatusReportable) {
    return m.statusReport
}
// GetVisibleOnDashboard gets the visibleOnDashboard property value. The visibleOnDashboard property
func (m *ProvisioningGroupProvisioningStatus) GetVisibleOnDashboard()(*bool) {
    return m.visibleOnDashboard
}
// Serialize serializes information the current object
func (m *ProvisioningGroupProvisioningStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("folder", m.GetFolder())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("group", m.GetGroup())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("provisioningDuration", m.GetProvisioningDuration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("provisioningEndTime", m.GetProvisioningEndTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("provisioningPermissionEndTime", m.GetProvisioningPermissionEndTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("statusReport", m.GetStatusReport())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("visibleOnDashboard", m.GetVisibleOnDashboard())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFolder sets the folder property value. The folder property
func (m *ProvisioningGroupProvisioningStatus) SetFolder(value GroupGroupFolderable)() {
    m.folder = value
}
// SetGroup sets the group property value. The group property
func (m *ProvisioningGroupProvisioningStatus) SetGroup(value GroupGroupable)() {
    m.group = value
}
// SetProvisioningDuration sets the provisioningDuration property value. The provisioningDuration property
func (m *ProvisioningGroupProvisioningStatus) SetProvisioningDuration(value ProvisioningGroupProvisioningStatus_provisioningDurationable)() {
    m.provisioningDuration = value
}
// SetProvisioningEndTime sets the provisioningEndTime property value. The provisioningEndTime property
func (m *ProvisioningGroupProvisioningStatus) SetProvisioningEndTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.provisioningEndTime = value
}
// SetProvisioningPermissionEndTime sets the provisioningPermissionEndTime property value. The provisioningPermissionEndTime property
func (m *ProvisioningGroupProvisioningStatus) SetProvisioningPermissionEndTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.provisioningPermissionEndTime = value
}
// SetStatusReport sets the statusReport property value. The statusReport property
func (m *ProvisioningGroupProvisioningStatus) SetStatusReport(value ProvisioningAccountProvisioningStatusReportable)() {
    m.statusReport = value
}
// SetVisibleOnDashboard sets the visibleOnDashboard property value. The visibleOnDashboard property
func (m *ProvisioningGroupProvisioningStatus) SetVisibleOnDashboard(value *bool)() {
    m.visibleOnDashboard = value
}
// ProvisioningGroupProvisioningStatusable 
type ProvisioningGroupProvisioningStatusable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFolder()(GroupGroupFolderable)
    GetGroup()(GroupGroupable)
    GetProvisioningDuration()(ProvisioningGroupProvisioningStatus_provisioningDurationable)
    GetProvisioningEndTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetProvisioningPermissionEndTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetStatusReport()(ProvisioningAccountProvisioningStatusReportable)
    GetVisibleOnDashboard()(*bool)
    SetFolder(value GroupGroupFolderable)()
    SetGroup(value GroupGroupable)()
    SetProvisioningDuration(value ProvisioningGroupProvisioningStatus_provisioningDurationable)()
    SetProvisioningEndTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetProvisioningPermissionEndTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetStatusReport(value ProvisioningAccountProvisioningStatusReportable)()
    SetVisibleOnDashboard(value *bool)()
}
