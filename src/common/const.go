package common

type ResponseStatus string
type ResponseMessage string

const (
	Success ResponseStatus = "Success"
	Failed  ResponseStatus = "Failed"
	Error   ResponseStatus = "Error"

	NotFound               ResponseMessage = "Data not found"
	SlugRequired           ResponseMessage = "Slug is required"
	SuccessFetch           ResponseMessage = "Successfully fetched data"
	GuestNotValid          ResponseMessage = "Guest not valid"
	GuestWishLimitExceeded ResponseMessage = "Guest wish limit exceeded"
	SuccessPost            ResponseMessage = "Successfully posted data"
	AuthorizationFailed    ResponseMessage = "Authorization failed"
)