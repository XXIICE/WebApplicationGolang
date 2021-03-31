//variable
//1.Full
var a,b,c bool;
var a,b,c string;
//2.add value
var a,b,c bool = false,false,true;
//3.None type variable
var a,b,c = false,10,"study";
//4.cut var and none type (popular)
a,b,c := false,10,"study";



// create function
func printSchoolAddress(){
	println("Bangkok");
}
//parameter function
 func printSchoolAddress(schoolAddress string)  {
	 println(schoolAddress);
 }
 //have return value
 func getSchoolAddress()string{
	return "bangkok";
 }
 //[การรับค่า Function ที่มีการ Return]
 schoolAddress := getSchoolAddress();

 //return many value
 func getSchoolAddress() (int,string){
	 code:= 1993;
	 address := "Bangkok";
	 return code,address;
 } 
 resultCode,resultAddress := getSchoolAddress();

 //Loop

//1. For Loop
for i:=1; i<9; i++{
	println(i);
}
//2.do while Loop
i:= 1
for{
	if i==2{
		println(i);
		break

	}
	i++
}
//3.While Loop
i := 1;
for i <= 5 {
    fmt.Println(i);
    i = i + 1;
}

//IF else condition
//1.normal if else
x := 10;
if x < 10 {
   println("x มีค่าน้อยกว่า 10");
} else {
   println("x มีค่ามากว่า หรือ เท่ากับ 10");
}
//2.have process before check ifelse
i := 2;
j := 3;
k := 0;
k = i + j;
if k == 5 {
   println("k มีค่าเท่ากับ 5");
}

//Array
// [การประกาศ Array เปล่าๆ ว่างๆ ]
var a [5]int; 
//a เป็น type int ที่มี Block การใส่ข้อมูล 5 ช่อง
// [การใส่ค่าใน Array แต่ละช่อง]
a[0] = 1;
a[1] = 2;
a[2] = 3;
//if value in array is null the default value is set at 0

//[การประกาศ Array แบบมีค่าเริ่มต้น]
var name = [3]string{"Chaiyarin", "Atikom", "Kritsana"};
//find size array
len(name)
//Slice (flexible array)
// [การประกาศตัวแปร Slice]
name := []string{}; 
// [การใส่ค่าใน Slice]
name = append(name, "Chaiyarin");
name = append(name, "Atikom");
name = append(name, "Kritsana");

//Map (Maps คือ Data Structure รูปแบบหนึ่ง ที่เก็บข้อมูลในลักษณะ Key Value)
//[การประกาศตัวแปร Map]
m := make(map[string]int); 
// [การใส่ค่าใน Map]
m["chaiyarin"] = 1;
m["atikom"] = 2;
//[สมมติว่าค่าใน Map มีดังนี้]
map[chaiyarin:1 atikom:2 kritsana:3]
// [เราต้องการลบ key value ของ atikom ออก เราจะใช้คำสั่ง]
delete(m, "atikom"); // delete(ตัวแปร map, key ที่ต้องการลบ)
