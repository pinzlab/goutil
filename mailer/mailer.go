package mailer

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"path/filepath"
)

// Mailer holds SMTP configuration and the directory for email templates.
type Mailer struct {
	SMTPServer  string // SMTP server address (e.g., "smtp.example.com")
	SMTPPort    string // SMTP server port (e.g., "587")
	Username    string // SMTP username (typically an email address)
	Password    string // SMTP password
	From        string // Email address used as the sender
	TemplateDir string // Directory where HTML email templates are stored
}

// NewMailer creates and returns a new Mailer instance configured with the
// specified SMTP server details and template directory.
//
// Parameters:
// - templateDir: Path to the directory containing HTML email templates.
// - server: SMTP server address.
// - port: SMTP server port (as a string, e.g., "587").
// - username: SMTP login username.
// - password: SMTP login password.
// - from: Sender's email address to appear in sent messages.
//
// Example:
//
//	mailer := NewMailer("./web/dist/email", "smtp.example.com", "587", "user@example.com", "password", "noreply@example.com")
func NewMailer(templateDir, server, port, username, password, from string) *Mailer {
	return &Mailer{
		SMTPServer:  server,
		SMTPPort:    port,
		Username:    username,
		Password:    password,
		From:        from,
		TemplateDir: templateDir,
	}
}

// SendTemplate loads and renders an HTML email template with dynamic data,
// then sends it to the specified recipient using SMTP.
//
// Parameters:
// - templateName: Name of the template file in the TemplateDir (e.g., "welcome.html").
// - to: Recipient's email address.
// - subject: Subject line of the email.
// - data: Arbitrary data (map or struct) to inject into the template.
//
// This method uses the Mailer's configured TemplateDir and SMTP credentials
// to build and send the message as HTML.
//
// Example:
//
//	data := map[string]interface{}{
//		"Name": "Alice",
//		"Link": "https://example.com/activate",
//	}
//	err := mailer.SendTemplate("welcome.html", "alice@example.com", "Welcome!", data)
//
// Returns an error if the template fails to load, render, or if the email cannot be sent.
func (m *Mailer) SendTemplate(templateName, to, subject string, data interface{}) error {
	// Build the full path to the template file
	tmplPath := filepath.Join(m.TemplateDir, templateName)
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return err
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		return err
	}

	// Construct the email with proper headers
	msg := "MIME-Version: 1.0\r\nContent-Type: text/html; charset=UTF-8\r\n" +
		fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n", m.From, to, subject) + body.String()

	// Set up SMTP authentication
	auth := smtp.PlainAuth("", m.Username, m.Password, m.SMTPServer)
	addr := fmt.Sprintf("%s:%s", m.SMTPServer, m.SMTPPort)

	// Send the email
	if err := smtp.SendMail(addr, auth, m.From, []string{to}, []byte(msg)); err != nil {
		return err
	}

	return nil
}
