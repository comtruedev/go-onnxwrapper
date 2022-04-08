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
	INVALID_ID_TYPE               = 107
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
func InitModels(modelFilePath string, key string, threadNum int) (ret ApiReturn) {
	_modelFilePath := C.CString(modelFilePath)
	defer C.free(unsafe.Pointer(_modelFilePath))
	_key := C.CString(key)
	defer C.free(unsafe.Pointer(_key))
	_threadNum := C.int32_t(threadNum)
	_ret := C.ComtrueInitModels(_modelFilePath, _key, _threadNum)
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
func ExtractText(data unsafe.Pointer, dataLength int) (retCode int, ret string, jsonOutLength int) {
	_dataLength := C.int64_t(dataLength)
	_jsonOutLength := (*C.int64_t)(unsafe.Pointer(&jsonOutLength))
	_retCode := (*C.int64_t)(unsafe.Pointer(&retCode))
	_ret := C.Comtrue_extract_text(data, _dataLength, _jsonOutLength, _retCode)
	defer C.free(unsafe.Pointer(_ret))
	ret = C.GoString(_ret)
	return
}

// Comtrue_identification
func Identification(data unsafe.Pointer, dataLength int, idType string) (retCode int, ret string, jsonOutLength int) {
	_dataLength := C.int64_t(dataLength)
	_idType := C.CString(idType)
	defer C.free(unsafe.Pointer(_idType))
	_jsonOutLength := (*C.int64_t)(unsafe.Pointer(&jsonOutLength))
	_retCode := (*C.int64_t)(unsafe.Pointer(&retCode))
	_ret := C.Comtrue_identification(data, _dataLength, _idType, _jsonOutLength, _retCode)
	defer C.free(unsafe.Pointer(_ret))
	ret = C.GoString(_ret)
	return
}

// Comtrue_identification_extract_text
func IdentificationExtractText(data unsafe.Pointer, dataLength int) (retCode int, ret string, jsonOutLength int) {
	_dataLength := C.int64_t(dataLength)
	_jsonOutLength := (*C.int64_t)(unsafe.Pointer(&jsonOutLength))
	_retCode := (*C.int64_t)(unsafe.Pointer(&retCode))
	_ret := C.Comtrue_identification_extract_text(data, _dataLength, _jsonOutLength, _retCode)
	defer C.free(unsafe.Pointer(_ret))
	ret = C.GoString(_ret)
	return
}

// Comtrue_face_detection
func FaceDetection(data unsafe.Pointer, dataLength int, modelType int) (retCode int, ret string, jsonOutLength int) {
	_dataLength := C.int64_t(dataLength)
	_modelType := C.int64_t(modelType)
	_jsonOutLength := (*C.int64_t)(unsafe.Pointer(&jsonOutLength))
	_retCode := (*C.int64_t)(unsafe.Pointer(&retCode))
	_ret := C.Comtrue_face_detection(data, _dataLength, _modelType, _jsonOutLength, _retCode)
	defer C.free(unsafe.Pointer(_ret))
	ret = C.GoString(_ret)
	return
}

// Comtrue_license_plate_detection
func LicensePlateDetection(data unsafe.Pointer, dataLength int) (retCode int, ret string, jsonOutLength int) {
	_dataLength := C.int64_t(dataLength)
	_jsonOutLength := (*C.int64_t)(unsafe.Pointer(&jsonOutLength))
	_retCode := (*C.int64_t)(unsafe.Pointer(&retCode))
	_ret := C.Comtrue_license_plate_detection(data, _dataLength, _jsonOutLength, _retCode)
	defer C.free(unsafe.Pointer(_ret))
	ret = C.GoString(_ret)
	return
}

// Comtrue_document_classification
func DocumentClassification(data unsafe.Pointer, dataLength int, detectOrientationType int32, ocrModelType int32) (retCode int, ret string, jsonOutLength int) {
	_data := data
	_dataLength := C.int64_t(dataLength)
	_detectOrientationType := C.int(detectOrientationType)
	_jsonOutLength := (*C.int64_t)(unsafe.Pointer(&jsonOutLength))
	_retCode := (*C.int64_t)(unsafe.Pointer(&retCode))
	_ocrModelType := C.int(ocrModelType)
	_ret := C.Comtrue_document_classification(_data, _dataLength, _detectOrientationType, _jsonOutLength, _retCode, _ocrModelType)
	defer C.free(unsafe.Pointer(_ret))
	ret = C.GoString(_ret)
	return
}

// Comtrue_id_liveness
func IDLiveness(data unsafe.Pointer, dataLength int, idType int) (retCode int, ret string, jsonOutLength int) {
	_dataLength := C.int64_t(dataLength)
	_idType := C.int32_t(idType)
	_jsonOutLength := (*C.int64_t)(unsafe.Pointer(&jsonOutLength))
	_retCode := (*C.int64_t)(unsafe.Pointer(&retCode))
	_ret := C.Comtrue_id_liveness(data, _dataLength, _idType, _jsonOutLength, _retCode)
	defer C.free(unsafe.Pointer(_ret))
	ret = C.GoString(_ret)
	return
}

// Comtrue_face_similarity
func FaceSimilarity(data1 unsafe.Pointer, data1Length int, data2 unsafe.Pointer, data2Length int, checkType int32) (retCode int, ret string, jsonOutLength int) {
	_data1Length := C.int64_t(data1Length)
	_data2Length := C.int64_t(data2Length)
	_checkType := C.int(checkType)
	_jsonOutLength := (*C.int64_t)(unsafe.Pointer(&jsonOutLength))
	_retCode := (*C.int64_t)(unsafe.Pointer(&retCode))
	_ret := C.Comtrue_face_similarity(data1, _data1Length, data2, _data2Length, _checkType, _jsonOutLength, _retCode)
	defer C.free(unsafe.Pointer(_ret))
	ret = C.GoString(_ret)
	return
}

// Comtrue_table_ocr
func TableOCR(data unsafe.Pointer, dataLength int, ocrModelType int) (retCode int, ret string, jsonOutLength int) {
	_dataLength := C.int64_t(dataLength)
	_jsonOutLength := (*C.int64_t)(unsafe.Pointer(&jsonOutLength))
	_ocrModelType := C.int64_t(ocrModelType)
	_retCode := (*C.int64_t)(unsafe.Pointer(&retCode))
	_ret := C.Comtrue_table_ocr(data, _dataLength, _ocrModelType, _jsonOutLength, _retCode)
	defer C.free(unsafe.Pointer(_ret))
	ret = C.GoString(_ret)
	return
}

func InitOCRCharSet(charSetFilePath string) (ret ApiReturn) {
	_charSetFilePath := C.CString(charSetFilePath)
	defer C.free(unsafe.Pointer(_charSetFilePath))
	_ret := C.ComtrueInitOCRCharSet(_charSetFilePath)
	ret = ApiReturn(_ret)
	return
}

// Comtrue_deidentification
func Deidentification(data unsafe.Pointer, dataLength int, idType string, didOption int) (retCode int, ret string, jsonOutLength int) {
	_data := data
	_dataLength := C.int64_t(dataLength)
	_idType := C.CString(idType)
	defer C.free(unsafe.Pointer(_idType))
	_didOption := C.int64_t(didOption)
	_jsonOutLength := (*C.int64_t)(unsafe.Pointer(&jsonOutLength))
	_retCode := (*C.int64_t)(unsafe.Pointer(&retCode))
	_ret := C.Comtrue_deidentification(_data, _dataLength, _idType, _didOption, _jsonOutLength, _retCode)
	defer C.free(unsafe.Pointer(_ret))
	ret = C.GoString(_ret)
	return
}

// Comtrue_document_info_extraction
func DocumentInfoExtraction(data unsafe.Pointer, dataLength int, docType int32, detectOrientationType int32, option int32) (retCode int, ret string, jsonOutLength int) {
	_data := data
	_dataLength := C.int64_t(dataLength)
	_docType := C.int(docType)
	_detectOrientationType := C.int(detectOrientationType)
	_jsonOutLength := (*C.int64_t)(unsafe.Pointer(&jsonOutLength))
	_retCode := (*C.int64_t)(unsafe.Pointer(&retCode))
	_option := C.int(option)
	_ret := C.Comtrue_document_info_extraction(_data, _dataLength, _docType, _detectOrientationType, _jsonOutLength, _retCode, _option)
	defer C.free(unsafe.Pointer(_ret))
	ret = C.GoString(_ret)
	return
}