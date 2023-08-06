from armulator.armv6.opcodes.abstract_opcodes.and_immediate import AndImmediate
from armulator.armv6.opcodes.opcode import Opcode
from armulator.armv6.shift import arm_expand_imm_c


class AndImmediateA1(AndImmediate, Opcode):
    def __init__(self, instruction, setflags, d, n, imm32, carry):
        Opcode.__init__(self, instruction)
        AndImmediate.__init__(self, setflags, d, n, imm32, carry)

    def is_pc_changing_opcode(self):
        return self.d == 15

    @staticmethod
    def from_bitarray(instr, processor):
        imm12 = instr[20:32]
        rd = instr[16:20]
        rn = instr[12:16]
        setflags = instr[11]
        imm32, carry = arm_expand_imm_c(imm12, processor.registers.cpsr.get_c())
        return AndImmediateA1(instr, **{"setflags": setflags, "d": rd.uint, "n": rn.uint, "imm32": imm32,
                                        "carry": carry})
