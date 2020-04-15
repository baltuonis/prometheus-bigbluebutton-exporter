# Prometheus BigBlueButton exporter

Exports gauges for BigBlueButton meetings/participants/streams + recording 

Output:

```text
# HELP bbb_participants Gauge for participants in BigBlueButton meetings
# TYPE bbb_participants gauge
bbb_participants{connection="interactive",meeting="MeetingName"} 4
bbb_participants{connection="listener",meeting="MeetingName"} 1
# HELP bbb_recording Gauge if BigBlueButton meetings are recorded
# TYPE bbb_recording gauge
bbb_recording{meeting="MeetingName"} 0
# HELP bbb_streams Gauge for active streams in BigBlueButton meetings
# TYPE bbb_streams gauge
bbb_streams{media="audio",meeting="MeetingName"} 5
bbb_streams{media="video",meeting="MeetingName"} 4
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

Add `--debug` for more debug info.

Add `--privacy` to use InternalMeetingId instead of MeetingName (for privacy reasons).

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
