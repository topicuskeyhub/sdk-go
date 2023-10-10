package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupGroupSegmentedLinkableWrapper 
type GroupGroupSegmentedLinkableWrapper struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The items property
    items []GroupGroupable
    // The segments property
    segments []SegmentCountable
}
// NewGroupGroupSegmentedLinkableWrapper instantiates a new groupGroupSegmentedLinkableWrapper and sets the default values.
func NewGroupGroupSegmentedLinkableWrapper()(*GroupGroupSegmentedLinkableWrapper) {
    m := &GroupGroupSegmentedLinkableWrapper{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGroupGroupSegmentedLinkableWrapperFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupGroupSegmentedLinkableWrapperFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupGroupSegmentedLinkableWrapper(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GroupGroupSegmentedLinkableWrapper) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupGroupSegmentedLinkableWrapper) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["items"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGroupGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GroupGroupable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(GroupGroupable)
                }
            }
            m.SetItems(res)
        }
        return nil
    }
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
// GetItems gets the items property value. The items property
func (m *GroupGroupSegmentedLinkableWrapper) GetItems()([]GroupGroupable) {
    return m.items
}
// GetSegments gets the segments property value. The segments property
func (m *GroupGroupSegmentedLinkableWrapper) GetSegments()([]SegmentCountable) {
    return m.segments
}
// Serialize serializes information the current object
func (m *GroupGroupSegmentedLinkableWrapper) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetItems() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetItems()))
        for i, v := range m.GetItems() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("items", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSegments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSegments()))
        for i, v := range m.GetSegments() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("segments", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GroupGroupSegmentedLinkableWrapper) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetItems sets the items property value. The items property
func (m *GroupGroupSegmentedLinkableWrapper) SetItems(value []GroupGroupable)() {
    m.items = value
}
// SetSegments sets the segments property value. The segments property
func (m *GroupGroupSegmentedLinkableWrapper) SetSegments(value []SegmentCountable)() {
    m.segments = value
}
// GroupGroupSegmentedLinkableWrapperable 
type GroupGroupSegmentedLinkableWrapperable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetItems()([]GroupGroupable)
    GetSegments()([]SegmentCountable)
    SetItems(value []GroupGroupable)()
    SetSegments(value []SegmentCountable)()
}
