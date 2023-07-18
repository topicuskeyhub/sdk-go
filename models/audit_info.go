package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuditInfo 
type AuditInfo struct {
    NonLinkable
    // The createdAt property
    createdAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The createdBy property
    createdBy *string
    // The lastModifiedAt property
    lastModifiedAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The lastModifiedBy property
    lastModifiedBy *string
}
// NewAuditInfo instantiates a new AuditInfo and sets the default values.
func NewAuditInfo()(*AuditInfo) {
    m := &AuditInfo{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "AuditInfo"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateAuditInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuditInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuditInfo(), nil
}
// GetCreatedAt gets the createdAt property value. The createdAt property
func (m *AuditInfo) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdAt
}
// GetCreatedBy gets the createdBy property value. The createdBy property
func (m *AuditInfo) GetCreatedBy()(*string) {
    return m.createdBy
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuditInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
    res["createdAt"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedAt(val)
        }
        return nil
    }
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val)
        }
        return nil
    }
    res["lastModifiedAt"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedAt(val)
        }
        return nil
    }
    res["lastModifiedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedBy(val)
        }
        return nil
    }
    return res
}
// GetLastModifiedAt gets the lastModifiedAt property value. The lastModifiedAt property
func (m *AuditInfo) GetLastModifiedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedAt
}
// GetLastModifiedBy gets the lastModifiedBy property value. The lastModifiedBy property
func (m *AuditInfo) GetLastModifiedBy()(*string) {
    return m.lastModifiedBy
}
// Serialize serializes information the current object
func (m *AuditInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("createdAt", m.GetCreatedAt())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("createdBy", m.GetCreatedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedAt", m.GetLastModifiedAt())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("lastModifiedBy", m.GetLastModifiedBy())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreatedAt sets the createdAt property value. The createdAt property
func (m *AuditInfo) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdAt = value
}
// SetCreatedBy sets the createdBy property value. The createdBy property
func (m *AuditInfo) SetCreatedBy(value *string)() {
    m.createdBy = value
}
// SetLastModifiedAt sets the lastModifiedAt property value. The lastModifiedAt property
func (m *AuditInfo) SetLastModifiedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedAt = value
}
// SetLastModifiedBy sets the lastModifiedBy property value. The lastModifiedBy property
func (m *AuditInfo) SetLastModifiedBy(value *string)() {
    m.lastModifiedBy = value
}
// AuditInfoable 
type AuditInfoable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCreatedBy()(*string)
    GetLastModifiedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastModifiedBy()(*string)
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCreatedBy(value *string)()
    SetLastModifiedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastModifiedBy(value *string)()
}
