package errors

const (
	UNKNOWN_ERROR                  ErrorCode = 0
	DATA_INVALID                   ErrorCode = 4_001_00_00001
	INVALID_HEADER_X_BUYER_ID      ErrorCode = 4_001_00_00002
	INVALID_HEADER_X_USER_ID       ErrorCode = 4_001_00_00003
	INVALID_HEADER_X_SELLER_ID     ErrorCode = 4_001_00_00004
	INVALID_PAYLOAD_CREATE_ORDER   ErrorCode = 4_001_00_00005
	INVALID_PAYLOAD_UPDATE_ORDER   ErrorCode = 4_001_00_00006
	INVALID_ORDER                  ErrorCode = 4_001_00_00007
	ORDER_NOT_FOUND                ErrorCode = 4_001_00_00008
	FAILED_RETRIEVE_ORDER          ErrorCode = 4_001_00_00009
	FAILED_CREATE_ORDER            ErrorCode = 4_001_00_00010
	FAILED_UPDATE_ORDER            ErrorCode = 4_001_00_00011
	FAILED_DELETE_ORDER            ErrorCode = 4_001_00_00012
	INVALID_REQUEST_RETRIEVE_ORDER ErrorCode = 4_001_00_00013
	INVALID_REQUEST_CREATE_ORDER   ErrorCode = 4_001_00_00014
	INVALID_REQUEST_UPDATE_ORDER   ErrorCode = 4_001_00_00015
	INVALID_REQUEST_DELETE_ORDER   ErrorCode = 4_001_00_00016
	INVALID_REQUEST_RETRIEVE_STORE ErrorCode = 4_001_00_00017
	FAILED_RETRIEVE_DATA           ErrorCode = 4_001_00_00018
	STATUS_PAGE_NOT_FOUND          ErrorCode = 4_001_00_00019
	INVALID_HEADER_X_PLATFORM      ErrorCode = 4_001_00_00020
	INVALID_HEADER_X_API_KEY       ErrorCode = 4_001_00_00021
	UNAUTHORIZED                   ErrorCode = 4_001_00_00022
	FAILED_CREATE_DATA             ErrorCode = 4_001_00_00023
)

var errorCodes = map[ErrorCode]*CommonError{
	UNKNOWN_ERROR: {
		ClientMessage: "Unknown error.",
		SystemMessage: "Unknown error.",
		ErrorCode:     UNKNOWN_ERROR,
	},
	DATA_INVALID: {
		ClientMessage: "Invalid Data Request",
		SystemMessage: "Some of query params has invalid value.",
		ErrorCode:     DATA_INVALID,
	},
	INVALID_HEADER_X_BUYER_ID: {
		ClientMessage: "Invalid buyer.",
		SystemMessage: "Invalid value of header X-Buyer-ID.",
		ErrorCode:     INVALID_HEADER_X_BUYER_ID,
	},
	INVALID_HEADER_X_USER_ID: {
		ClientMessage: "Invalid user.",
		SystemMessage: "Invalid value of header X-User-ID.",
		ErrorCode:     INVALID_HEADER_X_USER_ID,
	},
	INVALID_HEADER_X_SELLER_ID: {
		ClientMessage: "Invalid seller.",
		SystemMessage: "Invalid value of header X-Seller-ID.",
		ErrorCode:     INVALID_HEADER_X_SELLER_ID,
	},
	INVALID_PAYLOAD_CREATE_ORDER: {
		ClientMessage: "Failed to create order.",
		SystemMessage: "Request payload for create order has an invalid form.",
		ErrorCode:     INVALID_PAYLOAD_CREATE_ORDER,
	},
	INVALID_PAYLOAD_UPDATE_ORDER: {
		ClientMessage: "Failed to update order.",
		SystemMessage: "Request payload for update order has an invalid form.",
		ErrorCode:     INVALID_PAYLOAD_UPDATE_ORDER,
	},
	INVALID_ORDER: {
		ClientMessage: "Invalid order.",
		SystemMessage: "Invalid request order.",
		ErrorCode:     INVALID_ORDER,
	},
	ORDER_NOT_FOUND: {
		ClientMessage: "Invalid order.",
		SystemMessage: "Order not found.",
		ErrorCode:     ORDER_NOT_FOUND,
	},
	FAILED_RETRIEVE_ORDER: {
		ClientMessage: "Failed to retrieve order.",
		SystemMessage: "Something wrong happened while retrieve order.",
		ErrorCode:     FAILED_RETRIEVE_ORDER,
	},
	FAILED_CREATE_ORDER: {
		ClientMessage: "Failed to create order.",
		SystemMessage: "Something wrong happened while create order.",
		ErrorCode:     FAILED_CREATE_ORDER,
	},
	FAILED_UPDATE_ORDER: {
		ClientMessage: "Failed to update order.",
		SystemMessage: "Something wrong happened while update order.",
		ErrorCode:     FAILED_UPDATE_ORDER,
	},
	FAILED_DELETE_ORDER: {
		ClientMessage: "Failed to delete order.",
		SystemMessage: "Something wrong happened while delete order.",
		ErrorCode:     FAILED_DELETE_ORDER,
	},
	INVALID_REQUEST_RETRIEVE_ORDER: {
		ClientMessage: "Failed to request order.",
		SystemMessage: "Request has an invalid query params and/or payload to retrieve order.",
		ErrorCode:     INVALID_REQUEST_RETRIEVE_ORDER,
	},
	INVALID_REQUEST_CREATE_ORDER: {
		ClientMessage: "Failed to create order.",
		SystemMessage: "Request has an invalid query params and/or payload to create order.",
		ErrorCode:     INVALID_REQUEST_CREATE_ORDER,
	},
	INVALID_REQUEST_UPDATE_ORDER: {
		ClientMessage: "Failed to update order.",
		SystemMessage: "Request has an invalid query params and/or payload to update order.",
		ErrorCode:     INVALID_REQUEST_UPDATE_ORDER,
	},
	INVALID_REQUEST_DELETE_ORDER: {
		ClientMessage: "Failed to delete order.",
		SystemMessage: "Request has an invalid query params and/or payload to delete order.",
		ErrorCode:     INVALID_REQUEST_DELETE_ORDER,
	},
	INVALID_REQUEST_RETRIEVE_STORE: {
		ClientMessage: "Failed to request store.",
		SystemMessage: "Request has an invalid query params and/or payload to retrieve store.",
		ErrorCode:     INVALID_REQUEST_RETRIEVE_STORE,
	},
	FAILED_RETRIEVE_DATA: {
		ClientMessage: "Failed to retrieve Data.",
		SystemMessage: "Something wrong happened while retrieve Data.",
		ErrorCode:     FAILED_RETRIEVE_DATA,
	},
	STATUS_PAGE_NOT_FOUND: {
		ClientMessage: "Invalid Status Page.",
		SystemMessage: "Status Page Email Address not found.",
		ErrorCode:     STATUS_PAGE_NOT_FOUND,
	},
	INVALID_HEADER_X_PLATFORM: {
		ClientMessage: "Invalid platform.",
		SystemMessage: "Invalid value of header X-Platform.",
		ErrorCode:     INVALID_HEADER_X_PLATFORM,
	},
	INVALID_HEADER_X_API_KEY: {
		ClientMessage: "Invalid Api Key.",
		SystemMessage: "Invalid value of header X-Api-Key.",
		ErrorCode:     INVALID_HEADER_X_API_KEY,
	},
	UNAUTHORIZED: {
		ClientMessage: "Unauthorized",
		SystemMessage: "Unauthorized",
		ErrorCode:     UNAUTHORIZED,
	},
	FAILED_CREATE_DATA: {
		ClientMessage: "Failed to create data.",
		SystemMessage: "Something wrong happened while create data.",
		ErrorCode:     FAILED_CREATE_DATA,
	},
}
