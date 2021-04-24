package domain

type ReversionRequest struct {
	Message string `json:"message"`
}

type ReversionResponse struct {
	MessageToRevert string `json:"message_to_revert"`
	RevertedMessage string `json:"reverted_message"`
}
