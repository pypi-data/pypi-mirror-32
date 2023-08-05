# -*- coding: utf-8 -*-
#


class VolumeBase(object):
    _ID = 0
    dimension = 3

    def __init__(self, id0=None):
        if id0:
            self.id = id0
        else:
            self.id = 'vol{}'.format(VolumeBase._ID)
            VolumeBase._ID += 1
        return
