package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ProfileAccessProfileAccountWithAttributes struct {
    ProfileAccessProfileAccount
    // The attributes property
    attributes []ProfileAccessProfileAccountAttributeRuleStatusable
}
// NewProfileAccessProfileAccountWithAttributes instantiates a new ProfileAccessProfileAccountWithAttributes and sets the default values.
func NewProfileAccessProfileAccountWithAttributes()(*ProfileAccessProfileAccountWithAttributes) {
    m := &ProfileAccessProfileAccountWithAttributes{
        ProfileAccessProfileAccount: *NewProfileAccessProfileAccount(),
    }
    typeEscapedValue := "profile.AccessProfileAccountWithAttributes"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateProfileAccessProfileAccountWithAttributesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateProfileAccessProfileAccountWithAttributesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProfileAccessProfileAccountWithAttributes(), nil
}
// GetAttributes gets the attributes property value. The attributes property
// returns a []ProfileAccessProfileAccountAttributeRuleStatusable when successful
func (m *ProfileAccessProfileAccountWithAttributes) GetAttributes()([]ProfileAccessProfileAccountAttributeRuleStatusable) {
    return m.attributes
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ProfileAccessProfileAccountWithAttributes) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ProfileAccessProfileAccount.GetFieldDeserializers()
    res["attributes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateProfileAccessProfileAccountAttributeRuleStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ProfileAccessProfileAccountAttributeRuleStatusable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ProfileAccessProfileAccountAttributeRuleStatusable)
                }
            }
            m.SetAttributes(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ProfileAccessProfileAccountWithAttributes) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ProfileAccessProfileAccount.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// SetAttributes sets the attributes property value. The attributes property
func (m *ProfileAccessProfileAccountWithAttributes) SetAttributes(value []ProfileAccessProfileAccountAttributeRuleStatusable)() {
    m.attributes = value
}
type ProfileAccessProfileAccountWithAttributesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ProfileAccessProfileAccountable
    GetAttributes()([]ProfileAccessProfileAccountAttributeRuleStatusable)
    SetAttributes(value []ProfileAccessProfileAccountAttributeRuleStatusable)()
}