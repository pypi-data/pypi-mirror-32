from armulator.armv6.opcodes.abstract_opcodes.adr import Adr
from armulator.armv6.opcodes.opcode import Opcode
from armulator.armv6.shift import arm_expand_imm


class AdrA1(Adr, Opcode):
    def __init__(self, instruction, d, imm32):
        Opcode.__init__(self, instruction)
        Adr.__init__(self, True, d, imm32)

    def is_pc_changing_opcode(self):
        return self.d == 15

    @staticmethod
    def from_bitarray(instr, processor):
        imm12 = instr[20:32]
        imm32 = arm_expand_imm(imm12)
        rd = instr[16:20]
        return AdrA1(instr, **{"d": rd.uint, "imm32": imm32})
