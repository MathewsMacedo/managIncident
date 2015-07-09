<h2>{{ .title2 }}</h2><br/>
{{ if not_nil .mail }}
<h2 class="text-center">Vous êtes déjà connecté</h2><br/>
<h3 class="text-center"><a href="/incident-manager/">Revenir sur la page Principale</a></h3>
{{ else }}
<div class="container">
<form class="col-md-6 col-md-offset-3" action="/incident-manager/login" method="post">
    <div class="form-group">
        <input type="text" class="form-control input-lg username" placeholder="Email" name="username">
    </div>
    <div class="form-group">
        <input type="password" class="form-control input-lg password" placeholder="Password" name="password">
    </div>
    <div class="form-group">
        <button class="btn btn-primary btn-lg btn-block">Se connecter</button>
    </div>
</form>
<h4 class="col-md-6 col-md-offset-3 text-center">Si cela est la première fois que tu es sur ce site il faut en avoir l'accord.<br/><br/> <a href="/incident-manager/register">Cliques ici pour en demander l'accès.</a></h4>
</div>
{{ end }}
