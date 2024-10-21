package matrix

import (
	"context"
	"database/sql"
	"errors"

	"git.janky.solutions/finn/matrix-meshtastic-bridge-go/db"
	"maunium.net/go/mautrix/id"
)

type dbSyncStore struct{}

func (dbSyncStore) SaveFilterID(ctx context.Context, userID id.UserID, filterID string) error {
	queries, dbconn, err := db.Get()
	if err != nil {
		return err
	}
	defer dbconn.Close()

	err = queries.MatrixSaveFilterID(ctx, db.MatrixSaveFilterIDParams{
		UserID:   userID.String(),
		FilterID: filterID,
	})
	if err != nil {
		return err
	}

	return nil
}

func (dbSyncStore) LoadFilterID(ctx context.Context, userID id.UserID) (string, error) {
	queries, dbconn, err := db.Get()
	if err != nil {
		return "", err
	}
	defer dbconn.Close()

	filterID, err := queries.MatrixLoadFilterID(ctx, userID.String())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", nil
		}
		return "", err
	}

	return filterID, nil
}

func (dbSyncStore) SaveNextBatch(ctx context.Context, userID id.UserID, nextBatchToken string) error {
	queries, dbconn, err := db.Get()
	if err != nil {
		return err
	}
	defer dbconn.Close()

	err = queries.MatrixSaveNextBatch(ctx, db.MatrixSaveNextBatchParams{
		UserID: userID.String(),
		Token:  nextBatchToken,
	})
	if err != nil {
		return err
	}

	return nil
}

func (dbSyncStore) LoadNextBatch(ctx context.Context, userID id.UserID) (string, error) {
	queries, dbconn, err := db.Get()
	if err != nil {
		return "", err
	}
	defer dbconn.Close()

	token, err := queries.MatrixLoadNextBatch(ctx, userID.String())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", nil
		}
		return "", err
	}

	return token, nil
}
