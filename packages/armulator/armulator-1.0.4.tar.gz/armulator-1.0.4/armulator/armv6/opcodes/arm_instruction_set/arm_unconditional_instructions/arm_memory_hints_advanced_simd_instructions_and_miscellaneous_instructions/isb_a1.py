from armulator.armv6.opcodes.abstract_opcodes.isb import Isb
from armulator.armv6.opcodes.opcode import Opcode


class IsbA1(Isb, Opcode):
    def __init__(self, instruction):
        Opcode.__init__(self, instruction)
        Isb.__init__(self)

    def is_pc_changing_opcode(self):
        return False

    @staticmethod
    def from_bitarray(instr, processor):
        return IsbA1(instr)
