package errors

import (
	"net/http"
)

var httpCode = map[ErrorCode]int{
	UNKNOWN_ERROR:                  http.StatusInternalServerError,
	DATA_INVALID:                   http.StatusBadRequest,
	INVALID_HEADER_X_BUYER_ID:      http.StatusBadRequest,
	INVALID_HEADER_X_USER_ID:       http.StatusBadRequest,
	INVALID_HEADER_X_SELLER_ID:     http.StatusBadRequest,
	INVALID_PAYLOAD_CREATE_ORDER:   http.StatusBadRequest,
	INVALID_PAYLOAD_UPDATE_ORDER:   http.StatusBadRequest,
	INVALID_ORDER:                  http.StatusBadRequest,
	ORDER_NOT_FOUND:                http.StatusNotFound,
	FAILED_RETRIEVE_ORDER:          http.StatusInternalServerError,
	FAILED_CREATE_ORDER:            http.StatusInternalServerError,
	FAILED_UPDATE_ORDER:            http.StatusInternalServerError,
	FAILED_DELETE_ORDER:            http.StatusInternalServerError,
	INVALID_REQUEST_RETRIEVE_ORDER: http.StatusBadRequest,
	INVALID_REQUEST_CREATE_ORDER:   http.StatusBadRequest,
	INVALID_REQUEST_UPDATE_ORDER:   http.StatusBadRequest,
	INVALID_REQUEST_DELETE_ORDER:   http.StatusBadRequest,
	STATUS_PAGE_NOT_FOUND:          http.StatusNotFound,
	INVALID_HEADER_X_PLATFORM:      http.StatusBadRequest,
	INVALID_HEADER_X_API_KEY:       http.StatusBadRequest,
	UNAUTHORIZED:                   http.StatusUnauthorized,
	FAILED_RETRIEVE_DATA:           http.StatusInternalServerError,
}
