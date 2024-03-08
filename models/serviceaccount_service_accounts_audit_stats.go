package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ServiceaccountServiceAccountsAuditStats struct {
    NonLinkable
    // The disabledCount property
    disabledCount *int64
    // The passwordSchemeEnabledStats property
    passwordSchemeEnabledStats ServiceaccountServiceAccountsAuditStats_passwordSchemeEnabledStatsable
    // The systemStats property
    systemStats ServiceaccountServiceAccountsAuditStats_systemStatsable
}
// NewServiceaccountServiceAccountsAuditStats instantiates a new ServiceaccountServiceAccountsAuditStats and sets the default values.
func NewServiceaccountServiceAccountsAuditStats()(*ServiceaccountServiceAccountsAuditStats) {
    m := &ServiceaccountServiceAccountsAuditStats{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "serviceaccount.ServiceAccountsAuditStats"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateServiceaccountServiceAccountsAuditStatsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateServiceaccountServiceAccountsAuditStatsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServiceaccountServiceAccountsAuditStats(), nil
}
// GetDisabledCount gets the disabledCount property value. The disabledCount property
// returns a *int64 when successful
func (m *ServiceaccountServiceAccountsAuditStats) GetDisabledCount()(*int64) {
    return m.disabledCount
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ServiceaccountServiceAccountsAuditStats) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
    res["disabledCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisabledCount(val)
        }
        return nil
    }
    res["passwordSchemeEnabledStats"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateServiceaccountServiceAccountsAuditStats_passwordSchemeEnabledStatsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordSchemeEnabledStats(val.(ServiceaccountServiceAccountsAuditStats_passwordSchemeEnabledStatsable))
        }
        return nil
    }
    res["systemStats"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateServiceaccountServiceAccountsAuditStats_systemStatsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystemStats(val.(ServiceaccountServiceAccountsAuditStats_systemStatsable))
        }
        return nil
    }
    return res
}
// GetPasswordSchemeEnabledStats gets the passwordSchemeEnabledStats property value. The passwordSchemeEnabledStats property
// returns a ServiceaccountServiceAccountsAuditStats_passwordSchemeEnabledStatsable when successful
func (m *ServiceaccountServiceAccountsAuditStats) GetPasswordSchemeEnabledStats()(ServiceaccountServiceAccountsAuditStats_passwordSchemeEnabledStatsable) {
    return m.passwordSchemeEnabledStats
}
// GetSystemStats gets the systemStats property value. The systemStats property
// returns a ServiceaccountServiceAccountsAuditStats_systemStatsable when successful
func (m *ServiceaccountServiceAccountsAuditStats) GetSystemStats()(ServiceaccountServiceAccountsAuditStats_systemStatsable) {
    return m.systemStats
}
// Serialize serializes information the current object
func (m *ServiceaccountServiceAccountsAuditStats) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("disabledCount", m.GetDisabledCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("passwordSchemeEnabledStats", m.GetPasswordSchemeEnabledStats())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("systemStats", m.GetSystemStats())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisabledCount sets the disabledCount property value. The disabledCount property
func (m *ServiceaccountServiceAccountsAuditStats) SetDisabledCount(value *int64)() {
    m.disabledCount = value
}
// SetPasswordSchemeEnabledStats sets the passwordSchemeEnabledStats property value. The passwordSchemeEnabledStats property
func (m *ServiceaccountServiceAccountsAuditStats) SetPasswordSchemeEnabledStats(value ServiceaccountServiceAccountsAuditStats_passwordSchemeEnabledStatsable)() {
    m.passwordSchemeEnabledStats = value
}
// SetSystemStats sets the systemStats property value. The systemStats property
func (m *ServiceaccountServiceAccountsAuditStats) SetSystemStats(value ServiceaccountServiceAccountsAuditStats_systemStatsable)() {
    m.systemStats = value
}
type ServiceaccountServiceAccountsAuditStatsable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisabledCount()(*int64)
    GetPasswordSchemeEnabledStats()(ServiceaccountServiceAccountsAuditStats_passwordSchemeEnabledStatsable)
    GetSystemStats()(ServiceaccountServiceAccountsAuditStats_systemStatsable)
    SetDisabledCount(value *int64)()
    SetPasswordSchemeEnabledStats(value ServiceaccountServiceAccountsAuditStats_passwordSchemeEnabledStatsable)()
    SetSystemStats(value ServiceaccountServiceAccountsAuditStats_systemStatsable)()
}
