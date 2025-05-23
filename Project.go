package main

import (
	"fmt"
	"strings"
)

const (
	maxStartups    = 100
)

type TeamMember struct {
	Name string
	Role string
}

type STRP struct {
	Name     string
	Founded  int
	Funding  float64
	Category string
	Team     []TeamMember
	TeamSize int
}

func main() {
	var startups [maxStartups]STRP
	var N int

	for {
		fmt.Println("\n=== Startup Management Menu ===")
		fmt.Println("1. Add Startup")
		fmt.Println("2. View Startups")
		fmt.Println("3. Add Team Member")
		fmt.Println("4. Search Startup by Name")
		fmt.Println("5. Sort Startups by Funding(Descending)")
		fmt.Println("6. Sort Startups by Year(Ascending)")
		fmt.Println("7. Report by Category")
		fmt.Println("8. Delete Startup")
		fmt.Println("9. Exit")

		var choice int
		fmt.Print("Enter choice: ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			addStartup(&startups, &N)
		case 2:
			viewStartups(&startups, N)
		case 3:
			addTeamMember(&startups, N)
		case 4:
			searchByNameSequential(&startups, N)
		case 5:
			sortByFundingSelection(&startups, N)
		case 6:
			sortByYearSelection(&startups, N)
		case 7:
			reportByCategory(&startups, N)
		case 8:
			deleteStartup(&startups, &N)
		case 9:
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func addStartup(startups *[maxStartups]STRP, N *int) {
	if *N >= maxStartups {
		fmt.Println("Cannot add more startups.")
		return
	}
	var s STRP
	fmt.Print("Enter startup name: ")
	fmt.Scan(&s.Name)
	fmt.Print("Enter year founded: ")
	fmt.Scan(&s.Founded)
	fmt.Print("Enter funding amount: ")
	fmt.Scan(&s.Funding)
	fmt.Print("Enter category: ")
	fmt.Scan(&s.Category)
	startups[*N] = s
	(*N)++
	fmt.Println("Startup added successfully.")
}

func viewStartups(startups *[maxStartups]STRP, N int) {
	if N == 0 {
		fmt.Println("No startups available.")
		return
	}
	for i := 0; i < N; i++ {
		s := startups[i]
		fmt.Printf("[%d] %s (%d) - $%.2f - %s\n", i+1, s.Name, s.Founded, s.Funding, s.Category)
		for j := 0; j < s.TeamSize; j++ {
			t := s.Team[j]
			fmt.Printf("   - %s: %s\n", t.Name, t.Role)
		}
	}
}

func addTeamMember(startups *[maxStartups]STRP, N int) {
	var tm TeamMember
	var index int

    if N == 0 {
		fmt.Println("No startups to assign team members.")
		return
	}
	fmt.Print("Enter startup index: ")
	fmt.Scan(&index)

	if index <= 0 || index > N {
		fmt.Println("Invalid startup index.")
		return
	}
	s := &startups[index-1]
	fmt.Print("Enter team member name: ")
	fmt.Scan(&tm.Name)
	fmt.Print("Enter role: ")
	fmt.Scan(&tm.Role)
	s.Team[s.TeamSize] = tm
	s.TeamSize++
	fmt.Println("Team member added.")
}

func searchByNameSequential(startups *[maxStartups]STRP, N int) {
	var name string
	fmt.Print("Enter startup name to search: ")
	fmt.Scan(&name)
	found := false
	i := 0

	for i < N {
		if strings.EqualFold(startups[i].Name, name) {
			fmt.Printf("Found: %s (%d) - $%.2f - %s\n", startups[i].Name, startups[i].Founded, startups[i].Funding, startups[i].Category)
			found = true
			i = N
		} else {
			i++
		}
	}
	if found == false  {
		fmt.Println("Startup not found.")
	}
}

func sortByFundingSelection(startups *[maxStartups]STRP, N int) {
	var pass, idx, i int
	var temp STRP

	pass = 1
	for pass <= N-1 {
		idx = pass - 1
		i = pass
		for i < N {
			if startups[idx].Funding < startups[i].Funding {
				idx = i
			}
			i++
		}
		temp = startups[pass-1]
		startups[pass-1] = startups[idx]
		startups[idx] = temp
		pass++
	}
	fmt.Println("Startups sorted by funding(Descending).")
}

func sortByYearSelection(startups *[maxStartups]STRP, N int) {
	var pass, idx, i int
	var temp STRP

	pass = 1
	for pass <= N-1 {
		idx = pass - 1
		i = pass
		for i < N {
			if startups[idx].Founded > startups[i].Founded {
				idx = i
			}
			i++
		}
		temp = startups[pass-1]
		startups[pass-1] = startups[idx]
		startups[idx] = temp
		pass++
	}
	fmt.Println("Startups sorted by year (Ascending).")
}

func reportByCategory(startups *[maxStartups]STRP, N int) {
	if N == 0 {
		fmt.Println("No data available.")
		return
	}
	var categories [maxStartups]string
	var counts [maxStartups]int
	categoryCount := 0

	for i := 0; i < N; i++ {
		cat := startups[i].Category
		found := false

		for j := 0; j < categoryCount && !found; j++ {
			if categories[j] == cat {
				counts[j]++
				found = true
			}
		}
		if found == false {
			categories[categoryCount] = cat
			counts[categoryCount] = 1
			categoryCount++
		}
	}
	fmt.Println("Report: Number of Startups per Category")
	for i := 0; i < categoryCount; i++ {
		fmt.Printf("- %s: %d\n", categories[i], counts[i])
	}
}

func deleteStartup(startups *[maxStartups]STRP, N *int) {
	if *N == 0 {
		fmt.Println("No startups to delete.")
		return
	}
	var index int
	fmt.Print("Enter the index of the startup to delete: ")
	fmt.Scan(&index)

	if index <= 0 || index > *N {
		fmt.Println("Invalid index.")
		return
	}
	fmt.Println("Before deletion:")
	viewStartups(startups, *N)

	for i := index - 1; i < *N-1; i++ {
		startups[i] = startups[i+1]
	}
	*N-- 
	fmt.Println("After deletion:")
	viewStartups(startups, *N)
}