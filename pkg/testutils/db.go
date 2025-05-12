package testutils

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

func MustInsertPipelineRun(t *testing.T, db *sqlx.DB) (runID int64) {
	require.NoError(t, db.Get(&runID, `INSERT INTO pipeline_runs (state,pipeline_spec_id,pruning_key,created_at) VALUES ($1, 0, 0, NOW()) RETURNING id`, "running"))
	return runID
}

func MustInsertUnfinishedPipelineTaskRun(t *testing.T, db *sqlx.DB, pipelineRunID int64) (trID uuid.UUID) {
	/* #nosec G404 */
	require.NoError(t, db.Get(&trID, `INSERT INTO pipeline_task_runs (dot_id, pipeline_run_id, id, type, created_at) VALUES ($1,$2,$3, '', NOW()) RETURNING id`, strconv.Itoa(rand.Int()), pipelineRunID, uuid.New()))
	return trID
}
