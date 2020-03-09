HOW TO WORK LOCALLY: TO DO

HOW TO ACCESS THE DROPLET:
- "digo" terminal shortcut
- serverpass is the password
- nginx accessed here:
  - sudo nano /etc/nginx/sites-available/portfolio
- project accessed here


HOW TO START/STOP PROJECTS
- sudo systemctl start portfolio
- sudo systemctl enable jobs

Current resources
chi resources: https://github.com/go-chi/chi/blob/master/_examples/fileserver/main.go
nginx resources: https://www.digitalocean.com/community/tutorials/how-to-install-nginx-on-ubuntu-18-04
how to deploy golang: https://kenyaappexperts.com/blog/how-to-deploy-golang-to-production-step-by-step/

TODO:
- set up nginx to serve the static front end which requests data from the golang server
  - serve the react front end with nginx
  - set up a golang server that hosts some super basic api endpoints

