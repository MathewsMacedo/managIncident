<!DOCTYPE html>

<html>
<head>
	<title>{{ .title }}</title>
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	<link href="/static/css/bootstrap.min.css" rel="stylesheet">
	<link href="/static/css/index.css" rel="stylesheet">
	<!-- <link rel="stylesheet" type="text/css" href="/static/css/jquery.datetimepicker.css"/ > -->
	<link rel="stylesheet" type="text/css" href="/static/css/jquery.dataTables.min.css"/ >
	<link rel="stylesheet" href="/static/bower_components/eonasdan-bootstrap-datetimepicker/build/css/bootstrap-datetimepicker.min.css" />
</head>

<body>
	<div class="container-fluid">
		{{ .navbar }}

<div class="content">
		{{if .flash.success }}
				<div class="alert alert-success alert-dismissible text-center fade hide in" role="alert">
					{{.flash.success}}
					<button type="button" class="close" data-dismiss="alert" aria-label="Close">
						<span aria-hidden="true">&times;</span>
					</button>
				</div>
		{{ else if .flash.notice }}
				<div class="alert alert-info alert-dismissible text-center fade hide in" role="alert">
			{{.flash.notice}}
			<button type="button" class="close" data-dismiss="alert" aria-label="Close">
				<span aria-hidden="true">&times;</span>
			</button>
		</div>
		{{ else if .flash.error }}
				<div class="alert alert-danger alert-dismissible text-center fade hide in" role="alert">
			{{.flash.error}}
			<button type="button" class="close" data-dismiss="alert" aria-label="Close">
				<span aria-hidden="true">&times;</span>
			</button>
		</div>
		{{ else if .flash.warning }}
				<div class="alert alert-warning alert-dismissible text-center fade hide in" role="alert">
			{{.flash.warning}}
					<button type="button" class="close" data-dismiss="alert" aria-label="Close">
						<span aria-hidden="true">&times;</span>
					</button>
				</div>
		{{ end }}
	{{ .LayoutContent}}
	{{ .footer }}
	{{ .Username}}
	{{ .Email}}
	{{ .Count}}
</div>

	</div>
</body>
</html>
