package views

import (
    "bytes"
    "html/template"
)

const tpl = `
<!DOCTYPE html>
<html>
    <head>
        <!-- Required meta tags -->
        <meta charset="utf-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    
        <!-- Bootstrap CSS -->
        <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css" integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T" crossorigin="anonymous">
    
        <title>Quickstart Web Stacks</title>
    </head>
    <body>
        <a href="https://github.com/khatribharat/quickstart-web-stacks">
            <img style="position: absolute; top: 0; right: 0; border: 0;" src="https://s3.amazonaws.com/github/ribbons/forkme_right_orange_ff7600.png" alt="Fork me on GitHub">
        </a>
        <div class="jumbotron">
          <h1 class="display-4">Quickstart Web Stack for Go.</h1>
          <p class="lead">Quickstart Go template to launch your next web application. It supports HTTP/2, WebSocket, uses asynchronous I/O and can be run on Linux and FreeBSD.</p>
          <hr class="my-4">
          <p>It uses Go standard library as its web framework, application server and web server.</p>
          <a class="btn btn-primary btn-lg" href="https://github.com/khatribharat/quickstart-web-stacks" role="button">Learn more</a>
        </div>
        <div>Generated at {{.CurrentTime}}</div>
    </body>
</html>
`

func GenerateHTMLForHome(data interface{}) []byte {
    t, err := template.New("home").Parse(tpl)
    if err != nil {
        panic(err)
    }
    w := new(bytes.Buffer)
    err = t.Execute(w, data)
    if err != nil {
        panic(err)
    }
    return w.Bytes()
}
