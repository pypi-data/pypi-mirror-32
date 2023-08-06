# Sensor API for the Raspberry Pi

The Sensor API for the Raspberry Pi is a web service that give access to environmental
information of a Raspberry Pi. It uses the sensors of the
[Sense Hat](https://www.raspberrypi.org/products/sense-hat/) as well as the
internal CPU thermometer.

It returns the information as a JSON structure:

```
{
  "C": 50.464, 
  "H": 39.45915603637695, 
  "P": 941.78369140625, 
  "T": 34.720176696777344
}
```

* C is the CPU Temperature (in °C)
* H is the Humidity (in %rH)
* P is the Barometric Pressure (in Millibars)
* T is the Ambient Temperature (in °C)


## Installation

Install the required software:
```
sudo apt update
sudo apt install sense-hat libatlas-base-dev libopenjp2-7-dev gunicorn3
```
Install the app:
```
sudo pip install heiafr-hydrocontest-sensor_api
```

## Running

If you want to test the service you can just start it localy using gunicorn:

```
gunicorn3 heiafr.hydrocontest.sensor_api:app
```

You can now go to http://127.0.0.1:8000 and check the result.

In a production environment, you would rather start the service using systemd and use a reverse proxy (e.g. Nginx) in front of it.

You can use this systemd service file (you can name it `heiafr-hydrocontest-sensor_api.service`):

```
[Unit]
Description=HEIAFR Hydrocontest Sensor API
After=network.target

[Service]
User=root
Group=root
ExecStart=gunicorn --bind=0.0.0.0:8081 heiafr.hydrocontest.sensor_api:app

[Install]
WantedBy=multi-user.target
```

This is a possible Nginx configuration:

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
