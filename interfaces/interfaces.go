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

func (et ErrText) Display()    { fmt.Println("\033[91m" + et.Text) }
func (ct CyanText) Display()   { fmt.Println("\033[96m" + ct.Text) }
func (yt YellowText) Display() { fmt.Println("\033[93m" + yt.Text) }
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

func (ci CyanInput) Display()   { fmt.Println("\033[96m" + ci.Text) }
func (yi YellowInput) Display() { fmt.Println("\033[93m" + yi.Text) }
func (ci CyanInput) Read() string {
	fmt.Println("\033[96m" + ci.Text)
	var s string
	fmt.Scanln(&s)
	return s
}
func (yi YellowInput) Read() string {
	fmt.Println("\033[93m" + yi.Text)
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

func MakeSchemeManager(scheme string) (*SchemeManager, error) {

	var factory WidgetFactory

	switch scheme {
	case "Yellow":
		factory = CreateYellowFactory()
	case "Cyan":
		factory = CreateCyanFactory()
	default:
		return nil, fmt.Errorf("\"%s\" is an Invalid Scheme", scheme)
	}
	return &SchemeManager{Scheme: scheme, Factory: factory}, nil
}
