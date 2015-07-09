<h2>{{ .title2 }} : {{ .title }}</h2>

{{range $error := .errors}}
    {{$error.Key}}
    {{$error.Message}}
{{end}}

<div class="container">
<form class="form-horizontal" action="/incident-manager/{{ if not_nil .admin }}{{ .admin }}{{ else }}user/{{ end }}incident/update/{{ .id }}" method="post">
  <div class="form-group">
    <label for="user" class="col-sm-4 control-label">Utilisateur : </label>
    <div class="col-sm-8">
      <input class="form-control" type="text" value="{{ .user.Mail }}" name="userId" disabled>
    </div>
  </div>
  <div class="form-group">
    <label for="titre" class="col-sm-4 control-label">Titre : </label>
    <div class="col-sm-8">
      <input class="form-control" type="text" value="{{ .title }}" name="tiltle" maxlength="100" disabled>
    </div>
  </div>
  <div class="form-group">
    <label for="cat" class="col-sm-4 control-label">Catégorie : </label>
    <div class="col-sm-8">
      <select class="form-control" name="cat" {{ if compare_not .role "admin"}}disabled="disabled"{{ end }}>
        <option value="Informatique" {{ if compare .cat "Informatique"}}selected{{ end }}>Informatique</option>
        <option value="Commercial" {{ if compare .cat "Commercial"}}selected{{ end }}>Commercial</option>
        <option value="Réseau" {{ if compare .cat "Réseau"}}selected{{ end }}>Réseau</option>
      </select>
    </div>
  </div>
  <div class="form-group">
    <label for="description" class="col-sm-4 control-label">Description : </label>
    <div class="col-sm-8">
      <textarea  class="form-control" style="resize: none;" rows="5" maxlength="400" {{ if compare_not .role "admin"}}disabled="disabled"{{ end }} name="description">{{ .description }}</textarea>
    </div>
  </div>
  <div class="form-group">
    <label for="resolv" class="col-sm-4 control-label">Résolution : </label>
    <div class="col-sm-8">
      <textarea  class="form-control" style="resize: none;" rows="5" maxlength="400" {{ if compare_not .role "admin"}}disabled="disabled"{{ end }} name="resolution">{{ .resolv }}</textarea>
    </div>
  </div>
  <div class="form-group">
    <label for="dateRequest" class="col-sm-4 control-label">Date de Déclaration : </label>
    <div class="col-sm-8">
      <input class="form-control" type="text" value="{{ date .dateRequest "d-m-Y H:i:s" }}" name="dateRequest" disabled>
    </div>
  </div>
  <div class="form-group">
    <label for="dateEstimated" class="col-sm-4 control-label">Date estimée de résolution : </label>
    <div class="col-sm-{{ if compare .role "admin" }}7{{ else }}8{{ end }}">
      <div class='input-group date datetimepicker dtp1'>
      {{ if compare .dateEstimated "0001-01-01 00:00:00 +0000 UTC" }}
      <input class="form-control datetimepicker" type="text" placeholder="Pas encore pris en charge" {{ if compare_not .role "admin"}}disabled{{ end }} name="dateEstimated">
      {{ else }}
      <input class="form-control datetimepicker inputRemove1" type="text" value="{{ date .dateEstimated "d-m-Y H:i:s" }}" {{ if compare_not .role "admin"}}disabled{{ end }} name="dateEstimated">
      {{ end }}
      <span class="input-group-addon">
          <span class="glyphicon glyphicon-calendar"></span>
      </span>
    </div>
    </div>
    {{ if compare .role "admin" }}<button type="button" class="btn btn-default"><span class="glyphicon glyphicon-remove dateRemove1"></span></button>{{ end }}
  </div>
  <div class="form-group">
    <label for="dateResolution" class="col-sm-4 control-label">Date de résolution : </label>
    <div class="col-sm-{{ if compare .role "admin" }}7{{ else }}8{{ end }}">
      <div class='input-group date datetimepicker dtp2'>
      {{ if compare .dateResolution "0001-01-01 00:00:00 +0000 UTC" }}
      <input class="form-control" type="text" placeholder="Pas encore totalement Résolu" {{ if compare_not .role "admin"}}disabled{{ end }} name="dateResolution" disabled>
      {{ else }}
      <input class="form-control inputRemove2" type="text" value="{{ date .dateResolution "d-m-Y H:i:s" }}" {{ if compare_not .role "admin"}}disabled{{ end }} name="dateResolution">
      {{ end }}
      <span class="input-group-addon">
          <span class="glyphicon glyphicon-calendar"></span>
      </span>
    </div>
  </div>
  {{ if compare .role "admin" }}<button type="button" class="btn btn-default"><span class="glyphicon glyphicon-remove dateRemove2"></span></button>{{ end }}
  </div>
  <div class="form-group">
    <label for="priority" class="col-sm-4 control-label">Priorité : </label>
    <div class="col-sm-8">
      <select class="form-control" name="priority" value="{{ .priority }}" {{ if compare_not .role "admin"}}disabled{{ end }}>
        <option value="4" {{ if compare .priority "4"}}selected{{ end }}>Majeure</option>
        <option value="3" {{ if compare .priority "3"}}selected{{ end }}>Élevée</option>
        <option value="2" {{ if compare .priority "2"}}selected{{ end }}>Normale</option>
        <option value="1" {{ if compare .priority "1"}}selected{{ end }}>Basse</option>
      </select>
    </div>
  </div>
    <div class="form-group">
      <label for="confirmUser" class="col-sm-4 control-label">Confirmation Utilisateur de la résolution : </label>
      <div class="col-sm-8">
        <select class="form-control" name="confirmUser" value="{{ .confirmUser }}"
        {{ if compare .dateResolution "0001-01-01 00:00:00 +0000 UTC" }}
          disabled
        {{ else if compare_not .mail .user.Mail}}
            disabled
        {{ end }}>
          <option value="0" {{ if compare .confirmUser "0"}}selected{{ end }}></option>
          <option value="1" {{ if compare .confirmUser "1"}}selected{{ end }}>Incident Résolu</option>
          <option value="2" {{ if compare .confirmUser "2"}}selected{{ end }}>Un peu mieux mais encore quelques soucis</option>
          <option value="3" {{ if compare .confirmUser "3"}}selected{{ end }}>Toujours le même incident</option>
        </select>
      </div>
      </div>
      <a href="/incident-manager/admin/"><button type="button" class="btn btn-info col-sm-2">Retour</button></a>
    {{ if compare .role "admin" }}
      <a href="#myModalIncident-{{ .id }}" data-toggle="modal"><button type="button" class="btn btn-default col-sm-2 col-sm-offset-3">Mettre à jour</button></a>
      <a href="#myModalIncidentDelete-{{ .id }}" data-toggle="modal"><button type="button" class="btn btn-default col-sm-2 col-sm-offset-3">Supprimer</button></a>
    {{ end }}
    {{ if compare_not .dateResolution "0001-01-01 00:00:00 +0000 UTC" }}
      {{ if compare .role "user"}}
        {{ if compare .mail .user.Mail}}
          <a href="#myModalIncident-{{ .id }}" data-toggle="modal"><button type="button" class="btn btn-default col-sm-2 col-sm-offset-3">Mettre à jour</button></a>
          {{ end }}
          {{ end }}
    {{ end }}
    <!-- Modal window -->
    <div id="myModalIncident-{{ .id }}" class="modal fade">
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-header">
                    <button type="button" class="close" data-dismiss="modal" aria-hidden="true">&times;</button>
                    <h3 class="modal-title text-center">Modification de l'incident : {{ .title }}</h3>
                </div>
                <div class="modal-body">
                    <h3 class="text-center">Veux-tu vraiment mettre à jour l'incident ?</h3>
                </div>
                <div class="modal-footer">
                  <a href="/incident-manager/admin/"><button type="button" class="btn btn-default pull-left">Retour</button></a>
                    <button type="submit" class="btn btn-primary">Enregistrer les changements</button>
                    <button type="button" class="btn btn-default" data-dismiss="modal">Fermer</button>
                </div>
            </div>
        </div>
    </div>
    {{ if compare .role "admin" }}
    <div id="myModalIncidentDelete-{{ .id }}" class="modal fade">
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-header">
                    <button type="button" class="close" data-dismiss="modal" aria-hidden="true">&times;</button>
                    <h3 class="modal-title text-center">Suppression de l'incident : {{ .title }}</h3>
                </div>
                <div class="modal-body">
                      <h3 class="text-center text-uppercase">Es-tu sûr de vouloir effacer cet incident?</h3>
                </div>
                <div class="modal-footer">
                  <a href="/incident-manager/admin/"><button type="button" class="btn btn-default pull-left">Retour</button></a>
                    <a href="/incident-manager/admin/incident/delete/{{ .id }}"><button type="button" class="btn btn-danger">Supprimer l'incident</button></a>
                    <button type="button" class="btn btn-default" data-dismiss="modal">Fermer</button>
                </div>
            </div>
        </div>
    </div>
    {{ end }}
    <!-- Modal End -->
</form>
