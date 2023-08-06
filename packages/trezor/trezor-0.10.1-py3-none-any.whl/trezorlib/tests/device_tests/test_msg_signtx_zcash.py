# This file is part of the TREZOR project.
#
# Copyright (C) 2012-2016 Marek Palatinus <slush@satoshilabs.com>
# Copyright (C) 2012-2016 Pavol Rusnak <stick@satoshilabs.com>
#
# This library is free software: you can redistribute it and/or modify
# it under the terms of the GNU Lesser General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# This library is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU Lesser General Public License for more details.
#
# You should have received a copy of the GNU Lesser General Public License
# along with this library.  If not, see <http://www.gnu.org/licenses/>.
from binascii import unhexlify, hexlify
import pytest

from .common import TrezorTest

from trezorlib import coins
from trezorlib import messages as proto
from trezorlib.tools import parse_path

TxApiZcashTestnet = coins.tx_api['Zcash Testnet']

TXHASH_aaf51e = unhexlify('aaf51e4606c264e47e5c42c958fe4cf1539c5172684721e38e69f4ef634d75dc')


@pytest.mark.zcash
class TestMsgSigntxZcash(TrezorTest):

    def test_one_one_fee(self):
        self.setup_mnemonic_allallall()

        # tx: aaf51e4606c264e47e5c42c958fe4cf1539c5172684721e38e69f4ef634d75dc
        # input 1: 3.0 TAZ

        inp1 = proto.TxInputType(
            address_n=parse_path("m/Zcash Testnet/0h/0/0"),  # tmQoJ3PTXgQLaRRZZYT6xk8XtjRbr2kCqwu
            amount=300000000,
            prev_hash=TXHASH_aaf51e,
            prev_index=1,
        )

        out1 = proto.TxOutputType(
            address='tmJ1xYxP8XNTtCoDgvdmQPSrxh5qZJgy65Z',
            amount=300000000 - 1940,
            script_type=proto.OutputScriptType.PAYTOADDRESS,
        )

        with self.client:
            self.client.set_tx_api(TxApiZcashTestnet)
            self.client.set_expected_responses([
                proto.TxRequest(request_type=proto.RequestType.TXINPUT, details=proto.TxRequestDetailsType(request_index=0)),
                proto.TxRequest(request_type=proto.RequestType.TXOUTPUT, details=proto.TxRequestDetailsType(request_index=0)),
                proto.ButtonRequest(code=proto.ButtonRequestType.ConfirmOutput),
                proto.ButtonRequest(code=proto.ButtonRequestType.SignTx),
                proto.TxRequest(request_type=proto.RequestType.TXINPUT, details=proto.TxRequestDetailsType(request_index=0)),
                proto.TxRequest(request_type=proto.RequestType.TXOUTPUT, details=proto.TxRequestDetailsType(request_index=0)),
                proto.TxRequest(request_type=proto.RequestType.TXFINISHED),
            ])

            (signatures, serialized_tx) = self.client.sign_tx('Zcash Testnet', [inp1, ], [out1, ], version=3, overwintered=True)

        # Accepted by network: tx TODO
        assert hexlify(serialized_tx) == b'030000807082c40301dc754d63eff4698ee321476872519c53f14cfe58c9425c7ee464c206461ef5aa010000006a473044022036556e76f99d0456eb893c5e8cbcae7fd615210c0b998c77de23e685c3064e8102201c90df2ecb753604e6ae805f225550bd48e730c9743ef05a5e29c945a5b2f9b60121030e669acac1f280d1ddf441cd2ba5e97417bf2689e4bbec86df4f831bf9f7ffd0ffffffff016c9be111000000001976a9145b157a678a10021243307e4bb58f36375aa80e1088ac000000000000000000'
