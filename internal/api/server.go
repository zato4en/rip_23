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

	Spectrums := []rip2023.Spectrum{
		{1, "Космический микроволновый фон (CMB)", []float32{1, 2, 3}, "представляет собой радиацию с длиной волны около 1 мм (миллиметра). CMB является самым ранним излучением, которое мы можем наблюдать, и его спектр соответствует температуре около 2.7 Кельвина.\n",
			"relict.jpeg"},
		{2, "Инфракрасное реликтовое излучение (IRB)", []float32{1, 2, 3}, "соответствует длинам волн от 1 мм до 1 микрометра. IRB связано с тепловым излучением, испускаемым пылью и газом в галактиках.\n",
			"IRB.jpeg"},
		{3, "Рентгеновское реликтовое излучение (XRB)", []float32{1, 2, 3}, "соответствует длинам волн от 1 ангстрема до 1 пикометра. XRB связано с высокоэнергетическими процессами, такими как аккреция вокруг черных дыр и нейтронных звезд.\n",
			"xrb.jpeg"},
	}

	r.GET("/", func(c *gin.Context) {
		r.SetHTMLTemplate(template.Must(template.ParseFiles("./templates/mainpage.html")))
		c.HTML(http.StatusOK, "mainpage.html", gin.H{
			"Spectrum": Spectrums,
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

		// Здесь мы передаем выбранный AMS в шаблон
		c.HTML(http.StatusOK, "spectrum.html", gin.H{
			"Spectrum": selectedSpectrum,
		})
	})

	r.GET("/search", func(c *gin.Context) {
		searchQuery := c.Query("search")

		filteredSpectrum := []rip2023.Spectrum{}
		for _, c := range Spectrums {
			if strings.Contains(strings.ToLower(c.Name), strings.ToLower(searchQuery)) {
				filteredSpectrum = append(filteredSpectrum, c)
			}
		}

		c.HTML(http.StatusOK, "mainpage.html", gin.H{
			"Spectrum": filteredSpectrum,
		})
	})

	r.Run(":8080")
}
