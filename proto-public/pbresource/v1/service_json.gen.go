// Code generated by protoc-json-shim. DO NOT EDIT.
package resourcev1

import (
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// MarshalJSON is a custom marshaler for ReadRequest
func (this *ReadRequest) MarshalJSON() ([]byte, error) {
	str, err := ServiceMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ReadRequest
func (this *ReadRequest) UnmarshalJSON(b []byte) error {
	return ServiceUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for ReadResponse
func (this *ReadResponse) MarshalJSON() ([]byte, error) {
	str, err := ServiceMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ReadResponse
func (this *ReadResponse) UnmarshalJSON(b []byte) error {
	return ServiceUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for ListRequest
func (this *ListRequest) MarshalJSON() ([]byte, error) {
	str, err := ServiceMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ListRequest
func (this *ListRequest) UnmarshalJSON(b []byte) error {
	return ServiceUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for ListResponse
func (this *ListResponse) MarshalJSON() ([]byte, error) {
	str, err := ServiceMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ListResponse
func (this *ListResponse) UnmarshalJSON(b []byte) error {
	return ServiceUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for ListByOwnerRequest
func (this *ListByOwnerRequest) MarshalJSON() ([]byte, error) {
	str, err := ServiceMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ListByOwnerRequest
func (this *ListByOwnerRequest) UnmarshalJSON(b []byte) error {
	return ServiceUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for ListByOwnerResponse
func (this *ListByOwnerResponse) MarshalJSON() ([]byte, error) {
	str, err := ServiceMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ListByOwnerResponse
func (this *ListByOwnerResponse) UnmarshalJSON(b []byte) error {
	return ServiceUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for WriteRequest
func (this *WriteRequest) MarshalJSON() ([]byte, error) {
	str, err := ServiceMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for WriteRequest
func (this *WriteRequest) UnmarshalJSON(b []byte) error {
	return ServiceUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for WriteResponse
func (this *WriteResponse) MarshalJSON() ([]byte, error) {
	str, err := ServiceMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for WriteResponse
func (this *WriteResponse) UnmarshalJSON(b []byte) error {
	return ServiceUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for WriteStatusRequest
func (this *WriteStatusRequest) MarshalJSON() ([]byte, error) {
	str, err := ServiceMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for WriteStatusRequest
func (this *WriteStatusRequest) UnmarshalJSON(b []byte) error {
	return ServiceUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for WriteStatusResponse
func (this *WriteStatusResponse) MarshalJSON() ([]byte, error) {
	str, err := ServiceMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for WriteStatusResponse
func (this *WriteStatusResponse) UnmarshalJSON(b []byte) error {
	return ServiceUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for DeleteRequest
func (this *DeleteRequest) MarshalJSON() ([]byte, error) {
	str, err := ServiceMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for DeleteRequest
func (this *DeleteRequest) UnmarshalJSON(b []byte) error {
	return ServiceUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for DeleteResponse
func (this *DeleteResponse) MarshalJSON() ([]byte, error) {
	str, err := ServiceMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for DeleteResponse
func (this *DeleteResponse) UnmarshalJSON(b []byte) error {
	return ServiceUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for WatchListRequest
func (this *WatchListRequest) MarshalJSON() ([]byte, error) {
	str, err := ServiceMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for WatchListRequest
func (this *WatchListRequest) UnmarshalJSON(b []byte) error {
	return ServiceUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for WatchEvent
func (this *WatchEvent) MarshalJSON() ([]byte, error) {
	str, err := ServiceMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for WatchEvent
func (this *WatchEvent) UnmarshalJSON(b []byte) error {
	return ServiceUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for WatchEvent_Upsert
func (this *WatchEvent_Upsert) MarshalJSON() ([]byte, error) {
	str, err := ServiceMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for WatchEvent_Upsert
func (this *WatchEvent_Upsert) UnmarshalJSON(b []byte) error {
	return ServiceUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for WatchEvent_Delete
func (this *WatchEvent_Delete) MarshalJSON() ([]byte, error) {
	str, err := ServiceMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for WatchEvent_Delete
func (this *WatchEvent_Delete) UnmarshalJSON(b []byte) error {
	return ServiceUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for WatchEvent_EndOfSnapshot
func (this *WatchEvent_EndOfSnapshot) MarshalJSON() ([]byte, error) {
	str, err := ServiceMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for WatchEvent_EndOfSnapshot
func (this *WatchEvent_EndOfSnapshot) UnmarshalJSON(b []byte) error {
	return ServiceUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for MutateAndValidateRequest
func (this *MutateAndValidateRequest) MarshalJSON() ([]byte, error) {
	str, err := ServiceMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for MutateAndValidateRequest
func (this *MutateAndValidateRequest) UnmarshalJSON(b []byte) error {
	return ServiceUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for MutateAndValidateResponse
func (this *MutateAndValidateResponse) MarshalJSON() ([]byte, error) {
	str, err := ServiceMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for MutateAndValidateResponse
func (this *MutateAndValidateResponse) UnmarshalJSON(b []byte) error {
	return ServiceUnmarshaler.Unmarshal(b, this)
}

var (
	ServiceMarshaler   = &protojson.MarshalOptions{}
	ServiceUnmarshaler = &protojson.UnmarshalOptions{DiscardUnknown: false}
)