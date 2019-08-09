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

type SecretValueMetadata struct {
	Id              string            `json:"id,omitempty"`
	Name            string            `json:"name"`
	Namespace       string            `json:"namespace,omitempty"`
	CreatedAt       time.Time         `json:"createdAt,omitempty"`
	ModifiedAt      time.Time         `json:"modifiedAt,omitempty"`
	ResourceVersion int32             `json:"resourceVersion,omitempty"`
	Labels          map[string]string `json:"labels,omitempty"`
	Annotations     map[string]string `json:"annotations,omitempty"`
}
