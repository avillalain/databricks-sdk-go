/*
 * Databricks
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package models

type ServicePrincipalCreateRequest struct {
	ApplicationId string         `json:"applicationId,omitempty"`
	DisplayName   string         `json:"displayName,omitempty"`
	Entitlements  []Entitlements `json:"entitlements,omitempty"`
}
