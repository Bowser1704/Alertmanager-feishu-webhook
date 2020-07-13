{{ $var := .externalURL}}{{ range $k,$v:=.alerts }}
[Prometheus告警信息]({{$v.generatorURL}})
[{{$v.labels.alertname}}]({{$var}})
告警级别：{{$v.labels.level}}
开始时间：{{$v.startsAt}}
结束时间：{{$v.endsAt}}
故障主机IP：{{$v.labels.instance}}
{{$v.annotations.description}}
{{ end }}