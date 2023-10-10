package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProvisioningProvisionNumberSequence_additionalObjects 
type ProvisioningProvisionNumberSequence_additionalObjects struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The audit property
    audit AuditInfoable
    // The systems property
    systems ProvisioningProvisionedSystemPrimerLinkableWrapperable
}
// NewProvisioningProvisionNumberSequence_additionalObjects instantiates a new provisioningProvisionNumberSequence_additionalObjects and sets the default values.
func NewProvisioningProvisionNumberSequence_additionalObjects()(*ProvisioningProvisionNumberSequence_additionalObjects) {
    m := &ProvisioningProvisionNumberSequence_additionalObjects{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateProvisioningProvisionNumberSequence_additionalObjectsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProvisioningProvisionNumberSequence_additionalObjectsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProvisioningProvisionNumberSequence_additionalObjects(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ProvisioningProvisionNumberSequence_additionalObjects) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAudit gets the audit property value. The audit property
func (m *ProvisioningProvisionNumberSequence_additionalObjects) GetAudit()(AuditInfoable) {
    return m.audit
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProvisioningProvisionNumberSequence_additionalObjects) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["audit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuditInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAudit(val.(AuditInfoable))
        }
        return nil
    }
    res["systems"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProvisioningProvisionedSystemPrimerLinkableWrapperFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystems(val.(ProvisioningProvisionedSystemPrimerLinkableWrapperable))
        }
        return nil
    }
    return res
}
// GetSystems gets the systems property value. The systems property
func (m *ProvisioningProvisionNumberSequence_additionalObjects) GetSystems()(ProvisioningProvisionedSystemPrimerLinkableWrapperable) {
    return m.systems
}
// Serialize serializes information the current object
func (m *ProvisioningProvisionNumberSequence_additionalObjects) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("audit", m.GetAudit())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("systems", m.GetSystems())
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
func (m *ProvisioningProvisionNumberSequence_additionalObjects) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAudit sets the audit property value. The audit property
func (m *ProvisioningProvisionNumberSequence_additionalObjects) SetAudit(value AuditInfoable)() {
    m.audit = value
}
// SetSystems sets the systems property value. The systems property
func (m *ProvisioningProvisionNumberSequence_additionalObjects) SetSystems(value ProvisioningProvisionedSystemPrimerLinkableWrapperable)() {
    m.systems = value
}
// ProvisioningProvisionNumberSequence_additionalObjectsable 
type ProvisioningProvisionNumberSequence_additionalObjectsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAudit()(AuditInfoable)
    GetSystems()(ProvisioningProvisionedSystemPrimerLinkableWrapperable)
    SetAudit(value AuditInfoable)()
    SetSystems(value ProvisioningProvisionedSystemPrimerLinkableWrapperable)()
}
