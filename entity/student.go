package entity

import "github.com/asaskevich/govalidator"


type Student struct{
	Fullname string  `valid:"required~Fullname is required"`
	Age int  `valid:"required~Age is required, range(18|100)~Age must be between 18 and 100"`
	Email string  `valid:"required~Email is required, email~Email must be a valid email address"`
	GPA float32  `valid:"range(0|4)~GPA must be between 0.0 and 4.0"`
}

 func ValidationStudent(student *Student) (bool, error){
	return govalidator.ValidateStruct(student)
 }