package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RequestModificationRequest 
type RequestModificationRequest struct {
    Linkable
    // The account property
    account AuthAccountPrimerable
    // The additionalObjects property
    additionalObjects RequestModificationRequest_additionalObjectsable
    // The comment property
    comment *string
    // The feedback property
    feedback *string
    // The group property
    group GroupGroupPrimerable
    // The mailKey property
    mailKey *string
    // The type property
    requestModificationRequestType *RequestModificationRequestType
    // The status property
    status *RequestModificationRequestStatus
}
// NewRequestModificationRequest instantiates a new requestModificationRequest and sets the default values.
func NewRequestModificationRequest()(*RequestModificationRequest) {
    m := &RequestModificationRequest{
        Linkable: *NewLinkable(),
    }
    typeEscapedValue := "request.ModificationRequest"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateRequestModificationRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRequestModificationRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("$type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "request.AbstractApplicationModificationRequest":
                        return NewRequestAbstractApplicationModificationRequest(), nil
                    case "request.AbstractOrganizationalUnitModificationRequest":
                        return NewRequestAbstractOrganizationalUnitModificationRequest(), nil
                    case "request.AbstractProvisionedSystemModificationRequest":
                        return NewRequestAbstractProvisionedSystemModificationRequest(), nil
                    case "request.AddGroupAdminRequest":
                        return NewRequestAddGroupAdminRequest(), nil
                    case "request.CreateGroupOnSystemRequest":
                        return NewRequestCreateGroupOnSystemRequest(), nil
                    case "request.CreateGroupRequest":
                        return NewRequestCreateGroupRequest(), nil
                    case "request.CreateProvisionedNamespaceRequest":
                        return NewRequestCreateProvisionedNamespaceRequest(), nil
                    case "request.CreateServiceAccountRequest":
                        return NewRequestCreateServiceAccountRequest(), nil
                    case "request.Disable2FARequest":
                        return NewRequestDisable2FARequest(), nil
                    case "request.EnableTechnicalAdministrationRequest":
                        return NewRequestEnableTechnicalAdministrationRequest(), nil
                    case "request.ExtendAccessRequest":
                        return NewRequestExtendAccessRequest(), nil
                    case "request.GrantAccessRequest":
                        return NewRequestGrantAccessRequest(), nil
                    case "request.GrantApplicationRequest":
                        return NewRequestGrantApplicationRequest(), nil
                    case "request.GrantClientPermissionRequest":
                        return NewRequestGrantClientPermissionRequest(), nil
                    case "request.GrantGroupOnSystemRequest":
                        return NewRequestGrantGroupOnSystemRequest(), nil
                    case "request.GrantGroupOnSystemRequestRequest":
                        return NewRequestGrantGroupOnSystemRequestRequest(), nil
                    case "request.GrantServiceAccountGroupRequest":
                        return NewRequestGrantServiceAccountGroupRequest(), nil
                    case "request.JoinGroupRequest":
                        return NewRequestJoinGroupRequest(), nil
                    case "request.JoinVaultRequest":
                        return NewRequestJoinVaultRequest(), nil
                    case "request.PasswordResetRequest":
                        return NewRequestPasswordResetRequest(), nil
                    case "request.RemoveGroupRequest":
                        return NewRequestRemoveGroupRequest(), nil
                    case "request.RemoveOrganizationalUnitRequest":
                        return NewRequestRemoveOrganizationalUnitRequest(), nil
                    case "request.RemoveProvisionedSystemRequest":
                        return NewRequestRemoveProvisionedSystemRequest(), nil
                    case "request.ReviewAuditRequest":
                        return NewRequestReviewAuditRequest(), nil
                    case "request.RevokeAdminRequest":
                        return NewRequestRevokeAdminRequest(), nil
                    case "request.SetupAuthorizingGroupRequest":
                        return NewRequestSetupAuthorizingGroupRequest(), nil
                    case "request.SetupNestedGroupRequest":
                        return NewRequestSetupNestedGroupRequest(), nil
                    case "request.TransferApplicationAdministrationRequest":
                        return NewRequestTransferApplicationAdministrationRequest(), nil
                    case "request.TransferApplicationOwnershipRequest":
                        return NewRequestTransferApplicationOwnershipRequest(), nil
                    case "request.TransferAuditorGroupRequest":
                        return NewRequestTransferAuditorGroupRequest(), nil
                    case "request.TransferGroupOnSystemOwnershipRequest":
                        return NewRequestTransferGroupOnSystemOwnershipRequest(), nil
                    case "request.TransferOrganizationalUnitOwnershipRequest":
                        return NewRequestTransferOrganizationalUnitOwnershipRequest(), nil
                    case "request.TransferProvisionedSystemAdministrationRequest":
                        return NewRequestTransferProvisionedSystemAdministrationRequest(), nil
                    case "request.TransferProvisionedSystemContentAdministrationRequest":
                        return NewRequestTransferProvisionedSystemContentAdministrationRequest(), nil
                    case "request.TransferProvisionedSystemOwnershipRequest":
                        return NewRequestTransferProvisionedSystemOwnershipRequest(), nil
                    case "request.TransferServiceAccountAdministrationRequest":
                        return NewRequestTransferServiceAccountAdministrationRequest(), nil
                    case "request.UpdateGroupMembershipRequest":
                        return NewRequestUpdateGroupMembershipRequest(), nil
                    case "request.VerifyInternalAccountRequest":
                        return NewRequestVerifyInternalAccountRequest(), nil
                }
            }
        }
    }
    return NewRequestModificationRequest(), nil
}
// GetAccount gets the account property value. The account property
func (m *RequestModificationRequest) GetAccount()(AuthAccountPrimerable) {
    return m.account
}
// GetAdditionalObjects gets the additionalObjects property value. The additionalObjects property
func (m *RequestModificationRequest) GetAdditionalObjects()(RequestModificationRequest_additionalObjectsable) {
    return m.additionalObjects
}
// GetComment gets the comment property value. The comment property
func (m *RequestModificationRequest) GetComment()(*string) {
    return m.comment
}
// GetFeedback gets the feedback property value. The feedback property
func (m *RequestModificationRequest) GetFeedback()(*string) {
    return m.feedback
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RequestModificationRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Linkable.GetFieldDeserializers()
    res["account"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthAccountPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccount(val.(AuthAccountPrimerable))
        }
        return nil
    }
    res["additionalObjects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRequestModificationRequest_additionalObjectsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalObjects(val.(RequestModificationRequest_additionalObjectsable))
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
    res["feedback"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFeedback(val)
        }
        return nil
    }
    res["group"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupGroupPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroup(val.(GroupGroupPrimerable))
        }
        return nil
    }
    res["mailKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMailKey(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRequestModificationRequestType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestModificationRequestType(val.(*RequestModificationRequestType))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRequestModificationRequestStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*RequestModificationRequestStatus))
        }
        return nil
    }
    return res
}
// GetGroup gets the group property value. The group property
func (m *RequestModificationRequest) GetGroup()(GroupGroupPrimerable) {
    return m.group
}
// GetMailKey gets the mailKey property value. The mailKey property
func (m *RequestModificationRequest) GetMailKey()(*string) {
    return m.mailKey
}
// GetRequestModificationRequestType gets the type property value. The type property
func (m *RequestModificationRequest) GetRequestModificationRequestType()(*RequestModificationRequestType) {
    return m.requestModificationRequestType
}
// GetStatus gets the status property value. The status property
func (m *RequestModificationRequest) GetStatus()(*RequestModificationRequestStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *RequestModificationRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Linkable.Serialize(writer)
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
    {
        err = writer.WriteStringValue("feedback", m.GetFeedback())
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
    if m.GetRequestModificationRequestType() != nil {
        cast := (*m.GetRequestModificationRequestType()).String()
        err = writer.WriteStringValue("type", &cast)
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
// SetAccount sets the account property value. The account property
func (m *RequestModificationRequest) SetAccount(value AuthAccountPrimerable)() {
    m.account = value
}
// SetAdditionalObjects sets the additionalObjects property value. The additionalObjects property
func (m *RequestModificationRequest) SetAdditionalObjects(value RequestModificationRequest_additionalObjectsable)() {
    m.additionalObjects = value
}
// SetComment sets the comment property value. The comment property
func (m *RequestModificationRequest) SetComment(value *string)() {
    m.comment = value
}
// SetFeedback sets the feedback property value. The feedback property
func (m *RequestModificationRequest) SetFeedback(value *string)() {
    m.feedback = value
}
// SetGroup sets the group property value. The group property
func (m *RequestModificationRequest) SetGroup(value GroupGroupPrimerable)() {
    m.group = value
}
// SetMailKey sets the mailKey property value. The mailKey property
func (m *RequestModificationRequest) SetMailKey(value *string)() {
    m.mailKey = value
}
// SetRequestModificationRequestType sets the type property value. The type property
func (m *RequestModificationRequest) SetRequestModificationRequestType(value *RequestModificationRequestType)() {
    m.requestModificationRequestType = value
}
// SetStatus sets the status property value. The status property
func (m *RequestModificationRequest) SetStatus(value *RequestModificationRequestStatus)() {
    m.status = value
}
// RequestModificationRequestable 
type RequestModificationRequestable interface {
    Linkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccount()(AuthAccountPrimerable)
    GetAdditionalObjects()(RequestModificationRequest_additionalObjectsable)
    GetComment()(*string)
    GetFeedback()(*string)
    GetGroup()(GroupGroupPrimerable)
    GetMailKey()(*string)
    GetRequestModificationRequestType()(*RequestModificationRequestType)
    GetStatus()(*RequestModificationRequestStatus)
    SetAccount(value AuthAccountPrimerable)()
    SetAdditionalObjects(value RequestModificationRequest_additionalObjectsable)()
    SetComment(value *string)()
    SetFeedback(value *string)()
    SetGroup(value GroupGroupPrimerable)()
    SetMailKey(value *string)()
    SetRequestModificationRequestType(value *RequestModificationRequestType)()
    SetStatus(value *RequestModificationRequestStatus)()
}
