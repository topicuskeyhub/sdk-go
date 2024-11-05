package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GroupGroupAccessInfo struct {
    NonLinkable
    // The businessAccounts property
    businessAccounts *bool
}
// NewGroupGroupAccessInfo instantiates a new GroupGroupAccessInfo and sets the default values.
func NewGroupGroupAccessInfo()(*GroupGroupAccessInfo) {
    m := &GroupGroupAccessInfo{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "group.GroupAccessInfo"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateGroupGroupAccessInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGroupGroupAccessInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupGroupAccessInfo(), nil
}
// GetBusinessAccounts gets the businessAccounts property value. The businessAccounts property
// returns a *bool when successful
func (m *GroupGroupAccessInfo) GetBusinessAccounts()(*bool) {
    return m.businessAccounts
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GroupGroupAccessInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
    res["businessAccounts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBusinessAccounts(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *GroupGroupAccessInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// SetBusinessAccounts sets the businessAccounts property value. The businessAccounts property
func (m *GroupGroupAccessInfo) SetBusinessAccounts(value *bool)() {
    m.businessAccounts = value
}
type GroupGroupAccessInfoable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBusinessAccounts()(*bool)
    SetBusinessAccounts(value *bool)()
}