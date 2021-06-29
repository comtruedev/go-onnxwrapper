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
func InitOCRModels(ocrModel string) (ret ApiReturn) {
	_ocrModel := C.CString(ocrModel)
	defer C.free(unsafe.Pointer(_ocrModel))
	_ret := C.ComtrueInitOCRModels(_ocrModel)
	ret = ApiReturn(_ret)
	return
}

// Corner 모델 초기화
func InitCornerModels(cornerV1Onnx, cornerV1XML, cornerV1BIN, cornerV2Onnx, cornerV2XML, cornerV2BIN string) (ret ApiReturn) {
	_cornerV1Onnx := C.CString(cornerV1Onnx)
	defer C.free(unsafe.Pointer(_cornerV1Onnx))
	_cornerV1XML := C.CString(cornerV1XML)
	defer C.free(unsafe.Pointer(_cornerV1XML))
	_cornerV1BIN := C.CString(cornerV1BIN)
	defer C.free(unsafe.Pointer(_cornerV1BIN))
	_cornerV2Onnx := C.CString(cornerV2Onnx)
	defer C.free(unsafe.Pointer(_cornerV2Onnx))
	_cornerV2XML := C.CString(cornerV2XML)
	defer C.free(unsafe.Pointer(_cornerV2XML))
	_cornerV2BIN := C.CString(cornerV2BIN)
	defer C.free(unsafe.Pointer(_cornerV2BIN))
	_ret := C.ComtrueInitCornerModels(_cornerV1Onnx, _cornerV1XML, _cornerV1BIN, _cornerV2Onnx, _cornerV2XML, _cornerV2BIN)
	ret = ApiReturn(_ret)
	return
}

// DbNet 모델 초기화
func InitDbNetModels(dbnetOnnx, dbnetXML, dbnetBIN string) (ret ApiReturn) {
	_dbnetOnnx := C.CString(dbnetOnnx)
	defer C.free(unsafe.Pointer(_dbnetOnnx))
	_dbnetXML := C.CString(dbnetXML)
	defer C.free(unsafe.Pointer(_dbnetXML))
	_dbnetBIN := C.CString(dbnetBIN)
	defer C.free(unsafe.Pointer(_dbnetBIN))
	_ret := C.ComtrueInitDbNetModels(_dbnetOnnx, _dbnetXML, _dbnetBIN)
	ret = ApiReturn(_ret)
	return
}

// AngleNet 모델 초기화
func InitAngleNetModels(anglenetModel string) (ret ApiReturn) {
	_anglenetModel := C.CString(anglenetModel)
	defer C.free(unsafe.Pointer(_anglenetModel))
	_ret := C.ComtrueInitAngleNetModels(_anglenetModel)
	ret = ApiReturn(_ret)
	return
}

// text 모델 초기화
func InitTextModels(textOnnx, textXML, textBIN string) (ret ApiReturn) {
	_textOnnx := C.CString(textOnnx)
	defer C.free(unsafe.Pointer(_textOnnx))
	_textXML := C.CString(textXML)
	defer C.free(unsafe.Pointer(_textXML))
	_textBIN := C.CString(textBIN)
	defer C.free(unsafe.Pointer(_textBIN))
	_ret := C.ComtrueInitTextModels(_textOnnx, _textXML, _textBIN)
	ret = ApiReturn(_ret)
	return
}

// dcode 모델 초기화
func InitDcodeModels(dcodeModel string) (ret ApiReturn) {
	_dcodeModel := C.CString(dcodeModel)
	defer C.free(unsafe.Pointer(_dcodeModel))
	_ret := C.ComtrueInitDcodeModels(_dcodeModel)
	ret = ApiReturn(_ret)
	return
}

// mrz 모델 초기화
func InitMRZModels(mrzModel string) (ret ApiReturn) {
	_mrzModel := C.CString(mrzModel)
	defer C.free(unsafe.Pointer(_mrzModel))
	_ret := C.ComtrueInitMRZModels(_mrzModel)
	ret = ApiReturn(_ret)
	return
}

// 얼굴 인식 모델 초기화
func InitFaceModels(faceOnnx, faceXML, faceBIN string) (ret ApiReturn) {
	_faceOnnx := C.CString(faceOnnx)
	defer C.free(unsafe.Pointer(_faceOnnx))
	_faceXML := C.CString(faceXML)
	defer C.free(unsafe.Pointer(_faceXML))
	_faceBIN := C.CString(faceBIN)
	defer C.free(unsafe.Pointer(_faceBIN))
	_ret := C.ComtrueInitFaceModels(_faceOnnx, _faceXML, _faceBIN)
	ret = ApiReturn(_ret)
	return
}

// 자동차 번호판 인식 모델 초기화
func InitLicensePlateModels(licensePlateOnnx, licensePlateXML, licensePlateBIN string) (ret ApiReturn) {
	_licensePlateOnnx := C.CString(licensePlateOnnx)
	defer C.free(unsafe.Pointer(_licensePlateOnnx))
	_licensePlateXML := C.CString(licensePlateXML)
	defer C.free(unsafe.Pointer(_licensePlateXML))
	_licensePlateBIN := C.CString(licensePlateBIN)
	defer C.free(unsafe.Pointer(_licensePlateBIN))
	_ret := C.ComtrueInitLicensePlateModels(_licensePlateOnnx, _licensePlateXML, _licensePlateBIN)
	ret = ApiReturn(_ret)
	return
}

// 문서 분류 모델 초기화
func InitDocumentClassificationModels(documentClassificationOnnx string) (ret ApiReturn) {
	_documentClassificationOnnx := C.CString(documentClassificationOnnx)
	defer C.free(unsafe.Pointer(_documentClassificationOnnx))
	_ret := C.ComtrueInitDocumentClassificationModels(_documentClassificationOnnx)
	ret = ApiReturn(_ret)
	return
}

// OCR 모델 초기화 상태 체크
func CheckInitOCRModels() int {
	ret := int(0)
	_ret := (*C.int)(unsafe.Pointer(&ret))
	C.ComtrueCheckInitOCRModel(_ret)
	return ret
}

// Corner 모델 초기화 상태 체크
func CheckInitCornerModels() int {
	ret := int(0)
	_ret := (*C.int)(unsafe.Pointer(&ret))
	C.ComtrueCheckInitCornerModel(_ret)
	return ret
}

// DbNet 모델 초기화 상태 체크
func CheckInitDbNetModels() int {
	ret := int(0)
	_ret := (*C.int)(unsafe.Pointer(&ret))
	C.ComtrueCheckInitDbNetModel(_ret)
	return ret
}

// AngleNet 모델 초기화 상태 체크
func CheckInitAngleNetModels() int {
	ret := int(0)
	_ret := (*C.int)(unsafe.Pointer(&ret))
	C.ComtrueCheckInitAngleNetModel(_ret)
	return ret
}

// Text 모델 초기화 상태 체크
func CheckInitTextModels() int {
	ret := int(0)
	_ret := (*C.int)(unsafe.Pointer(&ret))
	C.ComtrueCheckInitTextModel(_ret)
	return ret
}

// Dcode 모델 초기화 상태 체크
func CheckInitDcodeModels() int {
	ret := int(0)
	_ret := (*C.int)(unsafe.Pointer(&ret))
	C.ComtrueCheckInitDcodeModel(_ret)
	return ret
}

// MRZ 모델 초기화 상태 체크
func CheckInitMRZModels() int {
	ret := int(0)
	_ret := (*C.int)(unsafe.Pointer(&ret))
	C.ComtrueCheckInitMRZModel(_ret)
	return ret
}

// 얼굴 인식 모델 초기화 상태 체크
func CheckInitFaceModels() int {
	ret := int(0)
	_ret := (*C.int)(unsafe.Pointer(&ret))
	C.ComtrueCheckInitFaceModel(_ret)
	return ret
}

// 자동차 번호판 인식 모델 초기화 상태 체크
func CheckInitLicensePlateModels() int {
	ret := int(0)
	_ret := (*C.int)(unsafe.Pointer(&ret))
	C.ComtrueCheckInitLicensePlateModel(_ret)
	return ret
}

// 문서 분류 모델 초기화 상태 체크
func CheckInitDocumentClassificationModels() int {
	ret := int(0)
	_ret := (*C.int)(unsafe.Pointer(&ret))
	C.ComtrueCheckInitDocumentClassificationModel(_ret)
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
	_ret := C.Comtrue_document_classification(_data, _dataLength, _jsonOutLength, _retCode)
	ret = C.GoString(_ret)
	return
}
