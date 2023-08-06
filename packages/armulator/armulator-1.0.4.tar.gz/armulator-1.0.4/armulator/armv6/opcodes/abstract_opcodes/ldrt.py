from armulator.armv6.opcodes.abstract_opcode import AbstractOpcode
from armulator.armv6.bits_ops import add as bits_add, sub as bits_sub
from bitstring import BitArray
from armulator.armv6.arm_exceptions import EndOfInstruction
from armulator.armv6.shift import shift, ror
from armulator.armv6.enums import InstrSet


class Ldrt(AbstractOpcode):
    def __init__(self, add, register_form, post_index, t, n, m="", shift_t="", shift_n="", imm32=""):
        super(Ldrt, self).__init__()
        self.add = add
        self.register_form = register_form
        self.post_index = post_index
        self.t = t
        self.n = n
        self.m = m
        self.shift_t = shift_t
        self.shift_n = shift_n
        self.imm32 = imm32

    def execute(self, processor):
        if processor.condition_passed():
            if processor.registers.current_mode_is_hyp():
                print "unpredictable"
            else:
                try:
                    processor.null_check_if_thumbee(self.n)
                except EndOfInstruction:
                    pass
                else:
                    offset = shift(processor.registers.get(self.m), self.shift_t, self.shift_n,
                                   processor.registers.cpsr.get_c()) if self.register_form else self.imm32
                    offset_addr = bits_add(processor.registers.get(self.n), offset, 32) if self.add else bits_sub(
                        processor.registers.get(self.n), offset, 32)
                    address = processor.registers.get(self.n) if self.post_index else offset_addr
                    data = processor.mem_u_unpriv_get(address, 4)
                    if self.post_index:
                        processor.registers.set(self.n, offset_addr)
                    if processor.unaligned_support() or address[30:32] == "0b00":
                        processor.registers.set(self.t, data)
                    else:
                        if processor.registers.current_instr_set() == InstrSet.InstrSet_ARM:
                            processor.registers.set(self.t, ror(data, 8 * address[30:32].uint))
                        else:
                            processor.registers.set(self.t, BitArray(length=32))  # unknown

    def instruction_syndrome(self):
        if self.t == 15:
            return BitArray(length=9)
        else:
            return BitArray(bin="11000") + BitArray(uint=self.t, length=4)
