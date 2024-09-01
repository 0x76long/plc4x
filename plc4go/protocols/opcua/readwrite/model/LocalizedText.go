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

// LocalizedText is the corresponding interface of LocalizedText
type LocalizedText interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetTextSpecified returns TextSpecified (property field)
	GetTextSpecified() bool
	// GetLocaleSpecified returns LocaleSpecified (property field)
	GetLocaleSpecified() bool
	// GetLocale returns Locale (property field)
	GetLocale() PascalString
	// GetText returns Text (property field)
	GetText() PascalString
}

// LocalizedTextExactly can be used when we want exactly this type and not a type which fulfills LocalizedText.
// This is useful for switch cases.
type LocalizedTextExactly interface {
	LocalizedText
	isLocalizedText() bool
}

// _LocalizedText is the data-structure of this message
type _LocalizedText struct {
	TextSpecified   bool
	LocaleSpecified bool
	Locale          PascalString
	Text            PascalString
	// Reserved Fields
	reservedField0 *uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_LocalizedText) GetTextSpecified() bool {
	return m.TextSpecified
}

func (m *_LocalizedText) GetLocaleSpecified() bool {
	return m.LocaleSpecified
}

func (m *_LocalizedText) GetLocale() PascalString {
	return m.Locale
}

func (m *_LocalizedText) GetText() PascalString {
	return m.Text
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewLocalizedText factory function for _LocalizedText
func NewLocalizedText(textSpecified bool, localeSpecified bool, locale PascalString, text PascalString) *_LocalizedText {
	return &_LocalizedText{TextSpecified: textSpecified, LocaleSpecified: localeSpecified, Locale: locale, Text: text}
}

// Deprecated: use the interface for direct cast
func CastLocalizedText(structType any) LocalizedText {
	if casted, ok := structType.(LocalizedText); ok {
		return casted
	}
	if casted, ok := structType.(*LocalizedText); ok {
		return *casted
	}
	return nil
}

func (m *_LocalizedText) GetTypeName() string {
	return "LocalizedText"
}

func (m *_LocalizedText) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Reserved Field (reserved)
	lengthInBits += 6

	// Simple field (textSpecified)
	lengthInBits += 1

	// Simple field (localeSpecified)
	lengthInBits += 1

	// Optional Field (locale)
	if m.Locale != nil {
		lengthInBits += m.Locale.GetLengthInBits(ctx)
	}

	// Optional Field (text)
	if m.Text != nil {
		lengthInBits += m.Text.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_LocalizedText) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func LocalizedTextParse(ctx context.Context, theBytes []byte) (LocalizedText, error) {
	return LocalizedTextParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func LocalizedTextParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (LocalizedText, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (LocalizedText, error) {
		return LocalizedTextParseWithBuffer(ctx, readBuffer)
	}
}

func LocalizedTextParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (LocalizedText, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("LocalizedText"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for LocalizedText")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(6)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}

	textSpecified, err := ReadSimpleField(ctx, "textSpecified", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'textSpecified' field"))
	}

	localeSpecified, err := ReadSimpleField(ctx, "localeSpecified", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'localeSpecified' field"))
	}

	_locale, err := ReadOptionalField[PascalString](ctx, "locale", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer), localeSpecified)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'locale' field"))
	}
	var locale PascalString
	if _locale != nil {
		locale = *_locale
	}

	_text, err := ReadOptionalField[PascalString](ctx, "text", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer), textSpecified)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'text' field"))
	}
	var text PascalString
	if _text != nil {
		text = *_text
	}

	if closeErr := readBuffer.CloseContext("LocalizedText"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for LocalizedText")
	}

	// Create the instance
	return &_LocalizedText{
		TextSpecified:   textSpecified,
		LocaleSpecified: localeSpecified,
		Locale:          locale,
		Text:            text,
		reservedField0:  reservedField0,
	}, nil
}

func (m *_LocalizedText) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_LocalizedText) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("LocalizedText"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for LocalizedText")
	}

	if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 6)); err != nil {
		return errors.Wrap(err, "Error serializing 'reserved' field number 1")
	}

	// Simple Field (textSpecified)
	textSpecified := bool(m.GetTextSpecified())
	_textSpecifiedErr := /*TODO: migrate me*/ writeBuffer.WriteBit("textSpecified", (textSpecified))
	if _textSpecifiedErr != nil {
		return errors.Wrap(_textSpecifiedErr, "Error serializing 'textSpecified' field")
	}

	// Simple Field (localeSpecified)
	localeSpecified := bool(m.GetLocaleSpecified())
	_localeSpecifiedErr := /*TODO: migrate me*/ writeBuffer.WriteBit("localeSpecified", (localeSpecified))
	if _localeSpecifiedErr != nil {
		return errors.Wrap(_localeSpecifiedErr, "Error serializing 'localeSpecified' field")
	}

	if err := WriteOptionalField[PascalString](ctx, "locale", GetRef(m.GetLocale()), WriteComplex[PascalString](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'locale' field")
	}

	if err := WriteOptionalField[PascalString](ctx, "text", GetRef(m.GetText()), WriteComplex[PascalString](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'text' field")
	}

	if popErr := writeBuffer.PopContext("LocalizedText"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for LocalizedText")
	}
	return nil
}

func (m *_LocalizedText) isLocalizedText() bool {
	return true
}

func (m *_LocalizedText) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
