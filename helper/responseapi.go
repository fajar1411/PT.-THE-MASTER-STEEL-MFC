package helper

import "github.com/google/uuid"

type Response struct {
	Code        int    `json:"-"`
	Status      string `json:"status"`
	IsError     bool   `json:"isError"`
	Data        any    `json:"data"`
	Description string `json:"deskripsi"`
}

func GetResponse(data any, code int, isError bool, desription string) *Response {

	if isError {
		return &Response{
			Code:        code,
			Status:      getStatus(code),
			IsError:     isError,
			Description: desription,
		}

	}
	return &Response{
		Code:        code,
		Status:      getStatus(code),
		IsError:     isError,
		Description: desription,
		Data:        data,
	}

}

func getStatus(code int) (desc string) {

	switch code {
	case 200:
		desc = "OK"
	case 201:
		desc = "Created"
	case 202:
		desc = "Accepted"
	case 304:
		desc = "Not Modified"
	case 400:
		desc = "Bad Request"
	case 401:
		desc = "Unauthorized"
	case 403:
		desc = "Forbidden"
	case 404:
		desc = "Not Found"
	case 415:
		desc = "Unsupported Media Type"
	case 500:
		desc = "Internal Server Error"
	case 502:
		desc = "Bad Gateway"
	default:
		desc = "Status Code Undefined"
	}

	return

}
func convertUid() string {
	return uuid.New().String()
}
