package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AuditAuditRecord struct {
    Linkable
    // The additionalObjects property
    additionalObjects AuditAuditRecord_additionalObjectsable
    // The type property
    auditAuditRecordType *AuditAuditRecordType
    // The dateTime property
    dateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The onAccount property
    onAccount *string
    // The onCertificate property
    onCertificate *string
    // The onClient property
    onClient *string
    // The onDirectory property
    onDirectory *string
    // The onGroup property
    onGroup *string
    // The onGroup2 property
    onGroup2 *string
    // The onGroupClassification property
    onGroupClassification *string
    // The onOrganizationalUnit property
    onOrganizationalUnit *string
    // The onServiceAccount property
    onServiceAccount *string
    // The onSystem property
    onSystem *string
    // The onVaultRecord property
    onVaultRecord *string
    // The onWebhook property
    onWebhook *string
    // The parameter1 property
    parameter1 *string
    // The parameter2 property
    parameter2 *string
    // The parameter3 property
    parameter3 *string
    // The performedBy property
    performedBy *string
    // The securityLevel property
    securityLevel *AuthSecurityLevel
}
// NewAuditAuditRecord instantiates a new AuditAuditRecord and sets the default values.
func NewAuditAuditRecord()(*AuditAuditRecord) {
    m := &AuditAuditRecord{
        Linkable: *NewLinkable(),
    }
    typeEscapedValue := "audit.AuditRecord"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateAuditAuditRecordFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAuditAuditRecordFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuditAuditRecord(), nil
}
// GetAdditionalObjects gets the additionalObjects property value. The additionalObjects property
// returns a AuditAuditRecord_additionalObjectsable when successful
func (m *AuditAuditRecord) GetAdditionalObjects()(AuditAuditRecord_additionalObjectsable) {
    return m.additionalObjects
}
// GetAuditAuditRecordType gets the type property value. The type property
// returns a *AuditAuditRecordType when successful
func (m *AuditAuditRecord) GetAuditAuditRecordType()(*AuditAuditRecordType) {
    return m.auditAuditRecordType
}
// GetDateTime gets the dateTime property value. The dateTime property
// returns a *Time when successful
func (m *AuditAuditRecord) GetDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.dateTime
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AuditAuditRecord) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Linkable.GetFieldDeserializers()
    res["additionalObjects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuditAuditRecord_additionalObjectsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalObjects(val.(AuditAuditRecord_additionalObjectsable))
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAuditAuditRecordType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuditAuditRecordType(val.(*AuditAuditRecordType))
        }
        return nil
    }
    res["dateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDateTime(val)
        }
        return nil
    }
    res["onAccount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnAccount(val)
        }
        return nil
    }
    res["onCertificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnCertificate(val)
        }
        return nil
    }
    res["onClient"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnClient(val)
        }
        return nil
    }
    res["onDirectory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnDirectory(val)
        }
        return nil
    }
    res["onGroup"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnGroup(val)
        }
        return nil
    }
    res["onGroup2"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnGroup2(val)
        }
        return nil
    }
    res["onGroupClassification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnGroupClassification(val)
        }
        return nil
    }
    res["onOrganizationalUnit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnOrganizationalUnit(val)
        }
        return nil
    }
    res["onServiceAccount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnServiceAccount(val)
        }
        return nil
    }
    res["onSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnSystem(val)
        }
        return nil
    }
    res["onVaultRecord"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnVaultRecord(val)
        }
        return nil
    }
    res["onWebhook"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnWebhook(val)
        }
        return nil
    }
    res["parameter1"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParameter1(val)
        }
        return nil
    }
    res["parameter2"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParameter2(val)
        }
        return nil
    }
    res["parameter3"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParameter3(val)
        }
        return nil
    }
    res["performedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPerformedBy(val)
        }
        return nil
    }
    res["securityLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAuthSecurityLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityLevel(val.(*AuthSecurityLevel))
        }
        return nil
    }
    return res
}
// GetOnAccount gets the onAccount property value. The onAccount property
// returns a *string when successful
func (m *AuditAuditRecord) GetOnAccount()(*string) {
    return m.onAccount
}
// GetOnCertificate gets the onCertificate property value. The onCertificate property
// returns a *string when successful
func (m *AuditAuditRecord) GetOnCertificate()(*string) {
    return m.onCertificate
}
// GetOnClient gets the onClient property value. The onClient property
// returns a *string when successful
func (m *AuditAuditRecord) GetOnClient()(*string) {
    return m.onClient
}
// GetOnDirectory gets the onDirectory property value. The onDirectory property
// returns a *string when successful
func (m *AuditAuditRecord) GetOnDirectory()(*string) {
    return m.onDirectory
}
// GetOnGroup gets the onGroup property value. The onGroup property
// returns a *string when successful
func (m *AuditAuditRecord) GetOnGroup()(*string) {
    return m.onGroup
}
// GetOnGroup2 gets the onGroup2 property value. The onGroup2 property
// returns a *string when successful
func (m *AuditAuditRecord) GetOnGroup2()(*string) {
    return m.onGroup2
}
// GetOnGroupClassification gets the onGroupClassification property value. The onGroupClassification property
// returns a *string when successful
func (m *AuditAuditRecord) GetOnGroupClassification()(*string) {
    return m.onGroupClassification
}
// GetOnOrganizationalUnit gets the onOrganizationalUnit property value. The onOrganizationalUnit property
// returns a *string when successful
func (m *AuditAuditRecord) GetOnOrganizationalUnit()(*string) {
    return m.onOrganizationalUnit
}
// GetOnServiceAccount gets the onServiceAccount property value. The onServiceAccount property
// returns a *string when successful
func (m *AuditAuditRecord) GetOnServiceAccount()(*string) {
    return m.onServiceAccount
}
// GetOnSystem gets the onSystem property value. The onSystem property
// returns a *string when successful
func (m *AuditAuditRecord) GetOnSystem()(*string) {
    return m.onSystem
}
// GetOnVaultRecord gets the onVaultRecord property value. The onVaultRecord property
// returns a *string when successful
func (m *AuditAuditRecord) GetOnVaultRecord()(*string) {
    return m.onVaultRecord
}
// GetOnWebhook gets the onWebhook property value. The onWebhook property
// returns a *string when successful
func (m *AuditAuditRecord) GetOnWebhook()(*string) {
    return m.onWebhook
}
// GetParameter1 gets the parameter1 property value. The parameter1 property
// returns a *string when successful
func (m *AuditAuditRecord) GetParameter1()(*string) {
    return m.parameter1
}
// GetParameter2 gets the parameter2 property value. The parameter2 property
// returns a *string when successful
func (m *AuditAuditRecord) GetParameter2()(*string) {
    return m.parameter2
}
// GetParameter3 gets the parameter3 property value. The parameter3 property
// returns a *string when successful
func (m *AuditAuditRecord) GetParameter3()(*string) {
    return m.parameter3
}
// GetPerformedBy gets the performedBy property value. The performedBy property
// returns a *string when successful
func (m *AuditAuditRecord) GetPerformedBy()(*string) {
    return m.performedBy
}
// GetSecurityLevel gets the securityLevel property value. The securityLevel property
// returns a *AuthSecurityLevel when successful
func (m *AuditAuditRecord) GetSecurityLevel()(*AuthSecurityLevel) {
    return m.securityLevel
}
// Serialize serializes information the current object
func (m *AuditAuditRecord) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Linkable.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("additionalObjects", m.GetAdditionalObjects())
        if err != nil {
            return err
        }
    }
    if m.GetAuditAuditRecordType() != nil {
        cast := (*m.GetAuditAuditRecordType()).String()
        err = writer.WriteStringValue("type", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSecurityLevel() != nil {
        cast := (*m.GetSecurityLevel()).String()
        err = writer.WriteStringValue("securityLevel", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalObjects sets the additionalObjects property value. The additionalObjects property
func (m *AuditAuditRecord) SetAdditionalObjects(value AuditAuditRecord_additionalObjectsable)() {
    m.additionalObjects = value
}
// SetAuditAuditRecordType sets the type property value. The type property
func (m *AuditAuditRecord) SetAuditAuditRecordType(value *AuditAuditRecordType)() {
    m.auditAuditRecordType = value
}
// SetDateTime sets the dateTime property value. The dateTime property
func (m *AuditAuditRecord) SetDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.dateTime = value
}
// SetOnAccount sets the onAccount property value. The onAccount property
func (m *AuditAuditRecord) SetOnAccount(value *string)() {
    m.onAccount = value
}
// SetOnCertificate sets the onCertificate property value. The onCertificate property
func (m *AuditAuditRecord) SetOnCertificate(value *string)() {
    m.onCertificate = value
}
// SetOnClient sets the onClient property value. The onClient property
func (m *AuditAuditRecord) SetOnClient(value *string)() {
    m.onClient = value
}
// SetOnDirectory sets the onDirectory property value. The onDirectory property
func (m *AuditAuditRecord) SetOnDirectory(value *string)() {
    m.onDirectory = value
}
// SetOnGroup sets the onGroup property value. The onGroup property
func (m *AuditAuditRecord) SetOnGroup(value *string)() {
    m.onGroup = value
}
// SetOnGroup2 sets the onGroup2 property value. The onGroup2 property
func (m *AuditAuditRecord) SetOnGroup2(value *string)() {
    m.onGroup2 = value
}
// SetOnGroupClassification sets the onGroupClassification property value. The onGroupClassification property
func (m *AuditAuditRecord) SetOnGroupClassification(value *string)() {
    m.onGroupClassification = value
}
// SetOnOrganizationalUnit sets the onOrganizationalUnit property value. The onOrganizationalUnit property
func (m *AuditAuditRecord) SetOnOrganizationalUnit(value *string)() {
    m.onOrganizationalUnit = value
}
// SetOnServiceAccount sets the onServiceAccount property value. The onServiceAccount property
func (m *AuditAuditRecord) SetOnServiceAccount(value *string)() {
    m.onServiceAccount = value
}
// SetOnSystem sets the onSystem property value. The onSystem property
func (m *AuditAuditRecord) SetOnSystem(value *string)() {
    m.onSystem = value
}
// SetOnVaultRecord sets the onVaultRecord property value. The onVaultRecord property
func (m *AuditAuditRecord) SetOnVaultRecord(value *string)() {
    m.onVaultRecord = value
}
// SetOnWebhook sets the onWebhook property value. The onWebhook property
func (m *AuditAuditRecord) SetOnWebhook(value *string)() {
    m.onWebhook = value
}
// SetParameter1 sets the parameter1 property value. The parameter1 property
func (m *AuditAuditRecord) SetParameter1(value *string)() {
    m.parameter1 = value
}
// SetParameter2 sets the parameter2 property value. The parameter2 property
func (m *AuditAuditRecord) SetParameter2(value *string)() {
    m.parameter2 = value
}
// SetParameter3 sets the parameter3 property value. The parameter3 property
func (m *AuditAuditRecord) SetParameter3(value *string)() {
    m.parameter3 = value
}
// SetPerformedBy sets the performedBy property value. The performedBy property
func (m *AuditAuditRecord) SetPerformedBy(value *string)() {
    m.performedBy = value
}
// SetSecurityLevel sets the securityLevel property value. The securityLevel property
func (m *AuditAuditRecord) SetSecurityLevel(value *AuthSecurityLevel)() {
    m.securityLevel = value
}
type AuditAuditRecordable interface {
    Linkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdditionalObjects()(AuditAuditRecord_additionalObjectsable)
    GetAuditAuditRecordType()(*AuditAuditRecordType)
    GetDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOnAccount()(*string)
    GetOnCertificate()(*string)
    GetOnClient()(*string)
    GetOnDirectory()(*string)
    GetOnGroup()(*string)
    GetOnGroup2()(*string)
    GetOnGroupClassification()(*string)
    GetOnOrganizationalUnit()(*string)
    GetOnServiceAccount()(*string)
    GetOnSystem()(*string)
    GetOnVaultRecord()(*string)
    GetOnWebhook()(*string)
    GetParameter1()(*string)
    GetParameter2()(*string)
    GetParameter3()(*string)
    GetPerformedBy()(*string)
    GetSecurityLevel()(*AuthSecurityLevel)
    SetAdditionalObjects(value AuditAuditRecord_additionalObjectsable)()
    SetAuditAuditRecordType(value *AuditAuditRecordType)()
    SetDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOnAccount(value *string)()
    SetOnCertificate(value *string)()
    SetOnClient(value *string)()
    SetOnDirectory(value *string)()
    SetOnGroup(value *string)()
    SetOnGroup2(value *string)()
    SetOnGroupClassification(value *string)()
    SetOnOrganizationalUnit(value *string)()
    SetOnServiceAccount(value *string)()
    SetOnSystem(value *string)()
    SetOnVaultRecord(value *string)()
    SetOnWebhook(value *string)()
    SetParameter1(value *string)()
    SetParameter2(value *string)()
    SetParameter3(value *string)()
    SetPerformedBy(value *string)()
    SetSecurityLevel(value *AuthSecurityLevel)()
}
