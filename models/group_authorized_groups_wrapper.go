package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupAuthorizedGroupsWrapper 
type GroupAuthorizedGroupsWrapper struct {
    GroupGroupLinkableWrapper
    // The groupCount property
    groupCount *int64
}
// NewGroupAuthorizedGroupsWrapper instantiates a new groupAuthorizedGroupsWrapper and sets the default values.
func NewGroupAuthorizedGroupsWrapper()(*GroupAuthorizedGroupsWrapper) {
    m := &GroupAuthorizedGroupsWrapper{
        GroupGroupLinkableWrapper: *NewGroupGroupLinkableWrapper(),
    }
    return m
}
// CreateGroupAuthorizedGroupsWrapperFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupAuthorizedGroupsWrapperFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupAuthorizedGroupsWrapper(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupAuthorizedGroupsWrapper) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.GroupGroupLinkableWrapper.GetFieldDeserializers()
    res["groupCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupCount(val)
        }
        return nil
    }
    return res
}
// GetGroupCount gets the groupCount property value. The groupCount property
func (m *GroupAuthorizedGroupsWrapper) GetGroupCount()(*int64) {
    return m.groupCount
}
// Serialize serializes information the current object
func (m *GroupAuthorizedGroupsWrapper) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.GroupGroupLinkableWrapper.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("groupCount", m.GetGroupCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetGroupCount sets the groupCount property value. The groupCount property
func (m *GroupAuthorizedGroupsWrapper) SetGroupCount(value *int64)() {
    m.groupCount = value
}
// GroupAuthorizedGroupsWrapperable 
type GroupAuthorizedGroupsWrapperable interface {
    GroupGroupLinkableWrapperable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetGroupCount()(*int64)
    SetGroupCount(value *int64)()
}
