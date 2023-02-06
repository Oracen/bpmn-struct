package process

type LaneSet struct {
}

func CreateLaneSet() LaneSet {
	return LaneSet{}
}

func (m LaneSet) Validate(name string) []error {
	return []error{}
}
