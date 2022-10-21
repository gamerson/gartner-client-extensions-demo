# Demo Steps

## Install Prototype Liferay Client Extension CLI

1. Download and install go: <https://go.dev/doc/install>
2. Clone repo <https://github.com/liferay/liferay-cli>
3. Go into repo dir and execute `./gow install`
4. Test with command `liferay --version`

## Mac/Windows Install/Configure Docker Desktop

1. Install Docker desktop: <https://www.docker.com/products/docker-desktop/>
2. Tweak docker desktop settings
2a. Turn off kubernetes
2b. Turn off experimental features  
2c. Remove any extra registrties (kind-registry)
2d. Up the docker resources to 8 CPUs and 16GBs of RAM
2e. Make sure the folder that you plan to clone this repo into is listed in the "file sharing" list

## Bring up demo environment

1. Clone repo <https://github.com/gamerson/gartner-client-extensions-demo>
2. Execute `liferay runtime mkcert` in order to generate a wildcart certificate for `*.lfr.dev`
3. Execute `liferay runtime mkcert --install` in order to install the local Root Certificate Authority into your OS's and browser's truststores
4. Execute `liferay extension start -d /path/to/repos/gartner-client-extensions-demo`
If this is your first time to run this command, this first time will take several minutes as images are downloaded and built
5. Click the link that is shown in the terminal for Tilt UI

After several minutes both DXP and all extensions should be green in the Tilt UI

## Demo Walkthrough

1. Load DXP at <https://dxp.lfr.dev>
2. Login as test@dxp.lfr.dev/test
3. Go to Applications > Coupons
4. Add a new coupon
5. Go to Tilt UI and look at the 3 function extensions and see the JSON payload of the added coupon
6. Edit the `coupon-function-nodejs/app.js` file change the payload message and save
7. The function will be rebuilt and redeployed in a few seconds.
8. Edit the same coupon in the UI (change issued checkbox) and save
9. Notice nodejs function is called again and with updated message.

## Shutdown environment

1. If you plan to run demo again shortly, you can just bring down extensions environment `liferay ext stop`
2. If you want to free up all docker desktop resources, you can shutdonw whole runtime `liferay runtime delete`
3. Run `docker image prune` and `docker builder prune` to reclaim some space