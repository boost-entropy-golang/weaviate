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

// Package rest with all rest API functions.
package rest

import (
	"crypto/tls"

	"github.com/go-openapi/swag"

	"github.com/semi-technologies/weaviate/adapters/handlers/rest/operations"
	"github.com/semi-technologies/weaviate/usecases/config"
)

var connectorOptionGroup *swag.CommandLineOptionsGroup

// configureAPI -> see configure_api.go

// configureServer -> see configure_server.go

func configureFlags(api *operations.WeaviateAPI) {
	connectorOptionGroup = config.GetConfigOptionGroup()

	api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{
		*connectorOptionGroup,
	}
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}
