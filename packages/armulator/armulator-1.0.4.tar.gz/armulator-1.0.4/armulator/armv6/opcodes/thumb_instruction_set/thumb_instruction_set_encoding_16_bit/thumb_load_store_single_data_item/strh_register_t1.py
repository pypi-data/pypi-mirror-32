from armulator.armv6.opcodes.abstract_opcodes.strh_register import StrhRegister
from armulator.armv6.opcodes.opcode import Opcode
from armulator.armv6.shift import SRType


class StrhRegisterT1(StrhRegister, Opcode):
    def __init__(self, instruction, add, wback, index, m, t, n, shift_t, shift_n):
        Opcode.__init__(self, instruction)
        StrhRegister.__init__(self, add, wback, index, m, t, n, shift_t, shift_n)

    def is_pc_changing_opcode(self):
        return False

    @staticmethod
    def from_bitarray(instr, processor):
        rt = instr[13:16]
        rn = instr[10:13]
        rm = instr[7:10]
        index = True
        add = True
        wback = False
        shift_t = SRType.SRType_LSL
        shift_n = 0
        return StrhRegisterT1(instr, **{"add": add, "wback": wback, "index": index, "m": rm.uint, "t": rt.uint,
                                        "n": rn.uint, "shift_t": shift_t, "shift_n": shift_n})
