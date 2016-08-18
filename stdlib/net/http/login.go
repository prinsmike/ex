package main

import (
	"code.google.com/p/go.crypto/bcrypt"
	"html/template"
	"labix.org/v2/mgo/bson"
)

type User struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Username string
	Password []byte
}

var router *mux.Router

func main() {
	router = mux.NewRouter()
}

func (u *User) SetPassword(password string) {
	hpass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	u.Password = hpass
}

func Login(ctx *Context, username, password string) (u *User, err error) {
	err = ctx.C("users").Find(bson.M{"username": username}).One(&u)
	if err != nil {
		return
	}

	err = bcrypt.CompareHashAndPassword(u.Password, []byte(password))
	if err != nil {
		u = nil
	}
	return
}

func loginForm(w http.ResponseWriter, req *http.Request, ctx *Context) (err error) {
	return T("login.html").Execute(w, map[string]interface{}{
		"ctx": ctx,
	})
}

func login(w http.ResponseWriter, req *http.Request, ctx *Context) (err error) {
	// Grab the username and password from the form.
	username, password := req.FormValue("username"), req.FormValue("password")

	// Log the user in.
	user, err := Login(ctx, username, password)
	if err != nil {
		ctx.Session.AddFlash("Invalid Username or Password")
		return loginForm(w, req, ctx)
	}

	// Store the user id in the values and redirect to index.
	ctx.Session.Values["user"] = user.ID
	http.Redirect(w, req, reverse("index"), http.StatusSeeOther)
	return nil
}
