// Code generated by protoc-gen-go-binary. DO NOT EDIT.
// source: pbcatalog/v1alpha1/mesh_gateway.proto

package catalogv1alpha1

import (
	"google.golang.org/protobuf/proto"
)

// MarshalBinary implements encoding.BinaryMarshaler
func (msg *MeshGateway) MarshalBinary() ([]byte, error) {
	return proto.Marshal(msg)
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler
func (msg *MeshGateway) UnmarshalBinary(b []byte) error {
	return proto.Unmarshal(b, msg)
}