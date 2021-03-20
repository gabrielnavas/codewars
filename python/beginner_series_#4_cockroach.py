def cockroach_speed(s):
    centrimeterPeerHour = s*100000
    centrimeterPeerSecond = centrimeterPeerHour/3600
    return centrimeterPeerSecond

    
print(cockroach_speed(1.08))