//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

package rest

import (
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/semi-technologies/weaviate/adapters/handlers/rest/operations"
	"github.com/semi-technologies/weaviate/adapters/handlers/rest/operations/schema"
	"github.com/semi-technologies/weaviate/adapters/handlers/rest/state"
	"github.com/semi-technologies/weaviate/adapters/repos/db"
	"github.com/semi-technologies/weaviate/entities/models"
	entitySchema "github.com/semi-technologies/weaviate/entities/schema"
	"github.com/semi-technologies/weaviate/entities/snapshots"
	"github.com/semi-technologies/weaviate/usecases/auth/authorization/errors"
	"github.com/semi-technologies/weaviate/usecases/backup"
	schemaUC "github.com/semi-technologies/weaviate/usecases/schema"
	"github.com/semi-technologies/weaviate/usecases/sharding"
)

func newSnapshotterProvider(db *db.DB) backup.SnapshotterProvider {
	return &snapshotterProvider{db}
}

type snapshotterProvider struct {
	db *db.DB
}

func (sp *snapshotterProvider) Snapshotter(className string) backup.Snapshotter {
	if idx := sp.db.GetIndex(entitySchema.ClassName(className)); idx != nil {
		return idx
	}
	return nil
}

type backupHandlers struct {
	manager *backup.Manager
}

func (s *backupHandlers) createSnapshot(params schema.SchemaObjectsSnapshotsCreateParams,
	principal *models.Principal,
) middleware.Responder {
	meta, err := s.manager.CreateSnapshot(params.HTTPRequest.Context(), principal,
		params.ClassName, params.StorageName, params.ID)
	if err != nil {
		switch err.(type) {
		case errors.Forbidden:
			return schema.NewSchemaObjectsSnapshotsCreateForbidden().
				WithPayload(errPayloadFromSingleErr(err))
		case snapshots.ErrUnprocessable:
			return schema.NewSchemaObjectsSnapshotsCreateUnprocessableEntity().
				WithPayload(errPayloadFromSingleErr(err))
		default:
			return schema.NewSchemaObjectsSnapshotsCreateInternalServerError().
				WithPayload(errPayloadFromSingleErr(err))
		}
	}

	return schema.NewSchemaObjectsSnapshotsCreateOK().WithPayload(meta)
}

func (s *backupHandlers) createSnapshotStatus(params schema.SchemaObjectsSnapshotsCreateStatusParams,
	principal *models.Principal,
) middleware.Responder {
	status, err := s.manager.CreateSnapshotStatus(params.HTTPRequest.Context(), principal,
		params.ClassName, params.StorageName, params.ID)
	if err != nil {
		switch err.(type) {
		case errors.Forbidden:
			return schema.NewSchemaObjectsSnapshotsCreateStatusForbidden().
				WithPayload(errPayloadFromSingleErr(err))
		case snapshots.ErrNotFound:
			return schema.NewSchemaObjectsSnapshotsCreateStatusNotFound().
				WithPayload(errPayloadFromSingleErr(err))
		default:
			return schema.NewSchemaObjectsSnapshotsCreateStatusInternalServerError().
				WithPayload(errPayloadFromSingleErr(err))
		}
	}
	return schema.NewSchemaObjectsSnapshotsCreateStatusOK().WithPayload(status)
}

func (s *backupHandlers) restoreSnapshot(params schema.SchemaObjectsSnapshotsRestoreParams,
	principal *models.Principal,
) middleware.Responder {
	meta, err := s.manager.RestoreSnapshot(params.HTTPRequest.Context(), principal,
		params.ClassName, params.StorageName, params.ID)
	if err != nil {
		switch err.(type) {
		case errors.Forbidden:
			return schema.NewSchemaObjectsSnapshotsRestoreForbidden().
				WithPayload(errPayloadFromSingleErr(err))
		case snapshots.ErrNotFound:
			return schema.NewSchemaObjectsSnapshotsRestoreNotFound().
				WithPayload(errPayloadFromSingleErr(err))
		case snapshots.ErrUnprocessable:
			return schema.NewSchemaObjectsSnapshotsRestoreUnprocessableEntity().
				WithPayload(errPayloadFromSingleErr(err))
		default:
			return schema.NewSchemaObjectsSnapshotsRestoreInternalServerError().
				WithPayload(errPayloadFromSingleErr(err))
		}
	}

	return schema.NewSchemaObjectsSnapshotsRestoreOK().WithPayload(meta)
}

func (s *backupHandlers) restoreSnapshotStatus(params schema.SchemaObjectsSnapshotsRestoreStatusParams,
	principal *models.Principal,
) middleware.Responder {
	status, restoreError, path, err := s.manager.RestoreSnapshotStatus(params.HTTPRequest.Context(), principal, params.ClassName, params.StorageName, params.ID)
	if err != nil {
		switch err.(type) {
		case errors.Forbidden:
			return schema.NewSchemaObjectsSnapshotsRestoreForbidden().
				WithPayload(errPayloadFromSingleErr(err))
		case snapshots.ErrNotFound:
			return schema.NewSchemaObjectsSnapshotsRestoreNotFound().
				WithPayload(errPayloadFromSingleErr(err))
		case snapshots.ErrUnprocessable:
			return schema.NewSchemaObjectsSnapshotsRestoreUnprocessableEntity().
				WithPayload(errPayloadFromSingleErr(err))
		default:
			return schema.NewSchemaObjectsSnapshotsRestoreInternalServerError().
				WithPayload(errPayloadFromSingleErr(err))
		}
	}

	return schema.
		NewSchemaObjectsSnapshotsRestoreStatusOK().
		WithPayload(&models.SnapshotRestoreMeta{
			Status:      &status,
			ClassName:   params.ClassName,
			Error:       restoreError,
			ID:          params.ID,
			Path:        path,
			StorageName: params.StorageName,
		})
}

func setupBackupHandlers(api *operations.WeaviateAPI, schemaManger *schemaUC.Manager, repo *db.DB, appState *state.State) {
	shardingStateFunc := func(className string) *sharding.State {
		return appState.SchemaManager.ShardingState(className)
	}
	snapshotterProvider := newSnapshotterProvider(repo)
	backupManager := backup.NewManager(appState.Logger, appState.Authorizer,
		schemaManger, snapshotterProvider,
		appState.Modules, shardingStateFunc)

	h := &backupHandlers{backupManager}
	api.SchemaSchemaObjectsSnapshotsCreateHandler = schema.
		SchemaObjectsSnapshotsCreateHandlerFunc(h.createSnapshot)
	api.SchemaSchemaObjectsSnapshotsCreateStatusHandler = schema.
		SchemaObjectsSnapshotsCreateStatusHandlerFunc(h.createSnapshotStatus)
	api.SchemaSchemaObjectsSnapshotsRestoreHandler = schema.
		SchemaObjectsSnapshotsRestoreHandlerFunc(h.restoreSnapshot)
	api.SchemaSchemaObjectsSnapshotsRestoreStatusHandler = schema.
		SchemaObjectsSnapshotsRestoreStatusHandlerFunc(h.restoreSnapshotStatus)
}
