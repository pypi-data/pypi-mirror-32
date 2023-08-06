from armulator.armv6.opcodes.abstract_opcode import AbstractOpcode


class MovRegisterThumb(AbstractOpcode):
    def __init__(self, setflags, m, d):
        super(MovRegisterThumb, self).__init__()
        self.setflags = setflags
        self.m = m
        self.d = d

    def execute(self, processor):
        if processor.condition_passed():
            result = processor.registers.get(self.m)
            if self.d == 15:
                processor.alu_write_pc(result)
            else:
                processor.registers.set(self.d, result)
                if self.setflags:
                    processor.registers.cpsr.set_n(result[0])
                    processor.registers.cpsr.set_z(result.all(0))
