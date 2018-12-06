/*
 * proto/v1beta1/grafeas.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// A note that indicates a type of analysis a provider would perform. This note exists in a provider's project. A `Discovery` occurrence is created in a consumer's project at the start of analysis.
type DiscoveryDiscovery struct {
	// Required. Immutable. The kind of analysis that is handled by this discovery.
	AnalysisKind *V1beta1NoteKind `json:"analysis_kind,omitempty"`
}
