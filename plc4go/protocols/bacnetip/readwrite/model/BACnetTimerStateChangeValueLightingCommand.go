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

// BACnetTimerStateChangeValueLightingCommand is the corresponding interface of BACnetTimerStateChangeValueLightingCommand
type BACnetTimerStateChangeValueLightingCommand interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetTimerStateChangeValue
	// GetLigthingCommandValue returns LigthingCommandValue (property field)
	GetLigthingCommandValue() BACnetLightingCommandEnclosed
	// IsBACnetTimerStateChangeValueLightingCommand is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetTimerStateChangeValueLightingCommand()
	// CreateBuilder creates a BACnetTimerStateChangeValueLightingCommandBuilder
	CreateBACnetTimerStateChangeValueLightingCommandBuilder() BACnetTimerStateChangeValueLightingCommandBuilder
}

// _BACnetTimerStateChangeValueLightingCommand is the data-structure of this message
type _BACnetTimerStateChangeValueLightingCommand struct {
	BACnetTimerStateChangeValueContract
	LigthingCommandValue BACnetLightingCommandEnclosed
}

var _ BACnetTimerStateChangeValueLightingCommand = (*_BACnetTimerStateChangeValueLightingCommand)(nil)
var _ BACnetTimerStateChangeValueRequirements = (*_BACnetTimerStateChangeValueLightingCommand)(nil)

// NewBACnetTimerStateChangeValueLightingCommand factory function for _BACnetTimerStateChangeValueLightingCommand
func NewBACnetTimerStateChangeValueLightingCommand(peekedTagHeader BACnetTagHeader, ligthingCommandValue BACnetLightingCommandEnclosed, objectTypeArgument BACnetObjectType) *_BACnetTimerStateChangeValueLightingCommand {
	if ligthingCommandValue == nil {
		panic("ligthingCommandValue of type BACnetLightingCommandEnclosed for BACnetTimerStateChangeValueLightingCommand must not be nil")
	}
	_result := &_BACnetTimerStateChangeValueLightingCommand{
		BACnetTimerStateChangeValueContract: NewBACnetTimerStateChangeValue(peekedTagHeader, objectTypeArgument),
		LigthingCommandValue:                ligthingCommandValue,
	}
	_result.BACnetTimerStateChangeValueContract.(*_BACnetTimerStateChangeValue)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetTimerStateChangeValueLightingCommandBuilder is a builder for BACnetTimerStateChangeValueLightingCommand
type BACnetTimerStateChangeValueLightingCommandBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(ligthingCommandValue BACnetLightingCommandEnclosed) BACnetTimerStateChangeValueLightingCommandBuilder
	// WithLigthingCommandValue adds LigthingCommandValue (property field)
	WithLigthingCommandValue(BACnetLightingCommandEnclosed) BACnetTimerStateChangeValueLightingCommandBuilder
	// WithLigthingCommandValueBuilder adds LigthingCommandValue (property field) which is build by the builder
	WithLigthingCommandValueBuilder(func(BACnetLightingCommandEnclosedBuilder) BACnetLightingCommandEnclosedBuilder) BACnetTimerStateChangeValueLightingCommandBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetTimerStateChangeValueBuilder
	// Build builds the BACnetTimerStateChangeValueLightingCommand or returns an error if something is wrong
	Build() (BACnetTimerStateChangeValueLightingCommand, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetTimerStateChangeValueLightingCommand
}

// NewBACnetTimerStateChangeValueLightingCommandBuilder() creates a BACnetTimerStateChangeValueLightingCommandBuilder
func NewBACnetTimerStateChangeValueLightingCommandBuilder() BACnetTimerStateChangeValueLightingCommandBuilder {
	return &_BACnetTimerStateChangeValueLightingCommandBuilder{_BACnetTimerStateChangeValueLightingCommand: new(_BACnetTimerStateChangeValueLightingCommand)}
}

type _BACnetTimerStateChangeValueLightingCommandBuilder struct {
	*_BACnetTimerStateChangeValueLightingCommand

	parentBuilder *_BACnetTimerStateChangeValueBuilder

	err *utils.MultiError
}

var _ (BACnetTimerStateChangeValueLightingCommandBuilder) = (*_BACnetTimerStateChangeValueLightingCommandBuilder)(nil)

func (b *_BACnetTimerStateChangeValueLightingCommandBuilder) setParent(contract BACnetTimerStateChangeValueContract) {
	b.BACnetTimerStateChangeValueContract = contract
	contract.(*_BACnetTimerStateChangeValue)._SubType = b._BACnetTimerStateChangeValueLightingCommand
}

func (b *_BACnetTimerStateChangeValueLightingCommandBuilder) WithMandatoryFields(ligthingCommandValue BACnetLightingCommandEnclosed) BACnetTimerStateChangeValueLightingCommandBuilder {
	return b.WithLigthingCommandValue(ligthingCommandValue)
}

func (b *_BACnetTimerStateChangeValueLightingCommandBuilder) WithLigthingCommandValue(ligthingCommandValue BACnetLightingCommandEnclosed) BACnetTimerStateChangeValueLightingCommandBuilder {
	b.LigthingCommandValue = ligthingCommandValue
	return b
}

func (b *_BACnetTimerStateChangeValueLightingCommandBuilder) WithLigthingCommandValueBuilder(builderSupplier func(BACnetLightingCommandEnclosedBuilder) BACnetLightingCommandEnclosedBuilder) BACnetTimerStateChangeValueLightingCommandBuilder {
	builder := builderSupplier(b.LigthingCommandValue.CreateBACnetLightingCommandEnclosedBuilder())
	var err error
	b.LigthingCommandValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetLightingCommandEnclosedBuilder failed"))
	}
	return b
}

func (b *_BACnetTimerStateChangeValueLightingCommandBuilder) Build() (BACnetTimerStateChangeValueLightingCommand, error) {
	if b.LigthingCommandValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'ligthingCommandValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetTimerStateChangeValueLightingCommand.deepCopy(), nil
}

func (b *_BACnetTimerStateChangeValueLightingCommandBuilder) MustBuild() BACnetTimerStateChangeValueLightingCommand {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetTimerStateChangeValueLightingCommandBuilder) Done() BACnetTimerStateChangeValueBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetTimerStateChangeValueBuilder().(*_BACnetTimerStateChangeValueBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetTimerStateChangeValueLightingCommandBuilder) buildForBACnetTimerStateChangeValue() (BACnetTimerStateChangeValue, error) {
	return b.Build()
}

func (b *_BACnetTimerStateChangeValueLightingCommandBuilder) DeepCopy() any {
	_copy := b.CreateBACnetTimerStateChangeValueLightingCommandBuilder().(*_BACnetTimerStateChangeValueLightingCommandBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetTimerStateChangeValueLightingCommandBuilder creates a BACnetTimerStateChangeValueLightingCommandBuilder
func (b *_BACnetTimerStateChangeValueLightingCommand) CreateBACnetTimerStateChangeValueLightingCommandBuilder() BACnetTimerStateChangeValueLightingCommandBuilder {
	if b == nil {
		return NewBACnetTimerStateChangeValueLightingCommandBuilder()
	}
	return &_BACnetTimerStateChangeValueLightingCommandBuilder{_BACnetTimerStateChangeValueLightingCommand: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetTimerStateChangeValueLightingCommand) GetParent() BACnetTimerStateChangeValueContract {
	return m.BACnetTimerStateChangeValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTimerStateChangeValueLightingCommand) GetLigthingCommandValue() BACnetLightingCommandEnclosed {
	return m.LigthingCommandValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetTimerStateChangeValueLightingCommand(structType any) BACnetTimerStateChangeValueLightingCommand {
	if casted, ok := structType.(BACnetTimerStateChangeValueLightingCommand); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTimerStateChangeValueLightingCommand); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTimerStateChangeValueLightingCommand) GetTypeName() string {
	return "BACnetTimerStateChangeValueLightingCommand"
}

func (m *_BACnetTimerStateChangeValueLightingCommand) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetTimerStateChangeValueContract.(*_BACnetTimerStateChangeValue).getLengthInBits(ctx))

	// Simple field (ligthingCommandValue)
	lengthInBits += m.LigthingCommandValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetTimerStateChangeValueLightingCommand) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetTimerStateChangeValueLightingCommand) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetTimerStateChangeValue, objectTypeArgument BACnetObjectType) (__bACnetTimerStateChangeValueLightingCommand BACnetTimerStateChangeValueLightingCommand, err error) {
	m.BACnetTimerStateChangeValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTimerStateChangeValueLightingCommand"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTimerStateChangeValueLightingCommand")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	ligthingCommandValue, err := ReadSimpleField[BACnetLightingCommandEnclosed](ctx, "ligthingCommandValue", ReadComplex[BACnetLightingCommandEnclosed](BACnetLightingCommandEnclosedParseWithBufferProducer((uint8)(uint8(3))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'ligthingCommandValue' field"))
	}
	m.LigthingCommandValue = ligthingCommandValue

	if closeErr := readBuffer.CloseContext("BACnetTimerStateChangeValueLightingCommand"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTimerStateChangeValueLightingCommand")
	}

	return m, nil
}

func (m *_BACnetTimerStateChangeValueLightingCommand) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetTimerStateChangeValueLightingCommand) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetTimerStateChangeValueLightingCommand"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetTimerStateChangeValueLightingCommand")
		}

		if err := WriteSimpleField[BACnetLightingCommandEnclosed](ctx, "ligthingCommandValue", m.GetLigthingCommandValue(), WriteComplex[BACnetLightingCommandEnclosed](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'ligthingCommandValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetTimerStateChangeValueLightingCommand"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetTimerStateChangeValueLightingCommand")
		}
		return nil
	}
	return m.BACnetTimerStateChangeValueContract.(*_BACnetTimerStateChangeValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetTimerStateChangeValueLightingCommand) IsBACnetTimerStateChangeValueLightingCommand() {
}

func (m *_BACnetTimerStateChangeValueLightingCommand) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetTimerStateChangeValueLightingCommand) deepCopy() *_BACnetTimerStateChangeValueLightingCommand {
	if m == nil {
		return nil
	}
	_BACnetTimerStateChangeValueLightingCommandCopy := &_BACnetTimerStateChangeValueLightingCommand{
		m.BACnetTimerStateChangeValueContract.(*_BACnetTimerStateChangeValue).deepCopy(),
		utils.DeepCopy[BACnetLightingCommandEnclosed](m.LigthingCommandValue),
	}
	_BACnetTimerStateChangeValueLightingCommandCopy.BACnetTimerStateChangeValueContract.(*_BACnetTimerStateChangeValue)._SubType = m
	return _BACnetTimerStateChangeValueLightingCommandCopy
}

func (m *_BACnetTimerStateChangeValueLightingCommand) String() string {
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
