package initialization

type SendLoadingMsg string

func (s SendLoadingMsg) String() string {
	return string(s)
}

type SendOutput string

func (s SendOutput) String() string {
	return string(s)
}
