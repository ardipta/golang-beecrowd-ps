package main

import (
	"fmt"
)

func main() {
	var n1, n2, n3, n4 float64
	var exam float64
	var messages []string

	fmt.Scanln(&n1, &n2, &n3, &n4)
	fmt.Scanln(&exam)

	media := ((n1 * 2) + (n2 * 3) + (n3 * 4) + (n4 * 1)) / (2 + 3 + 4 + 1)
	mediaMsg := fmt.Sprintf("Media: %.1f", media)
	messages = append(messages, mediaMsg)

	if media < 5.0 {
		messages = append(messages, "Aluno reprovado.")
	} else if media >= 5.0 && media <= 6.9 {
		messages = append(messages, "Aluno em exame.")
	} else if media >= 7.0 {
		messages = append(messages, "Aluno aprovado.")
	}
	if exam > 0 {
		examMsg := fmt.Sprintf("Nota do exame: %.1f", exam)
		messages = append(messages, examMsg)
		mediaFinal := (media + exam) / 2
		mediaFinalMsg := fmt.Sprintf("Media final: %.1f", mediaFinal)
		if mediaFinal >= 5.0 {
			messages = append(messages, "Aluno aprovado.")
		} else if mediaFinal <= 4.9 {
			messages = append(messages, "Aluno reprovado.")
		}
		messages = append(messages, mediaFinalMsg)
	}

	for _, msg := range messages {
		fmt.Println(msg)
	}
}
