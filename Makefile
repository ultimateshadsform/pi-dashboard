all:
	docker compose build --no-cache
	rm -rf app.tar
	docker save -o app.tar dashboard-app
	scp -i ~/.ssh/pi app.tar alex@raspberrypi:~/docker-images
	# stop the service
	ssh alex@raspberrypi -i ~/.ssh/pi 'sudo systemctl stop dashboard'
	# Delete old container
	ssh alex@raspberrypi -i ~/.ssh/pi 'docker rm -f dashboard'
	# Delete old image
	ssh alex@raspberrypi -i ~/.ssh/pi 'docker rmi -f dashboard-app'
	# Load new image
	ssh alex@raspberrypi -i ~/.ssh/pi 'docker load -i ~/docker-images/app.tar'
	ssh alex@raspberrypi -i ~/.ssh/pi 'docker run -d --restart on-failure:5 --name dashboard -p 127.0.0.1:3000:3000 dashboard-app'
	ssh alex@raspberrypi -i ~/.ssh/pi 'sudo systemctl restart dashboard'
	ssh alex@raspberrypi -i ~/.ssh/pi 'killall chromium-browser'
	ssh alex@raspberrypi -i ~/.ssh/pi 'DISPLAY=:0 ./start_browser.sh &'