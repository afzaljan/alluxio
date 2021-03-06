/*
 * The Alluxio Open Foundation licenses this work under the Apache License, version 2.0
 * (the "License"). You may not use this work except in compliance with the License, which is
 * available at www.apache.org/licenses/LICENSE-2.0
 *
 * This software is distributed on an "AS IS" basis, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied, as more fully set forth in the License.
 *
 * See the NOTICE file distributed with this work for information regarding copyright ownership.
 */

package alluxio.retry;

import static org.junit.Assert.assertEquals;

import org.junit.Test;

/**
 * Tests for the {@link CountingRetry} class.
 */
public final class CountingRetryTest {

  /**
   * Tests that the provided number of retries is equal the actual number of retries.
   */
  @Test
  public void testNumRetries() {
    int numTries = 10;
    CountingRetry countingRetry = new CountingRetry(numTries);
    assertEquals(0, countingRetry.getAttemptCount());
    int retryAttempts = 0;
    while (countingRetry.attempt()) {
      retryAttempts++;
    }
    assertEquals(numTries + 1, retryAttempts);
    assertEquals(numTries + 1, countingRetry.getAttemptCount());
  }
}
