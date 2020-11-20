package colour

var (
	//Red colour red
	Red = Colour{"red"}
	//Blue colour blue
	Blue = Colour{"blue"}
	//Green colour green
	Green = Colour{"green"}
	//White colour white
	White = Colour{"white"}
	//Black colour black
	Black = Colour{"black"}
)

// Colour Models a colour
type Colour struct {
	Name string `json:"name"`
}
