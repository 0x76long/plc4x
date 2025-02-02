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

public class DataSetMetaDataType extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public Integer getExtensionId() {
    return (int) 14525;
  }

  // Properties.
  protected final List<PascalString> namespaces;
  protected final List<StructureDescription> structureDataTypes;
  protected final List<EnumDescription> enumDataTypes;
  protected final List<SimpleTypeDescription> simpleDataTypes;
  protected final PascalString name;
  protected final LocalizedText description;
  protected final List<FieldMetaData> fields;
  protected final GuidValue dataSetClassId;
  protected final ConfigurationVersionDataType configurationVersion;

  public DataSetMetaDataType(
      List<PascalString> namespaces,
      List<StructureDescription> structureDataTypes,
      List<EnumDescription> enumDataTypes,
      List<SimpleTypeDescription> simpleDataTypes,
      PascalString name,
      LocalizedText description,
      List<FieldMetaData> fields,
      GuidValue dataSetClassId,
      ConfigurationVersionDataType configurationVersion) {
    super();
    this.namespaces = namespaces;
    this.structureDataTypes = structureDataTypes;
    this.enumDataTypes = enumDataTypes;
    this.simpleDataTypes = simpleDataTypes;
    this.name = name;
    this.description = description;
    this.fields = fields;
    this.dataSetClassId = dataSetClassId;
    this.configurationVersion = configurationVersion;
  }

  public List<PascalString> getNamespaces() {
    return namespaces;
  }

  public List<StructureDescription> getStructureDataTypes() {
    return structureDataTypes;
  }

  public List<EnumDescription> getEnumDataTypes() {
    return enumDataTypes;
  }

  public List<SimpleTypeDescription> getSimpleDataTypes() {
    return simpleDataTypes;
  }

  public PascalString getName() {
    return name;
  }

  public LocalizedText getDescription() {
    return description;
  }

  public List<FieldMetaData> getFields() {
    return fields;
  }

  public GuidValue getDataSetClassId() {
    return dataSetClassId;
  }

  public ConfigurationVersionDataType getConfigurationVersion() {
    return configurationVersion;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("DataSetMetaDataType");

    // Implicit Field (noOfNamespaces) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int noOfNamespaces = (int) ((((getNamespaces()) == (null)) ? -(1) : COUNT(getNamespaces())));
    writeImplicitField("noOfNamespaces", noOfNamespaces, writeSignedInt(writeBuffer, 32));

    // Array Field (namespaces)
    writeComplexTypeArrayField("namespaces", namespaces, writeBuffer);

    // Implicit Field (noOfStructureDataTypes) (Used for parsing, but its value is not stored as
    // it's implicitly given by the objects content)
    int noOfStructureDataTypes =
        (int) ((((getStructureDataTypes()) == (null)) ? -(1) : COUNT(getStructureDataTypes())));
    writeImplicitField(
        "noOfStructureDataTypes", noOfStructureDataTypes, writeSignedInt(writeBuffer, 32));

    // Array Field (structureDataTypes)
    writeComplexTypeArrayField("structureDataTypes", structureDataTypes, writeBuffer);

    // Implicit Field (noOfEnumDataTypes) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int noOfEnumDataTypes =
        (int) ((((getEnumDataTypes()) == (null)) ? -(1) : COUNT(getEnumDataTypes())));
    writeImplicitField("noOfEnumDataTypes", noOfEnumDataTypes, writeSignedInt(writeBuffer, 32));

    // Array Field (enumDataTypes)
    writeComplexTypeArrayField("enumDataTypes", enumDataTypes, writeBuffer);

    // Implicit Field (noOfSimpleDataTypes) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int noOfSimpleDataTypes =
        (int) ((((getSimpleDataTypes()) == (null)) ? -(1) : COUNT(getSimpleDataTypes())));
    writeImplicitField("noOfSimpleDataTypes", noOfSimpleDataTypes, writeSignedInt(writeBuffer, 32));

    // Array Field (simpleDataTypes)
    writeComplexTypeArrayField("simpleDataTypes", simpleDataTypes, writeBuffer);

    // Simple Field (name)
    writeSimpleField("name", name, writeComplex(writeBuffer));

    // Simple Field (description)
    writeSimpleField("description", description, writeComplex(writeBuffer));

    // Implicit Field (noOfFields) (Used for parsing, but its value is not stored as it's implicitly
    // given by the objects content)
    int noOfFields = (int) ((((getFields()) == (null)) ? -(1) : COUNT(getFields())));
    writeImplicitField("noOfFields", noOfFields, writeSignedInt(writeBuffer, 32));

    // Array Field (fields)
    writeComplexTypeArrayField("fields", fields, writeBuffer);

    // Simple Field (dataSetClassId)
    writeSimpleField("dataSetClassId", dataSetClassId, writeComplex(writeBuffer));

    // Simple Field (configurationVersion)
    writeSimpleField("configurationVersion", configurationVersion, writeComplex(writeBuffer));

    writeBuffer.popContext("DataSetMetaDataType");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    DataSetMetaDataType _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Implicit Field (noOfNamespaces)
    lengthInBits += 32;

    // Array field
    if (namespaces != null) {
      int i = 0;
      for (PascalString element : namespaces) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= namespaces.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    // Implicit Field (noOfStructureDataTypes)
    lengthInBits += 32;

    // Array field
    if (structureDataTypes != null) {
      int i = 0;
      for (StructureDescription element : structureDataTypes) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= structureDataTypes.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    // Implicit Field (noOfEnumDataTypes)
    lengthInBits += 32;

    // Array field
    if (enumDataTypes != null) {
      int i = 0;
      for (EnumDescription element : enumDataTypes) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= enumDataTypes.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    // Implicit Field (noOfSimpleDataTypes)
    lengthInBits += 32;

    // Array field
    if (simpleDataTypes != null) {
      int i = 0;
      for (SimpleTypeDescription element : simpleDataTypes) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= simpleDataTypes.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    // Simple field (name)
    lengthInBits += name.getLengthInBits();

    // Simple field (description)
    lengthInBits += description.getLengthInBits();

    // Implicit Field (noOfFields)
    lengthInBits += 32;

    // Array field
    if (fields != null) {
      int i = 0;
      for (FieldMetaData element : fields) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= fields.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    // Simple field (dataSetClassId)
    lengthInBits += dataSetClassId.getLengthInBits();

    // Simple field (configurationVersion)
    lengthInBits += configurationVersion.getLengthInBits();

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, Integer extensionId) throws ParseException {
    readBuffer.pullContext("DataSetMetaDataType");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    int noOfNamespaces = readImplicitField("noOfNamespaces", readSignedInt(readBuffer, 32));

    List<PascalString> namespaces =
        readCountArrayField(
            "namespaces",
            readComplex(() -> PascalString.staticParse(readBuffer), readBuffer),
            noOfNamespaces);

    int noOfStructureDataTypes =
        readImplicitField("noOfStructureDataTypes", readSignedInt(readBuffer, 32));

    List<StructureDescription> structureDataTypes =
        readCountArrayField(
            "structureDataTypes",
            readComplex(
                () ->
                    (StructureDescription)
                        ExtensionObjectDefinition.staticParse(readBuffer, (int) (15489)),
                readBuffer),
            noOfStructureDataTypes);

    int noOfEnumDataTypes = readImplicitField("noOfEnumDataTypes", readSignedInt(readBuffer, 32));

    List<EnumDescription> enumDataTypes =
        readCountArrayField(
            "enumDataTypes",
            readComplex(
                () ->
                    (EnumDescription)
                        ExtensionObjectDefinition.staticParse(readBuffer, (int) (15490)),
                readBuffer),
            noOfEnumDataTypes);

    int noOfSimpleDataTypes =
        readImplicitField("noOfSimpleDataTypes", readSignedInt(readBuffer, 32));

    List<SimpleTypeDescription> simpleDataTypes =
        readCountArrayField(
            "simpleDataTypes",
            readComplex(
                () ->
                    (SimpleTypeDescription)
                        ExtensionObjectDefinition.staticParse(readBuffer, (int) (15007)),
                readBuffer),
            noOfSimpleDataTypes);

    PascalString name =
        readSimpleField(
            "name", readComplex(() -> PascalString.staticParse(readBuffer), readBuffer));

    LocalizedText description =
        readSimpleField(
            "description", readComplex(() -> LocalizedText.staticParse(readBuffer), readBuffer));

    int noOfFields = readImplicitField("noOfFields", readSignedInt(readBuffer, 32));

    List<FieldMetaData> fields =
        readCountArrayField(
            "fields",
            readComplex(
                () ->
                    (FieldMetaData)
                        ExtensionObjectDefinition.staticParse(readBuffer, (int) (14526)),
                readBuffer),
            noOfFields);

    GuidValue dataSetClassId =
        readSimpleField(
            "dataSetClassId", readComplex(() -> GuidValue.staticParse(readBuffer), readBuffer));

    ConfigurationVersionDataType configurationVersion =
        readSimpleField(
            "configurationVersion",
            readComplex(
                () ->
                    (ConfigurationVersionDataType)
                        ExtensionObjectDefinition.staticParse(readBuffer, (int) (14595)),
                readBuffer));

    readBuffer.closeContext("DataSetMetaDataType");
    // Create the instance
    return new DataSetMetaDataTypeBuilderImpl(
        namespaces,
        structureDataTypes,
        enumDataTypes,
        simpleDataTypes,
        name,
        description,
        fields,
        dataSetClassId,
        configurationVersion);
  }

  public static class DataSetMetaDataTypeBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final List<PascalString> namespaces;
    private final List<StructureDescription> structureDataTypes;
    private final List<EnumDescription> enumDataTypes;
    private final List<SimpleTypeDescription> simpleDataTypes;
    private final PascalString name;
    private final LocalizedText description;
    private final List<FieldMetaData> fields;
    private final GuidValue dataSetClassId;
    private final ConfigurationVersionDataType configurationVersion;

    public DataSetMetaDataTypeBuilderImpl(
        List<PascalString> namespaces,
        List<StructureDescription> structureDataTypes,
        List<EnumDescription> enumDataTypes,
        List<SimpleTypeDescription> simpleDataTypes,
        PascalString name,
        LocalizedText description,
        List<FieldMetaData> fields,
        GuidValue dataSetClassId,
        ConfigurationVersionDataType configurationVersion) {
      this.namespaces = namespaces;
      this.structureDataTypes = structureDataTypes;
      this.enumDataTypes = enumDataTypes;
      this.simpleDataTypes = simpleDataTypes;
      this.name = name;
      this.description = description;
      this.fields = fields;
      this.dataSetClassId = dataSetClassId;
      this.configurationVersion = configurationVersion;
    }

    public DataSetMetaDataType build() {
      DataSetMetaDataType dataSetMetaDataType =
          new DataSetMetaDataType(
              namespaces,
              structureDataTypes,
              enumDataTypes,
              simpleDataTypes,
              name,
              description,
              fields,
              dataSetClassId,
              configurationVersion);
      return dataSetMetaDataType;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof DataSetMetaDataType)) {
      return false;
    }
    DataSetMetaDataType that = (DataSetMetaDataType) o;
    return (getNamespaces() == that.getNamespaces())
        && (getStructureDataTypes() == that.getStructureDataTypes())
        && (getEnumDataTypes() == that.getEnumDataTypes())
        && (getSimpleDataTypes() == that.getSimpleDataTypes())
        && (getName() == that.getName())
        && (getDescription() == that.getDescription())
        && (getFields() == that.getFields())
        && (getDataSetClassId() == that.getDataSetClassId())
        && (getConfigurationVersion() == that.getConfigurationVersion())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getNamespaces(),
        getStructureDataTypes(),
        getEnumDataTypes(),
        getSimpleDataTypes(),
        getName(),
        getDescription(),
        getFields(),
        getDataSetClassId(),
        getConfigurationVersion());
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
