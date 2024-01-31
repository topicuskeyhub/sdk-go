package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DirectoryAccountDirectoryStatusReport 
type DirectoryAccountDirectoryStatusReport struct {
    NonLinkable
    // The accounts property
    accounts *int64
    // The reason property
    reason *string
    // The status property
    status *DirectoryAccountDirectoryStatus
}
// NewDirectoryAccountDirectoryStatusReport instantiates a new directoryAccountDirectoryStatusReport and sets the default values.
func NewDirectoryAccountDirectoryStatusReport()(*DirectoryAccountDirectoryStatusReport) {
    m := &DirectoryAccountDirectoryStatusReport{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "directory.AccountDirectoryStatusReport"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateDirectoryAccountDirectoryStatusReportFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDirectoryAccountDirectoryStatusReportFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDirectoryAccountDirectoryStatusReport(), nil
}
// GetAccounts gets the accounts property value. The accounts property
func (m *DirectoryAccountDirectoryStatusReport) GetAccounts()(*int64) {
    return m.accounts
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DirectoryAccountDirectoryStatusReport) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
    res["accounts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccounts(val)
        }
        return nil
    }
    res["reason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReason(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDirectoryAccountDirectoryStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*DirectoryAccountDirectoryStatus))
        }
        return nil
    }
    return res
}
// GetReason gets the reason property value. The reason property
func (m *DirectoryAccountDirectoryStatusReport) GetReason()(*string) {
    return m.reason
}
// GetStatus gets the status property value. The status property
func (m *DirectoryAccountDirectoryStatusReport) GetStatus()(*DirectoryAccountDirectoryStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *DirectoryAccountDirectoryStatusReport) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccounts sets the accounts property value. The accounts property
func (m *DirectoryAccountDirectoryStatusReport) SetAccounts(value *int64)() {
    m.accounts = value
}
// SetReason sets the reason property value. The reason property
func (m *DirectoryAccountDirectoryStatusReport) SetReason(value *string)() {
    m.reason = value
}
// SetStatus sets the status property value. The status property
func (m *DirectoryAccountDirectoryStatusReport) SetStatus(value *DirectoryAccountDirectoryStatus)() {
    m.status = value
}
// DirectoryAccountDirectoryStatusReportable 
type DirectoryAccountDirectoryStatusReportable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccounts()(*int64)
    GetReason()(*string)
    GetStatus()(*DirectoryAccountDirectoryStatus)
    SetAccounts(value *int64)()
    SetReason(value *string)()
    SetStatus(value *DirectoryAccountDirectoryStatus)()
}
