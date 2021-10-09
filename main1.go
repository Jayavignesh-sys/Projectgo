package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"fmt"
    "GO_MAIN/helper"
    "GO_MAIN/models"
	"GO_MAIN/helper_user"
	"go.mongodb.org/mongo-driver/bson"
	"crypto/sha1"
	"sync"
	"time"
)

var collection = helper_user.ConnectDB()
var collection2 = helper.ConnectDB()
var lock sync.Mutex

func getPosts(w http.ResponseWriter, r *http.Request) {
	lock.Lock()
	if r.Method == "GET"{
		w.Header().Set("Content-Type", "application/json")

		var posts []models.Post

		last  := r.URL.Path[len(r.URL.Path)-3:]
		fmt.Printf(last)

		cur, err := collection2.Find(context.TODO(), bson.M{"User_id":last})

		if err != nil {
			helper.GetError(err, w)
			return
		}

		defer cur.Close(context.TODO())

		for cur.Next(context.TODO()) {

			var post models.Post
			err := cur.Decode(&post)
			if err != nil {
				log.Fatal(err)
			}

			// add item our array
			posts = append(posts, post)
		}

		if err := cur.Err(); err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(posts)
	}
	time.Sleep(1*time.Second)
	lock.Unlock()
}

func getPost(w http.ResponseWriter, r *http.Request) {
	lock.Lock()
	if r.Method == "GET"{

			fmt.Println("A GET Request given !!!")
			w.Header().Set("Content-Type", "application/json")

			var posts models.Post

			/*var params = mux.Vars(r)

			id, _ := primitive.ObjectIDFromHex(params["id"])*/

			last  := r.URL.Path[len(r.URL.Path)-3:]
			fmt.Printf(last)

			err := collection2.FindOne(context.TODO(), bson.M{"User_id":last}).Decode(&posts)

			if err != nil {
				helper.GetError(err, w)
				return
			}

			json.NewEncoder(w).Encode(posts)
		}
	time.Sleep(1*time.Second)
	lock.Unlock()
}

func createPost(w http.ResponseWriter, r *http.Request) {
	lock.Lock()
	var post models.Post

	_ = json.NewDecoder(r.Body).Decode(&post)


	fmt.Println(post.Caption)

	result,err := collection2.InsertOne(context.TODO(),post)

	if err != nil {
		helper.GetError(err, w)
		return
	}
	json.NewEncoder(w).Encode(result)
	time.Sleep(1*time.Second)
	lock.Unlock()
}

func getUser(w http.ResponseWriter, r *http.Request) {
	lock.Lock()
	if r.Method == "GET"{

			fmt.Println("A GET Request given !!!")
			w.Header().Set("Content-Type", "application/json")

			var user models.Users
			/*var params = mux.Vars(r)

			// string to primitive.ObjectID
			id, _ := primitive.ObjectIDFromHex(params["id"])*/

			//fmt.Println(id)
			/*var rNum = regexp.MustCompile("abc")
			switch{
			case rNum.MatchString(r.URL.Path):
				fmt.Println("Has digits")
			default:
				fmt.Println("Unknown pattern")
			}*/

			last  := r.URL.Path[len(r.URL.Path)-3:]
			fmt.Printf(last)

			err := collection.FindOne(context.TODO(), bson.M{"User_id":last}).Decode(&user)

			if err != nil {
				helper.GetError(err, w)
				return
			}

			json.NewEncoder(w).Encode(user)
		}
	time.Sleep(1*time.Second)
	lock.Unlock()
}

func createUser(w http.ResponseWriter, r *http.Request) {
	lock.Lock()
	if r.Method == "POST"{

			w.Header().Set("Content-Type", "application/json")

			var user models.Users

			_ = json.NewDecoder(r.Body).Decode(&user)

			fmt.Println(user.Name)

			//Hash passwords using sha Algo
			s := user.Password
			h := sha1.New()
			h.Write([]byte(s))
			bs := h.Sum(nil)
			user.Password = string(bs)

			result,err := collection.InsertOne(context.TODO(),user)


			if err != nil {
				helper.GetError(err, w)
				return
			}

			fmt.Println(result)

			json.NewEncoder(w).Encode(result)
		}
	time.Sleep(1*time.Second)
	lock.Unlock()
}

func main() {

	http.HandleFunc("/api/posts/users/", getPosts)  //The Id that we give while using a GET or POST request after '/' will be parsed and used for data retreival from mongodb
	http.HandleFunc("/api/users/", getUser)
	http.HandleFunc("/api/users", createUser)
	http.HandleFunc("/api/posts/", getPost)
	http.HandleFunc("/api/posts", createPost)


	log.Fatal(http.ListenAndServe(":8080", nil))

}
