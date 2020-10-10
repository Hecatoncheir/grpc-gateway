/*
 * examples/internal/proto/examplepb/generate_unbound_methods.proto
 *
 * Generate Unannotated Methods Echo Service Similar to echo_service.proto but without annotations and without external configuration.  Generate Unannotated Methods Echo Service API consists of a single service which returns a message.
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package generateunboundmethods

// GenerateUnboundMethodsSimpleMessage represents a simple message sent to the unannotated GenerateUnboundMethodsEchoService service.
type ExamplepbGenerateUnboundMethodsSimpleMessage struct {
	// Id represents the message identifier.
	Id string `json:"id,omitempty"`
	Num string `json:"num,omitempty"`
	Duration string `json:"duration,omitempty"`
}