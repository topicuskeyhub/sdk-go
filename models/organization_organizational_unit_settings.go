package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OrganizationOrganizationalUnitSettings struct {
    NonLinkable
    // The createGroupApproveGroup property
    createGroupApproveGroup GroupGroupable
    // The createGroupPlaceholder property
    createGroupPlaceholder *string
    // The enableTechAdminApproveGroup property
    enableTechAdminApproveGroup GroupGroupable
    // The removeGroupApproveGroup property
    removeGroupApproveGroup GroupGroupable
}
// NewOrganizationOrganizationalUnitSettings instantiates a new OrganizationOrganizationalUnitSettings and sets the default values.
func NewOrganizationOrganizationalUnitSettings()(*OrganizationOrganizationalUnitSettings) {
    m := &OrganizationOrganizationalUnitSettings{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "organization.OrganizationalUnitSettings"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateOrganizationOrganizationalUnitSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOrganizationOrganizationalUnitSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOrganizationOrganizationalUnitSettings(), nil
}
// GetCreateGroupApproveGroup gets the createGroupApproveGroup property value. The createGroupApproveGroup property
// returns a GroupGroupable when successful
func (m *OrganizationOrganizationalUnitSettings) GetCreateGroupApproveGroup()(GroupGroupable) {
    return m.createGroupApproveGroup
}
// GetCreateGroupPlaceholder gets the createGroupPlaceholder property value. The createGroupPlaceholder property
// returns a *string when successful
func (m *OrganizationOrganizationalUnitSettings) GetCreateGroupPlaceholder()(*string) {
    return m.createGroupPlaceholder
}
// GetEnableTechAdminApproveGroup gets the enableTechAdminApproveGroup property value. The enableTechAdminApproveGroup property
// returns a GroupGroupable when successful
func (m *OrganizationOrganizationalUnitSettings) GetEnableTechAdminApproveGroup()(GroupGroupable) {
    return m.enableTechAdminApproveGroup
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OrganizationOrganizationalUnitSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
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
// GetRemoveGroupApproveGroup gets the removeGroupApproveGroup property value. The removeGroupApproveGroup property
// returns a GroupGroupable when successful
func (m *OrganizationOrganizationalUnitSettings) GetRemoveGroupApproveGroup()(GroupGroupable) {
    return m.removeGroupApproveGroup
}
// Serialize serializes information the current object
func (m *OrganizationOrganizationalUnitSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
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
        err = writer.WriteObjectValue("enableTechAdminApproveGroup", m.GetEnableTechAdminApproveGroup())
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
// SetCreateGroupApproveGroup sets the createGroupApproveGroup property value. The createGroupApproveGroup property
func (m *OrganizationOrganizationalUnitSettings) SetCreateGroupApproveGroup(value GroupGroupable)() {
    m.createGroupApproveGroup = value
}
// SetCreateGroupPlaceholder sets the createGroupPlaceholder property value. The createGroupPlaceholder property
func (m *OrganizationOrganizationalUnitSettings) SetCreateGroupPlaceholder(value *string)() {
    m.createGroupPlaceholder = value
}
// SetEnableTechAdminApproveGroup sets the enableTechAdminApproveGroup property value. The enableTechAdminApproveGroup property
func (m *OrganizationOrganizationalUnitSettings) SetEnableTechAdminApproveGroup(value GroupGroupable)() {
    m.enableTechAdminApproveGroup = value
}
// SetRemoveGroupApproveGroup sets the removeGroupApproveGroup property value. The removeGroupApproveGroup property
func (m *OrganizationOrganizationalUnitSettings) SetRemoveGroupApproveGroup(value GroupGroupable)() {
    m.removeGroupApproveGroup = value
}
type OrganizationOrganizationalUnitSettingsable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreateGroupApproveGroup()(GroupGroupable)
    GetCreateGroupPlaceholder()(*string)
    GetEnableTechAdminApproveGroup()(GroupGroupable)
    GetRemoveGroupApproveGroup()(GroupGroupable)
    SetCreateGroupApproveGroup(value GroupGroupable)()
    SetCreateGroupPlaceholder(value *string)()
    SetEnableTechAdminApproveGroup(value GroupGroupable)()
    SetRemoveGroupApproveGroup(value GroupGroupable)()
}
