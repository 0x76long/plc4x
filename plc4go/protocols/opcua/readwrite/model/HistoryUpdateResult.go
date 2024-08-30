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
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// HistoryUpdateResult is the corresponding interface of HistoryUpdateResult
type HistoryUpdateResult interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetStatusCode returns StatusCode (property field)
	GetStatusCode() StatusCode
	// GetNoOfOperationResults returns NoOfOperationResults (property field)
	GetNoOfOperationResults() int32
	// GetOperationResults returns OperationResults (property field)
	GetOperationResults() []StatusCode
	// GetNoOfDiagnosticInfos returns NoOfDiagnosticInfos (property field)
	GetNoOfDiagnosticInfos() int32
	// GetDiagnosticInfos returns DiagnosticInfos (property field)
	GetDiagnosticInfos() []DiagnosticInfo
}

// HistoryUpdateResultExactly can be used when we want exactly this type and not a type which fulfills HistoryUpdateResult.
// This is useful for switch cases.
type HistoryUpdateResultExactly interface {
	HistoryUpdateResult
	isHistoryUpdateResult() bool
}

// _HistoryUpdateResult is the data-structure of this message
type _HistoryUpdateResult struct {
	*_ExtensionObjectDefinition
	StatusCode           StatusCode
	NoOfOperationResults int32
	OperationResults     []StatusCode
	NoOfDiagnosticInfos  int32
	DiagnosticInfos      []DiagnosticInfo
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_HistoryUpdateResult) GetIdentifier() string {
	return "697"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_HistoryUpdateResult) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_HistoryUpdateResult) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_HistoryUpdateResult) GetStatusCode() StatusCode {
	return m.StatusCode
}

func (m *_HistoryUpdateResult) GetNoOfOperationResults() int32 {
	return m.NoOfOperationResults
}

func (m *_HistoryUpdateResult) GetOperationResults() []StatusCode {
	return m.OperationResults
}

func (m *_HistoryUpdateResult) GetNoOfDiagnosticInfos() int32 {
	return m.NoOfDiagnosticInfos
}

func (m *_HistoryUpdateResult) GetDiagnosticInfos() []DiagnosticInfo {
	return m.DiagnosticInfos
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewHistoryUpdateResult factory function for _HistoryUpdateResult
func NewHistoryUpdateResult(statusCode StatusCode, noOfOperationResults int32, operationResults []StatusCode, noOfDiagnosticInfos int32, diagnosticInfos []DiagnosticInfo) *_HistoryUpdateResult {
	_result := &_HistoryUpdateResult{
		StatusCode:                 statusCode,
		NoOfOperationResults:       noOfOperationResults,
		OperationResults:           operationResults,
		NoOfDiagnosticInfos:        noOfDiagnosticInfos,
		DiagnosticInfos:            diagnosticInfos,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastHistoryUpdateResult(structType any) HistoryUpdateResult {
	if casted, ok := structType.(HistoryUpdateResult); ok {
		return casted
	}
	if casted, ok := structType.(*HistoryUpdateResult); ok {
		return *casted
	}
	return nil
}

func (m *_HistoryUpdateResult) GetTypeName() string {
	return "HistoryUpdateResult"
}

func (m *_HistoryUpdateResult) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (statusCode)
	lengthInBits += m.StatusCode.GetLengthInBits(ctx)

	// Simple field (noOfOperationResults)
	lengthInBits += 32

	// Array field
	if len(m.OperationResults) > 0 {
		for _curItem, element := range m.OperationResults {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.OperationResults), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (noOfDiagnosticInfos)
	lengthInBits += 32

	// Array field
	if len(m.DiagnosticInfos) > 0 {
		for _curItem, element := range m.DiagnosticInfos {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.DiagnosticInfos), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_HistoryUpdateResult) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func HistoryUpdateResultParse(ctx context.Context, theBytes []byte, identifier string) (HistoryUpdateResult, error) {
	return HistoryUpdateResultParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func HistoryUpdateResultParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (HistoryUpdateResult, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("HistoryUpdateResult"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for HistoryUpdateResult")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (statusCode)
	if pullErr := readBuffer.PullContext("statusCode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for statusCode")
	}
	_statusCode, _statusCodeErr := StatusCodeParseWithBuffer(ctx, readBuffer)
	if _statusCodeErr != nil {
		return nil, errors.Wrap(_statusCodeErr, "Error parsing 'statusCode' field of HistoryUpdateResult")
	}
	statusCode := _statusCode.(StatusCode)
	if closeErr := readBuffer.CloseContext("statusCode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for statusCode")
	}

	// Simple Field (noOfOperationResults)
	_noOfOperationResults, _noOfOperationResultsErr := /*TODO: migrate me*/ readBuffer.ReadInt32("noOfOperationResults", 32)
	if _noOfOperationResultsErr != nil {
		return nil, errors.Wrap(_noOfOperationResultsErr, "Error parsing 'noOfOperationResults' field of HistoryUpdateResult")
	}
	noOfOperationResults := _noOfOperationResults

	operationResults, err := ReadCountArrayField[StatusCode](ctx, "operationResults", ReadComplex[StatusCode](StatusCodeParseWithBuffer, readBuffer), uint64(noOfOperationResults))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'operationResults' field"))
	}

	// Simple Field (noOfDiagnosticInfos)
	_noOfDiagnosticInfos, _noOfDiagnosticInfosErr := /*TODO: migrate me*/ readBuffer.ReadInt32("noOfDiagnosticInfos", 32)
	if _noOfDiagnosticInfosErr != nil {
		return nil, errors.Wrap(_noOfDiagnosticInfosErr, "Error parsing 'noOfDiagnosticInfos' field of HistoryUpdateResult")
	}
	noOfDiagnosticInfos := _noOfDiagnosticInfos

	diagnosticInfos, err := ReadCountArrayField[DiagnosticInfo](ctx, "diagnosticInfos", ReadComplex[DiagnosticInfo](DiagnosticInfoParseWithBuffer, readBuffer), uint64(noOfDiagnosticInfos))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'diagnosticInfos' field"))
	}

	if closeErr := readBuffer.CloseContext("HistoryUpdateResult"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for HistoryUpdateResult")
	}

	// Create a partially initialized instance
	_child := &_HistoryUpdateResult{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		StatusCode:                 statusCode,
		NoOfOperationResults:       noOfOperationResults,
		OperationResults:           operationResults,
		NoOfDiagnosticInfos:        noOfDiagnosticInfos,
		DiagnosticInfos:            diagnosticInfos,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_HistoryUpdateResult) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_HistoryUpdateResult) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("HistoryUpdateResult"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for HistoryUpdateResult")
		}

		// Simple Field (statusCode)
		if pushErr := writeBuffer.PushContext("statusCode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for statusCode")
		}
		_statusCodeErr := writeBuffer.WriteSerializable(ctx, m.GetStatusCode())
		if popErr := writeBuffer.PopContext("statusCode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for statusCode")
		}
		if _statusCodeErr != nil {
			return errors.Wrap(_statusCodeErr, "Error serializing 'statusCode' field")
		}

		// Simple Field (noOfOperationResults)
		noOfOperationResults := int32(m.GetNoOfOperationResults())
		_noOfOperationResultsErr := /*TODO: migrate me*/ writeBuffer.WriteInt32("noOfOperationResults", 32, int32((noOfOperationResults)))
		if _noOfOperationResultsErr != nil {
			return errors.Wrap(_noOfOperationResultsErr, "Error serializing 'noOfOperationResults' field")
		}

		// Array Field (operationResults)
		if pushErr := writeBuffer.PushContext("operationResults", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for operationResults")
		}
		for _curItem, _element := range m.GetOperationResults() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetOperationResults()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'operationResults' field")
			}
		}
		if popErr := writeBuffer.PopContext("operationResults", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for operationResults")
		}

		// Simple Field (noOfDiagnosticInfos)
		noOfDiagnosticInfos := int32(m.GetNoOfDiagnosticInfos())
		_noOfDiagnosticInfosErr := /*TODO: migrate me*/ writeBuffer.WriteInt32("noOfDiagnosticInfos", 32, int32((noOfDiagnosticInfos)))
		if _noOfDiagnosticInfosErr != nil {
			return errors.Wrap(_noOfDiagnosticInfosErr, "Error serializing 'noOfDiagnosticInfos' field")
		}

		// Array Field (diagnosticInfos)
		if pushErr := writeBuffer.PushContext("diagnosticInfos", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for diagnosticInfos")
		}
		for _curItem, _element := range m.GetDiagnosticInfos() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetDiagnosticInfos()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'diagnosticInfos' field")
			}
		}
		if popErr := writeBuffer.PopContext("diagnosticInfos", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for diagnosticInfos")
		}

		if popErr := writeBuffer.PopContext("HistoryUpdateResult"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for HistoryUpdateResult")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_HistoryUpdateResult) isHistoryUpdateResult() bool {
	return true
}

func (m *_HistoryUpdateResult) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
