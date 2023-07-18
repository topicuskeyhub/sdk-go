package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProvisioningAbstractProvisionedLDAP_attributes 
type ProvisioningAbstractProvisionedLDAP_attributes struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
}
// NewProvisioningAbstractProvisionedLDAP_attributes instantiates a new provisioningAbstractProvisionedLDAP_attributes and sets the default values.
func NewProvisioningAbstractProvisionedLDAP_attributes()(*ProvisioningAbstractProvisionedLDAP_attributes) {
    m := &ProvisioningAbstractProvisionedLDAP_attributes{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateProvisioningAbstractProvisionedLDAP_attributesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProvisioningAbstractProvisionedLDAP_attributesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProvisioningAbstractProvisionedLDAP_attributes(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ProvisioningAbstractProvisionedLDAP_attributes) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProvisioningAbstractProvisionedLDAP_attributes) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    return res
}
// Serialize serializes information the current object
func (m *ProvisioningAbstractProvisionedLDAP_attributes) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ProvisioningAbstractProvisionedLDAP_attributes) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// ProvisioningAbstractProvisionedLDAP_attributesable 
type ProvisioningAbstractProvisionedLDAP_attributesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
