package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthAccountsAuditStats 
type AuthAccountsAuditStats struct {
    NonLinkable
    // The directoryStats property
    directoryStats AuthAccountsAuditStats_directoryStatsable
    // The twoFAStats property
    twoFAStats AuthAccountsAuditStats_twoFAStatsable
    // The validityStats property
    validityStats AuthAccountsAuditStats_validityStatsable
}
// NewAuthAccountsAuditStats instantiates a new authAccountsAuditStats and sets the default values.
func NewAuthAccountsAuditStats()(*AuthAccountsAuditStats) {
    m := &AuthAccountsAuditStats{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "auth.AccountsAuditStats"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateAuthAccountsAuditStatsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthAccountsAuditStatsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthAccountsAuditStats(), nil
}
// GetDirectoryStats gets the directoryStats property value. The directoryStats property
func (m *AuthAccountsAuditStats) GetDirectoryStats()(AuthAccountsAuditStats_directoryStatsable) {
    return m.directoryStats
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthAccountsAuditStats) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
    res["directoryStats"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthAccountsAuditStats_directoryStatsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDirectoryStats(val.(AuthAccountsAuditStats_directoryStatsable))
        }
        return nil
    }
    res["twoFAStats"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthAccountsAuditStats_twoFAStatsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTwoFAStats(val.(AuthAccountsAuditStats_twoFAStatsable))
        }
        return nil
    }
    res["validityStats"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthAccountsAuditStats_validityStatsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValidityStats(val.(AuthAccountsAuditStats_validityStatsable))
        }
        return nil
    }
    return res
}
// GetTwoFAStats gets the twoFAStats property value. The twoFAStats property
func (m *AuthAccountsAuditStats) GetTwoFAStats()(AuthAccountsAuditStats_twoFAStatsable) {
    return m.twoFAStats
}
// GetValidityStats gets the validityStats property value. The validityStats property
func (m *AuthAccountsAuditStats) GetValidityStats()(AuthAccountsAuditStats_validityStatsable) {
    return m.validityStats
}
// Serialize serializes information the current object
func (m *AuthAccountsAuditStats) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("directoryStats", m.GetDirectoryStats())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("twoFAStats", m.GetTwoFAStats())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("validityStats", m.GetValidityStats())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDirectoryStats sets the directoryStats property value. The directoryStats property
func (m *AuthAccountsAuditStats) SetDirectoryStats(value AuthAccountsAuditStats_directoryStatsable)() {
    m.directoryStats = value
}
// SetTwoFAStats sets the twoFAStats property value. The twoFAStats property
func (m *AuthAccountsAuditStats) SetTwoFAStats(value AuthAccountsAuditStats_twoFAStatsable)() {
    m.twoFAStats = value
}
// SetValidityStats sets the validityStats property value. The validityStats property
func (m *AuthAccountsAuditStats) SetValidityStats(value AuthAccountsAuditStats_validityStatsable)() {
    m.validityStats = value
}
// AuthAccountsAuditStatsable 
type AuthAccountsAuditStatsable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDirectoryStats()(AuthAccountsAuditStats_directoryStatsable)
    GetTwoFAStats()(AuthAccountsAuditStats_twoFAStatsable)
    GetValidityStats()(AuthAccountsAuditStats_validityStatsable)
    SetDirectoryStats(value AuthAccountsAuditStats_directoryStatsable)()
    SetTwoFAStats(value AuthAccountsAuditStats_twoFAStatsable)()
    SetValidityStats(value AuthAccountsAuditStats_validityStatsable)()
}
