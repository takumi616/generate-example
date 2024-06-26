// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package handler

import (
	"context"
	"github.com/takumi616/go-postgres-docker-restapi/entity"
	"sync"
)

// Ensure, that SentenceFetcherMock does implement SentenceFetcher.
// If this is not the case, regenerate this file with moq.
var _ SentenceFetcher = &SentenceFetcherMock{}

// SentenceFetcherMock is a mock implementation of SentenceFetcher.
//
//	func TestSomethingThatUsesSentenceFetcher(t *testing.T) {
//
//		// make and configure a mocked SentenceFetcher
//		mockedSentenceFetcher := &SentenceFetcherMock{
//			FetchSentenceListFunc: func(ctx context.Context) ([]entity.Sentence, error) {
//				panic("mock out the FetchSentenceList method")
//			},
//			FetchSingleSentenceFunc: func(ctx context.Context, id string) (entity.Sentence, error) {
//				panic("mock out the FetchSingleSentence method")
//			},
//		}
//
//		// use mockedSentenceFetcher in code that requires SentenceFetcher
//		// and then make assertions.
//
//	}
type SentenceFetcherMock struct {
	// FetchSentenceListFunc mocks the FetchSentenceList method.
	FetchSentenceListFunc func(ctx context.Context) ([]entity.Sentence, error)

	// FetchSingleSentenceFunc mocks the FetchSingleSentence method.
	FetchSingleSentenceFunc func(ctx context.Context, id string) (entity.Sentence, error)

	// calls tracks calls to the methods.
	calls struct {
		// FetchSentenceList holds details about calls to the FetchSentenceList method.
		FetchSentenceList []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// FetchSingleSentence holds details about calls to the FetchSingleSentence method.
		FetchSingleSentence []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID string
		}
	}
	lockFetchSentenceList   sync.RWMutex
	lockFetchSingleSentence sync.RWMutex
}

// FetchSentenceList calls FetchSentenceListFunc.
func (mock *SentenceFetcherMock) FetchSentenceList(ctx context.Context) ([]entity.Sentence, error) {
	if mock.FetchSentenceListFunc == nil {
		panic("SentenceFetcherMock.FetchSentenceListFunc: method is nil but SentenceFetcher.FetchSentenceList was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockFetchSentenceList.Lock()
	mock.calls.FetchSentenceList = append(mock.calls.FetchSentenceList, callInfo)
	mock.lockFetchSentenceList.Unlock()
	return mock.FetchSentenceListFunc(ctx)
}

// FetchSentenceListCalls gets all the calls that were made to FetchSentenceList.
// Check the length with:
//
//	len(mockedSentenceFetcher.FetchSentenceListCalls())
func (mock *SentenceFetcherMock) FetchSentenceListCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockFetchSentenceList.RLock()
	calls = mock.calls.FetchSentenceList
	mock.lockFetchSentenceList.RUnlock()
	return calls
}

// FetchSingleSentence calls FetchSingleSentenceFunc.
func (mock *SentenceFetcherMock) FetchSingleSentence(ctx context.Context, id string) (entity.Sentence, error) {
	if mock.FetchSingleSentenceFunc == nil {
		panic("SentenceFetcherMock.FetchSingleSentenceFunc: method is nil but SentenceFetcher.FetchSingleSentence was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  string
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockFetchSingleSentence.Lock()
	mock.calls.FetchSingleSentence = append(mock.calls.FetchSingleSentence, callInfo)
	mock.lockFetchSingleSentence.Unlock()
	return mock.FetchSingleSentenceFunc(ctx, id)
}

// FetchSingleSentenceCalls gets all the calls that were made to FetchSingleSentence.
// Check the length with:
//
//	len(mockedSentenceFetcher.FetchSingleSentenceCalls())
func (mock *SentenceFetcherMock) FetchSingleSentenceCalls() []struct {
	Ctx context.Context
	ID  string
} {
	var calls []struct {
		Ctx context.Context
		ID  string
	}
	mock.lockFetchSingleSentence.RLock()
	calls = mock.calls.FetchSingleSentence
	mock.lockFetchSingleSentence.RUnlock()
	return calls
}
