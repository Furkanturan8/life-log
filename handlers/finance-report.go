package handlers

import (
	"fmt"
	"github.com/bmdavis419/the-better-backend/models"
	"github.com/bmdavis419/the-better-backend/pkg/utils"
	"github.com/bmdavis419/the-better-backend/service"
	"github.com/gofiber/fiber/v2"
	"github.com/jung-kurt/gofpdf"
)

type ReportHandler struct {
	reportService service.IReportService
}

func NewReportHandler(s service.IReportService) ReportHandler {
	return ReportHandler{reportService: s}
}

func (h ReportHandler) GenerateReport(ctx *fiber.Ctx) error {
	var reportData models.GenerateReport

	if err := ctx.BodyParser(&reportData); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	report, err := h.reportService.GenerateReport(ctx.Context(), reportData.UserID, reportData.Month, reportData.Year)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Report not found"})
	}
	// PDF oluşturma
	pdf := gofpdf.New("L", "mm", "A4", "")
	pdf.AddPage()

	// Başlık
	pdf.SetFont("Times", "B", 28)
	pdf.Cell(40, 10, "PDF Report")
	pdf.SetTextColor(0, 0, 0) // Siyah renk
	pdf.Ln(12)

	// Raporun oluşturulma tarihi
	pdf.SetFont("Times", "B", 20)
	pdf.SetTextColor(0, 0, 255)
	pdf.Cell(40, 10, "Generated At: "+report.GeneratedAt.Format("Mon Jan 2, 2006 15:04:05"))
	pdf.Ln(20)

	// PDF tablosu başlık
	pdf.SetFillColor(240, 240, 240)
	pdf.SetTextColor(0, 0, 0) // Siyah renk
	pdf.SetFont("Times", "B", 16)
	headers := []string{"Adi Soyadi", "Email", "Islem Tarihleri", "Toplam Gelir", "Toplam Harcama", "Bakiye"}
	for _, str := range headers {
		pdf.CellFormat(45, 10, str, "1", 0, "C", true, 0, "")
	}
	pdf.Ln(-1)

	data := []string{
		report.User.Name + " " + report.User.Surname,
		report.User.Email,
		fmt.Sprintf("%02d/%d", report.Month, report.Year),
		fmt.Sprintf("%.2f", report.TotalIncome),
		fmt.Sprintf("%.2f", report.TotalExpense),
		fmt.Sprintf("%.2f", report.Balance),
	}

	pdf.SetFont("Times", "", 13) // Normal font ayarı
	maxHeight := 0.0

	colWidths := []float64{45, 45, 45, 45, 45, 45}
	rowHeights := make([]float64, len(data))

	for i, str := range data {
		x, y := pdf.GetX(), pdf.GetY()
		pdf.MultiCell(colWidths[i], 10, str, "1", "C", false)
		rowHeights[i] = pdf.GetY() - y // Bu sütunun yüksekliği
		pdf.SetXY(x+colWidths[i], y)

		if rowHeights[i] > maxHeight {
			maxHeight = rowHeights[i]
		}
	}

	randomString := utils.GenerateRandomString(2)
	filePath := fmt.Sprintf("./uploads/report-%d-%d-%s.pdf", report.Month, report.Year, randomString)

	err = pdf.OutputFileAndClose(filePath)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create PDF"})
	}

	return ctx.SendFile(filePath, false)
}
