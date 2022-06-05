package models

//ResponseError representa o response
type ResponseError struct {
	DeveloperMessage string `json:"developerMessage"`
	UserMessage      string `json:"userMessage"`
	ErrorCode        int    `json:"errorCode"`
	MoreInfo         string `json:"moreInfo"`
}
