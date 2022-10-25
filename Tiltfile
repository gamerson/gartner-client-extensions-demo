# for more docs how to use this see: https://docs.tilt.dev/buttons.html
load("ext://uibutton", "cmd_button", "text_input", "location")


def after_all():
    add_drop_db_button()


def add_drop_db_button():
    cmd_button(
        "Kill DXP && Drop Docker Volume!",
        argv=[
            "sh",
            "-c",
            "docker container rm -f dxp-server && docker volume rm $VOLUME",
        ],
        resource='dxp.lfr.dev',
        icon_name="delete",
        text="Kill DXP && Drop Docker Volume!",
        inputs=[
            text_input("VOLUME"),
        ],
    )
