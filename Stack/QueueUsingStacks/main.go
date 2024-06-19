package main

import (
	"fmt"
	"strconv"
	"strings"
)

type MyQueue struct {
	stack1 Stack
	stack2 Stack
}

func (q *MyQueue) push(value int) {
	for !q.stack1.isEmpty() {
		q.stack2.push(q.stack1.pop())
	}
	q.stack1.push(value)
	for !q.stack2.isEmpty() {
		q.stack1.push(q.stack2.pop())
	}
}

func (q *MyQueue) pop() int {
	return q.stack1.pop()
}

func (q *MyQueue) peek() int {
	return q.stack1.top()
}

func (q *MyQueue) empty() bool {
	return q.stack1.isEmpty()
}

func main() {
	inputQueues := [][][]string{
		{{"9", "3", "1", "", "", ""}, {"push", "push", "push", "pop", "peek", "empty"}},
		{{"10", "6", "", "", ""}, {"push", "push", "pop", "empty", "peek"}},
		{{"1", "2", "3", "", "", "", "", ""}, {"push", "push", "push", "peek", "pop", "pop", "pop", "empty"}},
		{{"14", "", "66", ""}, {"push", "pop", "push", "pop"}},
		{{"4", ""}, {"push", "peek"}},
	}

	for i := 0; i < len(inputQueues); i++ {
		fmt.Printf("%d.\t Starting operations:\n", i+1)

		// Initialize a queue
		queueObj := MyQueue{}

		// Loop over all the commands
		for j := 0; j < len(inputQueues[i][1]); j++ {
			if inputQueues[i][1][j] == "push" {
				inputStr := inputQueues[i][1][j] + "(" + inputQueues[i][0][j] + ")"
				fmt.Printf("\t\t%s\n", inputStr)
				x, err := strconv.Atoi(inputQueues[i][0][j])
				if err != nil {
					panic(err)
				}
				queueObj.push(x)
			} else if inputQueues[i][1][j] == "pop" {
				inputStr := inputQueues[i][1][j] + "(" + inputQueues[i][0][j] + ")"
				fmt.Printf("\t\t%s   returns %d\n", inputStr, queueObj.pop())
			} else if inputQueues[i][1][j] == "empty" {
				inputStr := inputQueues[i][1][j] + "(" + inputQueues[i][0][j] + ")"
				fmt.Printf("\t\t%s returns %v\n", inputStr, queueObj.empty())
			} else if inputQueues[i][1][j] == "peek" {
				inputStr := inputQueues[i][1][j] + "(" + inputQueues[i][0][j] + ")"
				fmt.Printf("\t\t%s  returns %d\n", inputStr, queueObj.peek())
			}
		}

		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
