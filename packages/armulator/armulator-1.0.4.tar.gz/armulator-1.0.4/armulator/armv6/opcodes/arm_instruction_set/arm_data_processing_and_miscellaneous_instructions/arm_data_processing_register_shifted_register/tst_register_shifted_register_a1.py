from armulator.armv6.opcodes.abstract_opcodes.tst_register_shifted_register import TstRegisterShiftedRegister
from armulator.armv6.opcodes.opcode import Opcode
from armulator.armv6.shift import decode_reg_shift


class TstRegisterShiftedRegisterA1(TstRegisterShiftedRegister, Opcode):
    def __init__(self, instruction, m, s, n, shift_t):
        Opcode.__init__(self, instruction)
        TstRegisterShiftedRegister.__init__(self, m, s, n, shift_t)

    def is_pc_changing_opcode(self):
        return False

    @staticmethod
    def from_bitarray(instr, processor):
        rm = instr[-4:]
        type_o = instr[25:27]
        rs = instr[20:24]
        rn = instr[12:16]
        if rn == "0b1111" or rm == "0b1111" or rs == "0b1111":
            print "unpredictable"
        else:
            shift_t = decode_reg_shift(type_o)
            return TstRegisterShiftedRegisterA1(instr, **{"m": rm.uint, "s": rs.uint,
                                                          "n": rn.uint, "shift_t": shift_t})
