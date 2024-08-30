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

// NLMUpdateKeyUpdateControlFlags is the corresponding interface of NLMUpdateKeyUpdateControlFlags
type NLMUpdateKeyUpdateControlFlags interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetSet1KeyRevisionActivationTimeExpirationTimePresent returns Set1KeyRevisionActivationTimeExpirationTimePresent (property field)
	GetSet1KeyRevisionActivationTimeExpirationTimePresent() bool
	// GetSet1KeyCountKeyParametersPresent returns Set1KeyCountKeyParametersPresent (property field)
	GetSet1KeyCountKeyParametersPresent() bool
	// GetSet1ShouldBeCleared returns Set1ShouldBeCleared (property field)
	GetSet1ShouldBeCleared() bool
	// GetSet2KeyRevisionActivationTimeExpirationTimePresent returns Set2KeyRevisionActivationTimeExpirationTimePresent (property field)
	GetSet2KeyRevisionActivationTimeExpirationTimePresent() bool
	// GetSet2KeyCountKeyParametersPresent returns Set2KeyCountKeyParametersPresent (property field)
	GetSet2KeyCountKeyParametersPresent() bool
	// GetSet2ShouldBeCleared returns Set2ShouldBeCleared (property field)
	GetSet2ShouldBeCleared() bool
	// GetMoreMessagesToBeExpected returns MoreMessagesToBeExpected (property field)
	GetMoreMessagesToBeExpected() bool
	// GetRemoveAllKeys returns RemoveAllKeys (property field)
	GetRemoveAllKeys() bool
}

// NLMUpdateKeyUpdateControlFlagsExactly can be used when we want exactly this type and not a type which fulfills NLMUpdateKeyUpdateControlFlags.
// This is useful for switch cases.
type NLMUpdateKeyUpdateControlFlagsExactly interface {
	NLMUpdateKeyUpdateControlFlags
	isNLMUpdateKeyUpdateControlFlags() bool
}

// _NLMUpdateKeyUpdateControlFlags is the data-structure of this message
type _NLMUpdateKeyUpdateControlFlags struct {
	Set1KeyRevisionActivationTimeExpirationTimePresent bool
	Set1KeyCountKeyParametersPresent                   bool
	Set1ShouldBeCleared                                bool
	Set2KeyRevisionActivationTimeExpirationTimePresent bool
	Set2KeyCountKeyParametersPresent                   bool
	Set2ShouldBeCleared                                bool
	MoreMessagesToBeExpected                           bool
	RemoveAllKeys                                      bool
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NLMUpdateKeyUpdateControlFlags) GetSet1KeyRevisionActivationTimeExpirationTimePresent() bool {
	return m.Set1KeyRevisionActivationTimeExpirationTimePresent
}

func (m *_NLMUpdateKeyUpdateControlFlags) GetSet1KeyCountKeyParametersPresent() bool {
	return m.Set1KeyCountKeyParametersPresent
}

func (m *_NLMUpdateKeyUpdateControlFlags) GetSet1ShouldBeCleared() bool {
	return m.Set1ShouldBeCleared
}

func (m *_NLMUpdateKeyUpdateControlFlags) GetSet2KeyRevisionActivationTimeExpirationTimePresent() bool {
	return m.Set2KeyRevisionActivationTimeExpirationTimePresent
}

func (m *_NLMUpdateKeyUpdateControlFlags) GetSet2KeyCountKeyParametersPresent() bool {
	return m.Set2KeyCountKeyParametersPresent
}

func (m *_NLMUpdateKeyUpdateControlFlags) GetSet2ShouldBeCleared() bool {
	return m.Set2ShouldBeCleared
}

func (m *_NLMUpdateKeyUpdateControlFlags) GetMoreMessagesToBeExpected() bool {
	return m.MoreMessagesToBeExpected
}

func (m *_NLMUpdateKeyUpdateControlFlags) GetRemoveAllKeys() bool {
	return m.RemoveAllKeys
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNLMUpdateKeyUpdateControlFlags factory function for _NLMUpdateKeyUpdateControlFlags
func NewNLMUpdateKeyUpdateControlFlags(set1KeyRevisionActivationTimeExpirationTimePresent bool, set1KeyCountKeyParametersPresent bool, set1ShouldBeCleared bool, set2KeyRevisionActivationTimeExpirationTimePresent bool, set2KeyCountKeyParametersPresent bool, set2ShouldBeCleared bool, moreMessagesToBeExpected bool, removeAllKeys bool) *_NLMUpdateKeyUpdateControlFlags {
	return &_NLMUpdateKeyUpdateControlFlags{Set1KeyRevisionActivationTimeExpirationTimePresent: set1KeyRevisionActivationTimeExpirationTimePresent, Set1KeyCountKeyParametersPresent: set1KeyCountKeyParametersPresent, Set1ShouldBeCleared: set1ShouldBeCleared, Set2KeyRevisionActivationTimeExpirationTimePresent: set2KeyRevisionActivationTimeExpirationTimePresent, Set2KeyCountKeyParametersPresent: set2KeyCountKeyParametersPresent, Set2ShouldBeCleared: set2ShouldBeCleared, MoreMessagesToBeExpected: moreMessagesToBeExpected, RemoveAllKeys: removeAllKeys}
}

// Deprecated: use the interface for direct cast
func CastNLMUpdateKeyUpdateControlFlags(structType any) NLMUpdateKeyUpdateControlFlags {
	if casted, ok := structType.(NLMUpdateKeyUpdateControlFlags); ok {
		return casted
	}
	if casted, ok := structType.(*NLMUpdateKeyUpdateControlFlags); ok {
		return *casted
	}
	return nil
}

func (m *_NLMUpdateKeyUpdateControlFlags) GetTypeName() string {
	return "NLMUpdateKeyUpdateControlFlags"
}

func (m *_NLMUpdateKeyUpdateControlFlags) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (set1KeyRevisionActivationTimeExpirationTimePresent)
	lengthInBits += 1

	// Simple field (set1KeyCountKeyParametersPresent)
	lengthInBits += 1

	// Simple field (set1ShouldBeCleared)
	lengthInBits += 1

	// Simple field (set2KeyRevisionActivationTimeExpirationTimePresent)
	lengthInBits += 1

	// Simple field (set2KeyCountKeyParametersPresent)
	lengthInBits += 1

	// Simple field (set2ShouldBeCleared)
	lengthInBits += 1

	// Simple field (moreMessagesToBeExpected)
	lengthInBits += 1

	// Simple field (removeAllKeys)
	lengthInBits += 1

	return lengthInBits
}

func (m *_NLMUpdateKeyUpdateControlFlags) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func NLMUpdateKeyUpdateControlFlagsParse(ctx context.Context, theBytes []byte) (NLMUpdateKeyUpdateControlFlags, error) {
	return NLMUpdateKeyUpdateControlFlagsParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func NLMUpdateKeyUpdateControlFlagsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (NLMUpdateKeyUpdateControlFlags, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("NLMUpdateKeyUpdateControlFlags"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NLMUpdateKeyUpdateControlFlags")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (set1KeyRevisionActivationTimeExpirationTimePresent)
	_set1KeyRevisionActivationTimeExpirationTimePresent, _set1KeyRevisionActivationTimeExpirationTimePresentErr := /*TODO: migrate me*/ readBuffer.ReadBit("set1KeyRevisionActivationTimeExpirationTimePresent")
	if _set1KeyRevisionActivationTimeExpirationTimePresentErr != nil {
		return nil, errors.Wrap(_set1KeyRevisionActivationTimeExpirationTimePresentErr, "Error parsing 'set1KeyRevisionActivationTimeExpirationTimePresent' field of NLMUpdateKeyUpdateControlFlags")
	}
	set1KeyRevisionActivationTimeExpirationTimePresent := _set1KeyRevisionActivationTimeExpirationTimePresent

	// Simple Field (set1KeyCountKeyParametersPresent)
	_set1KeyCountKeyParametersPresent, _set1KeyCountKeyParametersPresentErr := /*TODO: migrate me*/ readBuffer.ReadBit("set1KeyCountKeyParametersPresent")
	if _set1KeyCountKeyParametersPresentErr != nil {
		return nil, errors.Wrap(_set1KeyCountKeyParametersPresentErr, "Error parsing 'set1KeyCountKeyParametersPresent' field of NLMUpdateKeyUpdateControlFlags")
	}
	set1KeyCountKeyParametersPresent := _set1KeyCountKeyParametersPresent

	// Simple Field (set1ShouldBeCleared)
	_set1ShouldBeCleared, _set1ShouldBeClearedErr := /*TODO: migrate me*/ readBuffer.ReadBit("set1ShouldBeCleared")
	if _set1ShouldBeClearedErr != nil {
		return nil, errors.Wrap(_set1ShouldBeClearedErr, "Error parsing 'set1ShouldBeCleared' field of NLMUpdateKeyUpdateControlFlags")
	}
	set1ShouldBeCleared := _set1ShouldBeCleared

	// Simple Field (set2KeyRevisionActivationTimeExpirationTimePresent)
	_set2KeyRevisionActivationTimeExpirationTimePresent, _set2KeyRevisionActivationTimeExpirationTimePresentErr := /*TODO: migrate me*/ readBuffer.ReadBit("set2KeyRevisionActivationTimeExpirationTimePresent")
	if _set2KeyRevisionActivationTimeExpirationTimePresentErr != nil {
		return nil, errors.Wrap(_set2KeyRevisionActivationTimeExpirationTimePresentErr, "Error parsing 'set2KeyRevisionActivationTimeExpirationTimePresent' field of NLMUpdateKeyUpdateControlFlags")
	}
	set2KeyRevisionActivationTimeExpirationTimePresent := _set2KeyRevisionActivationTimeExpirationTimePresent

	// Simple Field (set2KeyCountKeyParametersPresent)
	_set2KeyCountKeyParametersPresent, _set2KeyCountKeyParametersPresentErr := /*TODO: migrate me*/ readBuffer.ReadBit("set2KeyCountKeyParametersPresent")
	if _set2KeyCountKeyParametersPresentErr != nil {
		return nil, errors.Wrap(_set2KeyCountKeyParametersPresentErr, "Error parsing 'set2KeyCountKeyParametersPresent' field of NLMUpdateKeyUpdateControlFlags")
	}
	set2KeyCountKeyParametersPresent := _set2KeyCountKeyParametersPresent

	// Simple Field (set2ShouldBeCleared)
	_set2ShouldBeCleared, _set2ShouldBeClearedErr := /*TODO: migrate me*/ readBuffer.ReadBit("set2ShouldBeCleared")
	if _set2ShouldBeClearedErr != nil {
		return nil, errors.Wrap(_set2ShouldBeClearedErr, "Error parsing 'set2ShouldBeCleared' field of NLMUpdateKeyUpdateControlFlags")
	}
	set2ShouldBeCleared := _set2ShouldBeCleared

	// Simple Field (moreMessagesToBeExpected)
	_moreMessagesToBeExpected, _moreMessagesToBeExpectedErr := /*TODO: migrate me*/ readBuffer.ReadBit("moreMessagesToBeExpected")
	if _moreMessagesToBeExpectedErr != nil {
		return nil, errors.Wrap(_moreMessagesToBeExpectedErr, "Error parsing 'moreMessagesToBeExpected' field of NLMUpdateKeyUpdateControlFlags")
	}
	moreMessagesToBeExpected := _moreMessagesToBeExpected

	// Simple Field (removeAllKeys)
	_removeAllKeys, _removeAllKeysErr := /*TODO: migrate me*/ readBuffer.ReadBit("removeAllKeys")
	if _removeAllKeysErr != nil {
		return nil, errors.Wrap(_removeAllKeysErr, "Error parsing 'removeAllKeys' field of NLMUpdateKeyUpdateControlFlags")
	}
	removeAllKeys := _removeAllKeys

	if closeErr := readBuffer.CloseContext("NLMUpdateKeyUpdateControlFlags"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NLMUpdateKeyUpdateControlFlags")
	}

	// Create the instance
	return &_NLMUpdateKeyUpdateControlFlags{
		Set1KeyRevisionActivationTimeExpirationTimePresent: set1KeyRevisionActivationTimeExpirationTimePresent,
		Set1KeyCountKeyParametersPresent:                   set1KeyCountKeyParametersPresent,
		Set1ShouldBeCleared:                                set1ShouldBeCleared,
		Set2KeyRevisionActivationTimeExpirationTimePresent: set2KeyRevisionActivationTimeExpirationTimePresent,
		Set2KeyCountKeyParametersPresent:                   set2KeyCountKeyParametersPresent,
		Set2ShouldBeCleared:                                set2ShouldBeCleared,
		MoreMessagesToBeExpected:                           moreMessagesToBeExpected,
		RemoveAllKeys:                                      removeAllKeys,
	}, nil
}

func (m *_NLMUpdateKeyUpdateControlFlags) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NLMUpdateKeyUpdateControlFlags) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("NLMUpdateKeyUpdateControlFlags"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for NLMUpdateKeyUpdateControlFlags")
	}

	// Simple Field (set1KeyRevisionActivationTimeExpirationTimePresent)
	set1KeyRevisionActivationTimeExpirationTimePresent := bool(m.GetSet1KeyRevisionActivationTimeExpirationTimePresent())
	_set1KeyRevisionActivationTimeExpirationTimePresentErr := /*TODO: migrate me*/ writeBuffer.WriteBit("set1KeyRevisionActivationTimeExpirationTimePresent", (set1KeyRevisionActivationTimeExpirationTimePresent))
	if _set1KeyRevisionActivationTimeExpirationTimePresentErr != nil {
		return errors.Wrap(_set1KeyRevisionActivationTimeExpirationTimePresentErr, "Error serializing 'set1KeyRevisionActivationTimeExpirationTimePresent' field")
	}

	// Simple Field (set1KeyCountKeyParametersPresent)
	set1KeyCountKeyParametersPresent := bool(m.GetSet1KeyCountKeyParametersPresent())
	_set1KeyCountKeyParametersPresentErr := /*TODO: migrate me*/ writeBuffer.WriteBit("set1KeyCountKeyParametersPresent", (set1KeyCountKeyParametersPresent))
	if _set1KeyCountKeyParametersPresentErr != nil {
		return errors.Wrap(_set1KeyCountKeyParametersPresentErr, "Error serializing 'set1KeyCountKeyParametersPresent' field")
	}

	// Simple Field (set1ShouldBeCleared)
	set1ShouldBeCleared := bool(m.GetSet1ShouldBeCleared())
	_set1ShouldBeClearedErr := /*TODO: migrate me*/ writeBuffer.WriteBit("set1ShouldBeCleared", (set1ShouldBeCleared))
	if _set1ShouldBeClearedErr != nil {
		return errors.Wrap(_set1ShouldBeClearedErr, "Error serializing 'set1ShouldBeCleared' field")
	}

	// Simple Field (set2KeyRevisionActivationTimeExpirationTimePresent)
	set2KeyRevisionActivationTimeExpirationTimePresent := bool(m.GetSet2KeyRevisionActivationTimeExpirationTimePresent())
	_set2KeyRevisionActivationTimeExpirationTimePresentErr := /*TODO: migrate me*/ writeBuffer.WriteBit("set2KeyRevisionActivationTimeExpirationTimePresent", (set2KeyRevisionActivationTimeExpirationTimePresent))
	if _set2KeyRevisionActivationTimeExpirationTimePresentErr != nil {
		return errors.Wrap(_set2KeyRevisionActivationTimeExpirationTimePresentErr, "Error serializing 'set2KeyRevisionActivationTimeExpirationTimePresent' field")
	}

	// Simple Field (set2KeyCountKeyParametersPresent)
	set2KeyCountKeyParametersPresent := bool(m.GetSet2KeyCountKeyParametersPresent())
	_set2KeyCountKeyParametersPresentErr := /*TODO: migrate me*/ writeBuffer.WriteBit("set2KeyCountKeyParametersPresent", (set2KeyCountKeyParametersPresent))
	if _set2KeyCountKeyParametersPresentErr != nil {
		return errors.Wrap(_set2KeyCountKeyParametersPresentErr, "Error serializing 'set2KeyCountKeyParametersPresent' field")
	}

	// Simple Field (set2ShouldBeCleared)
	set2ShouldBeCleared := bool(m.GetSet2ShouldBeCleared())
	_set2ShouldBeClearedErr := /*TODO: migrate me*/ writeBuffer.WriteBit("set2ShouldBeCleared", (set2ShouldBeCleared))
	if _set2ShouldBeClearedErr != nil {
		return errors.Wrap(_set2ShouldBeClearedErr, "Error serializing 'set2ShouldBeCleared' field")
	}

	// Simple Field (moreMessagesToBeExpected)
	moreMessagesToBeExpected := bool(m.GetMoreMessagesToBeExpected())
	_moreMessagesToBeExpectedErr := /*TODO: migrate me*/ writeBuffer.WriteBit("moreMessagesToBeExpected", (moreMessagesToBeExpected))
	if _moreMessagesToBeExpectedErr != nil {
		return errors.Wrap(_moreMessagesToBeExpectedErr, "Error serializing 'moreMessagesToBeExpected' field")
	}

	// Simple Field (removeAllKeys)
	removeAllKeys := bool(m.GetRemoveAllKeys())
	_removeAllKeysErr := /*TODO: migrate me*/ writeBuffer.WriteBit("removeAllKeys", (removeAllKeys))
	if _removeAllKeysErr != nil {
		return errors.Wrap(_removeAllKeysErr, "Error serializing 'removeAllKeys' field")
	}

	if popErr := writeBuffer.PopContext("NLMUpdateKeyUpdateControlFlags"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for NLMUpdateKeyUpdateControlFlags")
	}
	return nil
}

func (m *_NLMUpdateKeyUpdateControlFlags) isNLMUpdateKeyUpdateControlFlags() bool {
	return true
}

func (m *_NLMUpdateKeyUpdateControlFlags) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
