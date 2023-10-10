package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OrganizationOrganizationalUnitPrimerLinkableWrapper 
type OrganizationOrganizationalUnitPrimerLinkableWrapper struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The items property
    items []OrganizationOrganizationalUnitPrimerable
}
// NewOrganizationOrganizationalUnitPrimerLinkableWrapper instantiates a new organizationOrganizationalUnitPrimerLinkableWrapper and sets the default values.
func NewOrganizationOrganizationalUnitPrimerLinkableWrapper()(*OrganizationOrganizationalUnitPrimerLinkableWrapper) {
    m := &OrganizationOrganizationalUnitPrimerLinkableWrapper{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateOrganizationOrganizationalUnitPrimerLinkableWrapperFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOrganizationOrganizationalUnitPrimerLinkableWrapperFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOrganizationOrganizationalUnitPrimerLinkableWrapper(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OrganizationOrganizationalUnitPrimerLinkableWrapper) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OrganizationOrganizationalUnitPrimerLinkableWrapper) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["items"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOrganizationOrganizationalUnitPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OrganizationOrganizationalUnitPrimerable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(OrganizationOrganizationalUnitPrimerable)
                }
            }
            m.SetItems(res)
        }
        return nil
    }
    return res
}
// GetItems gets the items property value. The items property
func (m *OrganizationOrganizationalUnitPrimerLinkableWrapper) GetItems()([]OrganizationOrganizationalUnitPrimerable) {
    return m.items
}
// Serialize serializes information the current object
func (m *OrganizationOrganizationalUnitPrimerLinkableWrapper) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *OrganizationOrganizationalUnitPrimerLinkableWrapper) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetItems sets the items property value. The items property
func (m *OrganizationOrganizationalUnitPrimerLinkableWrapper) SetItems(value []OrganizationOrganizationalUnitPrimerable)() {
    m.items = value
}
// OrganizationOrganizationalUnitPrimerLinkableWrapperable 
type OrganizationOrganizationalUnitPrimerLinkableWrapperable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetItems()([]OrganizationOrganizationalUnitPrimerable)
    SetItems(value []OrganizationOrganizationalUnitPrimerable)()
}
