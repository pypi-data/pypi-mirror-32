from armulator.armv6.opcodes.abstract_opcode import AbstractOpcode
from armulator.armv6.shift import shift_c


class MvnRegisterShiftedRegister(AbstractOpcode):
    def __init__(self, setflags, m, s, d, shift_t):
        super(MvnRegisterShiftedRegister, self).__init__()
        self.setflags = setflags
        self.m = m
        self.s = s
        self.d = d
        self.shift_t = shift_t

    def execute(self, processor):
        shift_n = processor.registers.get(self.s)[24:32].uint
        shifted, carry = shift_c(processor.registers.get(self.m), self.shift_t, shift_n,
                                 processor.registers.cpsr.get_c())
        result = ~shifted
        processor.registers.set(self.d, result)
        if self.setflags:
            processor.registers.cpsr.set_n(result[0])
            processor.registers.cpsr.set_z(not result.any(True))
            processor.registers.cpsr.set_c(carry)
