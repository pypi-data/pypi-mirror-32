from armulator.armv6.opcodes.abstract_opcodes.sub_register import SubRegister
from armulator.armv6.opcodes.opcode import Opcode
from armulator.armv6.shift import SRType


class SubRegisterT1(SubRegister, Opcode):
    def __init__(self, instruction, setflags, m, d, n, shift_t, shift_n):
        Opcode.__init__(self, instruction)
        SubRegister.__init__(self, setflags, m, d, n, shift_t, shift_n)

    def is_pc_changing_opcode(self):
        return self.d == 15

    @staticmethod
    def from_bitarray(instr, processor):
        rd = instr[13:16]
        rn = instr[10:13]
        rm = instr[7:10]
        set_flags = not processor.in_it_block()
        shift_t = SRType.SRType_LSL
        shift_n = 0
        return SubRegisterT1(instr, **{"setflags": set_flags, "m": rm.uint, "d": rd.uint, "n": rn.uint,
                                       "shift_t": shift_t, "shift_n": shift_n})
