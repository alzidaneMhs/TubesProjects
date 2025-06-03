package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const MaxStartups = 100
const MaxMembers = 100

type TeamMember struct {
	Name string
	Role string
}

type Startup struct {
	Name     string `json:"name"`
	Founded  int
	Funding  float64
	Field    string
	Category string
	Team     [MaxMembers]TeamMember
	TeamSize int
}

type StartupList [MaxStartups]Startup

var startupCount = 0

func main() {
	var startups StartupList

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

router.GET("/", func(c *gin.Context) {
    c.HTML(http.StatusOK, "index.html", gin.H{
        "title": "Home",
    })
})


// This for View your Startup Router
router.GET("/view", func(c *gin.Context) {
	c.HTML(http.StatusOK, "view.html", gin.H{
		"title": "View Startups",
		"content": "view",
		"startups": startups[:startupCount],
	})
})

// This the Router of AddStartup
router.GET("/add", func(c *gin.Context) {
	c.HTML(http.StatusOK, "add.html", gin.H{
		"title": "Add Startup",
		"content": "add",
	})
})

// This the AddStartup
router.POST("/add", func(c *gin.Context) {
    var s Startup
	var fund = c.PostForm("funding")
	var found = c.PostForm("founded")
	number, err := strconv.ParseFloat(fund, 64)
 	namber, rre := strconv.Atoi(found)

	if err != nil {
		fmt.Print("error")
	} 

	if rre != nil {
		fmt.Print("error")
	}

    s.Name = c.PostForm("name")
	s.Funding = number
	s.Founded = namber
    s.Field = c.PostForm("field")	
    s.Category = c.PostForm("category")

    if startupCount < MaxStartups {
        startups[startupCount] = s
        startupCount++
        c.Redirect(http.StatusFound, "/view")
    } else {
        c.String(http.StatusBadRequest, "Startup list is full.")
    }
	fmt.Print(s)
})

// This for Add Member
router.POST("/add-member", func(c *gin.Context) {
    idxStr := c.PostForm("index")
    name := c.PostForm("member_name")
    role := c.PostForm("member_role")

    idx, err := strconv.Atoi(idxStr)
    if err != nil || idx < 0 || idx >= startupCount {
        c.String(http.StatusBadRequest, "Invalid startup index")
        return
    }

    if startups[idx].TeamSize >= MaxMembers {
        c.String(http.StatusBadRequest, "Team is full")
        return
    }

    startups[idx].Team[startups[idx].TeamSize] = TeamMember{Name: name, Role: role}
    startups[idx].TeamSize++

    c.Redirect(http.StatusFound, "/view")
})

// This for delete grrr
router.POST("/delete", func(c *gin.Context) {
    idxStr := c.PostForm("index")
    idx, err := strconv.Atoi(idxStr)
    if err != nil || idx < 0 || idx >= startupCount {
        c.String(http.StatusBadRequest, "Invalid index")
        return
    }

    for i := idx; i < startupCount-1; i++ {
        startups[i] = startups[i+1]
    }
    startupCount--

    c.Redirect(http.StatusFound, "/view")
})

// This for Search UI page
router.GET("/search", func(c *gin.Context) {
    c.HTML(http.StatusOK, "view.html", gin.H{
        "title": "Search",
        "content": "search",
    })
})

// Search by name
router.GET("/search-by-name", func(c *gin.Context) {
    name := c.Query("name")
    var result []Startup
    for i := 0; i < startupCount; i++ {
        if startups[i].Name == name {
            result = append(result, startups[i])
        }
    }
    c.HTML(http.StatusOK, "view.html", gin.H{
        "title": "Search Result",
        "results": result,
        "startups": startups[:startupCount],
    })
})

// Search by field
router.GET("/search-by-field", func(c *gin.Context) {
    field := c.Query("field")
    var result []Startup
    for i := 0; i < startupCount; i++ {
        if startups[i].Field == field {
            result = append(result, startups[i])
        }
    }
    c.HTML(http.StatusOK, "view.html", gin.H{
        "title": "Field Search",
        "results": result,
        "startups": startups[:startupCount],
    })
})

// Report by category
router.GET("/report-category", func(c *gin.Context) {
    categoryMap := make(map[string]int)
    for i := 0; i < startupCount; i++ {
        cat := startups[i].Category
        categoryMap[cat]++
    }
    c.HTML(http.StatusOK, "view.html", gin.H{
        "title": "Category Report",
        "report": categoryMap,
        "startups": startups[:startupCount],
    })
})

	router.Run()
	// var startups StartupList
	// var choice int

	// showMenu()
	// fmt.Scan(&choice)

	// for choice != 0 {
	// 	switch choice {
	// 	case 1:
	// 		addStartup(&startups, &startupCount)
	// 	case 2:
	// 		viewStartups(startups, startupCount)
	// 	case 3:
	// 		addTeamMember(&startups, startupCount)
	// 	case 4:
	// 		searchStartupByName(startups, startupCount)
	// 	case 5:
	// 		searchStartupByField(startups, startupCount)
	// 	case 6:
	// 		sortByFunding(&startups, startupCount)
	// 	case 7:
	// 		sortByYear(&startups, startupCount)
	// 	case 8:
	// 		reportByCategory(&startups, startupCount)
	// 	case 9:
	// 		deleteStartup(&startups, &startupCount)
	// 	default:
	// 		fmt.Println("Invalid option.")
	// 	}
	// 	showMenu()
	// 	fmt.Scan(&choice)
	// }
}

// func StrToInt(str string) (int, error) {
//     nonFractionalPart := strings.Split(str, ".")
//     return strconv.Atoi(nonFractionalPart[0])
// }

// func showMenu() {
// 	fmt.Println("=========================================")
// 	fmt.Println("1. Add Startup")
// 	fmt.Println("2. View Startups")
// 	fmt.Println("3. Add Team Member")
// 	fmt.Println("4. Search Startup by Name")
// 	fmt.Println("5. Search Startup by Field")
// 	fmt.Println("6. Sort by Funding (High to Low)")
// 	fmt.Println("7. Sort by Year (Old to New)")
// 	fmt.Println("8. Report by Category")
// 	fmt.Println("9. Delete Startup")
// 	fmt.Println("0. Exit")
// 	fmt.Print("Enter your choice: ")
// }

func addStartup(S *StartupList, count *int) {
	if *count >= MaxStartups {
		fmt.Println("Startup list is full.")
		return
	}

	var s Startup
	fmt.Print("Name: ")
	fmt.Scan(&s.Name)
	fmt.Print("Founded Year: ")
	fmt.Scan(&s.Founded)
	fmt.Print("Funding: ")
	fmt.Scan(&s.Funding)
	fmt.Print("Field: ")
	fmt.Scan(&s.Field)
	fmt.Print("Category: ")
	fmt.Scan(&s.Category)

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
		fmt.Printf("[%d] %s (%d) - $%.2f - %s - %s\n", i+1, S[i].Name, S[i].Founded, S[i].Funding, S[i].Field, S[i].Category)
		for j := 0; j < S[i].TeamSize; j++ {
			fmt.Printf("   -> %s (%s)\n", S[i].Team[j].Name, S[i].Team[j].Role)
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
	fmt.Scan(&member.Name)
	fmt.Print("Role: ")
	fmt.Scan(&member.Role)

	startup := &S[index-1]
	if startup.TeamSize >= MaxMembers {
		fmt.Println("Team is full.")
		return
	}
	startup.Team[startup.TeamSize] = member
	startup.TeamSize++
	fmt.Println("Team member added.")
}

func searchStartupByName(S StartupList, N int) {
	var name string
	fmt.Print("Enter name to search: ")
	fmt.Scan(&name)

	for i := 0; i < N; i++ {
		if S[i].Name == name {
			fmt.Printf("Found: %s (%d) - $%.2f - %s - %s\n", S[i].Name, S[i].Founded, S[i].Funding, S[i].Field, S[i].Category)
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
		if S[i].Field == field {
			fmt.Printf("Found: %s (%d) - $%.2f - %s - %s\n", S[i].Name, S[i].Founded, S[i].Funding, S[i].Field, S[i].Category)
			return
		}
	}
	fmt.Println("No startup found in that field.")
}

func sortByFunding(S *StartupList, N int) {
	var pass, idx, i int
	var temp Startup

	pass = 1
	for pass <= N-1 {
		idx = pass - 1
		i = pass
		for i < N {
			if S[idx].Funding < S[i].Funding {
				idx = i
			}
			i++
		}
		temp = S[pass-1]
		S[pass-1] = S[idx]
		S[idx] = temp
		pass++
	}
	fmt.Println("Startups sorted by funding (descending).")
}

func sortByYear(S *StartupList, N int) {
	var pass, j int
	var temp Startup

	for pass = 1; pass < N; pass++ {
		j = pass
		temp = S[pass]

		for j > 0 && S[j-1].Founded > temp.Founded {
			S[j] = S[j-1]
			j = j - 1
		}
		S[j] = temp
	}
	fmt.Println("Startups sorted by year (ascending).")
}

func reportByCategory(S *StartupList, N int) {
	if N == 0 {
		fmt.Println("No data available.")
		return
	}

	var categories [MaxStartups]string
	var counts [MaxStartups]int
	categoryCount := 0

	for i := 0; i < N; i++ {
		cat := S[i].Category
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

func deleteStartup(S *StartupList, N *int) {
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
	viewStartups(*S, *N)

	for i := index - 1; i < *N-1; i++ {
		S[i] = S[i+1]
	}
	*N--

	fmt.Println("After deletion:")
	viewStartups(*S, *N)
}
