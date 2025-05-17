package bmp

import (
	"errors"
	"fmt"
)

func PrintImage(color_array []RGB, width, height uint16) error {

	if len(color_array) != int(width*height) {
		return errors.New("data: the color_array dimensions do not match the informed dimensions")
	}

	for y := uint16(0); y < height - 1; y += 2 {
		rows := [2][]RGB{
			color_array[y*width : (y+1)*width],
			color_array[(y+1)*width : (y+2)*width],
		}

		for x := range width {
			fmt.Printf("\033[38;2;%d;%d;%d;48;2;%d;%d;%dm\u2580\033[m",
				rows[0][x].R, rows[0][x].G, rows[0][x].B,
				rows[1][x].R, rows[1][x].G, rows[1][x].B)
		}

		fmt.Println()
	}

	if height&1 == 1 {
		row := color_array[(height-1)*width:]

		for x := range width {
			fmt.Printf("\033[38;2;%d;%d;%dm\u2580\033[m", row[x].R, row[x].G, row[x].B)
		}

		fmt.Println()
	}

	return nil
}
