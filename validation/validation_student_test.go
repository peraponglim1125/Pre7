package validation_test

import (
	"testing"
	. "github.com/onsi/gomega"
	"github.com/peraponglim1125/Pre5/entity"
)

func TestValid(t *testing.T){
	g := NewGomegaWithT(t)
    student := entity.Student{
		Fullname: "John",
		Age: 20,
		Email: "Phirapong@example.com",
		GPA: 3.5,
	}
	ok, err := entity.ValidationStudent(&student)
	g.Expect(ok).To(BeTrue())
	g.Expect(err).To(BeNil())
}

func TestInValidName(t *testing.T){
	g := NewGomegaWithT(t)
    student := entity.Student{
		Fullname: "",
		Age: 20,
		Email: "Phirapong@example.com",
		GPA: 3.5,
	}
	ok, err := entity.ValidationStudent(&student)
	g.Expect(ok).To(BeFalse())
	g.Expect(err.Error()).To(Equal("Fullname is required"))
}

func TestInValidAge(t *testing.T){
	g := NewGomegaWithT(t)
    student := entity.Student{
		Fullname: "John",
		Age: 0,
		Email: "Phirapong@example.com",
		GPA: 3.5,
	}
	ok, err := entity.ValidationStudent(&student)
	g.Expect(ok).To(BeFalse())
	g.Expect(err.Error()).To(Equal("Age is required"))
}
func TestInvalidEmail(t *testing.T){
	g := NewGomegaWithT(t)
	student := entity.Student{
		Fullname: "John",
		Age: 20,
		Email: "invalid-email",		
		GPA: 3.5,
	}
	ok, err := entity.ValidationStudent(&student)
	g.Expect(ok).To(BeFalse())
	g.Expect(err.Error()).To(Equal("Email must be a valid email address"))
}
func TestInvalidEmail2(t *testing.T){
	g := NewGomegaWithT(t)
	student := entity.Student{
		Fullname: "John",
		Age: 20,
		Email: "",		
		GPA: 3.5,
	}
	ok, err := entity.ValidationStudent(&student)
	g.Expect(ok).To(BeFalse())
	g.Expect(err.Error()).To(Equal("Email is required"))
}

func TestInvalidGPA(t *testing.T){
	g := NewGomegaWithT(t)
	student := entity.Student{
		Fullname: "John",
		Age: 20,
		Email: "Phirapong@example.com",		
		GPA: 45.0 ,
	}
	ok, err := entity.ValidationStudent(&student)
	g.Expect(ok).To(BeFalse())
	g.Expect(err.Error()).To(Equal("GPA must be between 0.0 and 4.0"))
}
