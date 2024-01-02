package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

func (c *Course) isEmpty() bool {
	return c.CourseName == ""
}

var courseDb []Course

func main() {
	courseDb = append(courseDb, Course{CourseId: "2", CourseName: "MERN", CoursePrice: 799, Author: &Author{FullName: "Mihir Kumar", Website: "fb.com"}})
	courseDb = append(courseDb, Course{CourseId: "4", CourseName: "NextJS", CoursePrice: 999, Author: &Author{FullName: "Aaryan Bhardwaj", Website: "youtube.com"}})

	r := mux.NewRouter()
	addMiddleware := mux.NewRouter().PathPrefix("/course").Subrouter()
	addMiddleware.Use(checkDuplicate)
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getCourseById).Methods("GET")
	// r.HandleFunc("/course/add/", addCourse).Methods("POST")
	r.PathPrefix("/course").Handler(addMiddleware).Methods("POST")
	r.HandleFunc("/course/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}

func checkDuplicate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Body == nil {
			json.NewEncoder(w).Encode("empty body")
			return
		}

		var newCourse Course
		_ = json.NewDecoder(r.Body).Decode(&newCourse)
		if newCourse.isEmpty() {
			json.NewEncoder(w).Encode("name cannot be empty")
			return
		}

		for _, course := range courseDb {
			if course.CourseName == newCourse.CourseName {
				json.NewEncoder(w).Encode("course already exists")
				return
			}
		}

		next.ServeHTTP(w, r)

	})
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to this website"))

}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courseDb)
}

func getCourseById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, course := range courseDb {
		if params["id"] == course.CourseId {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("no course with such id")
	return
}

func addCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// if r.Body == nil {
	// 	json.NewEncoder(w).Encode("Provie some info")
	// }

	var newCourse Course
	_ = json.NewDecoder(r.Body).Decode(&newCourse)

	// if newCourse.isEmpty() {
	// 	json.NewEncoder(w).Encode("empty course name is not allowed")
	// }

	rand.Seed(time.Now().UnixNano())

	newCourse.CourseId = strconv.Itoa(rand.Intn(100))

	courseDb = append(courseDb, newCourse)
	json.NewEncoder(w).Encode(newCourse)
	return
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courseDb {
		if course.CourseId == params["id"] {
			courseDb = append(courseDb[:index], courseDb[index+1:]...)
			var updatedCourse Course
			updatedCourse.CourseId = params["id"]
			_ = json.NewDecoder(r.Body).Decode(&updatedCourse)
			courseDb = append(courseDb, updatedCourse)
			json.NewEncoder(w).Encode(updatedCourse)
			return
		}
	}

	json.NewEncoder(w).Encode("No course with the given ID exists")
	return
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courseDb {
		if course.CourseId == params["id"] {
			deletedCourse := courseDb[index]
			courseDb = append(courseDb[:index], courseDb[index+1:]...)
			json.NewEncoder(w).Encode(deletedCourse)
			return
		}
	}

	json.NewEncoder(w).Encode("No course with the given ID exists")
	return
}
