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
func ExtractText(data unsafe.Pointer, dataLength int, detectOrientationType int32) (retCode int, ret string, jsonOutLength int) {
	_data := data
	_dataLength := C.int64_t(dataLength)
	_detectOrientationType := C.int(detectOrientationType)
	_jsonOutLength := (*C.int64_t)(unsafe.Pointer(&jsonOutLength))
	_retCode := (*C.int64_t)(unsafe.Pointer(&retCode))
	_ret := C.Comtrue_extract_text(_data, _dataLength, _detectOrientationType, _jsonOutLength, _retCode)
	defer C.free(unsafe.Pointer(_ret))
	ret = C.GoString(_ret)
	return
}

// Comtrue_identification
func Identification(data unsafe.Pointer, dataLength int, idType string, detectionType int32) (retCode int, ret string, jsonOutLength int) {
	_data := data
	_dataLength := C.int64_t(dataLength)
	_idType := C.CString(idType)
	defer C.free(unsafe.Pointer(_idType))
	_detectionType := C.int(detectionType)
	_jsonOutLength := (*C.int64_t)(unsafe.Pointer(&jsonOutLength))
	_retCode := (*C.int64_t)(unsafe.Pointer(&retCode))
	_ret := C.Comtrue_identification(_data, _dataLength, _idType, _detectionType, _jsonOutLength, _retCode)
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
func Deidentification(data unsafe.Pointer, dataLength int, idType string, didOption int, detectionType int32) (retCode int, ret string, jsonOutLength int) {
	_data := data
	_dataLength := C.int64_t(dataLength)
	_idType := C.CString(idType)
	defer C.free(unsafe.Pointer(_idType))
	_didOption := C.int64_t(didOption)
	_detectionType := C.int(detectionType)
	_jsonOutLength := (*C.int64_t)(unsafe.Pointer(&jsonOutLength))
	_retCode := (*C.int64_t)(unsafe.Pointer(&retCode))
	_ret := C.Comtrue_deidentification(_data, _dataLength, _idType, _didOption, _detectionType, _jsonOutLength, _retCode)
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

// Comtrue_detect_text_with_dbnet_option
func DetectTextWithDbnetOption(data unsafe.Pointer, dataLength int, detectOrientationType int32, padding int32, maxSideLen int32, boxScoreThresh float32, boxThresh float32, unClipRatio float32, modelType int) (retCode int, ret string, jsonOutLength int) {
	_data := data
	_dataLength := C.int64_t(dataLength)
	_detectOrientationType := C.int(detectOrientationType)
	_jsonOutLength := (*C.int64_t)(unsafe.Pointer(&jsonOutLength))
	_padding := C.int(padding)
	_maxSideLen := C.int(maxSideLen)
	_boxScoreThresh := C.float(boxScoreThresh)
	_boxThresh := C.float(boxThresh)
	_unClipRatio := C.float(unClipRatio)
	_modelType := C.int64_t(modelType)
	_retCode := (*C.int64_t)(unsafe.Pointer(&retCode))
	_ret := C.Comtrue_detect_text_with_dbnet_option(_data, _dataLength, _detectOrientationType, _jsonOutLength, _padding, _maxSideLen, _boxScoreThresh, _boxThresh, _unClipRatio, _modelType, _retCode)
	defer C.free(unsafe.Pointer(_ret))
	ret = C.GoString(_ret)
	return
}

// Comtrue_extract_text_with_dbnet_option
func ExtractTextWithDbnetOption(data unsafe.Pointer, dataLength int, detectOrientationType int32, padding int32, maxSideLen int32, boxScoreThresh float32, boxThresh float32, unClipRatio float32) (retCode int, ret string, jsonOutLength int) {
	_data := data
	_dataLength := C.int64_t(dataLength)
	_detectOrientationType := C.int(detectOrientationType)
	_jsonOutLength := (*C.int64_t)(unsafe.Pointer(&jsonOutLength))
	_padding := C.int(padding)
	_maxSideLen := C.int(maxSideLen)
	_boxScoreThresh := C.float(boxScoreThresh)
	_boxThresh := C.float(boxThresh)
	_unClipRatio := C.float(unClipRatio)
	_retCode := (*C.int64_t)(unsafe.Pointer(&retCode))
	_ret := C.Comtrue_extract_text_with_dbnet_option(_data, _dataLength, _detectOrientationType, _jsonOutLength, _padding, _maxSideLen, _boxScoreThresh, _boxThresh, _unClipRatio, _retCode)
	defer C.free(unsafe.Pointer(_ret))
	ret = C.GoString(_ret)
	return
}

// Comtrue_receipt_info_extraction
func ReceiptInfoExtraction(data unsafe.Pointer, dataLength int) (retCode int, ret string, jsonOutLength int) {
	_data := data
	_dataLength := C.int64_t(dataLength)
	_jsonOutLength := (*C.int64_t)(unsafe.Pointer(&jsonOutLength))
	_retCode := (*C.int64_t)(unsafe.Pointer(&retCode))
	_ret := C.Comtrue_receipt_info_extraction(_data, _dataLength, _jsonOutLength, _retCode)
	defer C.free(unsafe.Pointer(_ret))
	ret = C.GoString(_ret)
	return
}

// Comtrue_fingerprint_detection
func FingerprintDetection(data unsafe.Pointer, dataLength int) (retCode int, ret string, jsonOutLength int) {
	_data := data
	_dataLength := C.int64_t(dataLength)
	_jsonOutLength := (*C.int64_t)(unsafe.Pointer(&jsonOutLength))
	_retCode := (*C.int64_t)(unsafe.Pointer(&retCode))
	_ret := C.Comtrue_fingerprint_detection(_data, _dataLength, _jsonOutLength, _retCode)
	defer C.free(unsafe.Pointer(_ret))
	ret = C.GoString(_ret)
	return
}

// Comtrue_detect_square
func DetectSquare(data unsafe.Pointer, dataLength int, type_ int) (retCode int, ret string, jsonOutLength int) {
	_data := data
	_dataLength := C.int64_t(dataLength)
	_type_ := C.int64_t(type_)
	_jsonOutLength := (*C.int64_t)(unsafe.Pointer(&jsonOutLength))
	_retCode := (*C.int64_t)(unsafe.Pointer(&retCode))
	_ret := C.Comtrue_detect_square(_data, _dataLength, _type_, _jsonOutLength, _retCode)
	defer C.free(unsafe.Pointer(_ret))
	ret = C.GoString(_ret)
	return
}

// Comtrue_identification_ekyc
func IdentificationEkyc(data unsafe.Pointer, dataLength int, idType string, detectionType int32, faceType int32) (retCode int, ret string, jsonOutLength int) {
	_data := data
	_dataLength := C.int64_t(dataLength)
	_idType := C.CString(idType)
	defer C.free(unsafe.Pointer(_idType))
	_detectionType := C.int(detectionType)
	_faceType := C.int(faceType)
	_jsonOutLength := (*C.int64_t)(unsafe.Pointer(&jsonOutLength))
	_retCode := (*C.int64_t)(unsafe.Pointer(&retCode))
	_ret := C.Comtrue_identification_ekyc(_data, _dataLength, _idType, _detectionType, _faceType, _jsonOutLength, _retCode)
	defer C.free(unsafe.Pointer(_ret))
	ret = C.GoString(_ret)
	return
}

// Comtrue_detect_identification
func DetectIdentification(data unsafe.Pointer, dataLength int, type_ int, modelType int32) (retCode int, ret string, jsonOutLength int) {
	_data := data
	_dataLength := C.int64_t(dataLength)
	_type_ := C.int64_t(type_)
	_modelType := C.int(modelType)
	_jsonOutLength := (*C.int64_t)(unsafe.Pointer(&jsonOutLength))
	_retCode := (*C.int64_t)(unsafe.Pointer(&retCode))
	_ret := C.Comtrue_detect_identification(_data, _dataLength, _type_, _modelType, _jsonOutLength, _retCode)
	defer C.free(unsafe.Pointer(_ret))
	ret = C.GoString(_ret)
	return
}

// Comtrue_face_detection_active_liveness
func FaceDetectionActiveLiveness(data unsafe.Pointer, dataLength int, type_ string) (retCode int, ret string, jsonOutLength int) {
	_data := data
	_dataLength := C.int64_t(dataLength)
	_type_ := C.CString(type_)
	defer C.free(unsafe.Pointer(_type_))
	_jsonOutLength := (*C.int64_t)(unsafe.Pointer(&jsonOutLength))
	_retCode := (*C.int64_t)(unsafe.Pointer(&retCode))
	_ret := C.Comtrue_face_detection_active_liveness(_data, _dataLength, _type_, _jsonOutLength, _retCode)
	defer C.free(unsafe.Pointer(_ret))
	ret = C.GoString(_ret)
	return
}

// Comtrue_face_detection_passive_liveness
func FaceDetectionPassiveLiveness(data unsafe.Pointer, dataLength int, type_ int) (retCode int, ret string, jsonOutLength int) {
	_data := data
	_dataLength := C.int64_t(dataLength)
	_type_ := C.int64_t(type_)
	_jsonOutLength := (*C.int64_t)(unsafe.Pointer(&jsonOutLength))
	_retCode := (*C.int64_t)(unsafe.Pointer(&retCode))
	_ret := C.Comtrue_face_detection_passive_liveness(_data, _dataLength, _type_, _jsonOutLength, _retCode)
	defer C.free(unsafe.Pointer(_ret))
	ret = C.GoString(_ret)
	return
}

// ComtrueSetMrzParserHost
func SetMrzParserHost(host string) (ret ApiReturn) {
	_host := C.CString(host)
	defer C.free(unsafe.Pointer(_host))
	_ret := C.ComtrueSetMrzParserHost(_host)
	ret = ApiReturn(_ret)
	return
}

// Comtrue_classify_idcard
func ClassifyIdcard(data unsafe.Pointer, dataLength int, checkType int32) (retCode int, ret string, jsonOutLength int) {
	_data := data
	_dataLength := C.int64_t(dataLength)
	_checkType := C.int(checkType)
	_jsonOutLength := (*C.int64_t)(unsafe.Pointer(&jsonOutLength))
	_retCode := (*C.int64_t)(unsafe.Pointer(&retCode))
	_ret := C.Comtrue_classify_idcard(_data, _dataLength, _checkType, _jsonOutLength, _retCode)
	defer C.free(unsafe.Pointer(_ret))
	ret = C.GoString(_ret)
	return
}

// Comtrue_credit_card_info_extraction
func CreditCardInfoExtraction(data unsafe.Pointer, dataLength int) (retCode int, ret string, jsonOutLength int) {
	_data := data
	_dataLength := C.int64_t(dataLength)
	_jsonOutLength := (*C.int64_t)(unsafe.Pointer(&jsonOutLength))
	_retCode := (*C.int64_t)(unsafe.Pointer(&retCode))
	_ret := C.Comtrue_credit_card_info_extraction(_data, _dataLength, _jsonOutLength, _retCode)
	defer C.free(unsafe.Pointer(_ret))
	ret = C.GoString(_ret)
	return
}

// Comtrue_classify_idcard_and_extract_text
func ClassifyIdcardAndExtractText(data unsafe.Pointer, dataLength int, checkType int32, detectionType int32, faceType int32) (retCode int, ret string, jsonOutLength int) {
	_data := data
	_dataLength := C.int64_t(dataLength)
	_checkType := C.int(checkType)
	_detectionType := C.int(detectionType)
	_faceType := C.int(faceType)
	_jsonOutLength := (*C.int64_t)(unsafe.Pointer(&jsonOutLength))
	_retCode := (*C.int64_t)(unsafe.Pointer(&retCode))
	_ret := C.Comtrue_classify_idcard_and_extract_text(_data, _dataLength, _checkType, _detectionType, _faceType, _jsonOutLength, _retCode)
	defer C.free(unsafe.Pointer(_ret))
	ret = C.GoString(_ret)
	return
}

// Comtrue_extract_imei
func ExtractImei(data unsafe.Pointer, dataLength int) (retCode int, ret string, jsonOutLength int) {
	_data := data
	_dataLength := C.int64_t(dataLength)
	_jsonOutLength := (*C.int64_t)(unsafe.Pointer(&jsonOutLength))
	_retCode := (*C.int64_t)(unsafe.Pointer(&retCode))
	_ret := C.Comtrue_extract_imei(_data, _dataLength, _jsonOutLength, _retCode)
	defer C.free(unsafe.Pointer(_ret))
	ret = C.GoString(_ret)
	return
}

// Comtrue_face_detection_embedding
func FaceDetectionEmbedding(data unsafe.Pointer, dataLength int) (retCode int, ret string, jsonOutLength int) {
	_data := data
	_dataLength := C.int64_t(dataLength)
	_jsonOutLength := (*C.int64_t)(unsafe.Pointer(&jsonOutLength))
	_retCode := (*C.int64_t)(unsafe.Pointer(&retCode))
	_ret := C.Comtrue_face_detection_embedding(_data, _dataLength, _jsonOutLength, _retCode)
	defer C.free(unsafe.Pointer(_ret))
	ret = C.GoString(_ret)
	return
}

// ComtrueInitExtraData
func InitExtraData(dataFilePath string) (ret ApiReturn) {
	_dataFilePath := C.CString(dataFilePath)
	defer C.free(unsafe.Pointer(_dataFilePath))
	_ret := C.ComtrueInitExtraData(_dataFilePath)
	ret = ApiReturn(_ret)
	return
}