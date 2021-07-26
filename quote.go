package status

// Quote define stock quote
type Quote struct {
	Timestamp uint64
	Open      float32
	Close     float32
	High      float32
	Low       float32
	Volume    uint64
}
