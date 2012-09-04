// Copyright 2012 Aaron Jacobs. All Rights Reserved.
// Author: aaronjjacobs@gmail.com (Aaron Jacobs)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package siv

import (
	. "github.com/jacobsa/oglematchers"
	. "github.com/jacobsa/ogletest"
	"testing"
)

func TestDbl(t *testing.T) { RunTests(t) }

////////////////////////////////////////////////////////////////////////
// Helpers
////////////////////////////////////////////////////////////////////////

type DblTest struct{}

func init() { RegisterTestSuite(&DblTest{}) }

////////////////////////////////////////////////////////////////////////
// Tests
////////////////////////////////////////////////////////////////////////

func (t *DblTest) NilBuffer() {
	f := func() { dbl(nil) }
	ExpectThat(f, Panics(HasSubstr("16-byte")))
}

func (t *DblTest) ShortBuffer() {
	f := func() { dbl(make([]byte, 15)) }
	ExpectThat(f, Panics(HasSubstr("16-byte")))
}

func (t *DblTest) LongBuffer() {
	f := func() { dbl(make([]byte, 17)) }
	ExpectThat(f, Panics(HasSubstr("16-byte")))
}

func (t *DblTest) RfcTestCases() {
	ExpectEq("TODO", "")
}

func (t *DblTest) GeneratedTestCases() {
	ExpectEq("TODO", "")
}
