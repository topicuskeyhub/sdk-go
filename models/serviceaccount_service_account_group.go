package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServiceaccountServiceAccountGroup 
type ServiceaccountServiceAccountGroup struct {
    ProvisioningGroupOnSystemPrimer
    // The additionalObjects property
    additionalObjects ServiceaccountServiceAccountGroup_additionalObjectsable
}
// NewServiceaccountServiceAccountGroup instantiates a new serviceaccountServiceAccountGroup and sets the default values.
func NewServiceaccountServiceAccountGroup()(*ServiceaccountServiceAccountGroup) {
    m := &ServiceaccountServiceAccountGroup{
        ProvisioningGroupOnSystemPrimer: *NewProvisioningGroupOnSystemPrimer(),
    }
    typeEscapedValue := "serviceaccount.ServiceAccountGroup"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateServiceaccountServiceAccountGroupFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateServiceaccountServiceAccountGroupFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServiceaccountServiceAccountGroup(), nil
}
// GetAdditionalObjects gets the additionalObjects property value. The additionalObjects property
func (m *ServiceaccountServiceAccountGroup) GetAdditionalObjects()(ServiceaccountServiceAccountGroup_additionalObjectsable) {
    return m.additionalObjects
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ServiceaccountServiceAccountGroup) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ProvisioningGroupOnSystemPrimer.GetFieldDeserializers()
    res["additionalObjects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateServiceaccountServiceAccountGroup_additionalObjectsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalObjects(val.(ServiceaccountServiceAccountGroup_additionalObjectsable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ServiceaccountServiceAccountGroup) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ProvisioningGroupOnSystemPrimer.Serialize(writer)
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
func (m *ServiceaccountServiceAccountGroup) SetAdditionalObjects(value ServiceaccountServiceAccountGroup_additionalObjectsable)() {
    m.additionalObjects = value
}
// ServiceaccountServiceAccountGroupable 
type ServiceaccountServiceAccountGroupable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ProvisioningGroupOnSystemPrimerable
    GetAdditionalObjects()(ServiceaccountServiceAccountGroup_additionalObjectsable)
    SetAdditionalObjects(value ServiceaccountServiceAccountGroup_additionalObjectsable)()
}
