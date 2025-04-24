package main

import "fmt"

type Task struct {
	Description string
	IsDone      bool
}

func main() {
	var task1 Task = Task{"install vscode", false}
	var task2 *Task = &task1

	task2.IsDone = true

	fmt.Println(task1)
	fmt.Println(task2)

	*task2 = Task{"update vscode", true}

	fmt.Println(task1)
	fmt.Println(task2)
}

/**
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	Description string
	IsDone      bool
}

var tasks []Task

func addTask(desription string) {
	newTask := Task{Description: desription, IsDone: false}
	tasks = append(tasks, newTask)
	fmt.Println("Tugas berhasil ditambahkan!")
}

func viewTasks() {
	if len(tasks) == 0 {
		fmt.Println("Belum ada tugas dalam daftar")
		return
	}

	fmt.Println("\nDaftar Tugas:")
	for i, task := range tasks {
		status := "[ ]"
		if task.IsDone {
			status = "[x]"
		}
		fmt.Printf("%d. %s %s\n", i+1, status, task.Description)
	}
	fmt.Println()
}

func markDone(taskNumber int) {
	if taskNumber < 1 || taskNumber > len(tasks) {
		fmt.Println("nomor tugas tidak valid")
		return
	}

	tasks[taskNumber-1].IsDone = true
	fmt.Printf("Tugas nomor %d berhasil ditandai selesai!\n", taskNumber)
}

func main() {
	for {
		fmt.Println("Menu:")
		fmt.Println("1. Tambah Tugas")
		fmt.Println("2. Lihat Daftar Tugas")
		fmt.Println("3. Tandai Selesai")
		fmt.Println("4. Keluar")
		fmt.Print("Pilih opsi: ")

		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			fmt.Print("Masukkan deskripsi tugas: ")
			var description string
			// Menggunakan fmt.Scanln untuk membaca seluruh baris input
			reader := bufio.NewReader(os.Stdin)
			description, _ = reader.ReadString('\n')
			description = strings.TrimSpace(description) // Hapus newline dan spasi berlebih
			addTask(description)
		case "2":
			viewTasks()
		case "3":
			fmt.Print("Masukkan nomor tugas yang ingin ditandai selesai: ")
			var taskNumberStr string
			fmt.Scanln(&taskNumberStr)
			taskNumber, err := strconv.Atoi(taskNumberStr)
			if err != nil {
				fmt.Println("Input tidak valid. Masukkan nomor.")
			} else {
				markDone(taskNumber)
			}
		case "4":
			fmt.Println("Terima kasih!")
			return
		default:
			fmt.Println("Opsi tidak valid. Silakan coba lagi.")
		}
		fmt.Println() // Baris kosong untuk pemisah
	}
}
**/
