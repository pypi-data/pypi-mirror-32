"""Const helpers of Epson projector module."""

HTTP_OK = 200

ACCEPT_ENCODING = "gzip, deflate"
ACCEPT_HEADER = "application/json, text/javascript"

POWER = "PWR"
CMODE = "CMODE"
SOURCE = "SOURCE"
VOLUME = "VOLUME"
MUTE = "MUTE"
VOL_UP = "VOL_UP"
VOL_DOWN = "VOL_DOWN"
PLAY = "PLAY"
PAUSE = "PAUSE"
FAST = "FAST"
BACK = "BACK"
TURN_ON = "TURN_ON"
TURN_OFF = "TURN_OFF"
ALL = "ALL"

EPSON_CODES = {
    'PWR': '01'
}

EPSON_KEY_COMMANDS = {
    "TURN_ON": [('KEY', '3B')],
    "TURN_OFF": [('KEY', '3B'), ('KEY', '3B')],
    "HDMILINK": [('jsoncallback', 'HDMILINK?')],
    "PWR": [('jsoncallback', 'PWR?')],
    "SOURCE": [('jsoncallback', 'SOURCE?')],
    "CMODE": [('jsoncallback', 'CMODE?')],
    "VOLUME": [('jsoncallback', 'VOL?')],
    "CMODE_AUTO": [('CMODE', '00')],
    "CMODE_CINEMA": [('CMODE', '15')],
    "CMODE_NATURAL": [('CMODE', '07')],
    "CMODE_BRIGHT": [('CMODE', '0C')],
    "CMODE_DYNAMIC": [('CMODE', '06')],
    "CMODE_3DDYNAMIC": [('CMODE', '18')],
    "CMODE_3DCINEMA": [('CMODE', '17')],
    "CMODE_3DTHX": [('CMODE', '19')],
    "CMODE_BWCINEMA": [('CMODE', '20')],
    "CMODE_ARGB": [('CMODE', '21')],
    "CMODE_DCINEMA": [('CMODE', '22')],
    "CMODE_THX": [('CMODE', '13')],
    "CMODE_GAME": [('CMODE', '0D')],
    "CMODE_STAGE": [('CMODE', '16')],
    "CMODE_AUTOCOLOR": [('CMODE', 'C1')],
    "CMODE_XV": [('CMODE', '0B')],
    "CMODE_THEATRE": [('CMODE', '05')],
    "CMODE_THEATREBLACK": [('CMODE', '09')],
    "CMODE_THEATREBLACK2": [('CMODE', '0A')],
    "VOL_UP": [('KEY', '56')],
    "VOL_DOWN": [('KEY', '57')],
    "MUTE": [('KEY', 'D8')],
    "HDMI1": [('KEY', '4D')],
    "HDMI2": [('KEY', '40')],
    "PC": [('KEY', '44')],
    "VIDEO": [('KEY', '46')],
    "USB": [('KEY', '85')],
    "LAN": [('KEY', '53')],
    "WFD": [('KEY', '56')],
    "PLAY": [('KEY', 'D1')],
    "PAUSE": [('KEY', 'D3')],
    "STOP": [('KEY', 'D2')],
    "BACK": [('KEY', 'D4')],
    "FAST": [('KEY', 'D5')],
}

TIMEOUT_TIMES = {
    'TURN_ON': 40,
    'TURN_OFF': 10,
    'SOURCE': 5,
    'ALL': 3
}

DEFAULT_SOURCES = {
    'HDMI1': 'HDMI1',
    'HDMI2': 'HDMI2',
    'PC': 'PC',
    'VIDEO': 'VIDEO',
    'USB': 'USB',
    'LAN': 'LAN',
    'WFD': 'WiFi Direct'
}

SOURCE_LIST = {
    '30': 'HDMI1',
    '10': 'PC',
    '40': 'VIDEO',
    '52': 'USB',
    '53': 'LAN',
    '56': 'WDF',
    'A0': 'HDMI2'
}

INV_SOURCES = {v: k for k, v in DEFAULT_SOURCES.items()}

CMODE_LIST = {
    '00': 'Auto',
    '15': 'Cinema',
    '07': 'Natural',
    '0C': 'Bright Cinema/Living',
    '06': 'Dynamic',
    '17': '3D Cinema',
    '18': '3D Dynamic',
    '19': '3D THX',
    '20': 'B&W Cinema',
    '21': 'Adobe RGB',
    '22': 'Digital Cinema',
    '13': 'THX',
    '0D': 'Game',
    '16': 'Stage',
    'C1': 'AutoColor',
    '0B': 'x.v. color',
    '05': 'Theatre',
    '09': 'Theatre Black 1/HD',
    '0A': 'Theatre Black 2/Silver Screen'
}

CMODE_LIST_SET = {
    'cinema': 'CMODE_CINEMA',
    'natural': 'CMODE_NATURAL',
    'bright cinema': 'CMODE_BRIGHT',
    'dynamic': 'CMODE_DYNAMIC',
    '3ddynamic': 'CMODE_3DDYNAMIC',
    '3dcinema': 'CMODE_3DCINEMA',
    'auto': 'CMODE_AUTO',
    '3dthx': 'CMODE_3DTHX',
    'bwcinema': 'CMODE_BWCINEMA',
    'adobe rgb': 'CMODE_ARGB',
    'digital cinema': 'CMODE_DCINEMA',
    'thx': 'CMODE_THX',
    'game': 'CMODE_GAME',
    'stage': 'CMODE_STAGE',
    'autocolor': 'CMODE_AUTOCOLOR',
    'xv': 'CMODE_XV',
    'theatre': 'CMODE_THEATRE',
    'theatre black': 'CMODE_THEATREBLACK',
    'theatre black 2': 'CMODE_THEATREBLACK2'
}
