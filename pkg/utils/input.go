package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// GetUserInput запрашивает ввод у пользователя
func GetUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// WaitForEnter ожидает нажатия Enter
func WaitForEnter() {
	fmt.Print("\nНажмите Enter для выхода...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}