/*
Impart Security v0 REST API

Imparts v0 REST API.

API version: 0.0.0
Contact: support@impart.security
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// checks if the ApiBinding type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiBinding{}

// ApiBinding struct for ApiBinding
type ApiBinding struct {
	// The ID for an API binding.
	Id string `json:"id"`
	// The name of an API binding.
	Name string `json:"name"`
	// The spec ID associated with an API binding.
	SpecId string `json:"spec_id"`
	// The hostname for an API Binding.
	Hostname string `json:"hostname"`
	// The port for an API binding.
	Port int32 `json:"port"`
	// The base path for an API binding.
	BasePath string `json:"base_path"`
	// The upstream origin for an API binding.
	UpstreamOrigin string `json:"upstream_origin"`
	// The date when traffic to the API binding was last observed.
	ObservedAt NullableTime `json:"observed_at"`
	// ID of the member that created the API binding.
	CreatedBy string `json:"created_by"`
	// The date the API binding was created.
	CreatedAt time.Time `json:"created_at"`
	// How many HTTP hops are before the inspector. `-1` means the inspector will try to figure this out based on heuristics.
	Hops int32 `json:"hops"`
	// Instructs the inspector to extract hop-to-hop info from the `forwarded` header.
	UseForwarded bool `json:"use_forwarded"`
	// A list of headers which contain the forwarded client IP. Providing more than 1 header may be a security risk.
	ForwardedFor []string `json:"forwarded_for"`
	// A list of headers which contain the forwarded host.
	ForwardedHost []string `json:"forwarded_host"`
	// A list of headers which contain the forwarded HTTP protocol. The value may be HTTP, HTTPS, on/off, true/false, or 0/1.
	ForwardedProto []string `json:"forwarded_proto"`
	// A list of headers which contain the request ID. If no request ID is found a random one will be generated.
	ForwardedId []string `json:"forwarded_id"`
}

// NewApiBinding instantiates a new ApiBinding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiBinding(id string, name string, specId string, hostname string, port int32, basePath string, upstreamOrigin string, observedAt NullableTime, createdBy string, createdAt time.Time, hops int32, useForwarded bool, forwardedFor []string, forwardedHost []string, forwardedProto []string, forwardedId []string) *ApiBinding {
	this := ApiBinding{}
	this.Id = id
	this.Name = name
	this.SpecId = specId
	this.Hostname = hostname
	this.Port = port
	this.BasePath = basePath
	this.UpstreamOrigin = upstreamOrigin
	this.ObservedAt = observedAt
	this.CreatedBy = createdBy
	this.CreatedAt = createdAt
	this.Hops = hops
	this.UseForwarded = useForwarded
	this.ForwardedFor = forwardedFor
	this.ForwardedHost = forwardedHost
	this.ForwardedProto = forwardedProto
	this.ForwardedId = forwardedId
	return &this
}

// NewApiBindingWithDefaults instantiates a new ApiBinding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiBindingWithDefaults() *ApiBinding {
	this := ApiBinding{}
	return &this
}

// GetId returns the Id field value
func (o *ApiBinding) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ApiBinding) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ApiBinding) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ApiBinding) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApiBinding) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApiBinding) SetName(v string) {
	o.Name = v
}

// GetSpecId returns the SpecId field value
func (o *ApiBinding) GetSpecId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SpecId
}

// GetSpecIdOk returns a tuple with the SpecId field value
// and a boolean to check if the value has been set.
func (o *ApiBinding) GetSpecIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpecId, true
}

// SetSpecId sets field value
func (o *ApiBinding) SetSpecId(v string) {
	o.SpecId = v
}

// GetHostname returns the Hostname field value
func (o *ApiBinding) GetHostname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value
// and a boolean to check if the value has been set.
func (o *ApiBinding) GetHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hostname, true
}

// SetHostname sets field value
func (o *ApiBinding) SetHostname(v string) {
	o.Hostname = v
}

// GetPort returns the Port field value
func (o *ApiBinding) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *ApiBinding) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *ApiBinding) SetPort(v int32) {
	o.Port = v
}

// GetBasePath returns the BasePath field value
func (o *ApiBinding) GetBasePath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BasePath
}

// GetBasePathOk returns a tuple with the BasePath field value
// and a boolean to check if the value has been set.
func (o *ApiBinding) GetBasePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BasePath, true
}

// SetBasePath sets field value
func (o *ApiBinding) SetBasePath(v string) {
	o.BasePath = v
}

// GetUpstreamOrigin returns the UpstreamOrigin field value
func (o *ApiBinding) GetUpstreamOrigin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpstreamOrigin
}

// GetUpstreamOriginOk returns a tuple with the UpstreamOrigin field value
// and a boolean to check if the value has been set.
func (o *ApiBinding) GetUpstreamOriginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpstreamOrigin, true
}

// SetUpstreamOrigin sets field value
func (o *ApiBinding) SetUpstreamOrigin(v string) {
	o.UpstreamOrigin = v
}

// GetObservedAt returns the ObservedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *ApiBinding) GetObservedAt() time.Time {
	if o == nil || o.ObservedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.ObservedAt.Get()
}

// GetObservedAtOk returns a tuple with the ObservedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiBinding) GetObservedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ObservedAt.Get(), o.ObservedAt.IsSet()
}

// SetObservedAt sets field value
func (o *ApiBinding) SetObservedAt(v time.Time) {
	o.ObservedAt.Set(&v)
}

// GetCreatedBy returns the CreatedBy field value
func (o *ApiBinding) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *ApiBinding) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *ApiBinding) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ApiBinding) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ApiBinding) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ApiBinding) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetHops returns the Hops field value
func (o *ApiBinding) GetHops() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Hops
}

// GetHopsOk returns a tuple with the Hops field value
// and a boolean to check if the value has been set.
func (o *ApiBinding) GetHopsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hops, true
}

// SetHops sets field value
func (o *ApiBinding) SetHops(v int32) {
	o.Hops = v
}

// GetUseForwarded returns the UseForwarded field value
func (o *ApiBinding) GetUseForwarded() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UseForwarded
}

// GetUseForwardedOk returns a tuple with the UseForwarded field value
// and a boolean to check if the value has been set.
func (o *ApiBinding) GetUseForwardedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UseForwarded, true
}

// SetUseForwarded sets field value
func (o *ApiBinding) SetUseForwarded(v bool) {
	o.UseForwarded = v
}

// GetForwardedFor returns the ForwardedFor field value
func (o *ApiBinding) GetForwardedFor() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ForwardedFor
}

// GetForwardedForOk returns a tuple with the ForwardedFor field value
// and a boolean to check if the value has been set.
func (o *ApiBinding) GetForwardedForOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ForwardedFor, true
}

// SetForwardedFor sets field value
func (o *ApiBinding) SetForwardedFor(v []string) {
	o.ForwardedFor = v
}

// GetForwardedHost returns the ForwardedHost field value
func (o *ApiBinding) GetForwardedHost() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ForwardedHost
}

// GetForwardedHostOk returns a tuple with the ForwardedHost field value
// and a boolean to check if the value has been set.
func (o *ApiBinding) GetForwardedHostOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ForwardedHost, true
}

// SetForwardedHost sets field value
func (o *ApiBinding) SetForwardedHost(v []string) {
	o.ForwardedHost = v
}

// GetForwardedProto returns the ForwardedProto field value
func (o *ApiBinding) GetForwardedProto() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ForwardedProto
}

// GetForwardedProtoOk returns a tuple with the ForwardedProto field value
// and a boolean to check if the value has been set.
func (o *ApiBinding) GetForwardedProtoOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ForwardedProto, true
}

// SetForwardedProto sets field value
func (o *ApiBinding) SetForwardedProto(v []string) {
	o.ForwardedProto = v
}

// GetForwardedId returns the ForwardedId field value
func (o *ApiBinding) GetForwardedId() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ForwardedId
}

// GetForwardedIdOk returns a tuple with the ForwardedId field value
// and a boolean to check if the value has been set.
func (o *ApiBinding) GetForwardedIdOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ForwardedId, true
}

// SetForwardedId sets field value
func (o *ApiBinding) SetForwardedId(v []string) {
	o.ForwardedId = v
}

func (o ApiBinding) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiBinding) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["spec_id"] = o.SpecId
	toSerialize["hostname"] = o.Hostname
	toSerialize["port"] = o.Port
	toSerialize["base_path"] = o.BasePath
	toSerialize["upstream_origin"] = o.UpstreamOrigin
	toSerialize["observed_at"] = o.ObservedAt.Get()
	toSerialize["created_by"] = o.CreatedBy
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["hops"] = o.Hops
	toSerialize["use_forwarded"] = o.UseForwarded
	toSerialize["forwarded_for"] = o.ForwardedFor
	toSerialize["forwarded_host"] = o.ForwardedHost
	toSerialize["forwarded_proto"] = o.ForwardedProto
	toSerialize["forwarded_id"] = o.ForwardedId
	return toSerialize, nil
}

type NullableApiBinding struct {
	value *ApiBinding
	isSet bool
}

func (v NullableApiBinding) Get() *ApiBinding {
	return v.value
}

func (v *NullableApiBinding) Set(val *ApiBinding) {
	v.value = val
	v.isSet = true
}

func (v NullableApiBinding) IsSet() bool {
	return v.isSet
}

func (v *NullableApiBinding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiBinding(val *ApiBinding) *NullableApiBinding {
	return &NullableApiBinding{value: val, isSet: true}
}

func (v NullableApiBinding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiBinding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}