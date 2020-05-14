/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package statsd_test

import (
	"testing"
)

func TestStatsd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Statsd Suite")
}
