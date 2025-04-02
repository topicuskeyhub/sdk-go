package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GroupProvisioningGroupLinkableWrapperWithCount struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The count property
    count *int64
    // The items property
    items []GroupProvisioningGroupable
}
// NewGroupProvisioningGroupLinkableWrapperWithCount instantiates a new GroupProvisioningGroupLinkableWrapperWithCount and sets the default values.
func NewGroupProvisioningGroupLinkableWrapperWithCount()(*GroupProvisioningGroupLinkableWrapperWithCount) {
    m := &GroupProvisioningGroupLinkableWrapperWithCount{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGroupProvisioningGroupLinkableWrapperWithCountFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGroupProvisioningGroupLinkableWrapperWithCountFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupProvisioningGroupLinkableWrapperWithCount(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GroupProvisioningGroupLinkableWrapperWithCount) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCount gets the count property value. The count property
// returns a *int64 when successful
func (m *GroupProvisioningGroupLinkableWrapperWithCount) GetCount()(*int64) {
    return m.count
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GroupProvisioningGroupLinkableWrapperWithCount) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCount(val)
        }
        return nil
    }
    res["items"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGroupProvisioningGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GroupProvisioningGroupable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(GroupProvisioningGroupable)
                }
            }
            m.SetItems(res)
        }
        return nil
    }
    return res
}
// GetItems gets the items property value. The items property
// returns a []GroupProvisioningGroupable when successful
func (m *GroupProvisioningGroupLinkableWrapperWithCount) GetItems()([]GroupProvisioningGroupable) {
    return m.items
}
// Serialize serializes information the current object
func (m *GroupProvisioningGroupLinkableWrapperWithCount) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GroupProvisioningGroupLinkableWrapperWithCount) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCount sets the count property value. The count property
func (m *GroupProvisioningGroupLinkableWrapperWithCount) SetCount(value *int64)() {
    m.count = value
}
// SetItems sets the items property value. The items property
func (m *GroupProvisioningGroupLinkableWrapperWithCount) SetItems(value []GroupProvisioningGroupable)() {
    m.items = value
}
type GroupProvisioningGroupLinkableWrapperWithCountable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCount()(*int64)
    GetItems()([]GroupProvisioningGroupable)
    SetCount(value *int64)()
    SetItems(value []GroupProvisioningGroupable)()
}
