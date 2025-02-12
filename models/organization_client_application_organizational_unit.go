package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OrganizationClientApplicationOrganizationalUnit struct {
    OrganizationOrganizationalUnitPrimer
}
// NewOrganizationClientApplicationOrganizationalUnit instantiates a new OrganizationClientApplicationOrganizationalUnit and sets the default values.
func NewOrganizationClientApplicationOrganizationalUnit()(*OrganizationClientApplicationOrganizationalUnit) {
    m := &OrganizationClientApplicationOrganizationalUnit{
        OrganizationOrganizationalUnitPrimer: *NewOrganizationOrganizationalUnitPrimer(),
    }
    typeEscapedValue := "organization.ClientApplicationOrganizationalUnit"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateOrganizationClientApplicationOrganizationalUnitFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOrganizationClientApplicationOrganizationalUnitFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOrganizationClientApplicationOrganizationalUnit(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OrganizationClientApplicationOrganizationalUnit) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.OrganizationOrganizationalUnitPrimer.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *OrganizationClientApplicationOrganizationalUnit) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.OrganizationOrganizationalUnitPrimer.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type OrganizationClientApplicationOrganizationalUnitable interface {
    OrganizationOrganizationalUnitPrimerable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
