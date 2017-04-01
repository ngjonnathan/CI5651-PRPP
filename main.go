package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Global variables
var mejorSol []*Edge
var beneficioDisponible int
var solParcial []*Edge

func main() {

	//file, _ := os.Open("./instanciasPRPP/CHRISTOFIDES/P14NoRPP")
	//file, _ := os.Open("./instanciasPRPP/RANDOM/R9NoRPP")
	//file, _ := os.Open("./instanciasPRPP/DEGREE/D6NoRPP")
	file, _ := os.Open("./instanciasPRPP/GRID/G15NoRPP")
	lineScanner := bufio.NewScanner(file)
	line := 0
	g := NewGraph(1)
	maxBenefit, b, c := 0, 0, 0
	for lineScanner.Scan() {
		contents := strings.Fields(lineScanner.Text())
		if line == 0 {
			number, _ := strconv.ParseInt(contents[len(contents)-1], 0, 0)
			g = NewGraph(int(number))
		}
		if _, err := strconv.Atoi(contents[0]); err == nil {
			startNode, _ := strconv.ParseInt(contents[0], 0, 0)
			endNode, _ := strconv.ParseInt(contents[1], 0, 0)
			cost, _ := strconv.ParseInt(contents[2], 0, 0)
			benefit, _ := strconv.ParseInt(contents[3], 0, 0)
			g.AddEdge(int(startNode), int(endNode), int(cost), int(benefit))
			b = g.Benefit(int(startNode), int(endNode))
			c = g.Cost(int(startNode), int(endNode))
			if b-c >= 0 {
				maxBenefit = maxBenefit + b - c
			}
		}
		line++
	}
	var path []*Edge
	path = getCycleGRASP(g)             // Get Greedy initial Path
	path = removeNegativeCycle(g, path) // Remove Negative Cycle
	mejorSol = make([]*Edge, 0)         // Global variable bestPath
	// Copy path to new array
	for _, elem := range path {
		mejorSol = append(mejorSol, elem)
	}
	beneficioDisponible = maxBenefit // Global variable maxBenefit
	g.branchAndBound(1)
	fmt.Println("Ciclo Branch and bound: ", mejorSol)
	fmt.Println("Total: ", getPathBenefit(mejorSol))
}
