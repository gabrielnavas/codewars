def cockroach_speed(s):
    centrimeterPeerHour = s*100000
    centrimeterPeerSecond = centrimeterPeerHour/3600
    return int(centrimeterPeerSecond)