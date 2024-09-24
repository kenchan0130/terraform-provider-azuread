package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChangeNotificationCollection struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Contains an array of JSON web tokens (JWT) generated by Microsoft Graph for the application to validate the origin of
	// the notifications. Microsoft Graph generates a single token for each distinct app and tenant pair for an item if it
	// exists in the value array. Keep in mind that notifications can contain a mix of items for various apps and tenants
	// that subscribed using the same notification URL. Only provided for change notifications with resource data. Optional.
	ValidationTokens *[]string `json:"validationTokens,omitempty"`

	// The set of notifications being sent to the notification URL. Required.
	Value []ChangeNotification `json:"value"`
}