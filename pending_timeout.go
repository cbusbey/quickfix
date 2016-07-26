package quickfix

type pendingTimeout struct {
	inSession
}

func (currentState pendingTimeout) SendQueued(session *session) (nextState sessionState) {
	nextState = currentState.inSession.SendQueued(session)

	if !nextState.IsLoggedOn() {
		return nextState
	}

	return currentState
}

func (currentState pendingTimeout) Timeout(session *session, event event) (nextState sessionState) {
	switch event {
	case peerTimeout:
		session.log.OnEvent("Session Timeout")
		return latentState{}
	}

	return currentState
}
