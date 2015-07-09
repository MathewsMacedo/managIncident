<h2>{{ .mail }}</h2>
<h2>{{ .title2 }}</h2><br/>

<div class="container">
<form class="col-md-6 col-md-offset-3 form-horizontal formPassLogin" action="/mail/confirmation/{{ .md5Mail }}" method="POST">
    <div class="form-group pass">
        <input type="password" class="form-control input-lg password pwd" placeholder="Mot de passe" name="password">
    </div>
    <div class="form-group repass">
        <input type="password" class="form-control input-lg password repwd" placeholder="Répéte le mot de passe" name="repassword">
    </div>
    <div class="form-group">
        <button class="btn btn-primary btn-lg btn-block">Se connecter</button>
    </div>
</form>
