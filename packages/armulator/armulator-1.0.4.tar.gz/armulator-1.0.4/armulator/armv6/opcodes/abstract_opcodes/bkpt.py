from armulator.armv6.opcodes.abstract_opcode import AbstractOpcode


class Bkpt(AbstractOpcode):
    def __init__(self):
        super(Bkpt, self).__init__()

    def execute(self, processor):
        processor.bkpt_instr_debug_event()
