package main

import "github.com/pesedr/sofe2016a/routes"

type Person struct {
	Name  string
	Phone string
}

func main() {

	//Start Databse
	// repo.InitDB()

	// session, err := mgo.Dial("mongodb://localhost")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer session.Close()

	// c := session.DB("test").C("people")
	// p := &Person{"Ale", "123"}
	// err = c.Insert(p)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// result := Person{}
	// err = c.Find(bson.M{"name": "Ale"}).One(&result)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println("Phone: ", result.Phone)

	//Start server
	routes.InitServer()
}
