package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WebhookWebhookPush 
type WebhookWebhookPush struct {
    NonLinkable
    // The account property
    account WebhookWebhookNameUuidable
    // The byParty property
    byParty WebhookWebhookNameUuidable
    // The certificate property
    certificate WebhookWebhookNameUuidable
    // The client property
    client WebhookWebhookNameUuidable
    // The directory property
    directory WebhookWebhookNameUuidable
    // The group property
    group WebhookWebhookNameUuidable
    // The group2 property
    group2 WebhookWebhookNameUuidable
    // The groupClassification property
    groupClassification WebhookWebhookNameUuidable
    // The modificationRequest property
    modificationRequest WebhookWebhookNameUuidable
    // The organizationalUnit property
    organizationalUnit WebhookWebhookNameUuidable
    // The parameter1 property
    parameter1 *string
    // The parameter2 property
    parameter2 *string
    // The parameter3 property
    parameter3 *string
    // The securityLevel property
    securityLevel *AuthSecurityLevel
    // The seq property
    seq *int64
    // The serviceAccount property
    serviceAccount WebhookWebhookNameUuidable
    // The system property
    system WebhookWebhookNameUuidable
    // The timestamp property
    timestamp *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The vaultRecord property
    vaultRecord WebhookWebhookNameUuidable
    // The webhook property
    webhook WebhookWebhookNameUuidable
    // The type property
    webhookWebhookPushType *AuditAuditRecordType
}
// NewWebhookWebhookPush instantiates a new webhookWebhookPush and sets the default values.
func NewWebhookWebhookPush()(*WebhookWebhookPush) {
    m := &WebhookWebhookPush{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "webhook.WebhookPush"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateWebhookWebhookPushFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWebhookWebhookPushFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWebhookWebhookPush(), nil
}
// GetAccount gets the account property value. The account property
func (m *WebhookWebhookPush) GetAccount()(WebhookWebhookNameUuidable) {
    return m.account
}
// GetByParty gets the byParty property value. The byParty property
func (m *WebhookWebhookPush) GetByParty()(WebhookWebhookNameUuidable) {
    return m.byParty
}
// GetCertificate gets the certificate property value. The certificate property
func (m *WebhookWebhookPush) GetCertificate()(WebhookWebhookNameUuidable) {
    return m.certificate
}
// GetClient gets the client property value. The client property
func (m *WebhookWebhookPush) GetClient()(WebhookWebhookNameUuidable) {
    return m.client
}
// GetDirectory gets the directory property value. The directory property
func (m *WebhookWebhookPush) GetDirectory()(WebhookWebhookNameUuidable) {
    return m.directory
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WebhookWebhookPush) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
    res["account"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWebhookWebhookNameUuidFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccount(val.(WebhookWebhookNameUuidable))
        }
        return nil
    }
    res["byParty"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWebhookWebhookNameUuidFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetByParty(val.(WebhookWebhookNameUuidable))
        }
        return nil
    }
    res["certificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWebhookWebhookNameUuidFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificate(val.(WebhookWebhookNameUuidable))
        }
        return nil
    }
    res["client"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWebhookWebhookNameUuidFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClient(val.(WebhookWebhookNameUuidable))
        }
        return nil
    }
    res["directory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWebhookWebhookNameUuidFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDirectory(val.(WebhookWebhookNameUuidable))
        }
        return nil
    }
    res["group"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWebhookWebhookNameUuidFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroup(val.(WebhookWebhookNameUuidable))
        }
        return nil
    }
    res["group2"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWebhookWebhookNameUuidFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroup2(val.(WebhookWebhookNameUuidable))
        }
        return nil
    }
    res["groupClassification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWebhookWebhookNameUuidFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupClassification(val.(WebhookWebhookNameUuidable))
        }
        return nil
    }
    res["modificationRequest"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWebhookWebhookNameUuidFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModificationRequest(val.(WebhookWebhookNameUuidable))
        }
        return nil
    }
    res["organizationalUnit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWebhookWebhookNameUuidFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationalUnit(val.(WebhookWebhookNameUuidable))
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
    res["seq"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSeq(val)
        }
        return nil
    }
    res["serviceAccount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWebhookWebhookNameUuidFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServiceAccount(val.(WebhookWebhookNameUuidable))
        }
        return nil
    }
    res["system"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWebhookWebhookNameUuidFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystem(val.(WebhookWebhookNameUuidable))
        }
        return nil
    }
    res["timestamp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimestamp(val)
        }
        return nil
    }
    res["vaultRecord"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWebhookWebhookNameUuidFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVaultRecord(val.(WebhookWebhookNameUuidable))
        }
        return nil
    }
    res["webhook"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWebhookWebhookNameUuidFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebhook(val.(WebhookWebhookNameUuidable))
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAuditAuditRecordType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebhookWebhookPushType(val.(*AuditAuditRecordType))
        }
        return nil
    }
    return res
}
// GetGroup gets the group property value. The group property
func (m *WebhookWebhookPush) GetGroup()(WebhookWebhookNameUuidable) {
    return m.group
}
// GetGroup2 gets the group2 property value. The group2 property
func (m *WebhookWebhookPush) GetGroup2()(WebhookWebhookNameUuidable) {
    return m.group2
}
// GetGroupClassification gets the groupClassification property value. The groupClassification property
func (m *WebhookWebhookPush) GetGroupClassification()(WebhookWebhookNameUuidable) {
    return m.groupClassification
}
// GetModificationRequest gets the modificationRequest property value. The modificationRequest property
func (m *WebhookWebhookPush) GetModificationRequest()(WebhookWebhookNameUuidable) {
    return m.modificationRequest
}
// GetOrganizationalUnit gets the organizationalUnit property value. The organizationalUnit property
func (m *WebhookWebhookPush) GetOrganizationalUnit()(WebhookWebhookNameUuidable) {
    return m.organizationalUnit
}
// GetParameter1 gets the parameter1 property value. The parameter1 property
func (m *WebhookWebhookPush) GetParameter1()(*string) {
    return m.parameter1
}
// GetParameter2 gets the parameter2 property value. The parameter2 property
func (m *WebhookWebhookPush) GetParameter2()(*string) {
    return m.parameter2
}
// GetParameter3 gets the parameter3 property value. The parameter3 property
func (m *WebhookWebhookPush) GetParameter3()(*string) {
    return m.parameter3
}
// GetSecurityLevel gets the securityLevel property value. The securityLevel property
func (m *WebhookWebhookPush) GetSecurityLevel()(*AuthSecurityLevel) {
    return m.securityLevel
}
// GetSeq gets the seq property value. The seq property
func (m *WebhookWebhookPush) GetSeq()(*int64) {
    return m.seq
}
// GetServiceAccount gets the serviceAccount property value. The serviceAccount property
func (m *WebhookWebhookPush) GetServiceAccount()(WebhookWebhookNameUuidable) {
    return m.serviceAccount
}
// GetSystem gets the system property value. The system property
func (m *WebhookWebhookPush) GetSystem()(WebhookWebhookNameUuidable) {
    return m.system
}
// GetTimestamp gets the timestamp property value. The timestamp property
func (m *WebhookWebhookPush) GetTimestamp()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.timestamp
}
// GetVaultRecord gets the vaultRecord property value. The vaultRecord property
func (m *WebhookWebhookPush) GetVaultRecord()(WebhookWebhookNameUuidable) {
    return m.vaultRecord
}
// GetWebhook gets the webhook property value. The webhook property
func (m *WebhookWebhookPush) GetWebhook()(WebhookWebhookNameUuidable) {
    return m.webhook
}
// GetWebhookWebhookPushType gets the type property value. The type property
func (m *WebhookWebhookPush) GetWebhookWebhookPushType()(*AuditAuditRecordType) {
    return m.webhookWebhookPushType
}
// Serialize serializes information the current object
func (m *WebhookWebhookPush) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("account", m.GetAccount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("byParty", m.GetByParty())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("certificate", m.GetCertificate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("client", m.GetClient())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("directory", m.GetDirectory())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("group", m.GetGroup())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("group2", m.GetGroup2())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("groupClassification", m.GetGroupClassification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("modificationRequest", m.GetModificationRequest())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("organizationalUnit", m.GetOrganizationalUnit())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("parameter1", m.GetParameter1())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("parameter2", m.GetParameter2())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("parameter3", m.GetParameter3())
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
    {
        err = writer.WriteInt64Value("seq", m.GetSeq())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("serviceAccount", m.GetServiceAccount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("system", m.GetSystem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("timestamp", m.GetTimestamp())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("vaultRecord", m.GetVaultRecord())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("webhook", m.GetWebhook())
        if err != nil {
            return err
        }
    }
    if m.GetWebhookWebhookPushType() != nil {
        cast := (*m.GetWebhookWebhookPushType()).String()
        err = writer.WriteStringValue("type", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccount sets the account property value. The account property
func (m *WebhookWebhookPush) SetAccount(value WebhookWebhookNameUuidable)() {
    m.account = value
}
// SetByParty sets the byParty property value. The byParty property
func (m *WebhookWebhookPush) SetByParty(value WebhookWebhookNameUuidable)() {
    m.byParty = value
}
// SetCertificate sets the certificate property value. The certificate property
func (m *WebhookWebhookPush) SetCertificate(value WebhookWebhookNameUuidable)() {
    m.certificate = value
}
// SetClient sets the client property value. The client property
func (m *WebhookWebhookPush) SetClient(value WebhookWebhookNameUuidable)() {
    m.client = value
}
// SetDirectory sets the directory property value. The directory property
func (m *WebhookWebhookPush) SetDirectory(value WebhookWebhookNameUuidable)() {
    m.directory = value
}
// SetGroup sets the group property value. The group property
func (m *WebhookWebhookPush) SetGroup(value WebhookWebhookNameUuidable)() {
    m.group = value
}
// SetGroup2 sets the group2 property value. The group2 property
func (m *WebhookWebhookPush) SetGroup2(value WebhookWebhookNameUuidable)() {
    m.group2 = value
}
// SetGroupClassification sets the groupClassification property value. The groupClassification property
func (m *WebhookWebhookPush) SetGroupClassification(value WebhookWebhookNameUuidable)() {
    m.groupClassification = value
}
// SetModificationRequest sets the modificationRequest property value. The modificationRequest property
func (m *WebhookWebhookPush) SetModificationRequest(value WebhookWebhookNameUuidable)() {
    m.modificationRequest = value
}
// SetOrganizationalUnit sets the organizationalUnit property value. The organizationalUnit property
func (m *WebhookWebhookPush) SetOrganizationalUnit(value WebhookWebhookNameUuidable)() {
    m.organizationalUnit = value
}
// SetParameter1 sets the parameter1 property value. The parameter1 property
func (m *WebhookWebhookPush) SetParameter1(value *string)() {
    m.parameter1 = value
}
// SetParameter2 sets the parameter2 property value. The parameter2 property
func (m *WebhookWebhookPush) SetParameter2(value *string)() {
    m.parameter2 = value
}
// SetParameter3 sets the parameter3 property value. The parameter3 property
func (m *WebhookWebhookPush) SetParameter3(value *string)() {
    m.parameter3 = value
}
// SetSecurityLevel sets the securityLevel property value. The securityLevel property
func (m *WebhookWebhookPush) SetSecurityLevel(value *AuthSecurityLevel)() {
    m.securityLevel = value
}
// SetSeq sets the seq property value. The seq property
func (m *WebhookWebhookPush) SetSeq(value *int64)() {
    m.seq = value
}
// SetServiceAccount sets the serviceAccount property value. The serviceAccount property
func (m *WebhookWebhookPush) SetServiceAccount(value WebhookWebhookNameUuidable)() {
    m.serviceAccount = value
}
// SetSystem sets the system property value. The system property
func (m *WebhookWebhookPush) SetSystem(value WebhookWebhookNameUuidable)() {
    m.system = value
}
// SetTimestamp sets the timestamp property value. The timestamp property
func (m *WebhookWebhookPush) SetTimestamp(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.timestamp = value
}
// SetVaultRecord sets the vaultRecord property value. The vaultRecord property
func (m *WebhookWebhookPush) SetVaultRecord(value WebhookWebhookNameUuidable)() {
    m.vaultRecord = value
}
// SetWebhook sets the webhook property value. The webhook property
func (m *WebhookWebhookPush) SetWebhook(value WebhookWebhookNameUuidable)() {
    m.webhook = value
}
// SetWebhookWebhookPushType sets the type property value. The type property
func (m *WebhookWebhookPush) SetWebhookWebhookPushType(value *AuditAuditRecordType)() {
    m.webhookWebhookPushType = value
}
// WebhookWebhookPushable 
type WebhookWebhookPushable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccount()(WebhookWebhookNameUuidable)
    GetByParty()(WebhookWebhookNameUuidable)
    GetCertificate()(WebhookWebhookNameUuidable)
    GetClient()(WebhookWebhookNameUuidable)
    GetDirectory()(WebhookWebhookNameUuidable)
    GetGroup()(WebhookWebhookNameUuidable)
    GetGroup2()(WebhookWebhookNameUuidable)
    GetGroupClassification()(WebhookWebhookNameUuidable)
    GetModificationRequest()(WebhookWebhookNameUuidable)
    GetOrganizationalUnit()(WebhookWebhookNameUuidable)
    GetParameter1()(*string)
    GetParameter2()(*string)
    GetParameter3()(*string)
    GetSecurityLevel()(*AuthSecurityLevel)
    GetSeq()(*int64)
    GetServiceAccount()(WebhookWebhookNameUuidable)
    GetSystem()(WebhookWebhookNameUuidable)
    GetTimestamp()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetVaultRecord()(WebhookWebhookNameUuidable)
    GetWebhook()(WebhookWebhookNameUuidable)
    GetWebhookWebhookPushType()(*AuditAuditRecordType)
    SetAccount(value WebhookWebhookNameUuidable)()
    SetByParty(value WebhookWebhookNameUuidable)()
    SetCertificate(value WebhookWebhookNameUuidable)()
    SetClient(value WebhookWebhookNameUuidable)()
    SetDirectory(value WebhookWebhookNameUuidable)()
    SetGroup(value WebhookWebhookNameUuidable)()
    SetGroup2(value WebhookWebhookNameUuidable)()
    SetGroupClassification(value WebhookWebhookNameUuidable)()
    SetModificationRequest(value WebhookWebhookNameUuidable)()
    SetOrganizationalUnit(value WebhookWebhookNameUuidable)()
    SetParameter1(value *string)()
    SetParameter2(value *string)()
    SetParameter3(value *string)()
    SetSecurityLevel(value *AuthSecurityLevel)()
    SetSeq(value *int64)()
    SetServiceAccount(value WebhookWebhookNameUuidable)()
    SetSystem(value WebhookWebhookNameUuidable)()
    SetTimestamp(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetVaultRecord(value WebhookWebhookNameUuidable)()
    SetWebhook(value WebhookWebhookNameUuidable)()
    SetWebhookWebhookPushType(value *AuditAuditRecordType)()
}
