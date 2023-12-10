iex (New-Object System.Net.WebClient).DownloadString('https://storage.yandexcloud.net/yandexcloud-yc/install.ps1')
curl -X POST `
 -H "Authorization: Bearer $(yc iam create-token)" `
 -H "Content-Type: application/json" `
 -H "X-Cloud-Org-ID: bpf4u818vq2a3k8iqlte" `
 --data '{"summary":"Упал билд в TeamCity", "queue":"TCBUILDFAILS", "type":"bug", "assignee":"8000000000000004"}' `
 https://api.tracker.yandex.net/v2/issues/
