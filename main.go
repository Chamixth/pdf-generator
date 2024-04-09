package main

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

type PaymentReceipt struct {
	PaymentId     string  `json:"PaymentId"`
	ApplicantId   string  `json:"ApplicantId"`
	ApplicantName string  `json:"ApplicantName"`
	Amount        float64 `json:"Amount"`
	PaymentDate   string  `json:"PaymentDate"`
	Status        string  `json:"Status"`
}

func main() {
	payment := PaymentReceipt{
		PaymentId:     "16434922-ce03-45b4-af55-58771568b1dc",
		ApplicantId:   "987654321",
		ApplicantName: "John Doe",
		Amount:        100.50,
		PaymentDate:   "2024-04-09",
		Status:        "Completed",
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// Add logo
	logoPath := "nextler.jpg"
	logoWidth := 50.0 // Width in millimeters
	logoHeight := 0.0 // Height 0 means proportional

	// Calculate center position for image
	pageWidth, _ := pdf.GetPageSize()
	imageX := (pageWidth - logoWidth) / 2

	pdf.ImageOptions(logoPath, imageX, 10, logoWidth, logoHeight, false, gofpdf.ImageOptions{}, 0, "")

	pdf.Ln(30)

	// Set title font and color
	pdf.SetFont("Arial", "B", 24)
	pdf.SetFillColor(173, 216, 230) // Light blue color
	pdf.SetTextColor(0, 0, 0)       // Black text color

	// Calculate title position
	_, lineHeight := pdf.GetFontSize()
	titleHeight := lineHeight * 2 // Double the line height for title

	// Print the title
	pdf.CellFormat(pageWidth-20, titleHeight, "Payment Receipt", "", 0, "C", false, 0, "")

	// Move the table down
	pdf.Ln(30)

	// Calculate the total width of your box
	boxWidth := 150.0
	boxHeight := 100.0

	// Calculate the starting position for the box to center it on the page
	pageWidth, _ = pdf.GetPageSize()
	boxX := (pageWidth - boxWidth) / 2.0
	boxY := 100.0

	pdf.Rect(boxX, boxY-25, boxWidth, boxHeight-25, "")

	// Write details inside the box
	pdf.SetFont("Arial", "", 12)
	pdf.SetTextColor(0, 0, 0) // Black text color

	fields := []string{"Applicant ID:", "Payment Date:", "Amount:", "Applicant Name:", "Status:"}
	values := []string{payment.ApplicantId, payment.PaymentDate, fmt.Sprintf("$%.2f", payment.Amount), payment.ApplicantName, payment.Status}

	lineHeight = 7.0
	currentY := boxY + 5.0 // Start 5 units from the top of the box
	for i, field := range fields {
		pdf.Ln(8)
		pdf.SetFont("Arial", "B", 12) // Set font to bold
		pdf.SetX(boxX + 5.0)          // Start 5 units from the left of the box
		pdf.CellFormat(40, lineHeight, field, "", 0, "", false, 0, "")

		pdf.SetX(boxX + 55.0)        // Start 55 units from the left of the box (create a gap)
		pdf.SetFont("Arial", "", 12) // Set font back to regular for the values

		// Add extra space between field and value
		pdf.CellFormat(40, lineHeight, "", "", 0, "", false, 0, "") // Empty cell with 10 units width

		pdf.CellFormat(0, lineHeight, values[i], "", 1, "", false, 0, "")

		currentY += lineHeight
	}
	pdf.Ln(30)
	pdf.SetFont("Arial", "B", 16)
	pdf.SetTextColor(0, 128, 0)     // Green text color
	pdf.SetFillColor(173, 216, 230) // Light blue background color
	pdf.SetY(currentY + 40)   
	pdf.SetX(30)   
	pdf.CellFormat(150, 15, "Payment is successful", "1", 1, "C", true, 0, "")

	err := pdf.OutputFileAndClose("payment_receipt.pdf")
	if err != nil {
		fmt.Println("Error creating PDF:", err)
		return
	}

	fmt.Println("PDF payment receipt generated successfully.")
}
