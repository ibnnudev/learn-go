package main

import "fmt"

type Siswa struct {
	ID     int
	Nama   string
	Gender string
	Kelas  int
}

func (s Siswa) DetailInfo() string {
	return fmt.Sprintf("ID: %d, Nama: %s, Gender: %s, Kelas: %d", s.ID, s.Nama, s.Gender, s.Kelas)
}

func (s *Siswa) UpdateInfo(nama string, gender string, kelas int) {
	s.Nama = nama
	s.Gender = gender
	s.Kelas = kelas
	fmt.Println("Info updated:", s.DetailInfo())
}

func main() {
	siswa1 := Siswa{
		ID:     1,
		Nama:   "Budi",
		Gender: "Laki-laki",
		Kelas:  10,
	}

	fmt.Println("Siswa 1 Info:", siswa1.DetailInfo())
	siswa1.UpdateInfo("Ani", "Wanita", 12)
	fmt.Println("Siswa 1 Info after update:", siswa1.DetailInfo())
}
