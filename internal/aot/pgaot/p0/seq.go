package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__equalCreateSeqStmt(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	v3 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v6 = F_equal(m, v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 == int32(0) {
			v27 = v3
			return v27
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v14 = F_equal(m, v12, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				if v14 == int32(0) {
					v27 = v3
				} else {
					v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					if v18 != v19 {
						v27 = v3
					} else {
						v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
						v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
						if v21 != v22 {
							v27 = v3
						} else {
							v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+17)))
							v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+17)))
							v27 = base.B2i32(v24 == v25)
						}
					}
				}
				return v27
			}
		}
	}
}
func F_seq_redo(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int64
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+48)))
	v15 = v13 & int32(240)
	if v15 == int32(0) {
		v18 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+64))
		v21 = F_XLogInitBufferForRedo(m, l0, int32(0))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			if v21 < int32(0) {
				v26 = *(*int32)(unsafe.Add(mBase, _c_F_seq_redo[0]))
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v26+(v21^int32(-1))<<(uint(int32(2))%32))))
				v40 = v32
			} else {
				v34 = *(*int32)(unsafe.Add(mBase, _c_F_seq_redo[1]))
				v40 = v34 + v21<<(uint(int32(13))%32) + int32(-8192)
			}
			v42 = F_palloc(m, int32(_a_F_seq_redo_0))
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return
			} else {
				v44 = int32(_a_F_seq_redo_0)
				v46 = int32(0)
				if v46|(v42&int32(3)|int32(1)) == v46 {
					v62 = v42 + v44
					v64 = v42 + int32(4)
					if base.Ui32(v64) < base.Ui32(v62) {
						v66 = v62
					} else {
						v66 = v64
					}
					v71 = (v42^int32(-1)+v66)&int32(-4) + int32(4)
					if v71 == int32(0) {
					} else {
						base.MemoryFill(m, v42, int32(0), v71)
					}
				} else {
					base.MemoryFill(m, v42, int32(0), v44)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v42)+10)) = int32(_a_F_seq_redo_1)
				v85 = int32(_a_F_seq_redo_2)
				*(*uint16)(unsafe.Add(mBase, uint32(v42)+18)) = uint16(v85)
				v91 = int32(_a_F_seq_redo_3)
				*(*uint16)(unsafe.Add(mBase, uint32(v42)+16)) = uint16(v91)
				*(*uint16)(unsafe.Add(mBase, uint32(v42)+14)) = uint16(v91)
				v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42)+16)))
				*(*int32)(unsafe.Add(mBase, uint32(v42+v94))) = int32(_a_F_seq_redo_4)
				v98 = int32(12)
				v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
				v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+68))
				v106 = F_PageAddItemExtended(m, v42, v19+v98, v101-v98, int32(1), int32(0))
				mBase = m.M
				v107 = m.ExcPending
				if v107 != 0 {
					return
				} else {
					if v106 == int32(0) {
						F_errstart_cold(m, int32(23), int32(0))
						mBase = m.M
						v140 = m.ExcPending
						if v140 != 0 {
							return
						} else {
							F_errmsg_internal(m, int32(_a_F_seq_redo_5), int32(0))
							mBase = m.M
							v144 = m.ExcPending
							if v144 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_seq_redo_6), int32(1930), int32(_a_F_seq_redo_7))
								mBase = m.M
								v149 = m.ExcPending
								if v149 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v42))) = base.I64_rotr(v18, int64(32))
						base.MemoryCopy(m, v40, v42, int32(_a_F_seq_redo_0))
						F_MarkBufferDirty(m, v21)
						mBase = m.M
						v116 = m.ExcPending
						if v116 != 0 {
							return
						} else {
							F_UnlockReleaseBuffer(m, v21)
							mBase = m.M
							v118 = m.ExcPending
							if v118 != 0 {
								return
							} else {
								F_pfree(m, v42)
								mBase = m.M
								v120 = m.ExcPending
								if v120 != 0 {
									return
								} else {
									m.G0 = v10 + int32(16)
									return
								}
							}
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(23), int32(0))
		mBase = m.M
		v127 = m.ExcPending
		if v127 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v10))) = v15
			F_errmsg_internal(m, int32(_a_F_seq_redo_8), v10)
			mBase = m.M
			v131 = m.ExcPending
			if v131 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_seq_redo_6), int32(1905), int32(_a_F_seq_redo_7))
				mBase = m.M
				v136 = m.ExcPending
				if v136 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
