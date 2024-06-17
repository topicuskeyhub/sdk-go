package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type DirectoryPendingAccountsDirectory struct {
    DirectoryAccountDirectory
}
// NewDirectoryPendingAccountsDirectory instantiates a new DirectoryPendingAccountsDirectory and sets the default values.
func NewDirectoryPendingAccountsDirectory()(*DirectoryPendingAccountsDirectory) {
    m := &DirectoryPendingAccountsDirectory{
        DirectoryAccountDirectory: *NewDirectoryAccountDirectory(),
    }
    typeEscapedValue := "directory.PendingAccountsDirectory"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateDirectoryPendingAccountsDirectoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDirectoryPendingAccountsDirectoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDirectoryPendingAccountsDirectory(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DirectoryPendingAccountsDirectory) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DirectoryAccountDirectory.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *DirectoryPendingAccountsDirectory) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DirectoryAccountDirectory.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type DirectoryPendingAccountsDirectoryable interface {
    DirectoryAccountDirectoryable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
