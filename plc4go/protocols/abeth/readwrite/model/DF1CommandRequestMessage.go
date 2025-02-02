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

// DF1CommandRequestMessage is the corresponding interface of DF1CommandRequestMessage
type DF1CommandRequestMessage interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	DF1RequestMessage
	// GetCommand returns Command (property field)
	GetCommand() DF1RequestCommand
	// IsDF1CommandRequestMessage is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsDF1CommandRequestMessage()
	// CreateBuilder creates a DF1CommandRequestMessageBuilder
	CreateDF1CommandRequestMessageBuilder() DF1CommandRequestMessageBuilder
}

// _DF1CommandRequestMessage is the data-structure of this message
type _DF1CommandRequestMessage struct {
	DF1RequestMessageContract
	Command DF1RequestCommand
}

var _ DF1CommandRequestMessage = (*_DF1CommandRequestMessage)(nil)
var _ DF1RequestMessageRequirements = (*_DF1CommandRequestMessage)(nil)

// NewDF1CommandRequestMessage factory function for _DF1CommandRequestMessage
func NewDF1CommandRequestMessage(destinationAddress uint8, sourceAddress uint8, status uint8, transactionCounter uint16, command DF1RequestCommand) *_DF1CommandRequestMessage {
	if command == nil {
		panic("command of type DF1RequestCommand for DF1CommandRequestMessage must not be nil")
	}
	_result := &_DF1CommandRequestMessage{
		DF1RequestMessageContract: NewDF1RequestMessage(destinationAddress, sourceAddress, status, transactionCounter),
		Command:                   command,
	}
	_result.DF1RequestMessageContract.(*_DF1RequestMessage)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// DF1CommandRequestMessageBuilder is a builder for DF1CommandRequestMessage
type DF1CommandRequestMessageBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(command DF1RequestCommand) DF1CommandRequestMessageBuilder
	// WithCommand adds Command (property field)
	WithCommand(DF1RequestCommand) DF1CommandRequestMessageBuilder
	// WithCommandBuilder adds Command (property field) which is build by the builder
	WithCommandBuilder(func(DF1RequestCommandBuilder) DF1RequestCommandBuilder) DF1CommandRequestMessageBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() DF1RequestMessageBuilder
	// Build builds the DF1CommandRequestMessage or returns an error if something is wrong
	Build() (DF1CommandRequestMessage, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() DF1CommandRequestMessage
}

// NewDF1CommandRequestMessageBuilder() creates a DF1CommandRequestMessageBuilder
func NewDF1CommandRequestMessageBuilder() DF1CommandRequestMessageBuilder {
	return &_DF1CommandRequestMessageBuilder{_DF1CommandRequestMessage: new(_DF1CommandRequestMessage)}
}

type _DF1CommandRequestMessageBuilder struct {
	*_DF1CommandRequestMessage

	parentBuilder *_DF1RequestMessageBuilder

	err *utils.MultiError
}

var _ (DF1CommandRequestMessageBuilder) = (*_DF1CommandRequestMessageBuilder)(nil)

func (b *_DF1CommandRequestMessageBuilder) setParent(contract DF1RequestMessageContract) {
	b.DF1RequestMessageContract = contract
	contract.(*_DF1RequestMessage)._SubType = b._DF1CommandRequestMessage
}

func (b *_DF1CommandRequestMessageBuilder) WithMandatoryFields(command DF1RequestCommand) DF1CommandRequestMessageBuilder {
	return b.WithCommand(command)
}

func (b *_DF1CommandRequestMessageBuilder) WithCommand(command DF1RequestCommand) DF1CommandRequestMessageBuilder {
	b.Command = command
	return b
}

func (b *_DF1CommandRequestMessageBuilder) WithCommandBuilder(builderSupplier func(DF1RequestCommandBuilder) DF1RequestCommandBuilder) DF1CommandRequestMessageBuilder {
	builder := builderSupplier(b.Command.CreateDF1RequestCommandBuilder())
	var err error
	b.Command, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "DF1RequestCommandBuilder failed"))
	}
	return b
}

func (b *_DF1CommandRequestMessageBuilder) Build() (DF1CommandRequestMessage, error) {
	if b.Command == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'command' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._DF1CommandRequestMessage.deepCopy(), nil
}

func (b *_DF1CommandRequestMessageBuilder) MustBuild() DF1CommandRequestMessage {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_DF1CommandRequestMessageBuilder) Done() DF1RequestMessageBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewDF1RequestMessageBuilder().(*_DF1RequestMessageBuilder)
	}
	return b.parentBuilder
}

func (b *_DF1CommandRequestMessageBuilder) buildForDF1RequestMessage() (DF1RequestMessage, error) {
	return b.Build()
}

func (b *_DF1CommandRequestMessageBuilder) DeepCopy() any {
	_copy := b.CreateDF1CommandRequestMessageBuilder().(*_DF1CommandRequestMessageBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateDF1CommandRequestMessageBuilder creates a DF1CommandRequestMessageBuilder
func (b *_DF1CommandRequestMessage) CreateDF1CommandRequestMessageBuilder() DF1CommandRequestMessageBuilder {
	if b == nil {
		return NewDF1CommandRequestMessageBuilder()
	}
	return &_DF1CommandRequestMessageBuilder{_DF1CommandRequestMessage: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DF1CommandRequestMessage) GetCommandCode() uint8 {
	return 0x0F
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DF1CommandRequestMessage) GetParent() DF1RequestMessageContract {
	return m.DF1RequestMessageContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DF1CommandRequestMessage) GetCommand() DF1RequestCommand {
	return m.Command
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastDF1CommandRequestMessage(structType any) DF1CommandRequestMessage {
	if casted, ok := structType.(DF1CommandRequestMessage); ok {
		return casted
	}
	if casted, ok := structType.(*DF1CommandRequestMessage); ok {
		return *casted
	}
	return nil
}

func (m *_DF1CommandRequestMessage) GetTypeName() string {
	return "DF1CommandRequestMessage"
}

func (m *_DF1CommandRequestMessage) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.DF1RequestMessageContract.(*_DF1RequestMessage).getLengthInBits(ctx))

	// Simple field (command)
	lengthInBits += m.Command.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_DF1CommandRequestMessage) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_DF1CommandRequestMessage) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_DF1RequestMessage) (__dF1CommandRequestMessage DF1CommandRequestMessage, err error) {
	m.DF1RequestMessageContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DF1CommandRequestMessage"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DF1CommandRequestMessage")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	command, err := ReadSimpleField[DF1RequestCommand](ctx, "command", ReadComplex[DF1RequestCommand](DF1RequestCommandParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'command' field"))
	}
	m.Command = command

	if closeErr := readBuffer.CloseContext("DF1CommandRequestMessage"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DF1CommandRequestMessage")
	}

	return m, nil
}

func (m *_DF1CommandRequestMessage) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DF1CommandRequestMessage) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DF1CommandRequestMessage"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DF1CommandRequestMessage")
		}

		if err := WriteSimpleField[DF1RequestCommand](ctx, "command", m.GetCommand(), WriteComplex[DF1RequestCommand](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'command' field")
		}

		if popErr := writeBuffer.PopContext("DF1CommandRequestMessage"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DF1CommandRequestMessage")
		}
		return nil
	}
	return m.DF1RequestMessageContract.(*_DF1RequestMessage).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DF1CommandRequestMessage) IsDF1CommandRequestMessage() {}

func (m *_DF1CommandRequestMessage) DeepCopy() any {
	return m.deepCopy()
}

func (m *_DF1CommandRequestMessage) deepCopy() *_DF1CommandRequestMessage {
	if m == nil {
		return nil
	}
	_DF1CommandRequestMessageCopy := &_DF1CommandRequestMessage{
		m.DF1RequestMessageContract.(*_DF1RequestMessage).deepCopy(),
		utils.DeepCopy[DF1RequestCommand](m.Command),
	}
	_DF1CommandRequestMessageCopy.DF1RequestMessageContract.(*_DF1RequestMessage)._SubType = m
	return _DF1CommandRequestMessageCopy
}

func (m *_DF1CommandRequestMessage) String() string {
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
