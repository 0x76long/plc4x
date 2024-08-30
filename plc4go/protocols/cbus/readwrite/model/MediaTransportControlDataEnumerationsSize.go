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

// MediaTransportControlDataEnumerationsSize is the corresponding interface of MediaTransportControlDataEnumerationsSize
type MediaTransportControlDataEnumerationsSize interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	MediaTransportControlData
	// GetSizeType returns SizeType (property field)
	GetSizeType() byte
	// GetStart returns Start (property field)
	GetStart() uint8
	// GetSize returns Size (property field)
	GetSize() uint8
	// GetIsListCategories returns IsListCategories (virtual field)
	GetIsListCategories() bool
	// GetIsListSelections returns IsListSelections (virtual field)
	GetIsListSelections() bool
	// GetIsListTracks returns IsListTracks (virtual field)
	GetIsListTracks() bool
	// GetIsReserved returns IsReserved (virtual field)
	GetIsReserved() bool
}

// MediaTransportControlDataEnumerationsSizeExactly can be used when we want exactly this type and not a type which fulfills MediaTransportControlDataEnumerationsSize.
// This is useful for switch cases.
type MediaTransportControlDataEnumerationsSizeExactly interface {
	MediaTransportControlDataEnumerationsSize
	isMediaTransportControlDataEnumerationsSize() bool
}

// _MediaTransportControlDataEnumerationsSize is the data-structure of this message
type _MediaTransportControlDataEnumerationsSize struct {
	*_MediaTransportControlData
	SizeType byte
	Start    uint8
	Size     uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MediaTransportControlDataEnumerationsSize) InitializeParent(parent MediaTransportControlData, commandTypeContainer MediaTransportControlCommandTypeContainer, mediaLinkGroup byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.MediaLinkGroup = mediaLinkGroup
}

func (m *_MediaTransportControlDataEnumerationsSize) GetParent() MediaTransportControlData {
	return m._MediaTransportControlData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MediaTransportControlDataEnumerationsSize) GetSizeType() byte {
	return m.SizeType
}

func (m *_MediaTransportControlDataEnumerationsSize) GetStart() uint8 {
	return m.Start
}

func (m *_MediaTransportControlDataEnumerationsSize) GetSize() uint8 {
	return m.Size
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_MediaTransportControlDataEnumerationsSize) GetIsListCategories() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetSizeType()) == (0x00)))
}

func (m *_MediaTransportControlDataEnumerationsSize) GetIsListSelections() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetSizeType()) == (0x01)))
}

func (m *_MediaTransportControlDataEnumerationsSize) GetIsListTracks() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetSizeType()) == (0x02)))
}

func (m *_MediaTransportControlDataEnumerationsSize) GetIsReserved() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool(bool(!(m.GetIsListCategories())) && bool(!(m.GetIsListSelections()))) && bool(!(m.GetIsListTracks())))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewMediaTransportControlDataEnumerationsSize factory function for _MediaTransportControlDataEnumerationsSize
func NewMediaTransportControlDataEnumerationsSize(sizeType byte, start uint8, size uint8, commandTypeContainer MediaTransportControlCommandTypeContainer, mediaLinkGroup byte) *_MediaTransportControlDataEnumerationsSize {
	_result := &_MediaTransportControlDataEnumerationsSize{
		SizeType:                   sizeType,
		Start:                      start,
		Size:                       size,
		_MediaTransportControlData: NewMediaTransportControlData(commandTypeContainer, mediaLinkGroup),
	}
	_result._MediaTransportControlData._MediaTransportControlDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastMediaTransportControlDataEnumerationsSize(structType any) MediaTransportControlDataEnumerationsSize {
	if casted, ok := structType.(MediaTransportControlDataEnumerationsSize); ok {
		return casted
	}
	if casted, ok := structType.(*MediaTransportControlDataEnumerationsSize); ok {
		return *casted
	}
	return nil
}

func (m *_MediaTransportControlDataEnumerationsSize) GetTypeName() string {
	return "MediaTransportControlDataEnumerationsSize"
}

func (m *_MediaTransportControlDataEnumerationsSize) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (sizeType)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// Simple field (start)
	lengthInBits += 8

	// Simple field (size)
	lengthInBits += 8

	return lengthInBits
}

func (m *_MediaTransportControlDataEnumerationsSize) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func MediaTransportControlDataEnumerationsSizeParse(ctx context.Context, theBytes []byte) (MediaTransportControlDataEnumerationsSize, error) {
	return MediaTransportControlDataEnumerationsSizeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func MediaTransportControlDataEnumerationsSizeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (MediaTransportControlDataEnumerationsSize, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("MediaTransportControlDataEnumerationsSize"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MediaTransportControlDataEnumerationsSize")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (sizeType)
	_sizeType, _sizeTypeErr := /*TODO: migrate me*/ readBuffer.ReadByte("sizeType")
	if _sizeTypeErr != nil {
		return nil, errors.Wrap(_sizeTypeErr, "Error parsing 'sizeType' field of MediaTransportControlDataEnumerationsSize")
	}
	sizeType := _sizeType

	// Virtual field
	_isListCategories := bool((sizeType) == (0x00))
	isListCategories := bool(_isListCategories)
	_ = isListCategories

	// Virtual field
	_isListSelections := bool((sizeType) == (0x01))
	isListSelections := bool(_isListSelections)
	_ = isListSelections

	// Virtual field
	_isListTracks := bool((sizeType) == (0x02))
	isListTracks := bool(_isListTracks)
	_ = isListTracks

	// Virtual field
	_isReserved := bool(bool(!(isListCategories)) && bool(!(isListSelections))) && bool(!(isListTracks))
	isReserved := bool(_isReserved)
	_ = isReserved

	// Simple Field (start)
	_start, _startErr := /*TODO: migrate me*/ readBuffer.ReadUint8("start", 8)
	if _startErr != nil {
		return nil, errors.Wrap(_startErr, "Error parsing 'start' field of MediaTransportControlDataEnumerationsSize")
	}
	start := _start

	// Simple Field (size)
	_size, _sizeErr := /*TODO: migrate me*/ readBuffer.ReadUint8("size", 8)
	if _sizeErr != nil {
		return nil, errors.Wrap(_sizeErr, "Error parsing 'size' field of MediaTransportControlDataEnumerationsSize")
	}
	size := _size

	if closeErr := readBuffer.CloseContext("MediaTransportControlDataEnumerationsSize"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MediaTransportControlDataEnumerationsSize")
	}

	// Create a partially initialized instance
	_child := &_MediaTransportControlDataEnumerationsSize{
		_MediaTransportControlData: &_MediaTransportControlData{},
		SizeType:                   sizeType,
		Start:                      start,
		Size:                       size,
	}
	_child._MediaTransportControlData._MediaTransportControlDataChildRequirements = _child
	return _child, nil
}

func (m *_MediaTransportControlDataEnumerationsSize) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MediaTransportControlDataEnumerationsSize) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MediaTransportControlDataEnumerationsSize"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MediaTransportControlDataEnumerationsSize")
		}

		// Simple Field (sizeType)
		sizeType := byte(m.GetSizeType())
		_sizeTypeErr := /*TODO: migrate me*/ writeBuffer.WriteByte("sizeType", (sizeType))
		if _sizeTypeErr != nil {
			return errors.Wrap(_sizeTypeErr, "Error serializing 'sizeType' field")
		}
		// Virtual field
		isListCategories := m.GetIsListCategories()
		_ = isListCategories
		if _isListCategoriesErr := writeBuffer.WriteVirtual(ctx, "isListCategories", m.GetIsListCategories()); _isListCategoriesErr != nil {
			return errors.Wrap(_isListCategoriesErr, "Error serializing 'isListCategories' field")
		}
		// Virtual field
		isListSelections := m.GetIsListSelections()
		_ = isListSelections
		if _isListSelectionsErr := writeBuffer.WriteVirtual(ctx, "isListSelections", m.GetIsListSelections()); _isListSelectionsErr != nil {
			return errors.Wrap(_isListSelectionsErr, "Error serializing 'isListSelections' field")
		}
		// Virtual field
		isListTracks := m.GetIsListTracks()
		_ = isListTracks
		if _isListTracksErr := writeBuffer.WriteVirtual(ctx, "isListTracks", m.GetIsListTracks()); _isListTracksErr != nil {
			return errors.Wrap(_isListTracksErr, "Error serializing 'isListTracks' field")
		}
		// Virtual field
		isReserved := m.GetIsReserved()
		_ = isReserved
		if _isReservedErr := writeBuffer.WriteVirtual(ctx, "isReserved", m.GetIsReserved()); _isReservedErr != nil {
			return errors.Wrap(_isReservedErr, "Error serializing 'isReserved' field")
		}

		// Simple Field (start)
		start := uint8(m.GetStart())
		_startErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("start", 8, uint8((start)))
		if _startErr != nil {
			return errors.Wrap(_startErr, "Error serializing 'start' field")
		}

		// Simple Field (size)
		size := uint8(m.GetSize())
		_sizeErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("size", 8, uint8((size)))
		if _sizeErr != nil {
			return errors.Wrap(_sizeErr, "Error serializing 'size' field")
		}

		if popErr := writeBuffer.PopContext("MediaTransportControlDataEnumerationsSize"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MediaTransportControlDataEnumerationsSize")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MediaTransportControlDataEnumerationsSize) isMediaTransportControlDataEnumerationsSize() bool {
	return true
}

func (m *_MediaTransportControlDataEnumerationsSize) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
