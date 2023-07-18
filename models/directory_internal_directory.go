package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DirectoryInternalDirectory 
type DirectoryInternalDirectory struct {
    DirectoryAccountDirectory
    // The owner property
    owner GroupGroupPrimerable
}
// NewDirectoryInternalDirectory instantiates a new directoryInternalDirectory and sets the default values.
func NewDirectoryInternalDirectory()(*DirectoryInternalDirectory) {
    m := &DirectoryInternalDirectory{
        DirectoryAccountDirectory: *NewDirectoryAccountDirectory(),
    }
    typeEscapedValue := "directory.InternalDirectory"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateDirectoryInternalDirectoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDirectoryInternalDirectoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDirectoryInternalDirectory(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DirectoryInternalDirectory) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DirectoryAccountDirectory.GetFieldDeserializers()
    res["owner"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupGroupPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwner(val.(GroupGroupPrimerable))
        }
        return nil
    }
    return res
}
// GetOwner gets the owner property value. The owner property
func (m *DirectoryInternalDirectory) GetOwner()(GroupGroupPrimerable) {
    return m.owner
}
// Serialize serializes information the current object
func (m *DirectoryInternalDirectory) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DirectoryAccountDirectory.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("owner", m.GetOwner())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOwner sets the owner property value. The owner property
func (m *DirectoryInternalDirectory) SetOwner(value GroupGroupPrimerable)() {
    m.owner = value
}
// DirectoryInternalDirectoryable 
type DirectoryInternalDirectoryable interface {
    DirectoryAccountDirectoryable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOwner()(GroupGroupPrimerable)
    SetOwner(value GroupGroupPrimerable)()
}
