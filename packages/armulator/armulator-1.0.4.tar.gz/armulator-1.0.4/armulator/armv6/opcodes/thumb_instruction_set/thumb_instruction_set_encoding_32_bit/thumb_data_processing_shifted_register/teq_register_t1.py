from armulator.armv6.opcodes.abstract_opcodes.teq_register import TeqRegister
from armulator.armv6.opcodes.opcode import Opcode
from armulator.armv6.shift import decode_imm_shift


class TeqRegisterT1(TeqRegister, Opcode):
    def __init__(self, instruction, m, n, shift_t, shift_n):
        Opcode.__init__(self, instruction)
        TeqRegister.__init__(self, m, n, shift_t, shift_n)

    def is_pc_changing_opcode(self):
        return False

    @staticmethod
    def from_bitarray(instr, processor):
        rm = instr[28:32]
        type_o = instr[26:28]
        imm2 = instr[24:26]
        imm3 = instr[17:20]
        rn = instr[12:16]
        shift_t, shift_n = decode_imm_shift(type_o, imm3 + imm2)
        if rn.uint in (13, 15) or rm.uint in (13, 15):
            print "unpredictable"
        else:
            return TeqRegisterT1(instr, **{"m": rm.uint, "n": rn.uint, "shift_t": shift_t, "shift_n": shift_n})
