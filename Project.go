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

type CategoryReport struct {
	Category string
	Count    int
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

var startups StartupList
var startupCount = 0

func main() {

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

router.GET("/", func(c *gin.Context) {
    c.HTML(http.StatusOK, "index.html", gin.H{
        "title": "Home",
    })
})


// This for View your Startup Router
router.GET("/view", func(c *gin.Context) {
	viewStartups(c, startups, startupCount)
})

// This the Router of AddStartup
router.GET("/add", func(c *gin.Context) {
	c.HTML(http.StatusOK, "add.html", gin.H{
		"title": "Add Startup",
		"content": "add",
	})
})

// This for the AddStartup
router.POST("/add", func(c *gin.Context) {
    var s Startup
    fundingStr := c.PostForm("funding")
    foundedStr := c.PostForm("founded")

    funding, err1 := strconv.ParseFloat(fundingStr, 64)
    founded, err2 := strconv.Atoi(foundedStr)

    if err1 != nil || err2 != nil {
        c.String(http.StatusBadRequest, "Invalid input")
        return
    }

    s.Name = c.PostForm("name")
    s.Funding = funding
    s.Founded = founded
    s.Field = c.PostForm("field")
    s.Category = c.PostForm("category")

    addStartup(&startups, &startupCount, s)

	fmt.Println("Added:", s.Name)
	fmt.Println("Total startups now:", startupCount)

    c.Redirect(http.StatusFound, "/view")
})

// This for Add Member
router.POST("/add-member", func(c *gin.Context) {
    idxStr := c.PostForm("index")
    idx, err := strconv.Atoi(idxStr)
    if err != nil || idx < 0 || idx >= startupCount {
        c.String(http.StatusBadRequest, "Invalid startup index")
        return
    }

    member := TeamMember{
        Name: c.PostForm("member_name"),
        Role: c.PostForm("member_role"),
    }

    addTeamMember(&startups, idx, member)
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

    deleteStartup(&startups, &startupCount, idx)
    c.Redirect(http.StatusFound, "/view")
})

// This for Search UI page
router.GET("/search", func(c *gin.Context) {
    c.HTML(http.StatusOK, "view.html", gin.H{
        "title": "Search",
        "content": "search",
    })
})

// This was the Sort your startup slice by founded year ascending
router.POST("/sort-by-year", func(c *gin.Context) {
    sortByYear(&startups, startupCount)
    c.HTML(http.StatusOK, "view.html", gin.H{
        "title": "Sorted by Year",
        "startups": startups[:startupCount],
        "results": []Startup{}, // NOTE "THIS FOR AVOID CRASH, DONT DELETE IT"
    })
})

// This was the Sort your startup slice by funding descending
router.POST("/sort-by-funding", func(c *gin.Context) {
    sortByFunding(&startups, startupCount)
    c.HTML(http.StatusOK, "view.html", gin.H{
        "title": "Sorted by Funding",
        "startups": startups[:startupCount],
        "results": []Startup{}, 
    })
})

// This is for the Search by name
router.GET("/search-by-name", func(c *gin.Context) {
	searchStartupByName(c, startups, startupCount)
})

// This is for the Search by field from the view html 
router.GET("/search-by-field", func(c *gin.Context) {
	searchStartupByField(c, startups, startupCount)
})

// This is for the Report by category from the view html
router.POST("/report-category", func(c *gin.Context) {
	reportByCategory(c, &startups, startupCount)
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

func addStartup(S *StartupList, count *int, s Startup) {
	if *count >= MaxStartups {
		fmt.Println("Startup list is full.")
		return
	}
	S[*count] = s
	*count++
	fmt.Println("Startup added successfully.")
}

func viewStartups(c *gin.Context, S StartupList, N int) {
	var lines []string

	if N == 0 {
		lines = append(lines, "No startups to show.")
	} else {
		for i := 0; i < N; i++ {
			lines = append(lines, fmt.Sprintf("[%d] %s (%d) - $%.2f - %s - %s",
				i+1, S[i].Name, S[i].Founded, S[i].Funding, S[i].Field, S[i].Category))

			for j := 0; j < S[i].TeamSize; j++ {
				lines = append(lines, fmt.Sprintf("   -> %s (%s)", S[i].Team[j].Name, S[i].Team[j].Role))
			}
		}
	}

	c.HTML(http.StatusOK, "view.html", gin.H{
		"title": "View Startups (Console Style)",
		"rawOutput": lines,
		"startups": S[:N],
	})
}

func addTeamMember(S *StartupList, index int, member TeamMember) {
	if index < 0 || index >= MaxStartups {
		fmt.Println("Invalid startup index.")
		return
	}

	startup := &S[index]
	if startup.TeamSize >= MaxMembers {
		fmt.Println("Team is full.")
		return
	}

	startup.Team[startup.TeamSize] = member
	startup.TeamSize++
	fmt.Println("Team member added.")
}

// and this was the search by name using the sequential
func searchStartupByName(c *gin.Context, S StartupList, N int) {
	name := c.Query("name")
	var results []Startup

	for i := 0; i < N; i++ {
		if S[i].Name == name {
			results = append(results, S[i])
		}
	}

	c.HTML(http.StatusOK, "view.html", gin.H{
		"title": "Search by Name",
		"results": results,        
		"startups": S[:N],
	})
}

// this was the search by field using binary search algorithm but with a little bit modified it
func searchStartupByField(c *gin.Context, S StartupList, N int) {
	field := c.Query("field")
	var results []Startup

	var R, L, M int

	L = 0 
	R = N
	for i := 0; i < N; i++ {
		M = (R + L) / 2
		if S[M].Field == field {
			results = append(results, S[i])
		}else if S[M].Field > field  {
			R = M - 1
		}else if S[M].Field < field{
			L = M + 1
		}
	}

	c.HTML(http.StatusOK, "view.html", gin.H{
		"title": "Search by Field",
		"results": results,
		"startups": S[:N],
	})
}

// SortByFunding Function using selection sort (Descending)
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

// SortByYear Function was using (Ascending)
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

func reportByCategory(c *gin.Context, S *StartupList, N int) {
	if N == 0 {
		c.HTML(http.StatusOK, "view.html", gin.H{
			"title": "Category Report",
			"message": "No data available.",
			"report": []CategoryReport{}, // 👈 changed from map to struct slice
			"startups": (*S)[:N],
			"results": []Startup{},
		})
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

		if !found {
			categories[categoryCount] = cat
			counts[categoryCount] = 1
			categoryCount++
		}
	}

	var report []CategoryReport
	for i := 0; i < categoryCount; i++ {
		report = append(report, CategoryReport{
			Category: categories[i],
			Count:    counts[i],
		})
	}

	c.HTML(http.StatusOK, "view.html", gin.H{
		"title": "Category Report",
		"report": report, // 👈 struct slice
		"startups": (*S)[:N],
		"results": []Startup{},
	})
}

func deleteStartup(S *StartupList, N *int, index int) {
	if *N == 0 || index < 0 || index >= *N {
		fmt.Println("Invalid index.")
		return
	}

	for i := index; i < *N-1; i++ {
		S[i] = S[i+1]
	}
	*N--
}