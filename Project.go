package main

import (
	"fmt"
	"strings"
)

const (
	maxStartups    = 100
	maxTeamMembers = 10
)

type TeamMember struct {
	Name string
	Role string
}

type Oger struct {
	Name     string
	Founded  int
	Funding  float64
	Category string
	Team     [maxTeamMembers]TeamMember
	TeamSize int
}

var startups [maxStartups]Oger
var startupCount int

func main() {
	for {
		fmt.Println("\n=== Startup Management Menu ===")
		fmt.Println("1. Add Startup")
		fmt.Println("2. View Startups")
		fmt.Println("3. Add Team Member")
		fmt.Println("4. Search Startup by Name")
		fmt.Println("5. Sort Startups by Funding (Selection Sort)")
		fmt.Println("6. Report by Category")
		fmt.Println("7. Delete Startup")
		fmt.Println("8. Exit")

		var choice int
		fmt.Print("Enter choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			addStartup()
		case 2:
			viewStartups()
		case 3:
			addTeamMember()
		case 4:
			searchByNameSequential()
		case 5:
			sortByFundingSelection()
		case 6:
			reportByCategory()
		case 7:
			deleteStartup()
		case 8:
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func addStartup() {
	if startupCount >= maxStartups {
		fmt.Println("Cannot add more startups.")
		return
	}

	var s Oger
	fmt.Print("Enter startup name: ")
	fmt.Scan(&s.Name)
	fmt.Print("Enter year founded: ")
	fmt.Scan(&s.Founded)
	fmt.Print("Enter funding amount: ")
	fmt.Scan(&s.Funding)
	fmt.Print("Enter category: ")
	fmt.Scan(&s.Category)

	startups[startupCount] = s
	startupCount++
	fmt.Println("Startup added successfully.")
}

func viewStartups() {
	if startupCount == 0 {
		fmt.Println("No startups available.")
		return
	}

	for i := 0; i < startupCount; i++ {
		s := startups[i]
		fmt.Printf("[%d] %s (%d) - $%.2f - %s\n", i+1, s.Name, s.Founded, s.Funding, s.Category)
		for j := 0; j < s.TeamSize; j++ {
			t := s.Team[j]
			fmt.Printf("   - %s: %s\n", t.Name, t.Role)
		}
	}
}

func addTeamMember() {
	if startupCount == 0 {
		fmt.Println("No startups to assign team members.")
		return
	}

	var index int
	fmt.Print("Enter startup index: ")
	fmt.Scan(&index)

	if index <= 0 || index > startupCount {
		fmt.Println("Invalid startup index.")
		return
	}

	s := &startups[index-1]
	if s.TeamSize >= maxTeamMembers {
		fmt.Println("Team is full.")
		return
	}

	var tm TeamMember
	fmt.Print("Enter team member name: ")
	fmt.Scan(&tm.Name)
	fmt.Print("Enter role: ")
	fmt.Scan(&tm.Role)

	s.Team[s.TeamSize] = tm
	s.TeamSize++
	fmt.Println("Team member added.")
}

func searchByNameSequential() {
	var name string
	fmt.Print("Enter startup name to search: ")
	fmt.Scan(&name)

	found := false
	i := 0

	for i < startupCount {
		if strings.EqualFold(startups[i].Name, name) {
			fmt.Printf("Found: %s (%d) - $%.2f - %s\n", startups[i].Name, startups[i].Founded, startups[i].Funding, startups[i].Category)
			found = true
			i = startupCount
		} else {
			i++
		}
	}

	if !found {
		fmt.Println("Startup not found.")
	}
}


func sortByFundingSelection() {
	for i := 0; i < startupCount-1; i++ {
		minIdx := i
		for j := i + 1; j < startupCount; j++ {
			if startups[j].Funding < startups[minIdx].Funding {
				minIdx = j
			}
		}
		startups[i], startups[minIdx] = startups[minIdx], startups[i]
	}
	fmt.Println("Startups sorted by funding.")
}

func reportByCategory() {
	if startupCount == 0 {
		fmt.Println("No data available.")
		return
	}

	var categories [maxStartups]string
	var counts [maxStartups]int
	categoryCount := 0

	for i := 0; i < startupCount; i++ {
		cat := startups[i].Category
		found := false

		j := 0
		for j < categoryCount && !found {
			if categories[j] == cat {
				counts[j]++
				found = true
			}
			j++
		}

		if !found {
			categories[categoryCount] = cat
			counts[categoryCount] = 1
			categoryCount++
		}
	}

	fmt.Println("Report: Number of Startups per Category")
	i := 0
	for i < categoryCount {
		fmt.Printf("- %s: %d\n", categories[i], counts[i])
		i++
	}
}

func deleteStartup() {
	if startupCount == 0 {
		fmt.Println("No startups to delete.")
		return
	}

	var index int
	fmt.Print("Enter the index of the startup to delete: ")
	fmt.Scan(&index)

	if index <= 0 || index > startupCount {
		fmt.Println("Invalid index.")
		return
	}

	fmt.Println("Before deletion:")
	viewStartups()

	for i := index - 1; i < startupCount-1; i++ {
		startups[i] = startups[i+1]
	}

	startupCount--

	fmt.Println("After deletion:")
	viewStartups()
}