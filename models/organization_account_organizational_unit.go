package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OrganizationAccountOrganizationalUnit 
type OrganizationAccountOrganizationalUnit struct {
    OrganizationOrganizationalUnitPrimer
    // The additionalObjects property
    additionalObjects OrganizationAccountOrganizationalUnit_additionalObjectsable
}
// NewOrganizationAccountOrganizationalUnit instantiates a new organizationAccountOrganizationalUnit and sets the default values.
func NewOrganizationAccountOrganizationalUnit()(*OrganizationAccountOrganizationalUnit) {
    m := &OrganizationAccountOrganizationalUnit{
        OrganizationOrganizationalUnitPrimer: *NewOrganizationOrganizationalUnitPrimer(),
    }
    typeEscapedValue := "organization.AccountOrganizationalUnit"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateOrganizationAccountOrganizationalUnitFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOrganizationAccountOrganizationalUnitFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOrganizationAccountOrganizationalUnit(), nil
}
// GetAdditionalObjects gets the additionalObjects property value. The additionalObjects property
func (m *OrganizationAccountOrganizationalUnit) GetAdditionalObjects()(OrganizationAccountOrganizationalUnit_additionalObjectsable) {
    return m.additionalObjects
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OrganizationAccountOrganizationalUnit) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.OrganizationOrganizationalUnitPrimer.GetFieldDeserializers()
    res["additionalObjects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOrganizationAccountOrganizationalUnit_additionalObjectsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalObjects(val.(OrganizationAccountOrganizationalUnit_additionalObjectsable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *OrganizationAccountOrganizationalUnit) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.OrganizationOrganizationalUnitPrimer.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("additionalObjects", m.GetAdditionalObjects())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalObjects sets the additionalObjects property value. The additionalObjects property
func (m *OrganizationAccountOrganizationalUnit) SetAdditionalObjects(value OrganizationAccountOrganizationalUnit_additionalObjectsable)() {
    m.additionalObjects = value
}
// OrganizationAccountOrganizationalUnitable 
type OrganizationAccountOrganizationalUnitable interface {
    OrganizationOrganizationalUnitPrimerable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdditionalObjects()(OrganizationAccountOrganizationalUnit_additionalObjectsable)
    SetAdditionalObjects(value OrganizationAccountOrganizationalUnit_additionalObjectsable)()
}
