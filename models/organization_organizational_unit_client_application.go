package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OrganizationOrganizationalUnitClientApplication struct {
    ClientClientApplicationPrimer
    // The additionalObjects property
    additionalObjects OrganizationOrganizationalUnitClientApplication_additionalObjectsable
}
// NewOrganizationOrganizationalUnitClientApplication instantiates a new OrganizationOrganizationalUnitClientApplication and sets the default values.
func NewOrganizationOrganizationalUnitClientApplication()(*OrganizationOrganizationalUnitClientApplication) {
    m := &OrganizationOrganizationalUnitClientApplication{
        ClientClientApplicationPrimer: *NewClientClientApplicationPrimer(),
    }
    typeEscapedValue := "organization.OrganizationalUnitClientApplication"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateOrganizationOrganizationalUnitClientApplicationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOrganizationOrganizationalUnitClientApplicationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOrganizationOrganizationalUnitClientApplication(), nil
}
// GetAdditionalObjects gets the additionalObjects property value. The additionalObjects property
// returns a OrganizationOrganizationalUnitClientApplication_additionalObjectsable when successful
func (m *OrganizationOrganizationalUnitClientApplication) GetAdditionalObjects()(OrganizationOrganizationalUnitClientApplication_additionalObjectsable) {
    return m.additionalObjects
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OrganizationOrganizationalUnitClientApplication) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ClientClientApplicationPrimer.GetFieldDeserializers()
    res["additionalObjects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOrganizationOrganizationalUnitClientApplication_additionalObjectsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalObjects(val.(OrganizationOrganizationalUnitClientApplication_additionalObjectsable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *OrganizationOrganizationalUnitClientApplication) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ClientClientApplicationPrimer.Serialize(writer)
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
func (m *OrganizationOrganizationalUnitClientApplication) SetAdditionalObjects(value OrganizationOrganizationalUnitClientApplication_additionalObjectsable)() {
    m.additionalObjects = value
}
type OrganizationOrganizationalUnitClientApplicationable interface {
    ClientClientApplicationPrimerable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdditionalObjects()(OrganizationOrganizationalUnitClientApplication_additionalObjectsable)
    SetAdditionalObjects(value OrganizationOrganizationalUnitClientApplication_additionalObjectsable)()
}
