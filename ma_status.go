package status

type MAStatus string

const (
	MAStatusRiseAbove MAStatus = "RiseAbove"
	MAStatusRiseBelow MAStatus = "RiseBelow"
	MAStatusRiseEqual MAStatus = "RisepEqual"
	MAStatusFallAbove MAStatus = "FallAbove"
	MAStatusFallBelow MAStatus = "FallBelow"
	MAStatusFallEqual MAStatus = "FallEqual"
	MAStatusKeepAbove MAStatus = "KeepAbove"
	MAStatusKeepBelow MAStatus = "KeepBelow"
	MAStatusKeepEqual MAStatus = "KeepEqual"
)

func (s MAStatus) String() string {
	return string(s)
}

func (s MAStatus) Equal(another Status) bool {
	maStatus, ok := another.(MAStatus)
	if !ok {
		return false
	}

	return s == maStatus
}
