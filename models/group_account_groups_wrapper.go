package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GroupAccountGroupsWrapper struct {
    GroupAccountGroupLinkableWrapper
    // The totalAccountGroupCount property
    totalAccountGroupCount *int64
}
// NewGroupAccountGroupsWrapper instantiates a new GroupAccountGroupsWrapper and sets the default values.
func NewGroupAccountGroupsWrapper()(*GroupAccountGroupsWrapper) {
    m := &GroupAccountGroupsWrapper{
        GroupAccountGroupLinkableWrapper: *NewGroupAccountGroupLinkableWrapper(),
    }
    return m
}
// CreateGroupAccountGroupsWrapperFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGroupAccountGroupsWrapperFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupAccountGroupsWrapper(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GroupAccountGroupsWrapper) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.GroupAccountGroupLinkableWrapper.GetFieldDeserializers()
    res["totalAccountGroupCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalAccountGroupCount(val)
        }
        return nil
    }
    return res
}
// GetTotalAccountGroupCount gets the totalAccountGroupCount property value. The totalAccountGroupCount property
// returns a *int64 when successful
func (m *GroupAccountGroupsWrapper) GetTotalAccountGroupCount()(*int64) {
    return m.totalAccountGroupCount
}
// Serialize serializes information the current object
func (m *GroupAccountGroupsWrapper) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.GroupAccountGroupLinkableWrapper.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// SetTotalAccountGroupCount sets the totalAccountGroupCount property value. The totalAccountGroupCount property
func (m *GroupAccountGroupsWrapper) SetTotalAccountGroupCount(value *int64)() {
    m.totalAccountGroupCount = value
}
type GroupAccountGroupsWrapperable interface {
    GroupAccountGroupLinkableWrapperable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetTotalAccountGroupCount()(*int64)
    SetTotalAccountGroupCount(value *int64)()
}
