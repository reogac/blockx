package event


type Where	struct {
	location	string
	name		string
	desc		string
}

func (w *Where) String () string {
	return w.name
}


