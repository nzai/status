package status

type Status interface {
	String() string
	Equal(Status) bool
}
