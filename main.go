package main

import (
	"fmt"
	"strings"

	"github.com/orus-dev/osui"
	"github.com/orus-dev/osui/ui"
)

var students map[string]string = map[string]string{}
var grades []string = []string{"1", "2", "3", "4", "5", "6", "7", "8"}

func main() {
	screen := osui.NewScreen(App())
	screen.Run()
}

func App() *ui.DivComponent {
	return renderRoot()
}

func renderRoot() *ui.DivComponent {
	var root ui.Id[*ui.DivComponent]

	return root.Id(ui.Div(
		ui.WithPosition(0, 0, ui.Text("student managment system (terminal edition permium) made by fefek")),
		ui.WithPosition(0, 2, ui.Button("add student").Params(ui.ButtonParams{OnClick: func(bc *ui.ButtonComponent) bool { // add student button function | ADD
			var nameInput ui.Id[*ui.InputBoxComponent]

			root.Component.Components = []osui.Component{}

			root.Component.Components = append(root.Component.Components, ui.WithPosition(0, 0, ui.Text("student's name and grade")))
			root.Component.Components = append(root.Component.Components, ui.WithPosition(0, 2, nameInput.Id(ui.InputBox(20))))
			root.Component.Components = append(root.Component.Components, ui.WithPosition(0, 5, ui.Button("sumbit").Params(ui.ButtonParams{OnClick: func(bc *ui.ButtonComponent) bool { // add student sumbit function | SUMBIT
				name := nameInput.Component.InputData
				students[name] = "no grade"

				screen := osui.NewScreen(App())
				screen.Run()

				return false
			}})))
			root.Component.Components = append(root.Component.Components, ui.WithPosition(0, 8, ui.Button("back").Params(ui.ButtonParams{OnClick: func(bc *ui.ButtonComponent) bool { // add student exit function | EXIT
				screen := osui.NewScreen(App())
				screen.Run()

				return false
			}})))

			return false
		}})),
		ui.WithPosition(25, 2, ui.Button("delete student").Params(ui.ButtonParams{OnClick: func(bc *ui.ButtonComponent) bool { // delete student button function | DELETE
			var nameInput ui.Id[*ui.InputBoxComponent]

			root.Component.Components = []osui.Component{}

			root.Component.Components = append(root.Component.Components, ui.WithPosition(0, 0, ui.Text("student's name")))
			root.Component.Components = append(root.Component.Components, ui.WithPosition(0, 2, nameInput.Id(ui.InputBox(20))))
			root.Component.Components = append(root.Component.Components, ui.WithPosition(0, 5, ui.Button("sumbit").Params(ui.ButtonParams{OnClick: func(bc *ui.ButtonComponent) bool { // delete student sumbit function | SUMBIT
				name := nameInput.Component.InputData

				delete(students, name)

				screen := osui.NewScreen(App())
				screen.Run()

				return false
			}})))
			root.Component.Components = append(root.Component.Components, ui.WithPosition(0, 8, ui.Button("back").Params(ui.ButtonParams{OnClick: func(bc *ui.ButtonComponent) bool { // add student exit function | EXIT
				screen := osui.NewScreen(App())
				screen.Run()

				return false
			}})))

			return false
		}})),
		ui.WithPosition(50, 2, ui.Button("add grade tag").Params(ui.ButtonParams{OnClick: func(bc *ui.ButtonComponent) bool { // add grade tag | GRADE
			var selectedStudent ui.Id[*ui.MenuComponent]
			var selectedGrade ui.Id[*ui.MenuComponent]

			var student string
			var grade string

			root.Component.Components = []osui.Component{}

			root.Component.Components = append(root.Component.Components, ui.WithPosition(0, 0, ui.Text("select student")))
			root.Component.Components = append(root.Component.Components, ui.WithPosition(50, 0, ui.Button("add new grade").Params(ui.ButtonParams{OnClick: func(bc *ui.ButtonComponent) bool { // add grade tag add new grade | GRADE NEW
				var newGrade ui.Id[*ui.InputBoxComponent]

				root.Component.Components = []osui.Component{}

				root.Component.Components = append(root.Component.Components, ui.WithPosition(0, 0, ui.Text("new grade's name")))
				root.Component.Components = append(root.Component.Components, ui.WithPosition(0, 2, newGrade.Id(ui.InputBox(20))))
				root.Component.Components = append(root.Component.Components, ui.WithPosition(0, 5, ui.Button("sumbit").Params(ui.ButtonParams{OnClick: func(bc *ui.ButtonComponent) bool { // add grade tag add new grade sumbit | GRADE NEW SUMBIT
					new := newGrade.Component.InputData
					grades = append(grades, new)

					screen := osui.NewScreen(App())
					screen.Run()

					return false
				}})))
				root.Component.Components = append(root.Component.Components, ui.WithPosition(0, 8, ui.Button("back").Params(ui.ButtonParams{OnClick: func(bc *ui.ButtonComponent) bool { // add grade tag add new grade back | GRADE NEW BACK
					screen := osui.NewScreen(App())
					screen.Run()
					return false
				}})))

				return false
			}})))
			root.Component.Components = append(root.Component.Components, ui.WithPosition(75, 0, ui.Button("back").Params(ui.ButtonParams{OnClick: func(bc *ui.ButtonComponent) bool { // add grade tag back | GRADE BACK
				screen := osui.NewScreen(App())
				screen.Run()

				return false
			}})))
			root.Component.Components = append(root.Component.Components, ui.WithPosition(0, 3, selectedStudent.Id(ui.Menu(getArray(students)...)).Params(ui.MenuParams{OnSelected: func(mc *ui.MenuComponent, b bool) { // add grade tag student menu | STUDENT MENU
				student = selectedStudent.Component.Items[selectedStudent.Component.SelectedItem]
				parts := strings.Split(student, "|")
				student = strings.TrimSpace(parts[0])
			}})))
			root.Component.Components = append(root.Component.Components, ui.WithPosition(50, 3, selectedGrade.Id(ui.Menu(grades...)).Params(ui.MenuParams{OnSelected: func(mc *ui.MenuComponent, b bool) { // add grade tag student menu | STUDENT MENU
				grade = selectedGrade.Component.Items[selectedGrade.Component.SelectedItem]

				for key := range students {
					if key == student || students[key] == student {
						students[key] = grade
						break
					}
				}

				screen := osui.NewScreen(App())
				screen.Run()
			}})))

			return false
		}})),
		ui.WithPosition(0, 5, ui.Text("students")),
		ui.WithPosition(0, 8, ui.Menu(getArray(students)...)),
	))
}

func getArray(m map[string]string) []string {
	var items []string = []string{}

	for key, value := range m {
		items = append(items, fmt.Sprintf("%v | %v", key, value))
	}

	return items
}
