/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"context"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// InstanceID is the corresponding interface of InstanceID
type InstanceID interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	LogicalSegmentType
	// GetFormat returns Format (property field)
	GetFormat() uint8
	// GetInstance returns Instance (property field)
	GetInstance() uint8
}

// InstanceIDExactly can be used when we want exactly this type and not a type which fulfills InstanceID.
// This is useful for switch cases.
type InstanceIDExactly interface {
	InstanceID
	isInstanceID() bool
}

// _InstanceID is the data-structure of this message
type _InstanceID struct {
	*_LogicalSegmentType
	Format   uint8
	Instance uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_InstanceID) GetLogicalSegmentType() uint8 {
	return 0x01
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_InstanceID) InitializeParent(parent LogicalSegmentType) {}

func (m *_InstanceID) GetParent() LogicalSegmentType {
	return m._LogicalSegmentType
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_InstanceID) GetFormat() uint8 {
	return m.Format
}

func (m *_InstanceID) GetInstance() uint8 {
	return m.Instance
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewInstanceID factory function for _InstanceID
func NewInstanceID(format uint8, instance uint8) *_InstanceID {
	_result := &_InstanceID{
		Format:              format,
		Instance:            instance,
		_LogicalSegmentType: NewLogicalSegmentType(),
	}
	_result._LogicalSegmentType._LogicalSegmentTypeChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastInstanceID(structType any) InstanceID {
	if casted, ok := structType.(InstanceID); ok {
		return casted
	}
	if casted, ok := structType.(*InstanceID); ok {
		return *casted
	}
	return nil
}

func (m *_InstanceID) GetTypeName() string {
	return "InstanceID"
}

func (m *_InstanceID) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (format)
	lengthInBits += 2

	// Simple field (instance)
	lengthInBits += 8

	return lengthInBits
}

func (m *_InstanceID) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func InstanceIDParse(ctx context.Context, theBytes []byte) (InstanceID, error) {
	return InstanceIDParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func InstanceIDParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (InstanceID, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("InstanceID"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for InstanceID")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (format)
	_format, _formatErr := /*TODO: migrate me*/ readBuffer.ReadUint8("format", 2)
	if _formatErr != nil {
		return nil, errors.Wrap(_formatErr, "Error parsing 'format' field of InstanceID")
	}
	format := _format

	// Simple Field (instance)
	_instance, _instanceErr := /*TODO: migrate me*/ readBuffer.ReadUint8("instance", 8)
	if _instanceErr != nil {
		return nil, errors.Wrap(_instanceErr, "Error parsing 'instance' field of InstanceID")
	}
	instance := _instance

	if closeErr := readBuffer.CloseContext("InstanceID"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for InstanceID")
	}

	// Create a partially initialized instance
	_child := &_InstanceID{
		_LogicalSegmentType: &_LogicalSegmentType{},
		Format:              format,
		Instance:            instance,
	}
	_child._LogicalSegmentType._LogicalSegmentTypeChildRequirements = _child
	return _child, nil
}

func (m *_InstanceID) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_InstanceID) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("InstanceID"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for InstanceID")
		}

		// Simple Field (format)
		format := uint8(m.GetFormat())
		_formatErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("format", 2, uint8((format)))
		if _formatErr != nil {
			return errors.Wrap(_formatErr, "Error serializing 'format' field")
		}

		// Simple Field (instance)
		instance := uint8(m.GetInstance())
		_instanceErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("instance", 8, uint8((instance)))
		if _instanceErr != nil {
			return errors.Wrap(_instanceErr, "Error serializing 'instance' field")
		}

		if popErr := writeBuffer.PopContext("InstanceID"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for InstanceID")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_InstanceID) isInstanceID() bool {
	return true
}

func (m *_InstanceID) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
