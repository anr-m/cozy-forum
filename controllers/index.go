package controllers

import (
	"fmt"
	"net/http"
	"path"
	"regexp"
	"strconv"
	"time"

	"../db"
	"../errorhandle"
	"../models"
	"../sessions"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// Index route
func Index(w http.ResponseWriter, r *http.Request) {
	if sessions.IsLoggedIn(w, r) {
		user := sessions.GetUser(w, r)
		w.Write([]byte("Welcome " + user.FirstName + " " + user.LastName))
	} else {
		w.Write([]byte("You are not logged in"))
	}
}

// Register route
func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var newUser models.User
		regex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

		newUser.Email = r.FormValue("email")
		newUser.Username = r.FormValue("username")
		newUser.FirstName = r.FormValue("firstname")
		newUser.LastName = r.FormValue("lastname")
		password := r.FormValue("password")

		if newUser.Email == "" {
			w.Write([]byte("Email must not be empty"))
			return
		} else if !regex.MatchString(newUser.Email) {
			w.Write([]byte("Enter valid email"))
			return
		} else if db.EmailExists(newUser.Email) {
			w.Write([]byte("Email exists"))
			return
		} else if newUser.Username == "" {
			w.Write([]byte("Username must not be empty"))
			return
		} else if db.UsernameExists(newUser.Username) {
			w.Write([]byte("Username exists"))
			return
		} else if len(password) < 8 {
			w.Write([]byte("Password must be at least 8 characters"))
			return
		} else if newUser.FirstName == "" {
			w.Write([]byte("First Name must not be empty"))
			return
		} else if newUser.LastName == "" {
			w.Write([]byte("Last Name must not be empty"))
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

		w.Write([]byte("Successfully registered"))
	} else if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`<form method="POST">
			<input type="text" name="username" placeholder="username" required><br>
			<input type="email" name="email" placeholder="email" required><br>
			<input type="password" name="password" placeholder="password" required><br>
			<input type="text" name="firstname" placeholder="first name" required><br>
			<input type="text" name="lastname" placeholder="last name" required><br>
			<button type="submit">Submit</button>
		</form>`))
	}
}

// Login route
func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		regex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == "" {
			w.Write([]byte("Username must not be empty"))
			return
		} else if password == "" {
			w.Write([]byte("Password must not be empty"))
			return
		}

		if regex.MatchString(username) {
			if !db.EmailExists(username) {
				w.Write([]byte("Invalid email"))
				return
			}
			user := db.GetUserByEmail(username)
			err := bcrypt.CompareHashAndPassword(user.Hash, []byte(password+user.Salt))
			errorhandle.Check(err)
			sessions.CreateSession(user.UserID, w)
		} else {
			if !db.UsernameExists(username) {
				w.Write([]byte("Invalid username"))
				return
			}
			user := db.GetUserByUsername(username)
			err := bcrypt.CompareHashAndPassword(user.Hash, []byte(password+user.Salt))
			errorhandle.Check(err)
			sessions.CreateSession(user.UserID, w)
		}

		http.Redirect(w, r, "/", http.StatusFound)

	} else if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`<form method="POST">
			<input type="text" name="username" placeholder="username or email" required><br>
			<input type="password" name="password" placeholder="password" required><br>
			<button type="submit">Submit</button>
		</form>`))
	}
}

// CreatePost route
func CreatePost(w http.ResponseWriter, r *http.Request) {
	if !sessions.IsLoggedIn(w, r) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

	if r.Method == http.MethodPost {
		regex := regexp.MustCompile(`^.*\.(jpg|JPG|jpeg|JPEG|gif|GIF|png|PNG|svg|SVG)$`)
		catregex := regexp.MustCompile(`^(gaming|technology|programming|books|music)$`)

		mf, fh, _ := r.FormFile("image")
		if fh != nil {
			defer mf.Close()
		}

		category := r.FormValue("category")
		title := r.FormValue("title")
		content := r.FormValue("content")

		if category == "" {
			w.Write([]byte("Category must not be empty."))
			return
		} else if !catregex.MatchString(category) {
			w.Write([]byte("Invalid category."))
			return
		} else if title == "" {
			w.Write([]byte("Title must not be empty."))
			return
		} else if content == "" {
			w.Write([]byte("Content must not be empty."))
			return
		} else if fh != nil && fh.Size > 20000000 {
			w.Write([]byte("File too large. Please limit size to 20MB."))
			return
		} else if fh != nil && !regex.MatchString(fh.Filename) {
			w.Write([]byte("Invalid file type. Please upload jpg, jpeg, png, gif, svg"))
			return
		}

		var bytes []byte
		if fh != nil {
			bytes = make([]byte, fh.Size)
			mf.Read(bytes)
		}

		newPost := models.Post{
			UserID:      sessions.GetUser(w, r).UserID,
			Category:    category,
			Title:       title,
			Content:     content,
			Image:       bytes,
			TimeCreated: time.Now(),
		}

		db.CreatePost(&newPost)

		http.Redirect(w, r, fmt.Sprintf("/posts/id/%d", newPost.PostID), http.StatusSeeOther)

	} else if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`<form method="POST" enctype="multipart/form-data">
			<input type="text" name="category" placeholder="category" required><br>
			<input type="text" name="title" placeholder="title" required><br>
			<input type="text" name="content" placeholder="content" required><br>
			<input type="file" name="image" accept=".jpg,.JPG,.png,.PNG,.gif,.GIF,.svg,.SVG"><br>
			<button type="submit">Submit</button>
		</form>`))
	}
}

func GetPostByID(w http.ResponseWriter, r *http.Request) {
	dir, endpoint := path.Split(r.URL.Path)
	postid, _ := strconv.Atoi(endpoint)

	if dir != "/posts/id/" || postid == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Not Found"))
		return
	}

	post := db.GetPostByID(postid)

	if post.PostID != postid {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Not Found"))
		return
	}

	fmt.Fprintln(w, post)
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

}
