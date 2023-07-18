package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProvisioningAccountProvisioningStatus 
type ProvisioningAccountProvisioningStatus struct {
    NonLinkable
    // The description property
    description *string
    // The result property
    result *ProvisioningAccountProvisioningResult
    // The system property
    system GroupProvisioningGroupable
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
    res["system"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupProvisioningGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystem(val.(GroupProvisioningGroupable))
        }
        return nil
    }
    return res
}
// GetResult gets the result property value. The result property
func (m *ProvisioningAccountProvisioningStatus) GetResult()(*ProvisioningAccountProvisioningResult) {
    return m.result
}
// GetSystem gets the system property value. The system property
func (m *ProvisioningAccountProvisioningStatus) GetSystem()(GroupProvisioningGroupable) {
    return m.system
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
    if m.GetResult() != nil {
        cast := (*m.GetResult()).String()
        err = writer.WriteStringValue("result", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("system", m.GetSystem())
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
// SetResult sets the result property value. The result property
func (m *ProvisioningAccountProvisioningStatus) SetResult(value *ProvisioningAccountProvisioningResult)() {
    m.result = value
}
// SetSystem sets the system property value. The system property
func (m *ProvisioningAccountProvisioningStatus) SetSystem(value GroupProvisioningGroupable)() {
    m.system = value
}
// ProvisioningAccountProvisioningStatusable 
type ProvisioningAccountProvisioningStatusable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetResult()(*ProvisioningAccountProvisioningResult)
    GetSystem()(GroupProvisioningGroupable)
    SetDescription(value *string)()
    SetResult(value *ProvisioningAccountProvisioningResult)()
    SetSystem(value GroupProvisioningGroupable)()
}
