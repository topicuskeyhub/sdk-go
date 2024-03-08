package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AuthPermission struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The full property
    full *string
    // The instances property
    instances []string
    // The operations property
    operations []AuthPermittedOperation
    // The type property
    typeEscaped *string
}
// NewAuthPermission instantiates a new AuthPermission and sets the default values.
func NewAuthPermission()(*AuthPermission) {
    m := &AuthPermission{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAuthPermissionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAuthPermissionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthPermission(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AuthPermission) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AuthPermission) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["full"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFull(val)
        }
        return nil
    }
    res["instances"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetInstances(res)
        }
        return nil
    }
    res["operations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseAuthPermittedOperation)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuthPermittedOperation, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*AuthPermittedOperation))
                }
            }
            m.SetOperations(res)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val)
        }
        return nil
    }
    return res
}
// GetFull gets the full property value. The full property
// returns a *string when successful
func (m *AuthPermission) GetFull()(*string) {
    return m.full
}
// GetInstances gets the instances property value. The instances property
// returns a []string when successful
func (m *AuthPermission) GetInstances()([]string) {
    return m.instances
}
// GetOperations gets the operations property value. The operations property
// returns a []AuthPermittedOperation when successful
func (m *AuthPermission) GetOperations()([]AuthPermittedOperation) {
    return m.operations
}
// GetTypeEscaped gets the type property value. The type property
// returns a *string when successful
func (m *AuthPermission) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *AuthPermission) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("full", m.GetFull())
        if err != nil {
            return err
        }
    }
    if m.GetInstances() != nil {
        err := writer.WriteCollectionOfStringValues("instances", m.GetInstances())
        if err != nil {
            return err
        }
    }
    if m.GetOperations() != nil {
        err := writer.WriteCollectionOfStringValues("operations", SerializeAuthPermittedOperation(m.GetOperations()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type", m.GetTypeEscaped())
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
func (m *AuthPermission) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetFull sets the full property value. The full property
func (m *AuthPermission) SetFull(value *string)() {
    m.full = value
}
// SetInstances sets the instances property value. The instances property
func (m *AuthPermission) SetInstances(value []string)() {
    m.instances = value
}
// SetOperations sets the operations property value. The operations property
func (m *AuthPermission) SetOperations(value []AuthPermittedOperation)() {
    m.operations = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *AuthPermission) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
type AuthPermissionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFull()(*string)
    GetInstances()([]string)
    GetOperations()([]AuthPermittedOperation)
    GetTypeEscaped()(*string)
    SetFull(value *string)()
    SetInstances(value []string)()
    SetOperations(value []AuthPermittedOperation)()
    SetTypeEscaped(value *string)()
}
