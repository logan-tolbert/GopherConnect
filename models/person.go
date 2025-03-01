package models

import (
	"fmt"
	"regexp"
	"time"
)

// Human defines a contract that all "person-like" types must follow
type Human interface {
	GetAge() int
	GetBirthday() string
	FormatPhone() string
}

// Person represents a contact in the system with personal details
type Person struct {
	FirstName string
	LastName  string
	BirthDate time.Time
	Phone     string
	Email     string
}

// GetAge calculates and returns the person's current age base on their birthdate
func (p Person) GetAge() int {

	if p.BirthDate.IsZero() {
		return 0
	}

	now := time.Now()
	age := now.Year() - p.BirthDate.Year()

	if now.YearDay() < p.BirthDate.YearDay() {
		age--
	}

	return age
}

// GetBirthday returns the person's birth date formatted as YYYY-MM-DD.
// If the birthdate is unknown, it returns "Birthdate unknown".
func (p Person) GetBirthday() string {
	if p.BirthDate.IsZero() {
		return "Birthdate unknown"
	}
	return p.BirthDate.Format("2006-01-02")
}

// IsValidPhoneNumber checks if the phone number is in a valid format.
// Supports formats like "1234567890", "(123) 456-7890", and "123-456-7890".
// Returns true if the phone number is valid, false otherwise.

func IsValidPhoneNumber(phoneNumber string) bool {
	phoneRegex := `^\(?\d{3}\)?[-.\s]?\d{3}[-.\s]?\d{4}$`
	re := regexp.MustCompile(phoneRegex)
	return re.MatchString(phoneNumber)
}

// FormatPhone returns the phone number formatted as (XXX) XXX-XXXX if valid.
// If the phone number is missing, it returns "No phone available".
func (p Person) FormatPhone() string {
	if p.Phone == "" {
		return "No phone available"
	}

	if len(p.Phone) == 10 {
		return fmt.Sprintf("(%s) %s-%s", p.Phone[:3], p.Phone[3:6], p.Phone[6:])
	}

	return p.Phone
}

// IsValidEmail checks if the given email address is in a valid format.
// Supports standard email formats such as "user@example.com" and "user.name@domain.com".
// The function ensures that the email contains a local part, an "@" symbol, and a valid domain.
// Returns true if the email is valid, false otherwise.
func IsValidEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

// String provides a human-readable representation of a Person.
func (p Person) String() string {
	return fmt.Sprintf("%s %s, Age: %d, Birthdate: %s, Phone: %s, Email: %s (Valid: %t)",
		p.FirstName, p.LastName, p.GetAge(), p.GetBirthday(), p.FormatPhone(), p.Email, IsValidEmail(p.Email))
}
