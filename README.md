# Prometheus BigBlueButton exporter

Exports gauge roomname/participant count

Output:

```text
bbb_meetings{meeting="MeetingName"} 20
```

## Linux service

1. `cp ./etc/systemd/system/bbb-exporter.service /etc/systemd/system/`
2. Replace BBB endpoint & secret
3. `sudo systemctl enable bbb-exporter.service`
4. `sudo systemctl start bbb-exporter.service`

## Docker

```bash
make docker

docker container run --rm -p 9688:9688 bigbluebutton-exporter --bbb.api=https://yoursite.com/bigbluebutton/api/ --bbb.secret=secret
```

Add `--debug` for more info

## Docker-compose

```yaml
  bbbexporter:
    image: baltuonis/prometheus-bigbluebutton-exporter
    restart: unless-stopped
    command: ["--bbb.api" ,"https://yoursite.com/bigbluebutton/api/", "--bbb.secret", "secret"]
    ports:
      - "9688:9688"
    networks:
      - monitor-net
```

## Todo

1. Cleanup
2. Setup CI/CD to dockerhub

## Credits

Took some code from https://github.com/MsloveDl/bbb4go
