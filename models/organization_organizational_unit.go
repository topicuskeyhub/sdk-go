package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OrganizationOrganizationalUnit struct {
    OrganizationOrganizationalUnitPrimer
    // The additionalObjects property
    additionalObjects OrganizationOrganizationalUnit_additionalObjectsable
    // The createGroupApproveGroup property
    createGroupApproveGroup GroupGroupable
    // The createGroupPlaceholder property
    createGroupPlaceholder *string
    // The depth property
    depth *int32
    // The description property
    description *string
    // The enableTechAdminApproveGroup property
    enableTechAdminApproveGroup GroupGroupable
    // The owner property
    owner GroupGroupPrimerable
    // The parent property
    parent OrganizationOrganizationalUnitPrimerable
    // The removeGroupApproveGroup property
    removeGroupApproveGroup GroupGroupable
}
// NewOrganizationOrganizationalUnit instantiates a new OrganizationOrganizationalUnit and sets the default values.
func NewOrganizationOrganizationalUnit()(*OrganizationOrganizationalUnit) {
    m := &OrganizationOrganizationalUnit{
        OrganizationOrganizationalUnitPrimer: *NewOrganizationOrganizationalUnitPrimer(),
    }
    typeEscapedValue := "organization.OrganizationalUnit"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateOrganizationOrganizationalUnitFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOrganizationOrganizationalUnitFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOrganizationOrganizationalUnit(), nil
}
// GetAdditionalObjects gets the additionalObjects property value. The additionalObjects property
// returns a OrganizationOrganizationalUnit_additionalObjectsable when successful
func (m *OrganizationOrganizationalUnit) GetAdditionalObjects()(OrganizationOrganizationalUnit_additionalObjectsable) {
    return m.additionalObjects
}
// GetCreateGroupApproveGroup gets the createGroupApproveGroup property value. The createGroupApproveGroup property
// returns a GroupGroupable when successful
func (m *OrganizationOrganizationalUnit) GetCreateGroupApproveGroup()(GroupGroupable) {
    return m.createGroupApproveGroup
}
// GetCreateGroupPlaceholder gets the createGroupPlaceholder property value. The createGroupPlaceholder property
// returns a *string when successful
func (m *OrganizationOrganizationalUnit) GetCreateGroupPlaceholder()(*string) {
    return m.createGroupPlaceholder
}
// GetDepth gets the depth property value. The depth property
// returns a *int32 when successful
func (m *OrganizationOrganizationalUnit) GetDepth()(*int32) {
    return m.depth
}
// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *OrganizationOrganizationalUnit) GetDescription()(*string) {
    return m.description
}
// GetEnableTechAdminApproveGroup gets the enableTechAdminApproveGroup property value. The enableTechAdminApproveGroup property
// returns a GroupGroupable when successful
func (m *OrganizationOrganizationalUnit) GetEnableTechAdminApproveGroup()(GroupGroupable) {
    return m.enableTechAdminApproveGroup
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
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
    res["createGroupApproveGroup"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreateGroupApproveGroup(val.(GroupGroupable))
        }
        return nil
    }
    res["createGroupPlaceholder"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreateGroupPlaceholder(val)
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
    res["enableTechAdminApproveGroup"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableTechAdminApproveGroup(val.(GroupGroupable))
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
    res["removeGroupApproveGroup"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemoveGroupApproveGroup(val.(GroupGroupable))
        }
        return nil
    }
    return res
}
// GetOwner gets the owner property value. The owner property
// returns a GroupGroupPrimerable when successful
func (m *OrganizationOrganizationalUnit) GetOwner()(GroupGroupPrimerable) {
    return m.owner
}
// GetParent gets the parent property value. The parent property
// returns a OrganizationOrganizationalUnitPrimerable when successful
func (m *OrganizationOrganizationalUnit) GetParent()(OrganizationOrganizationalUnitPrimerable) {
    return m.parent
}
// GetRemoveGroupApproveGroup gets the removeGroupApproveGroup property value. The removeGroupApproveGroup property
// returns a GroupGroupable when successful
func (m *OrganizationOrganizationalUnit) GetRemoveGroupApproveGroup()(GroupGroupable) {
    return m.removeGroupApproveGroup
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
        err = writer.WriteObjectValue("createGroupApproveGroup", m.GetCreateGroupApproveGroup())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("createGroupPlaceholder", m.GetCreateGroupPlaceholder())
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
        err = writer.WriteObjectValue("enableTechAdminApproveGroup", m.GetEnableTechAdminApproveGroup())
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
    {
        err = writer.WriteObjectValue("removeGroupApproveGroup", m.GetRemoveGroupApproveGroup())
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
// SetCreateGroupApproveGroup sets the createGroupApproveGroup property value. The createGroupApproveGroup property
func (m *OrganizationOrganizationalUnit) SetCreateGroupApproveGroup(value GroupGroupable)() {
    m.createGroupApproveGroup = value
}
// SetCreateGroupPlaceholder sets the createGroupPlaceholder property value. The createGroupPlaceholder property
func (m *OrganizationOrganizationalUnit) SetCreateGroupPlaceholder(value *string)() {
    m.createGroupPlaceholder = value
}
// SetDepth sets the depth property value. The depth property
func (m *OrganizationOrganizationalUnit) SetDepth(value *int32)() {
    m.depth = value
}
// SetDescription sets the description property value. The description property
func (m *OrganizationOrganizationalUnit) SetDescription(value *string)() {
    m.description = value
}
// SetEnableTechAdminApproveGroup sets the enableTechAdminApproveGroup property value. The enableTechAdminApproveGroup property
func (m *OrganizationOrganizationalUnit) SetEnableTechAdminApproveGroup(value GroupGroupable)() {
    m.enableTechAdminApproveGroup = value
}
// SetOwner sets the owner property value. The owner property
func (m *OrganizationOrganizationalUnit) SetOwner(value GroupGroupPrimerable)() {
    m.owner = value
}
// SetParent sets the parent property value. The parent property
func (m *OrganizationOrganizationalUnit) SetParent(value OrganizationOrganizationalUnitPrimerable)() {
    m.parent = value
}
// SetRemoveGroupApproveGroup sets the removeGroupApproveGroup property value. The removeGroupApproveGroup property
func (m *OrganizationOrganizationalUnit) SetRemoveGroupApproveGroup(value GroupGroupable)() {
    m.removeGroupApproveGroup = value
}
type OrganizationOrganizationalUnitable interface {
    OrganizationOrganizationalUnitPrimerable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdditionalObjects()(OrganizationOrganizationalUnit_additionalObjectsable)
    GetCreateGroupApproveGroup()(GroupGroupable)
    GetCreateGroupPlaceholder()(*string)
    GetDepth()(*int32)
    GetDescription()(*string)
    GetEnableTechAdminApproveGroup()(GroupGroupable)
    GetOwner()(GroupGroupPrimerable)
    GetParent()(OrganizationOrganizationalUnitPrimerable)
    GetRemoveGroupApproveGroup()(GroupGroupable)
    SetAdditionalObjects(value OrganizationOrganizationalUnit_additionalObjectsable)()
    SetCreateGroupApproveGroup(value GroupGroupable)()
    SetCreateGroupPlaceholder(value *string)()
    SetDepth(value *int32)()
    SetDescription(value *string)()
    SetEnableTechAdminApproveGroup(value GroupGroupable)()
    SetOwner(value GroupGroupPrimerable)()
    SetParent(value OrganizationOrganizationalUnitPrimerable)()
    SetRemoveGroupApproveGroup(value GroupGroupable)()
}
