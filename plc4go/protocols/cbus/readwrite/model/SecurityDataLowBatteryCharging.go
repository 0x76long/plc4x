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

// SecurityDataLowBatteryCharging is the corresponding interface of SecurityDataLowBatteryCharging
type SecurityDataLowBatteryCharging interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SecurityData
	// GetStartStop returns StartStop (property field)
	GetStartStop() byte
	// GetChargeStopped returns ChargeStopped (virtual field)
	GetChargeStopped() bool
	// GetChargeStarted returns ChargeStarted (virtual field)
	GetChargeStarted() bool
}

// SecurityDataLowBatteryChargingExactly can be used when we want exactly this type and not a type which fulfills SecurityDataLowBatteryCharging.
// This is useful for switch cases.
type SecurityDataLowBatteryChargingExactly interface {
	SecurityDataLowBatteryCharging
	isSecurityDataLowBatteryCharging() bool
}

// _SecurityDataLowBatteryCharging is the data-structure of this message
type _SecurityDataLowBatteryCharging struct {
	*_SecurityData
	StartStop byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SecurityDataLowBatteryCharging) InitializeParent(parent SecurityData, commandTypeContainer SecurityCommandTypeContainer, argument byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.Argument = argument
}

func (m *_SecurityDataLowBatteryCharging) GetParent() SecurityData {
	return m._SecurityData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SecurityDataLowBatteryCharging) GetStartStop() byte {
	return m.StartStop
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_SecurityDataLowBatteryCharging) GetChargeStopped() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetStartStop()) == (0x00)))
}

func (m *_SecurityDataLowBatteryCharging) GetChargeStarted() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetStartStop()) > (0xFE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSecurityDataLowBatteryCharging factory function for _SecurityDataLowBatteryCharging
func NewSecurityDataLowBatteryCharging(startStop byte, commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataLowBatteryCharging {
	_result := &_SecurityDataLowBatteryCharging{
		StartStop:     startStop,
		_SecurityData: NewSecurityData(commandTypeContainer, argument),
	}
	_result._SecurityData._SecurityDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSecurityDataLowBatteryCharging(structType any) SecurityDataLowBatteryCharging {
	if casted, ok := structType.(SecurityDataLowBatteryCharging); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataLowBatteryCharging); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataLowBatteryCharging) GetTypeName() string {
	return "SecurityDataLowBatteryCharging"
}

func (m *_SecurityDataLowBatteryCharging) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (startStop)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_SecurityDataLowBatteryCharging) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SecurityDataLowBatteryChargingParse(ctx context.Context, theBytes []byte) (SecurityDataLowBatteryCharging, error) {
	return SecurityDataLowBatteryChargingParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func SecurityDataLowBatteryChargingParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (SecurityDataLowBatteryCharging, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("SecurityDataLowBatteryCharging"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataLowBatteryCharging")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (startStop)
	_startStop, _startStopErr := /*TODO: migrate me*/ readBuffer.ReadByte("startStop")
	if _startStopErr != nil {
		return nil, errors.Wrap(_startStopErr, "Error parsing 'startStop' field of SecurityDataLowBatteryCharging")
	}
	startStop := _startStop

	// Virtual field
	_chargeStopped := bool((startStop) == (0x00))
	chargeStopped := bool(_chargeStopped)
	_ = chargeStopped

	// Virtual field
	_chargeStarted := bool((startStop) > (0xFE))
	chargeStarted := bool(_chargeStarted)
	_ = chargeStarted

	if closeErr := readBuffer.CloseContext("SecurityDataLowBatteryCharging"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataLowBatteryCharging")
	}

	// Create a partially initialized instance
	_child := &_SecurityDataLowBatteryCharging{
		_SecurityData: &_SecurityData{},
		StartStop:     startStop,
	}
	_child._SecurityData._SecurityDataChildRequirements = _child
	return _child, nil
}

func (m *_SecurityDataLowBatteryCharging) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataLowBatteryCharging) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataLowBatteryCharging"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataLowBatteryCharging")
		}

		// Simple Field (startStop)
		startStop := byte(m.GetStartStop())
		_startStopErr := /*TODO: migrate me*/ writeBuffer.WriteByte("startStop", (startStop))
		if _startStopErr != nil {
			return errors.Wrap(_startStopErr, "Error serializing 'startStop' field")
		}
		// Virtual field
		chargeStopped := m.GetChargeStopped()
		_ = chargeStopped
		if _chargeStoppedErr := writeBuffer.WriteVirtual(ctx, "chargeStopped", m.GetChargeStopped()); _chargeStoppedErr != nil {
			return errors.Wrap(_chargeStoppedErr, "Error serializing 'chargeStopped' field")
		}
		// Virtual field
		chargeStarted := m.GetChargeStarted()
		_ = chargeStarted
		if _chargeStartedErr := writeBuffer.WriteVirtual(ctx, "chargeStarted", m.GetChargeStarted()); _chargeStartedErr != nil {
			return errors.Wrap(_chargeStartedErr, "Error serializing 'chargeStarted' field")
		}

		if popErr := writeBuffer.PopContext("SecurityDataLowBatteryCharging"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataLowBatteryCharging")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityDataLowBatteryCharging) isSecurityDataLowBatteryCharging() bool {
	return true
}

func (m *_SecurityDataLowBatteryCharging) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
