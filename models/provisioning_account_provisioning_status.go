package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProvisioningAccountProvisioningStatus 
type ProvisioningAccountProvisioningStatus struct {
    NonLinkable
    // The description property
    description *string
    // The provisionedSystem property
    provisionedSystem ProvisioningProvisionedSystemPrimerable
    // The provisioningGroup property
    provisioningGroup GroupProvisioningGroupable
    // The result property
    result *ProvisioningAccountProvisioningResult
}
// NewProvisioningAccountProvisioningStatus instantiates a new provisioningAccountProvisioningStatus and sets the default values.
func NewProvisioningAccountProvisioningStatus()(*ProvisioningAccountProvisioningStatus) {
    m := &ProvisioningAccountProvisioningStatus{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "provisioning.AccountProvisioningStatus"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateProvisioningAccountProvisioningStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProvisioningAccountProvisioningStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProvisioningAccountProvisioningStatus(), nil
}
// GetDescription gets the description property value. The description property
func (m *ProvisioningAccountProvisioningStatus) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProvisioningAccountProvisioningStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
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
    res["provisionedSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProvisioningProvisionedSystemPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProvisionedSystem(val.(ProvisioningProvisionedSystemPrimerable))
        }
        return nil
    }
    res["provisioningGroup"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupProvisioningGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProvisioningGroup(val.(GroupProvisioningGroupable))
        }
        return nil
    }
    res["result"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseProvisioningAccountProvisioningResult)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResult(val.(*ProvisioningAccountProvisioningResult))
        }
        return nil
    }
    return res
}
// GetProvisionedSystem gets the provisionedSystem property value. The provisionedSystem property
func (m *ProvisioningAccountProvisioningStatus) GetProvisionedSystem()(ProvisioningProvisionedSystemPrimerable) {
    return m.provisionedSystem
}
// GetProvisioningGroup gets the provisioningGroup property value. The provisioningGroup property
func (m *ProvisioningAccountProvisioningStatus) GetProvisioningGroup()(GroupProvisioningGroupable) {
    return m.provisioningGroup
}
// GetResult gets the result property value. The result property
func (m *ProvisioningAccountProvisioningStatus) GetResult()(*ProvisioningAccountProvisioningResult) {
    return m.result
}
// Serialize serializes information the current object
func (m *ProvisioningAccountProvisioningStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("provisionedSystem", m.GetProvisionedSystem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("provisioningGroup", m.GetProvisioningGroup())
        if err != nil {
            return err
        }
    }
    if m.GetResult() != nil {
        cast := (*m.GetResult()).String()
        err = writer.WriteStringValue("result", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDescription sets the description property value. The description property
func (m *ProvisioningAccountProvisioningStatus) SetDescription(value *string)() {
    m.description = value
}
// SetProvisionedSystem sets the provisionedSystem property value. The provisionedSystem property
func (m *ProvisioningAccountProvisioningStatus) SetProvisionedSystem(value ProvisioningProvisionedSystemPrimerable)() {
    m.provisionedSystem = value
}
// SetProvisioningGroup sets the provisioningGroup property value. The provisioningGroup property
func (m *ProvisioningAccountProvisioningStatus) SetProvisioningGroup(value GroupProvisioningGroupable)() {
    m.provisioningGroup = value
}
// SetResult sets the result property value. The result property
func (m *ProvisioningAccountProvisioningStatus) SetResult(value *ProvisioningAccountProvisioningResult)() {
    m.result = value
}
// ProvisioningAccountProvisioningStatusable 
type ProvisioningAccountProvisioningStatusable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetProvisionedSystem()(ProvisioningProvisionedSystemPrimerable)
    GetProvisioningGroup()(GroupProvisioningGroupable)
    GetResult()(*ProvisioningAccountProvisioningResult)
    SetDescription(value *string)()
    SetProvisionedSystem(value ProvisioningProvisionedSystemPrimerable)()
    SetProvisioningGroup(value GroupProvisioningGroupable)()
    SetResult(value *ProvisioningAccountProvisioningResult)()
}
