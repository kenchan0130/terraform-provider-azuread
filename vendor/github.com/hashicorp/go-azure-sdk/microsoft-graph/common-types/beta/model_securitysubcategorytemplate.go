package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecurityFilePlanDescriptorTemplate = SecuritySubcategoryTemplate{}

type SecuritySubcategoryTemplate struct {

	// Fields inherited from SecurityFilePlanDescriptorTemplate

	// Represents the user who created the filePlanDescriptorTemplate column. Read-only.
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`

	// Represents the date and time in which the filePlanDescriptorTemplate is created. Read-only.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Unique string that defines a filePlanDescriptorTemplate name.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s SecuritySubcategoryTemplate) SecurityFilePlanDescriptorTemplate() BaseSecurityFilePlanDescriptorTemplateImpl {
	return BaseSecurityFilePlanDescriptorTemplateImpl{
		CreatedBy:       s.CreatedBy,
		CreatedDateTime: s.CreatedDateTime,
		DisplayName:     s.DisplayName,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s SecuritySubcategoryTemplate) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecuritySubcategoryTemplate{}

func (s SecuritySubcategoryTemplate) MarshalJSON() ([]byte, error) {
	type wrapper SecuritySubcategoryTemplate
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecuritySubcategoryTemplate: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecuritySubcategoryTemplate: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.subcategoryTemplate"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecuritySubcategoryTemplate: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &SecuritySubcategoryTemplate{}

func (s *SecuritySubcategoryTemplate) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`
		DisplayName     nullable.Type[string] `json:"displayName,omitempty"`
		Id              *string               `json:"id,omitempty"`
		ODataId         *string               `json:"@odata.id,omitempty"`
		ODataType       *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.CreatedDateTime = decoded.CreatedDateTime
	s.DisplayName = decoded.DisplayName
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SecuritySubcategoryTemplate into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'SecuritySubcategoryTemplate': %+v", err)
		}
		s.CreatedBy = &impl
	}

	return nil
}