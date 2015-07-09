<h2>{{ .title2 }}</h2>
<div class="row">

  <table class="table table-bordered table-hover table-responsive table-user">
    <thead>
    <th class="text-center">Id</th>
    <th class="text-center">Mail</th>
    <th class="text-center">Role</th>
    <th class="text-center">Password</th>
    <th class="text-center">Date Création</th>
    <th class="text-center">Supprimer / Modifier</th>
  </thead>
    {{ range $u := .users }}
    <tr class="admin-user vertical-align">
      <td>{{ $u.Id }}</td>
      <td>{{ $u.Mail }}</td>
      <td>{{ $u.Role }}</td>
      <td>{{ if compare_not $u.Pass ""}}********{{ end }}</td>
      <td>{{ date $u.Created "d-m-Y H:i"}}</td>
      <td class="text-center">
        <a href="#myModal-{{ $u.Id }}" data-toggle="modal"><span class="glyphicon glyphicon-pencil glyphicon-lg"></span></a>
        &nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
        <a href="/incident-manager/admin/user/delete/{{ $u.Id }}" title="Delete"><span class="glyphicon glyphicon-minus glyphicon-lg"></span></a>

        <!-- Modal window -->
        <div id="myModal-{{ $u.Id }}" class="modal fade">
          <form class="form-horizontal" action="/incident-manager/admin/user/update/{{ $u.Id }}" method="post">
            <div class="modal-dialog">
                <div class="modal-content">
                    <div class="modal-header">
                        <button type="button" class="close" data-dismiss="modal" aria-hidden="true">&times;</button>
                        <h4 class="modal-title">Modification de l'utilisateur : {{ $u.Mail }} </h4>
                    </div>
                    <div class="modal-body">
                          <div class="form-group">
                            <label for="Id" class="col-sm-3 control-label">Id</label>
                            <div class="col-sm-9">
                              <input type="text" class="form-control" id="Id" value="{{ $u.Id }}" disabled>
                            </div>
                          </div>
                          <div class="form-group">
                            <label for="inputEmail" class="col-sm-3 control-label">Email</label>
                            <div class="col-sm-9">
                              <input type="mail" class="form-control" id="inputEmail" value="{{ $u.Mail }}" disabled>
                            </div>
                          </div>
                          <div class="form-group">
                            <label for="inputRole" class="col-sm-3 control-label">Role</label>
                            <div class="col-sm-9">
                              <select class="form-control" name="role" id="inputRole">
                                <option value="" {{ if compare $u.Role ""}}selected{{ end }}></option>
                                <option value="admin" {{ if compare $u.Role "admin"}}selected{{ end }}>admin</option>
                              	<option value="user" {{ if compare $u.Role "user"}}selected{{ end }}>user</option>
                              </select>
                            </div>
                          </div>
                          <div class="form-group">
                            <label for="inputPassword" class="col-sm-3 control-label">Mot de Passe</label>
                            <div class="col-sm-9">
                              <input type="password" class="form-control" id="inputPassword" value="{{ $u.Pass }}" disabled>
                            </div>
                          </div>
                          <div class="form-group">
                            <label for="inputCreate" class="col-sm-3 control-label">Date de Création</label>
                            <div class="col-sm-9">
                              <input type="mail" class="form-control" id="inputCreate" value="{{ $u.Created }}" disabled>
                            </div>
                          </div>

                    </div>
                    <div class="modal-footer">
                        <button type="submit" class="btn btn-primary">Enregistrer les changements</button>
                        <button type="button" class="btn btn-default" data-dismiss="modal">Fermer</button>
                    </div>
                </div>
            </div>
          </form>
        </div>

        <!-- Modal End -->
      </td>
   </tr>
   {{ end }}
 </table>
<br/>
 <button type="button" class="btn btn-primary btn-lg col-xs-12" data-toggle="modal" data-target="#myModal-Add">
  Ajouter un Utilisateur
</button>

</div>


<!-- modal add -->
<div class="modal fade" id="myModal-Add" tabindex="-1" role="dialog" aria-labelledby="" aria-hidden="true">
<form class="form-horizontal" action="/incident-manager/admin/user/add" method="post">
  <div class="modal-dialog">
    <div class="modal-content">
      <div class="modal-header">
        <button type="button" class="close" data-dismiss="modal" aria-hidden="true">&times;</button>
        <h4 class="modal-title text-center">Ajouter un utilisateur</h4>
      </div>
      <div class="modal-body">
            <div class="form-group">
              <label for="Id" class="col-sm-3 control-label">Id</label>
              <div class="col-sm-9">
                <input type="text" class="form-control" id="Id" placeholder="Id générer par MySql" name="id" disabled>
              </div>
            </div>
            <div class="form-group">
              <label for="inputEmail" class="col-sm-3 control-label">Email</label>
              <div class="col-sm-9">
                <input type="mail" class="form-control" id="inputEmail" name="mail" placeholder="Email à enregistrer">
              </div>
            </div>
            <div class="form-group">
              <label for="inputRole" class="col-sm-3 control-label">Role</label>
              <div class="col-sm-9">
                <select class="form-control" name="role" id="inputRole">
                  <option value="admin">admin</option>
                  <option value="user" selected>user</option>
                </select>
              </div>
            </div>
            <div class="form-group">
              <label for="inputPassword" class="col-sm-3 control-label">Mot de Passe</label>
              <div class="col-sm-9">
                <input type="password" class="form-control" id="inputPassword" placeholder="Password" name="pass" disabled>
              </div>
            </div>
            <div class="form-group">
              <label for="inputCreate" class="col-sm-3 control-label">Date de Création</label>
              <div class="col-sm-9">
                <input type="datetime" class="form-control" id="inputCreate" placeholder="Générer à la création" name="created" disabled>
              </div>
            </div>

      </div>
      <div class="modal-footer">
        <button type="button" class="btn btn-default" data-dismiss="modal">Close</button>
        <button type="submit" class="btn btn-primary">Ajouter</button>
      </div>
    </div>
  </div>
</form>
</div>
<!-- modal end -->
