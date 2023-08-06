from armulator.armv6.opcodes.abstract_opcodes.smlsd import Smlsd
from armulator.armv6.opcodes.opcode import Opcode


class SmlsdA1(Smlsd, Opcode):
    def __init__(self, instruction, m_swap, m, a, d, n):
        Opcode.__init__(self, instruction)
        Smlsd.__init__(self, m_swap, m, a, d, n)

    def is_pc_changing_opcode(self):
        return False

    @staticmethod
    def from_bitarray(instr, processor):
        rn = instr[28:32]
        m_swap = instr[26]
        rm = instr[20:24]
        ra = instr[16:20]
        rd = instr[12:16]
        if rd.uint == 15 or rn.uint == 15 or rm.uint == 15:
            print "unpredictable"
        else:
            return SmlsdA1(instr, **{"m_swap": m_swap, "m": rm.uint, "a": ra.uint, "d": rd.uint, "n": rn.uint})
