package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ClientOAuth2ClientPermissionWithClientLinkableWrapper 
type ClientOAuth2ClientPermissionWithClientLinkableWrapper struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The items property
    items []ClientOAuth2ClientPermissionWithClientable
}
// NewClientOAuth2ClientPermissionWithClientLinkableWrapper instantiates a new clientOAuth2ClientPermissionWithClientLinkableWrapper and sets the default values.
func NewClientOAuth2ClientPermissionWithClientLinkableWrapper()(*ClientOAuth2ClientPermissionWithClientLinkableWrapper) {
    m := &ClientOAuth2ClientPermissionWithClientLinkableWrapper{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateClientOAuth2ClientPermissionWithClientLinkableWrapperFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateClientOAuth2ClientPermissionWithClientLinkableWrapperFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewClientOAuth2ClientPermissionWithClientLinkableWrapper(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ClientOAuth2ClientPermissionWithClientLinkableWrapper) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ClientOAuth2ClientPermissionWithClientLinkableWrapper) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["items"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateClientOAuth2ClientPermissionWithClientFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ClientOAuth2ClientPermissionWithClientable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ClientOAuth2ClientPermissionWithClientable)
                }
            }
            m.SetItems(res)
        }
        return nil
    }
    return res
}
// GetItems gets the items property value. The items property
func (m *ClientOAuth2ClientPermissionWithClientLinkableWrapper) GetItems()([]ClientOAuth2ClientPermissionWithClientable) {
    return m.items
}
// Serialize serializes information the current object
func (m *ClientOAuth2ClientPermissionWithClientLinkableWrapper) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ClientOAuth2ClientPermissionWithClientLinkableWrapper) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetItems sets the items property value. The items property
func (m *ClientOAuth2ClientPermissionWithClientLinkableWrapper) SetItems(value []ClientOAuth2ClientPermissionWithClientable)() {
    m.items = value
}
// ClientOAuth2ClientPermissionWithClientLinkableWrapperable 
type ClientOAuth2ClientPermissionWithClientLinkableWrapperable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetItems()([]ClientOAuth2ClientPermissionWithClientable)
    SetItems(value []ClientOAuth2ClientPermissionWithClientable)()
}
