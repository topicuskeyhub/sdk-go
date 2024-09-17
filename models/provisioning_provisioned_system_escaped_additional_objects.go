package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ProvisioningProvisionedSystem_additionalObjects struct {
    // The account property
    account ProvisioningProvisionedAccountable
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The audit property
    audit AuditInfoable
    // The issuedPermissions property
    issuedPermissions ClientOAuth2ClientPermissionWithClientLinkableWrapperable
    // The loginName property
    loginName *string
    // The managementPermissions property
    managementPermissions ProvisioningProvisioningManagementPermissionsable
    // The markers property
    markers MarkItemMarkersable
    // The statistics property
    statistics ProvisioningCircuitBreakerStatisticsable
    // The supportedGroupTypes property
    supportedGroupTypes ProvisioningGroupOnSystemTypesable
}
// NewProvisioningProvisionedSystem_additionalObjects instantiates a new ProvisioningProvisionedSystem_additionalObjects and sets the default values.
func NewProvisioningProvisionedSystem_additionalObjects()(*ProvisioningProvisionedSystem_additionalObjects) {
    m := &ProvisioningProvisionedSystem_additionalObjects{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateProvisioningProvisionedSystem_additionalObjectsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateProvisioningProvisionedSystem_additionalObjectsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProvisioningProvisionedSystem_additionalObjects(), nil
}
// GetAccount gets the account property value. The account property
// returns a ProvisioningProvisionedAccountable when successful
func (m *ProvisioningProvisionedSystem_additionalObjects) GetAccount()(ProvisioningProvisionedAccountable) {
    return m.account
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ProvisioningProvisionedSystem_additionalObjects) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAudit gets the audit property value. The audit property
// returns a AuditInfoable when successful
func (m *ProvisioningProvisionedSystem_additionalObjects) GetAudit()(AuditInfoable) {
    return m.audit
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ProvisioningProvisionedSystem_additionalObjects) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["account"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProvisioningProvisionedAccountFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccount(val.(ProvisioningProvisionedAccountable))
        }
        return nil
    }
    res["audit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuditInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAudit(val.(AuditInfoable))
        }
        return nil
    }
    res["issuedPermissions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateClientOAuth2ClientPermissionWithClientLinkableWrapperFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssuedPermissions(val.(ClientOAuth2ClientPermissionWithClientLinkableWrapperable))
        }
        return nil
    }
    res["loginName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLoginName(val)
        }
        return nil
    }
    res["managementPermissions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProvisioningProvisioningManagementPermissionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementPermissions(val.(ProvisioningProvisioningManagementPermissionsable))
        }
        return nil
    }
    res["markers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMarkItemMarkersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMarkers(val.(MarkItemMarkersable))
        }
        return nil
    }
    res["statistics"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProvisioningCircuitBreakerStatisticsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatistics(val.(ProvisioningCircuitBreakerStatisticsable))
        }
        return nil
    }
    res["supportedGroupTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProvisioningGroupOnSystemTypesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSupportedGroupTypes(val.(ProvisioningGroupOnSystemTypesable))
        }
        return nil
    }
    return res
}
// GetIssuedPermissions gets the issuedPermissions property value. The issuedPermissions property
// returns a ClientOAuth2ClientPermissionWithClientLinkableWrapperable when successful
func (m *ProvisioningProvisionedSystem_additionalObjects) GetIssuedPermissions()(ClientOAuth2ClientPermissionWithClientLinkableWrapperable) {
    return m.issuedPermissions
}
// GetLoginName gets the loginName property value. The loginName property
// returns a *string when successful
func (m *ProvisioningProvisionedSystem_additionalObjects) GetLoginName()(*string) {
    return m.loginName
}
// GetManagementPermissions gets the managementPermissions property value. The managementPermissions property
// returns a ProvisioningProvisioningManagementPermissionsable when successful
func (m *ProvisioningProvisionedSystem_additionalObjects) GetManagementPermissions()(ProvisioningProvisioningManagementPermissionsable) {
    return m.managementPermissions
}
// GetMarkers gets the markers property value. The markers property
// returns a MarkItemMarkersable when successful
func (m *ProvisioningProvisionedSystem_additionalObjects) GetMarkers()(MarkItemMarkersable) {
    return m.markers
}
// GetStatistics gets the statistics property value. The statistics property
// returns a ProvisioningCircuitBreakerStatisticsable when successful
func (m *ProvisioningProvisionedSystem_additionalObjects) GetStatistics()(ProvisioningCircuitBreakerStatisticsable) {
    return m.statistics
}
// GetSupportedGroupTypes gets the supportedGroupTypes property value. The supportedGroupTypes property
// returns a ProvisioningGroupOnSystemTypesable when successful
func (m *ProvisioningProvisionedSystem_additionalObjects) GetSupportedGroupTypes()(ProvisioningGroupOnSystemTypesable) {
    return m.supportedGroupTypes
}
// Serialize serializes information the current object
func (m *ProvisioningProvisionedSystem_additionalObjects) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("account", m.GetAccount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("audit", m.GetAudit())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("issuedPermissions", m.GetIssuedPermissions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("managementPermissions", m.GetManagementPermissions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("markers", m.GetMarkers())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("statistics", m.GetStatistics())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("supportedGroupTypes", m.GetSupportedGroupTypes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccount sets the account property value. The account property
func (m *ProvisioningProvisionedSystem_additionalObjects) SetAccount(value ProvisioningProvisionedAccountable)() {
    m.account = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ProvisioningProvisionedSystem_additionalObjects) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAudit sets the audit property value. The audit property
func (m *ProvisioningProvisionedSystem_additionalObjects) SetAudit(value AuditInfoable)() {
    m.audit = value
}
// SetIssuedPermissions sets the issuedPermissions property value. The issuedPermissions property
func (m *ProvisioningProvisionedSystem_additionalObjects) SetIssuedPermissions(value ClientOAuth2ClientPermissionWithClientLinkableWrapperable)() {
    m.issuedPermissions = value
}
// SetLoginName sets the loginName property value. The loginName property
func (m *ProvisioningProvisionedSystem_additionalObjects) SetLoginName(value *string)() {
    m.loginName = value
}
// SetManagementPermissions sets the managementPermissions property value. The managementPermissions property
func (m *ProvisioningProvisionedSystem_additionalObjects) SetManagementPermissions(value ProvisioningProvisioningManagementPermissionsable)() {
    m.managementPermissions = value
}
// SetMarkers sets the markers property value. The markers property
func (m *ProvisioningProvisionedSystem_additionalObjects) SetMarkers(value MarkItemMarkersable)() {
    m.markers = value
}
// SetStatistics sets the statistics property value. The statistics property
func (m *ProvisioningProvisionedSystem_additionalObjects) SetStatistics(value ProvisioningCircuitBreakerStatisticsable)() {
    m.statistics = value
}
// SetSupportedGroupTypes sets the supportedGroupTypes property value. The supportedGroupTypes property
func (m *ProvisioningProvisionedSystem_additionalObjects) SetSupportedGroupTypes(value ProvisioningGroupOnSystemTypesable)() {
    m.supportedGroupTypes = value
}
type ProvisioningProvisionedSystem_additionalObjectsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccount()(ProvisioningProvisionedAccountable)
    GetAudit()(AuditInfoable)
    GetIssuedPermissions()(ClientOAuth2ClientPermissionWithClientLinkableWrapperable)
    GetLoginName()(*string)
    GetManagementPermissions()(ProvisioningProvisioningManagementPermissionsable)
    GetMarkers()(MarkItemMarkersable)
    GetStatistics()(ProvisioningCircuitBreakerStatisticsable)
    GetSupportedGroupTypes()(ProvisioningGroupOnSystemTypesable)
    SetAccount(value ProvisioningProvisionedAccountable)()
    SetAudit(value AuditInfoable)()
    SetIssuedPermissions(value ClientOAuth2ClientPermissionWithClientLinkableWrapperable)()
    SetLoginName(value *string)()
    SetManagementPermissions(value ProvisioningProvisioningManagementPermissionsable)()
    SetMarkers(value MarkItemMarkersable)()
    SetStatistics(value ProvisioningCircuitBreakerStatisticsable)()
    SetSupportedGroupTypes(value ProvisioningGroupOnSystemTypesable)()
}
