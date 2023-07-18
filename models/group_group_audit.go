package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupGroupAudit 
type GroupGroupAudit struct {
    Linkable
    // The accounts property
    accounts []GroupGroupAuditAccountable
    // The additionalObjects property
    additionalObjects GroupGroupAudit_additionalObjectsable
    // The comment property
    comment *string
    // The createdAt property
    createdAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The createdBy property
    createdBy *string
    // The groupName property
    groupName *string
    // The nameOnAudit property
    nameOnAudit *string
    // The reviewedAt property
    reviewedAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The reviewedBy property
    reviewedBy *string
    // The status property
    status *GroupGroupAuditStatus
    // The submittedAt property
    submittedAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The submittedBy property
    submittedBy *string
}
// NewGroupGroupAudit instantiates a new groupGroupAudit and sets the default values.
func NewGroupGroupAudit()(*GroupGroupAudit) {
    m := &GroupGroupAudit{
        Linkable: *NewLinkable(),
    }
    typeEscapedValue := "group.GroupAudit"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateGroupGroupAuditFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupGroupAuditFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupGroupAudit(), nil
}
// GetAccounts gets the accounts property value. The accounts property
func (m *GroupGroupAudit) GetAccounts()([]GroupGroupAuditAccountable) {
    return m.accounts
}
// GetAdditionalObjects gets the additionalObjects property value. The additionalObjects property
func (m *GroupGroupAudit) GetAdditionalObjects()(GroupGroupAudit_additionalObjectsable) {
    return m.additionalObjects
}
// GetComment gets the comment property value. The comment property
func (m *GroupGroupAudit) GetComment()(*string) {
    return m.comment
}
// GetCreatedAt gets the createdAt property value. The createdAt property
func (m *GroupGroupAudit) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdAt
}
// GetCreatedBy gets the createdBy property value. The createdBy property
func (m *GroupGroupAudit) GetCreatedBy()(*string) {
    return m.createdBy
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupGroupAudit) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Linkable.GetFieldDeserializers()
    res["accounts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGroupGroupAuditAccountFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GroupGroupAuditAccountable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(GroupGroupAuditAccountable)
                }
            }
            m.SetAccounts(res)
        }
        return nil
    }
    res["additionalObjects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupGroupAudit_additionalObjectsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalObjects(val.(GroupGroupAudit_additionalObjectsable))
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
    res["groupName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupName(val)
        }
        return nil
    }
    res["nameOnAudit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNameOnAudit(val)
        }
        return nil
    }
    res["reviewedAt"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewedAt(val)
        }
        return nil
    }
    res["reviewedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewedBy(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseGroupGroupAuditStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*GroupGroupAuditStatus))
        }
        return nil
    }
    res["submittedAt"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubmittedAt(val)
        }
        return nil
    }
    res["submittedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubmittedBy(val)
        }
        return nil
    }
    return res
}
// GetGroupName gets the groupName property value. The groupName property
func (m *GroupGroupAudit) GetGroupName()(*string) {
    return m.groupName
}
// GetNameOnAudit gets the nameOnAudit property value. The nameOnAudit property
func (m *GroupGroupAudit) GetNameOnAudit()(*string) {
    return m.nameOnAudit
}
// GetReviewedAt gets the reviewedAt property value. The reviewedAt property
func (m *GroupGroupAudit) GetReviewedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.reviewedAt
}
// GetReviewedBy gets the reviewedBy property value. The reviewedBy property
func (m *GroupGroupAudit) GetReviewedBy()(*string) {
    return m.reviewedBy
}
// GetStatus gets the status property value. The status property
func (m *GroupGroupAudit) GetStatus()(*GroupGroupAuditStatus) {
    return m.status
}
// GetSubmittedAt gets the submittedAt property value. The submittedAt property
func (m *GroupGroupAudit) GetSubmittedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.submittedAt
}
// GetSubmittedBy gets the submittedBy property value. The submittedBy property
func (m *GroupGroupAudit) GetSubmittedBy()(*string) {
    return m.submittedBy
}
// Serialize serializes information the current object
func (m *GroupGroupAudit) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Linkable.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAccounts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAccounts()))
        for i, v := range m.GetAccounts() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("accounts", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("additionalObjects", m.GetAdditionalObjects())
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
func (m *GroupGroupAudit) SetAccounts(value []GroupGroupAuditAccountable)() {
    m.accounts = value
}
// SetAdditionalObjects sets the additionalObjects property value. The additionalObjects property
func (m *GroupGroupAudit) SetAdditionalObjects(value GroupGroupAudit_additionalObjectsable)() {
    m.additionalObjects = value
}
// SetComment sets the comment property value. The comment property
func (m *GroupGroupAudit) SetComment(value *string)() {
    m.comment = value
}
// SetCreatedAt sets the createdAt property value. The createdAt property
func (m *GroupGroupAudit) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdAt = value
}
// SetCreatedBy sets the createdBy property value. The createdBy property
func (m *GroupGroupAudit) SetCreatedBy(value *string)() {
    m.createdBy = value
}
// SetGroupName sets the groupName property value. The groupName property
func (m *GroupGroupAudit) SetGroupName(value *string)() {
    m.groupName = value
}
// SetNameOnAudit sets the nameOnAudit property value. The nameOnAudit property
func (m *GroupGroupAudit) SetNameOnAudit(value *string)() {
    m.nameOnAudit = value
}
// SetReviewedAt sets the reviewedAt property value. The reviewedAt property
func (m *GroupGroupAudit) SetReviewedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.reviewedAt = value
}
// SetReviewedBy sets the reviewedBy property value. The reviewedBy property
func (m *GroupGroupAudit) SetReviewedBy(value *string)() {
    m.reviewedBy = value
}
// SetStatus sets the status property value. The status property
func (m *GroupGroupAudit) SetStatus(value *GroupGroupAuditStatus)() {
    m.status = value
}
// SetSubmittedAt sets the submittedAt property value. The submittedAt property
func (m *GroupGroupAudit) SetSubmittedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.submittedAt = value
}
// SetSubmittedBy sets the submittedBy property value. The submittedBy property
func (m *GroupGroupAudit) SetSubmittedBy(value *string)() {
    m.submittedBy = value
}
// GroupGroupAuditable 
type GroupGroupAuditable interface {
    Linkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccounts()([]GroupGroupAuditAccountable)
    GetAdditionalObjects()(GroupGroupAudit_additionalObjectsable)
    GetComment()(*string)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCreatedBy()(*string)
    GetGroupName()(*string)
    GetNameOnAudit()(*string)
    GetReviewedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetReviewedBy()(*string)
    GetStatus()(*GroupGroupAuditStatus)
    GetSubmittedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSubmittedBy()(*string)
    SetAccounts(value []GroupGroupAuditAccountable)()
    SetAdditionalObjects(value GroupGroupAudit_additionalObjectsable)()
    SetComment(value *string)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCreatedBy(value *string)()
    SetGroupName(value *string)()
    SetNameOnAudit(value *string)()
    SetReviewedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetReviewedBy(value *string)()
    SetStatus(value *GroupGroupAuditStatus)()
    SetSubmittedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSubmittedBy(value *string)()
}
