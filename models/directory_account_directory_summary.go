package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DirectoryAccountDirectorySummary 
type DirectoryAccountDirectorySummary struct {
    Linkable
    // The type property
    directoryAccountDirectorySummaryType *DirectoryAccountDirectoryType
    // The domainRestriction property
    domainRestriction *string
    // The fullyResolvedIssuer property
    fullyResolvedIssuer *string
    // The name property
    name *string
    // The status property
    status DirectoryAccountDirectoryStatusReportable
    // The usernameCustomizable property
    usernameCustomizable *bool
}
// NewDirectoryAccountDirectorySummary instantiates a new directoryAccountDirectorySummary and sets the default values.
func NewDirectoryAccountDirectorySummary()(*DirectoryAccountDirectorySummary) {
    m := &DirectoryAccountDirectorySummary{
        Linkable: *NewLinkable(),
    }
    typeEscapedValue := "directory.AccountDirectorySummary"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateDirectoryAccountDirectorySummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDirectoryAccountDirectorySummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDirectoryAccountDirectorySummary(), nil
}
// GetDirectoryAccountDirectorySummaryType gets the type property value. The type property
func (m *DirectoryAccountDirectorySummary) GetDirectoryAccountDirectorySummaryType()(*DirectoryAccountDirectoryType) {
    return m.directoryAccountDirectorySummaryType
}
// GetDomainRestriction gets the domainRestriction property value. The domainRestriction property
func (m *DirectoryAccountDirectorySummary) GetDomainRestriction()(*string) {
    return m.domainRestriction
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DirectoryAccountDirectorySummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Linkable.GetFieldDeserializers()
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDirectoryAccountDirectoryType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDirectoryAccountDirectorySummaryType(val.(*DirectoryAccountDirectoryType))
        }
        return nil
    }
    res["domainRestriction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDomainRestriction(val)
        }
        return nil
    }
    res["fullyResolvedIssuer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFullyResolvedIssuer(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDirectoryAccountDirectoryStatusReportFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(DirectoryAccountDirectoryStatusReportable))
        }
        return nil
    }
    res["usernameCustomizable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsernameCustomizable(val)
        }
        return nil
    }
    return res
}
// GetFullyResolvedIssuer gets the fullyResolvedIssuer property value. The fullyResolvedIssuer property
func (m *DirectoryAccountDirectorySummary) GetFullyResolvedIssuer()(*string) {
    return m.fullyResolvedIssuer
}
// GetName gets the name property value. The name property
func (m *DirectoryAccountDirectorySummary) GetName()(*string) {
    return m.name
}
// GetStatus gets the status property value. The status property
func (m *DirectoryAccountDirectorySummary) GetStatus()(DirectoryAccountDirectoryStatusReportable) {
    return m.status
}
// GetUsernameCustomizable gets the usernameCustomizable property value. The usernameCustomizable property
func (m *DirectoryAccountDirectorySummary) GetUsernameCustomizable()(*bool) {
    return m.usernameCustomizable
}
// Serialize serializes information the current object
func (m *DirectoryAccountDirectorySummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Linkable.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDirectoryAccountDirectorySummaryType() != nil {
        cast := (*m.GetDirectoryAccountDirectorySummaryType()).String()
        err = writer.WriteStringValue("type", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("domainRestriction", m.GetDomainRestriction())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("fullyResolvedIssuer", m.GetFullyResolvedIssuer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("usernameCustomizable", m.GetUsernameCustomizable())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDirectoryAccountDirectorySummaryType sets the type property value. The type property
func (m *DirectoryAccountDirectorySummary) SetDirectoryAccountDirectorySummaryType(value *DirectoryAccountDirectoryType)() {
    m.directoryAccountDirectorySummaryType = value
}
// SetDomainRestriction sets the domainRestriction property value. The domainRestriction property
func (m *DirectoryAccountDirectorySummary) SetDomainRestriction(value *string)() {
    m.domainRestriction = value
}
// SetFullyResolvedIssuer sets the fullyResolvedIssuer property value. The fullyResolvedIssuer property
func (m *DirectoryAccountDirectorySummary) SetFullyResolvedIssuer(value *string)() {
    m.fullyResolvedIssuer = value
}
// SetName sets the name property value. The name property
func (m *DirectoryAccountDirectorySummary) SetName(value *string)() {
    m.name = value
}
// SetStatus sets the status property value. The status property
func (m *DirectoryAccountDirectorySummary) SetStatus(value DirectoryAccountDirectoryStatusReportable)() {
    m.status = value
}
// SetUsernameCustomizable sets the usernameCustomizable property value. The usernameCustomizable property
func (m *DirectoryAccountDirectorySummary) SetUsernameCustomizable(value *bool)() {
    m.usernameCustomizable = value
}
// DirectoryAccountDirectorySummaryable 
type DirectoryAccountDirectorySummaryable interface {
    Linkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDirectoryAccountDirectorySummaryType()(*DirectoryAccountDirectoryType)
    GetDomainRestriction()(*string)
    GetFullyResolvedIssuer()(*string)
    GetName()(*string)
    GetStatus()(DirectoryAccountDirectoryStatusReportable)
    GetUsernameCustomizable()(*bool)
    SetDirectoryAccountDirectorySummaryType(value *DirectoryAccountDirectoryType)()
    SetDomainRestriction(value *string)()
    SetFullyResolvedIssuer(value *string)()
    SetName(value *string)()
    SetStatus(value DirectoryAccountDirectoryStatusReportable)()
    SetUsernameCustomizable(value *bool)()
}
