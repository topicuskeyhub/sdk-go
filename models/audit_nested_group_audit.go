package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuditNestedGroupAudit 
type AuditNestedGroupAudit struct {
    Linkable
    // The action property
    action *AuditAuditNestedGroupAction
    // The comment property
    comment *string
    // The groupUuid property
    groupUuid *string
    // The name property
    name *string
}
// NewAuditNestedGroupAudit instantiates a new auditNestedGroupAudit and sets the default values.
func NewAuditNestedGroupAudit()(*AuditNestedGroupAudit) {
    m := &AuditNestedGroupAudit{
        Linkable: *NewLinkable(),
    }
    typeEscapedValue := "audit.NestedGroupAudit"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateAuditNestedGroupAuditFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuditNestedGroupAuditFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuditNestedGroupAudit(), nil
}
// GetAction gets the action property value. The action property
func (m *AuditNestedGroupAudit) GetAction()(*AuditAuditNestedGroupAction) {
    return m.action
}
// GetComment gets the comment property value. The comment property
func (m *AuditNestedGroupAudit) GetComment()(*string) {
    return m.comment
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuditNestedGroupAudit) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Linkable.GetFieldDeserializers()
    res["action"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAuditAuditNestedGroupAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAction(val.(*AuditAuditNestedGroupAction))
        }
        return nil
    }
    res["comment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComment(val)
        }
        return nil
    }
    res["groupUuid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupUuid(val)
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
    return res
}
// GetGroupUuid gets the groupUuid property value. The groupUuid property
func (m *AuditNestedGroupAudit) GetGroupUuid()(*string) {
    return m.groupUuid
}
// GetName gets the name property value. The name property
func (m *AuditNestedGroupAudit) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *AuditNestedGroupAudit) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Linkable.Serialize(writer)
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
        err = writer.WriteStringValue("comment", m.GetComment())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("groupUuid", m.GetGroupUuid())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAction sets the action property value. The action property
func (m *AuditNestedGroupAudit) SetAction(value *AuditAuditNestedGroupAction)() {
    m.action = value
}
// SetComment sets the comment property value. The comment property
func (m *AuditNestedGroupAudit) SetComment(value *string)() {
    m.comment = value
}
// SetGroupUuid sets the groupUuid property value. The groupUuid property
func (m *AuditNestedGroupAudit) SetGroupUuid(value *string)() {
    m.groupUuid = value
}
// SetName sets the name property value. The name property
func (m *AuditNestedGroupAudit) SetName(value *string)() {
    m.name = value
}
// AuditNestedGroupAuditable 
type AuditNestedGroupAuditable interface {
    Linkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAction()(*AuditAuditNestedGroupAction)
    GetComment()(*string)
    GetGroupUuid()(*string)
    GetName()(*string)
    SetAction(value *AuditAuditNestedGroupAction)()
    SetComment(value *string)()
    SetGroupUuid(value *string)()
    SetName(value *string)()
}
