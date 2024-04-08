// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package service

import (
	"context"
	"github.com/takumi616/english-example-sentence-api/entity"
	"sync"
)

// Ensure, that SentenceUpdaterMock does implement SentenceUpdater.
// If this is not the case, regenerate this file with moq.
var _ SentenceUpdater = &SentenceUpdaterMock{}

// SentenceUpdaterMock is a mock implementation of SentenceUpdater.
//
//	func TestSomethingThatUsesSentenceUpdater(t *testing.T) {
//
//		// make and configure a mocked SentenceUpdater
//		mockedSentenceUpdater := &SentenceUpdaterMock{
//			UpdateSentenceFunc: func(ctx context.Context, sentenceID int, body string) (entity.Sentence, error) {
//				panic("mock out the UpdateSentence method")
//			},
//		}
//
//		// use mockedSentenceUpdater in code that requires SentenceUpdater
//		// and then make assertions.
//
//	}
type SentenceUpdaterMock struct {
	// UpdateSentenceFunc mocks the UpdateSentence method.
	UpdateSentenceFunc func(ctx context.Context, sentenceID int, body string) (entity.Sentence, error)

	// calls tracks calls to the methods.
	calls struct {
		// UpdateSentence holds details about calls to the UpdateSentence method.
		UpdateSentence []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// SentenceID is the sentenceID argument value.
			SentenceID int
			// Body is the body argument value.
			Body string
		}
	}
	lockUpdateSentence sync.RWMutex
}

// UpdateSentence calls UpdateSentenceFunc.
func (mock *SentenceUpdaterMock) UpdateSentence(ctx context.Context, sentenceID int, body string) (entity.Sentence, error) {
	if mock.UpdateSentenceFunc == nil {
		panic("SentenceUpdaterMock.UpdateSentenceFunc: method is nil but SentenceUpdater.UpdateSentence was just called")
	}
	callInfo := struct {
		Ctx        context.Context
		SentenceID int
		Body       string
	}{
		Ctx:        ctx,
		SentenceID: sentenceID,
		Body:       body,
	}
	mock.lockUpdateSentence.Lock()
	mock.calls.UpdateSentence = append(mock.calls.UpdateSentence, callInfo)
	mock.lockUpdateSentence.Unlock()
	return mock.UpdateSentenceFunc(ctx, sentenceID, body)
}

// UpdateSentenceCalls gets all the calls that were made to UpdateSentence.
// Check the length with:
//
//	len(mockedSentenceUpdater.UpdateSentenceCalls())
func (mock *SentenceUpdaterMock) UpdateSentenceCalls() []struct {
	Ctx        context.Context
	SentenceID int
	Body       string
} {
	var calls []struct {
		Ctx        context.Context
		SentenceID int
		Body       string
	}
	mock.lockUpdateSentence.RLock()
	calls = mock.calls.UpdateSentence
	mock.lockUpdateSentence.RUnlock()
	return calls
}