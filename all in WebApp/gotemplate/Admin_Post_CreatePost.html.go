package gotemplate

import (
	"io"
	"main/datalayer"

	"github.com/shiyanhui/hero"
)

func AdminCreatePostHandler(messages []string, groups []datalayer.Group, w io.Writer) {
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
<div class="col-12 col-md-8">
    <div class="card">
        <div class="card-header">
            <h2>پست جدید</h2>
        </div>
        <div class="card-body">
            <form action="/Admin/AddPost" method="post" enctype="multipart/form-data">
                <div class="form-group">
                    <label class="control-label">عنوان</label>
                    <input class="form-control" name="title" />
                </div>
                <div class="form-group">
                    <label for="selectGroup">انتخاب گروه</label>
                    <select class="form-control" id="selectGroup" name="GroupID">
                        `)
	for _, value := range groups {
		_buffer.WriteString(`
                        <option value="`)
		hero.FormatUint(uint64(value.ID), _buffer)
		_buffer.WriteString(`">`)
		hero.EscapeHTML(value.Title, _buffer)
		_buffer.WriteString(`</option>
                      `)
	}
	_buffer.WriteString(`
                    </select>
                  </div>
                <div class="form-group">
                    <label class="control-label">توضیح کوتاه</label>
                    <textarea class="form-control" name="ShortDesc"></textarea>
                </div>
                <div class="form-group">
                    <label class="control-label">توضیح کامل </label>
                    <textarea class="form-control" name="LongDesc"></textarea>
                </div>
                <div class="form-group">
                    <label class="control-label">تصویر پست</label>
                    <input type="file" name="myFile"  />
                </div>
                

                <div class="form-group">
                    <input type="submit" value="ایجاد" class="btn btn-success" />
                </div>
            </form>
        </div>
        `)
	if len(messages) != 0 {
		_buffer.WriteString(`
        <div class="card-footer">
            `)
		for _, value := range messages {
			_buffer.WriteString(`
            <h6 class="text-danger">`)
			hero.EscapeHTML(value, _buffer)
			_buffer.WriteString(`</h6>
            `)
		}
		_buffer.WriteString(`
        </div>
        `)
	}
	_buffer.WriteString(`
        <div>
            <a href="/Admin/Posts" class="btn btn-primary">بازگشت به لیست</a>
        </div>
    </div>
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
