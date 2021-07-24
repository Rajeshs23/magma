// Code generated by go-swagger; DO NOT EDIT.

package lte_gateways

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// PutLTENetworkIDGatewaysGatewayIDDeviceReader is a Reader for the PutLTENetworkIDGatewaysGatewayIDDevice structure.
type PutLTENetworkIDGatewaysGatewayIDDeviceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutLTENetworkIDGatewaysGatewayIDDeviceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutLTENetworkIDGatewaysGatewayIDDeviceNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutLTENetworkIDGatewaysGatewayIDDeviceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutLTENetworkIDGatewaysGatewayIDDeviceNoContent creates a PutLTENetworkIDGatewaysGatewayIDDeviceNoContent with default headers values
func NewPutLTENetworkIDGatewaysGatewayIDDeviceNoContent() *PutLTENetworkIDGatewaysGatewayIDDeviceNoContent {
	return &PutLTENetworkIDGatewaysGatewayIDDeviceNoContent{}
}

/*PutLTENetworkIDGatewaysGatewayIDDeviceNoContent handles this case with default header values.

Success
*/
type PutLTENetworkIDGatewaysGatewayIDDeviceNoContent struct {
}

func (o *PutLTENetworkIDGatewaysGatewayIDDeviceNoContent) Error() string {
	return fmt.Sprintf("[PUT /lte/{network_id}/gateways/{gateway_id}/device][%d] putLteNetworkIdGatewaysGatewayIdDeviceNoContent ", 204)
}

func (o *PutLTENetworkIDGatewaysGatewayIDDeviceNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutLTENetworkIDGatewaysGatewayIDDeviceDefault creates a PutLTENetworkIDGatewaysGatewayIDDeviceDefault with default headers values
func NewPutLTENetworkIDGatewaysGatewayIDDeviceDefault(code int) *PutLTENetworkIDGatewaysGatewayIDDeviceDefault {
	return &PutLTENetworkIDGatewaysGatewayIDDeviceDefault{
		_statusCode: code,
	}
}

/*PutLTENetworkIDGatewaysGatewayIDDeviceDefault handles this case with default header values.

Unexpected Error
*/
type PutLTENetworkIDGatewaysGatewayIDDeviceDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put LTE network ID gateways gateway ID device default response
func (o *PutLTENetworkIDGatewaysGatewayIDDeviceDefault) Code() int {
	return o._statusCode
}

func (o *PutLTENetworkIDGatewaysGatewayIDDeviceDefault) Error() string {
	return fmt.Sprintf("[PUT /lte/{network_id}/gateways/{gateway_id}/device][%d] PutLTENetworkIDGatewaysGatewayIDDevice default  %+v", o._statusCode, o.Payload)
}

func (o *PutLTENetworkIDGatewaysGatewayIDDeviceDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutLTENetworkIDGatewaysGatewayIDDeviceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}