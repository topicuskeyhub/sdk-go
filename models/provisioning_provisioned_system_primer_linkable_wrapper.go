package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProvisioningProvisionedSystemPrimerLinkableWrapper 
type ProvisioningProvisionedSystemPrimerLinkableWrapper struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The items property
    items []ProvisioningProvisionedSystemPrimerable
}
// NewProvisioningProvisionedSystemPrimerLinkableWrapper instantiates a new provisioningProvisionedSystemPrimerLinkableWrapper and sets the default values.
func NewProvisioningProvisionedSystemPrimerLinkableWrapper()(*ProvisioningProvisionedSystemPrimerLinkableWrapper) {
    m := &ProvisioningProvisionedSystemPrimerLinkableWrapper{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateProvisioningProvisionedSystemPrimerLinkableWrapperFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProvisioningProvisionedSystemPrimerLinkableWrapperFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProvisioningProvisionedSystemPrimerLinkableWrapper(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ProvisioningProvisionedSystemPrimerLinkableWrapper) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProvisioningProvisionedSystemPrimerLinkableWrapper) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["items"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateProvisioningProvisionedSystemPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ProvisioningProvisionedSystemPrimerable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ProvisioningProvisionedSystemPrimerable)
                }
            }
            m.SetItems(res)
        }
        return nil
    }
    return res
}
// GetItems gets the items property value. The items property
func (m *ProvisioningProvisionedSystemPrimerLinkableWrapper) GetItems()([]ProvisioningProvisionedSystemPrimerable) {
    return m.items
}
// Serialize serializes information the current object
func (m *ProvisioningProvisionedSystemPrimerLinkableWrapper) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ProvisioningProvisionedSystemPrimerLinkableWrapper) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetItems sets the items property value. The items property
func (m *ProvisioningProvisionedSystemPrimerLinkableWrapper) SetItems(value []ProvisioningProvisionedSystemPrimerable)() {
    m.items = value
}
// ProvisioningProvisionedSystemPrimerLinkableWrapperable 
type ProvisioningProvisionedSystemPrimerLinkableWrapperable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetItems()([]ProvisioningProvisionedSystemPrimerable)
    SetItems(value []ProvisioningProvisionedSystemPrimerable)()
}
