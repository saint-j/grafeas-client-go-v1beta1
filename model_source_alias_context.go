/*
 * proto/v1beta1/grafeas.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// An alias to a repo revision.
type SourceAliasContext struct {
	// The alias kind.
	Kind *AliasContextKind `json:"kind,omitempty"`
	// The alias name.
	Name string `json:"name,omitempty"`
}
