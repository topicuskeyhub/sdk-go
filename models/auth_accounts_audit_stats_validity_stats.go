package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthAccountsAuditStats_validityStats 
type AuthAccountsAuditStats_validityStats struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
}
// NewAuthAccountsAuditStats_validityStats instantiates a new authAccountsAuditStats_validityStats and sets the default values.
func NewAuthAccountsAuditStats_validityStats()(*AuthAccountsAuditStats_validityStats) {
    m := &AuthAccountsAuditStats_validityStats{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAuthAccountsAuditStats_validityStatsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthAccountsAuditStats_validityStatsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthAccountsAuditStats_validityStats(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AuthAccountsAuditStats_validityStats) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthAccountsAuditStats_validityStats) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    return res
}
// Serialize serializes information the current object
func (m *AuthAccountsAuditStats_validityStats) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AuthAccountsAuditStats_validityStats) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// AuthAccountsAuditStats_validityStatsable 
type AuthAccountsAuditStats_validityStatsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
