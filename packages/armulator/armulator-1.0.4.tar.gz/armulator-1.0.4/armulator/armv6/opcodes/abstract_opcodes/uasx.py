from armulator.armv6.opcodes.abstract_opcode import AbstractOpcode
from bitstring import BitArray


class Uasx(AbstractOpcode):
    def __init__(self, m, d, n):
        super(Uasx, self).__init__()
        self.m = m
        self.d = d
        self.n = n

    def execute(self, processor):
        if processor.condition_passed():
            diff = processor.registers.get(self.n)[16:32].uint - processor.registers.get(self.m)[0:16].uint
            sum_ = processor.registers.get(self.n)[0:16].uint + processor.registers.get(self.m)[16:32].uint
            processor.registers.set(self.d, BitArray(int=sum_, length=16) + BitArray(int=diff, length=16))
            ge = "0b"
            ge += "11" if sum_ >= 0x10000 else "00"
            ge += "11" if diff >= 0 else "00"
            processor.registers.cpsr.set_ge(ge)
