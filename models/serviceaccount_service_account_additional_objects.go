package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ServiceaccountServiceAccount_additionalObjects struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The audit property
    audit AuditInfoable
    // The groups property
    groups ServiceaccountServiceAccountGroupLinkableWrapperable
    // The secret property
    secret GeneratedSecretable
}
// NewServiceaccountServiceAccount_additionalObjects instantiates a new ServiceaccountServiceAccount_additionalObjects and sets the default values.
func NewServiceaccountServiceAccount_additionalObjects()(*ServiceaccountServiceAccount_additionalObjects) {
    m := &ServiceaccountServiceAccount_additionalObjects{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateServiceaccountServiceAccount_additionalObjectsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateServiceaccountServiceAccount_additionalObjectsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServiceaccountServiceAccount_additionalObjects(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ServiceaccountServiceAccount_additionalObjects) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAudit gets the audit property value. The audit property
// returns a AuditInfoable when successful
func (m *ServiceaccountServiceAccount_additionalObjects) GetAudit()(AuditInfoable) {
    return m.audit
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ServiceaccountServiceAccount_additionalObjects) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["groups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateServiceaccountServiceAccountGroupLinkableWrapperFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroups(val.(ServiceaccountServiceAccountGroupLinkableWrapperable))
        }
        return nil
    }
    res["secret"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGeneratedSecretFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecret(val.(GeneratedSecretable))
        }
        return nil
    }
    return res
}
// GetGroups gets the groups property value. The groups property
// returns a ServiceaccountServiceAccountGroupLinkableWrapperable when successful
func (m *ServiceaccountServiceAccount_additionalObjects) GetGroups()(ServiceaccountServiceAccountGroupLinkableWrapperable) {
    return m.groups
}
// GetSecret gets the secret property value. The secret property
// returns a GeneratedSecretable when successful
func (m *ServiceaccountServiceAccount_additionalObjects) GetSecret()(GeneratedSecretable) {
    return m.secret
}
// Serialize serializes information the current object
func (m *ServiceaccountServiceAccount_additionalObjects) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("audit", m.GetAudit())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("groups", m.GetGroups())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("secret", m.GetSecret())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ServiceaccountServiceAccount_additionalObjects) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAudit sets the audit property value. The audit property
func (m *ServiceaccountServiceAccount_additionalObjects) SetAudit(value AuditInfoable)() {
    m.audit = value
}
// SetGroups sets the groups property value. The groups property
func (m *ServiceaccountServiceAccount_additionalObjects) SetGroups(value ServiceaccountServiceAccountGroupLinkableWrapperable)() {
    m.groups = value
}
// SetSecret sets the secret property value. The secret property
func (m *ServiceaccountServiceAccount_additionalObjects) SetSecret(value GeneratedSecretable)() {
    m.secret = value
}
type ServiceaccountServiceAccount_additionalObjectsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAudit()(AuditInfoable)
    GetGroups()(ServiceaccountServiceAccountGroupLinkableWrapperable)
    GetSecret()(GeneratedSecretable)
    SetAudit(value AuditInfoable)()
    SetGroups(value ServiceaccountServiceAccountGroupLinkableWrapperable)()
    SetSecret(value GeneratedSecretable)()
}
