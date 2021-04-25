# parrot-nag-bot

A Telegram bot (and Webserver) to feed an LED Matrix style display with messages.                                                   
## Features
* `/start` - Replies a greeting to show that he's ready to operate
* `/system` -  Print status message with information about how the bot is deployed

## Usage

### Plain Docker
Simply run:
```
docker run -d --restart=always -e TELEGRAM_TOKEN="yourSecureToken" technat/parrot-bot
```

### Kubernetes
Apply the kustomize files in the [kustomize](./kustomize) folder:
```
kubectl apply -f kustomize/base
```

## License
This Project is developed under the MIT-License.

Avatar Image photo by <a href="https://unsplash.com/@kriztheman?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText">Christopher Alvarenga</a> on <a href="https://unsplash.com/s/photos/parrot?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText">Unsplash</a>
  
