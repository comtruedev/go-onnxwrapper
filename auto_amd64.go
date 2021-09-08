package onnxwrapper

/*
#include <Comtrue/api.hpp>
#include <stdlib.h>
#cgo LDFLAGS: -lonnxwrapperdll
*/
import "C"

import (
	"unsafe"
)

// ComtrueAPI_RETURN
type ApiReturn uint32

const (
	DOCUMENT_CLASSIFICATION_OCR_DEFAULT = 0
	DOCUMENT_CLASSIFICATION_OCR_LIGHT   = 1
	DOCUMENT_CLASSIFICATION_OCR_ALT     = 2
)

// ErrorCode.h
const (
	OK                            = 0
	ONNX_EXCEPTION                = 100
	STD_EXCEPTION                 = 101
	CV_EXCEPTION                  = 102
	FILE_NOT_FOUND                = 103
	INVALID_INIT_KEY              = 104
	MODEL_INIT_FAIL               = 105
	MODEL_NOT_INITIALIZED         = 106
	CAN_NOT_FIND_TEXT             = 201
	CAN_NOT_FIND_VERTEX           = 202
	CAN_NOT_FIND_SQUARE           = 203
	NOT_ENOUGH_TO_BE_SQUARE       = 204
	CAN_NOT_FIND_MRZ              = 205
	NOT_ENOUGH_SIMILAR_AREA       = 206
	NOT_ENOUGH_CONTOUR            = 207
	NOT_ENOUGH_GROUP              = 208
	CAN_NOT_FIND_OBJECT           = 209
	IMAGE_CANNOT_BE_READ          = 400
	IMAGE_IS_EMPTY                = 401
	IMAGE_IS_TOO_SMALL_TO_PROCESS = 402
	IMAGE_PATH_NOT_VALID          = 501
	DATA_VALIDATION_FAIL          = 502
)

// 통합 모델 파일로부터 초기화
func InitModels(modelFilePath string, key string) (ret ApiReturn) {
	_modelFilePath := C.CString(modelFilePath)
	defer C.free(unsafe.Pointer(_modelFilePath))
	_key := C.CString(key)
	defer C.free(unsafe.Pointer(_key))
	_ret := C.ComtrueInitModels(_modelFilePath, _key)
	ret = ApiReturn(_ret)
	return
}

// 통합 모델 초기화 상태 체크
func CheckModelInitState() int {
	ret := int(0)
	_ret := (*C.int)(unsafe.Pointer(&ret))
	C.ComtrueCheckModelInitState(_ret)
	return ret
}

// Comtrue_extract_text
func ExtractText(data uintptr, dataLength int) (retCode int, ret string, jsonOutLength int) {
	_data := unsafe.Pointer(data)
	_dataLength := C.int64_t(dataLength)
	_jsonOutLength := (*C.int64_t)(unsafe.Pointer(&jsonOutLength))
	_retCode := (*C.int64_t)(unsafe.Pointer(&retCode))
	_ret := C.Comtrue_extract_text(_data, _dataLength, _jsonOutLength, _retCode)
	ret = C.GoString(_ret)
	return
}

// Comtrue_identification
func Identification(data uintptr, dataLength int, idType string) (retCode int, ret string, jsonOutLength int) {
	_data := unsafe.Pointer(data)
	_dataLength := C.int64_t(dataLength)
	_idType := C.CString(idType)
	defer C.free(unsafe.Pointer(_idType))
	_jsonOutLength := (*C.int64_t)(unsafe.Pointer(&jsonOutLength))
	_retCode := (*C.int64_t)(unsafe.Pointer(&retCode))
	_ret := C.Comtrue_identification(_data, _dataLength, _idType, _jsonOutLength, _retCode)
	ret = C.GoString(_ret)
	return
}

// Comtrue_identification_extract_text
func IdentificationExtractText(data uintptr, dataLength int) (retCode int, ret string, jsonOutLength int) {
	_data := unsafe.Pointer(data)
	_dataLength := C.int64_t(dataLength)
	_jsonOutLength := (*C.int64_t)(unsafe.Pointer(&jsonOutLength))
	_retCode := (*C.int64_t)(unsafe.Pointer(&retCode))
	_ret := C.Comtrue_identification_extract_text(_data, _dataLength, _jsonOutLength, _retCode)
	ret = C.GoString(_ret)
	return
}

// Comtrue_face_detection
func FaceDetection(data uintptr, dataLength int) (retCode int, ret string, jsonOutLength int) {
	_data := unsafe.Pointer(data)
	_dataLength := C.int64_t(dataLength)
	_jsonOutLength := (*C.int64_t)(unsafe.Pointer(&jsonOutLength))
	_retCode := (*C.int64_t)(unsafe.Pointer(&retCode))
	_ret := C.Comtrue_face_detection(_data, _dataLength, _jsonOutLength, _retCode)
	ret = C.GoString(_ret)
	return
}

// Comtrue_license_plate_detection
func LicensePlateDetection(data uintptr, dataLength int) (retCode int, ret string, jsonOutLength int) {
	_data := unsafe.Pointer(data)
	_dataLength := C.int64_t(dataLength)
	_jsonOutLength := (*C.int64_t)(unsafe.Pointer(&jsonOutLength))
	_retCode := (*C.int64_t)(unsafe.Pointer(&retCode))
	_ret := C.Comtrue_license_plate_detection(_data, _dataLength, _jsonOutLength, _retCode)
	ret = C.GoString(_ret)
	return
}

// Comtrue_document_classification
func DocumentClassification(data uintptr, dataLength int) (retCode int, ret string, jsonOutLength int) {
	_data := unsafe.Pointer(data)
	_dataLength := C.int64_t(dataLength)
	_jsonOutLength := (*C.int64_t)(unsafe.Pointer(&jsonOutLength))
	_retCode := (*C.int64_t)(unsafe.Pointer(&retCode))
	_ret := C.Comtrue_document_classification(_data, _dataLength, _jsonOutLength, _retCode, DOCUMENT_CLASSIFICATION_OCR_DEFAULT)
	ret = C.GoString(_ret)
	return
}
