package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Player struct {
	Cash           int
	Salary         int
	LivingExpenses int
	PassiveIncome  int
	Month          int
	GoalPassive    int
	EmergencyFund  int
}

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	player := Player{
		Cash:           2000,
		Salary:         3000,
		LivingExpenses: 2000,
		PassiveIncome:  0,
		Month:          1,
		GoalPassive:    5000,
		EmergencyFund:  1000,
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("=== Денежный поток: Мини-игра ===")
	fmt.Println("Цель: довести пассивный доход до 5000 в месяц.")
	fmt.Println("Каждый месяц выбирай действия: инвестировать, обучаться или копить.")

	for {
		fmt.Printf("\nМесяц %d\n", player.Month)
		monthIncome := player.Salary + player.PassiveIncome
		monthExpenses := player.LivingExpenses
		player.Cash += monthIncome - monthExpenses
		fmt.Printf("Доход: %d (зарплата %d + пассивный %d)\n", monthIncome, player.Salary, player.PassiveIncome)
		fmt.Printf("Расходы: %d\n", monthExpenses)
		fmt.Printf("Кэш: %d\n", player.Cash)

		if player.PassiveIncome >= player.GoalPassive {
			fmt.Println("\nПоздравляем! Ваш пассивный доход достиг цели.")
			break
		}

		event := randomEvent(&player)
		if event != "" {
			fmt.Println(event)
			if player.Cash < 0 {
				fmt.Println("Вы в долгах. Игра окончена.")
				break
			}
		}

		fmt.Println("\nВыберите действие:")
		fmt.Println("1) Инвестировать 1000 → шанс +300 пассивного")
		fmt.Println("2) Обучение 1500 → +500 к зарплате")
		fmt.Println("3) Копить 500 → +500 к резерву")
		fmt.Println("4) Завершить игру")
		choice := readInt(reader, "Ваш выбор: ")

		switch choice {
		case 1:
			invest(&player)
		case 2:
			train(&player)
		case 3:
			save(&player)
		case 4:
			fmt.Println("Игра завершена пользователем.")
			return
		default:
			fmt.Println("Неверный выбор, пропускаем ход.")
		}

		player.Month++
	}
}

func readInt(reader *bufio.Reader, prompt string) int {
	for {
		fmt.Print(prompt)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка чтения, попробуйте снова.")
			continue
		}
		input = strings.TrimSpace(input)
		value, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Введите число.")
			continue
		}
		return value
	}
}

func invest(player *Player) {
	if player.Cash < 1000 {
		fmt.Println("Недостаточно средств для инвестиций.")
		return
	}
	player.Cash -= 1000
	if rand.Intn(100) < 60 {
		player.PassiveIncome += 300
		fmt.Println("Инвестиция успешна! Пассивный доход +300.")
		return
	}
	fmt.Println("Инвестиция не сработала. Потрачено 1000.")
}

func train(player *Player) {
	if player.Cash < 1500 {
		fmt.Println("Недостаточно средств для обучения.")
		return
	}
	player.Cash -= 1500
	player.Salary += 500
	fmt.Println("Вы прошли обучение! Зарплата +500.")
}

func save(player *Player) {
	if player.Cash < 500 {
		fmt.Println("Недостаточно средств для накоплений.")
		return
	}
	player.Cash -= 500
	player.EmergencyFund += 500
	fmt.Printf("Резервный фонд: %d\n", player.EmergencyFund)
}

func randomEvent(player *Player) string {
	roll := rand.Intn(100)
	switch {
	case roll < 15:
		cost := 700
		player.Cash -= cost
		return fmt.Sprintf("Незапланированная поломка. Потрачено %d.", cost)
	case roll < 25:
		bonus := 800
		player.Cash += bonus
		return fmt.Sprintf("Премия на работе! Получено %d.", bonus)
	case roll < 30:
		cost := 1200
		player.Cash -= cost
		return fmt.Sprintf("Медицинские расходы. Потрачено %d.", cost)
	default:
		return ""
	}
}
