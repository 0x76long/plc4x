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
package org.apache.plc4x.java.opcua.readwrite;

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public class FindServersRequest extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public Integer getExtensionId() {
    return (int) 422;
  }

  // Properties.
  protected final RequestHeader requestHeader;
  protected final PascalString endpointUrl;
  protected final List<PascalString> localeIds;
  protected final List<PascalString> serverUris;

  public FindServersRequest(
      RequestHeader requestHeader,
      PascalString endpointUrl,
      List<PascalString> localeIds,
      List<PascalString> serverUris) {
    super();
    this.requestHeader = requestHeader;
    this.endpointUrl = endpointUrl;
    this.localeIds = localeIds;
    this.serverUris = serverUris;
  }

  public RequestHeader getRequestHeader() {
    return requestHeader;
  }

  public PascalString getEndpointUrl() {
    return endpointUrl;
  }

  public List<PascalString> getLocaleIds() {
    return localeIds;
  }

  public List<PascalString> getServerUris() {
    return serverUris;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("FindServersRequest");

    // Simple Field (requestHeader)
    writeSimpleField("requestHeader", requestHeader, writeComplex(writeBuffer));

    // Simple Field (endpointUrl)
    writeSimpleField("endpointUrl", endpointUrl, writeComplex(writeBuffer));

    // Implicit Field (noOfLocaleIds) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int noOfLocaleIds = (int) ((((getLocaleIds()) == (null)) ? -(1) : COUNT(getLocaleIds())));
    writeImplicitField("noOfLocaleIds", noOfLocaleIds, writeSignedInt(writeBuffer, 32));

    // Array Field (localeIds)
    writeComplexTypeArrayField("localeIds", localeIds, writeBuffer);

    // Implicit Field (noOfServerUris) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int noOfServerUris = (int) ((((getServerUris()) == (null)) ? -(1) : COUNT(getServerUris())));
    writeImplicitField("noOfServerUris", noOfServerUris, writeSignedInt(writeBuffer, 32));

    // Array Field (serverUris)
    writeComplexTypeArrayField("serverUris", serverUris, writeBuffer);

    writeBuffer.popContext("FindServersRequest");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    FindServersRequest _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (requestHeader)
    lengthInBits += requestHeader.getLengthInBits();

    // Simple field (endpointUrl)
    lengthInBits += endpointUrl.getLengthInBits();

    // Implicit Field (noOfLocaleIds)
    lengthInBits += 32;

    // Array field
    if (localeIds != null) {
      int i = 0;
      for (PascalString element : localeIds) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= localeIds.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    // Implicit Field (noOfServerUris)
    lengthInBits += 32;

    // Array field
    if (serverUris != null) {
      int i = 0;
      for (PascalString element : serverUris) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= serverUris.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, Integer extensionId) throws ParseException {
    readBuffer.pullContext("FindServersRequest");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    RequestHeader requestHeader =
        readSimpleField(
            "requestHeader",
            readComplex(
                () ->
                    (RequestHeader) ExtensionObjectDefinition.staticParse(readBuffer, (int) (391)),
                readBuffer));

    PascalString endpointUrl =
        readSimpleField(
            "endpointUrl", readComplex(() -> PascalString.staticParse(readBuffer), readBuffer));

    int noOfLocaleIds = readImplicitField("noOfLocaleIds", readSignedInt(readBuffer, 32));

    List<PascalString> localeIds =
        readCountArrayField(
            "localeIds",
            readComplex(() -> PascalString.staticParse(readBuffer), readBuffer),
            noOfLocaleIds);

    int noOfServerUris = readImplicitField("noOfServerUris", readSignedInt(readBuffer, 32));

    List<PascalString> serverUris =
        readCountArrayField(
            "serverUris",
            readComplex(() -> PascalString.staticParse(readBuffer), readBuffer),
            noOfServerUris);

    readBuffer.closeContext("FindServersRequest");
    // Create the instance
    return new FindServersRequestBuilderImpl(requestHeader, endpointUrl, localeIds, serverUris);
  }

  public static class FindServersRequestBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final RequestHeader requestHeader;
    private final PascalString endpointUrl;
    private final List<PascalString> localeIds;
    private final List<PascalString> serverUris;

    public FindServersRequestBuilderImpl(
        RequestHeader requestHeader,
        PascalString endpointUrl,
        List<PascalString> localeIds,
        List<PascalString> serverUris) {
      this.requestHeader = requestHeader;
      this.endpointUrl = endpointUrl;
      this.localeIds = localeIds;
      this.serverUris = serverUris;
    }

    public FindServersRequest build() {
      FindServersRequest findServersRequest =
          new FindServersRequest(requestHeader, endpointUrl, localeIds, serverUris);
      return findServersRequest;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof FindServersRequest)) {
      return false;
    }
    FindServersRequest that = (FindServersRequest) o;
    return (getRequestHeader() == that.getRequestHeader())
        && (getEndpointUrl() == that.getEndpointUrl())
        && (getLocaleIds() == that.getLocaleIds())
        && (getServerUris() == that.getServerUris())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(), getRequestHeader(), getEndpointUrl(), getLocaleIds(), getServerUris());
  }

  @Override
  public String toString() {
    WriteBufferBoxBased writeBufferBoxBased = new WriteBufferBoxBased(true, true);
    try {
      writeBufferBoxBased.writeSerializable(this);
    } catch (SerializationException e) {
      throw new RuntimeException(e);
    }
    return "\n" + writeBufferBoxBased.getBox().toString() + "\n";
  }
}
