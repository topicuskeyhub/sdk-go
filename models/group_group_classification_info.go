// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GroupGroupClassificationInfo struct {
    NonLinkable
    // The nrGroups property
    nrGroups *int32
}
// NewGroupGroupClassificationInfo instantiates a new GroupGroupClassificationInfo and sets the default values.
func NewGroupGroupClassificationInfo()(*GroupGroupClassificationInfo) {
    m := &GroupGroupClassificationInfo{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "group.GroupClassificationInfo"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateGroupGroupClassificationInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGroupGroupClassificationInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupGroupClassificationInfo(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GroupGroupClassificationInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
    res["nrGroups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNrGroups(val)
        }
        return nil
    }
    return res
}
// GetNrGroups gets the nrGroups property value. The nrGroups property
// returns a *int32 when successful
func (m *GroupGroupClassificationInfo) GetNrGroups()(*int32) {
    return m.nrGroups
}
// Serialize serializes information the current object
func (m *GroupGroupClassificationInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("nrGroups", m.GetNrGroups())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetNrGroups sets the nrGroups property value. The nrGroups property
func (m *GroupGroupClassificationInfo) SetNrGroups(value *int32)() {
    m.nrGroups = value
}
type GroupGroupClassificationInfoable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetNrGroups()(*int32)
    SetNrGroups(value *int32)()
}
