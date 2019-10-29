package messages

const alfredVersion uint8 = 0

// TODO: refactor and move to individual message files

type MessageType int

type AlfredRequestBody struct {
	Type          uint8
	TransactionId uint16
}

type AlfredStatusBody struct {
	Type            uint16
	NumberOfPackets uint16
}

type AlfredStatusErrorBody struct {
	Type      uint16
	ErrorCode uint16
}

const (
	alfredPushDataType       MessageType = 0
	alfredAccounceMasterType MessageType = 1
	alfredRequestType        MessageType = 2
	alfredStatusType         MessageType = 3
)
