<!DOCTYPE html>
<html>
	<head>
	<link rel="icon" href="images/play.png" sizes="16x16" type="image/png">
	<meta charset="UTF-8">
	<style>
		body, input, input[type="text"], input[type="email"], textarea 
		{font-family:'メイリオ',Meiryo;}

		.message_post {
			margin: auto;
			border:3px solid #8AC007;
			padding: 10px;
			text-align:center;}

		.display_messages {
			text-align:center;
			/*margin: auto;
			border:3px solid #8AC007;
			padding: 10px;*/}
	</style>
	<title>Go掲示板</title>
	</head>
	<body>
		<div class="message_post">
		<h1>
			<center>Go掲示板
				<img src="images/talks.png" alt="gopher talks img"-->
			</center>
		</h1>
		<form action="post_message" method="POST">
			<label for="name">おなまえ</label></br>
			<input type="text" id="name" name="name" value="" required placeholder="おなまえ" size="40"/></br>
			<label for="name">Eメール</label></br>
			<input type="email" id="email" name="email" value="" required placeholder="Eメール" size="40"/></br>
			<label for="title">タイトル</label></br>
			<input type="text" id="title" name="title" value="" required placeholder="タイトル" size="80"/></br>
			<label for="message">メッセージ</label></br>
			<textarea id="message" name="message" cols="80" rows="5" required placeholder="メッセージを入力"></textarea></br>
			
			<input type="submit" name="submit" value="送信"/>
			<input type="reset" name="reset" value="リセット" />
		</form>
		</div>

		<div class="display_messages">
			{{range .}}
			<p>{{.Title}} 投稿者: <a href="{{.Email}}">{{.Name}}</a> 投稿日: {{.Created}}<p>
			<p>{{.Message}}</p>
			<hr>
			{{end}}
		</div>

	</body>
</html>