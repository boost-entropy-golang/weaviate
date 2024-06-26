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

package locks

// ConnectorSchemaLock to be implemented through local (mutex) or distributed
// (etcd, redis, ...) means
//
// The returned functions are the respective unlock functions
type ConnectorSchemaLock interface {
	LockConnector() (func() error, error)
	LockSchema() (func() error, error)
}
