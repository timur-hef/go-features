package clock

import (
	"encoding/xml"
	"fmt"
	"io"
	"time"
)

type Svg struct {
	XMLName xml.Name `xml:"svg"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	ViewBox string   `xml:"viewBox,attr"`
	Version string   `xml:"version,attr"`
	Circle  struct {
		Text  string `xml:",chardata"`
		Cx    string `xml:"cx,attr"`
		Cy    string `xml:"cy,attr"`
		R     string `xml:"r,attr"`
		Style string `xml:"style,attr"`
	} `xml:"circle"`
	Line []struct {
		Text  string `xml:",chardata"`
		X1    string `xml:"x1,attr"`
		Y1    string `xml:"y1,attr"`
		X2    string `xml:"x2,attr"`
		Y2    string `xml:"y2,attr"`
		Style string `xml:"style,attr"`
	} `xml:"line"`
}

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`

const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`

const svgEnd = `</svg>`

func DrawHand(t time.Time) string {
	analogClock := AnalogClock(t)
	hands := fmt.Sprintf(`<line x1="150" y1="150" x2="%f" y2="%f" style="fill:none;stroke:#000;stroke-width:3px;"/>
	<line x1="150" y1="150" x2="%f" y2="%f" style="fill:none;stroke:#808080;stroke-width:2px;"/>
	<line x1="150" y1="150" x2="%f" y2="%f" style="fill:none;stroke:#f00;stroke-width:1px;"/>`,
		analogClock.hour.X, analogClock.hour.Y,
		analogClock.minute.X, analogClock.minute.Y,
		analogClock.second.X, analogClock.second.Y)

	return hands
}

func DrawSVG(w io.Writer, t time.Time) {
	clockHands := DrawHand(t)

	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	io.WriteString(w, clockHands)
	io.WriteString(w, svgEnd)
}
