package terminal

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

// The rows shown before the text system lets go of its fonts, as it does
// once they pile up with sizes, show again after, shaped anew: the glyphs
// of the fonts let go of are not drawn.
func TestRowsAfterFontsForgotten(t *testing.T) {
	loadLib(t)
	term, err := New(Options{Conn: newPipe()})
	if err != nil {
		t.Fatal(err)
	}
	defer term.Close()
	tt := ui.NewTester(func(c *ui.Context) { View(c, term).Fill() }, 400, 200)
	term.Feed([]byte("hello"))
	tt.Frame()
	for i := range 600 {
		ui.Shape("x", ui.Font{Family: "monospace", Size: 6 + float32(i)/64})
	}
	tt.Frame() // lets go of the fonts as it ends
	tt.Frame()
	v := term.v
	img, bg := tt.Image(), tt.Image().RGBAAt(v.ox+20*v.cellW, v.oy+v.cellH/2)
	ink := 0
	for y := v.oy; y < v.oy+v.cellH; y++ {
		for x := v.ox; x < v.ox+5*v.cellW; x++ {
			if img.RGBAAt(x, y) != bg {
				ink++
			}
		}
	}
	if ink == 0 {
		t.Error("the row drawn before the fonts were let go of is blank after")
	}
}
