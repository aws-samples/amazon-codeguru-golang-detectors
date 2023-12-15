package main

import "fmt"

// {fact rule=best-practices@v1.0 defects=1}
func hardcodedEqTrueOrFalseNoncompliant(num int){
	// Noncompliant: `if` statement with hardcoded `true` condition.
    if num == 5 && true {
        fmt.Println("Hello")
    }
}
// {/fact}

// {fact rule=best-practices@v1.0 defects=0}
func hardcodedEqTrueOrFalseCompliant(num int){
    // Compliant: `if` statement without hardcoded condition.
	if num == 5 {
		fmt.Println("Hello")
	}
}
// {/fact}
