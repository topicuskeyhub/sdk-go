package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AuthMoveInternalAccount struct {
    NonLinkable
    // The internalDirectory property
    internalDirectory DirectoryInternalDirectoryable
}
// NewAuthMoveInternalAccount instantiates a new AuthMoveInternalAccount and sets the default values.
func NewAuthMoveInternalAccount()(*AuthMoveInternalAccount) {
    m := &AuthMoveInternalAccount{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "auth.MoveInternalAccount"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateAuthMoveInternalAccountFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAuthMoveInternalAccountFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthMoveInternalAccount(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AuthMoveInternalAccount) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
    res["internalDirectory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDirectoryInternalDirectoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInternalDirectory(val.(DirectoryInternalDirectoryable))
        }
        return nil
    }
    return res
}
// GetInternalDirectory gets the internalDirectory property value. The internalDirectory property
// returns a DirectoryInternalDirectoryable when successful
func (m *AuthMoveInternalAccount) GetInternalDirectory()(DirectoryInternalDirectoryable) {
    return m.internalDirectory
}
// Serialize serializes information the current object
func (m *AuthMoveInternalAccount) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("internalDirectory", m.GetInternalDirectory())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetInternalDirectory sets the internalDirectory property value. The internalDirectory property
func (m *AuthMoveInternalAccount) SetInternalDirectory(value DirectoryInternalDirectoryable)() {
    m.internalDirectory = value
}
type AuthMoveInternalAccountable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetInternalDirectory()(DirectoryInternalDirectoryable)
    SetInternalDirectory(value DirectoryInternalDirectoryable)()
}
