from armulator.armv6.opcodes.abstract_opcode import AbstractOpcode


class It(AbstractOpcode):
    def __init__(self, firstcond, mask):
        super(It, self).__init__()
        self.firstcond = firstcond
        self.mask = mask

    def execute(self, processor):
        processor.registers.cpsr.set_it(self.firstcond + self.mask)
