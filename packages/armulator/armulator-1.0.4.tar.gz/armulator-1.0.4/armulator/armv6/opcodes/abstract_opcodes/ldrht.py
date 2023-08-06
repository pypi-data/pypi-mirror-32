from armulator.armv6.opcodes.abstract_opcode import AbstractOpcode
from armulator.armv6.bits_ops import add as bits_add, sub as bits_sub, zero_extend
from bitstring import BitArray
from armulator.armv6.arm_exceptions import EndOfInstruction


class Ldrht(AbstractOpcode):
    def __init__(self, add, register_form, post_index, t, n, m="", imm32=""):
        super(Ldrht, self).__init__()
        self.add = add
        self.register_form = register_form
        self.post_index = post_index
        self.t = t
        self.n = n
        self.m = m
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
                    offset = processor.registers.get(self.m) if self.register_form else self.imm32
                    offset_addr = bits_add(processor.registers.get(self.n), offset, 32) if self.add else bits_sub(
                        processor.registers.get(self.n), offset, 32)
                    address = processor.registers.get(self.n) if self.post_index else offset_addr
                    data = processor.mem_u_unpriv_get(address, 2)
                    if self.post_index:
                        processor.registers.set(self.n, offset_addr)
                    if processor.unaligned_support() or not address[31]:
                        processor.registers.set(self.t, zero_extend(data, 32))
                    else:
                        processor.registers.set(self.t, BitArray(length=32))  # unknown

    def instruction_syndrome(self):
        if self.t == 15:
            return BitArray(length=9)
        else:
            return BitArray(bin="10100") + BitArray(uint=self.t, length=4)
