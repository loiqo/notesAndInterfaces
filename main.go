package main

import (
	"bufio"
	"fmt"
	ns "notes/structs"
	"os"
	"strconv"
	"strings"

	"github.com/k0kubun/pp"
)

var commands = map[string]string{
	"create":      "Создать заметку",
	"show {id}":   "Показать заметку",
	"showall":     "Показать все заметки",
	"delete {id}": "Удалить заметку",
	"exit":        "Выход",
}

func main() {
	var i int = 1
	pp.Println(commands)
	for i != 0 {
		scanner := bufio.NewScanner(os.Stdin)

		if scannerValid := scanner.Scan(); !scannerValid {
			pp.Println(commands)
			return
		}

		commarr := strings.Fields(scanner.Text())

		switch commarr[0] {

		case "create":
			fmt.Println("Создаем...")
			fmt.Print("Напишите заголовок: ")
			scanName := bufio.NewScanner(os.Stdin)
			if scannerValid := scanName.Scan(); !scannerValid {
				return
			}
			fmt.Print("Напишите текст: ")
			scanText := bufio.NewScanner(os.Stdin)
			if scannerValid := scanText.Scan(); !scannerValid {
				return
			}
			id := ns.NoteRepo.NewNote(ns.MyNoteRepo1, scanName.Text(), scanText.Text())
			fmt.Println("Id созданной заметки:", id)
		case "show":
			fmt.Println("Показываем заметку", commarr[1])
			id, err := strconv.Atoi(commarr[1])
			if err == nil {
				pp.Println(ns.NoteRepo.GetNote(ns.MyNoteRepo1, id))
			}
		case "showall":
			fmt.Println("Показываем все...")
			pp.Println(ns.NoteRepo.GetAll(ns.MyNoteRepo1))
		case "delete":
			fmt.Println("Удаляем заметку", commarr[1])
			id, err := strconv.Atoi(commarr[1])
			if err == nil {
				ns.NoteRepo.Delete(ns.MyNoteRepo1, id)
				fmt.Println("Удалили заметку", commarr[1])
			}
		case "exit":
			i = 0
		default:
			pp.Println(commands)
		}
	}
}
