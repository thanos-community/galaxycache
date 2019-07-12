/*
Copyright 2012 Google Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package galaxycache

import (
	"github.com/golang/protobuf/proto"
)

// Codec includes both the BinaryMarshaler and BinaryUnmarshaler
// interfaces
type Codec interface {
	MarshalBinary() ([]byte, error)
	UnmarshalBinary(data []byte) error
}

// Note: to ensure that unmarshaling is a read-only operation, bytes
// are always cloned
func cloneBytes(b []byte) []byte {
	tmp := make([]byte, len(b))
	copy(tmp, b)
	return tmp
}

// ByteCodec is a byte slice type that implements Codec
type ByteCodec []byte

// MarshalBinary on a ByteCodec returns the bytes
func (c *ByteCodec) MarshalBinary() ([]byte, error) {
	return *c, nil
}

// UnmarshalBinary on a ByteCodec sets the ByteCodec to
// a copy of the provided data
func (c *ByteCodec) UnmarshalBinary(data []byte) error {
	*c = cloneBytes(data)
	return nil
}

// StringCodec is a string type that implements Codec
type StringCodec string

// MarshalBinary on a StringCodec returns the bytes underlying
// the string
func (c *StringCodec) MarshalBinary() ([]byte, error) {
	return []byte(*c), nil
}

// UnmarshalBinary on a StringCodec sets the StringCodec to
// a stringified copy of the provided data
func (c *StringCodec) UnmarshalBinary(data []byte) error {
	*c = StringCodec(cloneBytes(data))
	return nil
}

// ProtoCodec wraps a proto.Message and implements Codec
type ProtoCodec struct {
	Msg proto.Message
}

// MarshalBinary on a ProtoCodec returns the encoded proto message
func (c *ProtoCodec) MarshalBinary() ([]byte, error) {
	return proto.Marshal(c.Msg)
}

// UnmarshalBinary on a ProtoCodec unmarshals provided data into
// the proto message
func (c *ProtoCodec) UnmarshalBinary(data []byte) error {
	return proto.Unmarshal(data, c.Msg)
}
