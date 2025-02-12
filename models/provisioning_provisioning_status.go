package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ProvisioningProvisioningStatus struct {
    NonLinkable
    // The allLoaded property
    allLoaded *bool
    // The groups property
    groups []ProvisioningGroupProvisioningStatusable
    // The ignoreErrors property
    ignoreErrors *bool
    // The restrictedByLicense property
    restrictedByLicense *bool
}
// NewProvisioningProvisioningStatus instantiates a new ProvisioningProvisioningStatus and sets the default values.
func NewProvisioningProvisioningStatus()(*ProvisioningProvisioningStatus) {
    m := &ProvisioningProvisioningStatus{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "provisioning.ProvisioningStatus"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateProvisioningProvisioningStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateProvisioningProvisioningStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProvisioningProvisioningStatus(), nil
}
// GetAllLoaded gets the allLoaded property value. The allLoaded property
// returns a *bool when successful
func (m *ProvisioningProvisioningStatus) GetAllLoaded()(*bool) {
    return m.allLoaded
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ProvisioningProvisioningStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
    res["allLoaded"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllLoaded(val)
        }
        return nil
    }
    res["groups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateProvisioningGroupProvisioningStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ProvisioningGroupProvisioningStatusable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ProvisioningGroupProvisioningStatusable)
                }
            }
            m.SetGroups(res)
        }
        return nil
    }
    res["ignoreErrors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIgnoreErrors(val)
        }
        return nil
    }
    res["restrictedByLicense"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestrictedByLicense(val)
        }
        return nil
    }
    return res
}
// GetGroups gets the groups property value. The groups property
// returns a []ProvisioningGroupProvisioningStatusable when successful
func (m *ProvisioningProvisioningStatus) GetGroups()([]ProvisioningGroupProvisioningStatusable) {
    return m.groups
}
// GetIgnoreErrors gets the ignoreErrors property value. The ignoreErrors property
// returns a *bool when successful
func (m *ProvisioningProvisioningStatus) GetIgnoreErrors()(*bool) {
    return m.ignoreErrors
}
// GetRestrictedByLicense gets the restrictedByLicense property value. The restrictedByLicense property
// returns a *bool when successful
func (m *ProvisioningProvisioningStatus) GetRestrictedByLicense()(*bool) {
    return m.restrictedByLicense
}
// Serialize serializes information the current object
func (m *ProvisioningProvisioningStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("allLoaded", m.GetAllLoaded())
        if err != nil {
            return err
        }
    }
    if m.GetGroups() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetGroups()))
        for i, v := range m.GetGroups() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("groups", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("ignoreErrors", m.GetIgnoreErrors())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("restrictedByLicense", m.GetRestrictedByLicense())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllLoaded sets the allLoaded property value. The allLoaded property
func (m *ProvisioningProvisioningStatus) SetAllLoaded(value *bool)() {
    m.allLoaded = value
}
// SetGroups sets the groups property value. The groups property
func (m *ProvisioningProvisioningStatus) SetGroups(value []ProvisioningGroupProvisioningStatusable)() {
    m.groups = value
}
// SetIgnoreErrors sets the ignoreErrors property value. The ignoreErrors property
func (m *ProvisioningProvisioningStatus) SetIgnoreErrors(value *bool)() {
    m.ignoreErrors = value
}
// SetRestrictedByLicense sets the restrictedByLicense property value. The restrictedByLicense property
func (m *ProvisioningProvisioningStatus) SetRestrictedByLicense(value *bool)() {
    m.restrictedByLicense = value
}
type ProvisioningProvisioningStatusable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllLoaded()(*bool)
    GetGroups()([]ProvisioningGroupProvisioningStatusable)
    GetIgnoreErrors()(*bool)
    GetRestrictedByLicense()(*bool)
    SetAllLoaded(value *bool)()
    SetGroups(value []ProvisioningGroupProvisioningStatusable)()
    SetIgnoreErrors(value *bool)()
    SetRestrictedByLicense(value *bool)()
}
