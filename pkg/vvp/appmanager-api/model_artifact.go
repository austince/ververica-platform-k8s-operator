/*
 * Application Manager API
 *
 * Application Manager APIs to control Apache Flink jobs
 *
 * API version: 2.1.0
 * Contact: platform@ververica.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package appmanagerapi

type Artifact struct {
	Kind                 string `json:"kind"`
	JarUri               string `json:"jarUri"`
	MainArgs             string `json:"mainArgs,omitempty"`
	EntryClass           string `json:"entryClass,omitempty"`
	FlinkVersion         string `json:"flinkVersion,omitempty"`
	FlinkImageRegistry   string `json:"flinkImageRegistry,omitempty"`
	FlinkImageRepository string `json:"flinkImageRepository,omitempty"`
	FlinkImageTag        string `json:"flinkImageTag,omitempty"`
}