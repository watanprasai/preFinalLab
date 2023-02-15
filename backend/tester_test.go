package backend

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	Name string `valid:"required~Name cannot be blank"`
	Url  string `gorm:"uniqueIndex" valid:"url"`
}

func Test_first(t *testing.T) {
	g := NewGomegaWithT(t)

	Video := Video{
		Name: "",
		Url:  "facebook.com",
	}

	ok, err := govalidator.ValidateStruct(Video)

	g.Expect(ok).NotTo(BeTrue())
	g.Expect(err).NotTo(BeNil())
	g.Expect(err.Error()).To(Equal("Name cannot be blank"))
}
