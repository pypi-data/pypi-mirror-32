from armulator.armv6.opcodes.abstract_opcode import AbstractOpcode
from bitstring import BitArray
from armulator.armv6.configurations import arch_version


class Mul(AbstractOpcode):
    def __init__(self, setflags, m, d, n):
        super(Mul, self).__init__()
        self.setflags = setflags
        self.m = m
        self.d = d
        self.n = n

    def execute(self, processor):
        if processor.condition_passed():
            operand1 = processor.registers.get(self.n).int
            operand2 = processor.registers.get(self.m).int
            result = operand1 * operand2
            f_result = BitArray(int=result, length=64)[32:]
            processor.registers.set(self.d, f_result)
            if self.setflags:
                processor.registers.cpsr.set_n(f_result[0])
                processor.registers.cpsr.set_z(not f_result.any(True))
                if arch_version() == 4:
                    processor.registers.cpsr.set_c(False)  # uknown
