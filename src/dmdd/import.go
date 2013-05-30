//<AMPS><AMP><APID>526311000001106</APID><VPID>318248001</VPID><NM>Verapamil 160mg tablets</NM><DESC>Verapamil 160mg tablets (A A H Pharmaceuticals Ltd)</DESC><SUPPCD>3144701000001104</SUPPCD><LIC_AUTHCD>0001</LIC_AUTHCD><AVAIL_RESTRICTCD>0001</AVAIL_RESTRICTCD></AMP><AMP><APID>390711000001102</APID><VPID>318248001</VPID><NM>Verapamil 160mg tablets</NM><DESC>Verapamil 160mg tablets (Actavis UK Ltd)</DESC><SUPPCD>3875201000001104</SUPPCD><LIC_AUTHCD>0001</LIC_AUTHCD><AVAIL_RESTRICTCD>0001</AVAIL_RESTRICTCD></AMP><AMP><APID>108111000001106</APID><VPID>319996000</VPID><NM>Zocor 10mg tablets</NM><DESC>Zocor 10mg tablets (Merck Sharp &amp; Dohme Ltd)</DESC><SUPPCD>3146901000001106</SUPPCD><LIC_AUTHCD>0001</LIC_AUTHCD><CSM>0001</CSM><AVAIL_RESTRICTCD>0001</AVAIL_RESTRICTCD></AMP><AMP><APID>776811000001104</APID><VPID>319997009</VPID><NM>Zocor 20mg tablets</NM><DESC>Zocor 20mg tablets (Merck Sharp &amp; Dohme Ltd)</DESC><SUPPCD>3146901000001106</SUPPCD><LIC_AUTHCD>0001</LIC_AUTHCD><CSM>0001</CSM><AVAIL_RESTRICTCD>0001</AVAIL_RESTRICTCD></AMP><AMP><APID>859611000001107</APID><VPID>320000009</VPID><NM>Zocor 40mg tablets</NM><DESC>Zocor 40mg tablets (Merck Sharp &amp; Dohme Ltd)</DESC><SUPPCD>3146901000001106</SUPPCD><LIC_AUTHCD>0001</LIC_AUTHCD><CSM>0001</CSM><AVAIL_RESTRICTCD>0001</AVAIL_RESTRICTCD></AMP><AMP><APID>422011000001104</APID><VPID>318421004</VPID><NM>Atenolol 100mg tablets</NM><DESC>Atenolol 100mg tablets (A A H Pharmaceuticals Ltd)</DESC><SUPPCD>3144701000001104</SUPPCD><LIC_AUTHCD>0001</LIC_AUTHCD><AVAIL_RESTRICTCD>0001</AVAIL_RESTRICTCD></AMP><AMP><APID>275811000001103</APID><VPID>318421004</VPID><NM>Atenolol 100mg tablets</NM><DESC>Atenolol 100mg tablets (Actavis UK Ltd)</DESC><SUPPCD>3875201000001104</SUPPCD><LIC_AUTHCD>0001</LIC_AUTHCD><AVAIL_RESTRICTCD>0001</AVAIL_RESTRICTCD></AMP>
package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

func test_xml_parsing() {
	xmlFile, err := os.Open("/Users/ross/Desktop/Data/DMD/dmd/f_amp2_3160513.xml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	decoder := xml.NewDecoder(xmlFile)

	for {
		t, _ := decoder.Token()
		if t == nil {
			break
		}

		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "AMP" {
				var p AMP
				decoder.DecodeElement(&p, &se)
			}
		}
	}

}
