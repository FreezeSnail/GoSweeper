package mines

import (
	"github.com/nsf/termbox-go"
)

// Colors
const (
	boardColor        = termbox.ColorBlack
	instructionsColor = termbox.ColorYellow
)

var backgroundColor = termbox.RGBToAttribute(128, 128, 128)

var pieceColors = []termbox.Attribute{
	termbox.ColorBlack,
	termbox.ColorRed,
	termbox.ColorGreen,
	termbox.ColorYellow,
	termbox.ColorBlue,
	termbox.ColorMagenta,
	termbox.ColorCyan,
	termbox.ColorWhite,
	termbox.ColorWhite,
	termbox.RGBToAttribute(250, 128, 114), //salmon
	termbox.RGBToAttribute(153, 255, 255), //cyan
	termbox.RGBToAttribute(128, 128, 128), //grey
}

// Layout
const (
	defaultMarginWidth  = 2
	defaultMarginHeight = 1
	titleStartX         = defaultMarginWidth
	titleStartY         = defaultMarginHeight
	titleHeight         = 1
	titleEndY           = titleStartY + titleHeight
	boardStartX         = defaultMarginWidth
	boardStartY         = titleEndY + defaultMarginHeight
	boardWidth          = 10
	boardHeight         = 16
	cellWidth           = 2
	boardEndX           = boardStartX + boardWidth*cellWidth
	boardEndY           = boardStartY + boardHeight
	instructionsStartX  = boardEndX + defaultMarginWidth
	instructionsStartY  = boardStartY
)

func Render(g *Game) {
	termbox.Clear(backgroundColor, backgroundColor)
	for y := 0; y < g.b.h; y++ {
		for x := 0; x < g.b.w; x++ {
			cellValue := g.b.board[y][x]
			var cellColor termbox.Attribute
			absCellValue := cellValue.denom
			if absCellValue == 9 {
				cellColor = pieceColors[9]
			} else {
				cellColor = pieceColors[10]

			}
			for i := 0; i < g.b.w; i++ {
				for j := 0; j < g.b.w; j++ {
					termbox.SetCell(boardStartX+g.b.w*x+i, boardStartY*y+j, ' ', backgroundColor, cellColor)
				}
			}
		}
	}
	termbox.Flush()
}
