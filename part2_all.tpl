<!DOCTYPE html>
<html lang="ja">
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <!-- The above 3 meta tags *must* come first in the head; any other head content must come *after* these tags -->
    <meta name="description" content="">
    <meta name="author" content="">
    <link rel="icon" href="../../favicon.ico">

    <title>Message AP</title>

    <!-- Bootstrap core CSS -->
    <link href="/static/css/bootstrap.min.css" rel="stylesheet">
  </head>

  <body>
    <div class="container">
      <div class="header clearfix">
        <nav>
          <ul class="nav nav-pills pull-right">
            <li role="button"><a href="#">MyMessages</a></li>
            <li role="button"><a href="#">Settings</a></li>
            <li role="button"><a href="#" id="logout" onClick="alert(0);">Logout</a></li>
          </ul>
        </nav>
        <h3 class="text-primary">Message AP</h3>
      </div>

      <div class="row marketing">
        <div class="col-lg-4 well" style="background-color:#BFE0EC;">
          <h4>Name</h4>
          <div style="background:#ffeeee; padding:10px; border:1px solid #ff0000; border-radius:20px;"> {{ .User }} </div>

          <h4>Messages</h4>
          <a href="/"><span class="badge"> {{ .Message }} </span></a>


          <h4>follow</h4>
          <a href="/"><span class="badge">{{ .Follow }}</span></a>

          <h4>follower</h4>
					<a href="/"><span class="badge">{{ .Follower }}</span></a>
        </div>

        <div class="col-lg-8">
          <div class="well" style="background-color:#E5F2F7;">
            <form method="POST" action="/chapter1/write" >
              <div class="form-group">
                <textarea class="form-control" placeholder="Post a Message" rows="3" name="message"></textarea>
              </div>
              <div align="right">
                <button type="submit" class="btn btn-info" id="post-button">Post</button>
              </div>
            </form>
          </div>
          {{if .Messages_line}}
            <ul>
            {{range $message := .Messages_line}}
              <li> {{ $message.message }} ({{ $message.created_at }})</li>
            {{end}}
            </ul>
          {{end}}
        </div>
      </div>

      <footer class="footer">
        <p>&copy; Company 2015</p>
      </footer>

    </div> <!-- /container -->
		<script src="/static/js/jquery.min.js"></script>
    <script src="/static/js/bootstrap.min.js"></script>
    <script>
        $("#post-button").click(function(){
            $.ajax({
                method: "GET",
                url: "/messages",
                dataType: 'json',
            })
            .done(function( msg ) {
                var nesting = $("#ul").append('<li style="border-style:solid;border-width: thin;border-color:#e3e3e3;">' + msg.username + '</li>' );
            });
        });
    </script>
  </body>
</html>
