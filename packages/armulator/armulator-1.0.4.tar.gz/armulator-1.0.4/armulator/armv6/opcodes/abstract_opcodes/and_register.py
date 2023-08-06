from armulator.armv6.opcodes.abstract_opcode import AbstractOpcode
from armulator.armv6.shift import shift_c


class AndRegister(AbstractOpcode):
    def __init__(self, setflags, m, d, n, shift_t, shift_n):
        super(AndRegister, self).__init__()
        self.setflags = setflags
        self.m = m
        self.d = d
        self.n = n
        self.shift_t = shift_t
        self.shift_n = shift_n

    def execute(self, processor):
        if processor.condition_passed():
            shifted, carry = shift_c(
                    processor.registers.get(self.m), self.shift_t, self.shift_n,
                    processor.registers.cpsr.get_c()
            )
            result = processor.registers.get(self.n) & shifted
            if self.d == 15:
                processor.alu_write_pc(result)
            else:
                processor.registers.set(self.d, result)
                if self.setflags:
                    processor.registers.cpsr.set_n(result[0])
                    processor.registers.cpsr.set_z(not result.any(True))
                    processor.registers.cpsr.set_c(carry)
