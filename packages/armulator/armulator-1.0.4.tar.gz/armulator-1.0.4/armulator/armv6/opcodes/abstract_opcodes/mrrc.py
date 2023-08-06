from armulator.armv6.opcodes.abstract_opcode import AbstractOpcode


class Mrrc(AbstractOpcode):
    def __init__(self, cp, t, t2):
        super(Mrrc, self).__init__()
        self.cp = cp
        self.t = t
        self.t2 = t2

    def execute(self, processor):
        if processor.condition_passed():
            if not processor.coproc_accepted(self.cp, processor.this_instr()):
                processor.generate_coprocessor_exception()
            else:
                rt2, rt = processor.coproc_get_two_words(self.cp, processor.this_instr())
                processor.registers.set(self.t2, rt2)
                processor.registers.set(self.t, rt)
