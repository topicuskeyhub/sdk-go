package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ProvisioningProvisionedSystemSyncLog struct {
    Linkable
    // The additionalObjects property
    additionalObjects ProvisioningProvisionedSystemSyncLog_additionalObjectsable
    // The createdAt property
    createdAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The errors property
    errors *int32
    // The log property
    log *string
    // The modifications property
    modifications *int32
}
// NewProvisioningProvisionedSystemSyncLog instantiates a new ProvisioningProvisionedSystemSyncLog and sets the default values.
func NewProvisioningProvisionedSystemSyncLog()(*ProvisioningProvisionedSystemSyncLog) {
    m := &ProvisioningProvisionedSystemSyncLog{
        Linkable: *NewLinkable(),
    }
    typeEscapedValue := "provisioning.ProvisionedSystemSyncLog"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateProvisioningProvisionedSystemSyncLogFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateProvisioningProvisionedSystemSyncLogFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProvisioningProvisionedSystemSyncLog(), nil
}
// GetAdditionalObjects gets the additionalObjects property value. The additionalObjects property
// returns a ProvisioningProvisionedSystemSyncLog_additionalObjectsable when successful
func (m *ProvisioningProvisionedSystemSyncLog) GetAdditionalObjects()(ProvisioningProvisionedSystemSyncLog_additionalObjectsable) {
    return m.additionalObjects
}
// GetCreatedAt gets the createdAt property value. The createdAt property
// returns a *Time when successful
func (m *ProvisioningProvisionedSystemSyncLog) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdAt
}
// GetErrors gets the errors property value. The errors property
// returns a *int32 when successful
func (m *ProvisioningProvisionedSystemSyncLog) GetErrors()(*int32) {
    return m.errors
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ProvisioningProvisionedSystemSyncLog) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Linkable.GetFieldDeserializers()
    res["additionalObjects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProvisioningProvisionedSystemSyncLog_additionalObjectsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalObjects(val.(ProvisioningProvisionedSystemSyncLog_additionalObjectsable))
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
    res["errors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrors(val)
        }
        return nil
    }
    res["log"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLog(val)
        }
        return nil
    }
    res["modifications"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModifications(val)
        }
        return nil
    }
    return res
}
// GetLog gets the log property value. The log property
// returns a *string when successful
func (m *ProvisioningProvisionedSystemSyncLog) GetLog()(*string) {
    return m.log
}
// GetModifications gets the modifications property value. The modifications property
// returns a *int32 when successful
func (m *ProvisioningProvisionedSystemSyncLog) GetModifications()(*int32) {
    return m.modifications
}
// Serialize serializes information the current object
func (m *ProvisioningProvisionedSystemSyncLog) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    return nil
}
// SetAdditionalObjects sets the additionalObjects property value. The additionalObjects property
func (m *ProvisioningProvisionedSystemSyncLog) SetAdditionalObjects(value ProvisioningProvisionedSystemSyncLog_additionalObjectsable)() {
    m.additionalObjects = value
}
// SetCreatedAt sets the createdAt property value. The createdAt property
func (m *ProvisioningProvisionedSystemSyncLog) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdAt = value
}
// SetErrors sets the errors property value. The errors property
func (m *ProvisioningProvisionedSystemSyncLog) SetErrors(value *int32)() {
    m.errors = value
}
// SetLog sets the log property value. The log property
func (m *ProvisioningProvisionedSystemSyncLog) SetLog(value *string)() {
    m.log = value
}
// SetModifications sets the modifications property value. The modifications property
func (m *ProvisioningProvisionedSystemSyncLog) SetModifications(value *int32)() {
    m.modifications = value
}
type ProvisioningProvisionedSystemSyncLogable interface {
    Linkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdditionalObjects()(ProvisioningProvisionedSystemSyncLog_additionalObjectsable)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetErrors()(*int32)
    GetLog()(*string)
    GetModifications()(*int32)
    SetAdditionalObjects(value ProvisioningProvisionedSystemSyncLog_additionalObjectsable)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetErrors(value *int32)()
    SetLog(value *string)()
    SetModifications(value *int32)()
}
