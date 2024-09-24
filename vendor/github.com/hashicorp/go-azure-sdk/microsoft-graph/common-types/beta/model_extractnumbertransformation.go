package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CustomClaimTransformation = ExtractNumberTransformation{}

type ExtractNumberTransformation struct {
	Type *TransformationExtractType `json:"type,omitempty"`

	// Fields inherited from CustomClaimTransformation

	// The input attribute that provides the source for the transformation. This parameter is required if it's the first or
	// only transformation in the list of transformations to be applied. Subsequent transformations use the output of the
	// prior transformation as input.
	Input *TransformationAttribute `json:"input,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s ExtractNumberTransformation) CustomClaimTransformation() BaseCustomClaimTransformationImpl {
	return BaseCustomClaimTransformationImpl{
		Input:     s.Input,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ExtractNumberTransformation{}

func (s ExtractNumberTransformation) MarshalJSON() ([]byte, error) {
	type wrapper ExtractNumberTransformation
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ExtractNumberTransformation: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ExtractNumberTransformation: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.extractNumberTransformation"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ExtractNumberTransformation: %+v", err)
	}

	return encoded, nil
}