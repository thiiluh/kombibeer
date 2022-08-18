package handlers

import (
	"encoding/json"
	"strings"

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

var copyBeer = beers.Beer{
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

		Context("field validation when", func() {
			BeforeEach(func() {
				copyBeer = b
			})
			It("name less than 1 characters", func() {
				copyBeer.Name = ""
				resp, err := client.R().SetBody(copyBeer).Post("/beers")
				Expect(err).To(BeNil())
				var beerResp beers.Beer
				err = json.Unmarshal(resp.Body(), &beerResp)
				Expect(err).To(BeNil())
				Expect(resp.StatusCode()).To(Equal(400))
				Expect(beerResp.Name).To(Equal("cannot be blank"))
			})
			It("name greater than 50 characters", func() {
				copyBeer.Name = strings.Repeat("a", 51)
				resp, err := client.R().SetBody(copyBeer).Post("/beers")
				Expect(err).To(BeNil())
				var beerResp beers.Beer
				err = json.Unmarshal(resp.Body(), &beerResp)
				Expect(err).To(BeNil())
				Expect(resp.StatusCode()).To(Equal(400))
				Expect(beerResp.Name).To(Equal("the length must be between 1 and 50"))
			})
			It("ingredients less than 1 characters", func() {
				copyBeer.Ingredients = ""
				resp, err := client.R().SetBody(copyBeer).Post("/beers")
				Expect(err).To(BeNil())
				var beerResp beers.Beer
				err = json.Unmarshal(resp.Body(), &beerResp)
				Expect(err).To(BeNil())
				Expect(resp.StatusCode()).To(Equal(400))
				Expect(beerResp.Ingredients).To(Equal("cannot be blank"))
			})
			It("ingredients greater than 50 characters", func() {
				copyBeer.Ingredients = strings.Repeat("a", 51)
				resp, err := client.R().SetBody(copyBeer).Post("/beers")
				Expect(err).To(BeNil())
				var beerResp beers.Beer
				err = json.Unmarshal(resp.Body(), &beerResp)
				Expect(err).To(BeNil())
				Expect(resp.StatusCode()).To(Equal(400))
				Expect(beerResp.Ingredients).To(Equal("the length must be between 1 and 50"))
			})
			It("alcoholContent less than 1 characters", func() {
				copyBeer.AlcoholContent = ""
				resp, err := client.R().SetBody(copyBeer).Post("/beers")
				Expect(err).To(BeNil())
				var beerResp beers.Beer
				err = json.Unmarshal(resp.Body(), &beerResp)
				Expect(err).To(BeNil())
				Expect(resp.StatusCode()).To(Equal(400))
				Expect(beerResp.AlcoholContent).To(Equal("cannot be blank"))
			})
			It("alcoholContent greater than 50 characters", func() {
				copyBeer.AlcoholContent = strings.Repeat("a", 51)
				resp, err := client.R().SetBody(copyBeer).Post("/beers")
				Expect(err).To(BeNil())
				var beerResp beers.Beer
				err = json.Unmarshal(resp.Body(), &beerResp)
				Expect(err).To(BeNil())
				Expect(resp.StatusCode()).To(Equal(400))
				Expect(beerResp.AlcoholContent).To(Equal("the length must be between 1 and 50"))
			})
			It("price must be greater than or equal to 1", func() {
				copyBeer.Price = 0.9
				resp, err := client.R().SetBody(copyBeer).Post("/beers")
				Expect(err).To(BeNil())
				var beerResp interface{}
				err = json.Unmarshal(resp.Body(), &beerResp)
				Expect(err).To(BeNil())
				Expect(resp.StatusCode()).To(Equal(400))
				Expect(beerResp).To(ContainElement("must be no less than 1"))
			})
			It("price must be less than or equal to 100", func() {
				copyBeer.Price = 100.1
				resp, err := client.R().SetBody(copyBeer).Post("/beers")
				Expect(err).To(BeNil())
				var beerResp interface{}
				err = json.Unmarshal(resp.Body(), &beerResp)
				Expect(err).To(BeNil())
				Expect(resp.StatusCode()).To(Equal(400))
				Expect(beerResp).To(ContainElement("must be no greater than 100"))
			})
			It("category less than 1 characters", func() {
				copyBeer.Category = ""
				resp, err := client.R().SetBody(copyBeer).Post("/beers")
				Expect(err).To(BeNil())
				var beerResp beers.Beer
				err = json.Unmarshal(resp.Body(), &beerResp)
				Expect(err).To(BeNil())
				Expect(resp.StatusCode()).To(Equal(400))
				Expect(beerResp.Category).To(Equal("cannot be blank"))
			})
			It("category greater than 50 characters", func() {
				copyBeer.Category = strings.Repeat("a", 51)
				resp, err := client.R().SetBody(copyBeer).Post("/beers")
				Expect(err).To(BeNil())
				var beerResp beers.Beer
				err = json.Unmarshal(resp.Body(), &beerResp)
				Expect(err).To(BeNil())
				Expect(resp.StatusCode()).To(Equal(400))
				Expect(beerResp.Category).To(Equal("the length must be between 1 and 50"))
			})
		})
	})
})
