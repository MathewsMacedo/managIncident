<nav class="navbar navbar-default navbar-fixed-top">
	<div class="container-fluid">
		<!-- Brand and toggle get grouped for better mobile display -->
		<div class="navbar-header">

			<button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#bs-example-navbar-collapse-1">
				<span class="sr-only">Mobile navigation</span>
				<span class="icon-bar"></span>
				<span class="icon-bar"></span>
				<span class="icon-bar"></span>
			</button>
			<a class="navbar-brand" href="/incident-manager/{{ .admin }}">{{ if compare .role "admin"}}Administration : {{ end }}Incident Manager</a>
		</div>

		<!-- Collect the nav links, forms, and other content for toggling -->
		<div class="collapse navbar-collapse" id="bs-example-navbar-collapse-1">
			<ul class="nav navbar-nav">
				<li class="dropdown">
          <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-expanded="false">Incidents <span class="badge">{{ .badgeIncident }}</span><span class="caret"></span></a>
          <ul class="dropdown-menu" role="menu">
            <li><a href="/incident-manager/{{ .admin }}">Voir tous les incidents</a></li>
            <li class="divider"></li>
            <li><a href="/incident-manager/{{ if not_nil .admin }}{{ .admin }}{{ else }}user/{{ end }}declaration">Déclarer un incident</a></li>
          </ul>
        </li>
				<li class="dropdown pull-right">
          <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-expanded="false">
						{{ if not_nil .mail }}
							{{ .mail }}
						{{ else }}
							Pas connecté
							{{ end }}
						<span class="caret"></span></a>
          <ul class="dropdown-menu" role="menu">

            <li><a href="/incident-manager/{{ if not_nil .admin }}{{ .admin }}{{ else }}user/{{ end }}myincident">Voir mes incidents déclarés</a></li>
            <li class="divider"></li>
						{{ if not_nil .mail  }}
            <li><a href="/incident-manager/logout">Se déconnecter</a></li>
						{{ else }}
						<li><a href="/incident-manager/login">Se connecter</a></li>
            <li class="divider"></li>
						<li><a href="/incident-manager/register">Demande de connexion / Nouveau mot de passe</a></li>
						{{ end }}
          </ul>

        </li>
				{{ if not_nil .mail }}
					{{ if not_nil .admin }}
						<li class="dropdown">
							<a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-expanded="false">Utilisateurs <span class="badge">{{ .badgeUser }}</span><span class="caret"></span></a>
							<ul class="dropdown-menu" role="menu">
								<li><a href="/incident-manager/admin/user">Tous les utilisateurs</a></li>
								<li class="divider"></li>
								<li><a href="/incident-manager/admin/user/#myModal-Add" class="addUser" >Ajouter un utilisateur</a></li>
							</ul>
						</li>
						<li class="dropdown">
							<a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-expanded="false">Demande <span class="badge">{{ .badgeDemand }}</span><span class="caret"></span></a>
							<ul class="dropdown-menu" role="menu">
								<li><a href="/incident-manager/admin/register/" class="register">Verifier demande utilisateur</a></li>
							</ul>
						</li>
						{{ end }}
					{{ end }}
			</ul>
			<!-- {{ if compare .role "admin"}}<a href="/incident-manager/admin"><button type="button" class="btn btn-info navbar-btn">Admin</button></a>{{ end }} -->
		</div><!-- /.navbar-collapse -->
	</div><!-- /.container-fluid -->
</nav>
