package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type VaultMoveVaultRecord struct {
    NonLinkable
    // The account property
    account AuthAccountPrimerable
    // The action property
    action *VaultMoveVaultRecordAction
    // The group property
    group GroupGroupPrimerable
    // The shareDuration property
    shareDuration VaultMoveVaultRecord_shareDurationable
}
// NewVaultMoveVaultRecord instantiates a new VaultMoveVaultRecord and sets the default values.
func NewVaultMoveVaultRecord()(*VaultMoveVaultRecord) {
    m := &VaultMoveVaultRecord{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "vault.MoveVaultRecord"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateVaultMoveVaultRecordFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateVaultMoveVaultRecordFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVaultMoveVaultRecord(), nil
}
// GetAccount gets the account property value. The account property
// returns a AuthAccountPrimerable when successful
func (m *VaultMoveVaultRecord) GetAccount()(AuthAccountPrimerable) {
    return m.account
}
// GetAction gets the action property value. The action property
// returns a *VaultMoveVaultRecordAction when successful
func (m *VaultMoveVaultRecord) GetAction()(*VaultMoveVaultRecordAction) {
    return m.action
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *VaultMoveVaultRecord) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
    res["account"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthAccountPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccount(val.(AuthAccountPrimerable))
        }
        return nil
    }
    res["action"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVaultMoveVaultRecordAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAction(val.(*VaultMoveVaultRecordAction))
        }
        return nil
    }
    res["group"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupGroupPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroup(val.(GroupGroupPrimerable))
        }
        return nil
    }
    res["shareDuration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateVaultMoveVaultRecord_shareDurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShareDuration(val.(VaultMoveVaultRecord_shareDurationable))
        }
        return nil
    }
    return res
}
// GetGroup gets the group property value. The group property
// returns a GroupGroupPrimerable when successful
func (m *VaultMoveVaultRecord) GetGroup()(GroupGroupPrimerable) {
    return m.group
}
// GetShareDuration gets the shareDuration property value. The shareDuration property
// returns a VaultMoveVaultRecord_shareDurationable when successful
func (m *VaultMoveVaultRecord) GetShareDuration()(VaultMoveVaultRecord_shareDurationable) {
    return m.shareDuration
}
// Serialize serializes information the current object
func (m *VaultMoveVaultRecord) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("account", m.GetAccount())
        if err != nil {
            return err
        }
    }
    if m.GetAction() != nil {
        cast := (*m.GetAction()).String()
        err = writer.WriteStringValue("action", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("group", m.GetGroup())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("shareDuration", m.GetShareDuration())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccount sets the account property value. The account property
func (m *VaultMoveVaultRecord) SetAccount(value AuthAccountPrimerable)() {
    m.account = value
}
// SetAction sets the action property value. The action property
func (m *VaultMoveVaultRecord) SetAction(value *VaultMoveVaultRecordAction)() {
    m.action = value
}
// SetGroup sets the group property value. The group property
func (m *VaultMoveVaultRecord) SetGroup(value GroupGroupPrimerable)() {
    m.group = value
}
// SetShareDuration sets the shareDuration property value. The shareDuration property
func (m *VaultMoveVaultRecord) SetShareDuration(value VaultMoveVaultRecord_shareDurationable)() {
    m.shareDuration = value
}
type VaultMoveVaultRecordable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccount()(AuthAccountPrimerable)
    GetAction()(*VaultMoveVaultRecordAction)
    GetGroup()(GroupGroupPrimerable)
    GetShareDuration()(VaultMoveVaultRecord_shareDurationable)
    SetAccount(value AuthAccountPrimerable)()
    SetAction(value *VaultMoveVaultRecordAction)()
    SetGroup(value GroupGroupPrimerable)()
    SetShareDuration(value VaultMoveVaultRecord_shareDurationable)()
}
