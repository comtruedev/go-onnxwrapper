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

// ErrorCode.h
const (
	ONNX_EXCEPTION                = 100
	STD_EXCEPTION                 = 101
	CV_EXCEPTION                  = 102
	OK                            = 200
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

// OCR 모델 초기화
func InitOCRModels(cornerV1 string, cornerV2 string, textModel string, ocrModel string, dcodeModel string, mrzModel string) (ret ApiReturn) {
	_cornerV1 := C.CString(cornerV1)
	defer C.free(unsafe.Pointer(_cornerV1))
	_cornerV2 := C.CString(cornerV2)
	defer C.free(unsafe.Pointer(_cornerV2))
	_textModel := C.CString(textModel)
	defer C.free(unsafe.Pointer(_textModel))
	_ocrModel := C.CString(ocrModel)
	defer C.free(unsafe.Pointer(_ocrModel))
	_dcodeModel := C.CString(dcodeModel)
	defer C.free(unsafe.Pointer(_dcodeModel))
	_mrzModel := C.CString(mrzModel)
	defer C.free(unsafe.Pointer(_mrzModel))
	_ret := C.ComtrueInitOCRModels(_cornerV1, _cornerV2, _textModel, _ocrModel, _dcodeModel, _mrzModel)
	ret = ApiReturn(_ret)
	return
}

// 얼굴 인식 모델 초기화
func InitFaceModels(faceModel string) (ret ApiReturn) {
	_faceModel := C.CString(faceModel)
	defer C.free(unsafe.Pointer(_faceModel))
	_ret := C.ComtrueInitFaceModels(_faceModel)
	ret = ApiReturn(_ret)
	return
}

// 자동차 번호판 인식 모델 초기화
func InitLicensePlateModels(licensePlateModel string) (ret ApiReturn) {
	_licensePlateModel := C.CString(licensePlateModel)
	defer C.free(unsafe.Pointer(_licensePlateModel))
	_ret := C.ComtrueInitLicensePlateModels(_licensePlateModel)
	ret = ApiReturn(_ret)
	return
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
