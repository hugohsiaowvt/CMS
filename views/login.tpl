
<!doctype html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <meta name="description" content="">
    <meta name="author" content="Mark Otto, Jacob Thornton, and Bootstrap contributors">
    <meta name="generator" content="Jekyll v3.8.5">
    <title>Signin Template · Bootstrap</title>

    <!-- Bootstrap core CSS -->
<link href="/static/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-GJzZqFGwb1QTTN6wy59ffF1BuGJpLSa9DkKMp0DgiMDm4iYMj70gZWKYbI706tWS" crossorigin="anonymous">


    <style>
      .bd-placeholder-img {
        font-size: 1.125rem;
        text-anchor: middle;
      }

      @media (min-width: 768px) {
        .bd-placeholder-img-lg {
          font-size: 3.5rem;
        }
      }
    </style>
    <!-- Custom styles for this template -->
    <link href="/static/css/signin.css" rel="stylesheet">
  </head>
  <body class="text-center">

    <form class="form-signin" method="post">
      <img class="mb-4" src="/static/img/Logo.jpg" alt="" width="72" height="72">
      <h1 class="h3 mb-3 font-weight-normal">後台管理系統</h1>
      <label for="uname" class="sr-only">帳號</label>
      <input type="account" id="uname" name="uname" class="form-control" placeholder="帳號" value="{{.uname}}" required autofocus>
      <label for="upwd" class="sr-only">Password</label>
      <input type="password" id="upwd" name="upwd" class="form-control" placeholder="密碼" value="{{.upwd}}" required>
      {{if .isError}}
          <div class="alert alert-danger" role="alert">
            帳號或密碼錯誤
          </div>
      {{end}}
      <button class="btn btn-lg btn-primary btn-block" type="submit">登入</button>
      <p class="mt-5 mb-3 text-muted">&copy; 2019-2019</p>
    </form>
  </body>
</html>