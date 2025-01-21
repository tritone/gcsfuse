package local_file_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

// //////////////////////////////////////////////////////////////////////
// Boilerplate
// //////////////////////////////////////////////////////////////////////

type localFileTestSuite struct {
	commonLocalFileTestSuite
}

type localFileWithStreaminingWritesTestSuite struct {
	commonLocalFileTestSuite
}

////////////////////////////////////////////////////////////////////////
// Tests
////////////////////////////////////////////////////////////////////////

func TestCurrentTestSuite(t *testing.T) {
	suite.Run(t, currentTestSuite)
}

func TestCommonLocalFileTestSuite(t *testing.T) {
	suite.Run(t, &commonLocalFileTestSuite{})
}
