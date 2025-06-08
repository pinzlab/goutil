package mailer

// TemplateData represents the dynamic content passed to an email template.
//
// This generic struct allows flexible rendering of templates by combining
// standard metadata (like recipient, subject, and app info) with custom data.
//
// Type Parameter:
//   - D: The type of custom data passed to the template (e.g., string, struct).
type TemplateData[D any] struct {
	// To is the recipient's email address.
	To string

	// Subject is the subject line of the email.
	Subject string

	// Domain is the base URL or domain used in the template,
	// commonly for images, links, or branding (e.g., "https://example.com").
	Domain string

	// AppName is the name of the application or service sending the email,
	// used for branding and sign-off.
	AppName string

	// Data is the custom, template-specific payload.
	// For example, this could be an OTP string or a struct with user details.
	Data D
}
