//Challenge 4 · Issue Tracking Service

// Goal: Design a lightweight service similar to Jira.
// Tasks:
// Define an Issue struct with ID, Title, Status, Assignee, and CreatedAt.
// Implement CreateIssue(title string, assignee string) *Issue and ListIssuesByAssignee(assignee string) []*Issue.
// Store issues in map[int]*Issue with incremental IDs.
// Ensure concurrency safety; multiple goroutines may invoke the functions.
// Demonstrate usage in main().

package main

import (
	"fmt"
	"sync"
	"time"
)

type Issue struct {
	ID        int
	Title     string
	Status    string
	Assignee  string
	CreatedAt time.Time
}

type IssueStore struct {
	mu     sync.RWMutex
	lastID int
	items  map[int]*Issue
}

func (IssueStore *IssueStore) CreateIssue(title string, assignee string) *Issue {
	IssueStore.mu.Lock()
	defer IssueStore.mu.Unlock()
	newID := IssueStore.lastID + 1
	is := &Issue{ID: newID, Title: title,
		Assignee:  assignee,
		Status:    "Open",
		CreatedAt: time.Now()}
	IssueStore.items[newID] = is
	IssueStore.lastID = newID

	return is
}

func (IssueStore *IssueStore) ListIssuesByAssignee(assignee string) []*Issue {
	slice := []*Issue{}
	IssueStore.mu.RLock()
	defer IssueStore.mu.RUnlock()

	for _, j := range IssueStore.items {
		if j.Assignee == assignee {
			slice = append(slice, j)
		}
	}
	return slice
}

func main() {
	issueStore := IssueStore{items: make(map[int]*Issue)}

	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		issueStore.CreateIssue("first issue", "Himanshu")
	}()
	go func() {
		defer wg.Done()
		issueStore.CreateIssue("second issue", "Himanshu")
	}()
	go func() {
		defer wg.Done()
		issueStore.CreateIssue("third issue", "Mehta")
	}()

	wg.Wait()

	for _, iss := range issueStore.ListIssuesByAssignee("Himanshu") {
		fmt.Println(iss.ID, iss.Title, iss.Status, iss.CreatedAt)
	}

}
