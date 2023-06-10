// Code generated by hero.
// source: E:\GOLANG\src\TopLearn\WebApp\template\PostList.html
// DO NOT EDIT!
package gotemplate

import (
	"TopLearn/WebApp/datalayer"
	"TopLearn/WebApp/utility/time"
	"fmt"
	"io"

	"github.com/shiyanhui/hero"
)

func PostListHandler(posts []datalayer.Post, groups []datalayer.Group, w io.Writer) {
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
	_buffer.WriteString(`
    <div class="col-12 col-md-4 ">
        <div class="card">
            <div class="card-header">
                دسته بندی ها
            </div>
            <ul class="list-group list-group-flush text-primary p-0">
                `)
	for _, g := range groups {
		_buffer.WriteString(`
                    <li class="list-group-item">
                        <a class="text-primary" href="/Group/`)
		hero.FormatUint(uint64(g.ID), _buffer)
		_buffer.WriteString(`"> `)
		hero.EscapeHTML(g.Title, _buffer)
		_buffer.WriteString(`</a>
                    </li>
                `)
	}
	_buffer.WriteString(`
            </ul>
        </div>
    </div>
    <div class="col-12 col-md-8 ">
        `)
	for _, p := range posts {
		_buffer.WriteString(`
        <div id="content">
            <div class="card mb-4 ">
                <div class="card-body text-primary">
                    <h4> `)
		hero.EscapeHTML(p.Title, _buffer)
		_buffer.WriteString(`</h4>
                    <hr>
                    <img class="img-fluid img-thumbnail mx-auto d-block" src="/Images/Post/`)
		hero.EscapeHTML(p.Image, _buffer)
		_buffer.WriteString(`" alt="">
                    <p class="text-medium text-justify">
                    `)
		hero.EscapeHTML(p.ShortDesc, _buffer)
		_buffer.WriteString(`    
                    </p>
                    <hr>
                    <div class="row">
                        <div class="text-medium col">
                            <!-- <span>`)
		hero.EscapeHTML(fmt.Sprintf("%v", p.CreateDate), _buffer)
		_buffer.WriteString(` </span> -->
                            <span>`)
		hero.EscapeHTML(time.ToPersian(p.CreateDate), _buffer)
		_buffer.WriteString(` </span>
                        </div>
                        <div class="col">
                            <a href="/Post/`)
		hero.FormatUint(uint64(p.ID), _buffer)
		_buffer.WriteString(`" class="btn btn-sm btn-outline-primary float-left">ادامه مطلب</a>
                        </div>

                    </div>

                </div>
            </div>
        </div>
        `)
	}
	_buffer.WriteString(`
        <ul class="pagination mx-auto text-primary">

            <li class="page-item active">
                <a class="page-link" asp-action="Index" asp-route-pageId="1">1</a>
            </li>

            <li class="page-item">
                <a class="page-link" asp-action="Index" asp-route-pageId="@(Model.PageCount)">2</a>
            </li>
            <li class="page-item">
                <a class="page-link" asp-action="Index" asp-route-pageId="@(Model.PageCount)">3</a>
            </li>
            <li class="page-item">
                <a class="page-link" asp-action="Index" asp-route-pageId="@(Model.PageCount)">4</a>
            </li>

        </ul>

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