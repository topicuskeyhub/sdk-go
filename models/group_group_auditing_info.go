package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupGroupAuditingInfo 
type GroupGroupAuditingInfo struct {
    NonLinkable
    // The auditDueDate property
    auditDueDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The lastAuditDate property
    lastAuditDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The nrAccounts property
    nrAccounts *int64
    // The nrDisabledAccounts property
    nrDisabledAccounts *int64
    // The nrDisabledManagers property
    nrDisabledManagers *int64
    // The nrExpiredVaultRecords property
    nrExpiredVaultRecords *int64
    // The nrManagers property
    nrManagers *int64
    // The nrVaultRecordsWithEndDate property
    nrVaultRecordsWithEndDate *int64
}
// NewGroupGroupAuditingInfo instantiates a new groupGroupAuditingInfo and sets the default values.
func NewGroupGroupAuditingInfo()(*GroupGroupAuditingInfo) {
    m := &GroupGroupAuditingInfo{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "group.GroupAuditingInfo"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateGroupGroupAuditingInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupGroupAuditingInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupGroupAuditingInfo(), nil
}
// GetAuditDueDate gets the auditDueDate property value. The auditDueDate property
func (m *GroupGroupAuditingInfo) GetAuditDueDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.auditDueDate
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupGroupAuditingInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
    res["auditDueDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuditDueDate(val)
        }
        return nil
    }
    res["lastAuditDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastAuditDate(val)
        }
        return nil
    }
    res["nrAccounts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNrAccounts(val)
        }
        return nil
    }
    res["nrDisabledAccounts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNrDisabledAccounts(val)
        }
        return nil
    }
    res["nrDisabledManagers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNrDisabledManagers(val)
        }
        return nil
    }
    res["nrExpiredVaultRecords"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNrExpiredVaultRecords(val)
        }
        return nil
    }
    res["nrManagers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNrManagers(val)
        }
        return nil
    }
    res["nrVaultRecordsWithEndDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNrVaultRecordsWithEndDate(val)
        }
        return nil
    }
    return res
}
// GetLastAuditDate gets the lastAuditDate property value. The lastAuditDate property
func (m *GroupGroupAuditingInfo) GetLastAuditDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.lastAuditDate
}
// GetNrAccounts gets the nrAccounts property value. The nrAccounts property
func (m *GroupGroupAuditingInfo) GetNrAccounts()(*int64) {
    return m.nrAccounts
}
// GetNrDisabledAccounts gets the nrDisabledAccounts property value. The nrDisabledAccounts property
func (m *GroupGroupAuditingInfo) GetNrDisabledAccounts()(*int64) {
    return m.nrDisabledAccounts
}
// GetNrDisabledManagers gets the nrDisabledManagers property value. The nrDisabledManagers property
func (m *GroupGroupAuditingInfo) GetNrDisabledManagers()(*int64) {
    return m.nrDisabledManagers
}
// GetNrExpiredVaultRecords gets the nrExpiredVaultRecords property value. The nrExpiredVaultRecords property
func (m *GroupGroupAuditingInfo) GetNrExpiredVaultRecords()(*int64) {
    return m.nrExpiredVaultRecords
}
// GetNrManagers gets the nrManagers property value. The nrManagers property
func (m *GroupGroupAuditingInfo) GetNrManagers()(*int64) {
    return m.nrManagers
}
// GetNrVaultRecordsWithEndDate gets the nrVaultRecordsWithEndDate property value. The nrVaultRecordsWithEndDate property
func (m *GroupGroupAuditingInfo) GetNrVaultRecordsWithEndDate()(*int64) {
    return m.nrVaultRecordsWithEndDate
}
// Serialize serializes information the current object
func (m *GroupGroupAuditingInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// SetAuditDueDate sets the auditDueDate property value. The auditDueDate property
func (m *GroupGroupAuditingInfo) SetAuditDueDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.auditDueDate = value
}
// SetLastAuditDate sets the lastAuditDate property value. The lastAuditDate property
func (m *GroupGroupAuditingInfo) SetLastAuditDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.lastAuditDate = value
}
// SetNrAccounts sets the nrAccounts property value. The nrAccounts property
func (m *GroupGroupAuditingInfo) SetNrAccounts(value *int64)() {
    m.nrAccounts = value
}
// SetNrDisabledAccounts sets the nrDisabledAccounts property value. The nrDisabledAccounts property
func (m *GroupGroupAuditingInfo) SetNrDisabledAccounts(value *int64)() {
    m.nrDisabledAccounts = value
}
// SetNrDisabledManagers sets the nrDisabledManagers property value. The nrDisabledManagers property
func (m *GroupGroupAuditingInfo) SetNrDisabledManagers(value *int64)() {
    m.nrDisabledManagers = value
}
// SetNrExpiredVaultRecords sets the nrExpiredVaultRecords property value. The nrExpiredVaultRecords property
func (m *GroupGroupAuditingInfo) SetNrExpiredVaultRecords(value *int64)() {
    m.nrExpiredVaultRecords = value
}
// SetNrManagers sets the nrManagers property value. The nrManagers property
func (m *GroupGroupAuditingInfo) SetNrManagers(value *int64)() {
    m.nrManagers = value
}
// SetNrVaultRecordsWithEndDate sets the nrVaultRecordsWithEndDate property value. The nrVaultRecordsWithEndDate property
func (m *GroupGroupAuditingInfo) SetNrVaultRecordsWithEndDate(value *int64)() {
    m.nrVaultRecordsWithEndDate = value
}
// GroupGroupAuditingInfoable 
type GroupGroupAuditingInfoable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuditDueDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetLastAuditDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetNrAccounts()(*int64)
    GetNrDisabledAccounts()(*int64)
    GetNrDisabledManagers()(*int64)
    GetNrExpiredVaultRecords()(*int64)
    GetNrManagers()(*int64)
    GetNrVaultRecordsWithEndDate()(*int64)
    SetAuditDueDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetLastAuditDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetNrAccounts(value *int64)()
    SetNrDisabledAccounts(value *int64)()
    SetNrDisabledManagers(value *int64)()
    SetNrExpiredVaultRecords(value *int64)()
    SetNrManagers(value *int64)()
    SetNrVaultRecordsWithEndDate(value *int64)()
}
