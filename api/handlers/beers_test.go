package handlers

import (
	"encoding/json"

	"github.com/go-resty/resty/v2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/thiiluh/kombibeer/internal/config"
	"github.com/thiiluh/kombibeer/pkg/beers"
)

var b = beers.Beer{
	Name:           "Test Beer",
	Ingredients:    "Test Ingredients",
	AlcoholContent: "Test Alcohol Content",
	Price:          2.5,
	Category:       "Test Category",
}

var client = resty.New().SetBaseURL("http://localhost:8080")

var _ = Describe("Beers", Label("beers"), func() {

	Describe("GET Operations", func() {

		Context("when there are registered beers", func() {
			BeforeEach(func() {
				config.DB.Create(&b)
			})
			It("should return all beers", func() {
				resp, err := client.R().Get("/beers")
				Expect(err).To(BeNil())
				Expect(resp.StatusCode()).To(Equal(200))
			})
		})

		Context("when there are no registered beers", func() {
			BeforeEach(func() {
				config.DB.Delete(&b)
			})
			It("should return no content", func() {
				resp, err := client.R().Get("/beers")
				Expect(err).To(BeNil())
				Expect(resp.StatusCode()).To(Equal(204))
			})
		})
	})

	Describe("POST Operations", func() {
		Context("when successful", func() {
			AfterEach(func() {
				config.DB.Delete(&b)
			})
			It("should return all beers", func() {
				resp, err := client.R().SetBody(b).Post("/beers")
				Expect(err).To(BeNil())
				var beerResp beers.Beer
				err = json.Unmarshal(resp.Body(), &beerResp)

				Expect(err).To(BeNil())
				Expect(resp.StatusCode()).To(Equal(201))
				Expect(b).To(Equal(beerResp))

				resp, err = client.R().Get("/beers")
				Expect(err).To(BeNil())
				Expect(resp.StatusCode()).To(Equal(200))
			})
		})
	})
})
