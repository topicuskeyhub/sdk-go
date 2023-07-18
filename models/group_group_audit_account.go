package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupGroupAuditAccount 
type GroupGroupAuditAccount struct {
    Linkable
    // The accountUuid property
    accountUuid *string
    // The accountValid property
    accountValid *bool
    // The action property
    action *AuditAuditAccountAction
    // The comment property
    comment *string
    // The disconnectedNested property
    disconnectedNested *bool
    // The displayName property
    displayName *string
    // The endDate property
    endDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The lastActive property
    lastActive *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The lastUsed property
    lastUsed *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The nested property
    nested *bool
    // The rights property
    rights *GroupGroupRights
    // The username property
    username *string
}
// NewGroupGroupAuditAccount instantiates a new groupGroupAuditAccount and sets the default values.
func NewGroupGroupAuditAccount()(*GroupGroupAuditAccount) {
    m := &GroupGroupAuditAccount{
        Linkable: *NewLinkable(),
    }
    typeEscapedValue := "group.GroupAuditAccount"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateGroupGroupAuditAccountFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupGroupAuditAccountFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupGroupAuditAccount(), nil
}
// GetAccountUuid gets the accountUuid property value. The accountUuid property
func (m *GroupGroupAuditAccount) GetAccountUuid()(*string) {
    return m.accountUuid
}
// GetAccountValid gets the accountValid property value. The accountValid property
func (m *GroupGroupAuditAccount) GetAccountValid()(*bool) {
    return m.accountValid
}
// GetAction gets the action property value. The action property
func (m *GroupGroupAuditAccount) GetAction()(*AuditAuditAccountAction) {
    return m.action
}
// GetComment gets the comment property value. The comment property
func (m *GroupGroupAuditAccount) GetComment()(*string) {
    return m.comment
}
// GetDisconnectedNested gets the disconnectedNested property value. The disconnectedNested property
func (m *GroupGroupAuditAccount) GetDisconnectedNested()(*bool) {
    return m.disconnectedNested
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *GroupGroupAuditAccount) GetDisplayName()(*string) {
    return m.displayName
}
// GetEndDate gets the endDate property value. The endDate property
func (m *GroupGroupAuditAccount) GetEndDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.endDate
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupGroupAuditAccount) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Linkable.GetFieldDeserializers()
    res["accountUuid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountUuid(val)
        }
        return nil
    }
    res["accountValid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountValid(val)
        }
        return nil
    }
    res["action"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAuditAuditAccountAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAction(val.(*AuditAuditAccountAction))
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
    res["disconnectedNested"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisconnectedNested(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["endDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDate(val)
        }
        return nil
    }
    res["lastActive"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastActive(val)
        }
        return nil
    }
    res["lastUsed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastUsed(val)
        }
        return nil
    }
    res["nested"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNested(val)
        }
        return nil
    }
    res["rights"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseGroupGroupRights)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRights(val.(*GroupGroupRights))
        }
        return nil
    }
    res["username"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsername(val)
        }
        return nil
    }
    return res
}
// GetLastActive gets the lastActive property value. The lastActive property
func (m *GroupGroupAuditAccount) GetLastActive()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastActive
}
// GetLastUsed gets the lastUsed property value. The lastUsed property
func (m *GroupGroupAuditAccount) GetLastUsed()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.lastUsed
}
// GetNested gets the nested property value. The nested property
func (m *GroupGroupAuditAccount) GetNested()(*bool) {
    return m.nested
}
// GetRights gets the rights property value. The rights property
func (m *GroupGroupAuditAccount) GetRights()(*GroupGroupRights) {
    return m.rights
}
// GetUsername gets the username property value. The username property
func (m *GroupGroupAuditAccount) GetUsername()(*string) {
    return m.username
}
// Serialize serializes information the current object
func (m *GroupGroupAuditAccount) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Linkable.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("accountUuid", m.GetAccountUuid())
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
        err = writer.WriteStringValue("comment", m.GetComment())
        if err != nil {
            return err
        }
    }
    if m.GetRights() != nil {
        cast := (*m.GetRights()).String()
        err = writer.WriteStringValue("rights", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccountUuid sets the accountUuid property value. The accountUuid property
func (m *GroupGroupAuditAccount) SetAccountUuid(value *string)() {
    m.accountUuid = value
}
// SetAccountValid sets the accountValid property value. The accountValid property
func (m *GroupGroupAuditAccount) SetAccountValid(value *bool)() {
    m.accountValid = value
}
// SetAction sets the action property value. The action property
func (m *GroupGroupAuditAccount) SetAction(value *AuditAuditAccountAction)() {
    m.action = value
}
// SetComment sets the comment property value. The comment property
func (m *GroupGroupAuditAccount) SetComment(value *string)() {
    m.comment = value
}
// SetDisconnectedNested sets the disconnectedNested property value. The disconnectedNested property
func (m *GroupGroupAuditAccount) SetDisconnectedNested(value *bool)() {
    m.disconnectedNested = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *GroupGroupAuditAccount) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEndDate sets the endDate property value. The endDate property
func (m *GroupGroupAuditAccount) SetEndDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.endDate = value
}
// SetLastActive sets the lastActive property value. The lastActive property
func (m *GroupGroupAuditAccount) SetLastActive(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastActive = value
}
// SetLastUsed sets the lastUsed property value. The lastUsed property
func (m *GroupGroupAuditAccount) SetLastUsed(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.lastUsed = value
}
// SetNested sets the nested property value. The nested property
func (m *GroupGroupAuditAccount) SetNested(value *bool)() {
    m.nested = value
}
// SetRights sets the rights property value. The rights property
func (m *GroupGroupAuditAccount) SetRights(value *GroupGroupRights)() {
    m.rights = value
}
// SetUsername sets the username property value. The username property
func (m *GroupGroupAuditAccount) SetUsername(value *string)() {
    m.username = value
}
// GroupGroupAuditAccountable 
type GroupGroupAuditAccountable interface {
    Linkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccountUuid()(*string)
    GetAccountValid()(*bool)
    GetAction()(*AuditAuditAccountAction)
    GetComment()(*string)
    GetDisconnectedNested()(*bool)
    GetDisplayName()(*string)
    GetEndDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetLastActive()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastUsed()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetNested()(*bool)
    GetRights()(*GroupGroupRights)
    GetUsername()(*string)
    SetAccountUuid(value *string)()
    SetAccountValid(value *bool)()
    SetAction(value *AuditAuditAccountAction)()
    SetComment(value *string)()
    SetDisconnectedNested(value *bool)()
    SetDisplayName(value *string)()
    SetEndDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetLastActive(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastUsed(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetNested(value *bool)()
    SetRights(value *GroupGroupRights)()
    SetUsername(value *string)()
}
