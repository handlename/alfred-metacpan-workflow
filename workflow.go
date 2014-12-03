package wf

import (
	"encoding/xml"
	"fmt"

	"github.com/handlename/go-metacpan"
)

// The ModulesXML defines XML struct of list of distributions.
type ModulesXML struct {
	XMLName xml.Name    `xml:"items"`
	Item    []ModuleXML `xml:"item"`
}

// The ModuleXML defines XML struct of list of distributions.
type ModuleXML struct {
	XMLName  xml.Name `xml:"item"`
	Arg      string   `xml:"arg,attr"`
	Title    string   `xml:"title"`
	Subtitle string   `xml:"subtitle"`
}

// SearchModule returns search distribution by query(q) and returns results as XML.
func SearchModule(q string) string {
	results, err := metacpan.SearchAutocomplete(q)

	if err != nil {
		return errorToXML(err)
	}

	xmlType := ModulesXML{
		Item: []ModuleXML{},
	}

	for _, result := range results {
		xmlType.Item = append(xmlType.Item, ModuleXML{
			Arg:      result.Fields.Documentation,
			Title:    result.Fields.Documentation,
			Subtitle: fmt.Sprintf("%s/%s", result.Fields.Author, result.Fields.Release),
		})
	}

	xmlBytes, err := xml.Marshal(xmlType)

	if err != nil {
		return errorToXML(err)
	}

	return xml.Header + string(xmlBytes)
}

// errorToXML convert error to XML.
func errorToXML(err error) string {
	return xml.Header + `
<items>
  <item arg="">
    <title>ERROR</title>
    <subtitle>Something wrong</subtitle>
  </item>
</items>`
}
