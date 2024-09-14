package dbmongo_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	dbmongo "github.com/syedvasil/shaffra_skill_assessment_golang/database/mongo"
	"testing"
)

const (
	// Replace with your MongoDB URI for testing
	validURI = "mongodb://localhost:27017" // Ensure this URI is for a test database
	DB       = "testdb"
)

func TestInitDBConnect_Success(t *testing.T) {
	t.Parallel()

	// Initialize the database
	db, err := dbmongo.InitDBConnect(validURI)
	if err != nil {
		assert.Fail(t, "error should be nil")
	}

	require.NotNil(t, db)

	_, _ = db.ListCollectionNames(context.Background(), nil)
}
