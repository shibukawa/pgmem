package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecProcNodeFirst(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	F_check_stack_depth(m)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v8 == int32(0) {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v12 = v11
		} else {
			v12 = int32(679)
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v12
		v14 = m.T0[v12].(func(*base.Module, int32) int32)(m, l0)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			return v14
		}
	}
}
func F_ExecProcNodeInstr(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int64
	_ = v13
	var v18 int64
	_ = v18
	var v19 int64
	_ = v19
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v36 int64
	_ = v36
	var v39 int64
	_ = v39
	var v42 int64
	_ = v42
	var v45 int64
	_ = v45
	var v48 int64
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 float64
	_ = v61
	var v63 float64
	_ = v63
	var v64 float64
	_ = v64
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v84 int64
	_ = v84
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	if v10 == int32(1) {
		v13 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
		if v13 != int64(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v93 = m.ExcPending
			if v93 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_ExecProcNodeInstr_0), int32(0))
				mBase = m.M
				v97 = m.ExcPending
				if v97 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_ExecProcNodeInstr_1), int32(58), int32(_a_F_ExecProcNodeInstr_2))
					mBase = m.M
					v102 = m.ExcPending
					if v102 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			F___clock_gettime(m, int32(1), v7)
			mBase = m.M
			v18 = int64(*(*int32)(unsafe.Add(mBase, uint32(v7)+8)))
			v19 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
			*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v18 + v19*int64(1000000000)
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
			if v24 == int32(1) {
				base.MemoryCopy(m, v9+int32(16), int32(_a_F_ExecProcNodeInstr_3), int32(128))
			} else {
			}
			v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+2)))
			if v32 == int32(1) {
				v36 = *(*int64)(unsafe.Add(mBase, _c_F_ExecProcNodeInstr[0]))
				*(*int64)(unsafe.Add(mBase, uint32(v9)+176)) = v36
				v39 = *(*int64)(unsafe.Add(mBase, _c_F_ExecProcNodeInstr[1]))
				*(*int64)(unsafe.Add(mBase, uint32(v9)+168)) = v39
				v42 = *(*int64)(unsafe.Add(mBase, _c_F_ExecProcNodeInstr[2]))
				*(*int64)(unsafe.Add(mBase, uint32(v9)+160)) = v42
				v45 = *(*int64)(unsafe.Add(mBase, _c_F_ExecProcNodeInstr[3]))
				*(*int64)(unsafe.Add(mBase, uint32(v9)+152)) = v45
				v48 = *(*int64)(unsafe.Add(mBase, _c_F_ExecProcNodeInstr[4]))
				*(*int64)(unsafe.Add(mBase, uint32(v9)+144)) = v48
			} else {
			}
			v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v51 = m.T0[v50].(func(*base.Module, int32) int32)(m, l0)
			mBase = m.M
			v54 = m.ExcPending
			if v54 != 0 {
				return int32(0)
			} else {
				v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v51 != 0 {
					v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+4)))
					if v58&int32(2) != 0 {
						v61 = float64(0)
					} else {
						v61 = float64(1)
					}
					v63 = v61
				} else {
					v63 = float64(0)
				}
				v64 = *(*float64)(unsafe.Add(mBase, uint32(v55)+384))
				*(*float64)(unsafe.Add(mBase, uint32(v55)+384)) = base.F64_add(v63, v64)
				F_InstrStopCommon(m, v55, v55+int32(368))
				mBase = m.M
				v70 = m.ExcPending
				if v70 != 0 {
					return int32(0)
				} else {
					v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+361)))
					if v71 == int32(0) {
						v74 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v55)+361)) = uint8(v74)
						v84 = *(*int64)(unsafe.Add(mBase, uint32(v55)+368))
						*(*int64)(unsafe.Add(mBase, uint32(v55)+376)) = v84
					} else {
						v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+360)))
						if base.B2i32(base.F64_lt(v64, float64(1)) == int32(0))|base.B2i32(v80 != int32(1)) != 0 {
						} else {
							v84 = *(*int64)(unsafe.Add(mBase, uint32(v55)+368))
							*(*int64)(unsafe.Add(mBase, uint32(v55)+376)) = v84
						}
					}
					m.G0 = v7 + int32(16)
					return v51
				}
			}
		}
	} else {
		v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
		if v24 == int32(1) {
			base.MemoryCopy(m, v9+int32(16), int32(_a_F_ExecProcNodeInstr_3), int32(128))
		} else {
		}
		v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+2)))
		if v32 == int32(1) {
			v36 = *(*int64)(unsafe.Add(mBase, _c_F_ExecProcNodeInstr[0]))
			*(*int64)(unsafe.Add(mBase, uint32(v9)+176)) = v36
			v39 = *(*int64)(unsafe.Add(mBase, _c_F_ExecProcNodeInstr[1]))
			*(*int64)(unsafe.Add(mBase, uint32(v9)+168)) = v39
			v42 = *(*int64)(unsafe.Add(mBase, _c_F_ExecProcNodeInstr[2]))
			*(*int64)(unsafe.Add(mBase, uint32(v9)+160)) = v42
			v45 = *(*int64)(unsafe.Add(mBase, _c_F_ExecProcNodeInstr[3]))
			*(*int64)(unsafe.Add(mBase, uint32(v9)+152)) = v45
			v48 = *(*int64)(unsafe.Add(mBase, _c_F_ExecProcNodeInstr[4]))
			*(*int64)(unsafe.Add(mBase, uint32(v9)+144)) = v48
		} else {
		}
		v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v51 = m.T0[v50].(func(*base.Module, int32) int32)(m, l0)
		mBase = m.M
		v54 = m.ExcPending
		if v54 != 0 {
			return int32(0)
		} else {
			v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v51 != 0 {
				v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+4)))
				if v58&int32(2) != 0 {
					v61 = float64(0)
				} else {
					v61 = float64(1)
				}
				v63 = v61
			} else {
				v63 = float64(0)
			}
			v64 = *(*float64)(unsafe.Add(mBase, uint32(v55)+384))
			*(*float64)(unsafe.Add(mBase, uint32(v55)+384)) = base.F64_add(v63, v64)
			F_InstrStopCommon(m, v55, v55+int32(368))
			mBase = m.M
			v70 = m.ExcPending
			if v70 != 0 {
				return int32(0)
			} else {
				v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+361)))
				if v71 == int32(0) {
					v74 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v55)+361)) = uint8(v74)
					v84 = *(*int64)(unsafe.Add(mBase, uint32(v55)+368))
					*(*int64)(unsafe.Add(mBase, uint32(v55)+376)) = v84
				} else {
					v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+360)))
					if base.B2i32(base.F64_lt(v64, float64(1)) == int32(0))|base.B2i32(v80 != int32(1)) != 0 {
					} else {
						v84 = *(*int64)(unsafe.Add(mBase, uint32(v55)+368))
						*(*int64)(unsafe.Add(mBase, uint32(v55)+376)) = v84
					}
				}
				m.G0 = v7 + int32(16)
				return v51
			}
		}
	}
}
