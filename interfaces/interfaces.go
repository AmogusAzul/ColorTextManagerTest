package interfaces

import "fmt"

type WidgetFactory interface {
	CreateText(text string) Text
	CreateInput(text string) Input
}

type CyanWidgetFactory struct{}
type YellowWidgetFactory struct{}

func (CyanWidgetFactory) CreateText(text string) Text     { return *CreateCT(text) }
func (CyanWidgetFactory) CreateInput(text string) Input   { return *CreateCI(text) }
func (YellowWidgetFactory) CreateText(text string) Text   { return *CreateYT(text) }
func (YellowWidgetFactory) CreateInput(text string) Input { return *CreateYI(text) }
func CreateCyanFactory() WidgetFactory                    { return CyanWidgetFactory{} }
func CreateYellowFactory() WidgetFactory                  { return YellowWidgetFactory{} }

type Widget struct{ Text string }

// text
type Text interface {
	Display()
}
type ErrText struct{ Widget }
type CyanText struct{ Widget }
type YellowText struct{ Widget }

type Input interface {
	Display()
	Read() string
}
type CyanInput struct{ Widget }
type YellowInput struct{ Widget }

func DisplayText(text string, color_code string, leave_color ...bool) {
	fmt.Printf("\033[%s%s\n", color_code, text)
	if len(leave_color) == 0 || !leave_color[0] {
		fmt.Print("\033[0m")
	}
}

func (et ErrText) Display()    { DisplayText(et.Text, "91m") }
func (ct CyanText) Display()   { DisplayText(ct.Text, "96m") }
func (yt YellowText) Display() { DisplayText(yt.Text, "93m") }
func CreateET(err error) *ErrText {
	return &ErrText{
		Widget: Widget{Text: err.Error()}}
}
func CreateCT(text string) *CyanText {
	return &CyanText{
		Widget: Widget{Text: text}}
}
func CreateYT(text string) *YellowText {
	return &YellowText{
		Widget: Widget{Text: text}}
}

func (ci CyanInput) Display()   { DisplayText(ci.Text, "96m", true) }
func (yi YellowInput) Display() { DisplayText(yi.Text, "93m", true) }
func (ci CyanInput) Read() string {
	ci.Display()
	var s string
	fmt.Scanln(&s)
	return s
}
func (yi YellowInput) Read() string {
	yi.Display()
	var s string
	fmt.Scanln(&s)
	return s
}
func CreateCI(text string) *CyanInput {
	return &CyanInput{
		Widget: Widget{Text: text}}
}
func CreateYI(text string) *YellowInput {
	return &YellowInput{
		Widget: Widget{Text: text}}
}

type SchemeManager struct {
	Scheme  string
	Factory WidgetFactory
}

func MakeSchemeManager() *SchemeManager {
	var scheme string
	var err error = fmt.Errorf("")
	var factory WidgetFactory

	for err != nil {
		CreateET(err).Display()
		fmt.Println("Choose which color you want the text to be displayed: \033[96mCyan \033[0mor \033[93mYellow\033[0m")
		fmt.Scanln(&scheme)

		switch scheme {
		case "Yellow":
			factory, err = CreateYellowFactory(), nil
		case "Cyan":
			factory, err = CreateCyanFactory(), nil
		default:
			err = fmt.Errorf("\"%s\" is an Invalid Scheme", scheme)
		}
	}

	return &SchemeManager{Scheme: scheme, Factory: factory}
}
