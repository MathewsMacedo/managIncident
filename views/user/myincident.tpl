<h2>{{ .title2 }}</h2>
<div class="row">
  <table class="table table-bordered table-hover table-responsive">
    <thead>
    <th class="text-center">Utilisateurs</th>
    <th class="text-center">Catégories</th>
    <th class="text-center">Titre</th>
    <th class="text-center">Date Incident   </th>
    <th class="text-center">Date Résolu</th>
    <th class="text-center">Priority</th>
    <th class="text-center">Voir Détail</th>
  </thead>
    {{ range $incident := .incidents }}
    <tr class='
    {{ if compare $incident.DateResolution "0001-01-01 00:00:00 +0000 UTC" }}
      {{ if compare $incident.DateEstimated "0001-01-01 00:00:00 +0000 UTC" }}
        danger
      {{ else }}
      warning
    {{ end }}
    {{ else }}
    success
    {{ end }}'
    >
      <td>{{ $incident.User.Mail }}</td>
      <td>{{ $incident.Cat }}</td>
      <td>{{ $incident.Title }}</td>
      <td>{{ date $incident.DateRequest "d-m-Y H:i" }}</td>
      {{ if compare $incident.DateResolution "0001-01-01 00:00:00 +0000 UTC" }}
        {{ if compare $incident.DateEstimated "0001-01-01 00:00:00 +0000 UTC" }}
          <td></td>
          {{ else }}
          <td>En cours de résolution</td>
          {{ end }}
      {{ else }}
        <td>Résolu</td>
      {{ end }}
      {{ if compare $incident.Priority "4" }}
        <td>Majeure</td>
      {{ else if compare $incident.Priority "3" }}
        <td>Élevée</td>
      {{ else if compare $incident.Priority "2" }}
        <td>Normale</td>
      {{ else if compare $incident.Priority "1" }}
        <td>Basse</td>
      {{ else }}
      <td></td>
      {{ end }}

      <td class="text-center "><a href="/incident-manager/{{ if not_nil $.admin }}{{ $.admin }}{{ else }}user/{{ end }}incident/{{ $incident.Id }}" title="Voir"><span class="glyphicon glyphicon-eye-open"></span></a></td>
   </tr>
   {{ end }}
 </table>
</div>
