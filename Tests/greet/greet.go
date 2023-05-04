package greet

func ExtractWord(words []string) string {
    // Iterate over the words in the array
    for _, word := range words {
        // Check if the word is "extract"
        if word == "extract" {
            return word
        }
    }
    return ""
}