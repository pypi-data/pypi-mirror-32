from armulator.armv6.opcodes.abstract_opcodes.mul import Mul
from armulator.armv6.opcodes.opcode import Opcode


class MulT2(Mul, Opcode):
    def __init__(self, instruction, setflags, m, d, n):
        Opcode.__init__(self, instruction)
        Mul.__init__(self, setflags, m, d, n)

    def is_pc_changing_opcode(self):
        return False

    @staticmethod
    def from_bitarray(instr, processor):
        rm = instr[28:32]
        rd = instr[20:24]
        rn = instr[12:16]
        setflags = False
        if rd.uint in (13, 15) or rn.uint in (13, 15) or rm.uint in (13, 15):
            print "unpredictable"
        else:
            return MulT2(instr, **{"setflags": setflags, "m": rm.uint, "d": rd.uint, "n": rn.uint})
