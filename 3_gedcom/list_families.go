package main

import (
	"fmt"
	"strings"
	"github.com/elliotchance/gedcom"
	"os"
    "github.com/pborman/getopt"
)

func main() {
    optName := getopt.StringLong("name", 'n', "Prakash", "Your name")
	optPadding := getopt.StringLong("pad", 'p', "*", "Indentation")
	file := getopt.StringLong("file", 'f', "/Users/sarnobat/sarnobat.git/gedcom/rohidekar.ged", "Gedcom File")
    optHelp := getopt.BoolLong("help", 0, "Help")
    getopt.Parse()
	args := getopt.Args()

    if *optHelp {
        getopt.Usage()
        os.Exit(0)
    }

    fmt.Println("File:\t" + *file)
    fmt.Println("Name:\t" + *optName)
    
    // Get the remaining positional parameters
	
	fmt.Println("positional args: ", args)


	document, err := gedcom.NewDocumentFromGEDCOMFile("/Users/sarnobat/sarnobat.git/gedcom/rohidekar.ged")
	if err != nil {
		panic(err)
	}
	
	var root gedcom.FamilyNode
	var rootIndividual gedcom.IndividualNode
	individualFamilyMap  := make(map[string]gedcom.FamilyNode)


	for _, individual := range document.Individuals() {
	  fmt.Printf("individual: %s\n", individual.String())
	}

	for _, family := range document.Families() {
		fmt.Printf("%s\n", family)
		fmt.Printf("unique identifier: %s\n\n", family.Husband().Individual().UniqueIdentifiers().Strings()[0])
		if (family.Husband().Individual().UniqueIdentifiers().Strings()[0] == "799db437-e0d2-44cc-a8f9-afda533cb5b7") {
			root = *family;
			rootIndividual = *family.Husband().Individual()
			fmt.Printf("Found Root\n")
		}
		
		individualFamilyMap[family.Husband().Individual().String()] = *family
	}
	
	if (&root == nil) {
		panic("")
	}
	
	append(root, *optPadding, 1, individualFamilyMap)	
	
	printIndividual(rootIndividual)
}

func printIndividual(rootIndividual gedcom.IndividualNode) {

	fmt.Printf("%s\n", rootIndividual.String())

}

func append(familyNode gedcom.FamilyNode, indentation string, level int, individualFamilyMap map[string]gedcom.FamilyNode) {

	if (&familyNode != nil) {
		padding := strings.Repeat(indentation, level)
		for _, child := range familyNode.Children() {

			fmt.Printf("%s %s\n", padding, child)
			if val, ok := individualFamilyMap[child.Individual().String()]; ok {
				append(val, indentation, level + 1, individualFamilyMap)
			}
		}
	}	
}

