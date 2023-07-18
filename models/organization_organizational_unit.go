package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OrganizationOrganizationalUnit 
type OrganizationOrganizationalUnit struct {
    OrganizationOrganizationalUnitPrimer
    // The additionalObjects property
    additionalObjects OrganizationOrganizationalUnit_additionalObjectsable
    // The depth property
    depth *int32
    // The description property
    description *string
    // The owner property
    owner GroupGroupPrimerable
    // The parent property
    parent OrganizationOrganizationalUnitPrimerable
}
// NewOrganizationOrganizationalUnit instantiates a new organizationOrganizationalUnit and sets the default values.
func NewOrganizationOrganizationalUnit()(*OrganizationOrganizationalUnit) {
    m := &OrganizationOrganizationalUnit{
        OrganizationOrganizationalUnitPrimer: *NewOrganizationOrganizationalUnitPrimer(),
    }
    typeEscapedValue := "organization.OrganizationalUnit"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateOrganizationOrganizationalUnitFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOrganizationOrganizationalUnitFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOrganizationOrganizationalUnit(), nil
}
// GetAdditionalObjects gets the additionalObjects property value. The additionalObjects property
func (m *OrganizationOrganizationalUnit) GetAdditionalObjects()(OrganizationOrganizationalUnit_additionalObjectsable) {
    return m.additionalObjects
}
// GetDepth gets the depth property value. The depth property
func (m *OrganizationOrganizationalUnit) GetDepth()(*int32) {
    return m.depth
}
// GetDescription gets the description property value. The description property
func (m *OrganizationOrganizationalUnit) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OrganizationOrganizationalUnit) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.OrganizationOrganizationalUnitPrimer.GetFieldDeserializers()
    res["additionalObjects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOrganizationOrganizationalUnit_additionalObjectsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalObjects(val.(OrganizationOrganizationalUnit_additionalObjectsable))
        }
        return nil
    }
    res["depth"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDepth(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["owner"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupGroupPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwner(val.(GroupGroupPrimerable))
        }
        return nil
    }
    res["parent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOrganizationOrganizationalUnitPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParent(val.(OrganizationOrganizationalUnitPrimerable))
        }
        return nil
    }
    return res
}
// GetOwner gets the owner property value. The owner property
func (m *OrganizationOrganizationalUnit) GetOwner()(GroupGroupPrimerable) {
    return m.owner
}
// GetParent gets the parent property value. The parent property
func (m *OrganizationOrganizationalUnit) GetParent()(OrganizationOrganizationalUnitPrimerable) {
    return m.parent
}
// Serialize serializes information the current object
func (m *OrganizationOrganizationalUnit) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("owner", m.GetOwner())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("parent", m.GetParent())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalObjects sets the additionalObjects property value. The additionalObjects property
func (m *OrganizationOrganizationalUnit) SetAdditionalObjects(value OrganizationOrganizationalUnit_additionalObjectsable)() {
    m.additionalObjects = value
}
// SetDepth sets the depth property value. The depth property
func (m *OrganizationOrganizationalUnit) SetDepth(value *int32)() {
    m.depth = value
}
// SetDescription sets the description property value. The description property
func (m *OrganizationOrganizationalUnit) SetDescription(value *string)() {
    m.description = value
}
// SetOwner sets the owner property value. The owner property
func (m *OrganizationOrganizationalUnit) SetOwner(value GroupGroupPrimerable)() {
    m.owner = value
}
// SetParent sets the parent property value. The parent property
func (m *OrganizationOrganizationalUnit) SetParent(value OrganizationOrganizationalUnitPrimerable)() {
    m.parent = value
}
// OrganizationOrganizationalUnitable 
type OrganizationOrganizationalUnitable interface {
    OrganizationOrganizationalUnitPrimerable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdditionalObjects()(OrganizationOrganizationalUnit_additionalObjectsable)
    GetDepth()(*int32)
    GetDescription()(*string)
    GetOwner()(GroupGroupPrimerable)
    GetParent()(OrganizationOrganizationalUnitPrimerable)
    SetAdditionalObjects(value OrganizationOrganizationalUnit_additionalObjectsable)()
    SetDepth(value *int32)()
    SetDescription(value *string)()
    SetOwner(value GroupGroupPrimerable)()
    SetParent(value OrganizationOrganizationalUnitPrimerable)()
}
