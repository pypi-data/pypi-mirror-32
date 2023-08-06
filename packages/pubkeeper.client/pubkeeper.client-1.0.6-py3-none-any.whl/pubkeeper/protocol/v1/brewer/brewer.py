"""
    Copyright 2017 n.io Innovation, LLC | Patent Pending
"""
from pubkeeper.utils.crypto import PubCrypto
from pubkeeper.brewer.brewer import Brewer
from Crypto import Random
from Crypto.Cipher import AES
from binascii import hexlify, unhexlify
from uuid import uuid4
from pubkeeper.protocol.v1.frame import Frame
from tornado import ioloop


class ProtocolBrewer(Brewer):
    def __init__(self, *args, brewer_id=None, _ioloop=None, **kwargs):
        super().__init__(*args, **kwargs)
        self.brewer_id = brewer_id or uuid4().hex
        self.brewing = {}
        self.brews = []

        self._ioloop = _ioloop if _ioloop is not None else \
            ioloop.IOLoop.current()
        self._crypto = None
        if self.crypto:
            self._crypto = {
                'mode': AES.MODE_CBC,
                'key': hexlify(Random.new().read(16)).decode()
            }
            self._cipher = PubCrypto(
                self._crypto['mode']
            )

    def get_config(self):
        ret = {}

        if self.crypto:
            ret['cipher'] = self._crypto

        return ret

    def new_patrons(self, patrons):
        with self.topic_lock:
            for patron in patrons:
                brew = self.get_brew(patron['brew']['name'])

                if brew is None:
                    raise RuntimeError("Brewer could not match a parity "
                                       "brew for patron brew: {}".format(
                                           patron['brew']
                                       ))

                if brew in self.brewing and \
                   patron['patron_id'] in self.brewing[brew]:
                    self.remove_patron(patron['patron_id'])

                if brew not in self.brewing:
                    self.brewing[brew] = {
                        patron['patron_id']: patron['brew']
                    }
                else:
                    self.brewing[brew][patron['patron_id']] = patron['brew']

            for patron in patrons:
                brew.start_brewer(self.brewer_id,
                                  self.topic,
                                  patron['patron_id'],
                                  patron['brew'])

                self.logger.info("Started brewer for {}:{}".format(
                    self.topic, patron['patron_id']
                ))

    def remove_patron(self, patron_id):
        with self.topic_lock:
            for brew, patrons in self.brewing.copy().items():
                if patron_id in patrons:
                    brew.stop_brewer(self.brewer_id,
                                     self.topic,
                                     patron_id)

                    del(self.brewing[brew][patron_id])

                if len(self.brewing[brew]) == 0:
                    del(self.brewing[brew])

                self.logger.info("Stopped brewer for {}:{}".format(
                    self.topic, patron_id
                ))

    def brew(self, data):
        if self.crypto:
            data = self._cipher.encrypt(
                unhexlify(self._crypto['key']),
                data
            )

        self.logger.debug(
            "Brewer id: {} brewing on topic: {}, on {} brews".
            format(self.brewer_id, self.topic, len(self.brewing)))

        frame = Frame.pack(self.topic, self.brewer_id, data)
        with self.topic_lock:
            for brew, patrons in self.brewing.items():
                self._ioloop.add_callback(
                    brew.brew,
                    self.brewer_id,
                    self.topic,
                    frame,
                    patrons)
