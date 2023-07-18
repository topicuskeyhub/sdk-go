package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OrganizationOrganizationalUnitAccount_additionalObjects 
type OrganizationOrganizationalUnitAccount_additionalObjects struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The audit property
    audit AuditInfoable
}
// NewOrganizationOrganizationalUnitAccount_additionalObjects instantiates a new organizationOrganizationalUnitAccount_additionalObjects and sets the default values.
func NewOrganizationOrganizationalUnitAccount_additionalObjects()(*OrganizationOrganizationalUnitAccount_additionalObjects) {
    m := &OrganizationOrganizationalUnitAccount_additionalObjects{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateOrganizationOrganizationalUnitAccount_additionalObjectsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOrganizationOrganizationalUnitAccount_additionalObjectsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOrganizationOrganizationalUnitAccount_additionalObjects(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OrganizationOrganizationalUnitAccount_additionalObjects) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAudit gets the audit property value. The audit property
func (m *OrganizationOrganizationalUnitAccount_additionalObjects) GetAudit()(AuditInfoable) {
    return m.audit
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OrganizationOrganizationalUnitAccount_additionalObjects) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    return res
}
// Serialize serializes information the current object
func (m *OrganizationOrganizationalUnitAccount_additionalObjects) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("audit", m.GetAudit())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OrganizationOrganizationalUnitAccount_additionalObjects) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAudit sets the audit property value. The audit property
func (m *OrganizationOrganizationalUnitAccount_additionalObjects) SetAudit(value AuditInfoable)() {
    m.audit = value
}
// OrganizationOrganizationalUnitAccount_additionalObjectsable 
type OrganizationOrganizationalUnitAccount_additionalObjectsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAudit()(AuditInfoable)
    SetAudit(value AuditInfoable)()
}
