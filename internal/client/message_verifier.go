package client

type messageVerifier struct {
}

func (verifier *messageVerifier) isValid(message []string) bool {
	return true
}

