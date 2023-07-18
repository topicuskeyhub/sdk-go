package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OrganizationOrganizationalUnitLinkableWrapper 
type OrganizationOrganizationalUnitLinkableWrapper struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The items property
    items []OrganizationOrganizationalUnitable
}
// NewOrganizationOrganizationalUnitLinkableWrapper instantiates a new organizationOrganizationalUnitLinkableWrapper and sets the default values.
func NewOrganizationOrganizationalUnitLinkableWrapper()(*OrganizationOrganizationalUnitLinkableWrapper) {
    m := &OrganizationOrganizationalUnitLinkableWrapper{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateOrganizationOrganizationalUnitLinkableWrapperFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOrganizationOrganizationalUnitLinkableWrapperFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOrganizationOrganizationalUnitLinkableWrapper(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OrganizationOrganizationalUnitLinkableWrapper) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OrganizationOrganizationalUnitLinkableWrapper) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["items"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOrganizationOrganizationalUnitFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OrganizationOrganizationalUnitable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(OrganizationOrganizationalUnitable)
                }
            }
            m.SetItems(res)
        }
        return nil
    }
    return res
}
// GetItems gets the items property value. The items property
func (m *OrganizationOrganizationalUnitLinkableWrapper) GetItems()([]OrganizationOrganizationalUnitable) {
    return m.items
}
// Serialize serializes information the current object
func (m *OrganizationOrganizationalUnitLinkableWrapper) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *OrganizationOrganizationalUnitLinkableWrapper) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetItems sets the items property value. The items property
func (m *OrganizationOrganizationalUnitLinkableWrapper) SetItems(value []OrganizationOrganizationalUnitable)() {
    m.items = value
}
// OrganizationOrganizationalUnitLinkableWrapperable 
type OrganizationOrganizationalUnitLinkableWrapperable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetItems()([]OrganizationOrganizationalUnitable)
    SetItems(value []OrganizationOrganizationalUnitable)()
}