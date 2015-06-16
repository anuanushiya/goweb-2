<html>
<head>
<title></title>
</head>
<body>
<form action="/login" method="post">
<input type="checkbox" name="interest" value="football">football
<input type="checkbox" name="interest" value="basketball">basketball
<input type="checkbox" name="interest" value="tennis">tennis
    username: <input type="text" name="username">
    password: <input type="password" name="password">
    <input type="hidden" name="token" value="{{.}}">
    <input type="submit" value="Login">
</form>
</body>
</html>
