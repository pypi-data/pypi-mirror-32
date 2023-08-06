from armulator.armv6.opcodes.abstract_opcode import AbstractOpcode
from armulator.armv6.bits_ops import add as bits_add, sub as bits_sub, align, sign_extend
from armulator.armv6.arm_exceptions import EndOfInstruction
from bitstring import BitArray


class LdrsbLiteral(AbstractOpcode):
    def __init__(self, add, imm32, t):
        super(LdrsbLiteral, self).__init__()
        self.add = add
        self.imm32 = imm32
        self.t = t

    def execute(self, processor):
        if processor.condition_passed():
            try:
                processor.null_check_if_thumbee(15)
            except EndOfInstruction:
                pass
            else:
                base = align(processor.registers.get_pc(), 4)
                address = bits_add(base, self.imm32, 32) if self.add else bits_sub(base, self.imm32, 32)
                processor.registers.set(self.t, sign_extend(processor.mem_u_get(address, 1), 32))

    def instruction_syndrome(self):
        if self.t == 15:
            return BitArray(length=9)
        else:
            return BitArray(bin="10010") + BitArray(uint=self.t, length=4)
