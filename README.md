# Prometheus BigBlueButton exporter

Exports gauge roomname/participant count

Output: 

```
bbb_meetings{meeting="MeetingName"} 20 
```

## Docker

```bash
make docker

docker container run --rm -p 9688:9688 bigbluebutton-exporter --bbb.api=https://yoursite.com/bigbluebutton/api/ --bbb.secret=secret
```

Add `--debug` for more info

## Docker-compose

```yaml
  bbbexporter:
    image: bigbluebutton-exporter:latest
    restart: unless-stopped
    command: ["--bbb.api" ,"https://yoursite.com/bigbluebutton/api/", "--bbb.secret", "secret"]
    ports:
      - "9688:9688"
    networks:
      - monitor-net
```

## Credits

Took some code form https://github.com/MsloveDl/bbb4go