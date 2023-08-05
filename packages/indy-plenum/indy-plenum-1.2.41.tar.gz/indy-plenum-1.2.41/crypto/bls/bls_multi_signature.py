from collections import OrderedDict

from common.serializers.serialization import multi_signature_value_serializer


class MultiSignatureValue:
    """
    Data class for storing multi signature value,
    that is a value a BLS multi signature was calculated over.
    """

    # TODO: support ledger_size
    # It needs to be included into PRE-PREPARE
    def __init__(self,
                 ledger_id: int,
                 state_root_hash: list,
                 pool_state_root_hash: str,
                 txn_root_hash: str,
                 timestamp: int):
        '''
        :param ledger_id: id of the ledger multisig is created over
        :param state_root_hash: root hash (base58) of the state
        associated with ledger_id and multisig is created over
        :param pool_state_root_hash: root hash (base58) of pool state
        needed to anchor the state of Nodes participated in multisig
        :param txn_root_hash: root hash (base58) of the ledger associated
        with ledger_id and multisig is created over
        :param timestamp: timestamp of the state the multisig is created over
        '''
        assert ledger_id is not None
        assert state_root_hash is not None
        assert pool_state_root_hash is not None
        assert txn_root_hash is not None
        assert timestamp is not None
        self.ledger_id = ledger_id
        self.state_root_hash = state_root_hash
        self.pool_state_root_hash = pool_state_root_hash
        self.txn_root_hash = txn_root_hash
        self.timestamp = timestamp

    def as_single_value(self):
        return multi_signature_value_serializer.serialize(self.as_dict())

    def as_dict(self):
        return OrderedDict(sorted(self.__dict__.items()))

    def as_list(self):
        return [
            self.ledger_id,
            self.state_root_hash,
            self.pool_state_root_hash,
            self.txn_root_hash,
            self.timestamp,
        ]

    def __eq__(self, other):
        return isinstance(other, MultiSignatureValue) and self.as_dict() == other.as_dict()

    def __str__(self) -> str:
        return str(self.as_dict())


class MultiSignature:
    """
    Data class for storing multi signature and
    all data required for verification.
    """

    def __init__(self,
                 signature: str,
                 participants: list,
                 value: MultiSignatureValue):
        """
        :param signature: Multi signature itself
        :param participants: List of signers
        :param value: the value multi-signature was created over
        """
        assert signature is not None
        assert participants
        assert value is not None
        self.signature = signature
        self.participants = participants
        self.value = value

    @staticmethod
    def from_list(*args):
        value = MultiSignatureValue(*(args[2]))
        return MultiSignature(signature=args[0],
                              participants=args[1],
                              value=value)

    @staticmethod
    def from_dict(**kwargs):
        value = MultiSignatureValue(**(kwargs['value']))
        return MultiSignature(signature=kwargs['signature'],
                              participants=kwargs['participants'],
                              value=value)

    def as_dict(self):
        return {'signature': self.signature,
                'participants': self.participants,
                'value': self.value.as_dict()}

    def as_list(self):
        return [self.signature,
                self.participants,
                self.value.as_list()]

    def __eq__(self, other):
        return isinstance(other, MultiSignature) and self.as_dict() == other.as_dict()

    def __str__(self) -> str:
        return str(self.as_dict())
