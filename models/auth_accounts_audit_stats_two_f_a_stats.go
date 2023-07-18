package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthAccountsAuditStats_twoFAStats 
type AuthAccountsAuditStats_twoFAStats struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
}
// NewAuthAccountsAuditStats_twoFAStats instantiates a new authAccountsAuditStats_twoFAStats and sets the default values.
func NewAuthAccountsAuditStats_twoFAStats()(*AuthAccountsAuditStats_twoFAStats) {
    m := &AuthAccountsAuditStats_twoFAStats{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAuthAccountsAuditStats_twoFAStatsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthAccountsAuditStats_twoFAStatsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthAccountsAuditStats_twoFAStats(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AuthAccountsAuditStats_twoFAStats) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthAccountsAuditStats_twoFAStats) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    return res
}
// Serialize serializes information the current object
func (m *AuthAccountsAuditStats_twoFAStats) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AuthAccountsAuditStats_twoFAStats) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// AuthAccountsAuditStats_twoFAStatsable 
type AuthAccountsAuditStats_twoFAStatsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
