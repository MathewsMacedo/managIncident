<h2>{{ .title2 }}</h2>
<div class="row">

  <table class="table table-bordered table-hover table-responsive table-user">
    <thead>
    <th class="text-center">Id</th>
    <th class="text-center">Mail</th>
    <th class="text-center">IP</th>
    <th class="text-center">Date Création</th>
    <th class="text-center">Supprimer / Modifier</th>
  </thead>
    {{ range $r := .demand }}
    <tr class="admin-user vertical-align">
      <td>{{ $r.Id }}</td>
      <td>{{ $r.Mail }}</td>
      <td>{{ $r.IP }}</td>
      <td>{{ date $r.Created "d-m-Y H:i"}}</td>
      <td class="text-center">
        <a href="#myModal-register-{{ $r.Id }}" data-toggle="modal"><span class="glyphicon glyphicon-pencil glyphicon-lg"></span></a>
        &nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
        <a href="/incident-manager/admin/register/delete/{{ $r.Id }}" title="Delete"><span class="glyphicon glyphicon-minus glyphicon-lg"></span></a>
        <div id="myModal-register-{{ $r.Id }}" class="modal fade">
          <form class="form-horizontal" action="/incident-manager/admin/user/add" method="post">
            <div class="modal-dialog">
                <div class="modal-content">
                    <div class="modal-header">
                        <button type="button" class="close" data-dismiss="modal" aria-hidden="true">&times;</button>
                        <h4 class="modal-title">Ajout de l'utilisateur : {{ $r.Mail }} </h4>
                    </div>
                    <div class="modal-body">
                          <div class="form-group">
                            <label for="inputEmail" class="col-sm-3 control-label">ID</label>
                            <div class="col-sm-9">
                              <input type="mail" class="form-control" id="inputEmail" value="{{ $r.Id }}" name="register_id">
                            </div>
                          </div>
                          <div class="form-group">
                            <label for="inputEmail" class="col-sm-3 control-label">Email</label>
                            <div class="col-sm-9">
                              <input type="mail" class="form-control" id="inputEmail" value="{{ $r.Mail }}" name="mail">
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
                            <label for="inputCreate" class="col-sm-3 control-label">IP lors de la demande</label>
                            <div class="col-sm-9">
                              <input type="mail" class="form-control" id="inputCreate" value="{{ $r.IP }}" disabled>
                            </div>
                          </div>
                          <div class="form-group">
                            <label for="inputCreate" class="col-sm-3 control-label">Date de Création</label>
                            <div class="col-sm-9">
                              <input type="mail" class="form-control" id="inputCreate" value="{{ $r.Created }}" disabled>
                            </div>
                          </div>

                    </div>
                    <div class="modal-footer">
                        <button type="button" class="btn btn-default" data-dismiss="modal">Fermer</button>
                        <button type="submit" class="btn btn-primary">Enregistrer les changements</button>
                    </div>
                </div>
            </div>
          </form>
        </div>
      </td>
   </tr>



{{ end }}
</table>

</div>
