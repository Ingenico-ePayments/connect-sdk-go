package errors

import (
	"net/http"

	"github.com/Ingenico-ePayments/connect-sdk-go/communicator/communication"
	apiErrors "github.com/Ingenico-ePayments/connect-sdk-go/domain/errors"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/payment"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/payout"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/refund"
)

// CreateAPIError is used internally in order to create an API error after an HTTP request is done
func CreateAPIError(statusCode int, responseBody string, errorObject interface{}, context communication.CallContext) (APIError, error) {
	var errorID string
	var errors []apiErrors.APIError

	switch response := errorObject.(type) {
	case *payment.ErrorResponse:
		{
			paymentResult := response.PaymentResult

			if paymentResult != nil {
				return NewDeclinedPaymentError(statusCode, responseBody, response)
			}

			if response.ErrorID != nil {
				errorID = *response.ErrorID
			}
			if response.Errors != nil {
				errors = *response.Errors
			}
			break
		}
	case *payout.ErrorResponse:
		{
			payoutResult := response.PayoutResult

			if payoutResult != nil {
				return NewDeclinedPayoutError(statusCode, responseBody, response)
			}

			if response.ErrorID != nil {
				errorID = *response.ErrorID
			}
			if response.Errors != nil {
				errors = *response.Errors
			}

			break
		}
	case *refund.ErrorResponse:
		{
			refundResult := response.RefundResult

			if refundResult != nil {
				return NewDeclinedRefundError(statusCode, responseBody, response)
			}

			if response.ErrorID != nil {
				errorID = *response.ErrorID
			}
			if response.Errors != nil {
				errors = *response.Errors
			}

			break
		}
	case *apiErrors.ErrorResponse:
		{
			if response.ErrorID != nil {
				errorID = *response.ErrorID
			}
			if response.Errors != nil {
				errors = *response.Errors
			}

			break
		}
	}

	switch statusCode {
	case http.StatusBadRequest:
		{
			return NewValidateError(statusCode, responseBody, errorID, errors)
		}
	case http.StatusForbidden:
		{
			return NewAuthorizationError(statusCode, responseBody, errorID, errors)
		}
	case http.StatusNotFound:
		{
			return NewReferenceError(statusCode, responseBody, errorID, errors)
		}
	case http.StatusConflict:
		{
			if isIdempotenceError(errors, context) {
				idempotenceKey := context.GetIdempotenceKey()
				idempotenceTimeStamp := context.GetIdempotenceRequestTimestamp()

				return NewIdempotenceError(idempotenceKey, idempotenceTimeStamp, statusCode, responseBody, errorID, errors)
			}

			return NewReferenceError(statusCode, responseBody, errorID, errors)
		}
	case http.StatusGone:
		{
			return NewReferenceError(statusCode, responseBody, errorID, errors)
		}
	case http.StatusInternalServerError:
		fallthrough
	case http.StatusBadGateway:
		fallthrough
	case http.StatusServiceUnavailable:
		fallthrough
	default:
		{
			return NewGlobalCollectError(statusCode, responseBody, errorID, errors)
		}
	}
}

func isIdempotenceError(errors []apiErrors.APIError, context communication.CallContext) (ok bool) {
	ok = context != nil && len(context.GetIdempotenceKey()) != 0 && len(errors) == 1 && errors[0].Code != nil && *errors[0].Code == "1409"

	return
}
