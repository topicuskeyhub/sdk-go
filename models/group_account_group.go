package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupAccountGroup 
type GroupAccountGroup struct {
    GroupGroupPrimer
    // The additionalObjects property
    additionalObjects GroupAccountGroup_additionalObjectsable
    // The endDate property
    endDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The folder property
    folder GroupGroupFolderable
    // The lastUsed property
    lastUsed *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The provisioningEndTime property
    provisioningEndTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The rights property
    rights *GroupGroupRights
    // The visibleForProvisioning property
    visibleForProvisioning *bool
}
// NewGroupAccountGroup instantiates a new groupAccountGroup and sets the default values.
func NewGroupAccountGroup()(*GroupAccountGroup) {
    m := &GroupAccountGroup{
        GroupGroupPrimer: *NewGroupGroupPrimer(),
    }
    typeEscapedValue := "group.AccountGroup"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateGroupAccountGroupFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupAccountGroupFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupAccountGroup(), nil
}
// GetAdditionalObjects gets the additionalObjects property value. The additionalObjects property
func (m *GroupAccountGroup) GetAdditionalObjects()(GroupAccountGroup_additionalObjectsable) {
    return m.additionalObjects
}
// GetEndDate gets the endDate property value. The endDate property
func (m *GroupAccountGroup) GetEndDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.endDate
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupAccountGroup) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.GroupGroupPrimer.GetFieldDeserializers()
    res["additionalObjects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupAccountGroup_additionalObjectsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalObjects(val.(GroupAccountGroup_additionalObjectsable))
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
// GetFolder gets the folder property value. The folder property
func (m *GroupAccountGroup) GetFolder()(GroupGroupFolderable) {
    return m.folder
}
// GetLastUsed gets the lastUsed property value. The lastUsed property
func (m *GroupAccountGroup) GetLastUsed()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.lastUsed
}
// GetProvisioningEndTime gets the provisioningEndTime property value. The provisioningEndTime property
func (m *GroupAccountGroup) GetProvisioningEndTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.provisioningEndTime
}
// GetRights gets the rights property value. The rights property
func (m *GroupAccountGroup) GetRights()(*GroupGroupRights) {
    return m.rights
}
// GetVisibleForProvisioning gets the visibleForProvisioning property value. The visibleForProvisioning property
func (m *GroupAccountGroup) GetVisibleForProvisioning()(*bool) {
    return m.visibleForProvisioning
}
// Serialize serializes information the current object
func (m *GroupAccountGroup) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.GroupGroupPrimer.Serialize(writer)
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
        err = writer.WriteDateOnlyValue("endDate", m.GetEndDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("folder", m.GetFolder())
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
    return nil
}
// SetAdditionalObjects sets the additionalObjects property value. The additionalObjects property
func (m *GroupAccountGroup) SetAdditionalObjects(value GroupAccountGroup_additionalObjectsable)() {
    m.additionalObjects = value
}
// SetEndDate sets the endDate property value. The endDate property
func (m *GroupAccountGroup) SetEndDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.endDate = value
}
// SetFolder sets the folder property value. The folder property
func (m *GroupAccountGroup) SetFolder(value GroupGroupFolderable)() {
    m.folder = value
}
// SetLastUsed sets the lastUsed property value. The lastUsed property
func (m *GroupAccountGroup) SetLastUsed(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.lastUsed = value
}
// SetProvisioningEndTime sets the provisioningEndTime property value. The provisioningEndTime property
func (m *GroupAccountGroup) SetProvisioningEndTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.provisioningEndTime = value
}
// SetRights sets the rights property value. The rights property
func (m *GroupAccountGroup) SetRights(value *GroupGroupRights)() {
    m.rights = value
}
// SetVisibleForProvisioning sets the visibleForProvisioning property value. The visibleForProvisioning property
func (m *GroupAccountGroup) SetVisibleForProvisioning(value *bool)() {
    m.visibleForProvisioning = value
}
// GroupAccountGroupable 
type GroupAccountGroupable interface {
    GroupGroupPrimerable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdditionalObjects()(GroupAccountGroup_additionalObjectsable)
    GetEndDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetFolder()(GroupGroupFolderable)
    GetLastUsed()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetProvisioningEndTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRights()(*GroupGroupRights)
    GetVisibleForProvisioning()(*bool)
    SetAdditionalObjects(value GroupAccountGroup_additionalObjectsable)()
    SetEndDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetFolder(value GroupGroupFolderable)()
    SetLastUsed(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetProvisioningEndTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRights(value *GroupGroupRights)()
    SetVisibleForProvisioning(value *bool)()
}
