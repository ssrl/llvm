package llvm

/*
#cgo CFLAGS: -D__STDC_LIMIT_MACROS -D__STDC_CONSTANT_MACROS
#include <llvm-c/Core.h>
#include <stdlib.h>
*/
import "C"

type (
    Context         struct { C C.LLVMContextRef        }
    Module          struct { C C.LLVMModuleRef         }
    Type            struct { C C.LLVMTypeRef           }
    Value           struct { C C.LLVMValueRef          }
    BasicBlock      struct { C C.LLVMBasicBlockRef     }
    Builder         struct { C C.LLVMBuilderRef        }
    ModuleProvider  struct { C C.LLVMModuleProviderRef }
    MemoryBuffer    struct { C C.LLVMMemoryBufferRef   }
    PassManager     struct { C C.LLVMPassManagerRef    } // or PassManagerBase ?
    PassRegistry    struct { C C.LLVMPassRegistryRef   }
    Use             struct { C C.LLVMUseRef            }
)

func (this Context)        IsNil() bool { return this.C == nil }
func (this Module)         IsNil() bool { return this.C == nil }
func (this Type)           IsNil() bool { return this.C == nil }
func (this Value)          IsNil() bool { return this.C == nil }
func (this BasicBlock)     IsNil() bool { return this.C == nil }
func (this Builder)        IsNil() bool { return this.C == nil }
func (this ModuleProvider) IsNil() bool { return this.C == nil }
func (this MemoryBuffer)   IsNil() bool { return this.C == nil }
func (this PassManager)    IsNil() bool { return this.C == nil }
func (this PassRegistry)   IsNil() bool { return this.C == nil }
func (this Use)            IsNil() bool { return this.C == nil }

// LLVMAttribute
const (
    NoneAttribute            = 0
    ZExtAttribute            = C.LLVMZExtAttribute
    SExtAttribute            = C.LLVMSExtAttribute
    NoReturnAttribute        = C.LLVMNoReturnAttribute
    InRegAttribute           = C.LLVMInRegAttribute
    StructRetAttribute       = C.LLVMStructRetAttribute
    NoUnwindAttribute        = C.LLVMNoUnwindAttribute
    NoAliasAttribute         = C.LLVMNoAliasAttribute
    ByValAttribute           = C.LLVMByValAttribute
    NestAttribute            = C.LLVMNestAttribute
    ReadNoneAttribute        = C.LLVMReadNoneAttribute
    ReadOnlyAttribute        = C.LLVMReadOnlyAttribute
    NoInlineAttribute        = C.LLVMNoInlineAttribute
    AlwaysInlineAttribute    = C.LLVMAlwaysInlineAttribute
    OptimizeForSizeAttribute = C.LLVMOptimizeForSizeAttribute
    StackProtectAttribute    = C.LLVMStackProtectAttribute
    StackProtectReqAttribute = C.LLVMStackProtectReqAttribute
    Alignment                = C.LLVMAlignment
    NoCaptureAttribute       = C.LLVMNoCaptureAttribute
    NoRedZoneAttribute       = C.LLVMNoRedZoneAttribute
    NoImplicitFloatAttribute = C.LLVMNoImplicitFloatAttribute
    NakedAttribute           = C.LLVMNakedAttribute
    InlineHintAttribute      = C.LLVMInlineHintAttribute
    StackAlignment           = C.LLVMStackAlignment
    ReturnsTwice             = C.LLVMReturnsTwice
    UWTable                  = C.LLVMUWTable
    NonLazyBind              = C.LLVMNonLazyBind
)

// LLVMOpcode
const (
    Ret        =  C.LLVMRet
    Br         =  C.LLVMBr
    Switch     =  C.LLVMSwitch
    IndirectBr =  C.LLVMIndirectBr
    Invoke     =  C.LLVMInvoke

    /* removed 6 due to API changes */
    Unreachable = C.LLVMUnreachable

    /* Standard Binary Operators */
    Add  = C.LLVMAdd
    FAdd = C.LLVMFAdd
    Sub  = C.LLVMSub
    FSub = C.LLVMFSub
    Mul  = C.LLVMMul
    FMul = C.LLVMFMul
    UDiv = C.LLVMUDiv
    SDiv = C.LLVMSDiv
    FDiv = C.LLVMFDiv
    URem = C.LLVMURem
    SRem = C.LLVMSRem
    FRem = C.LLVMFRem

    /* Logical Operators */
    Shl  = C.LLVMShl
    LShr = C.LLVMLShr
    AShr = C.LLVMAShr
    And  = C.LLVMAnd
    Or   = C.LLVMOr
    Xor  = C.LLVMXor

    /* Memory Operators */
    Alloca        = C.LLVMAlloca
    Load          = C.LLVMLoad
    Store         = C.LLVMStore
    GetElementPtr = C.LLVMGetElementPtr

    /* Cast Operators */
    Trunc    = C.LLVMTrunc
    ZExt     = C.LLVMZExt
    SExt     = C.LLVMSExt
    FPToUI   = C.LLVMFPToUI
    FPToSI   = C.LLVMFPToSI
    UIToFP   = C.LLVMUIToFP
    SIToFP   = C.LLVMSIToFP
    FPTrunc  = C.LLVMFPTrunc
    FPExt    = C.LLVMFPExt
    PtrToInt = C.LLVMPtrToInt
    IntToPtr = C.LLVMIntToPtr
    BitCast  = C.LLVMBitCast

    /* Other Operators */
    ICmp           = C.LLVMICmp
    FCmp           = C.LLVMFCmp
    PHI            = C.LLVMPHI
    Call           = C.LLVMCall
    Select         = C.LLVMSelect
    UserOp1        = C.LLVMUserOp1
    UserOp2        = C.LLVMUserOp2
    VAArg          = C.LLVMVAArg
    ExtractElement = C.LLVMExtractElement
    InsertElement  = C.LLVMInsertElement
    ShuffleVector  = C.LLVMShuffleVector
    ExtractValue   = C.LLVMExtractValue
    InsertValue    = C.LLVMInsertValue

    /* Atomic operators */
    Fence         = C.LLVMFence
    AtomicCmpXchg = C.LLVMAtomicCmpXchg
    AtomicRMW     = C.LLVMAtomicRMW

    /* Exception Handling Operators */
    Resume     = C.LLVMResume
    LandingPad = C.LLVMLandingPad
)
