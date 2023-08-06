from armulator.armv6.opcodes.abstract_opcode import AbstractOpcode


class Rev16(AbstractOpcode):
    def __init__(self, m, d):
        super(Rev16, self).__init__()
        self.m = m
        self.d = d

    def execute(self, processor):
        if processor.condition_passed():
            result = processor.registers.get(self.m)[8:16]
            result += processor.registers.get(self.m)[0:8]
            result += processor.registers.get(self.m)[24:32]
            result += processor.registers.get(self.m)[16:24]
            processor.registers.set(self.d, result)
