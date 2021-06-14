# parrot-nag-bot

A Telegram bot (and Webserver) to feed an LED Matrix style display with messages.                                                   
## Features
* `/start` - Replies a greeting to show that he's ready to operate
* `/help` - Display a help message with all commands and their usage
* `/system` -  Print status message with information about how the bot is deployed

## Usage

### Plain Docker
Simply run:
```
docker run -d --restart=always -v  config/parrot-nag-bot_config.yaml:/etc/parrot-nag-bot/parrot-nag-bot_configy.yaml technat/parrot-bot:master
```

### Kubernetes
Apply the kustomize files in the [kustomize](./kustomize) folder:
```
kubectl apply -f kustomize/base
```

## License
This Project is developed under the MIT-License.

Avatar photo by <a href="https://unsplash.com/@kriztheman?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText">Christopher Alvarenga</a> on <a href="https://unsplash.com/s/photos/parrot?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText">Unsplash</a>
  
