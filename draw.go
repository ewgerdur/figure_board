package main

import (
	"github.com/fogleman/gg"
	//"fmt"
	//"github.com/ewgerdur/figure_board.git"
)

type Сircle struct {
	radius float64 // радиус
	count_of_circle int //кол-во кругов
}

type Rectangle struct {
	length float64 //длина
	width float64 //ширина
	count_of_rectangle int //кол-во квадратов
}

func ft_draw_rectangle(x int, y int, x_rec float64, y_rec float64, count int, dc *gg.Context) *gg.Context {
	var new_j float64 = y_rec
	var j float64 = 1
	var n int = (x / (int(x_rec))) // количество в строке
	var rasrad int = count / n // 16 / n
	//fmt.Println(count / n )
	for i := count; i > 0; i-- {
		for k := rasrad; k >= 0; k-- {
			j = float64(k)
			new_j = x_rec + j * (x_rec)
			if k == rasrad  && i > n * k {
				dc.DrawRectangle(x_rec + float64(count)*x_rec - float64(i)*x_rec, new_j, x_rec, y_rec)
				
				if i == count - count%n + 1 && k == rasrad {
					k = rasrad - 1
					i = count - count%n + 1
				}
			} else if k < rasrad && k > 0 {
				dc.DrawRectangle(x_rec + float64(count - (rasrad - k) * n - count%n)*x_rec - float64(i)*x_rec, new_j, x_rec, y_rec)
				//fmt.Println(r + float64(count - (rasrad - k) * n - count%n)*2*r - float64(i)*2*r)
			} else if k == 0 {
				dc.DrawRectangle(x_rec + float64(count - count%n)*x_rec - float64(i)*x_rec, new_j, x_rec, y_rec)
				//fmt.Println(r + float64(count  - count%n)*2*r - float64(i)*2*r)
			}
		}
	}
	return dc
}

func ft_draw_circle(x int, y int, r float64, count int, dc *gg.Context) *gg.Context {

	var new_j float64 = r
	var j float64 = 1
	var n int = (x / (2 * int(r))) // количество в строке*/
	var rasrad int = count / n // 16 / n
	//fmt.Println(count / n )
	for i := count; i > 0; i-- {
		for k := rasrad; k >= 0; k-- {
			j = float64(k)
			new_j = r + j * (2 * r)
			if k == rasrad  && i > n * k {
				dc.DrawCircle(r + float64(count)*2*r - float64(i)*2*r, new_j, r)
				
				if i == count - count%n + 1 && k == rasrad {
					k = rasrad - 1
					i = count - count%n + 1
				}
			} else if k < rasrad && k > 0 {
				dc.DrawCircle(r + float64(count - (rasrad - k) * n - count%n)*2*r - float64(i)*2*r, new_j, r)
				//fmt.Println(r + float64(count - (rasrad - k) * n - count%n)*2*r - float64(i)*2*r)
			} else if k == 0 {
				dc.DrawCircle(r + float64(count - count%n)*2*r - float64(i)*2*r, new_j, r)
				//fmt.Println(r + float64(count  - count%n)*2*r - float64(i)*2*r)
			}
		}
	}
	return dc
}



func main() {

	var x int = 500 // длина доски
	var y int = 500 // ширина доски

	//var first_circle = Сircle{radius: 50, count_of_circle: 4}
	var first_rectangle = Rectangle{length: 50, width: 50, count_of_rectangle: 19}
	dc := gg.NewContext(x, y)

	//dc = ft_draw_circle(x, y, first_circle.radius, first_circle.count_of_circle, dc)
	dc = ft_draw_rectangle(x, y, first_rectangle.length, first_rectangle.width, first_rectangle.count_of_rectangle ,dc)
    dc.SetRGB(2, 50, 70)//установка цвета
    dc.Fill()//закраска
    dc.SavePNG("out.png")//сохранить в пнг
}