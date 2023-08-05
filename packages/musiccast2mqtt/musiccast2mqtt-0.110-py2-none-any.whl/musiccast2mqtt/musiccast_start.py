'''Launcher for the MusicCast gateway.'''

import os.path

_APP_NAME = 'musiccast2mqtt'

import mqttgateway.utils.app_properties as app
app.Properties.init(app_path=os.path.realpath(__file__), app_name=_APP_NAME)

# import the module that initiates and starts the gateway
import mqttgateway.gateway.start_gateway as start_g

# import the module representing the interface *** add your import here ***
import musiccast2mqtt.musiccast_interface as mci

def main():
    # launch the gateway *** add your class here ***
    start_g.startgateway(mci.musiccastInterface)

if __name__ == '__main__':
    main()