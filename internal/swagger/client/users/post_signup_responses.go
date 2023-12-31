// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostSignupReader is a Reader for the PostSignup structure.
type PostSignupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSignupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostSignupCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostSignupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostSignupConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /signup] PostSignup", response, response.Code())
	}
}

// NewPostSignupCreated creates a PostSignupCreated with default headers values
func NewPostSignupCreated() *PostSignupCreated {
	return &PostSignupCreated{}
}

/*
PostSignupCreated describes a response with status code 201, with default header values.

User registered successfully
*/
type PostSignupCreated struct {
}

// IsSuccess returns true when this post signup created response has a 2xx status code
func (o *PostSignupCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post signup created response has a 3xx status code
func (o *PostSignupCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post signup created response has a 4xx status code
func (o *PostSignupCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post signup created response has a 5xx status code
func (o *PostSignupCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post signup created response a status code equal to that given
func (o *PostSignupCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the post signup created response
func (o *PostSignupCreated) Code() int {
	return 201
}

func (o *PostSignupCreated) Error() string {
	return fmt.Sprintf("[POST /signup][%d] postSignupCreated ", 201)
}

func (o *PostSignupCreated) String() string {
	return fmt.Sprintf("[POST /signup][%d] postSignupCreated ", 201)
}

func (o *PostSignupCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSignupBadRequest creates a PostSignupBadRequest with default headers values
func NewPostSignupBadRequest() *PostSignupBadRequest {
	return &PostSignupBadRequest{}
}

/*
PostSignupBadRequest describes a response with status code 400, with default header values.

Bad request. Invalid input data
*/
type PostSignupBadRequest struct {
}

// IsSuccess returns true when this post signup bad request response has a 2xx status code
func (o *PostSignupBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post signup bad request response has a 3xx status code
func (o *PostSignupBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post signup bad request response has a 4xx status code
func (o *PostSignupBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post signup bad request response has a 5xx status code
func (o *PostSignupBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post signup bad request response a status code equal to that given
func (o *PostSignupBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post signup bad request response
func (o *PostSignupBadRequest) Code() int {
	return 400
}

func (o *PostSignupBadRequest) Error() string {
	return fmt.Sprintf("[POST /signup][%d] postSignupBadRequest ", 400)
}

func (o *PostSignupBadRequest) String() string {
	return fmt.Sprintf("[POST /signup][%d] postSignupBadRequest ", 400)
}

func (o *PostSignupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSignupConflict creates a PostSignupConflict with default headers values
func NewPostSignupConflict() *PostSignupConflict {
	return &PostSignupConflict{}
}

/*
PostSignupConflict describes a response with status code 409, with default header values.

Conflict. User already exists
*/
type PostSignupConflict struct {
}

// IsSuccess returns true when this post signup conflict response has a 2xx status code
func (o *PostSignupConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post signup conflict response has a 3xx status code
func (o *PostSignupConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post signup conflict response has a 4xx status code
func (o *PostSignupConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this post signup conflict response has a 5xx status code
func (o *PostSignupConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this post signup conflict response a status code equal to that given
func (o *PostSignupConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the post signup conflict response
func (o *PostSignupConflict) Code() int {
	return 409
}

func (o *PostSignupConflict) Error() string {
	return fmt.Sprintf("[POST /signup][%d] postSignupConflict ", 409)
}

func (o *PostSignupConflict) String() string {
	return fmt.Sprintf("[POST /signup][%d] postSignupConflict ", 409)
}

func (o *PostSignupConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
PostSignupBody post signup body
swagger:model PostSignupBody
*/
type PostSignupBody struct {

	// The email address of the new user.
	// Required: true
	// Format: email
	Email *strfmt.Email `json:"email"`

	// The password for the new user.
	// Required: true
	Password *string `json:"password"`

	// The username of the new user.
	// Required: true
	Username *string `json:"username"`
}

// Validate validates this post signup body
func (o *PostSignupBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostSignupBody) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"email", "body", o.Email); err != nil {
		return err
	}

	if err := validate.FormatOf("body"+"."+"email", "body", "email", o.Email.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *PostSignupBody) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"password", "body", o.Password); err != nil {
		return err
	}

	return nil
}

func (o *PostSignupBody) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"username", "body", o.Username); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post signup body based on context it is used
func (o *PostSignupBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostSignupBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostSignupBody) UnmarshalBinary(b []byte) error {
	var res PostSignupBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
