package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AuditGroupAudit struct {
    Linkable
    // The accounts property
    accounts []AuditGroupAuditAccountable
    // The additionalObjects property
    additionalObjects AuditGroupAudit_additionalObjectsable
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
    // The nestedGroups property
    nestedGroups []AuditNestedGroupAuditable
    // The reviewedAt property
    reviewedAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The reviewedBy property
    reviewedBy *string
    // The status property
    status *AuditGroupAuditStatus
    // The submittedAt property
    submittedAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The submittedBy property
    submittedBy *string
}
// NewAuditGroupAudit instantiates a new AuditGroupAudit and sets the default values.
func NewAuditGroupAudit()(*AuditGroupAudit) {
    m := &AuditGroupAudit{
        Linkable: *NewLinkable(),
    }
    typeEscapedValue := "audit.GroupAudit"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateAuditGroupAuditFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAuditGroupAuditFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuditGroupAudit(), nil
}
// GetAccounts gets the accounts property value. The accounts property
// returns a []AuditGroupAuditAccountable when successful
func (m *AuditGroupAudit) GetAccounts()([]AuditGroupAuditAccountable) {
    return m.accounts
}
// GetAdditionalObjects gets the additionalObjects property value. The additionalObjects property
// returns a AuditGroupAudit_additionalObjectsable when successful
func (m *AuditGroupAudit) GetAdditionalObjects()(AuditGroupAudit_additionalObjectsable) {
    return m.additionalObjects
}
// GetComment gets the comment property value. The comment property
// returns a *string when successful
func (m *AuditGroupAudit) GetComment()(*string) {
    return m.comment
}
// GetCreatedAt gets the createdAt property value. The createdAt property
// returns a *Time when successful
func (m *AuditGroupAudit) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdAt
}
// GetCreatedBy gets the createdBy property value. The createdBy property
// returns a *string when successful
func (m *AuditGroupAudit) GetCreatedBy()(*string) {
    return m.createdBy
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AuditGroupAudit) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Linkable.GetFieldDeserializers()
    res["accounts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuditGroupAuditAccountFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuditGroupAuditAccountable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AuditGroupAuditAccountable)
                }
            }
            m.SetAccounts(res)
        }
        return nil
    }
    res["additionalObjects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuditGroupAudit_additionalObjectsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalObjects(val.(AuditGroupAudit_additionalObjectsable))
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
    res["nestedGroups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuditNestedGroupAuditFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuditNestedGroupAuditable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AuditNestedGroupAuditable)
                }
            }
            m.SetNestedGroups(res)
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
        val, err := n.GetEnumValue(ParseAuditGroupAuditStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*AuditGroupAuditStatus))
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
// returns a *string when successful
func (m *AuditGroupAudit) GetGroupName()(*string) {
    return m.groupName
}
// GetNameOnAudit gets the nameOnAudit property value. The nameOnAudit property
// returns a *string when successful
func (m *AuditGroupAudit) GetNameOnAudit()(*string) {
    return m.nameOnAudit
}
// GetNestedGroups gets the nestedGroups property value. The nestedGroups property
// returns a []AuditNestedGroupAuditable when successful
func (m *AuditGroupAudit) GetNestedGroups()([]AuditNestedGroupAuditable) {
    return m.nestedGroups
}
// GetReviewedAt gets the reviewedAt property value. The reviewedAt property
// returns a *Time when successful
func (m *AuditGroupAudit) GetReviewedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.reviewedAt
}
// GetReviewedBy gets the reviewedBy property value. The reviewedBy property
// returns a *string when successful
func (m *AuditGroupAudit) GetReviewedBy()(*string) {
    return m.reviewedBy
}
// GetStatus gets the status property value. The status property
// returns a *AuditGroupAuditStatus when successful
func (m *AuditGroupAudit) GetStatus()(*AuditGroupAuditStatus) {
    return m.status
}
// GetSubmittedAt gets the submittedAt property value. The submittedAt property
// returns a *Time when successful
func (m *AuditGroupAudit) GetSubmittedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.submittedAt
}
// GetSubmittedBy gets the submittedBy property value. The submittedBy property
// returns a *string when successful
func (m *AuditGroupAudit) GetSubmittedBy()(*string) {
    return m.submittedBy
}
// Serialize serializes information the current object
func (m *AuditGroupAudit) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetNestedGroups() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetNestedGroups()))
        for i, v := range m.GetNestedGroups() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("nestedGroups", cast)
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
func (m *AuditGroupAudit) SetAccounts(value []AuditGroupAuditAccountable)() {
    m.accounts = value
}
// SetAdditionalObjects sets the additionalObjects property value. The additionalObjects property
func (m *AuditGroupAudit) SetAdditionalObjects(value AuditGroupAudit_additionalObjectsable)() {
    m.additionalObjects = value
}
// SetComment sets the comment property value. The comment property
func (m *AuditGroupAudit) SetComment(value *string)() {
    m.comment = value
}
// SetCreatedAt sets the createdAt property value. The createdAt property
func (m *AuditGroupAudit) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdAt = value
}
// SetCreatedBy sets the createdBy property value. The createdBy property
func (m *AuditGroupAudit) SetCreatedBy(value *string)() {
    m.createdBy = value
}
// SetGroupName sets the groupName property value. The groupName property
func (m *AuditGroupAudit) SetGroupName(value *string)() {
    m.groupName = value
}
// SetNameOnAudit sets the nameOnAudit property value. The nameOnAudit property
func (m *AuditGroupAudit) SetNameOnAudit(value *string)() {
    m.nameOnAudit = value
}
// SetNestedGroups sets the nestedGroups property value. The nestedGroups property
func (m *AuditGroupAudit) SetNestedGroups(value []AuditNestedGroupAuditable)() {
    m.nestedGroups = value
}
// SetReviewedAt sets the reviewedAt property value. The reviewedAt property
func (m *AuditGroupAudit) SetReviewedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.reviewedAt = value
}
// SetReviewedBy sets the reviewedBy property value. The reviewedBy property
func (m *AuditGroupAudit) SetReviewedBy(value *string)() {
    m.reviewedBy = value
}
// SetStatus sets the status property value. The status property
func (m *AuditGroupAudit) SetStatus(value *AuditGroupAuditStatus)() {
    m.status = value
}
// SetSubmittedAt sets the submittedAt property value. The submittedAt property
func (m *AuditGroupAudit) SetSubmittedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.submittedAt = value
}
// SetSubmittedBy sets the submittedBy property value. The submittedBy property
func (m *AuditGroupAudit) SetSubmittedBy(value *string)() {
    m.submittedBy = value
}
type AuditGroupAuditable interface {
    Linkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccounts()([]AuditGroupAuditAccountable)
    GetAdditionalObjects()(AuditGroupAudit_additionalObjectsable)
    GetComment()(*string)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCreatedBy()(*string)
    GetGroupName()(*string)
    GetNameOnAudit()(*string)
    GetNestedGroups()([]AuditNestedGroupAuditable)
    GetReviewedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetReviewedBy()(*string)
    GetStatus()(*AuditGroupAuditStatus)
    GetSubmittedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSubmittedBy()(*string)
    SetAccounts(value []AuditGroupAuditAccountable)()
    SetAdditionalObjects(value AuditGroupAudit_additionalObjectsable)()
    SetComment(value *string)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCreatedBy(value *string)()
    SetGroupName(value *string)()
    SetNameOnAudit(value *string)()
    SetNestedGroups(value []AuditNestedGroupAuditable)()
    SetReviewedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetReviewedBy(value *string)()
    SetStatus(value *AuditGroupAuditStatus)()
    SetSubmittedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSubmittedBy(value *string)()
}
