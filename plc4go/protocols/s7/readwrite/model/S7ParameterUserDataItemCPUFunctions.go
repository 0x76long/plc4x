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
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// S7ParameterUserDataItemCPUFunctions is the corresponding interface of S7ParameterUserDataItemCPUFunctions
type S7ParameterUserDataItemCPUFunctions interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	S7ParameterUserDataItem
	// GetMethod returns Method (property field)
	GetMethod() uint8
	// GetCpuFunctionType returns CpuFunctionType (property field)
	GetCpuFunctionType() uint8
	// GetCpuFunctionGroup returns CpuFunctionGroup (property field)
	GetCpuFunctionGroup() uint8
	// GetCpuSubfunction returns CpuSubfunction (property field)
	GetCpuSubfunction() uint8
	// GetSequenceNumber returns SequenceNumber (property field)
	GetSequenceNumber() uint8
	// GetDataUnitReferenceNumber returns DataUnitReferenceNumber (property field)
	GetDataUnitReferenceNumber() *uint8
	// GetLastDataUnit returns LastDataUnit (property field)
	GetLastDataUnit() *uint8
	// GetErrorCode returns ErrorCode (property field)
	GetErrorCode() *uint16
}

// S7ParameterUserDataItemCPUFunctionsExactly can be used when we want exactly this type and not a type which fulfills S7ParameterUserDataItemCPUFunctions.
// This is useful for switch cases.
type S7ParameterUserDataItemCPUFunctionsExactly interface {
	S7ParameterUserDataItemCPUFunctions
	isS7ParameterUserDataItemCPUFunctions() bool
}

// _S7ParameterUserDataItemCPUFunctions is the data-structure of this message
type _S7ParameterUserDataItemCPUFunctions struct {
	*_S7ParameterUserDataItem
	Method                  uint8
	CpuFunctionType         uint8
	CpuFunctionGroup        uint8
	CpuSubfunction          uint8
	SequenceNumber          uint8
	DataUnitReferenceNumber *uint8
	LastDataUnit            *uint8
	ErrorCode               *uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7ParameterUserDataItemCPUFunctions) GetItemType() uint8 {
	return 0x12
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7ParameterUserDataItemCPUFunctions) InitializeParent(parent S7ParameterUserDataItem) {}

func (m *_S7ParameterUserDataItemCPUFunctions) GetParent() S7ParameterUserDataItem {
	return m._S7ParameterUserDataItem
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7ParameterUserDataItemCPUFunctions) GetMethod() uint8 {
	return m.Method
}

func (m *_S7ParameterUserDataItemCPUFunctions) GetCpuFunctionType() uint8 {
	return m.CpuFunctionType
}

func (m *_S7ParameterUserDataItemCPUFunctions) GetCpuFunctionGroup() uint8 {
	return m.CpuFunctionGroup
}

func (m *_S7ParameterUserDataItemCPUFunctions) GetCpuSubfunction() uint8 {
	return m.CpuSubfunction
}

func (m *_S7ParameterUserDataItemCPUFunctions) GetSequenceNumber() uint8 {
	return m.SequenceNumber
}

func (m *_S7ParameterUserDataItemCPUFunctions) GetDataUnitReferenceNumber() *uint8 {
	return m.DataUnitReferenceNumber
}

func (m *_S7ParameterUserDataItemCPUFunctions) GetLastDataUnit() *uint8 {
	return m.LastDataUnit
}

func (m *_S7ParameterUserDataItemCPUFunctions) GetErrorCode() *uint16 {
	return m.ErrorCode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7ParameterUserDataItemCPUFunctions factory function for _S7ParameterUserDataItemCPUFunctions
func NewS7ParameterUserDataItemCPUFunctions(method uint8, cpuFunctionType uint8, cpuFunctionGroup uint8, cpuSubfunction uint8, sequenceNumber uint8, dataUnitReferenceNumber *uint8, lastDataUnit *uint8, errorCode *uint16) *_S7ParameterUserDataItemCPUFunctions {
	_result := &_S7ParameterUserDataItemCPUFunctions{
		Method:                   method,
		CpuFunctionType:          cpuFunctionType,
		CpuFunctionGroup:         cpuFunctionGroup,
		CpuSubfunction:           cpuSubfunction,
		SequenceNumber:           sequenceNumber,
		DataUnitReferenceNumber:  dataUnitReferenceNumber,
		LastDataUnit:             lastDataUnit,
		ErrorCode:                errorCode,
		_S7ParameterUserDataItem: NewS7ParameterUserDataItem(),
	}
	_result._S7ParameterUserDataItem._S7ParameterUserDataItemChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastS7ParameterUserDataItemCPUFunctions(structType any) S7ParameterUserDataItemCPUFunctions {
	if casted, ok := structType.(S7ParameterUserDataItemCPUFunctions); ok {
		return casted
	}
	if casted, ok := structType.(*S7ParameterUserDataItemCPUFunctions); ok {
		return *casted
	}
	return nil
}

func (m *_S7ParameterUserDataItemCPUFunctions) GetTypeName() string {
	return "S7ParameterUserDataItemCPUFunctions"
}

func (m *_S7ParameterUserDataItemCPUFunctions) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Implicit Field (itemLength)
	lengthInBits += 8

	// Simple field (method)
	lengthInBits += 8

	// Simple field (cpuFunctionType)
	lengthInBits += 4

	// Simple field (cpuFunctionGroup)
	lengthInBits += 4

	// Simple field (cpuSubfunction)
	lengthInBits += 8

	// Simple field (sequenceNumber)
	lengthInBits += 8

	// Optional Field (dataUnitReferenceNumber)
	if m.DataUnitReferenceNumber != nil {
		lengthInBits += 8
	}

	// Optional Field (lastDataUnit)
	if m.LastDataUnit != nil {
		lengthInBits += 8
	}

	// Optional Field (errorCode)
	if m.ErrorCode != nil {
		lengthInBits += 16
	}

	return lengthInBits
}

func (m *_S7ParameterUserDataItemCPUFunctions) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func S7ParameterUserDataItemCPUFunctionsParse(ctx context.Context, theBytes []byte) (S7ParameterUserDataItemCPUFunctions, error) {
	return S7ParameterUserDataItemCPUFunctionsParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func S7ParameterUserDataItemCPUFunctionsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (S7ParameterUserDataItemCPUFunctions, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("S7ParameterUserDataItemCPUFunctions"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7ParameterUserDataItemCPUFunctions")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Implicit Field (itemLength) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	itemLength, _itemLengthErr := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint8("itemLength", 8)
	_ = itemLength
	if _itemLengthErr != nil {
		return nil, errors.Wrap(_itemLengthErr, "Error parsing 'itemLength' field of S7ParameterUserDataItemCPUFunctions")
	}

	// Simple Field (method)
	_method, _methodErr := /*TODO: migrate me*/ readBuffer.ReadUint8("method", 8)
	if _methodErr != nil {
		return nil, errors.Wrap(_methodErr, "Error parsing 'method' field of S7ParameterUserDataItemCPUFunctions")
	}
	method := _method

	// Simple Field (cpuFunctionType)
	_cpuFunctionType, _cpuFunctionTypeErr := /*TODO: migrate me*/ readBuffer.ReadUint8("cpuFunctionType", 4)
	if _cpuFunctionTypeErr != nil {
		return nil, errors.Wrap(_cpuFunctionTypeErr, "Error parsing 'cpuFunctionType' field of S7ParameterUserDataItemCPUFunctions")
	}
	cpuFunctionType := _cpuFunctionType

	// Simple Field (cpuFunctionGroup)
	_cpuFunctionGroup, _cpuFunctionGroupErr := /*TODO: migrate me*/ readBuffer.ReadUint8("cpuFunctionGroup", 4)
	if _cpuFunctionGroupErr != nil {
		return nil, errors.Wrap(_cpuFunctionGroupErr, "Error parsing 'cpuFunctionGroup' field of S7ParameterUserDataItemCPUFunctions")
	}
	cpuFunctionGroup := _cpuFunctionGroup

	// Simple Field (cpuSubfunction)
	_cpuSubfunction, _cpuSubfunctionErr := /*TODO: migrate me*/ readBuffer.ReadUint8("cpuSubfunction", 8)
	if _cpuSubfunctionErr != nil {
		return nil, errors.Wrap(_cpuSubfunctionErr, "Error parsing 'cpuSubfunction' field of S7ParameterUserDataItemCPUFunctions")
	}
	cpuSubfunction := _cpuSubfunction

	// Simple Field (sequenceNumber)
	_sequenceNumber, _sequenceNumberErr := /*TODO: migrate me*/ readBuffer.ReadUint8("sequenceNumber", 8)
	if _sequenceNumberErr != nil {
		return nil, errors.Wrap(_sequenceNumberErr, "Error parsing 'sequenceNumber' field of S7ParameterUserDataItemCPUFunctions")
	}
	sequenceNumber := _sequenceNumber

	// Optional Field (dataUnitReferenceNumber) (Can be skipped, if a given expression evaluates to false)
	var dataUnitReferenceNumber *uint8 = nil
	if bool((bool((cpuFunctionType) == (8)))) || bool((bool((bool((cpuFunctionType) == (0)))) && bool((bool((cpuFunctionGroup) == (2)))))) {
		currentPos = positionAware.GetPos()
		_val, _err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint8("dataUnitReferenceNumber", 8)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'dataUnitReferenceNumber' field of S7ParameterUserDataItemCPUFunctions")
		default:
			dataUnitReferenceNumber = &_val
		}
	}

	// Optional Field (lastDataUnit) (Can be skipped, if a given expression evaluates to false)
	var lastDataUnit *uint8 = nil
	if bool((bool((cpuFunctionType) == (8)))) || bool((bool((bool((cpuFunctionType) == (0)))) && bool((bool((cpuFunctionGroup) == (2)))))) {
		currentPos = positionAware.GetPos()
		_val, _err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint8("lastDataUnit", 8)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'lastDataUnit' field of S7ParameterUserDataItemCPUFunctions")
		default:
			lastDataUnit = &_val
		}
	}

	// Optional Field (errorCode) (Can be skipped, if a given expression evaluates to false)
	var errorCode *uint16 = nil
	if bool((bool((cpuFunctionType) == (8)))) || bool((bool((bool((cpuFunctionType) == (0)))) && bool((bool((cpuFunctionGroup) == (2)))))) {
		currentPos = positionAware.GetPos()
		_val, _err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint16("errorCode", 16)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'errorCode' field of S7ParameterUserDataItemCPUFunctions")
		default:
			errorCode = &_val
		}
	}

	if closeErr := readBuffer.CloseContext("S7ParameterUserDataItemCPUFunctions"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7ParameterUserDataItemCPUFunctions")
	}

	// Create a partially initialized instance
	_child := &_S7ParameterUserDataItemCPUFunctions{
		_S7ParameterUserDataItem: &_S7ParameterUserDataItem{},
		Method:                   method,
		CpuFunctionType:          cpuFunctionType,
		CpuFunctionGroup:         cpuFunctionGroup,
		CpuSubfunction:           cpuSubfunction,
		SequenceNumber:           sequenceNumber,
		DataUnitReferenceNumber:  dataUnitReferenceNumber,
		LastDataUnit:             lastDataUnit,
		ErrorCode:                errorCode,
	}
	_child._S7ParameterUserDataItem._S7ParameterUserDataItemChildRequirements = _child
	return _child, nil
}

func (m *_S7ParameterUserDataItemCPUFunctions) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7ParameterUserDataItemCPUFunctions) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7ParameterUserDataItemCPUFunctions"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7ParameterUserDataItemCPUFunctions")
		}

		// Implicit Field (itemLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		itemLength := uint8(uint8(uint8(m.GetLengthInBytes(ctx))) - uint8(uint8(2)))
		_itemLengthErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("itemLength", 8, uint8((itemLength)))
		if _itemLengthErr != nil {
			return errors.Wrap(_itemLengthErr, "Error serializing 'itemLength' field")
		}

		// Simple Field (method)
		method := uint8(m.GetMethod())
		_methodErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("method", 8, uint8((method)))
		if _methodErr != nil {
			return errors.Wrap(_methodErr, "Error serializing 'method' field")
		}

		// Simple Field (cpuFunctionType)
		cpuFunctionType := uint8(m.GetCpuFunctionType())
		_cpuFunctionTypeErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("cpuFunctionType", 4, uint8((cpuFunctionType)))
		if _cpuFunctionTypeErr != nil {
			return errors.Wrap(_cpuFunctionTypeErr, "Error serializing 'cpuFunctionType' field")
		}

		// Simple Field (cpuFunctionGroup)
		cpuFunctionGroup := uint8(m.GetCpuFunctionGroup())
		_cpuFunctionGroupErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("cpuFunctionGroup", 4, uint8((cpuFunctionGroup)))
		if _cpuFunctionGroupErr != nil {
			return errors.Wrap(_cpuFunctionGroupErr, "Error serializing 'cpuFunctionGroup' field")
		}

		// Simple Field (cpuSubfunction)
		cpuSubfunction := uint8(m.GetCpuSubfunction())
		_cpuSubfunctionErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("cpuSubfunction", 8, uint8((cpuSubfunction)))
		if _cpuSubfunctionErr != nil {
			return errors.Wrap(_cpuSubfunctionErr, "Error serializing 'cpuSubfunction' field")
		}

		// Simple Field (sequenceNumber)
		sequenceNumber := uint8(m.GetSequenceNumber())
		_sequenceNumberErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("sequenceNumber", 8, uint8((sequenceNumber)))
		if _sequenceNumberErr != nil {
			return errors.Wrap(_sequenceNumberErr, "Error serializing 'sequenceNumber' field")
		}

		// Optional Field (dataUnitReferenceNumber) (Can be skipped, if the value is null)
		var dataUnitReferenceNumber *uint8 = nil
		if m.GetDataUnitReferenceNumber() != nil {
			dataUnitReferenceNumber = m.GetDataUnitReferenceNumber()
			_dataUnitReferenceNumberErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("dataUnitReferenceNumber", 8, uint8(*(dataUnitReferenceNumber)))
			if _dataUnitReferenceNumberErr != nil {
				return errors.Wrap(_dataUnitReferenceNumberErr, "Error serializing 'dataUnitReferenceNumber' field")
			}
		}

		// Optional Field (lastDataUnit) (Can be skipped, if the value is null)
		var lastDataUnit *uint8 = nil
		if m.GetLastDataUnit() != nil {
			lastDataUnit = m.GetLastDataUnit()
			_lastDataUnitErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("lastDataUnit", 8, uint8(*(lastDataUnit)))
			if _lastDataUnitErr != nil {
				return errors.Wrap(_lastDataUnitErr, "Error serializing 'lastDataUnit' field")
			}
		}

		// Optional Field (errorCode) (Can be skipped, if the value is null)
		var errorCode *uint16 = nil
		if m.GetErrorCode() != nil {
			errorCode = m.GetErrorCode()
			_errorCodeErr := /*TODO: migrate me*/ writeBuffer.WriteUint16("errorCode", 16, uint16(*(errorCode)))
			if _errorCodeErr != nil {
				return errors.Wrap(_errorCodeErr, "Error serializing 'errorCode' field")
			}
		}

		if popErr := writeBuffer.PopContext("S7ParameterUserDataItemCPUFunctions"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7ParameterUserDataItemCPUFunctions")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7ParameterUserDataItemCPUFunctions) isS7ParameterUserDataItemCPUFunctions() bool {
	return true
}

func (m *_S7ParameterUserDataItemCPUFunctions) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
