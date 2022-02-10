package event

type Signer interface {
	Sign(dat []byte) ([]byte,error)
}

type EventProof interface {
}

type LogRoot []byte //root of log

type Verifier interface {
	Verify(*Event, EventProof, LogRoot) bool	
}

type EventLog interface {
	Add(e *Event) error
	Prove(id EventID) EventProof
	Signer() Signer
}
