package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OrganizationOrganizationalUnit_additionalObjects struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The audit property
    audit AuditInfoable
    // The createAsParentOf property
    createAsParentOf OrganizationOrganizationalUnitPrimerLinkableWrapperable
}
// NewOrganizationOrganizationalUnit_additionalObjects instantiates a new OrganizationOrganizationalUnit_additionalObjects and sets the default values.
func NewOrganizationOrganizationalUnit_additionalObjects()(*OrganizationOrganizationalUnit_additionalObjects) {
    m := &OrganizationOrganizationalUnit_additionalObjects{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateOrganizationOrganizationalUnit_additionalObjectsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOrganizationOrganizationalUnit_additionalObjectsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOrganizationOrganizationalUnit_additionalObjects(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *OrganizationOrganizationalUnit_additionalObjects) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAudit gets the audit property value. The audit property
// returns a AuditInfoable when successful
func (m *OrganizationOrganizationalUnit_additionalObjects) GetAudit()(AuditInfoable) {
    return m.audit
}
// GetCreateAsParentOf gets the createAsParentOf property value. The createAsParentOf property
// returns a OrganizationOrganizationalUnitPrimerLinkableWrapperable when successful
func (m *OrganizationOrganizationalUnit_additionalObjects) GetCreateAsParentOf()(OrganizationOrganizationalUnitPrimerLinkableWrapperable) {
    return m.createAsParentOf
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OrganizationOrganizationalUnit_additionalObjects) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["createAsParentOf"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOrganizationOrganizationalUnitPrimerLinkableWrapperFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreateAsParentOf(val.(OrganizationOrganizationalUnitPrimerLinkableWrapperable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *OrganizationOrganizationalUnit_additionalObjects) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("audit", m.GetAudit())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("createAsParentOf", m.GetCreateAsParentOf())
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
func (m *OrganizationOrganizationalUnit_additionalObjects) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAudit sets the audit property value. The audit property
func (m *OrganizationOrganizationalUnit_additionalObjects) SetAudit(value AuditInfoable)() {
    m.audit = value
}
// SetCreateAsParentOf sets the createAsParentOf property value. The createAsParentOf property
func (m *OrganizationOrganizationalUnit_additionalObjects) SetCreateAsParentOf(value OrganizationOrganizationalUnitPrimerLinkableWrapperable)() {
    m.createAsParentOf = value
}
type OrganizationOrganizationalUnit_additionalObjectsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAudit()(AuditInfoable)
    GetCreateAsParentOf()(OrganizationOrganizationalUnitPrimerLinkableWrapperable)
    SetAudit(value AuditInfoable)()
    SetCreateAsParentOf(value OrganizationOrganizationalUnitPrimerLinkableWrapperable)()
}
