package handlers

import (
    "github.com/jung-kurt/gofpdf"
    "net/http"
    "os"
    "path/filepath"
    "github.com/Hybrid-Codes/GoKit/models"
)

func GeneratePDFHandler(w http.ResponseWriter, r *http.Request) {
    var users []models.User
    if err := models.DB.Find(&users).Error; err != nil {
        http.Error(w, "Error fetching users", http.StatusInternalServerError)
        return
    }

    pdf := gofpdf.New("P", "mm", "A4", "")
    pdf.AddPage()
    pdf.SetFont("Arial", "B", 16)

    pdf.Cell(40, 10, "User Report")
    pdf.Ln(12)

    pdf.SetFont("Arial", "", 12)
    for _, user := range users {
        pdf.Cell(40, 10, "Name: "+user.Name)
        pdf.Ln(8)
        pdf.Cell(40, 10, "Email: "+user.Email)
        pdf.Ln(12)
    }

    // Save PDF to folder
    folderPath := "./reports"
    if _, err := os.Stat(folderPath); os.IsNotExist(err) {
        os.Mkdir(folderPath, os.ModePerm)
    }

    filePath := filepath.Join(folderPath, "user_report.pdf")
    if err := pdf.OutputFileAndClose(filePath); err != nil {
        http.Error(w, "Error saving PDF", http.StatusInternalServerError)
        return
    }

    // Serve the PDF to the client for download
    w.Header().Set("Content-Type", "application/pdf")
    w.Header().Set("Content-Disposition", "inline; filename=user_report.pdf")

    pdf.Output(w)
}
