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

// BACnetAccessRule is the corresponding interface of BACnetAccessRule
type BACnetAccessRule interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetTimeRangeSpecifier returns TimeRangeSpecifier (property field)
	GetTimeRangeSpecifier() BACnetAccessRuleTimeRangeSpecifierTagged
	// GetTimeRange returns TimeRange (property field)
	GetTimeRange() BACnetDeviceObjectPropertyReferenceEnclosed
	// GetLocationSpecifier returns LocationSpecifier (property field)
	GetLocationSpecifier() BACnetAccessRuleLocationSpecifierTagged
	// GetLocation returns Location (property field)
	GetLocation() BACnetDeviceObjectReferenceEnclosed
	// GetEnable returns Enable (property field)
	GetEnable() BACnetContextTagBoolean
}

// BACnetAccessRuleExactly can be used when we want exactly this type and not a type which fulfills BACnetAccessRule.
// This is useful for switch cases.
type BACnetAccessRuleExactly interface {
	BACnetAccessRule
	isBACnetAccessRule() bool
}

// _BACnetAccessRule is the data-structure of this message
type _BACnetAccessRule struct {
	TimeRangeSpecifier BACnetAccessRuleTimeRangeSpecifierTagged
	TimeRange          BACnetDeviceObjectPropertyReferenceEnclosed
	LocationSpecifier  BACnetAccessRuleLocationSpecifierTagged
	Location           BACnetDeviceObjectReferenceEnclosed
	Enable             BACnetContextTagBoolean
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetAccessRule) GetTimeRangeSpecifier() BACnetAccessRuleTimeRangeSpecifierTagged {
	return m.TimeRangeSpecifier
}

func (m *_BACnetAccessRule) GetTimeRange() BACnetDeviceObjectPropertyReferenceEnclosed {
	return m.TimeRange
}

func (m *_BACnetAccessRule) GetLocationSpecifier() BACnetAccessRuleLocationSpecifierTagged {
	return m.LocationSpecifier
}

func (m *_BACnetAccessRule) GetLocation() BACnetDeviceObjectReferenceEnclosed {
	return m.Location
}

func (m *_BACnetAccessRule) GetEnable() BACnetContextTagBoolean {
	return m.Enable
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetAccessRule factory function for _BACnetAccessRule
func NewBACnetAccessRule(timeRangeSpecifier BACnetAccessRuleTimeRangeSpecifierTagged, timeRange BACnetDeviceObjectPropertyReferenceEnclosed, locationSpecifier BACnetAccessRuleLocationSpecifierTagged, location BACnetDeviceObjectReferenceEnclosed, enable BACnetContextTagBoolean) *_BACnetAccessRule {
	return &_BACnetAccessRule{TimeRangeSpecifier: timeRangeSpecifier, TimeRange: timeRange, LocationSpecifier: locationSpecifier, Location: location, Enable: enable}
}

// Deprecated: use the interface for direct cast
func CastBACnetAccessRule(structType any) BACnetAccessRule {
	if casted, ok := structType.(BACnetAccessRule); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetAccessRule); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetAccessRule) GetTypeName() string {
	return "BACnetAccessRule"
}

func (m *_BACnetAccessRule) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (timeRangeSpecifier)
	lengthInBits += m.TimeRangeSpecifier.GetLengthInBits(ctx)

	// Optional Field (timeRange)
	if m.TimeRange != nil {
		lengthInBits += m.TimeRange.GetLengthInBits(ctx)
	}

	// Simple field (locationSpecifier)
	lengthInBits += m.LocationSpecifier.GetLengthInBits(ctx)

	// Optional Field (location)
	if m.Location != nil {
		lengthInBits += m.Location.GetLengthInBits(ctx)
	}

	// Simple field (enable)
	lengthInBits += m.Enable.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetAccessRule) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetAccessRuleParse(ctx context.Context, theBytes []byte) (BACnetAccessRule, error) {
	return BACnetAccessRuleParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetAccessRuleParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAccessRule, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetAccessRule"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetAccessRule")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (timeRangeSpecifier)
	if pullErr := readBuffer.PullContext("timeRangeSpecifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for timeRangeSpecifier")
	}
	_timeRangeSpecifier, _timeRangeSpecifierErr := BACnetAccessRuleTimeRangeSpecifierTaggedParseWithBuffer(ctx, readBuffer, uint8(uint8(0)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _timeRangeSpecifierErr != nil {
		return nil, errors.Wrap(_timeRangeSpecifierErr, "Error parsing 'timeRangeSpecifier' field of BACnetAccessRule")
	}
	timeRangeSpecifier := _timeRangeSpecifier.(BACnetAccessRuleTimeRangeSpecifierTagged)
	if closeErr := readBuffer.CloseContext("timeRangeSpecifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for timeRangeSpecifier")
	}

	// Optional Field (timeRange) (Can be skipped, if a given expression evaluates to false)
	var timeRange BACnetDeviceObjectPropertyReferenceEnclosed = nil
	if bool((timeRangeSpecifier) != (nil)) {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("timeRange"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for timeRange")
		}
		_val, _err := BACnetDeviceObjectPropertyReferenceEnclosedParseWithBuffer(ctx, readBuffer, uint8(1))
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'timeRange' field of BACnetAccessRule")
		default:
			timeRange = _val.(BACnetDeviceObjectPropertyReferenceEnclosed)
			if closeErr := readBuffer.CloseContext("timeRange"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for timeRange")
			}
		}
	}

	// Simple Field (locationSpecifier)
	if pullErr := readBuffer.PullContext("locationSpecifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for locationSpecifier")
	}
	_locationSpecifier, _locationSpecifierErr := BACnetAccessRuleLocationSpecifierTaggedParseWithBuffer(ctx, readBuffer, uint8(uint8(2)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _locationSpecifierErr != nil {
		return nil, errors.Wrap(_locationSpecifierErr, "Error parsing 'locationSpecifier' field of BACnetAccessRule")
	}
	locationSpecifier := _locationSpecifier.(BACnetAccessRuleLocationSpecifierTagged)
	if closeErr := readBuffer.CloseContext("locationSpecifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for locationSpecifier")
	}

	// Optional Field (location) (Can be skipped, if a given expression evaluates to false)
	var location BACnetDeviceObjectReferenceEnclosed = nil
	if bool((locationSpecifier) != (nil)) {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("location"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for location")
		}
		_val, _err := BACnetDeviceObjectReferenceEnclosedParseWithBuffer(ctx, readBuffer, uint8(3))
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'location' field of BACnetAccessRule")
		default:
			location = _val.(BACnetDeviceObjectReferenceEnclosed)
			if closeErr := readBuffer.CloseContext("location"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for location")
			}
		}
	}

	// Simple Field (enable)
	if pullErr := readBuffer.PullContext("enable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for enable")
	}
	_enable, _enableErr := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(uint8(4)), BACnetDataType(BACnetDataType_BOOLEAN))
	if _enableErr != nil {
		return nil, errors.Wrap(_enableErr, "Error parsing 'enable' field of BACnetAccessRule")
	}
	enable := _enable.(BACnetContextTagBoolean)
	if closeErr := readBuffer.CloseContext("enable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for enable")
	}

	if closeErr := readBuffer.CloseContext("BACnetAccessRule"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetAccessRule")
	}

	// Create the instance
	return &_BACnetAccessRule{
		TimeRangeSpecifier: timeRangeSpecifier,
		TimeRange:          timeRange,
		LocationSpecifier:  locationSpecifier,
		Location:           location,
		Enable:             enable,
	}, nil
}

func (m *_BACnetAccessRule) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetAccessRule) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetAccessRule"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetAccessRule")
	}

	// Simple Field (timeRangeSpecifier)
	if pushErr := writeBuffer.PushContext("timeRangeSpecifier"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for timeRangeSpecifier")
	}
	_timeRangeSpecifierErr := writeBuffer.WriteSerializable(ctx, m.GetTimeRangeSpecifier())
	if popErr := writeBuffer.PopContext("timeRangeSpecifier"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for timeRangeSpecifier")
	}
	if _timeRangeSpecifierErr != nil {
		return errors.Wrap(_timeRangeSpecifierErr, "Error serializing 'timeRangeSpecifier' field")
	}

	// Optional Field (timeRange) (Can be skipped, if the value is null)
	var timeRange BACnetDeviceObjectPropertyReferenceEnclosed = nil
	if m.GetTimeRange() != nil {
		if pushErr := writeBuffer.PushContext("timeRange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for timeRange")
		}
		timeRange = m.GetTimeRange()
		_timeRangeErr := writeBuffer.WriteSerializable(ctx, timeRange)
		if popErr := writeBuffer.PopContext("timeRange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for timeRange")
		}
		if _timeRangeErr != nil {
			return errors.Wrap(_timeRangeErr, "Error serializing 'timeRange' field")
		}
	}

	// Simple Field (locationSpecifier)
	if pushErr := writeBuffer.PushContext("locationSpecifier"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for locationSpecifier")
	}
	_locationSpecifierErr := writeBuffer.WriteSerializable(ctx, m.GetLocationSpecifier())
	if popErr := writeBuffer.PopContext("locationSpecifier"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for locationSpecifier")
	}
	if _locationSpecifierErr != nil {
		return errors.Wrap(_locationSpecifierErr, "Error serializing 'locationSpecifier' field")
	}

	// Optional Field (location) (Can be skipped, if the value is null)
	var location BACnetDeviceObjectReferenceEnclosed = nil
	if m.GetLocation() != nil {
		if pushErr := writeBuffer.PushContext("location"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for location")
		}
		location = m.GetLocation()
		_locationErr := writeBuffer.WriteSerializable(ctx, location)
		if popErr := writeBuffer.PopContext("location"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for location")
		}
		if _locationErr != nil {
			return errors.Wrap(_locationErr, "Error serializing 'location' field")
		}
	}

	// Simple Field (enable)
	if pushErr := writeBuffer.PushContext("enable"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for enable")
	}
	_enableErr := writeBuffer.WriteSerializable(ctx, m.GetEnable())
	if popErr := writeBuffer.PopContext("enable"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for enable")
	}
	if _enableErr != nil {
		return errors.Wrap(_enableErr, "Error serializing 'enable' field")
	}

	if popErr := writeBuffer.PopContext("BACnetAccessRule"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetAccessRule")
	}
	return nil
}

func (m *_BACnetAccessRule) isBACnetAccessRule() bool {
	return true
}

func (m *_BACnetAccessRule) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
