package runner

type Solution interface {
	Part1(input string) (string, error)
	Part2(input string) (string, error)
}

// Solution registry
type SolutionRegistry map[int]map[int]Solution

var Solutions = make(SolutionRegistry)

func Register(year, day int, solution Solution) {
	if _, exists := Solutions[year]; !exists {
		Solutions[year] = make(map[int]Solution)
	}
	Solutions[year][day] = solution
}
