/*
Elastic Email REST API

This API is based on the REST API architecture, allowing the user to easily manage their data with this resource-based approach.    Every API call is established on which specific request type (GET, POST, PUT, DELETE) will be used.    The API has a limit of 20 concurrent connections and a hard timeout of 600 seconds per request.    To start using this API, you will need your Access Token (available <a target=\"_blank\" href=\"https://app.elasticemail.com/marketing/settings/new/manage-api\">here</a>). Remember to keep it safe. Required access levels are listed in the given request’s description.    Downloadable library clients can be found in our Github repository <a target=\"_blank\" href=\"https://github.com/ElasticEmail?tab=repositories&q=%22rest+api%22+in%3Areadme\">here</a>

API version: 4.0.0
Contact: support@elasticemail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ElasticEmail

import (
	"encoding/json"
)

// checks if the CampaignRecipient type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignRecipient{}

// CampaignRecipient A set of lists and segments names to read recipients from
type CampaignRecipient struct {
	// Names of lists from your Account to read recipients from
	ListNames []string `json:"ListNames,omitempty"`
	// Names of segments from your Account to read recipients from
	SegmentNames []string `json:"SegmentNames,omitempty"`
}

// NewCampaignRecipient instantiates a new CampaignRecipient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignRecipient() *CampaignRecipient {
	this := CampaignRecipient{}
	return &this
}

// NewCampaignRecipientWithDefaults instantiates a new CampaignRecipient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignRecipientWithDefaults() *CampaignRecipient {
	this := CampaignRecipient{}
	return &this
}

// GetListNames returns the ListNames field value if set, zero value otherwise.
func (o *CampaignRecipient) GetListNames() []string {
	if o == nil || IsNil(o.ListNames) {
		var ret []string
		return ret
	}
	return o.ListNames
}

// GetListNamesOk returns a tuple with the ListNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignRecipient) GetListNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.ListNames) {
		return nil, false
	}
	return o.ListNames, true
}

// HasListNames returns a boolean if a field has been set.
func (o *CampaignRecipient) HasListNames() bool {
	if o != nil && !IsNil(o.ListNames) {
		return true
	}

	return false
}

// SetListNames gets a reference to the given []string and assigns it to the ListNames field.
func (o *CampaignRecipient) SetListNames(v []string) {
	o.ListNames = v
}

// GetSegmentNames returns the SegmentNames field value if set, zero value otherwise.
func (o *CampaignRecipient) GetSegmentNames() []string {
	if o == nil || IsNil(o.SegmentNames) {
		var ret []string
		return ret
	}
	return o.SegmentNames
}

// GetSegmentNamesOk returns a tuple with the SegmentNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignRecipient) GetSegmentNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.SegmentNames) {
		return nil, false
	}
	return o.SegmentNames, true
}

// HasSegmentNames returns a boolean if a field has been set.
func (o *CampaignRecipient) HasSegmentNames() bool {
	if o != nil && !IsNil(o.SegmentNames) {
		return true
	}

	return false
}

// SetSegmentNames gets a reference to the given []string and assigns it to the SegmentNames field.
func (o *CampaignRecipient) SetSegmentNames(v []string) {
	o.SegmentNames = v
}

func (o CampaignRecipient) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignRecipient) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ListNames) {
		toSerialize["ListNames"] = o.ListNames
	}
	if !IsNil(o.SegmentNames) {
		toSerialize["SegmentNames"] = o.SegmentNames
	}
	return toSerialize, nil
}

type NullableCampaignRecipient struct {
	value *CampaignRecipient
	isSet bool
}

func (v NullableCampaignRecipient) Get() *CampaignRecipient {
	return v.value
}

func (v *NullableCampaignRecipient) Set(val *CampaignRecipient) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignRecipient) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignRecipient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignRecipient(val *CampaignRecipient) *NullableCampaignRecipient {
	return &NullableCampaignRecipient{value: val, isSet: true}
}

func (v NullableCampaignRecipient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignRecipient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


