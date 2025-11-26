package pdft

import (
	"testing"

	gopdf "github.com/otaqwawi/pdft/minigopdf"
)

func TestAddExternalLink(t *testing.T) {
	var ipdf PDFt
	err := ipdf.Open("test/pdf/pdf_from_docx_with_f.pdf")
	if err != nil {
		t.Error(err)
		return
	}

	err = ipdf.AddFont("arial", "test/ttf/angsa.ttf")
	if err != nil {
		t.Error(err)
		return
	}

	err = ipdf.SetFont("arial", "", 14)
	if err != nil {
		t.Error(err)
		return
	}

	// Add text
	err = ipdf.Insert("Click here to visit Google", 1, 50, 50, 200, 30, gopdf.Left|gopdf.Top, nil)
	if err != nil {
		t.Error(err)
		return
	}

	// Add external link over the text
	err = ipdf.AddExternalLink("https://www.google.com", 1, 50, 50, 200, 30)
	if err != nil {
		t.Error(err)
		return
	}

	err = ipdf.Save("test/out/test_external_link.pdf")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestAddInternalLink(t *testing.T) {
	var ipdf PDFt
	err := ipdf.Open("test/pdf/pdf_from_docx_with_f.pdf")
	if err != nil {
		t.Error(err)
		return
	}

	err = ipdf.AddFont("arial", "test/ttf/angsa.ttf")
	if err != nil {
		t.Error(err)
		return
	}

	err = ipdf.SetFont("arial", "", 14)
	if err != nil {
		t.Error(err)
		return
	}

	// Add text on page 1
	err = ipdf.Insert("Click here to go to page 2", 1, 50, 50, 200, 30, gopdf.Left|gopdf.Top, nil)
	if err != nil {
		t.Error(err)
		return
	}

	// Add internal link to page 2
	err = ipdf.AddInternalLink(2, 1, 50, 50, 200, 30)
	if err != nil {
		t.Error(err)
		return
	}

	// Add text on page 2
	err = ipdf.Insert("You are now on page 2", 2, 50, 50, 200, 30, gopdf.Left|gopdf.Top, nil)
	if err != nil {
		t.Error(err)
		return
	}

	err = ipdf.Save("test/out/test_internal_link.pdf")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestMultipleLinks(t *testing.T) {
	var ipdf PDFt
	err := ipdf.Open("test/pdf/pdf_from_docx_with_f.pdf")
	if err != nil {
		t.Error(err)
		return
	}

	err = ipdf.AddFont("arial", "test/ttf/angsa.ttf")
	if err != nil {
		t.Error(err)
		return
	}

	err = ipdf.SetFont("arial", "", 14)
	if err != nil {
		t.Error(err)
		return
	}

	// Add multiple external links on page 1
	err = ipdf.Insert("Google", 1, 50, 50, 100, 20, gopdf.Left|gopdf.Top, nil)
	if err != nil {
		t.Error(err)
		return
	}
	err = ipdf.AddExternalLink("https://www.google.com", 1, 50, 50, 100, 20)
	if err != nil {
		t.Error(err)
		return
	}

	err = ipdf.Insert("GitHub", 1, 50, 80, 100, 20, gopdf.Left|gopdf.Top, nil)
	if err != nil {
		t.Error(err)
		return
	}
	err = ipdf.AddExternalLink("https://www.github.com", 1, 50, 80, 100, 20)
	if err != nil {
		t.Error(err)
		return
	}

	// Add internal link
	err = ipdf.Insert("Go to page 2", 1, 50, 110, 100, 20, gopdf.Left|gopdf.Top, nil)
	if err != nil {
		t.Error(err)
		return
	}
	err = ipdf.AddInternalLink(2, 1, 50, 110, 100, 20)
	if err != nil {
		t.Error(err)
		return
	}

	err = ipdf.Save("test/out/test_multiple_links.pdf")
	if err != nil {
		t.Error(err)
		return
	}
}
