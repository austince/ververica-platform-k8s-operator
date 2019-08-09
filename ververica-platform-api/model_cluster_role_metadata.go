/*
 * AppManager API
 *
 * HTTP REST API to connect to the AppManager
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ververicaplatformapi

import (
	"time"
)

type ClusterRoleMetadata struct {
	Id              string            `json:"id,omitempty"`
	Name            string            `json:"name"`
	CreatedAt       time.Time         `json:"createdAt,omitempty"`
	ModifiedAt      time.Time         `json:"modifiedAt,omitempty"`
	Labels          map[string]string `json:"labels,omitempty"`
	ResourceVersion int32             `json:"resourceVersion,omitempty"`
}
