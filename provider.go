package status

type Provider interface {
	GetStatus(quotes []*Quote) (Status, error)
}
