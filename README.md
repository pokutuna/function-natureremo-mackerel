function-natureremo-mackerel
===

## Deploy

- Configure `token.yaml`
- Create a topic
  - `$ gcloud pubsub topics create remo-to-mackerel`
- Deploy a function subscribing the topic
  - `$ gcloud functions deploy RemoToMackerel --runtime go111 --trigger-topic=remo-to-mackerel --env-vars-file token.yaml`
- Create a scheduler
  - `$ gcloud scheduler jobs create pubsub kick-remo-to-mackerel --topic=remo-to-mackerel --schedule="every 5 mins" --message-body="{}"`
