package shared

type BPMNStruct interface {
	Validate(string) []error
}

type BPMNEnum interface {
	String() string
	EnumIndex() int
	ToEnum(input string)
}
