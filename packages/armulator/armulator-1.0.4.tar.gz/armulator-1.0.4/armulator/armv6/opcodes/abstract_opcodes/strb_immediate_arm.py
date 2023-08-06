from armulator.armv6.opcodes.abstract_opcode import AbstractOpcode
from armulator.armv6.bits_ops import add as bits_add, sub as bits_sub
from bitstring import BitArray


class StrbImmediateArm(AbstractOpcode):
    def __init__(self, add, wback, index, t, n, imm32):
        super(StrbImmediateArm, self).__init__()
        self.add = add
        self.wback = wback
        self.index = index
        self.t = t
        self.n = n
        self.imm32 = imm32

    def execute(self, processor):
        if processor.condition_passed():
            offset_addr = bits_add(processor.registers.get(self.n), self.imm32, 32) if self.add else bits_sub(
                processor.registers.get(self.n), self.imm32, 32)
            address = offset_addr if self.index else processor.registers.get(self.n)
            processor.mem_u_set(address, 1, processor.registers.get(self.t)[24:32])
            if self.wback:
                processor.registers.set(self.n, offset_addr)

    def instruction_syndrome(self):
        if self.wback:
            return BitArray(length=9)
        else:
            return BitArray(bin="10000") + BitArray(uint=self.t, length=4)
