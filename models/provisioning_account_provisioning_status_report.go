package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProvisioningAccountProvisioningStatusReport 
type ProvisioningAccountProvisioningStatusReport struct {
    NonLinkable
    // The status property
    status []ProvisioningAccountProvisioningStatusable
}
// NewProvisioningAccountProvisioningStatusReport instantiates a new provisioningAccountProvisioningStatusReport and sets the default values.
func NewProvisioningAccountProvisioningStatusReport()(*ProvisioningAccountProvisioningStatusReport) {
    m := &ProvisioningAccountProvisioningStatusReport{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "provisioning.AccountProvisioningStatusReport"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateProvisioningAccountProvisioningStatusReportFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProvisioningAccountProvisioningStatusReportFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProvisioningAccountProvisioningStatusReport(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProvisioningAccountProvisioningStatusReport) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateProvisioningAccountProvisioningStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ProvisioningAccountProvisioningStatusable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ProvisioningAccountProvisioningStatusable)
                }
            }
            m.SetStatus(res)
        }
        return nil
    }
    return res
}
// GetStatus gets the status property value. The status property
func (m *ProvisioningAccountProvisioningStatusReport) GetStatus()([]ProvisioningAccountProvisioningStatusable) {
    return m.status
}
// Serialize serializes information the current object
func (m *ProvisioningAccountProvisioningStatusReport) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetStatus() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetStatus()))
        for i, v := range m.GetStatus() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("status", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetStatus sets the status property value. The status property
func (m *ProvisioningAccountProvisioningStatusReport) SetStatus(value []ProvisioningAccountProvisioningStatusable)() {
    m.status = value
}
// ProvisioningAccountProvisioningStatusReportable 
type ProvisioningAccountProvisioningStatusReportable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetStatus()([]ProvisioningAccountProvisioningStatusable)
    SetStatus(value []ProvisioningAccountProvisioningStatusable)()
}
