package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServiceaccountServiceAccountPrimerLinkableWrapper 
type ServiceaccountServiceAccountPrimerLinkableWrapper struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The items property
    items []ServiceaccountServiceAccountPrimerable
}
// NewServiceaccountServiceAccountPrimerLinkableWrapper instantiates a new serviceaccountServiceAccountPrimerLinkableWrapper and sets the default values.
func NewServiceaccountServiceAccountPrimerLinkableWrapper()(*ServiceaccountServiceAccountPrimerLinkableWrapper) {
    m := &ServiceaccountServiceAccountPrimerLinkableWrapper{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateServiceaccountServiceAccountPrimerLinkableWrapperFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateServiceaccountServiceAccountPrimerLinkableWrapperFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServiceaccountServiceAccountPrimerLinkableWrapper(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ServiceaccountServiceAccountPrimerLinkableWrapper) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ServiceaccountServiceAccountPrimerLinkableWrapper) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
func (m *ServiceaccountServiceAccountPrimerLinkableWrapper) GetItems()([]ServiceaccountServiceAccountPrimerable) {
    return m.items
}
// Serialize serializes information the current object
func (m *ServiceaccountServiceAccountPrimerLinkableWrapper) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ServiceaccountServiceAccountPrimerLinkableWrapper) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetItems sets the items property value. The items property
func (m *ServiceaccountServiceAccountPrimerLinkableWrapper) SetItems(value []ServiceaccountServiceAccountPrimerable)() {
    m.items = value
}
// ServiceaccountServiceAccountPrimerLinkableWrapperable 
type ServiceaccountServiceAccountPrimerLinkableWrapperable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetItems()([]ServiceaccountServiceAccountPrimerable)
    SetItems(value []ServiceaccountServiceAccountPrimerable)()
}
