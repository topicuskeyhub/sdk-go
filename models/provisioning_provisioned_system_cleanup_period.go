package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ProvisioningProvisionedSystem_cleanupPeriod struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The days property
    days *int32
    // The months property
    months *int32
    // The years property
    years *int32
}
// NewProvisioningProvisionedSystem_cleanupPeriod instantiates a new ProvisioningProvisionedSystem_cleanupPeriod and sets the default values.
func NewProvisioningProvisionedSystem_cleanupPeriod()(*ProvisioningProvisionedSystem_cleanupPeriod) {
    m := &ProvisioningProvisionedSystem_cleanupPeriod{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateProvisioningProvisionedSystem_cleanupPeriodFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateProvisioningProvisionedSystem_cleanupPeriodFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProvisioningProvisionedSystem_cleanupPeriod(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ProvisioningProvisionedSystem_cleanupPeriod) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDays gets the days property value. The days property
// returns a *int32 when successful
func (m *ProvisioningProvisionedSystem_cleanupPeriod) GetDays()(*int32) {
    return m.days
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ProvisioningProvisionedSystem_cleanupPeriod) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["days"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDays(val)
        }
        return nil
    }
    res["months"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMonths(val)
        }
        return nil
    }
    res["years"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetYears(val)
        }
        return nil
    }
    return res
}
// GetMonths gets the months property value. The months property
// returns a *int32 when successful
func (m *ProvisioningProvisionedSystem_cleanupPeriod) GetMonths()(*int32) {
    return m.months
}
// GetYears gets the years property value. The years property
// returns a *int32 when successful
func (m *ProvisioningProvisionedSystem_cleanupPeriod) GetYears()(*int32) {
    return m.years
}
// Serialize serializes information the current object
func (m *ProvisioningProvisionedSystem_cleanupPeriod) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("days", m.GetDays())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("months", m.GetMonths())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("years", m.GetYears())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ProvisioningProvisionedSystem_cleanupPeriod) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDays sets the days property value. The days property
func (m *ProvisioningProvisionedSystem_cleanupPeriod) SetDays(value *int32)() {
    m.days = value
}
// SetMonths sets the months property value. The months property
func (m *ProvisioningProvisionedSystem_cleanupPeriod) SetMonths(value *int32)() {
    m.months = value
}
// SetYears sets the years property value. The years property
func (m *ProvisioningProvisionedSystem_cleanupPeriod) SetYears(value *int32)() {
    m.years = value
}
type ProvisioningProvisionedSystem_cleanupPeriodable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDays()(*int32)
    GetMonths()(*int32)
    GetYears()(*int32)
    SetDays(value *int32)()
    SetMonths(value *int32)()
    SetYears(value *int32)()
}
