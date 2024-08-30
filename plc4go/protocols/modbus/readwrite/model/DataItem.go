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
	api "github.com/apache/plc4x/plc4go/pkg/api/values"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/apache/plc4x/plc4go/spi/values"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

func DataItemParse(ctx context.Context, theBytes []byte, dataType ModbusDataType, numberOfValues uint16, bigEndian bool) (api.PlcValue, error) {
	return DataItemParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), dataType, numberOfValues, bigEndian)
}

func DataItemParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, dataType ModbusDataType, numberOfValues uint16, bigEndian bool) (api.PlcValue, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	readBuffer.PullContext("DataItem")
	switch {
	case dataType == ModbusDataType_BOOL && numberOfValues == uint16(1) && bigEndian == bool(true): // BOOL
		// Reserved Field (Just skip the bytes)
		if _, _err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint16("reserved", 15); _err != nil {
			return nil, errors.Wrap(_err, "Error parsing reserved field")
		}

		// Simple Field (value)
		value, _valueErr := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadBit("value")
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcBOOL(value), nil
	case dataType == ModbusDataType_BOOL && numberOfValues == uint16(1) && bigEndian == bool(false): // BOOL
		// Reserved Field (Just skip the bytes)
		if _, _err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint8("reserved", 7); _err != nil {
			return nil, errors.Wrap(_err, "Error parsing reserved field")
		}

		// Simple Field (value)
		value, _valueErr := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadBit("value")
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)

		// Reserved Field (Just skip the bytes)
		if _, _err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint8("reserved", 8); _err != nil {
			return nil, errors.Wrap(_err, "Error parsing reserved field")
		}
		readBuffer.CloseContext("DataItem")
		return values.NewPlcBOOL(value), nil
	case dataType == ModbusDataType_BOOL: // List
		// Array Field (value)
		var value []api.PlcValue
		for i := 0; i < int(numberOfValues); i++ {
			_item, _itemErr := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadBit("value")
			if _itemErr != nil {
				return nil, errors.Wrap(_itemErr, "Error parsing 'value' field")
			}
			value = append(value, values.NewPlcBOOL(_item))
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcList(value), nil
	case dataType == ModbusDataType_BYTE && numberOfValues == uint16(1) && bigEndian == bool(true): // BYTE
		// Reserved Field (Just skip the bytes)
		if _, _err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint8("reserved", 8); _err != nil {
			return nil, errors.Wrap(_err, "Error parsing reserved field")
		}

		// Simple Field (value)
		value, _valueErr := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint8("value", 8)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcBYTE(value), nil
	case dataType == ModbusDataType_BYTE && numberOfValues == uint16(1) && bigEndian == bool(false): // BYTE
		// Simple Field (value)
		value, _valueErr := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint8("value", 8)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)

		// Reserved Field (Just skip the bytes)
		if _, _err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint8("reserved", 8); _err != nil {
			return nil, errors.Wrap(_err, "Error parsing reserved field")
		}
		readBuffer.CloseContext("DataItem")
		return values.NewPlcBYTE(value), nil
	case dataType == ModbusDataType_BYTE: // List
		// Array Field (value)
		var value []api.PlcValue
		for i := 0; i < int((numberOfValues)*(8)); i++ {
			_item, _itemErr := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadBit("value")
			if _itemErr != nil {
				return nil, errors.Wrap(_itemErr, "Error parsing 'value' field")
			}
			value = append(value, values.NewPlcBOOL(_item))
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcList(value), nil
	case dataType == ModbusDataType_WORD: // WORD
		// Simple Field (value)
		value, _valueErr := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint16("value", 16)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcWORD(value), nil
	case dataType == ModbusDataType_DWORD: // DWORD
		// Simple Field (value)
		value, _valueErr := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint32("value", 32)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcDWORD(value), nil
	case dataType == ModbusDataType_LWORD: // LWORD
		// Simple Field (value)
		value, _valueErr := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint64("value", 64)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcLWORD(value), nil
	case dataType == ModbusDataType_SINT && numberOfValues == uint16(1) && bigEndian == bool(true): // SINT
		// Reserved Field (Just skip the bytes)
		if _, _err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint8("reserved", 8); _err != nil {
			return nil, errors.Wrap(_err, "Error parsing reserved field")
		}

		// Simple Field (value)
		value, _valueErr := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadInt8("value", 8)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcSINT(value), nil
	case dataType == ModbusDataType_SINT && numberOfValues == uint16(1) && bigEndian == bool(false): // SINT
		// Simple Field (value)
		value, _valueErr := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadInt8("value", 8)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)

		// Reserved Field (Just skip the bytes)
		if _, _err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint8("reserved", 8); _err != nil {
			return nil, errors.Wrap(_err, "Error parsing reserved field")
		}
		readBuffer.CloseContext("DataItem")
		return values.NewPlcSINT(value), nil
	case dataType == ModbusDataType_SINT: // List
		// Array Field (value)
		var value []api.PlcValue
		for i := 0; i < int(numberOfValues); i++ {
			_item, _itemErr := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadInt8("value", 8)
			if _itemErr != nil {
				return nil, errors.Wrap(_itemErr, "Error parsing 'value' field")
			}
			value = append(value, values.NewPlcSINT(_item))
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcList(value), nil
	case dataType == ModbusDataType_INT && numberOfValues == uint16(1): // INT
		// Simple Field (value)
		value, _valueErr := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadInt16("value", 16)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcINT(value), nil
	case dataType == ModbusDataType_INT: // List
		// Array Field (value)
		var value []api.PlcValue
		for i := 0; i < int(numberOfValues); i++ {
			_item, _itemErr := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadInt16("value", 16)
			if _itemErr != nil {
				return nil, errors.Wrap(_itemErr, "Error parsing 'value' field")
			}
			value = append(value, values.NewPlcINT(_item))
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcList(value), nil
	case dataType == ModbusDataType_DINT && numberOfValues == uint16(1): // DINT
		// Simple Field (value)
		value, _valueErr := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadInt32("value", 32)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcDINT(value), nil
	case dataType == ModbusDataType_DINT: // List
		// Array Field (value)
		var value []api.PlcValue
		for i := 0; i < int(numberOfValues); i++ {
			_item, _itemErr := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadInt32("value", 32)
			if _itemErr != nil {
				return nil, errors.Wrap(_itemErr, "Error parsing 'value' field")
			}
			value = append(value, values.NewPlcDINT(_item))
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcList(value), nil
	case dataType == ModbusDataType_LINT && numberOfValues == uint16(1): // LINT
		// Simple Field (value)
		value, _valueErr := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadInt64("value", 64)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcLINT(value), nil
	case dataType == ModbusDataType_LINT: // List
		// Array Field (value)
		var value []api.PlcValue
		for i := 0; i < int(numberOfValues); i++ {
			_item, _itemErr := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadInt64("value", 64)
			if _itemErr != nil {
				return nil, errors.Wrap(_itemErr, "Error parsing 'value' field")
			}
			value = append(value, values.NewPlcLINT(_item))
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcList(value), nil
	case dataType == ModbusDataType_USINT && numberOfValues == uint16(1) && bigEndian == bool(true): // USINT
		// Reserved Field (Just skip the bytes)
		if _, _err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint8("reserved", 8); _err != nil {
			return nil, errors.Wrap(_err, "Error parsing reserved field")
		}

		// Simple Field (value)
		value, _valueErr := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint8("value", 8)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcUSINT(value), nil
	case dataType == ModbusDataType_USINT && numberOfValues == uint16(1) && bigEndian == bool(false): // USINT
		// Simple Field (value)
		value, _valueErr := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint8("value", 8)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)

		// Reserved Field (Just skip the bytes)
		if _, _err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint8("reserved", 8); _err != nil {
			return nil, errors.Wrap(_err, "Error parsing reserved field")
		}
		readBuffer.CloseContext("DataItem")
		return values.NewPlcUSINT(value), nil
	case dataType == ModbusDataType_USINT: // List
		// Array Field (value)
		var value []api.PlcValue
		for i := 0; i < int(numberOfValues); i++ {
			_item, _itemErr := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint8("value", 8)
			if _itemErr != nil {
				return nil, errors.Wrap(_itemErr, "Error parsing 'value' field")
			}
			value = append(value, values.NewPlcUSINT(_item))
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcList(value), nil
	case dataType == ModbusDataType_UINT && numberOfValues == uint16(1): // UINT
		// Simple Field (value)
		value, _valueErr := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint16("value", 16)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcUINT(value), nil
	case dataType == ModbusDataType_UINT: // List
		// Array Field (value)
		var value []api.PlcValue
		for i := 0; i < int(numberOfValues); i++ {
			_item, _itemErr := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint16("value", 16)
			if _itemErr != nil {
				return nil, errors.Wrap(_itemErr, "Error parsing 'value' field")
			}
			value = append(value, values.NewPlcUINT(_item))
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcList(value), nil
	case dataType == ModbusDataType_UDINT && numberOfValues == uint16(1): // UDINT
		// Simple Field (value)
		value, _valueErr := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint32("value", 32)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcUDINT(value), nil
	case dataType == ModbusDataType_UDINT: // List
		// Array Field (value)
		var value []api.PlcValue
		for i := 0; i < int(numberOfValues); i++ {
			_item, _itemErr := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint32("value", 32)
			if _itemErr != nil {
				return nil, errors.Wrap(_itemErr, "Error parsing 'value' field")
			}
			value = append(value, values.NewPlcUDINT(_item))
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcList(value), nil
	case dataType == ModbusDataType_ULINT && numberOfValues == uint16(1): // ULINT
		// Simple Field (value)
		value, _valueErr := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint64("value", 64)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcULINT(value), nil
	case dataType == ModbusDataType_ULINT: // List
		// Array Field (value)
		var value []api.PlcValue
		for i := 0; i < int(numberOfValues); i++ {
			_item, _itemErr := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint64("value", 64)
			if _itemErr != nil {
				return nil, errors.Wrap(_itemErr, "Error parsing 'value' field")
			}
			value = append(value, values.NewPlcULINT(_item))
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcList(value), nil
	case dataType == ModbusDataType_REAL && numberOfValues == uint16(1): // REAL
		// Simple Field (value)
		value, _valueErr := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadFloat32("value", 32)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcREAL(value), nil
	case dataType == ModbusDataType_REAL: // List
		// Array Field (value)
		var value []api.PlcValue
		for i := 0; i < int(numberOfValues); i++ {
			_item, _itemErr := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadFloat32("value", 32)
			if _itemErr != nil {
				return nil, errors.Wrap(_itemErr, "Error parsing 'value' field")
			}
			value = append(value, values.NewPlcREAL(_item))
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcList(value), nil
	case dataType == ModbusDataType_LREAL && numberOfValues == uint16(1): // LREAL
		// Simple Field (value)
		value, _valueErr := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadFloat64("value", 64)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcLREAL(value), nil
	case dataType == ModbusDataType_LREAL: // List
		// Array Field (value)
		var value []api.PlcValue
		for i := 0; i < int(numberOfValues); i++ {
			_item, _itemErr := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadFloat64("value", 64)
			if _itemErr != nil {
				return nil, errors.Wrap(_itemErr, "Error parsing 'value' field")
			}
			value = append(value, values.NewPlcLREAL(_item))
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcList(value), nil
	case dataType == ModbusDataType_CHAR && numberOfValues == uint16(1): // CHAR
		// Simple Field (value)
		value, _valueErr := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadString("value", uint32(8), utils.WithEncoding("UTF-8"))
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcCHAR(value), nil
	case dataType == ModbusDataType_CHAR: // List
		// Array Field (value)
		var value []api.PlcValue
		for i := 0; i < int(numberOfValues); i++ {
			_item, _itemErr := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadString("value", uint32(8), utils.WithEncoding("UTF-8"))
			if _itemErr != nil {
				return nil, errors.Wrap(_itemErr, "Error parsing 'value' field")
			}
			value = append(value, values.NewPlcSTRING(_item))
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcList(value), nil
	case dataType == ModbusDataType_WCHAR && numberOfValues == uint16(1): // WCHAR
		// Simple Field (value)
		value, _valueErr := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadString("value", uint32(16), utils.WithEncoding("UTF-16"))
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcWCHAR(value), nil
	case dataType == ModbusDataType_WCHAR: // List
		// Array Field (value)
		var value []api.PlcValue
		for i := 0; i < int(numberOfValues); i++ {
			_item, _itemErr := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadString("value", uint32(16), utils.WithEncoding("UTF-16"))
			if _itemErr != nil {
				return nil, errors.Wrap(_itemErr, "Error parsing 'value' field")
			}
			value = append(value, values.NewPlcSTRING(_item))
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcList(value), nil
	}
	// TODO: add more info which type it is actually
	return nil, errors.New("unsupported type")
}

func DataItemSerialize(value api.PlcValue, dataType ModbusDataType, numberOfValues uint16, bigEndian bool) ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := DataItemSerializeWithWriteBuffer(context.Background(), wb, value, dataType, numberOfValues, bigEndian); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func DataItemSerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer, value api.PlcValue, dataType ModbusDataType, numberOfValues uint16, bigEndian bool) error {
	log := zerolog.Ctx(ctx)
	_ = log
	m := struct {
		DataType       ModbusDataType
		NumberOfValues uint16
		BigEndian      bool
	}{
		DataType:       dataType,
		NumberOfValues: numberOfValues,
		BigEndian:      bigEndian,
	}
	_ = m
	writeBuffer.PushContext("DataItem")
	switch {
	case dataType == ModbusDataType_BOOL && numberOfValues == uint16(1) && bigEndian == bool(true): // BOOL
		// Reserved Field (Just skip the bytes)
		if _err := /*TODO: migrate me*/ writeBuffer.WriteUint16("reserved", 15, uint16(uint16(0x0000))); _err != nil {
			return errors.Wrap(_err, "Error serializing reserved field")
		}

		// Simple Field (value)
		if _err := /*TODO: migrate me*/ writeBuffer.WriteBit("value", value.GetBool()); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataType == ModbusDataType_BOOL && numberOfValues == uint16(1) && bigEndian == bool(false): // BOOL
		// Reserved Field (Just skip the bytes)
		if _err := /*TODO: migrate me*/ writeBuffer.WriteUint8("reserved", 7, uint8(uint8(0x00))); _err != nil {
			return errors.Wrap(_err, "Error serializing reserved field")
		}

		// Simple Field (value)
		if _err := /*TODO: migrate me*/ writeBuffer.WriteBit("value", value.GetBool()); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}

		// Reserved Field (Just skip the bytes)
		if _err := /*TODO: migrate me*/ writeBuffer.WriteUint8("reserved", 8, uint8(uint8(0x00))); _err != nil {
			return errors.Wrap(_err, "Error serializing reserved field")
		}
	case dataType == ModbusDataType_BOOL: // List
		// Array Field (value)
		for i := uint32(0); i < uint32(m.NumberOfValues); i++ {
			_itemErr := /*TODO: migrate me*/ writeBuffer.WriteBit("", value.GetIndex(i).GetBool())
			if _itemErr != nil {
				return errors.Wrap(_itemErr, "Error serializing 'value' field")
			}
		}
	case dataType == ModbusDataType_BYTE && numberOfValues == uint16(1) && bigEndian == bool(true): // BYTE
		// Reserved Field (Just skip the bytes)
		if _err := /*TODO: migrate me*/ writeBuffer.WriteUint8("reserved", 8, uint8(uint8(0x00))); _err != nil {
			return errors.Wrap(_err, "Error serializing reserved field")
		}

		// Simple Field (value)
		if _err := /*TODO: migrate me*/ writeBuffer.WriteUint8("value", 8, uint8(value.GetUint8())); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataType == ModbusDataType_BYTE && numberOfValues == uint16(1) && bigEndian == bool(false): // BYTE
		// Simple Field (value)
		if _err := /*TODO: migrate me*/ writeBuffer.WriteUint8("value", 8, uint8(value.GetUint8())); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}

		// Reserved Field (Just skip the bytes)
		if _err := /*TODO: migrate me*/ writeBuffer.WriteUint8("reserved", 8, uint8(uint8(0x00))); _err != nil {
			return errors.Wrap(_err, "Error serializing reserved field")
		}
	case dataType == ModbusDataType_BYTE: // List
		// Array Field (value)
		for i := uint32(0); i < uint32((m.NumberOfValues)*(8)); i++ {
			_itemErr := /*TODO: migrate me*/ writeBuffer.WriteBit("", value.GetIndex(i).GetBool())
			if _itemErr != nil {
				return errors.Wrap(_itemErr, "Error serializing 'value' field")
			}
		}
	case dataType == ModbusDataType_WORD: // WORD
		// Simple Field (value)
		if _err := /*TODO: migrate me*/ writeBuffer.WriteUint16("value", 16, uint16(value.GetUint16())); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataType == ModbusDataType_DWORD: // DWORD
		// Simple Field (value)
		if _err := /*TODO: migrate me*/ writeBuffer.WriteUint32("value", 32, uint32(value.GetUint32())); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataType == ModbusDataType_LWORD: // LWORD
		// Simple Field (value)
		if _err := /*TODO: migrate me*/ writeBuffer.WriteUint64("value", 64, uint64(value.GetUint64())); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataType == ModbusDataType_SINT && numberOfValues == uint16(1) && bigEndian == bool(true): // SINT
		// Reserved Field (Just skip the bytes)
		if _err := /*TODO: migrate me*/ writeBuffer.WriteUint8("reserved", 8, uint8(uint8(0x00))); _err != nil {
			return errors.Wrap(_err, "Error serializing reserved field")
		}

		// Simple Field (value)
		if _err := /*TODO: migrate me*/ writeBuffer.WriteInt8("value", 8, int8(value.GetInt8())); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataType == ModbusDataType_SINT && numberOfValues == uint16(1) && bigEndian == bool(false): // SINT
		// Simple Field (value)
		if _err := /*TODO: migrate me*/ writeBuffer.WriteInt8("value", 8, int8(value.GetInt8())); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}

		// Reserved Field (Just skip the bytes)
		if _err := /*TODO: migrate me*/ writeBuffer.WriteUint8("reserved", 8, uint8(uint8(0x00))); _err != nil {
			return errors.Wrap(_err, "Error serializing reserved field")
		}
	case dataType == ModbusDataType_SINT: // List
		// Array Field (value)
		for i := uint32(0); i < uint32(m.NumberOfValues); i++ {
			_itemErr := /*TODO: migrate me*/ writeBuffer.WriteInt8("", 8, int8(value.GetIndex(i).GetInt8()))
			if _itemErr != nil {
				return errors.Wrap(_itemErr, "Error serializing 'value' field")
			}
		}
	case dataType == ModbusDataType_INT && numberOfValues == uint16(1): // INT
		// Simple Field (value)
		if _err := /*TODO: migrate me*/ writeBuffer.WriteInt16("value", 16, int16(value.GetInt16())); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataType == ModbusDataType_INT: // List
		// Array Field (value)
		for i := uint32(0); i < uint32(m.NumberOfValues); i++ {
			_itemErr := /*TODO: migrate me*/ writeBuffer.WriteInt16("", 16, int16(value.GetIndex(i).GetInt16()))
			if _itemErr != nil {
				return errors.Wrap(_itemErr, "Error serializing 'value' field")
			}
		}
	case dataType == ModbusDataType_DINT && numberOfValues == uint16(1): // DINT
		// Simple Field (value)
		if _err := /*TODO: migrate me*/ writeBuffer.WriteInt32("value", 32, int32(value.GetInt32())); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataType == ModbusDataType_DINT: // List
		// Array Field (value)
		for i := uint32(0); i < uint32(m.NumberOfValues); i++ {
			_itemErr := /*TODO: migrate me*/ writeBuffer.WriteInt32("", 32, int32(value.GetIndex(i).GetInt32()))
			if _itemErr != nil {
				return errors.Wrap(_itemErr, "Error serializing 'value' field")
			}
		}
	case dataType == ModbusDataType_LINT && numberOfValues == uint16(1): // LINT
		// Simple Field (value)
		if _err := /*TODO: migrate me*/ writeBuffer.WriteInt64("value", 64, int64(value.GetInt64())); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataType == ModbusDataType_LINT: // List
		// Array Field (value)
		for i := uint32(0); i < uint32(m.NumberOfValues); i++ {
			_itemErr := /*TODO: migrate me*/ writeBuffer.WriteInt64("", 64, int64(value.GetIndex(i).GetInt64()))
			if _itemErr != nil {
				return errors.Wrap(_itemErr, "Error serializing 'value' field")
			}
		}
	case dataType == ModbusDataType_USINT && numberOfValues == uint16(1) && bigEndian == bool(true): // USINT
		// Reserved Field (Just skip the bytes)
		if _err := /*TODO: migrate me*/ writeBuffer.WriteUint8("reserved", 8, uint8(uint8(0x00))); _err != nil {
			return errors.Wrap(_err, "Error serializing reserved field")
		}

		// Simple Field (value)
		if _err := /*TODO: migrate me*/ writeBuffer.WriteUint8("value", 8, uint8(value.GetUint8())); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataType == ModbusDataType_USINT && numberOfValues == uint16(1) && bigEndian == bool(false): // USINT
		// Simple Field (value)
		if _err := /*TODO: migrate me*/ writeBuffer.WriteUint8("value", 8, uint8(value.GetUint8())); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}

		// Reserved Field (Just skip the bytes)
		if _err := /*TODO: migrate me*/ writeBuffer.WriteUint8("reserved", 8, uint8(uint8(0x00))); _err != nil {
			return errors.Wrap(_err, "Error serializing reserved field")
		}
	case dataType == ModbusDataType_USINT: // List
		// Array Field (value)
		for i := uint32(0); i < uint32(m.NumberOfValues); i++ {
			_itemErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("", 8, uint8(value.GetIndex(i).GetUint8()))
			if _itemErr != nil {
				return errors.Wrap(_itemErr, "Error serializing 'value' field")
			}
		}
	case dataType == ModbusDataType_UINT && numberOfValues == uint16(1): // UINT
		// Simple Field (value)
		if _err := /*TODO: migrate me*/ writeBuffer.WriteUint16("value", 16, uint16(value.GetUint16())); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataType == ModbusDataType_UINT: // List
		// Array Field (value)
		for i := uint32(0); i < uint32(m.NumberOfValues); i++ {
			_itemErr := /*TODO: migrate me*/ writeBuffer.WriteUint16("", 16, uint16(value.GetIndex(i).GetUint16()))
			if _itemErr != nil {
				return errors.Wrap(_itemErr, "Error serializing 'value' field")
			}
		}
	case dataType == ModbusDataType_UDINT && numberOfValues == uint16(1): // UDINT
		// Simple Field (value)
		if _err := /*TODO: migrate me*/ writeBuffer.WriteUint32("value", 32, uint32(value.GetUint32())); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataType == ModbusDataType_UDINT: // List
		// Array Field (value)
		for i := uint32(0); i < uint32(m.NumberOfValues); i++ {
			_itemErr := /*TODO: migrate me*/ writeBuffer.WriteUint32("", 32, uint32(value.GetIndex(i).GetUint32()))
			if _itemErr != nil {
				return errors.Wrap(_itemErr, "Error serializing 'value' field")
			}
		}
	case dataType == ModbusDataType_ULINT && numberOfValues == uint16(1): // ULINT
		// Simple Field (value)
		if _err := /*TODO: migrate me*/ writeBuffer.WriteUint64("value", 64, uint64(value.GetUint64())); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataType == ModbusDataType_ULINT: // List
		// Array Field (value)
		for i := uint32(0); i < uint32(m.NumberOfValues); i++ {
			_itemErr := /*TODO: migrate me*/ writeBuffer.WriteUint64("", 64, uint64(value.GetIndex(i).GetUint64()))
			if _itemErr != nil {
				return errors.Wrap(_itemErr, "Error serializing 'value' field")
			}
		}
	case dataType == ModbusDataType_REAL && numberOfValues == uint16(1): // REAL
		// Simple Field (value)
		if _err := /*TODO: migrate me*/ writeBuffer.WriteFloat32("value", 32, value.GetFloat32()); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataType == ModbusDataType_REAL: // List
		// Array Field (value)
		for i := uint32(0); i < uint32(m.NumberOfValues); i++ {
			_itemErr := /*TODO: migrate me*/ writeBuffer.WriteFloat32("", 32, value.GetIndex(i).GetFloat32())
			if _itemErr != nil {
				return errors.Wrap(_itemErr, "Error serializing 'value' field")
			}
		}
	case dataType == ModbusDataType_LREAL && numberOfValues == uint16(1): // LREAL
		// Simple Field (value)
		if _err := /*TODO: migrate me*/ writeBuffer.WriteFloat64("value", 64, value.GetFloat64()); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataType == ModbusDataType_LREAL: // List
		// Array Field (value)
		for i := uint32(0); i < uint32(m.NumberOfValues); i++ {
			_itemErr := /*TODO: migrate me*/ writeBuffer.WriteFloat64("", 64, value.GetIndex(i).GetFloat64())
			if _itemErr != nil {
				return errors.Wrap(_itemErr, "Error serializing 'value' field")
			}
		}
	case dataType == ModbusDataType_CHAR && numberOfValues == uint16(1): // CHAR
		// Simple Field (value)
		if _err := /*TODO: migrate me*/ writeBuffer.WriteString("value", uint32(8), value.GetString(), utils.WithEncoding("UTF-8)")); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataType == ModbusDataType_CHAR: // List
		// Array Field (value)
		for i := uint32(0); i < uint32(m.NumberOfValues); i++ {
			_itemErr := /*TODO: migrate me*/ writeBuffer.WriteString("", uint32(8), value.GetIndex(i).GetString(), utils.WithEncoding("UTF-8)"))
			if _itemErr != nil {
				return errors.Wrap(_itemErr, "Error serializing 'value' field")
			}
		}
	case dataType == ModbusDataType_WCHAR && numberOfValues == uint16(1): // WCHAR
		// Simple Field (value)
		if _err := /*TODO: migrate me*/ writeBuffer.WriteString("value", uint32(16), value.GetString(), utils.WithEncoding("UTF-16)")); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataType == ModbusDataType_WCHAR: // List
		// Array Field (value)
		for i := uint32(0); i < uint32(m.NumberOfValues); i++ {
			_itemErr := /*TODO: migrate me*/ writeBuffer.WriteString("", uint32(16), value.GetIndex(i).GetString(), utils.WithEncoding("UTF-16)"))
			if _itemErr != nil {
				return errors.Wrap(_itemErr, "Error serializing 'value' field")
			}
		}
	default:
		// TODO: add more info which type it is actually
		return errors.New("unsupported type")
	}
	writeBuffer.PopContext("DataItem")
	return nil
}
