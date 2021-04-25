FROM python:3-slim-buster
RUN pip3 install python-telegram-bot --upgrade
RUN mkdir -p /opt/tttbot
WORKDIR /opt/tttbot
#RUN apt update
#RUN apt install -y net-tools
ADD ttt-bot.py /opt/tttbot
CMD [ "python3", "./ttt-bot.py" ]
