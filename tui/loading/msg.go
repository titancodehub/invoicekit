package loading

type SendLoadingMsg string

func (s SendLoadingMsg) String() string {
	return string(s)
}
