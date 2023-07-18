package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupGroupSegmentedLinkableWrapper 
type GroupGroupSegmentedLinkableWrapper struct {
    LinkableWrapper
    // The segments property
    segments []SegmentCountable
}
// NewGroupGroupSegmentedLinkableWrapper instantiates a new groupGroupSegmentedLinkableWrapper and sets the default values.
func NewGroupGroupSegmentedLinkableWrapper()(*GroupGroupSegmentedLinkableWrapper) {
    m := &GroupGroupSegmentedLinkableWrapper{
        LinkableWrapper: *NewLinkableWrapper(),
    }
    return m
}
// CreateGroupGroupSegmentedLinkableWrapperFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupGroupSegmentedLinkableWrapperFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupGroupSegmentedLinkableWrapper(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupGroupSegmentedLinkableWrapper) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.LinkableWrapper.GetFieldDeserializers()
    res["segments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSegmentCountFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SegmentCountable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SegmentCountable)
                }
            }
            m.SetSegments(res)
        }
        return nil
    }
    return res
}
// GetSegments gets the segments property value. The segments property
func (m *GroupGroupSegmentedLinkableWrapper) GetSegments()([]SegmentCountable) {
    return m.segments
}
// Serialize serializes information the current object
func (m *GroupGroupSegmentedLinkableWrapper) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.LinkableWrapper.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetSegments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSegments()))
        for i, v := range m.GetSegments() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("segments", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSegments sets the segments property value. The segments property
func (m *GroupGroupSegmentedLinkableWrapper) SetSegments(value []SegmentCountable)() {
    m.segments = value
}
// GroupGroupSegmentedLinkableWrapperable 
type GroupGroupSegmentedLinkableWrapperable interface {
    LinkableWrapperable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSegments()([]SegmentCountable)
    SetSegments(value []SegmentCountable)()
}
