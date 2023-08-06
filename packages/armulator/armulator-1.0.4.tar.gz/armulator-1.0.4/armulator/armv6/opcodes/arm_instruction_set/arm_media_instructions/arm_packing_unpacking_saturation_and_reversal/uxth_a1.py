from armulator.armv6.opcodes.abstract_opcodes.uxth import Uxth
from armulator.armv6.opcodes.opcode import Opcode


class UxthA1(Uxth, Opcode):
    def __init__(self, instruction, m, d, rotation):
        Opcode.__init__(self, instruction)
        Uxth.__init__(self, m, d, rotation)

    def is_pc_changing_opcode(self):
        return False

    @staticmethod
    def from_bitarray(instr, processor):
        rm = instr[28:32]
        rotate = instr[20:22]
        rd = instr[16:20]
        rotation = rotate.uint * 8
        if rd.uint == 15 or rm.uint == 15:
            print "unpredictable"
        else:
            return UxthA1(instr, **{"m": rm.uint, "d": rd.uint, "rotation": rotation})
