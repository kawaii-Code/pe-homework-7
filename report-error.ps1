curl -X POST `
 -H "Authorization: Bearer t1.9euelZqQmpWWip2Vkp6QxpLGjsiLju3rnpWakZWMyJSTkJCdnp7JnMecnsjl9PdmeChU-e9YQFCl3fT3JicmVPnvWEBQpc3n9euelZrOyZmMx4nGj5GUlsiKkcrGie_8xeuelZrOyZmMx4nGj5GUlsiKkcrGiQ.iDBCZITFosdtxqTsLiW1h5g9qY8mKhvAgeZxQLQnnW85q-48zVt_sNjete7juM18qbvEb3UJlTk6_nRJ42bjDA" `
 -H "Content-Type: application/json" `
 -H "X-Cloud-Org-ID: bpf4u818vq2a3k8iqlte" `
 --data '{"summary":"Упал билд в TeamCity", "queue":"TCBUILDFAILS", "type":"bug", "assignee":"8000000000000004"}' `
 https://api.tracker.yandex.net/v2/issues/
