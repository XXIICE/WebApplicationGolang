//ref: https://medium.com/odds-team/%E0%B8%AA%E0%B8%A3%E0%B8%B8%E0%B8%9B%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B9%80%E0%B8%A3%E0%B8%B5%E0%B8%A2%E0%B8%99%E0%B8%9E%E0%B8%B7%E0%B9%89%E0%B8%99%E0%B8%90%E0%B8%B2%E0%B8%99%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2-go-%E0%B9%81%E0%B8%9A%E0%B8%9A-step-by-step-%E0%B8%88%E0%B8%B2%E0%B8%81-course-pre-ultimate-go-by-p-yod-%E0%B8%95%E0%B8%AD%E0%B8%99-3-%E0%B8%A3%E0%B8%B9%E0%B9%89%E0%B8%88%E0%B8%B1%E0%B8%81%E0%B8%81%E0%B8%B1%E0%B8%9A-1b8329a477cc
//normal none defer
package main

import (
	"fmt"
)

func main() {
	printFirst()
	printFinish()
	printSecond()
}
func printFirst() {
	fmt.Println("First")
}
func printSecond() {
	fmt.Println("Second")
}
func printFinish() {
	fmt.Println("Close")
}

//use defer คือทำอย่างอื่นก่อนแล้วค่อยมาทำ defer
func main() {
	printFirst()
	defer printFinish() // มีการเพิ่ม Defer มาที่ Function นี้
	printSecond()
}
func printFirst() {
	fmt.Println("First")
}
func printSecond() {
	fmt.Println("Second")
}
func printFinish() {
	fmt.Println("Close")
}

//have defer in many func ทำ defer อันที่ประกาศต่ำสุดก่อน
func main() {
    printFirst();
    defer printThird(); // -> ลำดับที่ 3
    defer printFourth(); // -> ลำดับที่ 2
    defer printFinish(); // -> ลำดับที่ 1
    printSecond();
}
func printFirst() {
    fmt.Println("First");
}
func printSecond() {
    fmt.Println("Second");
}
func printFinish() {
    fmt.Println("Close");
}
func printThird() {
    fmt.Println("Third");
}
func printFourth() {
    fmt.Println("Fourth");
}

//Struct ประเภทของข้อมูลประเภทหนึ่ง ที่รวมลักษณะต่างๆของข้อมูลเอาไว้ด้วยกัน
package main
import "fmt"
type customer struct { // การประกาศโครงสร้าง struct
    firstname string
    lastname  string
    code int
    phone string
}
func main() {
   cus := customer{
                    firstname: "Chaiyarin",
                    lastname: "Niamsuwan", 
                    code: 111990, 
                    phone: "085661234",
                  }; // การกำหนดค่าเริ่มต้น ให้ customer struct
   fmt.Println(cus);
}
// [ผลลัพธ์ที่ได้จากการแสดงตัวแปร cus]
// {Chaiyarin Niamsuwan 111990 085661234}

// [แสดง Field ทีละ Field ใน Struct]
fmt.Println(cus.firstname);
fmt.Println(cus.lastname);
fmt.Println(cus.code);
fmt.Println(cus.phone);
// [ผลลัพธ์จากการแสดงค่า]
// Chaiyarin
// Niamsuwan
// 111990
// 085661234

// [แสดง Field ทีละ Field ใน Struct]
cus.firstname = "Atikom";
cus.lastname = "Sombutjalearn"
fmt.Println(cus);
// [ผลลัพธ์จากการแสดงค่าจากตัวแปร cus]
// {Atikom Sombutjalearn 111990 085661234}


//Use  Struct with Array
customerLists := [3]customer{};
customerLists[0] = customer{
                      firstname: "Chaiyarin",
                      lastname: "Niamsuwan", 
                      code: 111990, 
                      phone: "085661234",
                   };
customerLists[1] = customer{
                      firstname: "Atikom",
                      lastname: "Sombutjalearn", 
                      code: 111991, 
                      phone: "085664321",
                   };
customerLists[2] = customer{
                      firstname: "Kritsana",
                      lastname: "Punyaphon", 
                      code: 111992, 
                      phone: "085662344",
                   };
fmt.Println(customerLists);
// [ผลลัพธ์จากการแสดงค่า customerLists]
// [
//   {Chaiyarin Niamsuwan 111990 085661234} 
//   {Atikom Sombutjalearn 111991 085664321} 
//   {Kritsana Punyaphon 111992 085662344}
// ]

//Use Struct with Slice

customerLists := []customer{};
customer1 := customer{
                      firstname: "Chaiyarin",
                      lastname: "Niamsuwan", 
                      code: 111990, 
                      phone: "085661234",
                   };
customer2 := customer{
                      firstname: "Atikom",
                      lastname: "Sombutjalearn", 
                      code: 111991, 
                      phone: "085664321",
                   };
customer3 := customer{
                      firstname: "Kritsana",
                      lastname: "Punyaphon", 
                      code: 111992, 
                      phone: "085662344",
                   };
customerLists = append(customerLists, customer1);
customerLists = append(customerLists, customer2);
customerLists = append(customerLists, customer3);
fmt.Println(customerLists);
// [ผลลัพธ์จากการแสดงค่า customerLists]
// [
//   {Chaiyarin Niamsuwan 111990 085661234} 
//   {Atikom Sombutjalearn 111991 085664321} 
//   {Kritsana Punyaphon 111992 085662344}
// ]


//Convert to JSON format
// [การทำ List จาก Array หรือ Slice เป็น JSON]
package main
import "fmt"
import "encoding/json" // มีการ Import Encoding JSON เข้ามา
type customer struct { // มีการปรับ Field ให้ขึ้นต้นตัวใหญ่
    Firstname string
    Lastname  string
    Code int
    Phone string
}
func main() {
  customerLists := []customer{};
  customer1 := customer{
                      Firstname: "Chaiyarin",
                      Lastname: "Niamsuwan", 
                      Code: 111990, 
                      Phone: "085661234",
                   };
  customer2 := customer{
                      Firstname: "Atikom",
                      Lastname: "Sombutjalearn", 
                      Code: 111991, 
                      Phone: "085664321",
                   };
  customer3 := customer{
                      Firstname: "Kritsana",
                      Lastname: "Punyaphon", 
                      Code: 111992, 
                      Phone: "085662344",
                   };
  customerLists = append(customerLists, customer1);
  customerLists = append(customerLists, customer2);
  customerLists = append(customerLists, customer3);
  customerListsJson, err := json.Marshal(customerLists)
  
  if err != nil {
    fmt.Println(err)
  }
  
  fmt.Println(string(customerListsJson));
}
// [ผลลัพธ์ที่ได้จากการแสดงค่าจากตัวแปร customerListsJson]
// [{"Firstname":"Chaiyarin","Lastname":"Niamsuwan","Code":111990,"Phone":"085661234"},{"Firstname":"Atikom","Lastname":"Sombutjalearn","Code":111991,"Phone":"085664321"},{"Firstname":"Kritsana","Lastname":"Punyaphon","Code":111992,"Phone":"085662344"}]