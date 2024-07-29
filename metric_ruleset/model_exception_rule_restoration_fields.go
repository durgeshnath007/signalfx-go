/*
Metric Ruleset API

 Metric ruleset API 

API version: 3.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metric_ruleset

import (
	"encoding/json"
)

// ExceptionRuleRestorationFields Fields of the exceptionRules restoration element. 
type ExceptionRuleRestorationFields struct {
	// ID of the restoration job. 
	RestorationId *string `json:"restorationId,omitempty"`
	// Time at which the API begins the restoration job. 
	StartTime *int32 `json:"startTime,omitempty"`
}

// NewExceptionRuleRestorationFields instantiates a new ExceptionRuleRestorationFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExceptionRuleRestorationFields() *ExceptionRuleRestorationFields {
	this := ExceptionRuleRestorationFields{}
	return &this
}

// NewExceptionRuleRestorationFieldsWithDefaults instantiates a new ExceptionRuleRestorationFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExceptionRuleRestorationFieldsWithDefaults() *ExceptionRuleRestorationFields {
	this := ExceptionRuleRestorationFields{}
	return &this
}

// GetRestorationId returns the RestorationId field value if set, zero value otherwise.
func (o *ExceptionRuleRestorationFields) GetRestorationId() string {
	if o == nil || isNil(o.RestorationId) {
		var ret string
		return ret
	}
	return *o.RestorationId
}

// GetRestorationIdOk returns a tuple with the RestorationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExceptionRuleRestorationFields) GetRestorationIdOk() (*string, bool) {
	if o == nil || isNil(o.RestorationId) {
		return nil, false
	}
	return o.RestorationId, true
}

// HasRestorationId returns a boolean if a field has been set.
func (o *ExceptionRuleRestorationFields) HasRestorationId() bool {
	if o != nil && !isNil(o.RestorationId) {
		return true
	}

	return false
}

// SetRestorationId gets a reference to the given string and assigns it to the RestorationId field.
func (o *ExceptionRuleRestorationFields) SetRestorationId(v string) {
	o.RestorationId = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *ExceptionRuleRestorationFields) GetStartTime() int32 {
	if o == nil || isNil(o.StartTime) {
		var ret int32
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExceptionRuleRestorationFields) GetStartTimeOk() (*int32, bool) {
	if o == nil || isNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *ExceptionRuleRestorationFields) HasStartTime() bool {
	if o != nil && !isNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given int32 and assigns it to the StartTime field.
func (o *ExceptionRuleRestorationFields) SetStartTime(v int32) {
	o.StartTime = &v
}

func (o ExceptionRuleRestorationFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExceptionRuleRestorationFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RestorationId) {
		toSerialize["restorationId"] = o.RestorationId
	}
	if !isNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	return toSerialize, nil
}

type NullableExceptionRuleRestorationFields struct {
	value *ExceptionRuleRestorationFields
	isSet bool
}

func (v NullableExceptionRuleRestorationFields) Get() *ExceptionRuleRestorationFields {
	return v.value
}

func (v *NullableExceptionRuleRestorationFields) Set(val *ExceptionRuleRestorationFields) {
	v.value = val
	v.isSet = true
}

func (v NullableExceptionRuleRestorationFields) IsSet() bool {
	return v.isSet
}

func (v *NullableExceptionRuleRestorationFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExceptionRuleRestorationFields(val *ExceptionRuleRestorationFields) *NullableExceptionRuleRestorationFields {
	return &NullableExceptionRuleRestorationFields{value: val, isSet: true}
}

func (v NullableExceptionRuleRestorationFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExceptionRuleRestorationFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


