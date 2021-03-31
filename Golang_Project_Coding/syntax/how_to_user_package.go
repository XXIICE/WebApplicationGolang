// Packages คืออะไร ?
// Packages คือ Source Code ของภาษา Go ที่อยู่ใน Directory ของระบบ โดยความสามารถของมันคือ สามารถให้นำ Source Code เหล่านั้น
// ไปให้ Application ที่พัฒนา โดยภาษา Go อื่นๆ นำไปใช้ ไป Reuse อะไรๆ ที่อยู่ภายใต้ Packages ได้


// [Filename: main.go]
// package main // มีการใส่ package main มาตอนเริ่มต้น
// func main() {
//    println("Hello World")
// }

//การใส่ package main ไปที่ไฟล์ไหนก็ตาม
// จะเป็นการบอกตัว Compiler ของภาษา Go ให้รู้ว่า 
// ไฟล์นั้นจะเป็นไฟล์ที่ถูก Execute หรือ เป็นไฟล์ที่ถูก Run นะ

//Condition to create 
// package ต้องชื่อเดียวกับชื่อ directory
// ชื่อ Variable ที่จะให้ คนที่ Import Package ของเราไปใช้ต้องมีตัวแรกตัวใหญ่
// ชื่อ Function ที่จะให้ คนที่ Import Package ของเราไปใช้ต้องมีตัวแรกตัวใหญ่

// [Filename: school.go]
package school

SchoolName := "PSC" ; 

func GetSchoolAddress() string{
	return "NakornPathom"
}
//นำ Package ไปใช้ใน main.go
// [Filename: main.go]
package main
import school
func main() {
   println("School Name", school.SchoolName);
   println("School Address:", school.GetSchoolAddress());
}