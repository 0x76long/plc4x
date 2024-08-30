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

// DF1RequestProtectedTypedLogicalRead is the corresponding interface of DF1RequestProtectedTypedLogicalRead
type DF1RequestProtectedTypedLogicalRead interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	DF1RequestCommand
	// GetByteSize returns ByteSize (property field)
	GetByteSize() uint8
	// GetFileNumber returns FileNumber (property field)
	GetFileNumber() uint8
	// GetFileType returns FileType (property field)
	GetFileType() uint8
	// GetElementNumber returns ElementNumber (property field)
	GetElementNumber() uint8
	// GetSubElementNumber returns SubElementNumber (property field)
	GetSubElementNumber() uint8
}

// DF1RequestProtectedTypedLogicalReadExactly can be used when we want exactly this type and not a type which fulfills DF1RequestProtectedTypedLogicalRead.
// This is useful for switch cases.
type DF1RequestProtectedTypedLogicalReadExactly interface {
	DF1RequestProtectedTypedLogicalRead
	isDF1RequestProtectedTypedLogicalRead() bool
}

// _DF1RequestProtectedTypedLogicalRead is the data-structure of this message
type _DF1RequestProtectedTypedLogicalRead struct {
	*_DF1RequestCommand
	ByteSize         uint8
	FileNumber       uint8
	FileType         uint8
	ElementNumber    uint8
	SubElementNumber uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DF1RequestProtectedTypedLogicalRead) GetFunctionCode() uint8 {
	return 0xA2
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DF1RequestProtectedTypedLogicalRead) InitializeParent(parent DF1RequestCommand) {}

func (m *_DF1RequestProtectedTypedLogicalRead) GetParent() DF1RequestCommand {
	return m._DF1RequestCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DF1RequestProtectedTypedLogicalRead) GetByteSize() uint8 {
	return m.ByteSize
}

func (m *_DF1RequestProtectedTypedLogicalRead) GetFileNumber() uint8 {
	return m.FileNumber
}

func (m *_DF1RequestProtectedTypedLogicalRead) GetFileType() uint8 {
	return m.FileType
}

func (m *_DF1RequestProtectedTypedLogicalRead) GetElementNumber() uint8 {
	return m.ElementNumber
}

func (m *_DF1RequestProtectedTypedLogicalRead) GetSubElementNumber() uint8 {
	return m.SubElementNumber
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewDF1RequestProtectedTypedLogicalRead factory function for _DF1RequestProtectedTypedLogicalRead
func NewDF1RequestProtectedTypedLogicalRead(byteSize uint8, fileNumber uint8, fileType uint8, elementNumber uint8, subElementNumber uint8) *_DF1RequestProtectedTypedLogicalRead {
	_result := &_DF1RequestProtectedTypedLogicalRead{
		ByteSize:           byteSize,
		FileNumber:         fileNumber,
		FileType:           fileType,
		ElementNumber:      elementNumber,
		SubElementNumber:   subElementNumber,
		_DF1RequestCommand: NewDF1RequestCommand(),
	}
	_result._DF1RequestCommand._DF1RequestCommandChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastDF1RequestProtectedTypedLogicalRead(structType any) DF1RequestProtectedTypedLogicalRead {
	if casted, ok := structType.(DF1RequestProtectedTypedLogicalRead); ok {
		return casted
	}
	if casted, ok := structType.(*DF1RequestProtectedTypedLogicalRead); ok {
		return *casted
	}
	return nil
}

func (m *_DF1RequestProtectedTypedLogicalRead) GetTypeName() string {
	return "DF1RequestProtectedTypedLogicalRead"
}

func (m *_DF1RequestProtectedTypedLogicalRead) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (byteSize)
	lengthInBits += 8

	// Simple field (fileNumber)
	lengthInBits += 8

	// Simple field (fileType)
	lengthInBits += 8

	// Simple field (elementNumber)
	lengthInBits += 8

	// Simple field (subElementNumber)
	lengthInBits += 8

	return lengthInBits
}

func (m *_DF1RequestProtectedTypedLogicalRead) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func DF1RequestProtectedTypedLogicalReadParse(ctx context.Context, theBytes []byte) (DF1RequestProtectedTypedLogicalRead, error) {
	return DF1RequestProtectedTypedLogicalReadParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func DF1RequestProtectedTypedLogicalReadParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (DF1RequestProtectedTypedLogicalRead, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("DF1RequestProtectedTypedLogicalRead"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DF1RequestProtectedTypedLogicalRead")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (byteSize)
	_byteSize, _byteSizeErr := /*TODO: migrate me*/ readBuffer.ReadUint8("byteSize", 8)
	if _byteSizeErr != nil {
		return nil, errors.Wrap(_byteSizeErr, "Error parsing 'byteSize' field of DF1RequestProtectedTypedLogicalRead")
	}
	byteSize := _byteSize

	// Simple Field (fileNumber)
	_fileNumber, _fileNumberErr := /*TODO: migrate me*/ readBuffer.ReadUint8("fileNumber", 8)
	if _fileNumberErr != nil {
		return nil, errors.Wrap(_fileNumberErr, "Error parsing 'fileNumber' field of DF1RequestProtectedTypedLogicalRead")
	}
	fileNumber := _fileNumber

	// Simple Field (fileType)
	_fileType, _fileTypeErr := /*TODO: migrate me*/ readBuffer.ReadUint8("fileType", 8)
	if _fileTypeErr != nil {
		return nil, errors.Wrap(_fileTypeErr, "Error parsing 'fileType' field of DF1RequestProtectedTypedLogicalRead")
	}
	fileType := _fileType

	// Simple Field (elementNumber)
	_elementNumber, _elementNumberErr := /*TODO: migrate me*/ readBuffer.ReadUint8("elementNumber", 8)
	if _elementNumberErr != nil {
		return nil, errors.Wrap(_elementNumberErr, "Error parsing 'elementNumber' field of DF1RequestProtectedTypedLogicalRead")
	}
	elementNumber := _elementNumber

	// Simple Field (subElementNumber)
	_subElementNumber, _subElementNumberErr := /*TODO: migrate me*/ readBuffer.ReadUint8("subElementNumber", 8)
	if _subElementNumberErr != nil {
		return nil, errors.Wrap(_subElementNumberErr, "Error parsing 'subElementNumber' field of DF1RequestProtectedTypedLogicalRead")
	}
	subElementNumber := _subElementNumber

	if closeErr := readBuffer.CloseContext("DF1RequestProtectedTypedLogicalRead"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DF1RequestProtectedTypedLogicalRead")
	}

	// Create a partially initialized instance
	_child := &_DF1RequestProtectedTypedLogicalRead{
		_DF1RequestCommand: &_DF1RequestCommand{},
		ByteSize:           byteSize,
		FileNumber:         fileNumber,
		FileType:           fileType,
		ElementNumber:      elementNumber,
		SubElementNumber:   subElementNumber,
	}
	_child._DF1RequestCommand._DF1RequestCommandChildRequirements = _child
	return _child, nil
}

func (m *_DF1RequestProtectedTypedLogicalRead) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DF1RequestProtectedTypedLogicalRead) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DF1RequestProtectedTypedLogicalRead"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DF1RequestProtectedTypedLogicalRead")
		}

		// Simple Field (byteSize)
		byteSize := uint8(m.GetByteSize())
		_byteSizeErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("byteSize", 8, uint8((byteSize)))
		if _byteSizeErr != nil {
			return errors.Wrap(_byteSizeErr, "Error serializing 'byteSize' field")
		}

		// Simple Field (fileNumber)
		fileNumber := uint8(m.GetFileNumber())
		_fileNumberErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("fileNumber", 8, uint8((fileNumber)))
		if _fileNumberErr != nil {
			return errors.Wrap(_fileNumberErr, "Error serializing 'fileNumber' field")
		}

		// Simple Field (fileType)
		fileType := uint8(m.GetFileType())
		_fileTypeErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("fileType", 8, uint8((fileType)))
		if _fileTypeErr != nil {
			return errors.Wrap(_fileTypeErr, "Error serializing 'fileType' field")
		}

		// Simple Field (elementNumber)
		elementNumber := uint8(m.GetElementNumber())
		_elementNumberErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("elementNumber", 8, uint8((elementNumber)))
		if _elementNumberErr != nil {
			return errors.Wrap(_elementNumberErr, "Error serializing 'elementNumber' field")
		}

		// Simple Field (subElementNumber)
		subElementNumber := uint8(m.GetSubElementNumber())
		_subElementNumberErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("subElementNumber", 8, uint8((subElementNumber)))
		if _subElementNumberErr != nil {
			return errors.Wrap(_subElementNumberErr, "Error serializing 'subElementNumber' field")
		}

		if popErr := writeBuffer.PopContext("DF1RequestProtectedTypedLogicalRead"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DF1RequestProtectedTypedLogicalRead")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DF1RequestProtectedTypedLogicalRead) isDF1RequestProtectedTypedLogicalRead() bool {
	return true
}

func (m *_DF1RequestProtectedTypedLogicalRead) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
