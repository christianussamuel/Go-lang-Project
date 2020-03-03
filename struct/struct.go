package main

import (
	"fmt"

)

type makanan struct {
	burger string
	harga  int
	rasa   string
}

type toko struct {
	namatoko   string
	alamattoko string
	status     bool
	makanan
}

var (
	s1       makanan
	s2       = toko{alamattoko: "tubagus ismail"}
	s3       toko
	datatoko = []struct {
		toko
		jumlahpenjual int
	}{
		{toko: toko{"wendy", "10", true, makanan{"a", 10, "A"}}, jumlahpenjual: 10},
	}
)

func save(s toko) {
	s.status = true
	fmt.Println("local save copy by value : ", s.status)
}

func save2(s *toko) {
	s.status = true
	fmt.Println("local save copy by ref : ", s.status)
}

func main() {

	s1.burger = "burger pedas"
	s3.makanan.rasa = "manis"

	fmt.Println("burger : ", s1.burger)
	fmt.Println("harga :", s1.harga)
	fmt.Println("alamat toko : ", s2.alamattoko)
	fmt.Println("rasanya : ", s3)
	fmt.Println("")

	fmt.Println("Status : ", s3.status)
	fmt.Println("")

	save(s3)

	fmt.Println("Status after: ", s3.status)
	fmt.Println("")

	save2(&s3)

	fmt.Println("Status after : ", s3.status)
	fmt.Println("")

	for _, toko1 := range datatoko {
		fmt.Println(toko1)
	}

}
