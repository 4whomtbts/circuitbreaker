package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var storageFileDir = "./testset"
var fileName = "ccb"

func removeFile() {
	os.Remove(storageFileDir + "/" + fileName)
}

func TestIsMetricHostAlreadyBraked_WHEN_file_not_contains_host_THEN_False(t *testing.T) {
	s := NewStorage(storageFileDir, fileName)
	result := s.isMetricHostAlreadyBraked("localhost")
	assert.False(t, result)
	removeFile()
}

func TestIsMetricHostAlreadyBraked_WHEN_file_contains_host_THEN_True(t *testing.T) {
	s := NewStorage(storageFileDir, fileName)

	result := s.isMetricHostAlreadyBraked("localhost")
	assert.True(t, result)
	removeFile()
}

