from armulator.armv6.opcodes.abstract_opcode import AbstractOpcode
from armulator.armv6.bits_ops import add, sub
from bitstring import BitArray
from armulator.armv6.arm_exceptions import EndOfInstruction


class Ldmdb(AbstractOpcode):
    def __init__(self, wback, registers, n):
        super(Ldmdb, self).__init__()
        self.wback = wback
        self.registers = registers
        self.n = n

    def execute(self, processor):
        if processor.condition_passed():
            try:
                processor.null_check_if_thumbee(self.n)
            except EndOfInstruction:
                pass
            else:
                address = sub(processor.registers.get(self.n),
                              BitArray(uint=(4 * self.registers.count(1)), length=32), 32)
                for i in xrange(15):
                    if self.registers[15 - i]:
                        processor.registers.set(i, processor.mem_a_get(address, 4))
                        address = add(address, BitArray(bin="100"), 32)
                if self.registers[0]:
                    processor.load_write_pc(processor.mem_a_get(address, 4))
                if self.wback and not self.registers[15 - self.n]:
                    processor.registers.set(
                        self.n,
                        sub(
                            processor.registers.get(self.n), BitArray(uint=(4 * self.registers.count(1)), length=32), 32
                        )
                    )
                if self.wback and self.registers[15 - self.n]:
                    processor.registers.set(self.n, BitArray(length=32))  # unknown
