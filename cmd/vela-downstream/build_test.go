// SPDX-License-Identifier: Apache-2.0

package main

import (
	"testing"

	"github.com/go-vela/types/constants"
)

func TestDownstream_Build_Validate(t *testing.T) {
	// setup types
	b := &Build{
		Branch: "master",
		Event:  constants.EventPush,
		Status: []string{constants.StatusSuccess},
	}

	// run test
	err := b.Validate()
	if err != nil {
		t.Errorf("Validate returned err: %v", err)
	}
}

func TestDownstream_Build_Validate_NoBranch(t *testing.T) {
	// setup types
	b := &Build{
		Event:  constants.EventPush,
		Status: []string{constants.StatusSuccess},
	}

	// run test
	err := b.Validate()
	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}

func TestDownstream_Build_Validate_NoEvent(t *testing.T) {
	// setup types
	b := &Build{
		Branch: "master",
		Status: []string{constants.StatusSuccess},
	}

	// run test
	err := b.Validate()
	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}

func TestDownstream_Build_Validate_NoStatus(t *testing.T) {
	// setup types
	b := &Build{
		Branch: "master",
		Event:  constants.EventPush,
	}

	// run test
	err := b.Validate()
	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}

func TestDownstream_Build_Validate_InvalidEvent(t *testing.T) {
	// setup types
	b := &Build{
		Branch: "master",
		Event:  "foo",
		Status: []string{constants.StatusSuccess},
	}

	// run test
	err := b.Validate()
	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}

func TestDownstream_Build_Validate_InvalidStatus(t *testing.T) {
	// setup types
	b := &Build{
		Branch: "master",
		Event:  constants.EventPush,
		Status: []string{"foo"},
	}

	// run test
	err := b.Validate()
	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}
