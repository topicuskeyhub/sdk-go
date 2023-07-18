package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthAccountsAuditStats_directoryStats 
type AuthAccountsAuditStats_directoryStats struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
}
// NewAuthAccountsAuditStats_directoryStats instantiates a new authAccountsAuditStats_directoryStats and sets the default values.
func NewAuthAccountsAuditStats_directoryStats()(*AuthAccountsAuditStats_directoryStats) {
    m := &AuthAccountsAuditStats_directoryStats{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAuthAccountsAuditStats_directoryStatsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthAccountsAuditStats_directoryStatsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthAccountsAuditStats_directoryStats(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AuthAccountsAuditStats_directoryStats) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthAccountsAuditStats_directoryStats) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    return res
}
// Serialize serializes information the current object
func (m *AuthAccountsAuditStats_directoryStats) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AuthAccountsAuditStats_directoryStats) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// AuthAccountsAuditStats_directoryStatsable 
type AuthAccountsAuditStats_directoryStatsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
