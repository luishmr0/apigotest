package main
 
import (
    "encoding/json"
    "log"
    "net/http" 
    "github.com/gorilla/mux"
)
 
type Person struct {
    ID        string   `json:"id,omitempty"`
    Firstname string   `json:"nombre,omitempty"`
    Lastname  string   `json:"apellido,omitempty"`
    Hobby     string   `json:"hobby,omitempty"`
    Address   *Address `json:"ubicacion,omitempty"`
}
 
type Address struct {
    City  string `json:"ciudad,omitempty"`
    State string `json:"departamento,omitempty"`
}
 
var people []Person
 
func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    for _, item := range people {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&Person{})
}
 
func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
    json.NewEncoder(w).Encode(people)
}
 
func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    var person Person
    _ = json.NewDecoder(req.Body).Decode(&person)
    person.ID = params["id"]
    people = append(people, person)
    json.NewEncoder(w).Encode(people)
}
 
func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    for index, item := range people {
        if item.ID == params["id"] {
            people = append(people[:index], people[index+1:]...)
            break
        }
    }
    json.NewEncoder(w).Encode(people)
}
 
func main() {
    router := mux.NewRouter()
    people = append(people, Person{ID: "1", Firstname: "Luis", Lastname: "Marin", Hobby: "Pensar",Address: &Address{City: "Tarapoto", State: "London"}})
    people = append(people, Person{ID: "2", Firstname: "Diego", Lastname: "Maradona", Hobby: "Pelotear"})
    router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
    router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
    router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
    router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")
    log.Fatal(http.ListenAndServe("", router))
}
