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

// BrokerTransportQualityOfService is an enum
type BrokerTransportQualityOfService uint32

type IBrokerTransportQualityOfService interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	BrokerTransportQualityOfService_brokerTransportQualityOfServiceNotSpecified BrokerTransportQualityOfService = 0
	BrokerTransportQualityOfService_brokerTransportQualityOfServiceBestEffort   BrokerTransportQualityOfService = 1
	BrokerTransportQualityOfService_brokerTransportQualityOfServiceAtLeastOnce  BrokerTransportQualityOfService = 2
	BrokerTransportQualityOfService_brokerTransportQualityOfServiceAtMostOnce   BrokerTransportQualityOfService = 3
	BrokerTransportQualityOfService_brokerTransportQualityOfServiceExactlyOnce  BrokerTransportQualityOfService = 4
)

var BrokerTransportQualityOfServiceValues []BrokerTransportQualityOfService

func init() {
	_ = errors.New
	BrokerTransportQualityOfServiceValues = []BrokerTransportQualityOfService{
		BrokerTransportQualityOfService_brokerTransportQualityOfServiceNotSpecified,
		BrokerTransportQualityOfService_brokerTransportQualityOfServiceBestEffort,
		BrokerTransportQualityOfService_brokerTransportQualityOfServiceAtLeastOnce,
		BrokerTransportQualityOfService_brokerTransportQualityOfServiceAtMostOnce,
		BrokerTransportQualityOfService_brokerTransportQualityOfServiceExactlyOnce,
	}
}

func BrokerTransportQualityOfServiceByValue(value uint32) (enum BrokerTransportQualityOfService, ok bool) {
	switch value {
	case 0:
		return BrokerTransportQualityOfService_brokerTransportQualityOfServiceNotSpecified, true
	case 1:
		return BrokerTransportQualityOfService_brokerTransportQualityOfServiceBestEffort, true
	case 2:
		return BrokerTransportQualityOfService_brokerTransportQualityOfServiceAtLeastOnce, true
	case 3:
		return BrokerTransportQualityOfService_brokerTransportQualityOfServiceAtMostOnce, true
	case 4:
		return BrokerTransportQualityOfService_brokerTransportQualityOfServiceExactlyOnce, true
	}
	return 0, false
}

func BrokerTransportQualityOfServiceByName(value string) (enum BrokerTransportQualityOfService, ok bool) {
	switch value {
	case "brokerTransportQualityOfServiceNotSpecified":
		return BrokerTransportQualityOfService_brokerTransportQualityOfServiceNotSpecified, true
	case "brokerTransportQualityOfServiceBestEffort":
		return BrokerTransportQualityOfService_brokerTransportQualityOfServiceBestEffort, true
	case "brokerTransportQualityOfServiceAtLeastOnce":
		return BrokerTransportQualityOfService_brokerTransportQualityOfServiceAtLeastOnce, true
	case "brokerTransportQualityOfServiceAtMostOnce":
		return BrokerTransportQualityOfService_brokerTransportQualityOfServiceAtMostOnce, true
	case "brokerTransportQualityOfServiceExactlyOnce":
		return BrokerTransportQualityOfService_brokerTransportQualityOfServiceExactlyOnce, true
	}
	return 0, false
}

func BrokerTransportQualityOfServiceKnows(value uint32) bool {
	for _, typeValue := range BrokerTransportQualityOfServiceValues {
		if uint32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBrokerTransportQualityOfService(structType any) BrokerTransportQualityOfService {
	castFunc := func(typ any) BrokerTransportQualityOfService {
		if sBrokerTransportQualityOfService, ok := typ.(BrokerTransportQualityOfService); ok {
			return sBrokerTransportQualityOfService
		}
		return 0
	}
	return castFunc(structType)
}

func (m BrokerTransportQualityOfService) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m BrokerTransportQualityOfService) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BrokerTransportQualityOfServiceParse(ctx context.Context, theBytes []byte) (BrokerTransportQualityOfService, error) {
	return BrokerTransportQualityOfServiceParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BrokerTransportQualityOfServiceParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BrokerTransportQualityOfService, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint32("BrokerTransportQualityOfService", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BrokerTransportQualityOfService")
	}
	if enum, ok := BrokerTransportQualityOfServiceByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for BrokerTransportQualityOfService")
		return BrokerTransportQualityOfService(val), nil
	} else {
		return enum, nil
	}
}

func (e BrokerTransportQualityOfService) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BrokerTransportQualityOfService) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteUint32("BrokerTransportQualityOfService", 32, uint32(uint32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BrokerTransportQualityOfService) PLC4XEnumName() string {
	switch e {
	case BrokerTransportQualityOfService_brokerTransportQualityOfServiceNotSpecified:
		return "brokerTransportQualityOfServiceNotSpecified"
	case BrokerTransportQualityOfService_brokerTransportQualityOfServiceBestEffort:
		return "brokerTransportQualityOfServiceBestEffort"
	case BrokerTransportQualityOfService_brokerTransportQualityOfServiceAtLeastOnce:
		return "brokerTransportQualityOfServiceAtLeastOnce"
	case BrokerTransportQualityOfService_brokerTransportQualityOfServiceAtMostOnce:
		return "brokerTransportQualityOfServiceAtMostOnce"
	case BrokerTransportQualityOfService_brokerTransportQualityOfServiceExactlyOnce:
		return "brokerTransportQualityOfServiceExactlyOnce"
	}
	return fmt.Sprintf("Unknown(%v)", uint32(e))
}

func (e BrokerTransportQualityOfService) String() string {
	return e.PLC4XEnumName()
}
