load("ext://uibutton", "cmd_button", "text_input", "location")

# this var will be used to customize the docker volume mapped to 
dxp_data_volume = "gartnerDemo"

# this map will add additional build args to the dxp-server base image build
more_dxp_buildargs = {
    "DXP_BASE_IMAGE": "liferay/dxp:7.4.13.nightly-d5.0.2-20221028072756"
}

# This map will add additional ENV VARs to the dxp-server docker container
#more_dxp_envs = {
#    "FOO": "BAR",
#    "FIZZ": "BUZZ",
#}

# api docs for adding UIButtons found here: https://docs.tilt.dev/buttons.html
cmd_button(
    "Kill DXP && Drop Docker Volume!",
    argv=[
        "sh",
        "-c",
        "docker container rm -f dxp-server && docker volume rm %s" % dxp_data_volume,
    ],
    resource="dxp.lfr.dev",
    icon_name="delete",
    text="Kill DXP && Drop Docker Volume!"
)