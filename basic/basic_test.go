package basicpaxos

import (
	"testing"
)

type TestClient struct {
	OnSendPrepare func(v interface{})
}

func (c TestClient) SendPrepare(v interface{}) {
	c.OnSendPrepare(v)
}
func TestProposer(t *testing.T) {
	t.Run("SendPrepare", func(t *testing.T) {
		proposeCalled := 0
		proposeValue := 0
		client := TestClient{
			OnSendPrepare: func(v interface{}) {
				proposeCalled++
				proposeValue = v.(int)
			},
		}
		proposer := Proposer{
			Comm: client,
		}
		err := proposer.Start(1234)
		if err != nil {
			t.Error("Error calling Start")
		}
		t.Run("SendPropose should have been called", func(t *testing.T) {})
		if proposeCalled != 1 {
			t.Error()
		}
		t.Run("Propose internal value should match the start value", func(t *testing.T) {
			if proposer.v != 1234 {
				t.Error("Got:", proposer.v)
			}
		})
		t.Run("Propose should not accept a second start", func(t *testing.T) {
			err = proposer.Start(2345)
			if err == nil {
				t.Error()
			}
		})
		t.Run("SendPropose should not have been called again", func(t *testing.T) {
			if proposeCalled > 1 {
				t.Error()
			}
		})
		t.Run("the value sent to sendpropose should be the propose interal value", func(t *testing.T) {
			if proposeValue != 1234 {
				t.Fail()
			}
		})
	})
}

func TestProposerConcurrency(t *testing.T) {
	t.Run("two clients using the same proposer concurrently", func(t *testing.T) {
		t.Skip("not implemented")
	})
}
