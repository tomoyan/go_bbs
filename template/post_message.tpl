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

		.error {color: red;}
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
		<form action="post" method="POST">
			<label for="name">おなまえ</label></br>
			{{ with .PostErrors.PostName }}
			<p class="error">{{ . }}</p>
			{{ end }}
			<input type="text" id="name" name="name" value="{{ .PostName }}" required placeholder="おなまえ" size="40"/></br>

			<label for="name">Eメール</label></br>
			{{ with .PostErrors.PostEmail }}
			<p class="error">{{ . }}</p>
			{{ end }}
			<input type="email" id="email" name="email" value="{{ .PostEmail }}" required placeholder="Eメール" size="40"/></br>

			<label for="title">タイトル</label></br>
			{{ with .PostErrors.PostTitle }}
			<p class="error">{{ . }}</p>
			{{ end }}
			<input type="text" id="title" name="title" value="{{ .PostTitle }}" required placeholder="タイトル" size="80"/></br>

			<label for="message">メッセージ</label></br>
			{{ with .PostErrors.PostMessage }}
			<p class="error">{{ . }}</p>
			{{ end }}
			<textarea id="message" name="message" cols="80" rows="5" required placeholder="メッセージを入力">{{ .PostMessage }}</textarea></br>
			
			<input type="submit" name="submit" value="送信"/>
			<input type="reset" name="reset" value="リセット" />
		</form>
		</div>

	</body>
</html>