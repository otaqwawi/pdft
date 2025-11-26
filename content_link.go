package pdft

import (
	"bytes"
	"fmt"
)

// ContentExternalLink represents an external link (hyperlink) in PDF
type ContentExternalLink struct {
	url        string
	pageNum    int
	x          float64
	y          float64
	w          float64
	h          float64
	pageHeight float64
}

// toSteram generates PDF stream for external link annotation
func (c *ContentExternalLink) toSteram() (*bytes.Buffer, error) {
	var buff bytes.Buffer

	// Calculate coordinates (PDF uses bottom-left origin)
	y1 := c.pageHeight - c.y - c.h
	y2 := c.pageHeight - c.y

	// Create link annotation dictionary
	// /Subtype /Link: defines this as a link annotation
	// /Rect: defines the clickable area [x1 y1 x2 y2]
	// /Border: [0 0 0] means no border
	// /A: action dictionary with /S /URI for external links
	buff.WriteString("<< /Type /Annot /Subtype /Link ")
	buff.WriteString(fmt.Sprintf("/Rect [%.2f %.2f %.2f %.2f] ", c.x, y1, c.x+c.w, y2))
	buff.WriteString("/Border [0 0 0] ")
	buff.WriteString(fmt.Sprintf("/A << /S /URI /URI (%s) >> ", c.url))
	buff.WriteString(">>\r\n")

	return &buff, nil
}

// page returns the page number where this link should be placed
func (c *ContentExternalLink) page() int {
	return c.pageNum
}

// ContentInternalLink represents an internal link (jump to another page) in PDF
type ContentInternalLink struct {
	targetPage int
	pageNum    int
	x          float64
	y          float64
	w          float64
	h          float64
	pageHeight float64
}

// toSteram generates PDF stream for internal link annotation
func (c *ContentInternalLink) toSteram() (*bytes.Buffer, error) {
	var buff bytes.Buffer

	// Calculate coordinates (PDF uses bottom-left origin)
	y1 := c.pageHeight - c.y - c.h
	y2 := c.pageHeight - c.y

	// Create link annotation dictionary for internal link
	// /Dest: destination array [page /Fit] to fit the target page in window
	// Note: page reference will be resolved during PDF build
	buff.WriteString("<< /Type /Annot /Subtype /Link ")
	buff.WriteString(fmt.Sprintf("/Rect [%.2f %.2f %.2f %.2f] ", c.x, y1, c.x+c.w, y2))
	buff.WriteString("/Border [0 0 0] ")
	buff.WriteString(fmt.Sprintf("/Dest [%d 0 R /Fit] ", c.targetPage))
	buff.WriteString(">>\r\n")

	return &buff, nil
}

// page returns the page number where this link should be placed
func (c *ContentInternalLink) page() int {
	return c.pageNum
}
