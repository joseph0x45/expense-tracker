package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type flow struct {
	ID        string `json:"id" db:"id"`
	FlowType  string `json:"flow_type" db:"flow_type"`
	CreatedAt string `json:"created_at" db:"created_at"`
	Amount    int    `json:"amount"`
	Method    string `json:"method" db:"method"`
	Planned   string `json:"planned" db:"planned"`
	Purpose   string `json:"purpose" db:"purpose"`
}

func registerFlowInteractive(newFlow *flow) {
	var err error
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("> Enter flow type: ")
	flowType, err := reader.ReadString('\n')
	trimLineFeed(&flowType)
	if err != nil {
		handleReaderError("flow type", err)
	}
	if flowType != "in" && flowType != "out" {
		fmt.Println("Flow type must be 'in' or 'out'")
		os.Exit(1)
	}
	fmt.Print("> Enter amount: ")
	amountStr, err := reader.ReadString('\n')
	if err != nil {
		handleReaderError("amount", err)
	}
	trimLineFeed(&amountStr)
	amount, err := strconv.Atoi(amountStr)
	if err != nil || amount == 0 {
		fmt.Println("Amount must be a positive integer than 0")
		os.Exit(1)
	}
	fmt.Print("> Enter method: ")
	method, err := reader.ReadString('\n')
	if err != nil {
		handleReaderError("method", err)
	}
	trimLineFeed(&method)
	if method != "cash" && method != "mobile money" && method != "bank" {
		fmt.Println("Method must be 'cash', 'mobile money' or 'bank'")
		os.Exit(1)
	}
	fmt.Print("> Enter purpose: ")
	purpose, err := reader.ReadString('\n')
	if err != nil {
		handleReaderError("purpose", err)
	}
	trimLineFeed(&purpose)
	if purpose == "" {
		fmt.Println("Can't have empty purpose'")
		os.Exit(1)
	}
	fmt.Print("> Was this planned? yes/no ")
	planned, err := reader.ReadString('\n')
	if err != nil {
		handleReaderError("planned", err)
	}
	trimLineFeed(&planned)
	planned = strings.ToLower(planned)
	if planned != "yes" && planned != "no" {
		fmt.Println("Please respond with either 'yes' or 'no'")
		os.Exit(1)
	}

	newFlow.FlowType = flowType
	newFlow.CreatedAt = getFormattedDateTime()
	newFlow.Amount = amount
	newFlow.Method = method
	newFlow.Purpose = purpose
	newFlow.Planned = planned
	printFlow(newFlow)
	fmt.Print("Do you want to save this flow? y/N ")
	save, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error while reading your response:", err.Error())
		os.Exit(1)
	}
	trimLineFeed(&save)
	if save != "y" {
		fmt.Println("Aborting...")
		os.Exit(1)
	}
}

func registerFlow(newFlow *flow) {
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Invalid Usage")
		return
	}
	initDB()
	newFlowCmd := flag.NewFlagSet("new", flag.ExitOnError)
	newFlowInteractive := newFlowCmd.Bool("i", false, "Interactive mode")
	cmd := os.Args[1]
	switch cmd {
	case "new":
		newFlowCmd.Parse(os.Args[2:])
		newFlow := &flow{}
		if *newFlowInteractive {
			registerFlowInteractive(newFlow)
		} else {
			registerFlow(newFlow)
		}
		saveFlow(newFlow)
	}
}
