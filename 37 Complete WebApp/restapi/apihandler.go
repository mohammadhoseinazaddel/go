package restapi

import (
	"TopLearn/WebApp/datalayer"
	"TopLearn/WebApp/gotemplate"
	"TopLearn/WebApp/security/jwt"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type PersonRestApiHandler struct {
	dbhandler datalayer.SQLHandler
}

func newPersonRestApiHandler(db datalayer.SQLHandler) *PersonRestApiHandler {
	return &PersonRestApiHandler{
		dbhandler: db,
	}
}

func (handler PersonRestApiHandler) Index(w http.ResponseWriter, r *http.Request) {

	groups, err := handler.dbhandler.GetGroups()
	if err != nil {
		fmt.Println(err)
	}

	posts, err := handler.dbhandler.GetPosts()
	if err != nil {
		fmt.Println(err)
	}
	gotemplate.PostListHandler(posts, groups, w)

}

func (handler PersonRestApiHandler) SinglePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postId, ok := vars["PostId"]
	if !ok {
		fmt.Println("Post ID Not Found")
		w.WriteHeader(http.StatusBadRequest)
		gotemplate.PostHandler(datalayer.Post{}, w)
		return
	}
	intPostId, err := strconv.Atoi(postId)
	if err != nil {
		fmt.Println("Post ID Not Found")
		w.WriteHeader(http.StatusBadRequest)
		gotemplate.PostHandler(datalayer.Post{}, w)
		return
	}

	post, err := handler.dbhandler.GetPostById(intPostId)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintln(w, "داده ای یافت نشد")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	gotemplate.PostHandler(post, w)

}

func (handler PersonRestApiHandler) Groups(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	groupId, ok := vars["GroupId"]
	if !ok {
		fmt.Println("Post ID Not Found")
		w.WriteHeader(http.StatusBadRequest)
		gotemplate.PostHandler(datalayer.Post{}, w)
		return
	}
	intGroupId, err := strconv.Atoi(groupId)
	if err != nil {
		fmt.Println("group ID Not Found")
		w.WriteHeader(http.StatusBadRequest)
		gotemplate.PostHandler(datalayer.Post{}, w)
		return
	}

	post, err := handler.dbhandler.GetPostsByGroupId(intGroupId)
	if err != nil {
		fmt.Println(err)
	}
	groups, err := handler.dbhandler.GetGroups()
	if err != nil {
		fmt.Println(err)
	}
	gotemplate.PostListHandler(post, groups, w)

}

func (handler PersonRestApiHandler) Menu(w http.ResponseWriter, r *http.Request) {
	groups, err := handler.dbhandler.GetGroups()
	if err != nil {
		fmt.Println(err)
	}

	login, ok := jwt.IsLogedin(r)

	fmt.Println(login, "  ", ok)

	gotemplate.MenuHandler(groups, w)

}

func (handler PersonRestApiHandler) GetLogin(w http.ResponseWriter, r *http.Request) {
	msg := []string{}
	gotemplate.LoginHandler(msg, w)
}

func (handler PersonRestApiHandler) PostLogin(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		msg := []string{"خطایی رخ داده است."}
		gotemplate.LoginHandler(msg, w)
		return
	}

	emailArray, ok := r.Form["email"]
	if !ok || emailArray[0] == "" {
		msg := []string{"لطفا ایمیل خود را وارد نمایید."}
		gotemplate.LoginHandler(msg, w)
		return
	}

	passArray, ok := r.Form["password"]
	if !ok || passArray[0] == "" {
		msg := []string{"لطفا رمز عبور خود را وارد نمایید."}
		gotemplate.LoginHandler(msg, w)
		return
	}
	email := emailArray[0]
	password := passArray[0]
	fmt.Println(email + "  " + password)

	ok = jwt.Signin(w, jwt.LoginInfo{
		Username: email,
		Password: password,
	}, handler.dbhandler)

	if !ok {
		msg := []string{"کاربری با این مشخصات یافت نشد"}
		gotemplate.LoginHandler(msg, w)
	}

	http.Redirect(w, r, "/", 302)
}
