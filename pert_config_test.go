package graphman

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

func ExampleFromPertConfig() {
	yamlConfig := `
actions:
  - id: "1"
    title: "Prepare foundation"
    estimate: [4, 6, 10]
  - id: "2"
    title: "Make & position door frames"
    estimate: [2, 4, 7]
  - id: "3"
    title: "Lay drains & floor base"
    estimate: [7, 9, 12]
  - id: "4"
    title: "Install service & settings"
    estimate: [2, 4, 5]
    depends_on: ["5"]
  - id: "5"
    title: "Erect walls"
    estimate: [7, 10, 15]
    depends_on: ["1", "2"]
  - id: "6"
    title: "Plaster ceilings"
    estimate: [1, 2, 4]
    depends_on: ["4", "7"]
  - id: "7"
    title: "Erect roof"
    estimate: [4, 6, 8]
    depends_on: ["5"]
  - id: "8"
    title: "Install door & windows"
    estimate: [7, 9, 11]
    depends_on: ["7"]
  - id: "9"
    title: "Fit gutters & pipes"
    estimate: [1, 2, 3]
    depends_on: ["3", "6"]
  - id: "10"
    title: "Paint outside"
    estimate: [1, 2, 3]
    depends_on: ["8", "9"]
`
	var config PertConfig
	if err := yaml.Unmarshal([]byte(yamlConfig), &config); err != nil {
		panic(err)
	}

	graph := FromPertConfig(config)
	fmt.Println(graph)
	// Output:
	// {(Start,pre_5)[[pert:To=4,Tm=6,Tp=10,Te=6.33,σe=1,Ve=1,title:Prepare foundation]],(Start,pre_5)[[pert:To=2,Tm=4,Tp=7,Te=4.17,σe=0.83,Ve=0.69,title:Make & position door frames]],(Start,pre_9)[[pert:To=7,Tm=9,Tp=12,Te=9.17,σe=0.83,Ve=0.69,title:Lay drains & floor base]],(post_5,pre_6)[[pert:To=2,Tm=4,Tp=5,Te=3.83,σe=0.5,Ve=0.25,title:Install service & settings]],(pre_5,post_5)[[pert:To=7,Tm=10,Tp=15,Te=10.33,σe=1.33,Ve=1.78,title:Erect walls]],(pre_6,pre_9)[[pert:To=1,Tm=2,Tp=4,Te=2.17,σe=0.5,Ve=0.25,title:Plaster ceilings]],(post_7,pre_6)[[]],(post_5,post_7)[[pert:To=4,Tm=6,Tp=8,Te=6,σe=0.67,Ve=0.44,title:Erect roof]],(post_7,pre_10)[[pert:To=7,Tm=9,Tp=11,Te=9,σe=0.67,Ve=0.44,title:Install door & windows]],(pre_9,pre_10)[[pert:To=1,Tm=2,Tp=3,Te=2,σe=0.33,Ve=0.11,title:Fit gutters & pipes]],(pre_10,Finish)[[pert:To=1,Tm=2,Tp=3,Te=2,σe=0.33,Ve=0.11,title:Paint outside]],Finish}
}
