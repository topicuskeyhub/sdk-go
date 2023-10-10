package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VaultVaultRecordPrimerLinkableWrapper 
type VaultVaultRecordPrimerLinkableWrapper struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The items property
    items []VaultVaultRecordPrimerable
}
// NewVaultVaultRecordPrimerLinkableWrapper instantiates a new vaultVaultRecordPrimerLinkableWrapper and sets the default values.
func NewVaultVaultRecordPrimerLinkableWrapper()(*VaultVaultRecordPrimerLinkableWrapper) {
    m := &VaultVaultRecordPrimerLinkableWrapper{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateVaultVaultRecordPrimerLinkableWrapperFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVaultVaultRecordPrimerLinkableWrapperFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVaultVaultRecordPrimerLinkableWrapper(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *VaultVaultRecordPrimerLinkableWrapper) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *VaultVaultRecordPrimerLinkableWrapper) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["items"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateVaultVaultRecordPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]VaultVaultRecordPrimerable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(VaultVaultRecordPrimerable)
                }
            }
            m.SetItems(res)
        }
        return nil
    }
    return res
}
// GetItems gets the items property value. The items property
func (m *VaultVaultRecordPrimerLinkableWrapper) GetItems()([]VaultVaultRecordPrimerable) {
    return m.items
}
// Serialize serializes information the current object
func (m *VaultVaultRecordPrimerLinkableWrapper) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *VaultVaultRecordPrimerLinkableWrapper) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetItems sets the items property value. The items property
func (m *VaultVaultRecordPrimerLinkableWrapper) SetItems(value []VaultVaultRecordPrimerable)() {
    m.items = value
}
// VaultVaultRecordPrimerLinkableWrapperable 
type VaultVaultRecordPrimerLinkableWrapperable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetItems()([]VaultVaultRecordPrimerable)
    SetItems(value []VaultVaultRecordPrimerable)()
}
