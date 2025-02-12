package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OrganizationOrganizationalUnitClientApplicationLinkableWrapper struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The items property
    items []OrganizationOrganizationalUnitClientApplicationable
}
// NewOrganizationOrganizationalUnitClientApplicationLinkableWrapper instantiates a new OrganizationOrganizationalUnitClientApplicationLinkableWrapper and sets the default values.
func NewOrganizationOrganizationalUnitClientApplicationLinkableWrapper()(*OrganizationOrganizationalUnitClientApplicationLinkableWrapper) {
    m := &OrganizationOrganizationalUnitClientApplicationLinkableWrapper{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateOrganizationOrganizationalUnitClientApplicationLinkableWrapperFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOrganizationOrganizationalUnitClientApplicationLinkableWrapperFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOrganizationOrganizationalUnitClientApplicationLinkableWrapper(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *OrganizationOrganizationalUnitClientApplicationLinkableWrapper) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OrganizationOrganizationalUnitClientApplicationLinkableWrapper) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["items"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOrganizationOrganizationalUnitClientApplicationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OrganizationOrganizationalUnitClientApplicationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(OrganizationOrganizationalUnitClientApplicationable)
                }
            }
            m.SetItems(res)
        }
        return nil
    }
    return res
}
// GetItems gets the items property value. The items property
// returns a []OrganizationOrganizationalUnitClientApplicationable when successful
func (m *OrganizationOrganizationalUnitClientApplicationLinkableWrapper) GetItems()([]OrganizationOrganizationalUnitClientApplicationable) {
    return m.items
}
// Serialize serializes information the current object
func (m *OrganizationOrganizationalUnitClientApplicationLinkableWrapper) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *OrganizationOrganizationalUnitClientApplicationLinkableWrapper) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetItems sets the items property value. The items property
func (m *OrganizationOrganizationalUnitClientApplicationLinkableWrapper) SetItems(value []OrganizationOrganizationalUnitClientApplicationable)() {
    m.items = value
}
type OrganizationOrganizationalUnitClientApplicationLinkableWrapperable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetItems()([]OrganizationOrganizationalUnitClientApplicationable)
    SetItems(value []OrganizationOrganizationalUnitClientApplicationable)()
}
