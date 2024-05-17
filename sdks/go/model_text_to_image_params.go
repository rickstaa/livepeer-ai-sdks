/*
Livepeer AI Runner

An application to run AI pipelines

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package livepeer_ai

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the TextToImageParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TextToImageParams{}

// TextToImageParams struct for TextToImageParams
type TextToImageParams struct {
	ModelId string `json:"model_id"`
	Prompt string `json:"prompt"`
	Height *int32 `json:"height,omitempty"`
	Width *int32 `json:"width,omitempty"`
	GuidanceScale *float32 `json:"guidance_scale,omitempty"`
	NegativePrompt *string `json:"negative_prompt,omitempty"`
	SafetyCheck *bool `json:"safety_check,omitempty"`
	Seed *int32 `json:"seed,omitempty"`
	NumInferenceSteps *int32 `json:"num_inference_steps,omitempty"`
	NumImagesPerPrompt *int32 `json:"num_images_per_prompt,omitempty"`
}

type _TextToImageParams TextToImageParams

// NewTextToImageParams instantiates a new TextToImageParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTextToImageParams(modelId string, prompt string) *TextToImageParams {
	this := TextToImageParams{}
	this.ModelId = modelId
	this.Prompt = prompt
	var guidanceScale float32 = 7.5
	this.GuidanceScale = &guidanceScale
	var negativePrompt string = ""
	this.NegativePrompt = &negativePrompt
	var safetyCheck bool = true
	this.SafetyCheck = &safetyCheck
	var numInferenceSteps int32 = 50
	this.NumInferenceSteps = &numInferenceSteps
	var numImagesPerPrompt int32 = 1
	this.NumImagesPerPrompt = &numImagesPerPrompt
	return &this
}

// NewTextToImageParamsWithDefaults instantiates a new TextToImageParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTextToImageParamsWithDefaults() *TextToImageParams {
	this := TextToImageParams{}
	var modelId string = ""
	this.ModelId = modelId
	var guidanceScale float32 = 7.5
	this.GuidanceScale = &guidanceScale
	var negativePrompt string = ""
	this.NegativePrompt = &negativePrompt
	var safetyCheck bool = true
	this.SafetyCheck = &safetyCheck
	var numInferenceSteps int32 = 50
	this.NumInferenceSteps = &numInferenceSteps
	var numImagesPerPrompt int32 = 1
	this.NumImagesPerPrompt = &numImagesPerPrompt
	return &this
}

// GetModelId returns the ModelId field value
func (o *TextToImageParams) GetModelId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModelId
}

// GetModelIdOk returns a tuple with the ModelId field value
// and a boolean to check if the value has been set.
func (o *TextToImageParams) GetModelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelId, true
}

// SetModelId sets field value
func (o *TextToImageParams) SetModelId(v string) {
	o.ModelId = v
}

// GetPrompt returns the Prompt field value
func (o *TextToImageParams) GetPrompt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Prompt
}

// GetPromptOk returns a tuple with the Prompt field value
// and a boolean to check if the value has been set.
func (o *TextToImageParams) GetPromptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prompt, true
}

// SetPrompt sets field value
func (o *TextToImageParams) SetPrompt(v string) {
	o.Prompt = v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *TextToImageParams) GetHeight() int32 {
	if o == nil || IsNil(o.Height) {
		var ret int32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextToImageParams) GetHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *TextToImageParams) HasHeight() bool {
	if o != nil && !IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given int32 and assigns it to the Height field.
func (o *TextToImageParams) SetHeight(v int32) {
	o.Height = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *TextToImageParams) GetWidth() int32 {
	if o == nil || IsNil(o.Width) {
		var ret int32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextToImageParams) GetWidthOk() (*int32, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *TextToImageParams) HasWidth() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given int32 and assigns it to the Width field.
func (o *TextToImageParams) SetWidth(v int32) {
	o.Width = &v
}

// GetGuidanceScale returns the GuidanceScale field value if set, zero value otherwise.
func (o *TextToImageParams) GetGuidanceScale() float32 {
	if o == nil || IsNil(o.GuidanceScale) {
		var ret float32
		return ret
	}
	return *o.GuidanceScale
}

// GetGuidanceScaleOk returns a tuple with the GuidanceScale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextToImageParams) GetGuidanceScaleOk() (*float32, bool) {
	if o == nil || IsNil(o.GuidanceScale) {
		return nil, false
	}
	return o.GuidanceScale, true
}

// HasGuidanceScale returns a boolean if a field has been set.
func (o *TextToImageParams) HasGuidanceScale() bool {
	if o != nil && !IsNil(o.GuidanceScale) {
		return true
	}

	return false
}

// SetGuidanceScale gets a reference to the given float32 and assigns it to the GuidanceScale field.
func (o *TextToImageParams) SetGuidanceScale(v float32) {
	o.GuidanceScale = &v
}

// GetNegativePrompt returns the NegativePrompt field value if set, zero value otherwise.
func (o *TextToImageParams) GetNegativePrompt() string {
	if o == nil || IsNil(o.NegativePrompt) {
		var ret string
		return ret
	}
	return *o.NegativePrompt
}

// GetNegativePromptOk returns a tuple with the NegativePrompt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextToImageParams) GetNegativePromptOk() (*string, bool) {
	if o == nil || IsNil(o.NegativePrompt) {
		return nil, false
	}
	return o.NegativePrompt, true
}

// HasNegativePrompt returns a boolean if a field has been set.
func (o *TextToImageParams) HasNegativePrompt() bool {
	if o != nil && !IsNil(o.NegativePrompt) {
		return true
	}

	return false
}

// SetNegativePrompt gets a reference to the given string and assigns it to the NegativePrompt field.
func (o *TextToImageParams) SetNegativePrompt(v string) {
	o.NegativePrompt = &v
}

// GetSafetyCheck returns the SafetyCheck field value if set, zero value otherwise.
func (o *TextToImageParams) GetSafetyCheck() bool {
	if o == nil || IsNil(o.SafetyCheck) {
		var ret bool
		return ret
	}
	return *o.SafetyCheck
}

// GetSafetyCheckOk returns a tuple with the SafetyCheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextToImageParams) GetSafetyCheckOk() (*bool, bool) {
	if o == nil || IsNil(o.SafetyCheck) {
		return nil, false
	}
	return o.SafetyCheck, true
}

// HasSafetyCheck returns a boolean if a field has been set.
func (o *TextToImageParams) HasSafetyCheck() bool {
	if o != nil && !IsNil(o.SafetyCheck) {
		return true
	}

	return false
}

// SetSafetyCheck gets a reference to the given bool and assigns it to the SafetyCheck field.
func (o *TextToImageParams) SetSafetyCheck(v bool) {
	o.SafetyCheck = &v
}

// GetSeed returns the Seed field value if set, zero value otherwise.
func (o *TextToImageParams) GetSeed() int32 {
	if o == nil || IsNil(o.Seed) {
		var ret int32
		return ret
	}
	return *o.Seed
}

// GetSeedOk returns a tuple with the Seed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextToImageParams) GetSeedOk() (*int32, bool) {
	if o == nil || IsNil(o.Seed) {
		return nil, false
	}
	return o.Seed, true
}

// HasSeed returns a boolean if a field has been set.
func (o *TextToImageParams) HasSeed() bool {
	if o != nil && !IsNil(o.Seed) {
		return true
	}

	return false
}

// SetSeed gets a reference to the given int32 and assigns it to the Seed field.
func (o *TextToImageParams) SetSeed(v int32) {
	o.Seed = &v
}

// GetNumInferenceSteps returns the NumInferenceSteps field value if set, zero value otherwise.
func (o *TextToImageParams) GetNumInferenceSteps() int32 {
	if o == nil || IsNil(o.NumInferenceSteps) {
		var ret int32
		return ret
	}
	return *o.NumInferenceSteps
}

// GetNumInferenceStepsOk returns a tuple with the NumInferenceSteps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextToImageParams) GetNumInferenceStepsOk() (*int32, bool) {
	if o == nil || IsNil(o.NumInferenceSteps) {
		return nil, false
	}
	return o.NumInferenceSteps, true
}

// HasNumInferenceSteps returns a boolean if a field has been set.
func (o *TextToImageParams) HasNumInferenceSteps() bool {
	if o != nil && !IsNil(o.NumInferenceSteps) {
		return true
	}

	return false
}

// SetNumInferenceSteps gets a reference to the given int32 and assigns it to the NumInferenceSteps field.
func (o *TextToImageParams) SetNumInferenceSteps(v int32) {
	o.NumInferenceSteps = &v
}

// GetNumImagesPerPrompt returns the NumImagesPerPrompt field value if set, zero value otherwise.
func (o *TextToImageParams) GetNumImagesPerPrompt() int32 {
	if o == nil || IsNil(o.NumImagesPerPrompt) {
		var ret int32
		return ret
	}
	return *o.NumImagesPerPrompt
}

// GetNumImagesPerPromptOk returns a tuple with the NumImagesPerPrompt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextToImageParams) GetNumImagesPerPromptOk() (*int32, bool) {
	if o == nil || IsNil(o.NumImagesPerPrompt) {
		return nil, false
	}
	return o.NumImagesPerPrompt, true
}

// HasNumImagesPerPrompt returns a boolean if a field has been set.
func (o *TextToImageParams) HasNumImagesPerPrompt() bool {
	if o != nil && !IsNil(o.NumImagesPerPrompt) {
		return true
	}

	return false
}

// SetNumImagesPerPrompt gets a reference to the given int32 and assigns it to the NumImagesPerPrompt field.
func (o *TextToImageParams) SetNumImagesPerPrompt(v int32) {
	o.NumImagesPerPrompt = &v
}

func (o TextToImageParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TextToImageParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["model_id"] = o.ModelId
	toSerialize["prompt"] = o.Prompt
	if !IsNil(o.Height) {
		toSerialize["height"] = o.Height
	}
	if !IsNil(o.Width) {
		toSerialize["width"] = o.Width
	}
	if !IsNil(o.GuidanceScale) {
		toSerialize["guidance_scale"] = o.GuidanceScale
	}
	if !IsNil(o.NegativePrompt) {
		toSerialize["negative_prompt"] = o.NegativePrompt
	}
	if !IsNil(o.SafetyCheck) {
		toSerialize["safety_check"] = o.SafetyCheck
	}
	if !IsNil(o.Seed) {
		toSerialize["seed"] = o.Seed
	}
	if !IsNil(o.NumInferenceSteps) {
		toSerialize["num_inference_steps"] = o.NumInferenceSteps
	}
	if !IsNil(o.NumImagesPerPrompt) {
		toSerialize["num_images_per_prompt"] = o.NumImagesPerPrompt
	}
	return toSerialize, nil
}

func (o *TextToImageParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"model_id",
		"prompt",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varTextToImageParams := _TextToImageParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTextToImageParams)

	if err != nil {
		return err
	}

	*o = TextToImageParams(varTextToImageParams)

	return err
}

type NullableTextToImageParams struct {
	value *TextToImageParams
	isSet bool
}

func (v NullableTextToImageParams) Get() *TextToImageParams {
	return v.value
}

func (v *NullableTextToImageParams) Set(val *TextToImageParams) {
	v.value = val
	v.isSet = true
}

func (v NullableTextToImageParams) IsSet() bool {
	return v.isSet
}

func (v *NullableTextToImageParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTextToImageParams(val *TextToImageParams) *NullableTextToImageParams {
	return &NullableTextToImageParams{value: val, isSet: true}
}

func (v NullableTextToImageParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTextToImageParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

