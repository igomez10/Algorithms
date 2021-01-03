package main

import (
	"container/heap"
)

type Logger struct {
	Messages       map[string]int
	Stream         *LogStream
	BannedMessages map[string]bool
}

/** Initialize your data structure here. */
func Constructor() Logger {
	logger := Logger{}

	logger.Stream = &LogStream{}
	logger.BannedMessages = map[string]bool{}
	return logger
}

/** Returns true if the message should be printed in the given timestamp, otherwise returns false.
  If this method returns false, the message will not be printed.
  The timestamp is in seconds granularity. */
func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	// prune messages older than timestamp + 10
	for this.Stream.Len() > 0 && timestamp-this.Stream.GetOldestLog().Timestamp >= 10 {
		delete(this.BannedMessages, this.Stream.GetOldestLog().Message)
		heap.Pop(this.Stream)
	}

	// check if message is in heap
	if _, ok := this.BannedMessages[message]; ok {
		return false
	}

	// add message to heap and map
	heap.Push(this.Stream, Log{Timestamp: timestamp, Message: message})
	this.BannedMessages[message] = true

	return true
}

type LogStream []Log

type Log struct {
	Timestamp int
	Message   string
}

func (l LogStream) Less(i, j int) bool {
	return l[i].Timestamp < l[j].Timestamp
}

func (l LogStream) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l LogStream) Len() int {
	return len(l)
}

func (l *LogStream) Pop() interface{} {
	lDeref := *l
	old := lDeref[len(lDeref)-1]
	*l = lDeref[:len(lDeref)-1]
	return old
}

func (l *LogStream) Push(n interface{}) {
	*l = append(*l, n.(Log))
}

func Answer(events []string, logs []Log) []*bool {
	res := make([]*bool, len(events))

	var l Logger
	for i := range events {
		switch events[i] {
		case "Logger":
			l = Constructor()

		case "shouldPrintMessage":
			ans := l.ShouldPrintMessage(logs[i].Timestamp, logs[i].Message)
			res[i] = &ans
		}
	}

	return res
}

func (l LogStream) GetOldestLog() Log {
	return l[0]
}
