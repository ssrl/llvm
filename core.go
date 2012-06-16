package llvm

/**
  LLVM 3.1 Go Bindings
  (c) 2012 Chanwit Kaewkasi / SUT
  Inspired by on the Go-LLVM package of nsf
**/

/*
#cgo CFLAGS: -D__STDC_LIMIT_MACROS -D__STDC_CONSTANT_MACROS
#cgo LDFLAGS: -lLLVMCore
#include <llvm-c/Core.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

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

// LLVMTypeKind
const (
    VoidTypeKind      = C.LLVMVoidTypeKind
    HalfTypeKind      = C.LLVMHalfTypeKind
    FloatTypeKind     = C.LLVMFloatTypeKind
    DoubleTypeKind    = C.LLVMDoubleTypeKind
    X86_FP80TypeKind  = C.LLVMX86_FP80TypeKind
    FP128TypeKind     = C.LLVMFP128TypeKind
    PPC_FP128TypeKind = C.LLVMPPC_FP128TypeKind
    LabelTypeKind     = C.LLVMLabelTypeKind
    IntegerTypeKind   = C.LLVMIntegerTypeKind
    FunctionTypeKind  = C.LLVMFunctionTypeKind
    StructTypeKind    = C.LLVMStructTypeKind
    ArrayTypeKind     = C.LLVMArrayTypeKind
    PointerTypeKind   = C.LLVMPointerTypeKind
    VectorTypeKind    = C.LLVMVectorTypeKind
    MetadataTypeKind  = C.LLVMMetadataTypeKind
    X86_MMXTypeKind   = C.LLVMX86_MMXTypeKind
)

// LLVMLinkage
const (
    ExternalLinkage            = C.LLVMExternalLinkage
    AvailableExternallyLinkage = C.LLVMAvailableExternallyLinkage
    LinkOnceAnyLinkage         = C.LLVMLinkOnceAnyLinkage
    LinkOnceODRLinkage         = C.LLVMLinkOnceODRLinkage

    WeakAnyLinkage   = C.LLVMWeakAnyLinkage
    WeakODRLinkage   = C.LLVMWeakODRLinkage

    AppendingLinkage = C.LLVMAppendingLinkage
    InternalLinkage  = C.LLVMInternalLinkage

    PrivateLinkage                  = C.LLVMPrivateLinkage
    DLLImportLinkage                = C.LLVMDLLImportLinkage
    DLLExportLinkage                = C.LLVMDLLExportLinkage
    ExternalWeakLinkage             = C.LLVMExternalWeakLinkage
    GhostLinkage                    = C.LLVMGhostLinkage
    CommonLinkage                   = C.LLVMCommonLinkage
    LinkerPrivateLinkage            = C.LLVMLinkerPrivateLinkage
    LinkerPrivateWeakLinkage        = C.LLVMLinkerPrivateWeakLinkage
    LinkerPrivateWeakDefAutoLinkage = C.LLVMLinkerPrivateWeakDefAutoLinkage
)

// LLVMVisibility
const (
    DefaultVisibility   = C.LLVMDefaultVisibility
    HiddenVisibility    = C.LLVMHiddenVisibility
    ProtectedVisibility = C.LLVMProtectedVisibility
)

// LLVMCallConv
const (
    CCallConv           = C.LLVMCCallConv
    FastCallConv        = C.LLVMFastCallConv
    ColdCallConv        = C.LLVMColdCallConv
    X86StdcallCallConv  = C.LLVMX86StdcallCallConv
    X86FastcallCallConv = C.LLVMX86FastcallCallConv
)

// LLVMIntPredicate
const (
    IntEQ  = C.LLVMIntEQ
    IntNE  = C.LLVMIntNE
    IntUGT = C.LLVMIntUGT
    IntUGE = C.LLVMIntUGE
    IntULT = C.LLVMIntULT
    IntULE = C.LLVMIntULE
    IntSGT = C.LLVMIntSGT
    IntSGE = C.LLVMIntSGE
    IntSLT = C.LLVMIntSLT
    IntSLE = C.LLVMIntSLE
)

// LLVMRealPredicate
const (
    RealPredicateFalse  = C.LLVMRealPredicateFalse
    RealOEQ             = C.LLVMRealOEQ
    RealOGT             = C.LLVMRealOGT
    RealOGE             = C.LLVMRealOGE
    RealOLT             = C.LLVMRealOLT
    RealOLE             = C.LLVMRealOLE
    RealONE             = C.LLVMRealONE
    RealORD             = C.LLVMRealORD
    RealUNO             = C.LLVMRealUNO
    RealUEQ             = C.LLVMRealUEQ
    RealUGT             = C.LLVMRealUGT
    RealUGE             = C.LLVMRealUGE
    RealULT             = C.LLVMRealULT
    RealULE             = C.LLVMRealULE
    RealUNE             = C.LLVMRealUNE
    RealPredicateTrue   = C.LLVMRealPredicateTrue
)

// LLVMLandingPadClauseTy
const (
    LLVMLandingPadCatch  = C.LLVMLandingPadCatch
    LLVMLandingPadFilter = C.LLVMLandingPadFilter
)



func InitializeCore(registry PassRegistry) {
    C.LLVMInitializeCore(registry.C)
}

//
// Error Handling
//

func DisposeMessage(message string) {
    cmessage := C.CString(message)
    C.LLVMDisposeMessage(cmessage)
    C.free(unsafe.Pointer(cmessage))
}


//
// llvm.Context
//

func NewContext()       Context { return Context{C.LLVMContextCreate()}    }
func GetGlobalContext() Context { return Context{C.LLVMGetGlobalContext()} }

func (c Context) Dispose() { C.LLVMContextDispose(c.C) }

func (c Context) GetMDKindIDInContext(name string) (id int) {
    cname := C.CString(name)
    id := C.LLVMGetMDKindIDInContext(c.C, cname, C.unsigned(len(name)))
    C.free(unsafe.Pointer(cname))
    return
}

func GetMDKindID(name string) (id int) {
    cname := C.CString(name)
    C.LLVMGetMDKindID(cname, C.unsigned(len(name)))
    C.free(unsafe.Pointer(cname))
    return
}

//
// llvm.Module
//

func NewModule(name string) (m Module) {
    cname := C.CString(name)
    m := Module{C.LLVMModuleCreateWithName(cname)}
    C.free(unsafe.Pointer(cname))
    return
}

func NewModuleInContext(name string, context Context) (m Module) {
    cname := C.CString(name)
    m := Module{C.LLVMModuleCreateWithNameInContext(cname, context.C)}
    C.free(unsafe.Pointer(cname))
    return
}

func (m Module) Dispose() {
    C.LLVMDisposeModule(m.C)
}

func (m Module) GetDataLayout() string {
    cdataLayout := C.LLVMGetDataLayout(m.C)
    return C.GoString(cdataLayout)
}

func (m Module) SetDataLayout(triple string) {
    ctriple := C.CString(triple)
    C.LLVMSetDataLayout(m.C, ctriple)
    C.free(unsafe.Pointer(ctriple))
}

func (m Module) GetTargetTriple() string {
    ctargetTriple := C.LLVMGetTarget(m.C)
    return C.GoString(ctargetTriple)
}

func (m Module) SetTargetTriple(triple string) {
    ctriple := C.CString(triple)
    C.LLVMSetTarget(m.C, ctriple)
    C.free(unsafe.Pointer(ctriple))
}