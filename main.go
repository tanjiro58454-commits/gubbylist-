package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
)


// Global memory buffer to hold unique combinations securely in RAM
var uniqueWords = make(map[string]bool)
var totalCombinations int64 = 0

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("==================================================")
	fmt.Println("🚀        WELCOME TO GUBBY.LIST GENERATOR        🚀")
	fmt.Println("==================================================")

	// Step 1: Input Modes
	fmt.Print("\nDo you want to enter Target Information? (y/n): ")
	mode, _ := reader.ReadString('\n')
	mode = strings.TrimSpace(strings.ToLower(mode))

	var baseWords []string

	if mode == "y" {
		// CUPP Profiling Mode
		baseWords = gatherCuppData(reader)
	} else {
		// Completely Flexible Mode
		fmt.Println("\n[Flexible Mode Activated]")
		fmt.Print("Enter base characters/words separated by commas: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		for _, item := range strings.Split(input, ",") {
			trimmed := strings.TrimSpace(item)
			if trimmed != "" {
				baseWords = append(baseWords, trimmed)
			}
		}
	}

	if len(baseWords) == 0 {
		fmt.Println("❌ Error: No base inputs provided. Exiting.")
		return
	}

	// Channel to capture the Ctrl+C (Interrupt Signal)
	stopSignal := make(chan os.Signal, 1)
	signal.Notify(stopSignal, os.Interrupt, syscall.SIGTERM)

	fmt.Println("\n--------------------------------------------------")
	fmt.Println("🔥 INF-GEN MODE ACTIVATED!")
	fmt.Println("⏳ Generating infinite unique combinations in RAM...")
	fmt.Println("🛑 PRESS [Ctrl + C] TO STOP GENERATION AND GET LINK.")
	fmt.Println("--------------------------------------------------")

	// Channel to smoothly break out of the infinite generation loop
	done := make(chan bool)

	// Step 2: Run Generator in a background routine
	go func() {
		depth := 1
		for {
			select {
			case <-done:
				return
			default:
				// Keeps increasing complexity depth over time dynamically
				generateInfiniteCombinations(baseWords, "", depth, done)
				depth++
			}
		}
	}()

	// Wait explicitly for the user to press Ctrl + C
	<-stopSignal
	close(done) // Signals the generator loop to stop safely

	// Step 3: Serve the data over HTTP straight from memory
	totalCombinations = int64(len(uniqueWords))
	port := ":9999"
	endpoint := "/gubby.list"
	serverURL := fmt.Sprintf("http://localhost%s%s", port, endpoint)

	fmt.Println("\n\n==================================================")
	fmt.Println("🛑 GENERATION STOPPED BY USER")
	fmt.Printf("📊 Total Unique Combinations Generated: %d\n", totalCombinations)
	fmt.Println("==================================================")
	fmt.Println("🌐 [LIVE STORAGE-LESS LINK GENERATED]")
	fmt.Printf("🔗 Wordlist Endpoint Link: %s\n", serverURL)
	fmt.Println("ℹ️  Copy this link directly into your other security tools.")
	fmt.Println("🛑 Press Ctrl+C once more in the terminal to close the server.")
	fmt.Println("==================================================")

	// HTTP Handler to serve data directly from memory
	http.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		writer := bufio.NewWriter(w)
		for word := range uniqueWords {
			_, _ = writer.WriteString(word + "\n")
		}
		writer.Flush()
	})

	// Start the local micro-server
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Printf("❌ Server Error: %v\n", err)
	}
}

// CUPP data gathering flow
func gatherCuppData(reader *bufio.Reader) []string {
	var words []string
	add := func(label string) {
		fmt.Printf("%s: ", label)
		val, _ := reader.ReadString('\n')
		val = strings.TrimSpace(val)
		if val != "" {
			words = append(words, val)
			words = append(words, strings.ToLower(val))
			words = append(words, strings.ToUpper(val))
		}
	}

	fmt.Println("\n--- Enter Target Details (Press Enter to skip) ---")
	add("First Name")
	add("Last Name")
	add("Nickname")
	add("Birth Day (DD)")
	add("Birth Month (MM)")
	add("Birth Year (YYYY)")
	add("Partner's Name")
	add("Pet's Name")
	add("Company Name")

	// Standard structural components used in typical combinations
	words = append(words, []string{"123", "!", "@", "2026", "2025", "12345"}...)
	return words
}

// Core mathematical generator loop with live deduplication checking
func generateInfiniteCombinations(elements []string, current string, depth int, done chan bool) {
	select {
	case <-done:
		return
	default:
		if current != "" {
			// Deduplication check: Map handles this natively in O(1) time complexity
			uniqueWords[current] = true
		}

		if depth == 0 {
			return
		}

		for _, el := range elements {
			generateInfiniteCombinations(elements, current+el, depth-1, done)
		}
	}
}

