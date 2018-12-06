/*
 * proto/v1beta1/grafeas.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Response for listing occurrences.
type V1beta1ListOccurrencesResponse struct {
	// The occurrences requested.
	Occurrences []V1beta1Occurrence `json:"occurrences,omitempty"`
	// The next pagination token in the list response. It should be used as `page_token` for the following request. An empty value means no more results.
	NextPageToken string `json:"next_page_token,omitempty"`
}
