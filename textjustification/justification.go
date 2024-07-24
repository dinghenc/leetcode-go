package textjustification

import "strings"

func fullJustify(words []string, maxWidth int) []string {
	var ans []string
	beginWordIndex := 0
	for beginWordIndex < len(words) {
		line := ""
		endWordIndex := beginWordIndex
		lineLength := 0
		wordLength := 0
		for endWordIndex < len(words) && lineLength < maxWidth {
			lineLength += len(words[endWordIndex]) + 1
			wordLength += len(words[endWordIndex])
			endWordIndex++
		}
		if lineLength > maxWidth {
			endWordIndex -= 1
			wordLength -= len(words[endWordIndex])
			blanketCount := maxWidth - wordLength
			wordCount := endWordIndex - beginWordIndex
			if wordCount == 1 {
				line += words[beginWordIndex] + strings.Repeat(" ", blanketCount)
			} else {
				avgBlanketCount := blanketCount / (wordCount - 1)
				exceedBlanketCount := blanketCount - (wordCount-1)*avgBlanketCount
				for i := beginWordIndex; i < endWordIndex; i++ {
					if i == endWordIndex-1 {
						line += words[i]
						continue
					}
					line += words[i] + strings.Repeat(" ", avgBlanketCount)
					if exceedBlanketCount > 0 {
						line += " "
						exceedBlanketCount--
					}
				}
			}
		} else {
			endLine := strings.Join(words[beginWordIndex:endWordIndex], " ")
			line += endLine + strings.Repeat(" ", maxWidth-len(endLine))
		}
		beginWordIndex = endWordIndex
		ans = append(ans, line)
	}
	return ans
}
