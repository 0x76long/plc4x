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

// SALDataTelephonyStatusAndControl is the corresponding interface of SALDataTelephonyStatusAndControl
type SALDataTelephonyStatusAndControl interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SALData
	// GetTelephonyData returns TelephonyData (property field)
	GetTelephonyData() TelephonyData
}

// SALDataTelephonyStatusAndControlExactly can be used when we want exactly this type and not a type which fulfills SALDataTelephonyStatusAndControl.
// This is useful for switch cases.
type SALDataTelephonyStatusAndControlExactly interface {
	SALDataTelephonyStatusAndControl
	isSALDataTelephonyStatusAndControl() bool
}

// _SALDataTelephonyStatusAndControl is the data-structure of this message
type _SALDataTelephonyStatusAndControl struct {
	*_SALData
	TelephonyData TelephonyData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SALDataTelephonyStatusAndControl) GetApplicationId() ApplicationId {
	return ApplicationId_TELEPHONY_STATUS_AND_CONTROL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SALDataTelephonyStatusAndControl) InitializeParent(parent SALData, salData SALData) {
	m.SalData = salData
}

func (m *_SALDataTelephonyStatusAndControl) GetParent() SALData {
	return m._SALData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SALDataTelephonyStatusAndControl) GetTelephonyData() TelephonyData {
	return m.TelephonyData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSALDataTelephonyStatusAndControl factory function for _SALDataTelephonyStatusAndControl
func NewSALDataTelephonyStatusAndControl(telephonyData TelephonyData, salData SALData) *_SALDataTelephonyStatusAndControl {
	_result := &_SALDataTelephonyStatusAndControl{
		TelephonyData: telephonyData,
		_SALData:      NewSALData(salData),
	}
	_result._SALData._SALDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSALDataTelephonyStatusAndControl(structType any) SALDataTelephonyStatusAndControl {
	if casted, ok := structType.(SALDataTelephonyStatusAndControl); ok {
		return casted
	}
	if casted, ok := structType.(*SALDataTelephonyStatusAndControl); ok {
		return *casted
	}
	return nil
}

func (m *_SALDataTelephonyStatusAndControl) GetTypeName() string {
	return "SALDataTelephonyStatusAndControl"
}

func (m *_SALDataTelephonyStatusAndControl) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (telephonyData)
	lengthInBits += m.TelephonyData.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_SALDataTelephonyStatusAndControl) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SALDataTelephonyStatusAndControlParse(ctx context.Context, theBytes []byte, applicationId ApplicationId) (SALDataTelephonyStatusAndControl, error) {
	return SALDataTelephonyStatusAndControlParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), applicationId)
}

func SALDataTelephonyStatusAndControlParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, applicationId ApplicationId) (SALDataTelephonyStatusAndControl, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("SALDataTelephonyStatusAndControl"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SALDataTelephonyStatusAndControl")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (telephonyData)
	if pullErr := readBuffer.PullContext("telephonyData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for telephonyData")
	}
	_telephonyData, _telephonyDataErr := TelephonyDataParseWithBuffer(ctx, readBuffer)
	if _telephonyDataErr != nil {
		return nil, errors.Wrap(_telephonyDataErr, "Error parsing 'telephonyData' field of SALDataTelephonyStatusAndControl")
	}
	telephonyData := _telephonyData.(TelephonyData)
	if closeErr := readBuffer.CloseContext("telephonyData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for telephonyData")
	}

	if closeErr := readBuffer.CloseContext("SALDataTelephonyStatusAndControl"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SALDataTelephonyStatusAndControl")
	}

	// Create a partially initialized instance
	_child := &_SALDataTelephonyStatusAndControl{
		_SALData:      &_SALData{},
		TelephonyData: telephonyData,
	}
	_child._SALData._SALDataChildRequirements = _child
	return _child, nil
}

func (m *_SALDataTelephonyStatusAndControl) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SALDataTelephonyStatusAndControl) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SALDataTelephonyStatusAndControl"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SALDataTelephonyStatusAndControl")
		}

		// Simple Field (telephonyData)
		if pushErr := writeBuffer.PushContext("telephonyData"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for telephonyData")
		}
		_telephonyDataErr := writeBuffer.WriteSerializable(ctx, m.GetTelephonyData())
		if popErr := writeBuffer.PopContext("telephonyData"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for telephonyData")
		}
		if _telephonyDataErr != nil {
			return errors.Wrap(_telephonyDataErr, "Error serializing 'telephonyData' field")
		}

		if popErr := writeBuffer.PopContext("SALDataTelephonyStatusAndControl"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SALDataTelephonyStatusAndControl")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SALDataTelephonyStatusAndControl) isSALDataTelephonyStatusAndControl() bool {
	return true
}

func (m *_SALDataTelephonyStatusAndControl) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
