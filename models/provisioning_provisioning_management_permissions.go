package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProvisioningProvisioningManagementPermissions 
type ProvisioningProvisioningManagementPermissions struct {
    NonLinkable
    // The createNewGroupsAllowed property
    createNewGroupsAllowed *bool
    // The createServiceAccountsAllowed property
    createServiceAccountsAllowed *bool
    // The reuseExistingGroupsAllowed property
    reuseExistingGroupsAllowed *bool
}
// NewProvisioningProvisioningManagementPermissions instantiates a new provisioningProvisioningManagementPermissions and sets the default values.
func NewProvisioningProvisioningManagementPermissions()(*ProvisioningProvisioningManagementPermissions) {
    m := &ProvisioningProvisioningManagementPermissions{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "provisioning.ProvisioningManagementPermissions"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateProvisioningProvisioningManagementPermissionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProvisioningProvisioningManagementPermissionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProvisioningProvisioningManagementPermissions(), nil
}
// GetCreateNewGroupsAllowed gets the createNewGroupsAllowed property value. The createNewGroupsAllowed property
func (m *ProvisioningProvisioningManagementPermissions) GetCreateNewGroupsAllowed()(*bool) {
    return m.createNewGroupsAllowed
}
// GetCreateServiceAccountsAllowed gets the createServiceAccountsAllowed property value. The createServiceAccountsAllowed property
func (m *ProvisioningProvisioningManagementPermissions) GetCreateServiceAccountsAllowed()(*bool) {
    return m.createServiceAccountsAllowed
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProvisioningProvisioningManagementPermissions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
    res["createNewGroupsAllowed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreateNewGroupsAllowed(val)
        }
        return nil
    }
    res["createServiceAccountsAllowed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreateServiceAccountsAllowed(val)
        }
        return nil
    }
    res["reuseExistingGroupsAllowed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReuseExistingGroupsAllowed(val)
        }
        return nil
    }
    return res
}
// GetReuseExistingGroupsAllowed gets the reuseExistingGroupsAllowed property value. The reuseExistingGroupsAllowed property
func (m *ProvisioningProvisioningManagementPermissions) GetReuseExistingGroupsAllowed()(*bool) {
    return m.reuseExistingGroupsAllowed
}
// Serialize serializes information the current object
func (m *ProvisioningProvisioningManagementPermissions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// SetCreateNewGroupsAllowed sets the createNewGroupsAllowed property value. The createNewGroupsAllowed property
func (m *ProvisioningProvisioningManagementPermissions) SetCreateNewGroupsAllowed(value *bool)() {
    m.createNewGroupsAllowed = value
}
// SetCreateServiceAccountsAllowed sets the createServiceAccountsAllowed property value. The createServiceAccountsAllowed property
func (m *ProvisioningProvisioningManagementPermissions) SetCreateServiceAccountsAllowed(value *bool)() {
    m.createServiceAccountsAllowed = value
}
// SetReuseExistingGroupsAllowed sets the reuseExistingGroupsAllowed property value. The reuseExistingGroupsAllowed property
func (m *ProvisioningProvisioningManagementPermissions) SetReuseExistingGroupsAllowed(value *bool)() {
    m.reuseExistingGroupsAllowed = value
}
// ProvisioningProvisioningManagementPermissionsable 
type ProvisioningProvisioningManagementPermissionsable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreateNewGroupsAllowed()(*bool)
    GetCreateServiceAccountsAllowed()(*bool)
    GetReuseExistingGroupsAllowed()(*bool)
    SetCreateNewGroupsAllowed(value *bool)()
    SetCreateServiceAccountsAllowed(value *bool)()
    SetReuseExistingGroupsAllowed(value *bool)()
}
