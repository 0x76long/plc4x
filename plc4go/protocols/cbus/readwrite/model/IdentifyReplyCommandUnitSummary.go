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

// IdentifyReplyCommandUnitSummary is the corresponding interface of IdentifyReplyCommandUnitSummary
type IdentifyReplyCommandUnitSummary interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetAssertingNetworkBurden returns AssertingNetworkBurden (property field)
	GetAssertingNetworkBurden() bool
	// GetRestrikeTimingActive returns RestrikeTimingActive (property field)
	GetRestrikeTimingActive() bool
	// GetRemoteOFFInputAsserted returns RemoteOFFInputAsserted (property field)
	GetRemoteOFFInputAsserted() bool
	// GetRemoteONInputAsserted returns RemoteONInputAsserted (property field)
	GetRemoteONInputAsserted() bool
	// GetLocalToggleEnabled returns LocalToggleEnabled (property field)
	GetLocalToggleEnabled() bool
	// GetLocalToggleActiveState returns LocalToggleActiveState (property field)
	GetLocalToggleActiveState() bool
	// GetClockGenerationEnabled returns ClockGenerationEnabled (property field)
	GetClockGenerationEnabled() bool
	// GetUnitGeneratingClock returns UnitGeneratingClock (property field)
	GetUnitGeneratingClock() bool
	// IsIdentifyReplyCommandUnitSummary is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsIdentifyReplyCommandUnitSummary()
	// CreateBuilder creates a IdentifyReplyCommandUnitSummaryBuilder
	CreateIdentifyReplyCommandUnitSummaryBuilder() IdentifyReplyCommandUnitSummaryBuilder
}

// _IdentifyReplyCommandUnitSummary is the data-structure of this message
type _IdentifyReplyCommandUnitSummary struct {
	AssertingNetworkBurden bool
	RestrikeTimingActive   bool
	RemoteOFFInputAsserted bool
	RemoteONInputAsserted  bool
	LocalToggleEnabled     bool
	LocalToggleActiveState bool
	ClockGenerationEnabled bool
	UnitGeneratingClock    bool
}

var _ IdentifyReplyCommandUnitSummary = (*_IdentifyReplyCommandUnitSummary)(nil)

// NewIdentifyReplyCommandUnitSummary factory function for _IdentifyReplyCommandUnitSummary
func NewIdentifyReplyCommandUnitSummary(assertingNetworkBurden bool, restrikeTimingActive bool, remoteOFFInputAsserted bool, remoteONInputAsserted bool, localToggleEnabled bool, localToggleActiveState bool, clockGenerationEnabled bool, unitGeneratingClock bool) *_IdentifyReplyCommandUnitSummary {
	return &_IdentifyReplyCommandUnitSummary{AssertingNetworkBurden: assertingNetworkBurden, RestrikeTimingActive: restrikeTimingActive, RemoteOFFInputAsserted: remoteOFFInputAsserted, RemoteONInputAsserted: remoteONInputAsserted, LocalToggleEnabled: localToggleEnabled, LocalToggleActiveState: localToggleActiveState, ClockGenerationEnabled: clockGenerationEnabled, UnitGeneratingClock: unitGeneratingClock}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// IdentifyReplyCommandUnitSummaryBuilder is a builder for IdentifyReplyCommandUnitSummary
type IdentifyReplyCommandUnitSummaryBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(assertingNetworkBurden bool, restrikeTimingActive bool, remoteOFFInputAsserted bool, remoteONInputAsserted bool, localToggleEnabled bool, localToggleActiveState bool, clockGenerationEnabled bool, unitGeneratingClock bool) IdentifyReplyCommandUnitSummaryBuilder
	// WithAssertingNetworkBurden adds AssertingNetworkBurden (property field)
	WithAssertingNetworkBurden(bool) IdentifyReplyCommandUnitSummaryBuilder
	// WithRestrikeTimingActive adds RestrikeTimingActive (property field)
	WithRestrikeTimingActive(bool) IdentifyReplyCommandUnitSummaryBuilder
	// WithRemoteOFFInputAsserted adds RemoteOFFInputAsserted (property field)
	WithRemoteOFFInputAsserted(bool) IdentifyReplyCommandUnitSummaryBuilder
	// WithRemoteONInputAsserted adds RemoteONInputAsserted (property field)
	WithRemoteONInputAsserted(bool) IdentifyReplyCommandUnitSummaryBuilder
	// WithLocalToggleEnabled adds LocalToggleEnabled (property field)
	WithLocalToggleEnabled(bool) IdentifyReplyCommandUnitSummaryBuilder
	// WithLocalToggleActiveState adds LocalToggleActiveState (property field)
	WithLocalToggleActiveState(bool) IdentifyReplyCommandUnitSummaryBuilder
	// WithClockGenerationEnabled adds ClockGenerationEnabled (property field)
	WithClockGenerationEnabled(bool) IdentifyReplyCommandUnitSummaryBuilder
	// WithUnitGeneratingClock adds UnitGeneratingClock (property field)
	WithUnitGeneratingClock(bool) IdentifyReplyCommandUnitSummaryBuilder
	// Build builds the IdentifyReplyCommandUnitSummary or returns an error if something is wrong
	Build() (IdentifyReplyCommandUnitSummary, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() IdentifyReplyCommandUnitSummary
}

// NewIdentifyReplyCommandUnitSummaryBuilder() creates a IdentifyReplyCommandUnitSummaryBuilder
func NewIdentifyReplyCommandUnitSummaryBuilder() IdentifyReplyCommandUnitSummaryBuilder {
	return &_IdentifyReplyCommandUnitSummaryBuilder{_IdentifyReplyCommandUnitSummary: new(_IdentifyReplyCommandUnitSummary)}
}

type _IdentifyReplyCommandUnitSummaryBuilder struct {
	*_IdentifyReplyCommandUnitSummary

	err *utils.MultiError
}

var _ (IdentifyReplyCommandUnitSummaryBuilder) = (*_IdentifyReplyCommandUnitSummaryBuilder)(nil)

func (b *_IdentifyReplyCommandUnitSummaryBuilder) WithMandatoryFields(assertingNetworkBurden bool, restrikeTimingActive bool, remoteOFFInputAsserted bool, remoteONInputAsserted bool, localToggleEnabled bool, localToggleActiveState bool, clockGenerationEnabled bool, unitGeneratingClock bool) IdentifyReplyCommandUnitSummaryBuilder {
	return b.WithAssertingNetworkBurden(assertingNetworkBurden).WithRestrikeTimingActive(restrikeTimingActive).WithRemoteOFFInputAsserted(remoteOFFInputAsserted).WithRemoteONInputAsserted(remoteONInputAsserted).WithLocalToggleEnabled(localToggleEnabled).WithLocalToggleActiveState(localToggleActiveState).WithClockGenerationEnabled(clockGenerationEnabled).WithUnitGeneratingClock(unitGeneratingClock)
}

func (b *_IdentifyReplyCommandUnitSummaryBuilder) WithAssertingNetworkBurden(assertingNetworkBurden bool) IdentifyReplyCommandUnitSummaryBuilder {
	b.AssertingNetworkBurden = assertingNetworkBurden
	return b
}

func (b *_IdentifyReplyCommandUnitSummaryBuilder) WithRestrikeTimingActive(restrikeTimingActive bool) IdentifyReplyCommandUnitSummaryBuilder {
	b.RestrikeTimingActive = restrikeTimingActive
	return b
}

func (b *_IdentifyReplyCommandUnitSummaryBuilder) WithRemoteOFFInputAsserted(remoteOFFInputAsserted bool) IdentifyReplyCommandUnitSummaryBuilder {
	b.RemoteOFFInputAsserted = remoteOFFInputAsserted
	return b
}

func (b *_IdentifyReplyCommandUnitSummaryBuilder) WithRemoteONInputAsserted(remoteONInputAsserted bool) IdentifyReplyCommandUnitSummaryBuilder {
	b.RemoteONInputAsserted = remoteONInputAsserted
	return b
}

func (b *_IdentifyReplyCommandUnitSummaryBuilder) WithLocalToggleEnabled(localToggleEnabled bool) IdentifyReplyCommandUnitSummaryBuilder {
	b.LocalToggleEnabled = localToggleEnabled
	return b
}

func (b *_IdentifyReplyCommandUnitSummaryBuilder) WithLocalToggleActiveState(localToggleActiveState bool) IdentifyReplyCommandUnitSummaryBuilder {
	b.LocalToggleActiveState = localToggleActiveState
	return b
}

func (b *_IdentifyReplyCommandUnitSummaryBuilder) WithClockGenerationEnabled(clockGenerationEnabled bool) IdentifyReplyCommandUnitSummaryBuilder {
	b.ClockGenerationEnabled = clockGenerationEnabled
	return b
}

func (b *_IdentifyReplyCommandUnitSummaryBuilder) WithUnitGeneratingClock(unitGeneratingClock bool) IdentifyReplyCommandUnitSummaryBuilder {
	b.UnitGeneratingClock = unitGeneratingClock
	return b
}

func (b *_IdentifyReplyCommandUnitSummaryBuilder) Build() (IdentifyReplyCommandUnitSummary, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._IdentifyReplyCommandUnitSummary.deepCopy(), nil
}

func (b *_IdentifyReplyCommandUnitSummaryBuilder) MustBuild() IdentifyReplyCommandUnitSummary {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_IdentifyReplyCommandUnitSummaryBuilder) DeepCopy() any {
	_copy := b.CreateIdentifyReplyCommandUnitSummaryBuilder().(*_IdentifyReplyCommandUnitSummaryBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateIdentifyReplyCommandUnitSummaryBuilder creates a IdentifyReplyCommandUnitSummaryBuilder
func (b *_IdentifyReplyCommandUnitSummary) CreateIdentifyReplyCommandUnitSummaryBuilder() IdentifyReplyCommandUnitSummaryBuilder {
	if b == nil {
		return NewIdentifyReplyCommandUnitSummaryBuilder()
	}
	return &_IdentifyReplyCommandUnitSummaryBuilder{_IdentifyReplyCommandUnitSummary: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_IdentifyReplyCommandUnitSummary) GetAssertingNetworkBurden() bool {
	return m.AssertingNetworkBurden
}

func (m *_IdentifyReplyCommandUnitSummary) GetRestrikeTimingActive() bool {
	return m.RestrikeTimingActive
}

func (m *_IdentifyReplyCommandUnitSummary) GetRemoteOFFInputAsserted() bool {
	return m.RemoteOFFInputAsserted
}

func (m *_IdentifyReplyCommandUnitSummary) GetRemoteONInputAsserted() bool {
	return m.RemoteONInputAsserted
}

func (m *_IdentifyReplyCommandUnitSummary) GetLocalToggleEnabled() bool {
	return m.LocalToggleEnabled
}

func (m *_IdentifyReplyCommandUnitSummary) GetLocalToggleActiveState() bool {
	return m.LocalToggleActiveState
}

func (m *_IdentifyReplyCommandUnitSummary) GetClockGenerationEnabled() bool {
	return m.ClockGenerationEnabled
}

func (m *_IdentifyReplyCommandUnitSummary) GetUnitGeneratingClock() bool {
	return m.UnitGeneratingClock
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastIdentifyReplyCommandUnitSummary(structType any) IdentifyReplyCommandUnitSummary {
	if casted, ok := structType.(IdentifyReplyCommandUnitSummary); ok {
		return casted
	}
	if casted, ok := structType.(*IdentifyReplyCommandUnitSummary); ok {
		return *casted
	}
	return nil
}

func (m *_IdentifyReplyCommandUnitSummary) GetTypeName() string {
	return "IdentifyReplyCommandUnitSummary"
}

func (m *_IdentifyReplyCommandUnitSummary) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (assertingNetworkBurden)
	lengthInBits += 1

	// Simple field (restrikeTimingActive)
	lengthInBits += 1

	// Simple field (remoteOFFInputAsserted)
	lengthInBits += 1

	// Simple field (remoteONInputAsserted)
	lengthInBits += 1

	// Simple field (localToggleEnabled)
	lengthInBits += 1

	// Simple field (localToggleActiveState)
	lengthInBits += 1

	// Simple field (clockGenerationEnabled)
	lengthInBits += 1

	// Simple field (unitGeneratingClock)
	lengthInBits += 1

	return lengthInBits
}

func (m *_IdentifyReplyCommandUnitSummary) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func IdentifyReplyCommandUnitSummaryParse(ctx context.Context, theBytes []byte) (IdentifyReplyCommandUnitSummary, error) {
	return IdentifyReplyCommandUnitSummaryParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func IdentifyReplyCommandUnitSummaryParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (IdentifyReplyCommandUnitSummary, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (IdentifyReplyCommandUnitSummary, error) {
		return IdentifyReplyCommandUnitSummaryParseWithBuffer(ctx, readBuffer)
	}
}

func IdentifyReplyCommandUnitSummaryParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (IdentifyReplyCommandUnitSummary, error) {
	v, err := (&_IdentifyReplyCommandUnitSummary{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_IdentifyReplyCommandUnitSummary) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__identifyReplyCommandUnitSummary IdentifyReplyCommandUnitSummary, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("IdentifyReplyCommandUnitSummary"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for IdentifyReplyCommandUnitSummary")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	assertingNetworkBurden, err := ReadSimpleField(ctx, "assertingNetworkBurden", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'assertingNetworkBurden' field"))
	}
	m.AssertingNetworkBurden = assertingNetworkBurden

	restrikeTimingActive, err := ReadSimpleField(ctx, "restrikeTimingActive", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'restrikeTimingActive' field"))
	}
	m.RestrikeTimingActive = restrikeTimingActive

	remoteOFFInputAsserted, err := ReadSimpleField(ctx, "remoteOFFInputAsserted", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'remoteOFFInputAsserted' field"))
	}
	m.RemoteOFFInputAsserted = remoteOFFInputAsserted

	remoteONInputAsserted, err := ReadSimpleField(ctx, "remoteONInputAsserted", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'remoteONInputAsserted' field"))
	}
	m.RemoteONInputAsserted = remoteONInputAsserted

	localToggleEnabled, err := ReadSimpleField(ctx, "localToggleEnabled", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'localToggleEnabled' field"))
	}
	m.LocalToggleEnabled = localToggleEnabled

	localToggleActiveState, err := ReadSimpleField(ctx, "localToggleActiveState", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'localToggleActiveState' field"))
	}
	m.LocalToggleActiveState = localToggleActiveState

	clockGenerationEnabled, err := ReadSimpleField(ctx, "clockGenerationEnabled", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'clockGenerationEnabled' field"))
	}
	m.ClockGenerationEnabled = clockGenerationEnabled

	unitGeneratingClock, err := ReadSimpleField(ctx, "unitGeneratingClock", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'unitGeneratingClock' field"))
	}
	m.UnitGeneratingClock = unitGeneratingClock

	if closeErr := readBuffer.CloseContext("IdentifyReplyCommandUnitSummary"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for IdentifyReplyCommandUnitSummary")
	}

	return m, nil
}

func (m *_IdentifyReplyCommandUnitSummary) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_IdentifyReplyCommandUnitSummary) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("IdentifyReplyCommandUnitSummary"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for IdentifyReplyCommandUnitSummary")
	}

	if err := WriteSimpleField[bool](ctx, "assertingNetworkBurden", m.GetAssertingNetworkBurden(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'assertingNetworkBurden' field")
	}

	if err := WriteSimpleField[bool](ctx, "restrikeTimingActive", m.GetRestrikeTimingActive(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'restrikeTimingActive' field")
	}

	if err := WriteSimpleField[bool](ctx, "remoteOFFInputAsserted", m.GetRemoteOFFInputAsserted(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'remoteOFFInputAsserted' field")
	}

	if err := WriteSimpleField[bool](ctx, "remoteONInputAsserted", m.GetRemoteONInputAsserted(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'remoteONInputAsserted' field")
	}

	if err := WriteSimpleField[bool](ctx, "localToggleEnabled", m.GetLocalToggleEnabled(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'localToggleEnabled' field")
	}

	if err := WriteSimpleField[bool](ctx, "localToggleActiveState", m.GetLocalToggleActiveState(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'localToggleActiveState' field")
	}

	if err := WriteSimpleField[bool](ctx, "clockGenerationEnabled", m.GetClockGenerationEnabled(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'clockGenerationEnabled' field")
	}

	if err := WriteSimpleField[bool](ctx, "unitGeneratingClock", m.GetUnitGeneratingClock(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'unitGeneratingClock' field")
	}

	if popErr := writeBuffer.PopContext("IdentifyReplyCommandUnitSummary"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for IdentifyReplyCommandUnitSummary")
	}
	return nil
}

func (m *_IdentifyReplyCommandUnitSummary) IsIdentifyReplyCommandUnitSummary() {}

func (m *_IdentifyReplyCommandUnitSummary) DeepCopy() any {
	return m.deepCopy()
}

func (m *_IdentifyReplyCommandUnitSummary) deepCopy() *_IdentifyReplyCommandUnitSummary {
	if m == nil {
		return nil
	}
	_IdentifyReplyCommandUnitSummaryCopy := &_IdentifyReplyCommandUnitSummary{
		m.AssertingNetworkBurden,
		m.RestrikeTimingActive,
		m.RemoteOFFInputAsserted,
		m.RemoteONInputAsserted,
		m.LocalToggleEnabled,
		m.LocalToggleActiveState,
		m.ClockGenerationEnabled,
		m.UnitGeneratingClock,
	}
	return _IdentifyReplyCommandUnitSummaryCopy
}

func (m *_IdentifyReplyCommandUnitSummary) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}
