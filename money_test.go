package money

import (
	"fmt"
	"image"
	"testing"
)

func TestGridItems(t *testing.T) {
	terminal := image.Rect(0, 0, 120, 40)
	rootGridItemsType := []bool{true, true}                 // две строки
	rootGridItemsRatio := []float64{1.0 / 10, 1.0 - 1.0/10} // две строки
	levelOneGridItems := [][]bool{{}, {false, false}}

	fmt.Printf(
		"Terminal: Xmin: %d Ymin: %d Xmax: %d Ymax: %d\n",
		terminal.Min.X,
		terminal.Min.Y,
		terminal.Max.X,
		terminal.Max.Y)
	fmt.Printf("Terminal DD: DX = %d, DY = %d\n", terminal.Dx(), terminal.Dy())

	for idx, val := range rootGridItemsType {
		itemSize := GridItemSizeHelper(terminal, rootGridItemsRatio[idx], val)
		if val {
			fmt.Printf(" Root Item %d is row\n", idx)
		} else {
			fmt.Printf(" Root Item %d is col\n", idx)
		}

		fmt.Printf(
			"About Root Item %d: Xmin: %d Ymin: %d Xmax: %d Ymax: %d DX = %d DY = %d\n",
			idx,
			itemSize.Min.X,
			itemSize.Min.Y,
			itemSize.Max.X,
			itemSize.Max.Y,
			itemSize.Dx(),
			itemSize.Dy())
		levelOneItems := levelOneGridItems[idx]
		if len(levelOneItems) == 0 {
			continue
		}
		for oneIdx, oneVal := range levelOneItems {
			oneItemSize := GridItemSizeHelper(itemSize, 1.0/2, oneVal)
			if val {
				fmt.Printf(" One Item %d is row\n", oneIdx)
			} else {
				fmt.Printf(" One Item %d is col\n", oneIdx)
			}
			fmt.Printf(
				"About Child of Root %d OneItem %d: Xmin: %d Ymin: %d Xmax: %d Ymax: %d DX = %d DY = %d\n",
				idx,
				oneIdx,
				oneItemSize.Min.X,
				oneItemSize.Min.Y,
				oneItemSize.Max.X,
				oneItemSize.Max.Y,
				oneItemSize.Dx(),
				oneItemSize.Dy())
		}
	}
}

func GridItemSizeHelper(parent image.Rectangle, ratio float64, isRow bool) image.Rectangle {
	width := float64(parent.Dx()) + 1
	height := float64(parent.Dy()) + 1
	wRatio := 1.0
	hRatio := 1.0
	if isRow {
		hRatio = ratio
	} else {
		wRatio = ratio
	}
	//x := int(width*ratio) + parent.Min.X
	//y := int(height*ratio) + parent.Min.Y
	x := parent.Min.X
	y := parent.Min.Y
	w := int(width * wRatio)
	h := int(height * hRatio)

	if x+w > parent.Dx() {
		w--
	}
	if y+h > parent.Dy() {
		h--
	}

	return image.Rect(x, y, x+w, y+h)
}
