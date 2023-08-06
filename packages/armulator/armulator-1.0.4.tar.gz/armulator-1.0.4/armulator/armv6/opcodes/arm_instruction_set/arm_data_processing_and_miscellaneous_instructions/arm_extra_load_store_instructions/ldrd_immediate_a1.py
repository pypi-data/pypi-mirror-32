from armulator.armv6.opcodes.abstract_opcodes.ldrd_immediate import LdrdImmediate
from armulator.armv6.opcodes.opcode import Opcode


class LdrdImmediateA1(LdrdImmediate, Opcode):
    def __init__(self, instruction, add, wback, index, imm32, t, t2, n):
        Opcode.__init__(self, instruction)
        LdrdImmediate.__init__(self, add, wback, index, imm32, t, t2, n)

    def is_pc_changing_opcode(self):
        return False

    @staticmethod
    def from_bitarray(instr, processor):
        w = instr[10]
        index = instr[7]
        imm4_l = instr[-4:]
        imm4_h = instr[20:24]
        rt = instr[16:20]
        rn = instr[12:16]
        add = instr[8]
        imm32 = "0b000000000000000000000000" + imm4_h + imm4_l
        t2 = rt.uint + 1
        wback = (not index) or w
        if rt[3] or (not index and w) or (wback and (rn.uint == rt.uint or rn.uint == t2)) or t2 == 15:
            print "unpredictable"
        else:
            return LdrdImmediateA1(instr, **{"add": add, "wback": wback, "index": index, "imm32": imm32, "t": rt.uint,
                                             "t2": t2, "n": rn.uint})
