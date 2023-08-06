from armulator.armv6.opcodes.abstract_opcodes.adc_immediate import AdcImmediate
from armulator.armv6.opcodes.opcode import Opcode
from armulator.armv6.shift import arm_expand_imm


class AdcImmediateA1(AdcImmediate, Opcode):
    def __init__(self, instruction, setflags, d, n, imm32):
        Opcode.__init__(self, instruction)
        AdcImmediate.__init__(self, setflags, d, n, imm32)

    def is_pc_changing_opcode(self):
        return self.d == 15

    @staticmethod
    def from_bitarray(instr, processor):
        imm12 = instr[20:32]
        rd = instr[16:20]
        rn = instr[12:16]
        setflags = instr[11]
        imm32 = arm_expand_imm(imm12)
        return AdcImmediateA1(instr, **{"setflags": setflags, "d": rd.uint, "n": rn.uint, "imm32": imm32})
