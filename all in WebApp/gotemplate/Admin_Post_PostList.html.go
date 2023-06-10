package gotemplate

import (
	"io"
	"main/datalayer"

	"github.com/shiyanhui/hero"
)

func AdminPostListHandler(posts []datalayer.Post, w io.Writer) {
	_buffer := hero.GetBuffer()
	defer hero.PutBuffer(_buffer)
	_buffer.WriteString(`<!DOCTYPE html>

<html lang="fa">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>بلاگ</title>
    <link rel="stylesheet" href="/Static/assets/dist/css/bootstrap.min.1.css" />
    <link rel="stylesheet" href="/Static/css/customCss.css" />
    `)
	_buffer.WriteString(`
</head>

<body>
    <div class="div-mainMenu"></div>
    <div class="container">

        <div class="row mt-4">
            `)
	_buffer.WriteString(`<div class="col-12 col-md-4 ">
    <div class="card">
        <div class="card-header">
            پنل مدیریت
        </div>
        <ul class="list-group list-group-flush text-primary p-0">
             <li class="list-group-item">
                <a class="text-primary" href="/Admin/Posts"> مدیریت پست ها</a>
            </li> 
            <li class="list-group-item">
                <a class="text-primary" href="/Admin/Groups"> مدیریت گروه ها</a>
            </li>
        </ul>
    </div>
</div>
`)
	_buffer.WriteString(`
<div class="col-12 col-md-8 ">
    <p>
        <a href="/Admin/AddPost" class="btn btn-success" >پست جدید</a>
    </p>
        <table class="table">
            <thead>
                <td>تصویر</td>
                <td>عنوان پست</td>
                <td>عملیات</td>
            </thead>
            <tbody>
                `)
	for _, p := range posts {
		_buffer.WriteString(`
                <tr>
                    <td><img class="img-thumbnail img-fluid" src="/Images/Post/`)
		hero.EscapeHTML(p.Image, _buffer)
		_buffer.WriteString(`" style="height: 100px;width: 100px" ></td>
                    <td>`)
		hero.EscapeHTML(p.Title, _buffer)
		_buffer.WriteString(`</td>
                    <td>
                        <a href="/Admin/EditPost/`)
		hero.FormatUint(uint64(p.ID), _buffer)
		_buffer.WriteString(`" class="btn btn-warning">ویرایش</a> |
                    </td>
                </tr>
                `)
	}
	_buffer.WriteString(`
            </tbody>
        </table>
</div>
`)

	_buffer.WriteString(`
        </div>

    </div>
    <script src="/Static/js/jquery-3.3.1.min.js"></script>
    <script src="/Static/js/popper.min.js"></script>
    <script src="/Static/js/bootstrap.min.js"></script>
    <script src="/Static/js/bootstrap.bundle.min.js"></script>
    <script>
        $.get("/GetMenu",function (data) {
            $(".div-mainMenu").html(data)
        });
    </script>
    `)
	_buffer.WriteString(`
</body>

</html>`)
	w.Write(_buffer.Bytes())

}
