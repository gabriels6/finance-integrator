package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"google.golang.org/genai"
)

// ParsePDFHandler receives a multipart/form-data request with field `file` containing a PDF.
// It validates the file looks like a PDF and forwards the bytes to the service layer.
func ParsePDFHandler(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "file is required", "error": err.Error()})
		return
	}
	defer file.Close()

	err = EnableClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error on AI parser", "error": err.Error()})
		return
	}

	aiFile, err := UploadFile(file, "application/pdf")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to upload PDF", "error": err.Error()})
		return
	}

	filePart := genai.NewPartFromURI(aiFile.URI, aiFile.MIMEType)

	fileContent := genai.Content{
		Parts: []*genai.Part{
			filePart,
		},
		Role: "user",
	}

	textContent := genai.Content{
		Parts: []*genai.Part{
			genai.NewPartFromText("\n \n"),
			genai.NewPartFromText("The PDF file is a portfolio. Extract it to a json as the following { \"broker\":\"Broker\" \"portfolioCode\": \"1234\", \"financialAmount\": 123456.78, \"positions\": [ { \"assetType\": \"Cash\", \"assetCode\": \"TEST123\", \"assetDescription\":\"The stock TEST1234\", \"currency\":\"USD\", \"quantity\": 100.00, \"price\": 10.00, \"financialAmount\": 1000.0, \"maturityDate\":\"2025-05-06\"  } ]  }. The broker code must be the broker which generated the file, on upper case, with _ instead of spaces between words and no special characters."),
			genai.NewPartFromText("Also, for the asset type, divide in current: FIRF (Fundos renda fixa), FIC FIRF, FIA (Fundos acoes), FIC FIA, FIC FIRF, FIM (Fundos Multimercado), FII (Fundos Imobiliarios), OTHER_FUNDS, FIXED INCOME PUBLIC (Government bonds), FIXED INCOME PRIVATE (Corporate Bonds), ETF, STOCKS, REITS, CASH, FUTURES, OPTIONS, DERIVATIVES, OTHERS. Also, if the sum of positions.financialAmount - portfolio.financialAmount is not zero, add the difference to a new position of assetType OTHER, assetDescription \"assets or liabities\", price = 1, quantity = 1 and financialAmount = the difference"),
		},
		Role: "user",
	}

	result, err := AskForResponse([]*genai.Content{
		&fileContent,
		&textContent,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to parse PDF", "error": err.Error()})
		return
	}

	c.Data(http.StatusOK, "application/json", ClearAiJSON(result))
}

func ClearAiJSON(jsonStr string) []byte {
	resp := strings.ReplaceAll(strings.ReplaceAll(jsonStr, "`", ""), "json", "")

	return []byte(resp)
}
