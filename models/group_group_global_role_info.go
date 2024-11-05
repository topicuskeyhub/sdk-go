package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GroupGroupGlobalRoleInfo struct {
    NonLinkable
    // The createGroupApproveGroupFor property
    createGroupApproveGroupFor []OrganizationOrganizationalUnitPrimerable
    // The enableTechAdminApproveGroupFor property
    enableTechAdminApproveGroupFor []OrganizationOrganizationalUnitPrimerable
    // The recoveryFallbackGroupFor property
    recoveryFallbackGroupFor []OrganizationOrganizationalUnitPrimerable
    // The removeGroupApproveGroupFor property
    removeGroupApproveGroupFor []OrganizationOrganizationalUnitPrimerable
}
// NewGroupGroupGlobalRoleInfo instantiates a new GroupGroupGlobalRoleInfo and sets the default values.
func NewGroupGroupGlobalRoleInfo()(*GroupGroupGlobalRoleInfo) {
    m := &GroupGroupGlobalRoleInfo{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "group.GroupGlobalRoleInfo"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateGroupGroupGlobalRoleInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGroupGroupGlobalRoleInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupGroupGlobalRoleInfo(), nil
}
// GetCreateGroupApproveGroupFor gets the createGroupApproveGroupFor property value. The createGroupApproveGroupFor property
// returns a []OrganizationOrganizationalUnitPrimerable when successful
func (m *GroupGroupGlobalRoleInfo) GetCreateGroupApproveGroupFor()([]OrganizationOrganizationalUnitPrimerable) {
    return m.createGroupApproveGroupFor
}
// GetEnableTechAdminApproveGroupFor gets the enableTechAdminApproveGroupFor property value. The enableTechAdminApproveGroupFor property
// returns a []OrganizationOrganizationalUnitPrimerable when successful
func (m *GroupGroupGlobalRoleInfo) GetEnableTechAdminApproveGroupFor()([]OrganizationOrganizationalUnitPrimerable) {
    return m.enableTechAdminApproveGroupFor
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GroupGroupGlobalRoleInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
    res["createGroupApproveGroupFor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOrganizationOrganizationalUnitPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OrganizationOrganizationalUnitPrimerable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(OrganizationOrganizationalUnitPrimerable)
                }
            }
            m.SetCreateGroupApproveGroupFor(res)
        }
        return nil
    }
    res["enableTechAdminApproveGroupFor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOrganizationOrganizationalUnitPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OrganizationOrganizationalUnitPrimerable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(OrganizationOrganizationalUnitPrimerable)
                }
            }
            m.SetEnableTechAdminApproveGroupFor(res)
        }
        return nil
    }
    res["recoveryFallbackGroupFor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOrganizationOrganizationalUnitPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OrganizationOrganizationalUnitPrimerable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(OrganizationOrganizationalUnitPrimerable)
                }
            }
            m.SetRecoveryFallbackGroupFor(res)
        }
        return nil
    }
    res["removeGroupApproveGroupFor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOrganizationOrganizationalUnitPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OrganizationOrganizationalUnitPrimerable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(OrganizationOrganizationalUnitPrimerable)
                }
            }
            m.SetRemoveGroupApproveGroupFor(res)
        }
        return nil
    }
    return res
}
// GetRecoveryFallbackGroupFor gets the recoveryFallbackGroupFor property value. The recoveryFallbackGroupFor property
// returns a []OrganizationOrganizationalUnitPrimerable when successful
func (m *GroupGroupGlobalRoleInfo) GetRecoveryFallbackGroupFor()([]OrganizationOrganizationalUnitPrimerable) {
    return m.recoveryFallbackGroupFor
}
// GetRemoveGroupApproveGroupFor gets the removeGroupApproveGroupFor property value. The removeGroupApproveGroupFor property
// returns a []OrganizationOrganizationalUnitPrimerable when successful
func (m *GroupGroupGlobalRoleInfo) GetRemoveGroupApproveGroupFor()([]OrganizationOrganizationalUnitPrimerable) {
    return m.removeGroupApproveGroupFor
}
// Serialize serializes information the current object
func (m *GroupGroupGlobalRoleInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCreateGroupApproveGroupFor() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCreateGroupApproveGroupFor()))
        for i, v := range m.GetCreateGroupApproveGroupFor() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("createGroupApproveGroupFor", cast)
        if err != nil {
            return err
        }
    }
    if m.GetEnableTechAdminApproveGroupFor() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEnableTechAdminApproveGroupFor()))
        for i, v := range m.GetEnableTechAdminApproveGroupFor() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("enableTechAdminApproveGroupFor", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRecoveryFallbackGroupFor() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRecoveryFallbackGroupFor()))
        for i, v := range m.GetRecoveryFallbackGroupFor() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("recoveryFallbackGroupFor", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRemoveGroupApproveGroupFor() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRemoveGroupApproveGroupFor()))
        for i, v := range m.GetRemoveGroupApproveGroupFor() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("removeGroupApproveGroupFor", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreateGroupApproveGroupFor sets the createGroupApproveGroupFor property value. The createGroupApproveGroupFor property
func (m *GroupGroupGlobalRoleInfo) SetCreateGroupApproveGroupFor(value []OrganizationOrganizationalUnitPrimerable)() {
    m.createGroupApproveGroupFor = value
}
// SetEnableTechAdminApproveGroupFor sets the enableTechAdminApproveGroupFor property value. The enableTechAdminApproveGroupFor property
func (m *GroupGroupGlobalRoleInfo) SetEnableTechAdminApproveGroupFor(value []OrganizationOrganizationalUnitPrimerable)() {
    m.enableTechAdminApproveGroupFor = value
}
// SetRecoveryFallbackGroupFor sets the recoveryFallbackGroupFor property value. The recoveryFallbackGroupFor property
func (m *GroupGroupGlobalRoleInfo) SetRecoveryFallbackGroupFor(value []OrganizationOrganizationalUnitPrimerable)() {
    m.recoveryFallbackGroupFor = value
}
// SetRemoveGroupApproveGroupFor sets the removeGroupApproveGroupFor property value. The removeGroupApproveGroupFor property
func (m *GroupGroupGlobalRoleInfo) SetRemoveGroupApproveGroupFor(value []OrganizationOrganizationalUnitPrimerable)() {
    m.removeGroupApproveGroupFor = value
}
type GroupGroupGlobalRoleInfoable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreateGroupApproveGroupFor()([]OrganizationOrganizationalUnitPrimerable)
    GetEnableTechAdminApproveGroupFor()([]OrganizationOrganizationalUnitPrimerable)
    GetRecoveryFallbackGroupFor()([]OrganizationOrganizationalUnitPrimerable)
    GetRemoveGroupApproveGroupFor()([]OrganizationOrganizationalUnitPrimerable)
    SetCreateGroupApproveGroupFor(value []OrganizationOrganizationalUnitPrimerable)()
    SetEnableTechAdminApproveGroupFor(value []OrganizationOrganizationalUnitPrimerable)()
    SetRecoveryFallbackGroupFor(value []OrganizationOrganizationalUnitPrimerable)()
    SetRemoveGroupApproveGroupFor(value []OrganizationOrganizationalUnitPrimerable)()
}
