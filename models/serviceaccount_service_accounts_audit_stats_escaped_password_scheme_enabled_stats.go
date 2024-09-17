package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ServiceaccountServiceAccountsAuditStats_passwordSchemeEnabledStats struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
}
// NewServiceaccountServiceAccountsAuditStats_passwordSchemeEnabledStats instantiates a new ServiceaccountServiceAccountsAuditStats_passwordSchemeEnabledStats and sets the default values.
func NewServiceaccountServiceAccountsAuditStats_passwordSchemeEnabledStats()(*ServiceaccountServiceAccountsAuditStats_passwordSchemeEnabledStats) {
    m := &ServiceaccountServiceAccountsAuditStats_passwordSchemeEnabledStats{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateServiceaccountServiceAccountsAuditStats_passwordSchemeEnabledStatsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateServiceaccountServiceAccountsAuditStats_passwordSchemeEnabledStatsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServiceaccountServiceAccountsAuditStats_passwordSchemeEnabledStats(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ServiceaccountServiceAccountsAuditStats_passwordSchemeEnabledStats) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ServiceaccountServiceAccountsAuditStats_passwordSchemeEnabledStats) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    return res
}
// Serialize serializes information the current object
func (m *ServiceaccountServiceAccountsAuditStats_passwordSchemeEnabledStats) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ServiceaccountServiceAccountsAuditStats_passwordSchemeEnabledStats) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
type ServiceaccountServiceAccountsAuditStats_passwordSchemeEnabledStatsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
