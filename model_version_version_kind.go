/*
 * proto/v1beta1/grafeas.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger
// VersionVersionKind : Whether this is an ordinary package version or a sentinel MIN/MAX version.   - VERSION_KIND_UNSPECIFIED: Unknown.  - NORMAL: A standard package version.  - MINIMUM: A special version representing negative infinity.  - MAXIMUM: A special version representing positive infinity.
type VersionVersionKind string

// List of VersionVersionKind
const (
	VERSION_KIND_UNSPECIFIED VersionVersionKind = "VERSION_KIND_UNSPECIFIED"
	NORMAL VersionVersionKind = "NORMAL"
	MINIMUM VersionVersionKind = "MINIMUM"
	MAXIMUM VersionVersionKind = "MAXIMUM"
)
