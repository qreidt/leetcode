package main

import (
	"fmt"
	"math"
)

func longestPalindrome(s string) string {
	var longest []int32

	for i, char1 := range s {
		current := []int32{char1}
		if len(current) > len(longest) && isPalindrome(current) {
			longest = current
		}

		for _, char2 := range s[i+1:] {
			current = append(current, char2)
			if len(current) > len(longest) && isPalindrome(current) {
				longest = current
			}
		}
	}

	// build back to string
	result := ""
	for _, char := range longest {
		result += string(char)
	}

	return result
}

func isPalindrome(s []int32) bool {
	half := int(math.Floor(float64(len(s)) / 2))
	for i := 0; i < half; i++ {
		v1 := s[i]
		v2 := s[len(s)-1-i]

		if v1 != v2 {
			return false
		}
	}

	return true
}

//goland:noinspection SpellCheckingInspection
func main() {
	// Example 1
	fmt.Println("Output:", longestPalindrome("babad")) // Expected: "bab" or "aba"

	// Example 2
	fmt.Println("Output:", longestPalindrome("cbbd")) // Expected: "bb"

	// Example 3
	fmt.Println("Output:", longestPalindrome("a")) // Expected: "a"

	// Example 4
	fmt.Println("Output:", longestPalindrome("ac")) // Expected: "a"

	// Example 5
	fmt.Println("Output:", longestPalindrome("bb")) // Expected: "bb"

	// Example 6
	fmt.Println("Output:", longestPalindrome("ccc")) // Expected: "ccc"

	// Example 7
	fmt.Println("Output:", longestPalindrome("yfikrcvmuegdciuqahlsjesplljlswxaejgdzhubzqkiroxyhtjvazcwcnsvdzjiainmiyobyfclyugttaswlntwukkfbebcdaxdpaxwqenkxxphxdcgrnpruoaetvunwyskswvvmjmltncsdukwzlpfodhgxkjvzppwpvmqlfbojgbdiryleskemhjfoxxzjqihcykpgzhaugwwbqtddjzpmrgdncgzsttqenmbnnavfjkiennwxtguywoaiuungqpyfcffzmljfianapawiayywuvazrnxouvndzqbmmyntkkdyykgodjbeojtpnsyhfrltuazgznddaaibupephvgrcjpzvjttmhtnydwvrpgijselaukwrcosxpcbptebalkheymuyblffahvbszotmutmmqhlgoskuoejvavlprvgyozpylsnqhqrnqpabgbwzwxyibpmsauxcfnbtwwbosksuzqzmobijytxxtyjibomzqzusksobwwtbnfcxuasmpbiyxwzwbgbapqnrqhqnslypzoygvrplvavjeouksoglhqmmtumtozsbvhafflbyumyehklabetpbcpxsocrwkualesjigprvwdynthmttjvzpjcrgvhpepubiaaddnzgzautlrfhysnptjoebjdogkyydkktnymmbqzdnvuoxnrzavuwyyaiwapanaifjlmzffcfypqgnuuiaowyugtxwnneikjfvannbmneqttszgcndgrmpzjddtqbwwguahzgpkychiqjzxxofjhmekselyridbgjobflqmvpwppzvjkxghdofplzwkudscntlmjmvvwsksywnuvteaourpnrgcdxhpxxkneqwxapdxadcbebfkkuwtnlwsattguylcfyboyimniaijzdvsncwczavjthyxorikqzbuhzdgjeaxwsljllpsejslhaquicdgeumvcrkify")) // Expected: "ccc"
}
