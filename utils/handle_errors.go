package utils

import "github.com/gin-gonic/gin"

// declare errorReponse type 
type ErrorResponse struct {
    Error   string `json:"error"` // specify the struct to response 
    Message string `json:"message"`
}

// SendError is a utility function to send a standardized error response
// gin.Context handle Http request(create update insert delete) just waiter hanlde request form customer
func SendError(handleRequest *gin.Context, statusCode int, message string, err_message error) {
    handleRequest.JSON(statusCode, ErrorResponse{
        Error:   err_message.Error(),
        Message: message,
    })
    handleRequest.Abort() // Abort use for ensuring response is sent immediately!
}

// custome success message 
func SendSuccess(handleRequest *gin.Context, statusCode int, data interface{}) {
    handleRequest.JSON(statusCode, data);
}

// custome success message 
func SendDeleteSuccess(handleRequest *gin.Context, statusCode int, message string) {
    handleRequest.JSON(statusCode,message);
}
