<DOCTYPE html>
<html>
<head>
    <title>DashBoard</title>
</head>
<body>
    <table border="1" style="width:100%;">
        {{ range.students }}
        <tr>
            <td>{{.ID}}</td>
            <td>{{.FirstName}}</td>
            <td>{{.LastName}}</td>
        </tr>
        {{end}}
    </table>
</body>
</html>
