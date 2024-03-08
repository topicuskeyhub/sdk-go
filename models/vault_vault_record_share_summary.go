package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type VaultVaultRecordShareSummary struct {
    NonLinkable
    // The children property
    children []VaultVaultRecordShareable
    // The parent property
    parent VaultVaultRecordShareable
}
// NewVaultVaultRecordShareSummary instantiates a new VaultVaultRecordShareSummary and sets the default values.
func NewVaultVaultRecordShareSummary()(*VaultVaultRecordShareSummary) {
    m := &VaultVaultRecordShareSummary{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "vault.VaultRecordShareSummary"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateVaultVaultRecordShareSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateVaultVaultRecordShareSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVaultVaultRecordShareSummary(), nil
}
// GetChildren gets the children property value. The children property
// returns a []VaultVaultRecordShareable when successful
func (m *VaultVaultRecordShareSummary) GetChildren()([]VaultVaultRecordShareable) {
    return m.children
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *VaultVaultRecordShareSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
    res["children"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateVaultVaultRecordShareFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]VaultVaultRecordShareable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(VaultVaultRecordShareable)
                }
            }
            m.SetChildren(res)
        }
        return nil
    }
    res["parent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateVaultVaultRecordShareFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParent(val.(VaultVaultRecordShareable))
        }
        return nil
    }
    return res
}
// GetParent gets the parent property value. The parent property
// returns a VaultVaultRecordShareable when successful
func (m *VaultVaultRecordShareSummary) GetParent()(VaultVaultRecordShareable) {
    return m.parent
}
// Serialize serializes information the current object
func (m *VaultVaultRecordShareSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetChildren() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetChildren()))
        for i, v := range m.GetChildren() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("children", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("parent", m.GetParent())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetChildren sets the children property value. The children property
func (m *VaultVaultRecordShareSummary) SetChildren(value []VaultVaultRecordShareable)() {
    m.children = value
}
// SetParent sets the parent property value. The parent property
func (m *VaultVaultRecordShareSummary) SetParent(value VaultVaultRecordShareable)() {
    m.parent = value
}
type VaultVaultRecordShareSummaryable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetChildren()([]VaultVaultRecordShareable)
    GetParent()(VaultVaultRecordShareable)
    SetChildren(value []VaultVaultRecordShareable)()
    SetParent(value VaultVaultRecordShareable)()
}
