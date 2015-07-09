<h2>{{ .title2 }}</h2>
<div class="container">
  {{range $error := .errors}}
  <div class="alert alert-danger text-center" role="alert">
    <span>{{$error.Key}}  {{$error.Message}}</span>
    <button type="button" class="close" data-dismiss="alert" aria-label="Close">
    <span aria-hidden="true">&times;</span>
  </button>
  </div>
  {{end}}
<form class="form-horizontal" id="incidents" method="post">
  <div class="form-group">
    <label for="user" class="col-sm-3 control-label">Utilisateur : </label>

    <div class="col-sm-6">
      <input class="form-control" type="text" value="{{ .mail }}" name="user" disabled>
    </div>
  </div>
  <div class="form-group">
    <label for="titre" class="col-sm-3 control-label">Titre : </label>
    <div class="col-sm-6">
      <input class="form-control" type="text" name="title" maxlength="100">
    </div>
  </div>
  <div class="form-group">
    <label for="cat" class="col-sm-3 control-label">Catégorie : </label>
    <div class="col-sm-6">
      <select class="form-control" name="cat">
        <option>Informatique</option>
      	<option>Commercial</option>
      	<option>Réseau</option>
      </select>
    </div>

  </div>

  <div class="form-group">
    <label for="description" class="col-sm-3 control-label">Description : </label>
    <div class="col-sm-6">
      <textarea  class="form-control" style="resize: none;" rows="5" maxlength="400"  name="description" ></textarea>
    </div>
  </div>
  <div class="form-group">
    <label for="dateRequest" class="col-sm-3 control-label">Date de Déclaration : </label>
    <div class="col-sm-6">
      <div class='input-group date datetimepicker'>
      <input class="form-control datetimepicker" type="text" name="dateRequest" value="{{ date .dateRequest "d-m-Y H:i:s" }}">
    <span class="input-group-addon">
        <span class="glyphicon glyphicon-calendar"></span>
    </span>
    </div>
  </div>
  </div>
  <div class="form-group">
    <label for="priority" class="col-sm-3 control-label">Priorité : </label>
    <div class="col-sm-6">
      <select class="form-control" name="priority">
        <option value="4">Majeure</option>
        <option value="3">Élevée</option>
        <option value="2" selected>Normale</option>
        <option value="1">Basse</option>
      </select>
    </div>
  </div>
  <p class="col-sm-offset-3">
    <button class="btn btn-info">Déclarer l'incident</button>
  </p>
</form>
</div>
