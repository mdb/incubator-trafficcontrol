package tc

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

import (
	"testing"
	"time"
)

var jsonTests = []struct {
	time Time
	json string
}{
	{Time{Time: time.Date(9999, 4, 12, 23, 20, 50, 520*1e6, time.UTC)}, `"9999-04-12 23:20:50+00"`},
	{Time{Time: time.Date(1996, 12, 19, 16, 39, 57, 0, time.UTC)}, `"1996-12-19 16:39:57+00"`},
}

// TestJSON tests that we get format tc uses for lastUpdated fields
func TestJSON(t *testing.T) {
	for _, tm := range jsonTests {
		got, err := tm.time.MarshalJSON()
		if err != nil {
			t.Errorf("MarshalJSON error: %+v", err)
		}

		if string(got) != tm.json {
			t.Errorf("expected %s, got %s", tm.json, got)
		}
	}
}
