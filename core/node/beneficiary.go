package node

func Address(address string) func() string {
	return func() string {
		return address
	}
}