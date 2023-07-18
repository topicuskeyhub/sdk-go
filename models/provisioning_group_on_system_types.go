package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProvisioningGroupOnSystemTypes 
type ProvisioningGroupOnSystemTypes struct {
    NonLinkable
    // The types property
    types []ProvisioningGroupOnSystemType
}
// NewProvisioningGroupOnSystemTypes instantiates a new provisioningGroupOnSystemTypes and sets the default values.
func NewProvisioningGroupOnSystemTypes()(*ProvisioningGroupOnSystemTypes) {
    m := &ProvisioningGroupOnSystemTypes{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "provisioning.GroupOnSystemTypes"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateProvisioningGroupOnSystemTypesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProvisioningGroupOnSystemTypesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProvisioningGroupOnSystemTypes(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProvisioningGroupOnSystemTypes) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
    res["types"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseProvisioningGroupOnSystemType)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ProvisioningGroupOnSystemType, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*ProvisioningGroupOnSystemType))
                }
            }
            m.SetTypes(res)
        }
        return nil
    }
    return res
}
// GetTypes gets the types property value. The types property
func (m *ProvisioningGroupOnSystemTypes) GetTypes()([]ProvisioningGroupOnSystemType) {
    return m.types
}
// Serialize serializes information the current object
func (m *ProvisioningGroupOnSystemTypes) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetTypes() != nil {
        err = writer.WriteCollectionOfStringValues("types", SerializeProvisioningGroupOnSystemType(m.GetTypes()))
        if err != nil {
            return err
        }
    }
    return nil
}
// SetTypes sets the types property value. The types property
func (m *ProvisioningGroupOnSystemTypes) SetTypes(value []ProvisioningGroupOnSystemType)() {
    m.types = value
}
// ProvisioningGroupOnSystemTypesable 
type ProvisioningGroupOnSystemTypesable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetTypes()([]ProvisioningGroupOnSystemType)
    SetTypes(value []ProvisioningGroupOnSystemType)()
}
