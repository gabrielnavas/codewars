import math

def distance_between_points(a, b):
    return math.sqrt(math.pow(b.x-a.x, 2) + math.pow(b.y-a.y, 2))