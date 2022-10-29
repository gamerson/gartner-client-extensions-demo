load("ext://uibutton", "cmd_button", "text_input", "location")
include('ext://cancel')

#dxp_data_volume = "liferayGartnerDemoData"

more_dxp_buildargs = {
    "DXP_BASE_IMAGE": "liferay/dxp:7.4.13.nightly-d5.0.2-20221028072756"
}

# This map will add additional ENV VARs to the dxp-server docker container
more_dxp_envs = {
    #"FOO": "BAR",
    #"FIZZ": "BUZZ",
}

# The localdev Tiltfile will look for a Tiltfile in the root of the client extensions directory
# if it exists, and there is a after_all() function it will be called after main Tiltfile is processed
def after_all():
    add_drop_db_button()


def add_drop_db_button():
    # api docs for adding UIButtons found here: https://docs.tilt.dev/buttons.html
    cmd_button(
        "Kill DXP && Drop Docker Volume!",
        argv=[
            "sh",
            "-c",
            "docker container rm -f dxp-server && docker volume rm $VOLUME",
        ],
        resource="dxp.lfr.dev",
        icon_name="delete",
        text="Kill DXP && Drop Docker Volume!",
        inputs=[
            text_input("VOLUME")#, default=dxp_data_volume)
        ],
    )
