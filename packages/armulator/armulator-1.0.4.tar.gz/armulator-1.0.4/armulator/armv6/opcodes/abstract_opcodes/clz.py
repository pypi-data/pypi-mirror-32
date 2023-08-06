from armulator.armv6.opcodes.abstract_opcode import AbstractOpcode
from bitstring import BitArray
from armulator.armv6.bits_ops import count_leading_zero_bits


class Clz(AbstractOpcode):
    def __init__(self, m, d):
        super(Clz, self).__init__()
        self.m = m
        self.d = d

    def execute(self, processor):
        if processor.condition_passed():
            processor.registers.set(self.d, BitArray(
                    uint=count_leading_zero_bits(processor.registers.get(self.m)), length=32))
