// Code generated by protoc-gen-go-binary. DO NOT EDIT.
// source: pbhcp/v1/hcp_link.proto

package hcpv1

import (
	"google.golang.org/protobuf/proto"
)

// MarshalBinary implements encoding.BinaryMarshaler
func (msg *LinkConfiguration) MarshalBinary() ([]byte, error) {
	return proto.Marshal(msg)
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler
func (msg *LinkConfiguration) UnmarshalBinary(b []byte) error {
	return proto.Unmarshal(b, msg)
}
