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
	"github.com/jacobsa/aes/siv"
	aes_testing "github.com/jacobsa/aes/testing"
	. "github.com/jacobsa/oglematchers"
	. "github.com/jacobsa/ogletest"
	"testing"
)

func TestDecrypt(t *testing.T) { RunTests(t) }

////////////////////////////////////////////////////////////////////////
// Helpers
////////////////////////////////////////////////////////////////////////

type DecryptTest struct{}

func init() { RegisterTestSuite(&DecryptTest{}) }

////////////////////////////////////////////////////////////////////////
// Tests
////////////////////////////////////////////////////////////////////////

func (t *DecryptTest) NilKey() {
	key := []byte(nil)
	ciphertext := []byte{}

	_, err := siv.Decrypt(key, ciphertext, nil)
	ExpectThat(err, Error(HasSubstr("-byte")))
}

func (t *DecryptTest) ShortKey() {
	ExpectEq("TODO", "")
}

func (t *DecryptTest) LongKey() {
	ExpectEq("TODO", "")
}

func (t *DecryptTest) TooMuchAssociatedData() {
	ExpectEq("TODO", "")
}

func (t *DecryptTest) JustLittleEnoughAssociatedData() {
	ExpectEq("TODO", "")
}

func (t *DecryptTest) DoesntClobberAssociatedSlice() {
	ExpectEq("TODO", "")
}

func (t *DecryptTest) WrongKey() {
	ExpectEq("TODO", "")
}

func (t *DecryptTest) Rfc5297TestCaseA1() {
	key := aes_testing.FromRfcHex(
		"fffefdfc fbfaf9f8 f7f6f5f4 f3f2f1f0" +
		"f0f1f2f3 f4f5f6f7 f8f9fafb fcfdfeff")

	ciphertext := aes_testing.FromRfcHex(
		"85632d07 c6e8f37f 950acd32 0a2ecc93" +
		"40c02b96 90c4dc04 daef7f6a fe5c")

	associated := [][]byte{
		aes_testing.FromRfcHex(
			"10111213 14151617 18191a1b 1c1d1e1f" +
			"20212223 24252627"),
	}

	expected := aes_testing.FromRfcHex(
		"11223344 55667788 99aabbcc ddee")

	output, err := siv.Decrypt(key, ciphertext, associated)
	AssertEq(nil, err)
	ExpectThat(output, DeepEquals(expected))
}

func (t *DecryptTest) Rfc5297TestCaseA2() {
	key := aes_testing.FromRfcHex(
		"7f7e7d7c 7b7a7978 77767574 73727170" +
		"40414243 44454647 48494a4b 4c4d4e4f")

	ciphertext := aes_testing.FromRfcHex(
		"7bdb6e3b 432667eb 06f4d14b ff2fbd0f" +
		"cb900f2f ddbe4043 26601965 c889bf17" +
		"dba77ceb 094fa663 b7a3f748 ba8af829" +
		"ea64ad54 4a272e9c 485b62a3 fd5c0d")

	associated := [][]byte{
		aes_testing.FromRfcHex(
			"00112233 44556677 8899aabb ccddeeff" +
			"deaddada deaddada ffeeddcc bbaa9988" +
			"77665544 33221100"),
		aes_testing.FromRfcHex(
			"10203040 50607080 90a0"),
		aes_testing.FromRfcHex(
			"09f91102 9d74e35b d84156c5 635688c0"),
	}

	expected := aes_testing.FromRfcHex(
		"74686973 20697320 736f6d65 20706c61" +
		"696e7465 78742074 6f20656e 63727970" +
		"74207573 696e6720 5349562d 414553")

	output, err := siv.Decrypt(key, ciphertext, associated)
	AssertEq(nil, err)
	ExpectThat(output, DeepEquals(expected))
}

func (t *DecryptTest) GeneratedTestCases() {
	ExpectEq("TODO", "")
}