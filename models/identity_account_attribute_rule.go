package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type IdentityAccountAttributeRule struct {
    Linkable
    // The additionalObjects property
    additionalObjects IdentityAccountAttributeRule_additionalObjectsable
    // The allowOverride property
    allowOverride *bool
    // The allowSelfService property
    allowSelfService *bool
    // The attribute property
    attribute IdentityAccountAttributeDefinitionable
    // The defaultValue property
    defaultValue *string
    // The priorityDirectory property
    priorityDirectory *int32
    // The priorityExternalSource property
    priorityExternalSource *int32
    // The priorityFormula property
    priorityFormula *int32
    // The prioritySCIM property
    prioritySCIM *int32
    // The script property
    script *string
    // The updateAutomatically property
    updateAutomatically *bool
}
// NewIdentityAccountAttributeRule instantiates a new IdentityAccountAttributeRule and sets the default values.
func NewIdentityAccountAttributeRule()(*IdentityAccountAttributeRule) {
    m := &IdentityAccountAttributeRule{
        Linkable: *NewLinkable(),
    }
    typeEscapedValue := "identity.AccountAttributeRule"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateIdentityAccountAttributeRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIdentityAccountAttributeRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIdentityAccountAttributeRule(), nil
}
// GetAdditionalObjects gets the additionalObjects property value. The additionalObjects property
// returns a IdentityAccountAttributeRule_additionalObjectsable when successful
func (m *IdentityAccountAttributeRule) GetAdditionalObjects()(IdentityAccountAttributeRule_additionalObjectsable) {
    return m.additionalObjects
}
// GetAllowOverride gets the allowOverride property value. The allowOverride property
// returns a *bool when successful
func (m *IdentityAccountAttributeRule) GetAllowOverride()(*bool) {
    return m.allowOverride
}
// GetAllowSelfService gets the allowSelfService property value. The allowSelfService property
// returns a *bool when successful
func (m *IdentityAccountAttributeRule) GetAllowSelfService()(*bool) {
    return m.allowSelfService
}
// GetAttribute gets the attribute property value. The attribute property
// returns a IdentityAccountAttributeDefinitionable when successful
func (m *IdentityAccountAttributeRule) GetAttribute()(IdentityAccountAttributeDefinitionable) {
    return m.attribute
}
// GetDefaultValue gets the defaultValue property value. The defaultValue property
// returns a *string when successful
func (m *IdentityAccountAttributeRule) GetDefaultValue()(*string) {
    return m.defaultValue
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *IdentityAccountAttributeRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Linkable.GetFieldDeserializers()
    res["additionalObjects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentityAccountAttributeRule_additionalObjectsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalObjects(val.(IdentityAccountAttributeRule_additionalObjectsable))
        }
        return nil
    }
    res["allowOverride"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowOverride(val)
        }
        return nil
    }
    res["allowSelfService"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowSelfService(val)
        }
        return nil
    }
    res["attribute"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentityAccountAttributeDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttribute(val.(IdentityAccountAttributeDefinitionable))
        }
        return nil
    }
    res["defaultValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultValue(val)
        }
        return nil
    }
    res["priorityDirectory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPriorityDirectory(val)
        }
        return nil
    }
    res["priorityExternalSource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPriorityExternalSource(val)
        }
        return nil
    }
    res["priorityFormula"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPriorityFormula(val)
        }
        return nil
    }
    res["prioritySCIM"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrioritySCIM(val)
        }
        return nil
    }
    res["script"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScript(val)
        }
        return nil
    }
    res["updateAutomatically"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdateAutomatically(val)
        }
        return nil
    }
    return res
}
// GetPriorityDirectory gets the priorityDirectory property value. The priorityDirectory property
// returns a *int32 when successful
func (m *IdentityAccountAttributeRule) GetPriorityDirectory()(*int32) {
    return m.priorityDirectory
}
// GetPriorityExternalSource gets the priorityExternalSource property value. The priorityExternalSource property
// returns a *int32 when successful
func (m *IdentityAccountAttributeRule) GetPriorityExternalSource()(*int32) {
    return m.priorityExternalSource
}
// GetPriorityFormula gets the priorityFormula property value. The priorityFormula property
// returns a *int32 when successful
func (m *IdentityAccountAttributeRule) GetPriorityFormula()(*int32) {
    return m.priorityFormula
}
// GetPrioritySCIM gets the prioritySCIM property value. The prioritySCIM property
// returns a *int32 when successful
func (m *IdentityAccountAttributeRule) GetPrioritySCIM()(*int32) {
    return m.prioritySCIM
}
// GetScript gets the script property value. The script property
// returns a *string when successful
func (m *IdentityAccountAttributeRule) GetScript()(*string) {
    return m.script
}
// GetUpdateAutomatically gets the updateAutomatically property value. The updateAutomatically property
// returns a *bool when successful
func (m *IdentityAccountAttributeRule) GetUpdateAutomatically()(*bool) {
    return m.updateAutomatically
}
// Serialize serializes information the current object
func (m *IdentityAccountAttributeRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err = writer.WriteBoolValue("allowOverride", m.GetAllowOverride())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("allowSelfService", m.GetAllowSelfService())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("attribute", m.GetAttribute())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("defaultValue", m.GetDefaultValue())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("priorityDirectory", m.GetPriorityDirectory())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("priorityExternalSource", m.GetPriorityExternalSource())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("priorityFormula", m.GetPriorityFormula())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("prioritySCIM", m.GetPrioritySCIM())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("script", m.GetScript())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("updateAutomatically", m.GetUpdateAutomatically())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalObjects sets the additionalObjects property value. The additionalObjects property
func (m *IdentityAccountAttributeRule) SetAdditionalObjects(value IdentityAccountAttributeRule_additionalObjectsable)() {
    m.additionalObjects = value
}
// SetAllowOverride sets the allowOverride property value. The allowOverride property
func (m *IdentityAccountAttributeRule) SetAllowOverride(value *bool)() {
    m.allowOverride = value
}
// SetAllowSelfService sets the allowSelfService property value. The allowSelfService property
func (m *IdentityAccountAttributeRule) SetAllowSelfService(value *bool)() {
    m.allowSelfService = value
}
// SetAttribute sets the attribute property value. The attribute property
func (m *IdentityAccountAttributeRule) SetAttribute(value IdentityAccountAttributeDefinitionable)() {
    m.attribute = value
}
// SetDefaultValue sets the defaultValue property value. The defaultValue property
func (m *IdentityAccountAttributeRule) SetDefaultValue(value *string)() {
    m.defaultValue = value
}
// SetPriorityDirectory sets the priorityDirectory property value. The priorityDirectory property
func (m *IdentityAccountAttributeRule) SetPriorityDirectory(value *int32)() {
    m.priorityDirectory = value
}
// SetPriorityExternalSource sets the priorityExternalSource property value. The priorityExternalSource property
func (m *IdentityAccountAttributeRule) SetPriorityExternalSource(value *int32)() {
    m.priorityExternalSource = value
}
// SetPriorityFormula sets the priorityFormula property value. The priorityFormula property
func (m *IdentityAccountAttributeRule) SetPriorityFormula(value *int32)() {
    m.priorityFormula = value
}
// SetPrioritySCIM sets the prioritySCIM property value. The prioritySCIM property
func (m *IdentityAccountAttributeRule) SetPrioritySCIM(value *int32)() {
    m.prioritySCIM = value
}
// SetScript sets the script property value. The script property
func (m *IdentityAccountAttributeRule) SetScript(value *string)() {
    m.script = value
}
// SetUpdateAutomatically sets the updateAutomatically property value. The updateAutomatically property
func (m *IdentityAccountAttributeRule) SetUpdateAutomatically(value *bool)() {
    m.updateAutomatically = value
}
type IdentityAccountAttributeRuleable interface {
    Linkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdditionalObjects()(IdentityAccountAttributeRule_additionalObjectsable)
    GetAllowOverride()(*bool)
    GetAllowSelfService()(*bool)
    GetAttribute()(IdentityAccountAttributeDefinitionable)
    GetDefaultValue()(*string)
    GetPriorityDirectory()(*int32)
    GetPriorityExternalSource()(*int32)
    GetPriorityFormula()(*int32)
    GetPrioritySCIM()(*int32)
    GetScript()(*string)
    GetUpdateAutomatically()(*bool)
    SetAdditionalObjects(value IdentityAccountAttributeRule_additionalObjectsable)()
    SetAllowOverride(value *bool)()
    SetAllowSelfService(value *bool)()
    SetAttribute(value IdentityAccountAttributeDefinitionable)()
    SetDefaultValue(value *string)()
    SetPriorityDirectory(value *int32)()
    SetPriorityExternalSource(value *int32)()
    SetPriorityFormula(value *int32)()
    SetPrioritySCIM(value *int32)()
    SetScript(value *string)()
    SetUpdateAutomatically(value *bool)()
}
