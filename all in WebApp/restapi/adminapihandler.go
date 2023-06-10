package restapi

import (
	"fmt"
	"main/datalayer"
	"main/gotemplate"
	"main/utility/uploader"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	guuid "github.com/google/uuid"
	"github.com/gorilla/mux"
)

type AdminRestApiHandler struct {
	dbhandler datalayer.SQLHandler
}

func newAdminRestApiHandler(db datalayer.SQLHandler) *AdminRestApiHandler {
	return &AdminRestApiHandler{
		dbhandler: db,
	}
}

func (handler AdminRestApiHandler) Dashboard(w http.ResponseWriter, r *http.Request) {
	gotemplate.AdminDashboardHandler(w)
}

func (handler AdminRestApiHandler) PostList(w http.ResponseWriter, r *http.Request) {
	posts, err := handler.dbhandler.GetPosts()
	if err != nil {
		fmt.Println(err)
	}
	gotemplate.AdminPostListHandler(posts, w)
}

func (handler AdminRestApiHandler) AddPost(w http.ResponseWriter, r *http.Request) {
	groups, err := handler.dbhandler.GetGroups()
	if err != nil {
		fmt.Println(err)
		fmt.Fprintln(w, "سرور با خطا مواجه شد.")
		return
	}
	gotemplate.AdminCreatePostHandler([]string{}, groups, w)
}

func (handler AdminRestApiHandler) PostAddPost(w http.ResponseWriter, r *http.Request) {
	//Bind Form In Struct
	post := new(datalayer.Post)
	post.Title = r.FormValue("title")
	post.ShortDesc = r.FormValue("ShortDesc")
	post.LongDesc = r.FormValue("LongDesc")
	intGroupID, _ := strconv.ParseUint(r.FormValue("GroupID"), 10, 32)
	post.GroupID = uint(intGroupID)
	// END Bind Form In Struct

	//validate Struct
	_, msgs, _ := post.Validate()
	if len(msgs) != 0 {
		groups, _ := handler.dbhandler.GetGroups()
		gotemplate.AdminCreatePostHandler(msgs, groups, w)
		return
	}
	//END validate Struct

	//Bind and Upload file
	file, header, err := r.FormFile("myFile")
	if err != nil {
		msg := []string{"خطایی رخ داده است."}
		groups, _ := handler.dbhandler.GetGroups()
		gotemplate.AdminCreatePostHandler(msg, groups, w)
		return
	}
	defer file.Close()
	filename := guuid.New().String() + filepath.Ext(header.Filename)
	path := filepath.Join("./Content/Images/Post/", filename)

	ok, err := uploader.Uploader(file, header, path)
	if !ok {
		msg := []string{"خطایی رخ داده است."}
		groups, _ := handler.dbhandler.GetGroups()
		gotemplate.AdminCreatePostHandler(msg, groups, w)
		return
	}

	post.Image = filename

	//END Bind and Upload file

	post.CreateDate = time.Now()

	//Insert in to DataBase
	err = handler.dbhandler.InsertPosts(*post)
	if err != nil {
		msg := []string{"خطایی رخ داده است."}
		groups, _ := handler.dbhandler.GetGroups()
		gotemplate.AdminCreatePostHandler(msg, groups, w)
		return
	}
	//END Insert in to DataBase

	http.Redirect(w, r, "/Admin/Posts", 302)
}
func (handler AdminRestApiHandler) EditPost(w http.ResponseWriter, r *http.Request) {
	//Get GroupId From URL
	vars := mux.Vars(r)
	postId, ok := vars["PostId"]
	if !ok {
		fmt.Println("PostId Not Found")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "PostId Not Found")
		return
	}
	//End Get GroupId From URL

	//Convert To int
	intPostId, err := strconv.Atoi(postId)
	if err != nil {
		fmt.Println("Post ID Not Found")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Post ID Not Found")
		return
	}
	//End Convert To int

	groups, err := handler.dbhandler.GetGroups()
	if err != nil {
		fmt.Println(err)
		fmt.Fprintln(w, "سرور با خطا مواجه شد.")
		return
	}

	post, err := handler.dbhandler.GetPostById(intPostId)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintln(w, "سرور با خطا مواجه شد.")
		return
	}
	gotemplate.AdminEditPostHandler([]string{}, groups, post, w)
}

func (handler AdminRestApiHandler) PostEditPost(w http.ResponseWriter, r *http.Request) {

	//Get GroupId From URL
	vars := mux.Vars(r)
	postId, ok := vars["PostId"]
	if !ok {
		fmt.Println("PostId Not Found")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "PostId Not Found")
		return
	}
	//End Get GroupId From URL

	//Convert To int
	intPostId, err := strconv.Atoi(postId)
	if err != nil {
		fmt.Println("Post ID Not Found")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Post ID Not Found")
		return
	}
	//End Convert To int

	//Bind Form In Struct
	post := new(datalayer.Post)
	post.Title = r.FormValue("title")
	post.ShortDesc = r.FormValue("ShortDesc")
	post.LongDesc = r.FormValue("LongDesc")
	intPostID, _ := strconv.ParseUint(r.FormValue("postId"), 10, 32)
	post.ID = uint(intPostID)
	post.Image = r.FormValue("oldImageName")
	intGroupID, _ := strconv.ParseUint(r.FormValue("GroupID"), 10, 32)
	post.GroupID = uint(intGroupID)
	// END Bind Form In Struct

	//validate Struct
	_, msgs, _ := post.Validate()
	if len(msgs) != 0 {
		groups, _ := handler.dbhandler.GetGroups()

		post, err := handler.dbhandler.GetPostById(intPostId)
		if err != nil {
			fmt.Println(err)
			fmt.Fprintln(w, "سرور با خطا مواجه شد.")
			return
		}
		gotemplate.AdminEditPostHandler(msgs, groups, post, w)
		return
	}
	//END validate Struct

	//Bind and Upload file
	file, header, err := r.FormFile("myFile")
	if err == nil {
		defer file.Close()
		filename := guuid.New().String() + filepath.Ext(header.Filename)
		path := filepath.Join("./Content/Images/Post/", filename)

		ok, err = uploader.Uploader(file, header, path)
		if !ok {
			msg := []string{"خطایی رخ داده است."}
			groups, _ := handler.dbhandler.GetGroups()

			post, err := handler.dbhandler.GetPostById(intPostId)
			if err != nil {
				fmt.Println(err)
				fmt.Fprintln(w, "سرور با خطا مواجه شد.")
				return
			}
			gotemplate.AdminEditPostHandler(msg, groups, post, w)
			return
		}

		post.Image = filename
	}

	//END Bind and Upload file

	post.CreateDate = time.Now()

	//Insert in to DataBase
	err = handler.dbhandler.UpdatePosts(*post)
	if err != nil {
		msg := []string{"خطایی رخ داده است."}
		groups, _ := handler.dbhandler.GetGroups()

		post, err := handler.dbhandler.GetPostById(intPostId)
		if err != nil {
			fmt.Println(err)
			fmt.Fprintln(w, "سرور با خطا مواجه شد.")
			return
		}
		gotemplate.AdminEditPostHandler(msg, groups, post, w)
		return
	}
	//END Insert in to DataBase

	http.Redirect(w, r, "/Admin/Posts", 302)
}

func (handler AdminRestApiHandler) GroupList(w http.ResponseWriter, r *http.Request) {
	groups, err := handler.dbhandler.GetGroups()
	if err != nil {
		fmt.Fprintln(w, err)
		fmt.Println(err)
		return
	}
	gotemplate.AdminGroupListHandler(groups, w)
}

func (handler AdminRestApiHandler) AddGroup(w http.ResponseWriter, r *http.Request) {
	gotemplate.AdminCreateGroupHandler([]string{}, w)
}

func (handler AdminRestApiHandler) PostAddGroup(w http.ResponseWriter, r *http.Request) {

	//Bind Form In Struct
	group := new(datalayer.Group)
	group.Title = r.FormValue("title")
	group.EnTitle = r.FormValue("enTitle")
	// END Bind Form In Struct

	//validate Struct
	_, msgs, _ := group.Validate()
	if len(msgs) != 0 {
		gotemplate.AdminCreateGroupHandler(msgs, w)
		return
	}
	//END validate Struct

	//Insert in to DataBase
	err := handler.dbhandler.InsertGroups(*group) // *Group   ===> Group
	if err != nil {
		msg := []string{"خطایی رخ داده است."}
		gotemplate.AdminCreateGroupHandler(msg, w)
		fmt.Println(err)
		return
	}
	//END Insert in to DataBase

	// groups, err := handler.dbhandler.GetGroups()
	// if err != nil {
	// 	fmt.Fprintln(w, err)
	// 	fmt.Println(err)
	// 	return
	// }
	// gotemplate.AdminGroupListHandler(groups, w)

	http.Redirect(w, r, "/Admin/Groups", 302)
}

func (handler AdminRestApiHandler) EditGroup(w http.ResponseWriter, r *http.Request) {

	//Get GroupId From URL
	vars := mux.Vars(r)
	groupId, ok := vars["GroupId"]
	if !ok {
		fmt.Println("GroupId Not Found")
		w.WriteHeader(http.StatusBadRequest)
		gotemplate.PostHandler(datalayer.Post{}, w)
		return
	}
	//End Get GroupId From URL

	//Convert To int
	intGroupId, err := strconv.Atoi(groupId)
	if err != nil {
		fmt.Println("group ID Not Found")
		w.WriteHeader(http.StatusBadRequest)
		gotemplate.PostHandler(datalayer.Post{}, w)
		return
	}
	//End Convert To int

	group, err := handler.dbhandler.GetGroupById(intGroupId)
	if err != nil {
		fmt.Fprintln(w, err)
		fmt.Println(err)
		return
	}

	gotemplate.AdminEditGroupHandler([]string{}, group, w)
}

func (handler AdminRestApiHandler) PostEditGroup(w http.ResponseWriter, r *http.Request) {
	group := new(datalayer.Group)
	group.Title = r.FormValue("title")
	group.EnTitle = r.FormValue("enTitle")

	groupId := r.FormValue("groupId")

	intGroupID, _ := strconv.ParseUint(groupId, 10, 32)
	group.ID = uint(intGroupID)
	// END Bind Form In Struct

	//validate Struct
	_, msgs, _ := group.Validate()
	if len(msgs) != 0 {
		gotemplate.AdminEditGroupHandler(msgs, *group, w)
		return
	}
	//END validate Struct

	//Update DataBase
	err := handler.dbhandler.UpdateGroup(*group)
	if err != nil {
		msg := []string{"خطایی رخ داده است."}
		gotemplate.AdminCreateGroupHandler(msg, w)
		fmt.Println(err)
		return
	}
	//END Update in to DataBase

	http.Redirect(w, r, "/Admin/Groups", 302)
}

func (handler AdminRestApiHandler) DeleteGroup(w http.ResponseWriter, r *http.Request) {

	//Get GroupId From URL
	vars := mux.Vars(r)
	groupId, ok := vars["GroupId"]
	if !ok {
		fmt.Println("GroupId Not Found")
		w.WriteHeader(http.StatusBadRequest)
		gotemplate.PostHandler(datalayer.Post{}, w)
		return
	}
	//End Get GroupId From URL

	//Convert To int
	intGroupId, err := strconv.Atoi(groupId)
	if err != nil {
		fmt.Println("group ID Not Found")
		w.WriteHeader(http.StatusBadRequest)
		gotemplate.PostHandler(datalayer.Post{}, w)
		return
	}
	//End Convert To int

	//Delete Groups
	err = handler.dbhandler.DeleteGroupById(intGroupId)
	if err != nil {
		fmt.Fprintln(w, err)
		fmt.Println(err)
		return
	}
	//End Delete Groups

	http.Redirect(w, r, "/Admin/Groups", 302)
}
