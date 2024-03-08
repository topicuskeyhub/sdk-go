package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AuthAccountBulkUpdate struct {
    NonLinkable
    // The action property
    action *AuthAccountBulkUpdateAction
    // The exclusiveSelection property
    exclusiveSelection *bool
    // The selectedAccounts property
    selectedAccounts []int64
}
// NewAuthAccountBulkUpdate instantiates a new AuthAccountBulkUpdate and sets the default values.
func NewAuthAccountBulkUpdate()(*AuthAccountBulkUpdate) {
    m := &AuthAccountBulkUpdate{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "auth.AccountBulkUpdate"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateAuthAccountBulkUpdateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAuthAccountBulkUpdateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthAccountBulkUpdate(), nil
}
// GetAction gets the action property value. The action property
// returns a *AuthAccountBulkUpdateAction when successful
func (m *AuthAccountBulkUpdate) GetAction()(*AuthAccountBulkUpdateAction) {
    return m.action
}
// GetExclusiveSelection gets the exclusiveSelection property value. The exclusiveSelection property
// returns a *bool when successful
func (m *AuthAccountBulkUpdate) GetExclusiveSelection()(*bool) {
    return m.exclusiveSelection
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AuthAccountBulkUpdate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
    res["action"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAuthAccountBulkUpdateAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAction(val.(*AuthAccountBulkUpdateAction))
        }
        return nil
    }
    res["exclusiveSelection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExclusiveSelection(val)
        }
        return nil
    }
    res["selectedAccounts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("int64")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]int64, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*int64))
                }
            }
            m.SetSelectedAccounts(res)
        }
        return nil
    }
    return res
}
// GetSelectedAccounts gets the selectedAccounts property value. The selectedAccounts property
// returns a []int64 when successful
func (m *AuthAccountBulkUpdate) GetSelectedAccounts()([]int64) {
    return m.selectedAccounts
}
// Serialize serializes information the current object
func (m *AuthAccountBulkUpdate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAction() != nil {
        cast := (*m.GetAction()).String()
        err = writer.WriteStringValue("action", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("exclusiveSelection", m.GetExclusiveSelection())
        if err != nil {
            return err
        }
    }
    if m.GetSelectedAccounts() != nil {
        err = writer.WriteCollectionOfInt64Values("selectedAccounts", m.GetSelectedAccounts())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAction sets the action property value. The action property
func (m *AuthAccountBulkUpdate) SetAction(value *AuthAccountBulkUpdateAction)() {
    m.action = value
}
// SetExclusiveSelection sets the exclusiveSelection property value. The exclusiveSelection property
func (m *AuthAccountBulkUpdate) SetExclusiveSelection(value *bool)() {
    m.exclusiveSelection = value
}
// SetSelectedAccounts sets the selectedAccounts property value. The selectedAccounts property
func (m *AuthAccountBulkUpdate) SetSelectedAccounts(value []int64)() {
    m.selectedAccounts = value
}
type AuthAccountBulkUpdateable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAction()(*AuthAccountBulkUpdateAction)
    GetExclusiveSelection()(*bool)
    GetSelectedAccounts()([]int64)
    SetAction(value *AuthAccountBulkUpdateAction)()
    SetExclusiveSelection(value *bool)()
    SetSelectedAccounts(value []int64)()
}
