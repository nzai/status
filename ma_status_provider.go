package status

const MAEqualMaxDegree float32 = 0.003

type MAStatusProvicedr struct {
	peroid int
}

func NewMAStatusProvider(peroid int) *MAStatusProvicedr {
	if peroid < 2 {
		panic("peroid must be greater than 2")
	}

	return &MAStatusProvicedr{peroid: peroid}
}

func (p MAStatusProvicedr) GetStatus(quotes []*Quote) (Status, error) {
	count := len(quotes)
	if count <= 1 {
		return MAStatusKeepEqual, nil
	}

	mas := p.caculatorMA(quotes)

	quoteDifference := quotes[count-1].Close - mas[count-1]
	maDifference := mas[count-1] - mas[count-2]
	maDegree := maDifference / mas[count-2]

	if maDegree <= MAEqualMaxDegree && maDegree >= -MAEqualMaxDegree {
		if quoteDifference == 0 {
			return MAStatusKeepEqual, nil
		}

		if quoteDifference > 0 {
			return MAStatusKeepAbove, nil
		}

		return MAStatusKeepBelow, nil
	}

	if maDifference > 0 {
		if quoteDifference == 0 {
			return MAStatusRiseEqual, nil
		}

		if quoteDifference > 0 {
			return MAStatusRiseAbove, nil
		}

		return MAStatusRiseBelow, nil
	}

	if quoteDifference == 0 {
		return MAStatusFallEqual, nil
	}

	if quoteDifference > 0 {
		return MAStatusFallAbove, nil
	}

	return MAStatusFallBelow, nil
}

func (p MAStatusProvicedr) caculatorMA(quotes []*Quote) []float32 {
	mas := make([]float32, len(quotes))
	for index, quote := range quotes {
		if index == 0 {

			mas[0] = quote.Close / float32(p.peroid)
			continue
		}

		mas[index] = (mas[index-1]*(float32(p.peroid-1)) + quote.Close) / float32(p.peroid)
	}

	return mas
}
