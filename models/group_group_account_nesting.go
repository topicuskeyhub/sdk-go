package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GroupGroupAccountNesting struct {
    NonLinkable
    // The connected property
    connected *bool
}
// NewGroupGroupAccountNesting instantiates a new GroupGroupAccountNesting and sets the default values.
func NewGroupGroupAccountNesting()(*GroupGroupAccountNesting) {
    m := &GroupGroupAccountNesting{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "group.GroupAccountNesting"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateGroupGroupAccountNestingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGroupGroupAccountNestingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupGroupAccountNesting(), nil
}
// GetConnected gets the connected property value. The connected property
// returns a *bool when successful
func (m *GroupGroupAccountNesting) GetConnected()(*bool) {
    return m.connected
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GroupGroupAccountNesting) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
    res["connected"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnected(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *GroupGroupAccountNesting) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("connected", m.GetConnected())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConnected sets the connected property value. The connected property
func (m *GroupGroupAccountNesting) SetConnected(value *bool)() {
    m.connected = value
}
type GroupGroupAccountNestingable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConnected()(*bool)
    SetConnected(value *bool)()
}
