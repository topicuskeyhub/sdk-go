package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ServiceaccountServiceAccountPrimerLinkableWrapperWithCount struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The count property
    count *int64
    // The items property
    items []ServiceaccountServiceAccountPrimerable
}
// NewServiceaccountServiceAccountPrimerLinkableWrapperWithCount instantiates a new ServiceaccountServiceAccountPrimerLinkableWrapperWithCount and sets the default values.
func NewServiceaccountServiceAccountPrimerLinkableWrapperWithCount()(*ServiceaccountServiceAccountPrimerLinkableWrapperWithCount) {
    m := &ServiceaccountServiceAccountPrimerLinkableWrapperWithCount{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateServiceaccountServiceAccountPrimerLinkableWrapperWithCountFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateServiceaccountServiceAccountPrimerLinkableWrapperWithCountFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServiceaccountServiceAccountPrimerLinkableWrapperWithCount(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ServiceaccountServiceAccountPrimerLinkableWrapperWithCount) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCount gets the count property value. The count property
// returns a *int64 when successful
func (m *ServiceaccountServiceAccountPrimerLinkableWrapperWithCount) GetCount()(*int64) {
    return m.count
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ServiceaccountServiceAccountPrimerLinkableWrapperWithCount) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
        val, err := n.GetCollectionOfObjectValues(CreateServiceaccountServiceAccountPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ServiceaccountServiceAccountPrimerable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ServiceaccountServiceAccountPrimerable)
                }
            }
            m.SetItems(res)
        }
        return nil
    }
    return res
}
// GetItems gets the items property value. The items property
// returns a []ServiceaccountServiceAccountPrimerable when successful
func (m *ServiceaccountServiceAccountPrimerLinkableWrapperWithCount) GetItems()([]ServiceaccountServiceAccountPrimerable) {
    return m.items
}
// Serialize serializes information the current object
func (m *ServiceaccountServiceAccountPrimerLinkableWrapperWithCount) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ServiceaccountServiceAccountPrimerLinkableWrapperWithCount) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCount sets the count property value. The count property
func (m *ServiceaccountServiceAccountPrimerLinkableWrapperWithCount) SetCount(value *int64)() {
    m.count = value
}
// SetItems sets the items property value. The items property
func (m *ServiceaccountServiceAccountPrimerLinkableWrapperWithCount) SetItems(value []ServiceaccountServiceAccountPrimerable)() {
    m.items = value
}
type ServiceaccountServiceAccountPrimerLinkableWrapperWithCountable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCount()(*int64)
    GetItems()([]ServiceaccountServiceAccountPrimerable)
    SetCount(value *int64)()
    SetItems(value []ServiceaccountServiceAccountPrimerable)()
}
