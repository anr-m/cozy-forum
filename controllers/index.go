package controllers

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html/template"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"path"
	"regexp"
	"strconv"
	"time"

	"../db"
	"../errorhandle"
	"../models"
	"../sessions"
	"../tpl"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type pageData struct {
	PageTitle string
	User      models.User
	Data      interface{}
}

// Index route
func Index(w http.ResponseWriter, r *http.Request) {
	data := pageData{"Home", sessions.GetUser(w, r), nil}
	tpl.ExecuteTemplate(w, "index.html", data)
}

// Register route
func Register(w http.ResponseWriter, r *http.Request) {

	data := pageData{"Register", models.User{}, nil}

	if r.Method == http.MethodPost {
		var newUser models.User
		regex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

		newUser.Email = r.FormValue("email")
		newUser.Username = r.FormValue("username")
		newUser.FirstName = r.FormValue("firstname")
		newUser.LastName = r.FormValue("lastname")
		password := r.FormValue("password")

		if newUser.Email == "" {
			data.Data = "Email must not be empty"
			w.WriteHeader(http.StatusUnprocessableEntity)
			tpl.ExecuteTemplate(w, "register.html", data)
			return
		} else if !regex.MatchString(newUser.Email) {
			data.Data = "Invalid email"
			w.WriteHeader(http.StatusUnprocessableEntity)
			tpl.ExecuteTemplate(w, "register.html", data)
			return
		} else if db.EmailExists(newUser.Email) {
			data.Data = "Email already exists"
			w.WriteHeader(http.StatusUnprocessableEntity)
			tpl.ExecuteTemplate(w, "register.html", data)
			return
		} else if newUser.Username == "" {
			data.Data = "Username must not be empty"
			w.WriteHeader(http.StatusUnprocessableEntity)
			tpl.ExecuteTemplate(w, "register.html", data)
			return
		} else if db.UsernameExists(newUser.Username) {
			data.Data = "Username exists"
			w.WriteHeader(http.StatusUnprocessableEntity)
			tpl.ExecuteTemplate(w, "register.html", data)
			return
		} else if len(password) < 8 {
			data.Data = "Password must be at least 8 characters"
			w.WriteHeader(http.StatusUnprocessableEntity)
			tpl.ExecuteTemplate(w, "register.html", data)
			return
		} else if newUser.FirstName == "" {
			data.Data = "First Name must not be empty"
			w.WriteHeader(http.StatusUnprocessableEntity)
			tpl.ExecuteTemplate(w, "register.html", data)
			return
		} else if newUser.LastName == "" {
			data.Data = "Last Name must not be empty"
			w.WriteHeader(http.StatusUnprocessableEntity)
			tpl.ExecuteTemplate(w, "register.html", data)
			return
		}

		salt, err := uuid.NewV4()
		errorhandle.Check(err)
		hash, err := bcrypt.GenerateFromPassword([]byte(password+salt.String()), bcrypt.MinCost)
		errorhandle.Check(err)

		newUser.Hash = hash
		newUser.Salt = salt.String()

		db.CreateUser(&newUser)
		sessions.CreateSession(newUser.UserID, w)

		http.Redirect(w, r, "/index", http.StatusFound)
	} else if r.Method == http.MethodGet {
		tpl.ExecuteTemplate(w, "register.html", data)
	}
}

// Login route
func Login(w http.ResponseWriter, r *http.Request) {

	data := pageData{"Login", models.User{}, nil}

	if r.Method == http.MethodPost {
		regex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == "" {
			data.Data = "Username must not be empty"
			w.WriteHeader(http.StatusUnprocessableEntity)
			tpl.ExecuteTemplate(w, "login.html", data)
			return
		} else if password == "" {
			data.Data = "Password must not be empty"
			w.WriteHeader(http.StatusUnprocessableEntity)
			tpl.ExecuteTemplate(w, "login.html", data)
			return
		}

		if regex.MatchString(username) {
			if !db.EmailExists(username) {
				data.Data = "Invalid email"
				w.WriteHeader(http.StatusUnprocessableEntity)
				tpl.ExecuteTemplate(w, "login.html", data)
				return
			}
			user := db.GetUserByEmail(username)
			err := bcrypt.CompareHashAndPassword(user.Hash, []byte(password+user.Salt))
			if err != nil {
				data.Data = "Incorrect password"
				w.WriteHeader(http.StatusUnauthorized)
				tpl.ExecuteTemplate(w, "login.html", data)
			}
			sessions.CreateSession(user.UserID, w)
		} else {
			if !db.UsernameExists(username) {
				data.Data = "Invalid username"
				w.WriteHeader(http.StatusUnprocessableEntity)
				tpl.ExecuteTemplate(w, "login.html", data)
				return
			}
			user := db.GetUserByUsername(username)
			err := bcrypt.CompareHashAndPassword(user.Hash, []byte(password+user.Salt))
			if err != nil {
				data.Data = "Incorrect password"
				w.WriteHeader(http.StatusUnauthorized)
				tpl.ExecuteTemplate(w, "login.html", data)
			}
			sessions.CreateSession(user.UserID, w)
		}

		http.Redirect(w, r, "/index", http.StatusFound)

	} else if r.Method == http.MethodGet {
		tpl.ExecuteTemplate(w, "login.html", data)
	}
}

// Logout route
func Logout(w http.ResponseWriter, r *http.Request) {
	sessions.Logout(w, r)
	http.Redirect(w, r, "/index", http.StatusSeeOther)
}

// CreatePost route
func CreatePost(w http.ResponseWriter, r *http.Request) {
	if !sessions.IsLoggedIn(w, r) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

	data := pageData{"Create Post", sessions.GetUser(w, r), nil}

	if r.Method == http.MethodPost {
		regex := regexp.MustCompile(`^.*\.(jpg|JPG|jpeg|JPEG|gif|GIF|png|PNG)$`)
		catregex := regexp.MustCompile(`^(gaming|technology|programming|books|music)$`)

		mf, fh, _ := r.FormFile("image")
		if fh != nil {
			defer mf.Close()
		}

		category := r.FormValue("category")
		title := r.FormValue("title")
		content := r.FormValue("content")

		if category == "" {
			data.Data = "Category must not be empty"
			w.WriteHeader(http.StatusUnprocessableEntity)
			tpl.ExecuteTemplate(w, "createpost.html", data)
			return
		} else if !catregex.MatchString(category) {
			data.Data = "Invalid category"
			w.WriteHeader(http.StatusUnprocessableEntity)
			tpl.ExecuteTemplate(w, "createpost.html", data)
			return
		} else if title == "" {
			data.Data = "Title must not be empty"
			w.WriteHeader(http.StatusUnprocessableEntity)
			tpl.ExecuteTemplate(w, "createpost.html", data)
			return
		} else if content == "" {
			data.Data = "Content must not be empty"
			w.WriteHeader(http.StatusUnprocessableEntity)
			tpl.ExecuteTemplate(w, "createpost.html", data)
			return
		} else if fh != nil && fh.Size > 20000000 {
			data.Data = "File too large, please limit size to 20MB"
			w.WriteHeader(http.StatusUnprocessableEntity)
			tpl.ExecuteTemplate(w, "createpost.html", data)
			return
		} else if fh != nil && !regex.MatchString(fh.Filename) {
			data.Data = "Invalid file type, please upload jpg, jpeg, png, gif"
			w.WriteHeader(http.StatusUnprocessableEntity)
			tpl.ExecuteTemplate(w, "createpost.html", data)
			return
		}

		var encodedImage string
		if fh != nil {
			var buff bytes.Buffer
			img, ext, err := image.Decode(mf)
			errorhandle.Check(err)
			switch ext {
			case "png":
				png.Encode(&buff, img)
			case "jpg", "jpeg":
				jpeg.Encode(&buff, img, nil)
			case "gif":
				mf.Seek(0, io.SeekStart)
				gifimg, err := gif.DecodeAll(mf)
				errorhandle.Check(err)
				gif.EncodeAll(&buff, gifimg)
			}
			png.Encode(&buff, img)
			encodedImage = fmt.Sprintf(`<img src="data:image/%v;base64, %v"/>`, ext, base64.StdEncoding.EncodeToString(buff.Bytes()))
		}

		newPost := models.Post{
			UserID:      sessions.GetUser(w, r).UserID,
			Category:    category,
			Title:       title,
			Content:     content,
			HTMLImage:   template.HTML(encodedImage),
			TimeCreated: time.Now(),
		}

		db.CreatePost(&newPost)

		http.Redirect(w, r, fmt.Sprintf("/posts/id/%d", newPost.PostID), http.StatusSeeOther)

	} else if r.Method == http.MethodGet {
		tpl.ExecuteTemplate(w, "createpost.html", data)
	}
}

func GetPostByID(w http.ResponseWriter, r *http.Request) {
	dir, endpoint := path.Split(r.URL.Path)
	postid, _ := strconv.Atoi(endpoint)

	data := pageData{"Not found", sessions.GetUser(w, r), nil}

	if dir != "/posts/id/" || postid == 0 {
		NotFoundHandler(w, r)
		return
	}

	post := db.GetPostByID(postid)

	if post.PostID != postid {
		NotFoundHandler(w, r)
		return
	}

	data.PageTitle = post.Title
	data.Data = post

	tpl.ExecuteTemplate(w, "post.html", data)
}

func CreateComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Write([]byte("Invalid method"))
		return
	}

	if !sessions.IsLoggedIn(w, r) {
		w.Write([]byte("Not logged in"))
		return
	}

	dir, endpoint := path.Split(r.URL.Path)
	postid, _ := strconv.Atoi(endpoint)

	if dir != "/comment/" || postid == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Invalid path"))
		return
	}

	text := r.FormValue("text")
	user := sessions.GetUser(w, r)

	if text == "" {
		w.Write([]byte("Invalid comment"))
		return
	}

	newComment := models.Comment{
		PostID:      postid,
		UserID:      user.UserID,
		Text:        text,
		TimeCreated: time.Now(),
	}

	db.CreateComment(&newComment)

	http.Redirect(w, r, fmt.Sprintf("/posts/id/%d", newComment.PostID), http.StatusSeeOther)
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	posts := db.GetPosts()
	fmt.Fprintln(w, posts)
}

func GetPostsByCategory(w http.ResponseWriter, r *http.Request) {
	dir, category := path.Split(r.URL.Path)
	catregex := regexp.MustCompile(`^(gaming|technology|programming|books|music)$`)

	if dir != "/posts/" || !catregex.MatchString(category) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Invalid path/category"))
		return
	}

	posts := db.GetPostsByCategory(category)

	fmt.Fprintln(w, posts)
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	tpl.ExecuteTemplate(w, "notfound.html", pageData{"Not Found", sessions.GetUser(w, r), nil})
}
