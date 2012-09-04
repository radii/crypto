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

package siv_test

import (
	. "github.com/jacobsa/ogletest"
	"testing"
)

func TestEncrypt(t *testing.T) { RunTests(t) }

////////////////////////////////////////////////////////////////////////
// Helpers
////////////////////////////////////////////////////////////////////////

type EncryptTest struct{}

func init() { RegisterTestSuite(&EncryptTest{}) }

////////////////////////////////////////////////////////////////////////
// Tests
////////////////////////////////////////////////////////////////////////

func (t *EncryptTest) NilKey() {
	ExpectEq("TODO", "")
}

func (t *EncryptTest) ShortKey() {
	ExpectEq("TODO", "")
}

func (t *EncryptTest) LongKey() {
	ExpectEq("TODO", "")
}

func (t *EncryptTest) TooMuchAssociatedData() {
	ExpectEq("TODO", "")
}

func (t *EncryptTest) JustLittleEnoughAssociatedData() {
	ExpectEq("TODO", "")
}

func (t *EncryptTest) OutputIsDeterministic() {
	ExpectEq("TODO", "")

func (t *EncryptTest) Rfc5297TestCaseA1() {
	ExpectEq("TODO", "")
}

func (t *EncryptTest) Rfc5297TestCaseA2() {
	ExpectEq("TODO", "")
}

func (t *EncryptTest) GeneratedTestCases() {
	ExpectEq("TODO", "")
}