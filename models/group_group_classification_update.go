package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupGroupClassificationUpdate 
type GroupGroupClassificationUpdate struct {
    NonLinkable
    // The classificaton property
    classificaton GroupGroupClassificationPrimerable
}
// NewGroupGroupClassificationUpdate instantiates a new groupGroupClassificationUpdate and sets the default values.
func NewGroupGroupClassificationUpdate()(*GroupGroupClassificationUpdate) {
    m := &GroupGroupClassificationUpdate{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "group.GroupClassificationUpdate"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateGroupGroupClassificationUpdateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupGroupClassificationUpdateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupGroupClassificationUpdate(), nil
}
// GetClassificaton gets the classificaton property value. The classificaton property
func (m *GroupGroupClassificationUpdate) GetClassificaton()(GroupGroupClassificationPrimerable) {
    return m.classificaton
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupGroupClassificationUpdate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
    res["classificaton"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupGroupClassificationPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClassificaton(val.(GroupGroupClassificationPrimerable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *GroupGroupClassificationUpdate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("classificaton", m.GetClassificaton())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetClassificaton sets the classificaton property value. The classificaton property
func (m *GroupGroupClassificationUpdate) SetClassificaton(value GroupGroupClassificationPrimerable)() {
    m.classificaton = value
}
// GroupGroupClassificationUpdateable 
type GroupGroupClassificationUpdateable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClassificaton()(GroupGroupClassificationPrimerable)
    SetClassificaton(value GroupGroupClassificationPrimerable)()
}
