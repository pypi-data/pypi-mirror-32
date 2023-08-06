from armulator.armv6.opcodes.abstract_opcodes.uxth import Uxth
from armulator.armv6.opcodes.opcode import Opcode


class UxthT1(Uxth, Opcode):
    def __init__(self, instruction, m, d, rotation):
        Opcode.__init__(self, instruction)
        Uxth.__init__(self, m, d, rotation)

    def is_pc_changing_opcode(self):
        return False

    @staticmethod
    def from_bitarray(instr, processor):
        rd = instr[13:16]
        rm = instr[10:13]
        rotation = 0
        return UxthT1(instr, **{"m": rm.uint, "d": rd.uint, "rotation": rotation})
