// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostLoginReader is a Reader for the PostLogin structure.
type PostLoginReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLoginReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLoginOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPostLoginUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /login] PostLogin", response, response.Code())
	}
}

// NewPostLoginOK creates a PostLoginOK with default headers values
func NewPostLoginOK() *PostLoginOK {
	return &PostLoginOK{}
}

/*
PostLoginOK describes a response with status code 200, with default header values.

User logged in successfully
*/
type PostLoginOK struct {
	Payload *PostLoginOKBody
}

// IsSuccess returns true when this post login o k response has a 2xx status code
func (o *PostLoginOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post login o k response has a 3xx status code
func (o *PostLoginOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post login o k response has a 4xx status code
func (o *PostLoginOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post login o k response has a 5xx status code
func (o *PostLoginOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post login o k response a status code equal to that given
func (o *PostLoginOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post login o k response
func (o *PostLoginOK) Code() int {
	return 200
}

func (o *PostLoginOK) Error() string {
	return fmt.Sprintf("[POST /login][%d] postLoginOK  %+v", 200, o.Payload)
}

func (o *PostLoginOK) String() string {
	return fmt.Sprintf("[POST /login][%d] postLoginOK  %+v", 200, o.Payload)
}

func (o *PostLoginOK) GetPayload() *PostLoginOKBody {
	return o.Payload
}

func (o *PostLoginOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostLoginOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLoginUnauthorized creates a PostLoginUnauthorized with default headers values
func NewPostLoginUnauthorized() *PostLoginUnauthorized {
	return &PostLoginUnauthorized{}
}

/*
PostLoginUnauthorized describes a response with status code 401, with default header values.

Unauthorized. Invalid credentials
*/
type PostLoginUnauthorized struct {
}

// IsSuccess returns true when this post login unauthorized response has a 2xx status code
func (o *PostLoginUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post login unauthorized response has a 3xx status code
func (o *PostLoginUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post login unauthorized response has a 4xx status code
func (o *PostLoginUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post login unauthorized response has a 5xx status code
func (o *PostLoginUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post login unauthorized response a status code equal to that given
func (o *PostLoginUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the post login unauthorized response
func (o *PostLoginUnauthorized) Code() int {
	return 401
}

func (o *PostLoginUnauthorized) Error() string {
	return fmt.Sprintf("[POST /login][%d] postLoginUnauthorized ", 401)
}

func (o *PostLoginUnauthorized) String() string {
	return fmt.Sprintf("[POST /login][%d] postLoginUnauthorized ", 401)
}

func (o *PostLoginUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
PostLoginBody post login body
swagger:model PostLoginBody
*/
type PostLoginBody struct {

	// The email address of the user.
	// Required: true
	// Format: email
	Email *strfmt.Email `json:"email"`

	// The user's password.
	// Required: true
	Password *string `json:"password"`
}

// Validate validates this post login body
func (o *PostLoginBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostLoginBody) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"email", "body", o.Email); err != nil {
		return err
	}

	if err := validate.FormatOf("body"+"."+"email", "body", "email", o.Email.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *PostLoginBody) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"password", "body", o.Password); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post login body based on context it is used
func (o *PostLoginBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostLoginBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostLoginBody) UnmarshalBinary(b []byte) error {
	var res PostLoginBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
PostLoginOKBody post login o k body
swagger:model PostLoginOKBody
*/
type PostLoginOKBody struct {

	// Authentication token for the user's session
	Token string `json:"token,omitempty"`
}

// Validate validates this post login o k body
func (o *PostLoginOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post login o k body based on context it is used
func (o *PostLoginOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostLoginOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostLoginOKBody) UnmarshalBinary(b []byte) error {
	var res PostLoginOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}