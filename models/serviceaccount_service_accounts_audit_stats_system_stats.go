package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServiceaccountServiceAccountsAuditStats_systemStats 
type ServiceaccountServiceAccountsAuditStats_systemStats struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
}
// NewServiceaccountServiceAccountsAuditStats_systemStats instantiates a new serviceaccountServiceAccountsAuditStats_systemStats and sets the default values.
func NewServiceaccountServiceAccountsAuditStats_systemStats()(*ServiceaccountServiceAccountsAuditStats_systemStats) {
    m := &ServiceaccountServiceAccountsAuditStats_systemStats{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateServiceaccountServiceAccountsAuditStats_systemStatsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateServiceaccountServiceAccountsAuditStats_systemStatsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServiceaccountServiceAccountsAuditStats_systemStats(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ServiceaccountServiceAccountsAuditStats_systemStats) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ServiceaccountServiceAccountsAuditStats_systemStats) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    return res
}
// Serialize serializes information the current object
func (m *ServiceaccountServiceAccountsAuditStats_systemStats) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ServiceaccountServiceAccountsAuditStats_systemStats) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// ServiceaccountServiceAccountsAuditStats_systemStatsable 
type ServiceaccountServiceAccountsAuditStats_systemStatsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
