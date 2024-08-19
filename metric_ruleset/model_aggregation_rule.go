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

// AggregationRule A single aggregation rule in a metric ruleset 
type AggregationRule struct {
	// Name of the aggregation rule
	Name *string `json:"name,omitempty"`
	Matcher MetricMatcher `json:"matcher"`
	// Status of this aggregation rule 
	Enabled bool `json:"enabled"`
	Aggregator MetricAggregator `json:"aggregator"`
	// Information about an aggregation rule. 
	Description *string `json:"description,omitempty"`
}

// NewAggregationRule instantiates a new AggregationRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAggregationRule(matcher MetricMatcher, enabled bool, aggregator MetricAggregator) *AggregationRule {
	this := AggregationRule{}
	var name string = ""
	this.Name = &name
	this.Matcher = matcher
	this.Enabled = enabled
	this.Aggregator = aggregator
	return &this
}

// NewAggregationRuleWithDefaults instantiates a new AggregationRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAggregationRuleWithDefaults() *AggregationRule {
	this := AggregationRule{}
	var name string = ""
	this.Name = &name
	var enabled bool = false
	this.Enabled = enabled
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AggregationRule) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregationRule) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AggregationRule) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AggregationRule) SetName(v string) {
	o.Name = &v
}

// GetMatcher returns the Matcher field value
func (o *AggregationRule) GetMatcher() MetricMatcher {
	if o == nil {
		var ret MetricMatcher
		return ret
	}

	return o.Matcher
}

// GetMatcherOk returns a tuple with the Matcher field value
// and a boolean to check if the value has been set.
func (o *AggregationRule) GetMatcherOk() (*MetricMatcher, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Matcher, true
}

// SetMatcher sets field value
func (o *AggregationRule) SetMatcher(v MetricMatcher) {
	o.Matcher = v
}

// GetEnabled returns the Enabled field value
func (o *AggregationRule) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AggregationRule) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AggregationRule) SetEnabled(v bool) {
	o.Enabled = v
}

// GetAggregator returns the Aggregator field value
func (o *AggregationRule) GetAggregator() MetricAggregator {
	if o == nil {
		var ret MetricAggregator
		return ret
	}

	return o.Aggregator
}

// GetAggregatorOk returns a tuple with the Aggregator field value
// and a boolean to check if the value has been set.
func (o *AggregationRule) GetAggregatorOk() (*MetricAggregator, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Aggregator, true
}

// SetAggregator sets field value
func (o *AggregationRule) SetAggregator(v MetricAggregator) {
	o.Aggregator = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AggregationRule) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregationRule) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AggregationRule) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AggregationRule) SetDescription(v string) {
	o.Description = &v
}

func (o AggregationRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["matcher"] = o.Matcher
	toSerialize["enabled"] = o.Enabled
	toSerialize["aggregator"] = o.Aggregator
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableAggregationRule struct {
	value *AggregationRule
	isSet bool
}

func (v NullableAggregationRule) Get() *AggregationRule {
	return v.value
}

func (v *NullableAggregationRule) Set(val *AggregationRule) {
	v.value = val
	v.isSet = true
}

func (v NullableAggregationRule) IsSet() bool {
	return v.isSet
}

func (v *NullableAggregationRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAggregationRule(val *AggregationRule) *NullableAggregationRule {
	return &NullableAggregationRule{value: val, isSet: true}
}

func (v NullableAggregationRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAggregationRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


