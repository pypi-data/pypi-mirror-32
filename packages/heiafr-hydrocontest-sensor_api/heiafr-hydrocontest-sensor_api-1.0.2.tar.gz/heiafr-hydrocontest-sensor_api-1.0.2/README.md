# Sensor API for the Raspberry Pi

## Installation:

Install the required software:
```
sudo apt update
sudo apt install libatlas-base-dev libopenjp2-7-dev gunicorn3
```

## Running

```
gunicorn --bind=0.0.0.0:8081 heiafr.hydrocontest.sensor_api:app
```

or start using sysdemd (e.g. hydro-sensors.service):

```
[Unit]
Description=Hydrocontest Sensors server
After=network.target

[Service]
User=root
Group=root
ExecStart=gunicorn --bind=0.0.0.0:8081 heiafr.hydrocontest.sensor_api:app

[Install]
WantedBy=multi-user.target
```

## Nginx Configuration

```
location /sensors {
    proxy_set_header    Host $host;
    proxy_set_header    X-Real-IP $remote_addr;
    proxy_set_header    X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header    X-Forwarded-Proto $scheme;
    proxy_pass          http://localhost:8081/;
    proxy_read_timeout  90;
}
```
