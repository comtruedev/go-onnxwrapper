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

// ComtrueAPICode
type APICode int32

const (
	ApiOk              = 0x0
	ApiInitFail        = 0x1
	ApiJsonLoadFail    = 0x2
	ApiFail            = 0x3
	ApiDetectSqureFail = 0x4
	ApiDetectTextFail  = 0x5
	ApiOcrFail         = 0x6
	ApiImageLoadFail   = 0x7
)

// ComtrueAPI_RETURN
type ApiReturn uint32

// ComtrueInitFromFile
func InitFromFile(cornerV1 string, cornerV2 string, textModel string, ocrModel string, dcodeModel string) (ret ApiReturn) {
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
	_ret := C.ComtrueInitFromFile(_cornerV1, _cornerV2, _textModel, _ocrModel, _dcodeModel)
	ret = ApiReturn(_ret)
	return
}

// Comtrue_extract_text
func ExtractText(data uintptr, dataLength int) (jsonOutLength int, ret string) {
	_data := unsafe.Pointer(data)
	_dataLength := C.int64_t(dataLength)
	_jsonOutLength := (*C.int64_t)(unsafe.Pointer(&jsonOutLength))
	_ret := C.Comtrue_extract_text(_data, _dataLength, _jsonOutLength)
	ret = C.GoString(_ret)
	return
}

// Comtrue_identification
func Identification(data uintptr, dataLength int, idType string) (jsonOutLength int, ret string) {
	_data := unsafe.Pointer(data)
	_dataLength := C.int64_t(dataLength)
	_idType := C.CString(idType)
	defer C.free(unsafe.Pointer(_idType))
	_jsonOutLength := (*C.int64_t)(unsafe.Pointer(&jsonOutLength))
	_ret := C.Comtrue_identification(_data, _dataLength, _idType, _jsonOutLength)
	ret = C.GoString(_ret)
	return
}

// Comtrue_identification_extract_text
func IdentificationExtractText(data uintptr, dataLength int) (jsonOutLength int, ret string) {
	_data := unsafe.Pointer(data)
	_dataLength := C.int64_t(dataLength)
	_jsonOutLength := (*C.int64_t)(unsafe.Pointer(&jsonOutLength))
	_ret := C.Comtrue_identification_extract_text(_data, _dataLength, _jsonOutLength)
	ret = C.GoString(_ret)
	return
}
