/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package org.apache.plc4x.java.scraper.triggeredscraper;

import org.apache.commons.lang3.tuple.Pair;
import org.apache.plc4x.java.PlcDriverManager;
import org.apache.plc4x.java.api.types.PlcResponseCode;
import org.apache.plc4x.java.spi.messages.items.DefaultBooleanFieldItem;
import org.apache.plc4x.java.spi.messages.items.DefaultLongFieldItem;
import org.apache.plc4x.java.mock.MockDevice;
import org.apache.plc4x.java.mock.PlcMockConnection;
import org.apache.plc4x.java.scraper.config.ScraperConfiguration;
import org.apache.plc4x.java.scraper.config.ScraperConfigurationClassicImpl;
import org.apache.plc4x.java.scraper.exception.ScraperException;
import org.apache.plc4x.java.scraper.triggeredscraper.triggerhandler.collector.TriggerCollector;
import org.apache.plc4x.java.scraper.triggeredscraper.triggerhandler.collector.TriggerCollectorImpl;
import org.junit.Before;
import org.junit.Test;
import org.mockito.Mockito;
import org.mockito.invocation.InvocationOnMock;
import org.mockito.stubbing.Answer;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.io.IOException;
import java.util.Random;

import static org.mockito.ArgumentMatchers.any;
import static org.mockito.ArgumentMatchers.eq;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;

/**
 *
 * @author julian
 * Created by julian on 2019-05-08
 */
public class TriggeredScraperImplTest {

    private static final Logger LOGGER = LoggerFactory.getLogger(TriggeredScraperImplTest.class);
    private PlcDriverManager driverManager;
    private MockDevice mockDevice1;
    private MockDevice mockDevice2;

    @Before
    public void setUp() throws Exception {
        driverManager = new PlcDriverManager();
        PlcMockConnection mock1Connection = ((PlcMockConnection) driverManager.getConnection("mock:1"));
        PlcMockConnection mock2Connection = ((PlcMockConnection) driverManager.getConnection("mock:2"));

        // Create Mocks
        mockDevice1 = Mockito.mock(MockDevice.class);
        mockDevice2 = Mockito.mock(MockDevice.class);
        // Assign to Connections
        mock1Connection.setDevice(mockDevice1);
        mock2Connection.setDevice(mockDevice2);
    }

    /**
     * Test is added because we assume some strange behavior.
     */
    @Test
    public void scrapeMultipleTargets() throws ScraperException, IOException, InterruptedException {
        // Prepare the Mocking
        // Scrate Jobs 1 and 2
        when(mockDevice1.read(eq("%DB810:DBB0:USINT"))).thenReturn(Pair.of(PlcResponseCode.OK, new DefaultLongFieldItem(1L)));
        when(mockDevice2.read(eq("%DB810:DBB0:USINT"))).thenReturn(Pair.of(PlcResponseCode.OK, new DefaultLongFieldItem(2L)));
        // Trigger Jobs
        // Trigger var
        Random rand = new Random();
        when(mockDevice1.read(eq("%M0.3:BOOL"))).thenAnswer(new Answer<Object>() {
            @Override
            public Object answer(InvocationOnMock invocationOnMock) throws Throwable {
                boolean trigger = rand.nextBoolean();
                System.out.println(trigger);
                return Pair.of(PlcResponseCode.OK, new DefaultBooleanFieldItem(trigger));
            }
        });
        when(mockDevice2.read(eq("%M0.3:BOOL"))).thenAnswer(new Answer<Object>() {
            @Override
            public Object answer(InvocationOnMock invocationOnMock) throws Throwable {
                boolean trigger = rand.nextBoolean();
                System.out.println("\t\t" + trigger);
                return Pair.of(PlcResponseCode.OK, new DefaultBooleanFieldItem(trigger));
            }
        });
        // Read var
        when(mockDevice1.read(eq("%DB810:DBW0:INT"))).thenReturn(Pair.of(PlcResponseCode.OK, new DefaultLongFieldItem(3L)));
        when(mockDevice2.read(eq("%DB810:DBW0:INT"))).thenReturn(Pair.of(PlcResponseCode.OK, new DefaultLongFieldItem(4L)));

        ScraperConfiguration configuration = ScraperConfiguration.fromFile("src/test/resources/mock-scraper-config.yml", ScraperConfigurationClassicImpl.class);
        TriggerCollector triggerCollector = new TriggerCollectorImpl(driverManager);
        TriggeredScraperImpl scraper = new TriggeredScraperImpl((j, a, m) -> System.out.println(String.format("Results from %s/%s: %s", j, a, m)), driverManager, configuration.getJobs(),triggerCollector,1000);

        scraper.start();

        Thread.sleep(2_000);

        scraper.stop();
    }
}