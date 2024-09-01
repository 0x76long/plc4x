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

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// TriggerControlLabelOptions is the corresponding interface of TriggerControlLabelOptions
type TriggerControlLabelOptions interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetLabelFlavour returns LabelFlavour (property field)
	GetLabelFlavour() TriggerControlLabelFlavour
	// GetLabelType returns LabelType (property field)
	GetLabelType() TriggerControlLabelType
}

// TriggerControlLabelOptionsExactly can be used when we want exactly this type and not a type which fulfills TriggerControlLabelOptions.
// This is useful for switch cases.
type TriggerControlLabelOptionsExactly interface {
	TriggerControlLabelOptions
	isTriggerControlLabelOptions() bool
}

// _TriggerControlLabelOptions is the data-structure of this message
type _TriggerControlLabelOptions struct {
	LabelFlavour TriggerControlLabelFlavour
	LabelType    TriggerControlLabelType
	// Reserved Fields
	reservedField0 *bool
	reservedField1 *bool
	reservedField2 *bool
	reservedField3 *bool
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_TriggerControlLabelOptions) GetLabelFlavour() TriggerControlLabelFlavour {
	return m.LabelFlavour
}

func (m *_TriggerControlLabelOptions) GetLabelType() TriggerControlLabelType {
	return m.LabelType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewTriggerControlLabelOptions factory function for _TriggerControlLabelOptions
func NewTriggerControlLabelOptions(labelFlavour TriggerControlLabelFlavour, labelType TriggerControlLabelType) *_TriggerControlLabelOptions {
	return &_TriggerControlLabelOptions{LabelFlavour: labelFlavour, LabelType: labelType}
}

// Deprecated: use the interface for direct cast
func CastTriggerControlLabelOptions(structType any) TriggerControlLabelOptions {
	if casted, ok := structType.(TriggerControlLabelOptions); ok {
		return casted
	}
	if casted, ok := structType.(*TriggerControlLabelOptions); ok {
		return *casted
	}
	return nil
}

func (m *_TriggerControlLabelOptions) GetTypeName() string {
	return "TriggerControlLabelOptions"
}

func (m *_TriggerControlLabelOptions) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Reserved Field (reserved)
	lengthInBits += 1

	// Simple field (labelFlavour)
	lengthInBits += 2

	// Reserved Field (reserved)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 1

	// Simple field (labelType)
	lengthInBits += 2

	// Reserved Field (reserved)
	lengthInBits += 1

	return lengthInBits
}

func (m *_TriggerControlLabelOptions) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func TriggerControlLabelOptionsParse(ctx context.Context, theBytes []byte) (TriggerControlLabelOptions, error) {
	return TriggerControlLabelOptionsParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func TriggerControlLabelOptionsParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (TriggerControlLabelOptions, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (TriggerControlLabelOptions, error) {
		return TriggerControlLabelOptionsParseWithBuffer(ctx, readBuffer)
	}
}

func TriggerControlLabelOptionsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (TriggerControlLabelOptions, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TriggerControlLabelOptions"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TriggerControlLabelOptions")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadBoolean(readBuffer), bool(false))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}

	labelFlavour, err := ReadEnumField[TriggerControlLabelFlavour](ctx, "labelFlavour", "TriggerControlLabelFlavour", ReadEnum(TriggerControlLabelFlavourByValue, ReadUnsignedByte(readBuffer, uint8(2))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'labelFlavour' field"))
	}

	reservedField1, err := ReadReservedField(ctx, "reserved", ReadBoolean(readBuffer), bool(false))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}

	reservedField2, err := ReadReservedField(ctx, "reserved", ReadBoolean(readBuffer), bool(false))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}

	labelType, err := ReadEnumField[TriggerControlLabelType](ctx, "labelType", "TriggerControlLabelType", ReadEnum(TriggerControlLabelTypeByValue, ReadUnsignedByte(readBuffer, uint8(2))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'labelType' field"))
	}

	reservedField3, err := ReadReservedField(ctx, "reserved", ReadBoolean(readBuffer), bool(false))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}

	if closeErr := readBuffer.CloseContext("TriggerControlLabelOptions"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TriggerControlLabelOptions")
	}

	// Create the instance
	return &_TriggerControlLabelOptions{
		LabelFlavour:   labelFlavour,
		LabelType:      labelType,
		reservedField0: reservedField0,
		reservedField1: reservedField1,
		reservedField2: reservedField2,
		reservedField3: reservedField3,
	}, nil
}

func (m *_TriggerControlLabelOptions) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TriggerControlLabelOptions) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("TriggerControlLabelOptions"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for TriggerControlLabelOptions")
	}

	if err := WriteReservedField[bool](ctx, "reserved", bool(false), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'reserved' field number 1")
	}

	// Simple Field (labelFlavour)
	if pushErr := writeBuffer.PushContext("labelFlavour"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for labelFlavour")
	}
	_labelFlavourErr := writeBuffer.WriteSerializable(ctx, m.GetLabelFlavour())
	if popErr := writeBuffer.PopContext("labelFlavour"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for labelFlavour")
	}
	if _labelFlavourErr != nil {
		return errors.Wrap(_labelFlavourErr, "Error serializing 'labelFlavour' field")
	}

	if err := WriteReservedField[bool](ctx, "reserved", bool(false), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'reserved' field number 2")
	}

	if err := WriteReservedField[bool](ctx, "reserved", bool(false), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'reserved' field number 3")
	}

	// Simple Field (labelType)
	if pushErr := writeBuffer.PushContext("labelType"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for labelType")
	}
	_labelTypeErr := writeBuffer.WriteSerializable(ctx, m.GetLabelType())
	if popErr := writeBuffer.PopContext("labelType"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for labelType")
	}
	if _labelTypeErr != nil {
		return errors.Wrap(_labelTypeErr, "Error serializing 'labelType' field")
	}

	if err := WriteReservedField[bool](ctx, "reserved", bool(false), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'reserved' field number 4")
	}

	if popErr := writeBuffer.PopContext("TriggerControlLabelOptions"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for TriggerControlLabelOptions")
	}
	return nil
}

func (m *_TriggerControlLabelOptions) isTriggerControlLabelOptions() bool {
	return true
}

func (m *_TriggerControlLabelOptions) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
