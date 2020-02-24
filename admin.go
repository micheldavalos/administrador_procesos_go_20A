package main

import (
	"fmt"
	"time"
	// "time"
)

func main() {
	id := uint(0)
	op := 1
	flag := false
	c := make(chan uint)

	for op != 0 {
		fmt.Println("1) Agregar proceso")
		fmt.Println("2) Mostrar proceso")
		fmt.Println("3) Eliminar proceso")
		fmt.Println("0) Salir")
		fmt.Scanln(&op)

		switch op {
		case 1:
			go func(id_ext uint, ch chan uint) {
				id := id_ext
				i := 0
				for {

					select {
					case msg := <-ch:
						if msg == id {
							fmt.Println(id)
							return
						} else {
							ch <- msg
						}
					default:
						if flag {
							fmt.Println(id, ":", i)
						}
						i++
						time.Sleep(time.Millisecond * 500)
					}
				}
			}(id, c)
			id++
		case 2:
			flag = !flag
		case 3:
			var id_del uint
			fmt.Scanln(&id_del)
			c <- id_del
		case 0:
			op = 0
		}
	}

}
