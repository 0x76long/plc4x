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

// MonitoringFilterResult is the corresponding interface of MonitoringFilterResult
type MonitoringFilterResult interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
}

// MonitoringFilterResultExactly can be used when we want exactly this type and not a type which fulfills MonitoringFilterResult.
// This is useful for switch cases.
type MonitoringFilterResultExactly interface {
	MonitoringFilterResult
	isMonitoringFilterResult() bool
}

// _MonitoringFilterResult is the data-structure of this message
type _MonitoringFilterResult struct {
	*_ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_MonitoringFilterResult) GetIdentifier() string {
	return "733"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MonitoringFilterResult) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_MonitoringFilterResult) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

// NewMonitoringFilterResult factory function for _MonitoringFilterResult
func NewMonitoringFilterResult() *_MonitoringFilterResult {
	_result := &_MonitoringFilterResult{
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastMonitoringFilterResult(structType any) MonitoringFilterResult {
	if casted, ok := structType.(MonitoringFilterResult); ok {
		return casted
	}
	if casted, ok := structType.(*MonitoringFilterResult); ok {
		return *casted
	}
	return nil
}

func (m *_MonitoringFilterResult) GetTypeName() string {
	return "MonitoringFilterResult"
}

func (m *_MonitoringFilterResult) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_MonitoringFilterResult) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func MonitoringFilterResultParse(ctx context.Context, theBytes []byte, identifier string) (MonitoringFilterResult, error) {
	return MonitoringFilterResultParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func MonitoringFilterResultParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (MonitoringFilterResult, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("MonitoringFilterResult"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MonitoringFilterResult")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("MonitoringFilterResult"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MonitoringFilterResult")
	}

	// Create a partially initialized instance
	_child := &_MonitoringFilterResult{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_MonitoringFilterResult) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MonitoringFilterResult) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MonitoringFilterResult"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MonitoringFilterResult")
		}

		if popErr := writeBuffer.PopContext("MonitoringFilterResult"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MonitoringFilterResult")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MonitoringFilterResult) isMonitoringFilterResult() bool {
	return true
}

func (m *_MonitoringFilterResult) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
