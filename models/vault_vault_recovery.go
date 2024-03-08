package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type VaultVaultRecovery struct {
    NonLinkable
    // The account property
    account AuthAccountPrimerable
    // The privateKey property
    privateKey *string
}
// NewVaultVaultRecovery instantiates a new VaultVaultRecovery and sets the default values.
func NewVaultVaultRecovery()(*VaultVaultRecovery) {
    m := &VaultVaultRecovery{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "vault.VaultRecovery"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateVaultVaultRecoveryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateVaultVaultRecoveryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVaultVaultRecovery(), nil
}
// GetAccount gets the account property value. The account property
// returns a AuthAccountPrimerable when successful
func (m *VaultVaultRecovery) GetAccount()(AuthAccountPrimerable) {
    return m.account
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *VaultVaultRecovery) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["privateKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivateKey(val)
        }
        return nil
    }
    return res
}
// GetPrivateKey gets the privateKey property value. The privateKey property
// returns a *string when successful
func (m *VaultVaultRecovery) GetPrivateKey()(*string) {
    return m.privateKey
}
// Serialize serializes information the current object
func (m *VaultVaultRecovery) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err = writer.WriteStringValue("privateKey", m.GetPrivateKey())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccount sets the account property value. The account property
func (m *VaultVaultRecovery) SetAccount(value AuthAccountPrimerable)() {
    m.account = value
}
// SetPrivateKey sets the privateKey property value. The privateKey property
func (m *VaultVaultRecovery) SetPrivateKey(value *string)() {
    m.privateKey = value
}
type VaultVaultRecoveryable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccount()(AuthAccountPrimerable)
    GetPrivateKey()(*string)
    SetAccount(value AuthAccountPrimerable)()
    SetPrivateKey(value *string)()
}
