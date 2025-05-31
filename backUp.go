package main
import "fmt"

const MaxStartups = 100
const MaxMembers = 100

type TeamMember struct {
	name string
	role string
}

type Startup struct {
	name     string
	founded  int
	funding  float64
	field    string
	category string
	team     [MaxMembers]TeamMember
	teamSize int
}

type StartupList [MaxStartups]Startup

var startupCount = 0

func main() {
	var startups StartupList
	var choice int

	showMenu()
	fmt.Scan(&choice)

	for choice != 0 {
		switch choice {
		case 1:
			addStartup(&startups, &startupCount)
		case 2:
			viewStartups(startups, startupCount)
		case 3:
			addTeamMember(&startups, startupCount)
		case 4:
			searchStartupByName(startups, startupCount)
		case 5:
			searchStartupByField(startups, startupCount)
		case 6:
			sortByFunding(&startups, startupCount)
		case 7:
			sortByYear(&startups, startupCount)
		case 8:
			reportByCategory(startups, startupCount)
		case 9:
			deleteStartup(&startups, &startupCount)
		default:
			fmt.Println("Invalid option.")
		}
		showMenu()
		fmt.Scan(&choice)
	}
}

func showMenu() {
	fmt.Println("=========================================")
	fmt.Println("1. Add Startup")
	fmt.Println("2. View Startups")
	fmt.Println("3. Add Team Member")
	fmt.Println("4. Search Startup by Name")
	fmt.Println("5. Search Startup by Field")
	fmt.Println("6. Sort by Funding (High to Low)")
	fmt.Println("7. Sort by Year (Old to New)")
	fmt.Println("8. Report by Category")
	fmt.Println("9. Delete Startup")
	fmt.Println("0. Exit")
	fmt.Print("Enter your choice: ")
}

func addStartup(S *StartupList, count *int) {
	if *count >= MaxStartups {
		fmt.Println("Startup list is full.")
		return
	}

	var s Startup
	fmt.Print("Name: ")
	fmt.Scan(&s.name)
	fmt.Print("Founded Year: ")
	fmt.Scan(&s.founded)
	fmt.Print("Funding: ")
	fmt.Scan(&s.funding)
	fmt.Print("Field: ")
	fmt.Scan(&s.field)
	fmt.Print("Category: ")
	fmt.Scan(&s.category)

	S[*count] = s
	*count++
	fmt.Println("Startup added successfully.")
}

func viewStartups(S StartupList, N int) {
	if N == 0 {
		fmt.Println("No startups to show.")
		return
	}
	for i := 0; i < N; i++ {
		fmt.Printf("[%d] %s (%d) - $%.2f - %s - %s\n", i+1, S[i].name, S[i].founded, S[i].funding, S[i].field, S[i].category)
		for j := 0; j < S[i].teamSize; j++ {
			fmt.Printf("   -> %s (%s)\n", S[i].team[j].name, S[i].team[j].role)
		}
	}
}

func addTeamMember(S *StartupList, N int) {
	var index int
	var member TeamMember
	
	if N == 0 {
		fmt.Println("No startups available.")
		return
	}
	fmt.Print("Enter startup number: ")
	fmt.Scan(&index)
	if index < 1 || index > N {
		fmt.Println("Invalid startup number.")
		return
	}

	fmt.Print("Team member name: ")
	fmt.Scan(&member.name)
	fmt.Print("Role: ")
	fmt.Scan(&member.role)

	startup := &S[index-1]
	if startup.teamSize >= MaxMembers {
		fmt.Println("Team is full.")
		return
	}
	startup.team[startup.teamSize] = member
	startup.teamSize++
	fmt.Println("Team member added.")
}

func searchStartupByName(S StartupList, N int) {
	var name string
	fmt.Print("Enter name to search: ")
	fmt.Scan(&name)

	for i := 0; i < N; i++ {
		if S[i].name == name {
			fmt.Printf("Found: %s (%d) - $%.2f - %s - %s\n", S[i].name, S[i].founded, S[i].funding, S[i].field, S[i].category)
			return
		}
	}
	fmt.Println("Startup not found.")
}

func searchStartupByField(S StartupList, N int) {
	var field string
	fmt.Print("Enter field to search: ")
	fmt.Scan(&field)

	for i := 0; i < N; i++ {
		if S[i].field == field {
			fmt.Printf("Found: %s (%d) - $%.2f - %s - %s\n", S[i].name, S[i].founded, S[i].funding, S[i].field, S[i].category)
			return
		}
	}
	fmt.Println("No startup found in that field.")
}

func sortByFunding(S *StartupList, N int) {
	for i := 0; i < N-1; i++ {
		for j := i + 1; j < N; j++ {
			if S[i].funding < S[j].funding {
				S[i], S[j] = S[j], S[i]
			}
		}
	}
	fmt.Println("Sorted by funding (descending).")
}

func sortByYear(S *StartupList, N int) {
	for i := 0; i < N-1; i++ {
		for j := i + 1; j < N; j++ {
			if S[i].founded > S[j].founded {
				S[i], S[j] = S[j], S[i]
			}
		}
	}
	fmt.Println("Sorted by year (ascending).")
}

func reportByCategory(S StartupList, N int) {
	if N == 0 {
		fmt.Println("No startups to report.")
		return
	}

	var categories [MaxStartups]string
	var count [MaxStartups]int
	var catCount int

	for i := 0; i < N; i++ {
		found := false
		for j := 0; j < catCount; j++ {
			if S[i].category == categories[j] {
				count[j]++
				found = true
				break
			}
		}
		if !found {
			categories[catCount] = S[i].category
			count[catCount] = 1
			catCount++
		}
	}

	fmt.Println("Report by Category:")
	for i := 0; i < catCount; i++ {
		fmt.Printf("- %s: %d startups\n", categories[i], count[i])
	}
}

func deleteStartup(S *StartupList, N *int) {
	var index int
	if *N == 0 {
		fmt.Println("No startups to delete.")
		return
	}

	fmt.Print("Enter startup number to delete: ")
	fmt.Scan(&index)
	if index < 1 || index > *N {
		fmt.Println("Invalid index.")
		return
	}

	for i := index - 1; i < *N-1; i++ {
		S[i] = S[i+1]
	}
	*N--
	fmt.Println("Startup deleted.")
}
