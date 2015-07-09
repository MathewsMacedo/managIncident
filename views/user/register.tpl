<h2>{{ .title2 }}{{ .ip }}</h2><br/>
{{ if not_nil .mail }}
<h2 class="text-center">Vous êtes déjà connecté</h2><br/>
<h3 class="text-center"><a href="/incident-manager/">Revenir sur la page Principale</a></h3>
{{ else }}
<div class="container">
<form class="col-md-6 col-md-offset-3" action="/incident-manager/register" method="post">
    <div class="form-group">
        <input type="email" class="form-control input-lg" placeholder="Email" name="mail">
    </div>
    <div class="form-group">
        <button class="btn btn-primary btn-lg btn-block">Envoi de la demande à un administrateur</button>
    </div>
</form>
</div>
{{ end }}
