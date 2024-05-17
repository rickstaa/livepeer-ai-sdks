/*
Livepeer AI Runner

An application to run AI pipelines

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package livepeer_ai

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"os"
)


// DefaultAPIService DefaultAPI service
type DefaultAPIService service

type ApiImageToImageRequest struct {
	ctx context.Context
	ApiService *DefaultAPIService
	prompt *string
	image *os.File
	modelId *string
	strength *float32
	guidanceScale *float32
	negativePrompt *string
	safetyCheck *bool
	seed *int32
	numImagesPerPrompt *int32
}

func (r ApiImageToImageRequest) Prompt(prompt string) ApiImageToImageRequest {
	r.prompt = &prompt
	return r
}

func (r ApiImageToImageRequest) Image(image *os.File) ApiImageToImageRequest {
	r.image = image
	return r
}

func (r ApiImageToImageRequest) ModelId(modelId string) ApiImageToImageRequest {
	r.modelId = &modelId
	return r
}

func (r ApiImageToImageRequest) Strength(strength float32) ApiImageToImageRequest {
	r.strength = &strength
	return r
}

func (r ApiImageToImageRequest) GuidanceScale(guidanceScale float32) ApiImageToImageRequest {
	r.guidanceScale = &guidanceScale
	return r
}

func (r ApiImageToImageRequest) NegativePrompt(negativePrompt string) ApiImageToImageRequest {
	r.negativePrompt = &negativePrompt
	return r
}

func (r ApiImageToImageRequest) SafetyCheck(safetyCheck bool) ApiImageToImageRequest {
	r.safetyCheck = &safetyCheck
	return r
}

func (r ApiImageToImageRequest) Seed(seed int32) ApiImageToImageRequest {
	r.seed = &seed
	return r
}

func (r ApiImageToImageRequest) NumImagesPerPrompt(numImagesPerPrompt int32) ApiImageToImageRequest {
	r.numImagesPerPrompt = &numImagesPerPrompt
	return r
}

func (r ApiImageToImageRequest) Execute() (*ImageResponse, *http.Response, error) {
	return r.ApiService.ImageToImageExecute(r)
}

/*
ImageToImage Image To Image

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiImageToImageRequest
*/
func (a *DefaultAPIService) ImageToImage(ctx context.Context) ApiImageToImageRequest {
	return ApiImageToImageRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ImageResponse
func (a *DefaultAPIService) ImageToImageExecute(r ApiImageToImageRequest) (*ImageResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ImageResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultAPIService.ImageToImage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/image-to-image"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.prompt == nil {
		return localVarReturnValue, nil, reportError("prompt is required and must be specified")
	}
	if r.image == nil {
		return localVarReturnValue, nil, reportError("image is required and must be specified")
	}
	if r.modelId == nil {
		return localVarReturnValue, nil, reportError("modelId is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "prompt", r.prompt, "")
	var imageLocalVarFormFileName string
	var imageLocalVarFileName     string
	var imageLocalVarFileBytes    []byte

	imageLocalVarFormFileName = "image"
	imageLocalVarFile := r.image

	if imageLocalVarFile != nil {
		fbs, _ := io.ReadAll(imageLocalVarFile)

		imageLocalVarFileBytes = fbs
		imageLocalVarFileName = imageLocalVarFile.Name()
		imageLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: imageLocalVarFileBytes, fileName: imageLocalVarFileName, formFileName: imageLocalVarFormFileName})
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "model_id", r.modelId, "")
	if r.strength != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "strength", r.strength, "")
	}
	if r.guidanceScale != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "guidance_scale", r.guidanceScale, "")
	}
	if r.negativePrompt != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "negative_prompt", r.negativePrompt, "")
	}
	if r.safetyCheck != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "safety_check", r.safetyCheck, "")
	}
	if r.seed != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "seed", r.seed, "")
	}
	if r.numImagesPerPrompt != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "num_images_per_prompt", r.numImagesPerPrompt, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v HTTPError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v HTTPError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v HTTPValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiImageToVideoRequest struct {
	ctx context.Context
	ApiService *DefaultAPIService
	image *os.File
	modelId *string
	height *int32
	width *int32
	fps *int32
	motionBucketId *int32
	noiseAugStrength *float32
	seed *int32
}

func (r ApiImageToVideoRequest) Image(image *os.File) ApiImageToVideoRequest {
	r.image = image
	return r
}

func (r ApiImageToVideoRequest) ModelId(modelId string) ApiImageToVideoRequest {
	r.modelId = &modelId
	return r
}

func (r ApiImageToVideoRequest) Height(height int32) ApiImageToVideoRequest {
	r.height = &height
	return r
}

func (r ApiImageToVideoRequest) Width(width int32) ApiImageToVideoRequest {
	r.width = &width
	return r
}

func (r ApiImageToVideoRequest) Fps(fps int32) ApiImageToVideoRequest {
	r.fps = &fps
	return r
}

func (r ApiImageToVideoRequest) MotionBucketId(motionBucketId int32) ApiImageToVideoRequest {
	r.motionBucketId = &motionBucketId
	return r
}

func (r ApiImageToVideoRequest) NoiseAugStrength(noiseAugStrength float32) ApiImageToVideoRequest {
	r.noiseAugStrength = &noiseAugStrength
	return r
}

func (r ApiImageToVideoRequest) Seed(seed int32) ApiImageToVideoRequest {
	r.seed = &seed
	return r
}

func (r ApiImageToVideoRequest) Execute() (*VideoResponse, *http.Response, error) {
	return r.ApiService.ImageToVideoExecute(r)
}

/*
ImageToVideo Image To Video

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiImageToVideoRequest
*/
func (a *DefaultAPIService) ImageToVideo(ctx context.Context) ApiImageToVideoRequest {
	return ApiImageToVideoRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return VideoResponse
func (a *DefaultAPIService) ImageToVideoExecute(r ApiImageToVideoRequest) (*VideoResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *VideoResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultAPIService.ImageToVideo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/image-to-video"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.image == nil {
		return localVarReturnValue, nil, reportError("image is required and must be specified")
	}
	if r.modelId == nil {
		return localVarReturnValue, nil, reportError("modelId is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	var imageLocalVarFormFileName string
	var imageLocalVarFileName     string
	var imageLocalVarFileBytes    []byte

	imageLocalVarFormFileName = "image"
	imageLocalVarFile := r.image

	if imageLocalVarFile != nil {
		fbs, _ := io.ReadAll(imageLocalVarFile)

		imageLocalVarFileBytes = fbs
		imageLocalVarFileName = imageLocalVarFile.Name()
		imageLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: imageLocalVarFileBytes, fileName: imageLocalVarFileName, formFileName: imageLocalVarFormFileName})
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "model_id", r.modelId, "")
	if r.height != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "height", r.height, "")
	}
	if r.width != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "width", r.width, "")
	}
	if r.fps != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "fps", r.fps, "")
	}
	if r.motionBucketId != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "motion_bucket_id", r.motionBucketId, "")
	}
	if r.noiseAugStrength != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "noise_aug_strength", r.noiseAugStrength, "")
	}
	if r.seed != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "seed", r.seed, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v HTTPError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v HTTPError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v HTTPValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiTextToImageRequest struct {
	ctx context.Context
	ApiService *DefaultAPIService
	textToImageParams *TextToImageParams
}

func (r ApiTextToImageRequest) TextToImageParams(textToImageParams TextToImageParams) ApiTextToImageRequest {
	r.textToImageParams = &textToImageParams
	return r
}

func (r ApiTextToImageRequest) Execute() (*ImageResponse, *http.Response, error) {
	return r.ApiService.TextToImageExecute(r)
}

/*
TextToImage Text To Image

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiTextToImageRequest
*/
func (a *DefaultAPIService) TextToImage(ctx context.Context) ApiTextToImageRequest {
	return ApiTextToImageRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ImageResponse
func (a *DefaultAPIService) TextToImageExecute(r ApiTextToImageRequest) (*ImageResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ImageResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultAPIService.TextToImage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/text-to-image"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.textToImageParams == nil {
		return localVarReturnValue, nil, reportError("textToImageParams is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.textToImageParams
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v HTTPError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v HTTPError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v HTTPValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
