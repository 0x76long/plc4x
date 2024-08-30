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

// FirmataCommandProtocolVersion is the corresponding interface of FirmataCommandProtocolVersion
type FirmataCommandProtocolVersion interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	FirmataCommand
	// GetMajorVersion returns MajorVersion (property field)
	GetMajorVersion() uint8
	// GetMinorVersion returns MinorVersion (property field)
	GetMinorVersion() uint8
}

// FirmataCommandProtocolVersionExactly can be used when we want exactly this type and not a type which fulfills FirmataCommandProtocolVersion.
// This is useful for switch cases.
type FirmataCommandProtocolVersionExactly interface {
	FirmataCommandProtocolVersion
	isFirmataCommandProtocolVersion() bool
}

// _FirmataCommandProtocolVersion is the data-structure of this message
type _FirmataCommandProtocolVersion struct {
	*_FirmataCommand
	MajorVersion uint8
	MinorVersion uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_FirmataCommandProtocolVersion) GetCommandCode() uint8 {
	return 0x9
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_FirmataCommandProtocolVersion) InitializeParent(parent FirmataCommand) {}

func (m *_FirmataCommandProtocolVersion) GetParent() FirmataCommand {
	return m._FirmataCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_FirmataCommandProtocolVersion) GetMajorVersion() uint8 {
	return m.MajorVersion
}

func (m *_FirmataCommandProtocolVersion) GetMinorVersion() uint8 {
	return m.MinorVersion
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewFirmataCommandProtocolVersion factory function for _FirmataCommandProtocolVersion
func NewFirmataCommandProtocolVersion(majorVersion uint8, minorVersion uint8, response bool) *_FirmataCommandProtocolVersion {
	_result := &_FirmataCommandProtocolVersion{
		MajorVersion:    majorVersion,
		MinorVersion:    minorVersion,
		_FirmataCommand: NewFirmataCommand(response),
	}
	_result._FirmataCommand._FirmataCommandChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastFirmataCommandProtocolVersion(structType any) FirmataCommandProtocolVersion {
	if casted, ok := structType.(FirmataCommandProtocolVersion); ok {
		return casted
	}
	if casted, ok := structType.(*FirmataCommandProtocolVersion); ok {
		return *casted
	}
	return nil
}

func (m *_FirmataCommandProtocolVersion) GetTypeName() string {
	return "FirmataCommandProtocolVersion"
}

func (m *_FirmataCommandProtocolVersion) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (majorVersion)
	lengthInBits += 8

	// Simple field (minorVersion)
	lengthInBits += 8

	return lengthInBits
}

func (m *_FirmataCommandProtocolVersion) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func FirmataCommandProtocolVersionParse(ctx context.Context, theBytes []byte, response bool) (FirmataCommandProtocolVersion, error) {
	return FirmataCommandProtocolVersionParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), response)
}

func FirmataCommandProtocolVersionParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, response bool) (FirmataCommandProtocolVersion, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("FirmataCommandProtocolVersion"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for FirmataCommandProtocolVersion")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (majorVersion)
	_majorVersion, _majorVersionErr := /*TODO: migrate me*/ readBuffer.ReadUint8("majorVersion", 8)
	if _majorVersionErr != nil {
		return nil, errors.Wrap(_majorVersionErr, "Error parsing 'majorVersion' field of FirmataCommandProtocolVersion")
	}
	majorVersion := _majorVersion

	// Simple Field (minorVersion)
	_minorVersion, _minorVersionErr := /*TODO: migrate me*/ readBuffer.ReadUint8("minorVersion", 8)
	if _minorVersionErr != nil {
		return nil, errors.Wrap(_minorVersionErr, "Error parsing 'minorVersion' field of FirmataCommandProtocolVersion")
	}
	minorVersion := _minorVersion

	if closeErr := readBuffer.CloseContext("FirmataCommandProtocolVersion"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for FirmataCommandProtocolVersion")
	}

	// Create a partially initialized instance
	_child := &_FirmataCommandProtocolVersion{
		_FirmataCommand: &_FirmataCommand{
			Response: response,
		},
		MajorVersion: majorVersion,
		MinorVersion: minorVersion,
	}
	_child._FirmataCommand._FirmataCommandChildRequirements = _child
	return _child, nil
}

func (m *_FirmataCommandProtocolVersion) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_FirmataCommandProtocolVersion) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("FirmataCommandProtocolVersion"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for FirmataCommandProtocolVersion")
		}

		// Simple Field (majorVersion)
		majorVersion := uint8(m.GetMajorVersion())
		_majorVersionErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("majorVersion", 8, uint8((majorVersion)))
		if _majorVersionErr != nil {
			return errors.Wrap(_majorVersionErr, "Error serializing 'majorVersion' field")
		}

		// Simple Field (minorVersion)
		minorVersion := uint8(m.GetMinorVersion())
		_minorVersionErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("minorVersion", 8, uint8((minorVersion)))
		if _minorVersionErr != nil {
			return errors.Wrap(_minorVersionErr, "Error serializing 'minorVersion' field")
		}

		if popErr := writeBuffer.PopContext("FirmataCommandProtocolVersion"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for FirmataCommandProtocolVersion")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_FirmataCommandProtocolVersion) isFirmataCommandProtocolVersion() bool {
	return true
}

func (m *_FirmataCommandProtocolVersion) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
