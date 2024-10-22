package main

// import "fmt"

// type LgTV struct {
// 	Status bool
// 	Model  string
// }

// type SamsungTV struct {
// 	Status bool
// 	Model  string
// }

// type TVer interface {
// 	switchOFF() bool
// 	GetStatus() bool
// 	GetModel() string
// }

// func (lg LgTV) switchOFF() bool {
// 	lg.Status = false
// 	fmt.Println(lg.Model, "TV is off")
// 	return true
// }

// // func switchON(tv TVer) bool {
// // 	tv.GetStatus()
// // 	return true
// // }

// func (tv LgTV) GetStatus() bool {
// 	return tv.Status
// }

// func (tv LgTV) GetModel() string {
// 	return tv.Model
// }

// func (tv SamsungTV) switchOFF() bool {
// 	tv.Status = false
// 	fmt.Println(tv.Model, "TV is off")
// 	return true
// }

// // func (tv SamsungTV) switchON () bool {
// // 	tv.Status = true
// // 	fmt.Println(tv.Model, "TV is on")
// // 	return true
// // }

// func (tv SamsungTV) GetStatus() bool {
// 	return tv.Status
// }

// func (tv SamsungTV) GetModel() string {
// 	return tv.Model
// }

// func (ss SamsungTV) SamsungHub() {

// }

// func (lg LgTV) LgHub() {

// }

// func main() {
// 	tvS := SamsungTV{true, "SXL-783"}
// 	tvL := LgTV{false, "LG-X75"}
// 	switchON(tvS)
// 	switchON(tvL)

// }
