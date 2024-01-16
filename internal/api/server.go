package api

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"rip2023"
	"strconv"
	"strings"
)

func StartServer() {
	r := gin.Default()

	r.Static("/resources", "./resources")
	r.Static("/static", "./static")

	//МОКИ ЗДЕСЬ
	Spectrums := []rip2023.Spectrum{
		{1, "CMB1", []float32{1, 2, 3}, "Описание спектра 1\n",
			"relict.jpeg"},
		{2, "CMB2", []float32{1, 2, 3}, "Описание спектра 2\n",
			"IRB.jpeg"},
		{3, "CMB3", []float32{1, 2, 3}, "Описание спектра 3\n",
			"xrb.jpeg"},
	}

	r.GET("/", func(c *gin.Context) {
		searchQuery := c.Query("search")

		filteredSpectrum := []rip2023.Spectrum{}
		for _, c := range Spectrums {
			if strings.Contains(strings.ToLower(c.Name), strings.ToLower(searchQuery)) {
				filteredSpectrum = append(filteredSpectrum, c)
			}
		}

		r.SetHTMLTemplate(template.Must(template.ParseFiles("./templates/mainpage.html")))
		c.HTML(http.StatusOK, "mainpage.html", gin.H{
			"Spectrum": filteredSpectrum,
		})
	})

	r.GET("/Spectrum/:id", func(c *gin.Context) {
		r.SetHTMLTemplate(template.Must(template.ParseFiles("./templates/spectrum.html")))
		id := c.Param("id")

		var selectedSpectrum rip2023.Spectrum

		for _, Spectrum := range Spectrums {
			if strconv.Itoa(Spectrum.ID) == id {
				selectedSpectrum = Spectrum
				break
			}
		}
		//c
		c.HTML(http.StatusOK, "spectrum.html", gin.H{
			"Spectrum": selectedSpectrum,
		})
	})

	r.Run(":8080")
}
